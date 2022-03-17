package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/norbertruff/go-graphql/graphql/model"
	gqlgen "github.com/norbertruff/go-graphql/graphql/server"
	auth2 "github.com/norbertruff/go-graphql/internal/pkg/auth"
)

func (r *movieResolver) User(ctx context.Context, obj *model.Movie) (*model.User, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return user, nil
}

// Movie returns gqlgen.MovieResolver implementation.
func (r *Resolver) Movie() gqlgen.MovieResolver { return &movieResolver{r} }

type movieResolver struct{ *Resolver }
