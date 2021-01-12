package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"test/SkyScraper/graph/generated"
	"test/SkyScraper/graph/model"
)

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {
	var persons []*model.Person
	person := model.Person{
		Firstname: "Sam",
		Lastname:  "van der Sloot",
	}
	persons = append(persons, &person)

	person2 := model.Person{
		Firstname: "Firstname",
		Lastname:  "Lastname",
	}
	persons = append(persons, &person2)
	return persons, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
