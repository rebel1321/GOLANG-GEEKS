package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: "1", Name: "John Doe", Age: 30},
	{Id: "2", Name: "Jane Smith", Age: 25},
	{Id: "3", Name: "Sam Brown", Age: 22},
}

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"age":  &graphql.Field{Type: graphql.Int},
	},
})
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return users, nil
			},
		},
		"users": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string) 
					for _, user := range users {
						if user.Id == id {
							return user, nil
						}
					}
				  return nil, fmt.Errorf("User not found")
			},
		},
	},
})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"addUser": &graphql.Field{
			Type: userType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				age := p.Args["age"].(int)

				user := User{
					Id:   fmt.Sprintf("%d", len(users)+1),
					Name: name,
					Age:  age,
				}
				users = append(users, user)
				return user, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queryType,
	Mutation: mutationType,
})



func main() {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.Handle("/graphql", h)
	fmt.Println("Our first GraphQL server is running on http://localhost:8080/graphql")
	http.ListenAndServe(":8080", nil)
}