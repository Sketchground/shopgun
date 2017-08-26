package shopgun_test

import (
	"fmt"
	"os"

	"github.com/sketchground/shopgun"
)

func ExampleNew() {
	key := os.Getenv("SHOPGUN_KEY")

	client := shopgun.New(key)

	dealers, err := client.DealerSearch("netto", shopgun.DefaultPagination())
	if err != nil {
		panic(err)
	}
	for _, d := range dealers {
		fmt.Println(d)
	}
}
