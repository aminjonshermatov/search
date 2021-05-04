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
	files = append(files, "cmd/data/info21.txt")
	//files = append(files, "cmd/data/info2.txt")
	//files = append(files, "cmd/data/info3.txt")

	ch := search.All(ctx, phrase, files)

	//for k := range ch {
	//	log.Printf("result %#v", k)
	//	cancel()
	//}
	for {
		select {
		case res, ok := <-ch:
			if !ok {
				log.Print("done")
				cancel()
				return
			}
			log.Printf("result %#v", res)
		}
	}
}
