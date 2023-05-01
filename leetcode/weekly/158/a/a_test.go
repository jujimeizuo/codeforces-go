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
			`"RLRRLLRLRL"`, 
			`4`,
		},
		{
			`"RLLLLRRRLR"`, 
			`3`,
		},
		{
			`"LLLLRRRR"`, 
			`1`,
		},
		{
			`"RLRRRLLRLL"`, 
			`2`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, balancedStringSplit, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-158/problems/split-a-string-in-balanced-strings/
