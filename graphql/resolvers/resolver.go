package resolvers

import "github.com/norbertruff/go-graphql/graphql/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	movies []*model.Movie
	shows  []*model.TvShow
	links  []*model.Link
	user   []*model.User
}
