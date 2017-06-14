package models

import "github.com/graphql-go/graphql"

// Endereco struct
type Endereco struct {
	Rua string `json:"rua"`
}

var endereco = &Endereco{
	Rua: "Avenida Brasil",
}

var enderecoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Endereco",
	Fields: graphql.Fields{
		"rua": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// GetQueryType function
func (u Endereco) GetQueryType() *graphql.Object {
	return enderecoType
}

// GetEndereco getter
func (u Endereco) GetEndereco() *Endereco {
	return endereco
}
