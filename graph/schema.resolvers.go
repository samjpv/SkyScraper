package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"test/SkyScraper/dbClient"
	"test/SkyScraper/graph/model"

	generated1 "github.com/samjpv/SkyScraper/graph/generated"
	model1 "github.com/samjpv/SkyScraper/graph/model"
)

func (r *queryResolver) Persons(ctx context.Context) ([]*model1.Person, error) {
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

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }
