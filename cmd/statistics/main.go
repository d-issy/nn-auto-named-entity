package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
	"unicode"

	proto "github.com/golang/protobuf/proto"
	pb "github.com/nakalab/nn-auto-named-entity/protobuf"
)

type Pair struct {
	Key   string
	Value uint64
}

type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Less(i int, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) Swap(i int, j int) {
	p[i].Key, p[j].Key = p[j].Key, p[i].Key
	p[i].Value, p[j].Value = p[j].Value, p[i].Value
}

type Counter struct {
	d  map[string]uint64
	mu *sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		d:  make(map[string]uint64),
		mu: &sync.RWMutex{},
	}
}

func (c *Counter) CountRune(s string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, r := range s {
		k := string(r)
		if unicode.In(r, unicode.Number) {
			k = "number"
		} else if unicode.In(r, unicode.Latin) {
			k = "alphabet"
		} else if unicode.In(r, unicode.Hiragana) {
			k = "ひらがな"
		} else if unicode.In(r, unicode.Katakana) {
			k = "カタカナ"
		}
		if _, ok := c.d[k]; !ok {
			c.d[k] = 0
		}
		c.d[k] += 1
	}
}

func (c *Counter) CountTemplate(ss []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ex := make(map[string]struct{})
	for _, s := range ss {
		if _, ok := ex[s]; ok {
			continue
		}
		ex[s] = struct{}{}
		if _, ok := c.d[s]; !ok {
			c.d[s] = 0
		}
		c.d[s] += 1
	}
}

func (c *Counter) GetSortPairs() PairList {
	d := make(PairList, 0)
	for k, v := range c.d {
		d = append(d, Pair{
			Key:   k,
			Value: v,
		})
	}
	sort.Sort(d)
	return d
}

func CountPage(p *pb.Page, cm, lm, tm *Counter) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		for _, c := range p.Categories {
			cm.CountRune(c)
		}
		wg.Done()
	}()
	go func() {

		for _, l := range p.Links {
			lm.CountRune(l)
		}
		wg.Done()
	}()
	go func() {
		tm.CountTemplate(p.Templates)
		wg.Done()
	}()
	wg.Wait()
}

func main() {
	cpuNum := runtime.NumCPU()
	worker := make(chan struct{}, cpuNum)
	wg := &sync.WaitGroup{}

	wiki := &pb.Wiki{}
	in, err := ioutil.ReadFile(filepath.Join("data", "wiki.bin"))
	if err != nil {
		log.Fatalf("file read error: ", err)
	}
	err = proto.Unmarshal(in, wiki)
	if err != nil {
		log.Fatalf("proto file parse error: ", err)
	}

	cm := NewCounter()
	lm := NewCounter()
	tm := NewCounter()

	for _, page := range wiki.Pages {
		wg.Add(1)
		worker <- struct{}{}
		go func() {
			defer wg.Done()
			CountPage(page, cm, lm, tm)
			<-worker
		}()
	}
	wg.Wait()

	stat := &pb.Statistics{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for _, c := range cm.GetSortPairs() {
			stat.Categories = append(stat.Categories, &pb.Counter{
				Key:   c.Key,
				Count: c.Value,
			})
		}
	}()

	go func() {
		defer wg.Done()
		for _, l := range lm.GetSortPairs() {
			stat.Links = append(stat.Links, &pb.Counter{
				Key:   l.Key,
				Count: l.Value,
			})
		}
	}()

	go func() {
		defer wg.Done()
		for _, t := range tm.GetSortPairs() {
			stat.Templates = append(stat.Templates, &pb.Counter{
				Key:   t.Key,
				Count: t.Value,
			})
		}
	}()
	wg.Wait()

	out, err := proto.Marshal(stat)
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile("data/stat.bin", out, 0644); err != nil {
		log.Fatalln(err)
	}
}
