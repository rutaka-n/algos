#include <stdlib.h>
#include <stdio.h>
#include "bsearch.h"

#ifdef BSEARCH_STD
	int int_cmp(const void * a, const void * b) {
		return *(int*)a - *(int*)b;
	}
#endif

int
main(int argc, char **argv) {
	size_t n = 0;
	scanf("%zd", &n);
	int *arr = calloc(n, sizeof(int));
	for (size_t i = 0; i < n; i++) {
		scanf("%d", &arr[i]);
	}
	int target = 0;
	scanf("%d", &target);

#ifdef BSEARCH_STD
	int *res = NULL;
	res = bsearch((void*)(&target), (void*)arr, n, sizeof(int), int_cmp);
	if (res) {
		// need to return idx in array
		printf("%d\n", (int)((int*)res - arr));
	} else {
		printf("-1\n");
	}
#else
	printf("%d\n", _bsearch(arr, n, target));
#endif
	return 0;
}
