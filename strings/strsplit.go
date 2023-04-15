package strings

import stdstrings "strings"

type StrSplit struct {
	inputStr string
	sep      string
	current  int
	runes    []rune
}

type Iterator interface {
	Next() string
}

// NewStrSplit ...
func NewStrSplit(s string, sep string) *StrSplit {
	return &StrSplit{
		inputStr: s,
		sep:      sep,
		runes:    []rune(s),
	}
}

func (ss *StrSplit) Next() string {
	if ss.current >= len(ss.runes) {
		return ""
	}

	start := ss.current

	subStr := string(ss.runes[start:])
	index := stdstrings.Index(subStr, ss.sep)
	if index >= 0 {
		ss.current = start + index + len(ss.sep)
		return string(ss.runes[start : start+index])
	}

	ss.current = len(ss.runes)
	ss.runes = []rune{}
	return subStr
}

func (ss *StrSplit) Split() Iterator {
	return ss
}
