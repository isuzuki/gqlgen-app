package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
	"math/rand"
	"strconv"
)

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return []*model.Item{{
		ID:   strconv.Itoa(rand.Int()),
		Name: "item",
	}}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
