package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type arguments struct {
	strs    []string
	numbers []int
	cond    bool
	expRess []string
	expErr  string
}

func TestDoInt(t *testing.T) {
	lettersLow := []string{"a", "b", "c", "d"}
	lettersAny := []string{"x", "y", "w", "z"}
	lettersHi := []string{"A", "B", "C", "D"}
	integers := []int{1, 2, 3, 5, 8}
	integersInvalid := []int{4, 6, 7, 10, 11}
	empty := make([]string, 5)
	tests := map[string]arguments{
		"low": {strs: lettersLow, numbers: integers, cond: false,
			expRess: lettersLow, expErr: ""},
		"upper": {strs: lettersLow, numbers: integers, cond: true,
			expRess: lettersHi, expErr: ""},
		"lowErr": {strs: lettersLow, numbers: integersInvalid, cond: true,
			expRess: empty, expErr: "invalid s"},
		"upperErr": {strs: lettersAny, numbers: integers, cond: true,
			expRess: empty, expErr: "invalid s"},
	}
	for name, tt := range tests {
		check(t, name, tt)
	}
}

func check(t *testing.T, name string, tt arguments) {
	t.Run(name, func(t *testing.T) {
		for i := 0; i < len(tt.strs); i++ {
			for _, tw := range tt.numbers {
				actual, err := Do(tt.strs[i], tw, tt.cond)
				if tt.expErr != "" {
					assert.EqualError(t, err, tt.expErr)
				} else {
					assert.NoError(t, err)
				}
				l := len(actual)
				assert.GreaterOrEqual(t, 2, l)
				assert.Equal(t, tt.expRess[i], actual)
			}
		}
	})
}
