package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AdanJSuarez/go-graphql/graph/generated"
	"github.com/AdanJSuarez/go-graphql/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	meetup, err := r.DB.InsertMeetup(input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	return meetup, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.DB.InsertUser(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.NewUser) (*model.User, error) {
	user := model.User{ID: fmt.Sprint(id), Username: input.Name, Email: input.Email}
	_, err := r.DB.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	meetups, err := r.DB.GetMeetups()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.DB.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Meetup(ctx context.Context, meetupID int) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, userID *int) (*model.User, error) {
	user, err := r.DB.GetUserWithMeetups(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
