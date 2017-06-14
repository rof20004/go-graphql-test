package models

import "github.com/graphql-go/graphql"

// Endereco struct
type Endereco struct {
	Rua string `json:"rua"`
}

// Mock de endereco
var endereco = &Endereco{
	Rua: "Avenida Brasil",
}

// Definindo enderecoType, referenciado a estrutura completa no graphql
var enderecoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Endereco",
	Fields: graphql.Fields{
		"rua": &graphql.Field{
			Type: graphql.String,
		},
	},
})

// GetQueryType function
// Retorna o enderecoType para ser usado na main queryType
func (u Endereco) GetQueryType() *graphql.Object {
	return enderecoType
}

// GetEndereco getter
// Retorna a estrutura endereco
func (u Endereco) GetEndereco() *Endereco {
	return endereco
}
