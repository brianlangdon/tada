package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/brianlangdon/tada/graph/model"
)

// UpsertProgrammer is the resolver for the upsertProgrammer field.
func (r *mutationResolver) UpsertProgrammer(ctx context.Context, input model.ProgrammerInput) (*model.Programmer, error) {
	panic(fmt.Errorf("not implemented: UpsertProgrammer - upsertProgrammer"))
}

// Programmers is the resolver for the programmers field.
func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
	return r.DB.GetProgrammers(skill)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
