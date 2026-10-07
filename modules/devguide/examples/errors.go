package main

import (
	"context"
	"errors"
	"fmt"

	cbinsights "github.com/couchbase/gocbinsights"
)

func errorHandling() {
	cluster, err := cbinsights.NewCluster(
		connStr,
		cbinsights.NewBasicAuthCredential(username, password),
	)
	handleErr(err)

	analyticsError(context.Background(), cluster)
	queryError(context.Background(), cluster)

	err = cluster.Close()
	handleErr(err)
}

func analyticsError(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::analyticsError[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select 1=1",
	)
	if err != nil {
		var insightsErr cbinsights.InsightsError
		if errors.As(err, &insightsErr) {
			// Do something with this error.
		}

		// This error occurred out of a client-server interaction.
	}
	// end::analyticsError[]

	handleResult(result)
}

func queryError(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::queryError[]
	handleQueryError := func(err error) {
		if err != nil {
			var queryErr cbinsights.QueryError
			if errors.As(err, &queryErr) {
				fmt.Printf("Error code: %d, error message: %s", queryErr.Code(), queryErr.Message())
				return
			}

			// This error isn't a result of query processing, possibly something like a connection error.
		}
	}

	result, err := cluster.ExecuteQuery(
		ctx,
		"selec 1=1", // Syntax error
	)
	handleQueryError(err)
	// end::queryError[]

	if result == nil {
		return
	}

	for row := result.NextRow(); row != nil; row = result.NextRow() {
	}

	err = result.Err()
	handleQueryError(err)

	handleResult(result)
}
