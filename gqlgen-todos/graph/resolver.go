package graph

//This magic comment tells go generate what command to run when we want to regenerate our code.
//To run go generate recursively over your entire project, use this command: go generate ./...
//go:generate go run github.com/99designs/gqlgen

import "ram/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}
