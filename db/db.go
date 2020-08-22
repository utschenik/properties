package db

import (
	"time"

	"github.com/couchbase/gocb"
)

// Collection is a reference to the bucket collection
var Collection *gocb.Collection

// InitClusterGetCollection initializes the cluster for you
func InitClusterGetCollection() {
	cluster, err := gocb.Connect("localhost", gocb.ClusterOptions{
		Username: "Administrator",
		Password: "1234,56Q",
	})
	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("properties")

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	Collection = bucket.DefaultCollection()
}
