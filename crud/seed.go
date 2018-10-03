package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	riak "github.com/basho/riak-go-client"
)

type Record struct {
	Datetime time.Time
	Body     []byte
}

func Seed(count int, cluster *riak.Cluster) {
	for i := 0; i < count; i++ {
		Store(cluster)
		count += i
		fmt.Println(i)
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("Done?")
}

func Store(cluster *riak.Cluster) {
	now := time.Now()
	body, _ := ioutil.ReadFile("../tmp/ipsum.txt")

	record := &Record{
		Datetime: now,
		Body:     body,
	}
	r, err := json.Marshal(record)

	if err != nil {
		fmt.Println(err)
		return
	}

	obj := &riak.Object{
		ContentType:     "application/json",
		Charset:         "utf-8",
		ContentEncoding: "utf-8",
		Value:           []byte(r),
	}

	// key := strconv.FormatInt(now.Unix(), 10)
	// key := strconv.Itoa(index)
	// fmt.Println(key)

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithBucket("refactor").
		// WithKey(key).
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
}
