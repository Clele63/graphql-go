package wrapper

import (
	"log"
	"workbench/graphql-app/queries/generated"
)

type WrappedQueries struct {
	inner  generated.Querier
	logger *log.Logger
}

func NewWrappedQueries(inner generated.Querier, logger *log.Logger) *WrappedQueries {
	return &WrappedQueries{
		inner:  inner,
		logger: logger,
	}
}
