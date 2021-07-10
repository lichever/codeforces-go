// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3
1 3 3
1 2 2
1 3 2`,
			`3`,
		},
		{
			`5
6 3 2 5 0
1 2 10
2 3 3
2 4 1
1 5 1`,
			`7`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
