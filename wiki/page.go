package wiki

import (
	"errors"
	"strings"
	"sync"

	pb "github.com/nakalab/nn-auto-named-entity/protobuf"
)

func (p *XMLPage) GetProtoBuf() (*pb.Page, error) {
	if len(p.Revisions) == 0 {
		return nil, errors.New("text is not contained")
	}
	page := &pb.Page{}
	page.Id = p.ID
	page.Title = p.Title

	text := p.Revisions[0].Text
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func() {
		page.Links = parseLinks(text)
		wg.Done()
	}()
	go func() {
		page.Categories = parseCategories(text)
		wg.Done()
	}()
	go func() {
		page.Templates = parseTemplates(text)
		wg.Done()
	}()
	wg.Wait()
	return page, nil
}

func parseLinks(text string) []string {
	rv := make([]string, 0)
	dupCheck := make(map[string]struct{})
	for {
		si := strings.Index(text, "[[")
		if si == -1 {
			break
		}
		text = text[si:]
		c := 0
		i := 0
		t := []rune(text)
		for i+1 < len(t) {
			if t[i] == '[' && t[i+1] == '[' {
				c += 1
				i += 2
				continue
			} else if t[i] == ']' && t[i+1] == ']' {
				c -= 1
				if c == 0 {
					break
				}
				i += 2
				continue
			}
			i += 1
		}
		ei := len(string(t[:i]))
		s := text[2:ei]
		if ei+2 >= len(text) {
			break
		}
		text = text[ei+2:]
		if strings.Contains(s, ":") {
			continue
		}
		if i := strings.IndexAny(s, "#|"); i != -1 {
			s = s[:i]
		}
		s = strings.TrimSpace(s)
		if _, ok := dupCheck[s]; !ok {
			rv = append(rv, strings.TrimSpace(s))
			dupCheck[s] = struct{}{}
		}
	}
	return rv
}

func parseCategories(text string) []string {
	rv := make([]string, 0)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[[Category:") {
			line = line[11:]
		} else if strings.HasPrefix(line, "[[カテゴリ:") {
			line = line[15:]
		} else {
			continue
		}
		i := strings.IndexAny(line, "|]")
		if i != -1 {
			line = line[:i]
		}
		line = strings.TrimSpace(line)
		rv = append(rv, line)
	}
	return rv
}

func parseTemplates(text string) []string {
	rv := make([]string, 0)
	dupCheck := make(map[string]struct{})
	for {
		si := strings.Index(text, "{{")
		if si == -1 {
			break
		}
		text = text[si:]
		c := 0
		i := 0
		t := []rune(text)
		for i+1 < len(t) {
			if t[i] == '{' && t[i+1] == '{' {
				c += 1
				i += 2
				continue
			} else if t[i] == '}' && t[i+1] == '}' {
				c -= 1
				if c == 0 {
					break
				}
				i += 2
				continue
			}
			i += 1
		}
		ei := len(string(t[:i]))
		s := text[2:ei]
		if ei+2 >= len(text) {
			break
		}
		text = text[ei+2:]
		if strings.HasPrefix(s, "#") {
			continue
		}
		if i := strings.IndexAny(s, "|\n"); i != -1 {
			s = s[:i]
		}
		if _, ok := dupCheck[s]; !ok {
			rv = append(rv, s)
			dupCheck[s] = struct{}{}
		}
	}
	return rv
}
