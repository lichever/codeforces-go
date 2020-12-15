// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`[40,10,20,30]`, 
			`[4,1,2,3]`,
		},
		{
			`[100,100,100]`, 
			`[1,1,1]`,
		},
		{
			`[37,12,28,9,100,56,80,5,12]`, 
			`[5,3,4,2,8,6,7,1,3]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, arrayRankTransform, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-18/problems/rank-transform-of-an-array/
