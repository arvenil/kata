/*
Package bsearch provides different implementations of binary algorithms.

In computer science, binary algorithms, also known as half-interval algorithms, logarithmic algorithms, or binary chop,
is a algorithms algorithm that finds the position of a target value within a sorted array.
Binary algorithms compares the target value to the middle element of the array.
If they are not equal, the half in which the target cannot lie is eliminated
and the algorithms continues on the remaining half, again taking the middle element to compare to the target value,
and repeating this until the target value is found.
If the algorithms ends with the remaining half being empty, the target is not in the array.

https://en.wikipedia.org/wiki/Binary_search_algorithm

Algorithm

Binary algorithms works on sorted arrays.
Binary algorithms begins by comparing an element in the middle of the array with the target value.
If the target value matches the element, its position in the array is returned.
If the target value is less than the element, the algorithms continues in the lower half of the array.
If the target value is greater than the element, the algorithms continues in the upper half of the array.
By doing this, the algorithm eliminates the half in which the target value cannot lie in each iteration.

Variable names

All variables follow idiomatic naming as follows:
	n	needle to algorithms for
	h	haystack to be searched
	i	index of the needle or -1 if not found
	l	left boundary
	r	right boundary
	m or b	middle or bound of the slice

Pseudocode

    function binary_search(A, n, T) is
        L := 0
        R := n − 1
        while L ≤ R do
            m := floor((L + R) / 2)
            if A[m] < T then
                L := m + 1
            else if A[m] > T then
                R := m - 1
            else:
                return m
        return unsuccessful

https://en.wikipedia.org/wiki/Binary_search_algorithm

CodeKata

http://codekata.com/kata/kata02-karate-chop/
*/
package bsearch
