package main

import (
	"time"

	cbinsights "github.com/couchbase/gocbinsights"
)

func configuration() {
	// #tag::configuration[]
	cluster, err := cbinsights.NewCluster(
		connStr,
		cbinsights.NewBasicAuthCredential(username, password),
		cbinsights.NewClusterOptions().
			SetTimeoutOptions(
				cbinsights.NewTimeoutOptions().
					SetConnectTimeout(30*time.Second).
					SetQueryTimeout(2*time.Minute),
			),
	)
	handleErr(err)
	// #end::configuration[]

	err = cluster.Close()
	handleErr(err)
}
