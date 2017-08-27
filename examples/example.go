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

	fmt.Println("-----")

	cats, err := client.Catalogs([]string{"90f2VL"}, shopgun.DefaultPagination())
	if err != nil {
		panic(err)
	}
	for _, c := range cats {
		fmt.Println(c)
	}

	fmt.Println("-----")

	offers, err := client.Offers([]string{"5744o8E"}, shopgun.DefaultPagination())
	if err != nil {
		panic(err)
	}
	for _, o := range offers {
		fmt.Println(o)
	}
}
