package main

import (
	"context"
	"github.com/aminjonshermatov/search/pkg/search"
	"log"
)

func main() {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)

	phrase := "foo"
	files := make([]string, 0)
	files = append(files, "cmd/data/info1.txt")
	files = append(files, "cmd/data/info2.txt")
	files = append(files, "cmd/data/info3.txt")

	ch := search.All(ctx, phrase, files)

	for k := range ch {
		log.Printf("result %#v", k)
	}
	log.Print("done")
	cancel()
}
