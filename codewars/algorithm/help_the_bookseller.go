package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

// StockList from
//   https://www.codewars.com/kata/54dc6f5a224c26032800005c/train/go
func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	set := make(map[string]bool)
	counter := make(map[string]int)
	for _, c := range listCat {
		set[c] = true
	}

	for _, art := range listArt {
		c := art[:1]
		if !set[c] {
			continue
		}
		items := strings.Split(art, " ")
		n, _ := strconv.Atoi(items[1])
		counter[c] += n
	}

	vals := make([]string, 0, len(listCat))
	for _, c := range listCat {
		vals = append(vals, fmt.Sprintf("(%s : %d)", c, counter[c]))
	}
	return strings.Join(vals, " - ")
}
