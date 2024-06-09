package bsearch

func bsearch(arr []int, target int) int {
    for l, r := 0, len(arr)-1; l <= r; {
        mid := l + ((r - l) / 2)
        // target less than middle, so go left side
        if target < arr[mid] {
            r = mid-1
            continue
        // target greater than middle, so go right side
        } else if target > arr[mid] {
            l = mid+1
            continue
        }
        // target equals middle, return it
        return mid
    }
    // target does not exist in the arr
    return -1
}
