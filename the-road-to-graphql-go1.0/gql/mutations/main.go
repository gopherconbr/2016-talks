package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/chris-ramon/talks/gql/gql/queries/schema"
)

func main() {
	query := `{ hello }`
	result := graphql.Do(graphql.Params{
		Schema:        schema.Schema,
		Mutation: rootMutation,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	fmt.Println(result)
}

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
		*/
		"createTodo": &graphql.Field{
			Type: todoType, // the return type for this field
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {

				// marshall and cast the argument value
				text, _ := params.Args["text"].(string)

				// figure out new id
				newID := RandStringRunes(8)

				// perform mutation operation here
				// for e.g. create a Todo and save to DB.
				newTodo := Todo{
					ID:   newID,
					Text: text,
					Done: false,
				}

				TodoList = append(TodoList, newTodo)

				// return the new Todo object that we supposedly save to DB
				// Note here that
				// - we are returning a `Todo` struct instance here
				// - we previously specified the return Type to be `todoType`
				// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
				return newTodo, nil
			},
		},
	},
})
