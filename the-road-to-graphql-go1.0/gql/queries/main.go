package main

import (
	"fmt"
	"log"

	"github.com/chris-ramon/talks/gql/gql/queries/schema"
	"github.com/graphql-go/graphql"
	"github.com/kr/pretty"
)

func main() {
	query := `{ repositoryOwner(login: "octocat") { login url } }`
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	fmt.Printf("%# v", pretty.Formatter(result))
}
