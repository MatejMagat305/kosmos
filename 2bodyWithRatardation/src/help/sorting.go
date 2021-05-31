package help

import (
	"sort"
	"strings"
)

type alphabetic []string

func (list alphabetic) Len() int { return len(list) }

func (list alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list alphabetic) Less(i, j int) bool {
	si := list[i]
	sj := list[j]
	var siLower = strings.ToLower(si)
	var sjLower = strings.ToLower(sj)
	if siLower == sjLower {
		return si < sj
	}
	return siLower < sjLower
}

func Sort(slice []string) {
	sort.Sort(alphabetic(slice))
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
