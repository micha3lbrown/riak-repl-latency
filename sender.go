package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	riak "github.com/basho/riak-go-client"
	util "github.com/basho/taste-of-riak/go/util"
)

type Tracer struct {
	Index      int64
	Epoch_time int64
	Datetime   time.Time
}

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
	}

	storeRecord(cluster)
}

func storeRecord(cluster *riak.Cluster) {
	now := time.Now()
	tracer := &Tracer{
		Index:      1,
		Epoch_time: now.Unix(),
		Datetime:   now,
	}

	t, err := json.Marshal(tracer)

	if err != nil {
		fmt.Println(err)
		return
	}

	obj := &riak.Object{
		ContentType:     "application/json",
		Charset:         "utf-8",
		ContentEncoding: "utf-8",
		Value:           []byte(t),
	}

	key := strconv.FormatInt(now.Unix(), 10)
	fmt.Println(key)

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucketType("default").
		WithBucket("test").
		WithKey(key).
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

	// svc := cmd.(*riak.StoreValueCommand)
	// rsp := svc.Response
}
