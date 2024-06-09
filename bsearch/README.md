Binary Search
---
Binary search is an algorithm to search for elements within a sorted array.

The array is devided by the middle and target element is compared with the middle.
If they equal, element is found, otherwise the left or right side of the array is used for the next iteration.
if the target value less, we take the left side, if greate, we take the right one.

To find a middle element we need to calculate:
```
mid = (L + R)/2
```

or

```
mid = L + ((R - L) / 2)
```
the last formula garantees that result value does not exeeded the max integer value.

checkout implementation in GO and C.
to run tests in GO:
```
cd go && go test -v .
```
to run tests in C:
```
cd c && make && testrun
```
and for stdlib example:
```
cd c && make CFLAGS="-DBSEARCH_STD" && make testrun
```
