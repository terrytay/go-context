package internal

import (
	"context"
	"os"

	tools "github.com/terrytay/go-context/tools/json"
)

type Product struct {
	Name        string
	Description string
}

func WithValue() {
	ctx := context.Background()

	ctx = addValue(ctx, Product{}, Product{Name: "Milo", Description: "A sweet drink"})

	product := readValue(ctx, Product{})
	tools.ToJSON(product, os.Stdout)
}

func addValue(c context.Context, key interface{}, value interface{}) context.Context {
	return context.WithValue(c, key, value)
}

func readValue(c context.Context, key interface{}) interface{} {
	return c.Value(key)
}
