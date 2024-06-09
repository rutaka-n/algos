package bsearch

import (
	"testing"
)

func TestBSearch(t *testing.T) {
	cases := []struct {
        desc string
		arr      []int
		target   int
		expected int
	}{
        {
            desc: "target element exists",
            arr: []int{0,1,2,3,4,5,6,7,8,9},
            target: 0,
            expected: 0,
        },
        {
            desc: "target element exists",
            arr: []int{0,1,2,3,4,5,6,7,8,9},
            target: 9,
            expected: 9,
        },
        {
            desc: "target element exists",
            arr: []int{0,10,20,30,40,50,60,70,80,90},
            target: 70,
            expected: 7,
        },
        {
            desc: "empty array",
            arr: []int{},
            target: 0,
            expected: -1,
        },
        {
            desc: "no target element in array",
            arr: []int{5,10,25,30,45,50,65,75,80,95},
            target: 36,
            expected: -1,
        },
    }

    for _, tc := range cases {
        tc := tc
        t.Run(tc.desc, func(t *testing.T) {
            got := bsearch(tc.arr, tc.target)
            if got != tc.expected {
                t.Errorf("expected %d, got %d\n", tc.expected, got)
            }
        })
    }
}

func TestBSearchStd(t *testing.T) {
	cases := []struct {
        desc string
		arr      []int
		target   int
		expected int
	}{
        {
            desc: "target element exists",
            arr: []int{0,1,2,3,4,5,6,7,8,9},
            target: 0,
            expected: 0,
        },
        {
            desc: "target element exists",
            arr: []int{0,1,2,3,4,5,6,7,8,9},
            target: 9,
            expected: 9,
        },
        {
            desc: "target element exists",
            arr: []int{0,10,20,30,40,50,60,70,80,90},
            target: 70,
            expected: 7,
        },
        {
            desc: "empty array",
            arr: []int{},
            target: 0,
            expected: -1,
        },
        {
            desc: "no target element in array",
            arr: []int{5,10,25,30,45,50,65,75,80,95},
            target: 36,
            expected: -1,
        },
    }

    for _, tc := range cases {
        tc := tc
        t.Run(tc.desc, func(t *testing.T) {
            got := bsearch_std(tc.arr, tc.target)
            if got != tc.expected {
                t.Errorf("expected %d, got %d\n", tc.expected, got)
            }
        })
    }
}
