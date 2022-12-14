// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc132/submit?taskScreenName=arc132_c
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`4 2
3 -1 1 -1`,
			`2`,
		},
		{
			`5 1
2 3 4 5 -1`,
			`0`,
		},
		{
			`16 5
-1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1 -1`,
			`794673086`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc132/tasks/arc132_c
