package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/norbertruff/go-graphql/graphql/model"
	gqlgen "github.com/norbertruff/go-graphql/graphql/server"
	linkService "github.com/norbertruff/go-graphql/internal/services/links"
	movieService "github.com/norbertruff/go-graphql/internal/services/movies"
	tvShowService "github.com/norbertruff/go-graphql/internal/services/tvshows"
)

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	dbLinks, err := linkService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbLinks, nil
}

func (r *queryResolver) Movies(ctx context.Context) ([]*model.Movie, error) {
	dbMovies, err := movieService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbMovies, nil
}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TvShows(ctx context.Context) ([]*model.TvShow, error) {
	dbShows, err := tvShowService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbShows, nil
}

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
