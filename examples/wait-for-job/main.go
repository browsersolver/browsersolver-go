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
	submitted, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com", IsAsync: browsersolver.Bool(true)})
	if err != nil {
		log.Fatal(err)
	}
	var accepted struct {
		JobID string `json:"job_id"`
	}
	if err := submitted.Decode(&accepted); err != nil {
		log.Fatal(err)
	}
	job, err := client.WaitForJob(ctx, accepted.JobID, &browsersolver.WaitForJobOptions{
		OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(job.Body))
}
