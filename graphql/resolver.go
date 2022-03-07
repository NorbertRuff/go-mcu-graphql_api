package graphql

import "github.com/norbertruff/go-graphql/graphql/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	movies []*models.Movie
	shows  []*models.TvShow
	links  []*models.Link
	user   []*models.User
}
