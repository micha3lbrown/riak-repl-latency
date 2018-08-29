# Riak Replication Latency
An application that will create Riak records at a cofigurable interval in a test bucket that is replicated between clusters.

### Riak Conf
By default Riak client uses Protocol Buffers, Riak listens for Protocol Buffers only on port `8087` by default. Riak also provides a Web UI that is available at http://localhost:8098/

### Retrieve object over HTTP
```
curl 127.0.0.1:8098/buckets\?buckets=true
curl 127.0.0.1:32770/buckets/bucket1/keys/key1
```

### Run
1. `go run sender.go`

### Build
1. `go build sender.go`
- or -
1. `GOOS=linux GOARCH=amd64 go build sender.go`
^ These options are only required if you wish to run the resulting binary on a different system. 
