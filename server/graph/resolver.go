package graph

import (
	"github.com/brianlangdon/tada/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB db.DB
}

//func (r *Resolver) Query() graph.QueryResolver {
//	return &queryResolver{r}
//}

//type queryResolver struct{ *Resolver }

//func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
//	return r.DB.GetProgrammers(skill)
//}
