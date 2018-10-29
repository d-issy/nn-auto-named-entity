package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/golang/protobuf/proto"
	pb "github.com/nakalab/nn-auto-named-entity/protobuf"
	"github.com/nakalab/nn-auto-named-entity/wiki"
)

func main() {
	cpuNum := runtime.NumCPU()
	worker := make(chan struct{}, cpuNum)
	wg := &sync.WaitGroup{}
	mu := new(sync.RWMutex)

	pbWiki := &pb.Wiki{}
	pbWiki.Pages = make(map[string]*pb.Page)

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
		if count > 300 {
			break
		}
		wg.Add(1)
		worker <- struct{}{}
		go func() {
			defer wg.Done()
			count += 1
			a, err := xmlPage.GetProtoBuf()
			if err == nil {
				fmt.Println(count, a.Title)
				mu.Lock()
				pbWiki.Pages[a.Title] = a
				mu.Unlock()
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
