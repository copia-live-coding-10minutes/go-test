# Live Coding Challenge — Word Frequency Counter

**Time limit: 10 minutes**

---

## Problem

Implement the `TopN` function in `wordfreq.go`.

```go
func TopN(text string, n int) []WordCount
```

Given a string of text and an integer `n`, return the `n` most frequently occurring words in the text.

---

## Requirements

- Words are **case-insensitive** — `"Go"` and `"go"` are the same word.
- **Punctuation** directly attached to a word should be stripped — `"hello,"` counts as `"hello"`.
- Results must be sorted by **count descending**.
- Words with the **same count** must be sorted **alphabetically ascending**.
- If `n` is larger than the number of unique words, return all words.
- If the input is empty, return an empty slice.

---

## Example

Input:
```
text := "the quick brown fox jumps over the lazy dog the fox"
n    := 3
```

Output:
```
[{the 3} {fox 2} {brown 1}]
```

---

## Project Structure

```
.
├── go.mod
├── wordfreq.go       ← implement TopN here
├── wordfreq_test.go  ← do not modify
└── README.md
```

---

## Running Tests

```bash
go test ./...
```

To see detailed output:

```bash
go test -v ./...
```
