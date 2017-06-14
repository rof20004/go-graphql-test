package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
	"github.com/rof20004/go-graphql-test/models"
)

var usuario models.Usuario
var endereco models.Endereco

// Main query, to get all queries types
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"usuario": &graphql.Field{
			Type: usuario.GetQueryType(),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, isOK := p.Args["id"].(int)
				if isOK {
					if id == 1 {
						return usuario.GetUsuario(), nil
					}
				}
				return nil, nil
			},
		},
		"endereco": &graphql.Field{
			Type: endereco.GetQueryType(),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return endereco.GetEndereco(), nil
			},
		},
	},
})

// Schema instance
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {
	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":8080", nil)
}
