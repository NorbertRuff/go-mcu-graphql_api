package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/norbertruff/go-graphql/graphql/generated/gqlgen"
	"github.com/norbertruff/go-graphql/graphql/model"
	auth2 "github.com/norbertruff/go-graphql/internal/pkg/auth"
	customError "github.com/norbertruff/go-graphql/internal/pkg/errors"
	linkService "github.com/norbertruff/go-graphql/internal/services/links"
	movieService "github.com/norbertruff/go-graphql/internal/services/movies"
	tvShowService "github.com/norbertruff/go-graphql/internal/services/tvshows"
	userService "github.com/norbertruff/go-graphql/internal/services/users"
)

func (r *linkResolver) User(ctx context.Context, obj *model.Link) (*model.User, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return &model.User{
		UserID: obj.UserID,
	}, nil
}

func (r *movieResolver) User(ctx context.Context, obj *model.Movie) (*model.User, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return user, nil
}

func (r *mutationResolver) CreateLink(ctx context.Context, variables model.NewLink) (*model.Link, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	newLink := &model.Link{
		LinkID:  uuid.New().String(),
		Title:   variables.Title,
		Address: variables.Address,
		UserID:  user.UserID,
	}

	err := linkService.Save(newLink)
	if err != nil {
		return nil, err
	}

	return newLink, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, variables model.NewUser) (string, error) {
	err := userService.Create(variables)
	if err != nil {
		return "", err
	}
	token, err := auth2.GenerateToken(variables.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreateMovie(ctx context.Context, variables model.NewMovie) (*model.Movie, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.Movie{}, fmt.Errorf("access denied")
	}

	newMovie := &model.Movie{
		MovieID:          uuid.New().String(),
		Title:            variables.Title,
		ReleaseDate:      variables.ReleaseDate,
		BoxOffice:        variables.BoxOffice,
		Duration:         variables.Duration,
		Overview:         variables.Overview,
		CoverURL:         variables.CoverURL,
		TrailerURL:       variables.TrailerURL,
		DirectedBy:       variables.DirectedBy,
		Phase:            variables.Phase,
		Saga:             variables.Saga,
		Chronology:       variables.Chronology,
		PostCreditScenes: variables.PostCreditScenes,
		ImdbID:           variables.ImdbID,
		UserID:           user.UserID,
	}
	err := movieService.Save(newMovie)
	if err != nil {
		return nil, err
	}
	r.movies = append(r.movies, newMovie)

	return newMovie, nil
}

func (r *mutationResolver) CreateTvShow(ctx context.Context, variables model.NewTvShow) (*model.TvShow, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.TvShow{}, fmt.Errorf("access denied")
	}

	graphQlUser := &model.User{
		UserID:   user.UserID,
		Username: user.Username,
	}

	newTvShow := &model.TvShow{
		ShowID:         new(uuid.UUID).String(),
		Title:          variables.Title,
		ReleaseDate:    variables.ReleaseDate,
		LastAiredDate:  variables.LastAiredDate,
		NumberEpisodes: variables.NumberEpisodes,
		NumberSeasons:  variables.NumberSeasons,
		Overview:       variables.Overview,
		CoverURL:       variables.CoverURL,
		TrailerURL:     variables.TrailerURL,
		DirectedBy:     variables.DirectedBy,
		Phase:          variables.Phase,
		Saga:           variables.Saga,
		ImdbID:         variables.ImdbID,
		UserID:         graphQlUser.UserID,
	}
	r.shows = append(r.shows, newTvShow)
	err := tvShowService.Save(newTvShow)
	if err != nil {
		return nil, err
	}
	return newTvShow, nil
}

func (r *mutationResolver) Login(ctx context.Context, variables model.Login) (string, error) {
	var user model.User
	user.Username = variables.Username
	user.Password = variables.Password
	correct := userService.Authenticate(user)
	if !correct {
		return "", &customError.WrongUsernameOrPasswordError{}
	}
	token, err := auth2.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, variables model.RefreshTokenInput) (string, error) {
	username, err := auth2.ParseToken(variables.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := auth2.GenerateToken(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

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

func (r *queryResolver) TvShows(ctx context.Context) ([]*model.TvShow, error) {
	dbShows, err := tvShowService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbShows, nil
}

func (r *tvShowResolver) User(ctx context.Context, obj *model.TvShow) (*model.User, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &model.User{}, fmt.Errorf("access denied")
	}
	return user, nil
}

// Link returns gqlgen.LinkResolver implementation.
func (r *Resolver) Link() gqlgen.LinkResolver { return &linkResolver{r} }

// Movie returns gqlgen.MovieResolver implementation.
func (r *Resolver) Movie() gqlgen.MovieResolver { return &movieResolver{r} }

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

// TvShow returns gqlgen.TvShowResolver implementation.
func (r *Resolver) TvShow() gqlgen.TvShowResolver { return &tvShowResolver{r} }

type linkResolver struct{ *Resolver }
type movieResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tvShowResolver struct{ *Resolver }
