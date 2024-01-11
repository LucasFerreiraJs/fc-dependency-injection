//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/lucasferreirajs/19-DI/product"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	wire.Build(
		// product.NewProductRepository,
		setRepositoryDependency,
		product.NewProductUsecase,
	)
	return &product.ProductUsecase{}
}
