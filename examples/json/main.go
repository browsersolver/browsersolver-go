package main

import (
	"context"
	"fmt"
	"log"

	"github.com/browsersolver/browsersolver-go"
)

func main() {
	client, err := browsersolver.NewClient("", nil) // reads BROWSERSOLVER_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
