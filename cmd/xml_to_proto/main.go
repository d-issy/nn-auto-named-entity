package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
	pb "github.com/nakalab/nn-auto-named-entity/protobuf"
	"github.com/nakalab/nn-auto-named-entity/wiki"
)

func main() {
	worker := make(chan struct{}, 2)
	wg := &sync.WaitGroup{}

	pbWiki := &pb.Wiki{}
	parser, err := wiki.NewXMLParser(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for {
		xmlPage, err := parser.Next()
		if err != nil {
			break
		}
		if xmlPage.Namespace != 0 {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker <- struct{}{}
			a, err := xmlPage.GetProtoBuf()
			if err == nil {
				count += 1
				fmt.Println(count, a.Title)
				pbWiki.Pages = append(pbWiki.Pages, a)
			}
			<-worker
		}()
	}
	wg.Wait()

	out, err := proto.Marshal(pbWiki)
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("data/wiki.bit", out, 0644); err != nil {
		log.Fatalln(err)
	}
}
