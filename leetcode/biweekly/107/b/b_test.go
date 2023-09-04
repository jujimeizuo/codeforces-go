// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, longestString, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunFuncWithRandomInput(t, longestString); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-107/problems/construct-the-longest-new-string/
// https://leetcode.cn/problems/construct-the-longest-new-string/
