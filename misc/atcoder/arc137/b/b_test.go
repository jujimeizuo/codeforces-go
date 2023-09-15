// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc137/tasks/arc137_b
// 提交：https://atcoder.jp/contests/arc137/submit?taskScreenName=arc137_b
// 对拍：https://atcoder.jp/contests/arc137/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc137_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`4
0 1 1 0`,
			`4`,
		},
		{
			`5
0 0 0 0 0`,
			`6`,
		},
		{
			`6
0 1 0 1 0 1`,
			`3`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
