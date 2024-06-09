#include <stdlib.h>
#include <stdio.h>
#include "bsearch.h"

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

	printf("%d\n", _bsearch(arr, n, target));
	return 0;
}
