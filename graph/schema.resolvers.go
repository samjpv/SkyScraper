package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"test/SkyScraper/graph/generated"
	"test/SkyScraper/graph/model"
	"test/SkyScraper/util/dbClient"
)

func (r *queryResolver) Persons(ctx context.Context) ([]*model.Person, error) {

	var db = dbClient.OpenDb()
	var (
		firstName string
		lastName  string
	)
	people, err := db.Query("SELECT * FROM People")
	if err != nil {
		panic(err.Error())
	}
	var peopleResponse []*model.Person
	for people.Next() {
		err := people.Scan(&firstName, &lastName)
		if err != nil {
			panic(err.Error())
		}
		person := model.Person{
			Firstname: firstName,
			Lastname:  lastName,
		}
		peopleResponse = append(peopleResponse, &person)
	}
	db.Close()
	return peopleResponse, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
