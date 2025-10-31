package queries

import (
	"log"
	"workbench/graphql-app/queries/generated"
	"workbench/graphql-app/queries/wrapper"
)

func InitWrappedQueries(db generated.Querier) *wrapper.WrappedQueries {
	return wrapper.NewWrappedQueries(db, log.Default())
}
