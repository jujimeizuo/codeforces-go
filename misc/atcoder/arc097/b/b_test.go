// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc097/submit?taskScreenName=arc097_b
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`5 2
5 3 1 4 2
1 3
5 4`,
			`2`,
		},
		{
			`3 2
3 2 1
1 2
2 3`,
			`3`,
		},
		{
			`10 8
5 3 6 8 7 10 9 1 2 4
3 1
4 1
5 9
2 5
6 5
3 5
8 9
7 9`,
			`8`,
		},
		{
			`5 1
1 2 3 4 5
1 5`,
			`5`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc097/tasks/arc097_b
