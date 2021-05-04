package search

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"
)

type Result struct {
	Phrase	string
	Line	string
	LineNum	int64
	ColNum	int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)

	log.Print(len(files))
	for _, file := range files {
		log.Print("file")
		go func(ctx context.Context, file string, phrase string, ch chan<- []Result) {
			log.Print("go routine")
			select {
			case <-ctx.Done():
				close(ch)
			default:
			}
			_, err := os.Stat(file)
			if !os.IsNotExist(err) {
				src, err := os.Open(file)
				if err != nil {
					log.Print(err)
					return
				}
				defer func() {
					if cerr := src.Close(); cerr != nil {
						if err == nil {
							err = cerr
						}
					}
				}()

				reader := bufio.NewReader(src)
				lineNum := 1

				resultArr := make([]Result, 0)

				for {
					line, err := reader.ReadString('\n')
					if err == io.EOF {
						break
					}
					if err != nil {
						log.Print(err)
						return
					}

					line = strings.ReplaceAll(line, "\r\n", "")
					line = strings.ReplaceAll(line, "\n", "")

					idx := strings.Index(line, phrase)

					if idx != -1 {
						result := Result{
							Phrase: 	phrase,
							Line: 		line,
							LineNum: 	int64(lineNum),
							ColNum: 	int64(idx + 1),
						}

						resultArr = append(resultArr, result)
					}

					lineNum++
				}

				if len(resultArr) > 0 {
					log.Print("send")
					ch <- resultArr
				}
			}
		}(ctx, file, phrase, ch)
	}

	return ch
}
