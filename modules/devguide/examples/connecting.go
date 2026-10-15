package main

import (
	"time"

	cbinsights "github.com/couchbase/gocbinsights"
)

func connecting() {
	// tag::connecting[]
	cluster, err := cbinsights.NewCluster(
		connStr,
		cbinsights.NewBasicAuthCredential(username, password),
		// The third parameter is optional.
		// This example sets the default server query timeout to 3 minutes,
		// that is the timeout value sent to the query server.
		cbinsights.NewClusterOptions().SetTimeoutOptions(
			cbinsights.NewTimeoutOptions().SetQueryTimeout(3*time.Minute),
		),
	)
	handleErr(err)
	// end::connecting[]

	err = cluster.Close()
	handleErr(err)
}
