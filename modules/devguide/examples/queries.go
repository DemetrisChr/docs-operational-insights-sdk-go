package main

import (
	"context"
	"fmt"

	cbinsights "github.com/couchbase/gocbinsights"
)

func queries() {
	cluster, err := cbinsights.NewCluster(
		connStr,
		cbinsights.NewBasicAuthCredential(username, password),
	)
	handleErr(err)

	scopeLevelQuery(context.Background(), cluster)
	clusterLevelQuery(context.Background(), cluster)
	positionalParamQuery(context.Background(), cluster)
	namedParamQuery(context.Background(), cluster)

	result, err := cluster.ExecuteQuery(context.Background(), "select 1")
	handleErr(err)
	handleResult(result)
	metadata(result)

	err = cluster.Close()
	handleErr(err)
}

func scopeLevelQuery(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::scopeLevelQuery[]
	scope := cluster.Database("my_database").Scope("my_scope")
	result, err := scope.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::scopeLevelQuery[]
}

func clusterLevelQuery(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::clusterLevelQuery[]
	result, err := cluster.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::clusterLevelQuery[]
}

func positionalParamQuery(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::positionalParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select ?=1",
		cbinsights.NewQueryOptions().SetPositionalParameters([]interface{}{1}),
	)
	handleErr(err)
	// end::positionalParamQuery[]

	handleResult(result)
}

func namedParamQuery(ctx context.Context, cluster *cbinsights.Cluster) {
	// tag::namedParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select $foo=1",
		cbinsights.NewQueryOptions().SetNamedParameters(map[string]interface{}{"foo": 1}),
	)
	handleErr(err)
	// end::namedParamQuery[]

	handleResult(result)
}

func handleResult(result *cbinsights.QueryResult) {
	// tag::handleResults[]
	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err := row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}

	if err := result.Err(); err != nil {
		handleErr(err)
	}
	// end::handleResults[]
}

func metadata(result *cbinsights.QueryResult) {
	// tag::metadata[]
	meta, err := result.MetaData()
	handleErr(err)

	fmt.Printf("Got meta: %v", meta)
	// end::metadata[]
}
