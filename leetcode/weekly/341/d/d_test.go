// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_d(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, minimumTotalPrice, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
	if err := testutil.RunFuncWithRandomInput(t, minimumTotalPrice); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-341/problems/minimize-the-total-price-of-the-trips/
// https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
