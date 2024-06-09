#include "bsearch.h"
#include <stddef.h>
#include <stdlib.h>

int
_bsearch(int const *arr, size_t n, int target) {
	size_t l,r;
	l = 0;
	r = n;
	while (l <= r) {
		size_t mid = l + ((r - l)/2);
		if (target < arr[mid]) { // search in the fleft side
			r = mid-1;
			continue;
		} else if (target > arr[mid]) { // search in the right side
			l = mid+1;
			continue;
		}
		return mid;
	}
	return -1;
}
