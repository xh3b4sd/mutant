package wave

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/xh3b4sd/mutant"
)

func Test_Wave_Lifecycle(t *testing.T) {
	testCases := []struct {
		shift func(p mutant.Interface)
		index []int
	}{
		// Case 0 ensures the permutation works as expected.
		{
			shift: func(p mutant.Interface) {
			},
			index: []int{0, 0, 0},
		},
		// Case 1 ensures the permutation works as expected.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
			},
			index: []int{0, 0, 1},
		},
		// Case 2 ensures the permutation works as expected.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
				p.Shift()
			},
			index: []int{0, 1, 0},
		},
		// Case 3 ensures the permutation works as expected.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
				p.Shift()
				p.Shift()
			},
			index: []int{1, 0, 0},
		},
		// Case 4 ensures mutating beyond capacity is idempotent.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
			},
			index: []int{1, 0, 0},
		},
		// Case 5 ensures resetting starts over again.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()

				p.Reset()
			},
			index: []int{0, 0, 0},
		},
		// Case 6 ensures resetting starts over again.
		{
			shift: func(p mutant.Interface) {
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()
				p.Shift()

				p.Reset()

				p.Shift()
			},
			index: []int{0, 0, 1},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var err error

			var p mutant.Interface
			{
				c := Config{
					Length: 3,
				}

				p, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			tc.shift(p)

			index := p.Index()

			if !reflect.DeepEqual(index, tc.index) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.index, index))
			}
		})
	}
}
