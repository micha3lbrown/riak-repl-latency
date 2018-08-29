package main

import (
	"fmt"

	riak "github.com/basho/riak-go-client"
	util "github.com/basho/taste-of-riak/go/util"
)

func main() {
	riak.EnableDebugLogging = true

	nodeOpts := &riak.NodeOptions{
		RemoteAddress: "127.0.0.1:8087",
	}

	var node *riak.Node
	var err error
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

	defer func() {
		if err := cluster.Stop(); err != nil {
			util.ErrExit(err)
			fmt.Println(err.Error())
		}
	}()

	if err = cluster.Start(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ping passed")
	}

	storeRecord(cluster)
}

func storeRecord(cluster *riak.Cluster) {
	obj := &riak.Object{
		ContentType:     "text/plain",
		Charset:         "utf-8",
		ContentEncoding: "utf-8",
		Value:           []byte("riak stored object"),
	}

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucketType("default").
		WithBucket("test").
		WithKey("riakobj1").
		WithContent(obj).
		Build()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = cluster.Execute(cmd); err != nil {
		fmt.Println(err.Error())
		return
	}

	svc := cmd.(*riak.StoreValueCommand)
	rsp := svc.Response
	fmt.Println(svc)
	fmt.Println(rsp)
	fmt.Println(rsp.VClock)

}
