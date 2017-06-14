package models

import "github.com/graphql-go/graphql"

// Usuario struct
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Endereco
}

// Mock de usuario
var usuario = &Usuario{
	ID:    1,
	Nome:  "Rodolfo Azevedo",
	Idade: 33,
	Endereco: Endereco{
		Rua: "Rua do Cabocão",
	},
}

// Definindo usuarioType, referenciado a estrutura completa no graphql
var usuarioType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Usuario",
	Fields: graphql.Fields{
		"nome": &graphql.Field{
			Type: graphql.String,
		},
		"idade": &graphql.Field{
			Type: graphql.Int,
		},
		"endereco": &graphql.Field{
			Type: enderecoType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return usuario.Endereco, nil
			},
		},
	},
})

// GetQueryType function
// Retorna o usuarioType para ser usado na main queryType
func (u Usuario) GetQueryType() *graphql.Object {
	return usuarioType
}

// GetUsuario getter
// Retorna a estrutura usuario
func (u Usuario) GetUsuario() *Usuario {
	return usuario
}
