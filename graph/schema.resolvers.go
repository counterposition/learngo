package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/counterposition/learngo/graph/model"
)

func (r *groupResolver) Users(ctx context.Context, obj *model.Group, matching *string) ([]*model.User, error) {
	if matching == nil {
		return []*model.User{
			{
				Name: "Harish",
			},
			{
				Name: "Picard",
			},
			{
				Name: "Eun-Ji",
			},
		}, nil
	}

	if strings.Contains(*matching, "Harish") {
		return []*model.User{
			{Name: "Harish"},
		}, nil
	} else {
		return []*model.User{}, nil
	}
}

func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	// Typically, these would be fetched from a database
	var users = []*model.User{
		{
			Name: "Harish",
		},
		{
			Name: "Picard",
		},
		{
			Name: "Eun-Ji",
		},
	}
	var group1 = model.Group{
		Name:  "First",
		Users: users[:2],
	}
	var group2 = model.Group{
		Name:  "Second",
		Users: users[1:],
	}
	var groups = []*model.Group{&group1, &group2}

	return groups, nil
}

// Group returns GroupResolver implementation.
func (r *Resolver) Group() GroupResolver { return &groupResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type groupResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
