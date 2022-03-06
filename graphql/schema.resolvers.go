package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/norbertruff/go-graphql/graphql/generated/gqlgen"
	"github.com/norbertruff/go-graphql/graphql/generated/models"
	auth2 "github.com/norbertruff/go-graphql/internal/pkg/auth"
	customError "github.com/norbertruff/go-graphql/internal/pkg/errors"
	linkService "github.com/norbertruff/go-graphql/internal/services/links"
	movieService "github.com/norbertruff/go-graphql/internal/services/movies"
	tvShowService "github.com/norbertruff/go-graphql/internal/services/tvshows"
	userService "github.com/norbertruff/go-graphql/internal/services/users"
)

// CreateLink Creates a new link.
func (r *mutationResolver) CreateLink(ctx context.Context, variables models.NewLink) (*models.Link, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &models.Link{}, fmt.Errorf("access denied")
	}
	graphqlUser := &models.User{
		UserID:   user.UserID,
		Username: user.Username,
	}
	var newLink = new(models.Link)
	newLink.LinkID = uuid.New().String()
	newLink.PublishedBy = graphqlUser
	newLink.Title = variables.Title
	newLink.Address = variables.Address

	err := linkService.Save(newLink)
	if err != nil {
		return nil, err
	}
	return &models.Link{
			LinkID:      newLink.LinkID,
			Title:       variables.Title,
			Address:     variables.Address,
			PublishedBy: graphqlUser,
		},
		nil
}

// CreateUser Creates a new user.
func (r *mutationResolver) CreateUser(ctx context.Context, variables models.NewUser) (string, error) {
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

// CreateMovie Creates a new movie.
func (r *mutationResolver) CreateMovie(ctx context.Context, variables models.NewMovie) (*models.Movie, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &models.Movie{}, fmt.Errorf("access denied")
	}

	grahpqlUser := &models.User{
		UserID:   user.UserID,
		Username: user.Username,
	}

	newMovie := &models.Movie{
		MovieID:          new(uuid.UUID).String(),
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
		PublishedBy:      grahpqlUser,
	}
	r.movies = append(r.movies, newMovie)
	err := movieService.Save(*newMovie)
	if err != nil {
		return nil, err
	}

	return newMovie, nil
}

// CreateTvShow Creates a new Tv show.
func (r *mutationResolver) CreateTvShow(ctx context.Context, variables models.NewTvShow) (*models.TvShow, error) {
	user := auth2.ForContext(ctx)
	if user == nil {
		return &models.TvShow{}, fmt.Errorf("access denied")
	}

	grahpqlUser := &models.User{
		UserID:   user.UserID,
		Username: user.Username,
	}

	newTvShow := &models.TvShow{
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
		PublishedBy:    grahpqlUser,
	}

	return newTvShow, nil
}

// Login validating user and sending back jwt token.
func (r *mutationResolver) Login(ctx context.Context, variables models.Login) (string, error) {
	var user models.User
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

// RefreshToken refreshes token.
func (r *mutationResolver) RefreshToken(ctx context.Context, variables models.RefreshTokenInput) (string, error) {
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

// Links returns all links.
func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	dbLinks, err := linkService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbLinks, nil
}

// Movies returns all tv movies.
func (r *queryResolver) Movies(ctx context.Context) ([]*models.Movie, error) {
	dbMovies, err := movieService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbMovies, nil
}

// TvShows returns all tv shows.
func (r *queryResolver) TvShows(ctx context.Context) ([]*models.TvShow, error) {
	dbShows, err := tvShowService.GetAll()
	if err != nil {
		return nil, err
	}
	return dbShows, nil
}

// Mutation returns gqlgen.MutationResolver implementation.
func (r *Resolver) Mutation() gqlgen.MutationResolver { return &mutationResolver{r} }

// Query returns gqlgen.QueryResolver implementation.
func (r *Resolver) Query() gqlgen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
