package wiki

import (
	"strings"
)

func parseLinks(text string) []string {
	rv := make([]string, 0)
	dupCheck := make(map[string]struct{})
	for {
		si := strings.Index(text, "[[")
		ei := strings.Index(text, "]]")
		if si == -1 || ei == -1 {
			break
		}
		t := text[si+2 : ei]
		text = text[ei+2:]
		if strings.Contains(t, ":") {
			continue
		}
		if i := strings.IndexAny(t, "#|"); i != -1 {
			t = t[:i]
		}
		s := strings.TrimSpace(t)
		if _, ok := dupCheck[s]; !ok {
			rv = append(rv, strings.TrimSpace(t))
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
		if 1 != -1 {
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
