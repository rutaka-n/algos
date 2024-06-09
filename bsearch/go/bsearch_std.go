package bsearch

import (
    "slices"
)

func bsearch_std(s []int, target int) int {
    idx, found := slices.BinarySearch(s, target)
    if !found {
        return -1
    }
    return idx
}
