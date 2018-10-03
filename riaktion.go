package main

import (
	"flag"
	"fmt"

	crud "./crud"
	riak "github.com/basho/riak-go-client"
	util "github.com/basho/taste-of-riak/go/util"
)

func main() {

	var (
		// flagTrace   = flag.String("trace", "trace", "Create tracing records in Riak to monitor Realtime Replication")
		// flagSeed    = flag.String("seed", "seed", "Seed records into an existing cluster")
		// flagGet     = flag.String("get", "get", "Check that keys exist in a Riak cluster")
		// flagKeys    = flag.String("keys", "./tmp/riak-keys.json", "Path to file with JSON formatted Riak keys")
		flagVerbose = flag.Bool("verbose", false, "Turn on additional logging. e.g. riak.EnableDebugLogging")
		node        *riak.Node
		err         error
	)

	riak.EnableDebugLogging = *flagVerbose

	nodeOpts := &riak.NodeOptions{
		RemoteAddress: "127.0.0.1:8087",
	}

	if node, err = riak.NewNode(nodeOpts); err != nil {
		fmt.Println(err.Error())
	}

	nodes := []*riak.Node{node}
	opts := &riak.ClusterOptions{
		Nodes: nodes,
	}

	cluster, err := riak.NewCluster(opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	if err = cluster.Start(); err != nil {
		fmt.Println(err.Error())
	}

	crud.Seed(1000, cluster)

	defer func() {
		if err := cluster.Stop(); err != nil {
			util.ErrExit(err)
			fmt.Println(err.Error())
		}
	}()
}
