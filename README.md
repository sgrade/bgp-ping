# bgp-ping
Connectivity test for a BGP (Border Gateway Protocol) peer

Original idea and some code to start from is from [https://github.com/cloverstd/tcping](https://github.com/cloverstd/tcping). It is being simplified for BGP and re-written as a client/server app.

## Sample report of monolithic bgp-ping (inherited from tcping)
```
./bgp-ping
Ping tcp://10.0.12.11:179(10.0.12.11:179) - Connected - time=982.225µs
Ping tcp://10.0.12.11:179(10.0.12.11:179) - Connected - time=345.467µs
Ping tcp://10.0.12.11:179(10.0.12.11:179) - Connected - time=411.968µs
Ping tcp://10.0.12.11:179(10.0.12.11:179) - Connected - time=356.94µs

Ping statistics tcp://10.0.12.11:179
        4 probes sent.
        4 successful, 0 failed.
Approximate trip times:
        Minimum = 345.467µs, Maximum = 982.225µs, Average = 524.15µs
```

## Client/server version
In development