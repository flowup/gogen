package resource

import "github.com/flowup/gogen/generator-model"

// Feedback is model for the feedback entity
var Feedback = &model.Schema{
	Name: "Feedback",
	Fields: []model.Field{
		{
			Name: "Description",
			Type: model.String,
		},
	},
}
