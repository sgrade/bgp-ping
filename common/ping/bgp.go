package ping

import (
	"fmt"
	"net"
	"time"
)

// TCPing ...
type BgpPing struct {
	target   *Target
	progress chan struct{}
	done     chan struct{}
	result   *Result
}

var _ Pinger = (*BgpPing)(nil)

// NewBgpPing return a new BgpPing
func NewBgpPing() *BgpPing {
	bgpping := BgpPing{
		progress: make(chan struct{}),
		done:     make(chan struct{}),
	}
	return &bgpping
}

// SetTarget set target for BgpPing
func (bgpping *BgpPing) SetTarget(target *Target) {
	bgpping.target = target
	if bgpping.result == nil {
		bgpping.result = &Result{Target: target}
	}
}

// Report progress after each ping attempt
func (bgpping BgpPing) ReportProgress() {
	bgpping.progress <- struct{}{}
}

// Result return the result
func (bgpping BgpPing) Result() *Result {
	return bgpping.result
}

// Start a bgpping
func (bgpping BgpPing) Start() <-chan struct{} {
	go func() {
		// cheap workaround to not wait for Interval seconds to get the 1st ping: set short wait time and ...
		t := time.NewTicker(1)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				// ... change to real Interval after the 1st run
				t.Reset(bgpping.target.Interval)

				if bgpping.result.Counter >= bgpping.target.Counter && bgpping.target.Counter != 0 {
					bgpping.Stop()
					return
				}
				duration, remoteAddr, err := bgpping.ping()
				bgpping.result.Counter++

				if err != nil {
					fmt.Printf("Ping %s - failed: %s\n", bgpping.target, err)
				} else {
					fmt.Printf("Ping %s(%s) - Connected - time=%s\n", bgpping.target, remoteAddr, duration)

					if bgpping.result.MinDuration == 0 {
						bgpping.result.MinDuration = duration
					}
					if bgpping.result.MaxDuration == 0 {
						bgpping.result.MaxDuration = duration
					}
					bgpping.result.SuccessCounter++
					if duration > bgpping.result.MaxDuration {
						bgpping.result.MaxDuration = duration
					} else if duration < bgpping.result.MinDuration {
						bgpping.result.MinDuration = duration
					}
					bgpping.result.TotalDuration += duration
				}

				bgpping.ReportProgress()

			case <-bgpping.done:
				return
			}
		}
	}()
	return bgpping.done
}

// Stop the bgpping
func (bgpping *BgpPing) Stop() {
	bgpping.done <- struct{}{}
}

func (bgpping BgpPing) ping() (time.Duration, net.Addr, error) {
	var remoteAddr net.Addr
	duration, errIfce := timeIt(func() interface{} {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", bgpping.target.Host, bgpping.target.Port), bgpping.target.Timeout)
		if err != nil {
			return err
		}
		remoteAddr = conn.RemoteAddr()
		conn.Close()
		return nil
	})
	if errIfce != nil {
		err := errIfce.(error)
		return 0, remoteAddr, err
	}
	return time.Duration(duration), remoteAddr, nil
}
