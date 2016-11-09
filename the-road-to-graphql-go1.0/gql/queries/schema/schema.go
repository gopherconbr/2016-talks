package schema

import (
	"log"

	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

type RepositoryOwner struct {
	Login string `json:"login"`
	Url   string `json:"url"`
}

var repositoryOwnerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Viewer",
	Fields: graphql.Fields{
		"login": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func init() {
	fields := graphql.Fields{
		"repositoryOwner": &graphql.Field{
			Type: repositoryOwnerType,
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return RepositoryOwner{
					Login: "octocat",
					Url: "https://api.github.com/users/octocat",
				}, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	Schema = s
}
