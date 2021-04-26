package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/models"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *itemResolver) Category(ctx context.Context, obj *models.Item) (*generated.Category, error) {
	log.Println("itemResolver: Category")
	return &generated.Category{
		ID:        generateID(),
		Name:      "item_category",
		CreatedAt: time.Now(),
	}, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input generated.CreateItemInput) (*models.Item, error) {
	log.Println("mutationResolver: CreateItem")
	item := &models.Item{
		ID:         generateID(),
		Name:       input.Name,
		CategoryID: input.CategoryID,
		CreatedAt:  time.Now(),
	}
	r.items = append(r.items, item)
	return item, nil
}

func (r *mutationResolver) UpdateItem(ctx context.Context, input generated.UpdateItemInput) (*models.Item, error) {
	log.Println("mutationResolver: UpdateItem")
	ok := false
	updated := &models.Item{}
	for i, item := range r.items {
		if item.ID == input.ID {
			item.Name = input.Name
			item.CategoryID = input.CategoryID
			r.items[i] = item

			updated = item
			ok = true
			break
		}
	}

	if !ok {
		return nil, errors.New("not found")
	}
	return updated, nil
}

func (r *mutationResolver) DeleteItem(ctx context.Context, input generated.DeleteItemInput) (string, error) {
	log.Println("mutationResolver: DeleteItem")
	ok := false
	deletedID := ""
	for i, item := range r.items {
		if item.ID == input.ID {
			r.items = append(r.items[:i], r.items[i+1:]...)

			deletedID = item.ID
			ok = true
			break
		}
	}

	if !ok {
		return "", errors.New("not found")
	}
	return deletedID, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*models.Item, error) {
	log.Println("queryResolver: Items")
	return r.items, nil
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type itemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
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
