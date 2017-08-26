package main

import (
	"fmt"
	"os"

	"github.com/sketchground/shopgun"
)

func main() {
	key := os.Getenv("SHOPGUN_KEY")
	secret := os.Getenv("SHOPGUN_SECRET")

	client := shopgun.New(key, secret)

	dealers, err := client.DealerSearch("netto", shopgun.DefaultPagination())
	if err != nil {
		panic(err)
	}
	for _, d := range dealers {
		fmt.Println(d)
	}
}
