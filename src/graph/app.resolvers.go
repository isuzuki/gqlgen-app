package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/models"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func (r *itemResolver) Category(ctx context.Context, obj *models.Item) (*generated.Category, error) {
	log.Println("itemResolver: Category")
	return &generated.Category{
		ID:        generateID(),
		Name:      "item_category",
		CreatedAt: time.Now(),
	}, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*models.Item, error) {
	log.Println("queryResolver: Items")
	return []*models.Item{{
		ID:         generateID(),
		Name:       "item",
		CategoryID: generateID(),
		CreatedAt:  time.Now(),
	}}, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*generated.Category, error) {
	log.Println("queryResolver: Categories")

	allItems := []*models.Item{{
		ID:         generateID(),
		Name:       "item_" + generateID(),
		CategoryID: generateID(),
		CreatedAt:  time.Now(),
	}}

	catItems := make([]*models.Item, 0)
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		for _, a := range f.Arguments {
			if a.Name == "createdSince" {
				t, _ := time.Parse(time.RFC3339, a.Value.Raw)
				for _, i := range allItems {
					if i.CreatedAt.After(t) {
						catItems = append(catItems, i)
					}
				}
			}

			fmt.Println(a.Name, a.Value)
		}
	}

	return []*generated.Category{{
		ID:        generateID(),
		Name:      "category",
		CreatedAt: time.Now(),
		Items:     catItems,
	}}, nil
}

func (r *queryResolver) Item(ctx context.Context, id string) (*models.Item, error) {
	log.Println("queryResolver: Item")
	item := &models.Item{
		ID:         id,
		Name:       "item_" + id,
		CategoryID: generateID(),
		CreatedAt:  time.Now(),
	}
	return item, nil
}

// Item returns generated.ItemResolver implementation.
func (r *Resolver) Item() generated.ItemResolver { return &itemResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type itemResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func generateID() string {
	return strconv.Itoa(rand.Int())
}
