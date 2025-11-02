package resolver

import (
	"log"
	"time"
	"workbench/graphql-app/graph/model"
	"workbench/graphql-app/graph/resolver/scalar"
)

func badgerUserToGraphUser(u map[string]string) *model.User {
	t, err := time.Parse("2006-01-02", u["creation_date"])
	if err != nil {
		log.Panicf("Error during parses the creation date string into time.Time  : %v", err)
		return nil
	}
	return &model.User{
		ID:           u["id"],
		Name:         u["name"],
		Email:        u["email"],
		CreationDate: scalar.Date{Time: &t},
	}
}
