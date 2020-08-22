package db

import (
	"time"

	"github.com/couchbase/gocb"
)

// InitClusterGetCollection initializes the clust for you, and returns the collection from the bucket
func InitClusterGetCollection() *gocb.Collection {
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

	collection := bucket.DefaultCollection()

	return collection
}
