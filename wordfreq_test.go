package main

import (
	"reflect"
	"testing"
)

func TestTopN_Basic(t *testing.T) {
	text := "the quick brown fox jumps over the lazy dog the fox"
	got := TopN(text, 3)
	want := []WordCount{
		{"the", 3},
		{"fox", 2},
		{"brown", 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopN_TieBreakAlphabetical(t *testing.T) {
	text := "cat bat ant"
	got := TopN(text, 3)
	want := []WordCount{
		{"ant", 1},
		{"bat", 1},
		{"cat", 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopN_CaseInsensitive(t *testing.T) {
	text := "Go go GO"
	got := TopN(text, 1)
	want := []WordCount{
		{"go", 3},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopN_Punctuation(t *testing.T) {
	text := "hello, hello world."
	got := TopN(text, 2)
	want := []WordCount{
		{"hello", 2},
		{"world", 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopN_NLargerThanUniqueWords(t *testing.T) {
	text := "one two"
	got := TopN(text, 10)
	want := []WordCount{
		{"one", 1},
		{"two", 1},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTopN_EmptyString(t *testing.T) {
	got := TopN("", 5)
	if len(got) != 0 {
		t.Errorf("expected empty result, got %v", got)
	}
}
