package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gql-exp/graph/generated"
	"gql-exp/graph/model"
)

func (r *queryResolver) AssetSymbols(ctx context.Context) ([]*model.AssetSymbol, error) {
	symbol1 := model.AssetSymbol{
		Currency:      "USD",
		Description:   "PLEXUS HOLDINGS PLC",
		DisplaySymbol: "PLXXF",
		Figi:          "BBG003PMR1G8",
		Mic:           "OOTC",
		Symbol:        "PLXXF",
		AssetType:     "Common Stock",
	}
	r.assetSymbols = append(r.assetSymbols, &symbol1)
	return r.assetSymbols, nil
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
type todoResolver struct{ *Resolver }
