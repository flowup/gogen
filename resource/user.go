package resource

import "github.com/flowup/gogen/generator-model"

// User is entity definition for the user
var User = &model.Schema{
	Name: "User",
	Fields: []model.Field{
		{
			Name: "Username",
			Type: model.String,
		},
		{
			Name: "Password",
			Type: model.String,
		},
		{
			Name: "Email",
			Type: model.String,
		},
	},
}
