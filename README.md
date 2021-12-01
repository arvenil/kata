# Kata 形🤺
[![GoDoc](https://godoc.org/github.com/arvenil/kata?status.svg)](https://pkg.go.dev/github.com/arvenil/kata?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/arvenil/kata)](https://goreportcard.com/report/github.com/arvenil/kata)
[![BuildStatus](https://github.com/arvenil/kata/workflows/go/badge.svg)](https://github.com/arvenil/kata/actions?query=workflow%3Ago)
[![codecov.io](https://codecov.io/gh/arvenil/kata/branch/master/graph/badge.svg)](https://codecov.io/gh/arvenil/kata)
[![coveralls.io](https://coveralls.io/repos/github/arvenil/kata/badge.svg?branch=master)](https://coveralls.io/github/arvenil/kata?branch=master)

> Kata is a Japanese word (型 or 形) meaning literally "form"
> referring to a detailed choreographed pattern of movements made to be practised alone,
> and within groups and in unison when training.
> It is practised in Japanese martial arts as a way to memorize and perfect the movements being executed.

# Install

```bash
brew tap arvenil/kata
brew install kata
```

# Usage

## bsearch

In computer science, binary algorithms, also known as half-interval algorithms, logarithmic algorithms, or binary chop,
is a algorithms algorithm that finds the position of a target value within a sorted array.
Binary algorithms compares the target value to the middle element of the array.
If they are not equal, the half in which the target cannot lie is eliminated
and the algorithms continues on the remaining half, again taking the middle element to compare to the target value,
and repeating this until the target value is found.
If the algorithms ends with the remaining half being empty, the target is not in the array.

```bash
$ bsearch
Usage of bsearch:
  -algorithm string
        exponential,
        interpolation,
        loop,
        loopslicing,
        recursive,
        recursiveslicing,
        standard (default "interpolation")
  -haystack value
        comma-separated, sorted, list of integers e.g. 1,5,7
  -needle int
        an integer to search for in haystack e.g. 5
```

```bash
$ bsearch --needle 5 --haystack 1,3,5,7
2
```

```bash
$ make bench
...
pkg: github.com/arvenil/kata/bsearch/algorithms
BenchmarkSearch/interpolation-12                 2656200               449 ns/op
BenchmarkSearch/loop-12                          3742628               327 ns/op
BenchmarkSearch/exponential-12                   3817282               308 ns/op
BenchmarkSearch/loopslicing-12                   4105732               293 ns/op
BenchmarkSearch/recursive-12                     2615310               459 ns/op
BenchmarkSearch/recursiveslicing-12              3311100               358 ns/op
BenchmarkSearch/standard-12                      2132502               558 ns/op
...
```

## ladder

A word-ladder puzzle begins with two words,
and to solve the puzzle one must find a chain of other words to link the two,
in which two adjacent words (that is, words in successive steps) differ by one letter.

```bash
$ ladder
Usage of ladder:
  -d string
        path to dictionary (default "/usr/share/dict/words")
  -json
        print results as json
  -p value
        two words separated by comma e.g. dog,cat
  -template string
        pretty-print results using a Go template
```

```bash
$ ladder -p dog,cat -p gold,lead
start  end   word-ladder            error (if any)
dog    cat   [dog dot cot cat]     
gold   lead  [gold goad load lead]  
```

```bash
$ ladder -p above,below -json | jq .
[
  {
    "Start": "above",
    "End": "below",
    "Words": [
      "above",
      "amove",
      "amoke",
      "smoke",
      "smoky",
      "sooky",
      "booky",
      "booly",
      "bolly",
      "bally",
      "balli",
      "balai",
      "balao",
      "baloo",
      "balow",
      "below"
    ]
  }
]
```

# Release

```sh
GITHUB_TOKEN=secret_token make release
```
