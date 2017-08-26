# Shopgun API client

This is not a generic client for the shopgun api.
It only does what i need it to do but could easily be extended to cover more endpoints and support
for more parameters on the individual endpoints.

Example usage:

```
package main

func main() {
	client := shopgun.New("my fancy api key", "my fancy secret key")

	dealers, err := client.Dealers("netto", shopgun.DefaultPagination())
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range dealers {
		// Do something with dealers
		fmt.Println(d)
	}
}
```
