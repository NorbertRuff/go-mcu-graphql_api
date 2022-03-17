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

func (r *tvShowResolver) User(ctx context.Context, obj *model.TvShow) (*model.User, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return user, nil
}

// TvShow returns gqlgen.TvShowResolver implementation.
func (r *Resolver) TvShow() gqlgen.TvShowResolver { return &tvShowResolver{r} }

type tvShowResolver struct{ *Resolver }
