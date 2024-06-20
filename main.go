package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/edgedb/edgedb-go"
)

var errors = sync.Map{}

func work(ctx context.Context, edb *edgedb.Client) {
	var result []string
	for {
		err := edb.Query(ctx, "SELECT User.name", &result)
		if err == nil {
			continue
		}

		msg := err.Error()
		if _, ok := errors.LoadOrStore(msg, struct{}{}); !ok {
			fmt.Println(msg)
		}
	}
}

func main() {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	go work(ctx, client)
	go work(ctx, client)
	work(ctx, client)
}
