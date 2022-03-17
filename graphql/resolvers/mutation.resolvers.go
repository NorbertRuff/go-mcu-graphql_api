package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/norbertruff/go-graphql/graphql/model"
	gqlgen "github.com/norbertruff/go-graphql/graphql/server"
	auth2 "github.com/norbertruff/go-graphql/internal/pkg/auth"
	customError "github.com/norbertruff/go-graphql/internal/pkg/errors"
	linkService "github.com/norbertruff/go-graphql/internal/services/links"
	movieService "github.com/norbertruff/go-graphql/internal/services/movies"
	tvShowService "github.com/norbertruff/go-graphql/internal/services/tvshows"
	userService "github.com/norbertruff/go-graphql/internal/services/users"
)

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

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
