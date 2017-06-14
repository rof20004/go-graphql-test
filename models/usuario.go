package models

import "github.com/graphql-go/graphql"

// Usuario struct
type Usuario struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Endereco
}

var usuario = &Usuario{
	Nome:  "Rodolfo Azevedo",
	Idade: 33,
	Endereco: Endereco{
		Rua: "Rua do Cabocão",
	},
}

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
func (u Usuario) GetQueryType() *graphql.Object {
	return usuarioType
}

// GetUsuario getter
func (u Usuario) GetUsuario() *Usuario {
	return usuario
}
