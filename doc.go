/*
Package kata 形🤺 provides playground for various code exercises.

Kata is a Japanese word (型 or 形) meaning literally "form"
referring to a detailed choreographed pattern of movements made to be practised alone,
and within groups and in unison when training.
It is practised in Japanese martial arts as a way to memorize and perfect the movements being executed.

Commands

bsearch

In computer science, binary algorithms, also known as half-interval algorithms, logarithmic algorithms, or binary chop,
is a algorithms algorithm that finds the position of a target value within a sorted array.
Binary algorithms compares the target value to the middle element of the array.
If they are not equal, the half in which the target cannot lie is eliminated
and the algorithms continues on the remaining half, again taking the middle element to compare to the target value,
and repeating this until the target value is found.
If the algorithms ends with the remaining half being empty, the target is not in the array.

	Usage of bsearch:
	  -algorithm string
			choose from: exponential, interpolation (default), loop, loopslicing, recursive, recursiveslicing, standard
	  -haystack value
			comma-separated, sorted, list of integers e.g. 1,5,7
	  -needle int
			an integer to search for in haystack e.g. 5

ladder

A ladder puzzle begins with two words,
and to solve the puzzle one must find a chain of other words to link the two,
in which two adjacent words (that is, words in successive steps) differ by one letter.

	Usage of ladder:
	  -d string
			path to dictionary (default "/usr/share/dict/words")
	  -json
			print results as json
	  -p value
			two words separated by comma e.g. dog,cat
	  -template string
			pretty-print results using a Go template

Benchmark

cmd/bsearch:
	pkg: github.com/arvenil/kata/bsearch/algorithms
	BenchmarkSearch/interpolation-12                 2656200               449 ns/op
	BenchmarkSearch/loop-12                          3742628               327 ns/op
	BenchmarkSearch/exponential-12                   3817282               308 ns/op
	BenchmarkSearch/loopslicing-12                   4105732               293 ns/op
	BenchmarkSearch/recursive-12                     2615310               459 ns/op
	BenchmarkSearch/recursiveslicing-12              3311100               358 ns/op
	BenchmarkSearch/standard-12                      2132502               558 ns/op

cmd/ladder:
	$ time ladder -p dog,cat -p godo,loto -p kot,pies -p gold,lead -p above,below -p soup,rice
	start  end    ladder                                                                                        error (if any)
	dog    cat    [dog dot cot cat]
	godo   loto   []                                                                                                 could not find 'godo' in dictionary
	kot    pies   []                                                                                                 words 'kot'(3) and 'pies'(4) have different length
	gold   lead   [gold goad load lead]
	above  below  [above amove amoke smoke smoky sooky booky booly bolly bally balli balai balao baloo balow below]
	soup   rice   [soup roup roue role rile rice]
	5.44s user 0.50s system 197% cpu 3.010 total
*/
package kata
