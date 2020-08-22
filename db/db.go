package db

import (
	"time"

	"github.com/couchbase/gocb"
)

// Collection is a reference to the bucket collection
var Collection *gocb.Collection

// Cluster is the connection to the database
var Cluster *gocb.Cluster

// InitClusterGetCollection initializes the cluster for you
func InitClusterGetCollection() {
	Cluster, _ = gocb.Connect("localhost", gocb.ClusterOptions{
		Username: "Administrator",
		Password: "1234,56Q",
	})

	bucket := Cluster.Bucket("properties")

	err := bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	Collection = bucket.DefaultCollection()
}
