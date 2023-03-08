package gql

import (
	"context"

	"github.com/brianlangdon/tada/db"
	"github.com/brianlangdon/tada/gql/gen"
	"github.com/brianlangdon/tada/model"
)

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
	return r.DB.GetProgrammers(skill)
}
