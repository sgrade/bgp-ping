# bgp-ping
Connectivity test for a BGP (Border Gateway Protocol) peer. 

## What is in the repo

Two bgp-ping implementations:
- Monolithic CLI app
- Client/Server interacting through gRPC

Two BIRD routers are integrated in the dev environment for easier testing. They start automatically, if the root bgpping folder is opened in [VS Code devcontainer](https://code.visualstudio.com/docs/remote/containers) and Docker Desktop is installed.

## Monolith

Currently it is a prototype, that pings 10.0.12.11:179 and accepts no parameters. Based on [cloverstd/tcping](https://github.com/cloverstd/tcping).

### Build
go build monolith/bgp-ping.go

### Use
`./bgp-ping`

### Sample report (inherited from tcping)
```
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

## Client/server
The prototype is in development

### Server
go run server/server.go

### Client
go run client/client.go
