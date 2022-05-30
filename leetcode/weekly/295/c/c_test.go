// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_c(t *testing.T) {
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithFile(t, totalSteps, "c.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/weekly-contest-295/problems/steps-to-make-array-non-decreasing/

func TestCompareInf(t *testing.T) {
	testutil.DebugTLE = 0

	inputGenerator := func() (a []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 5)
		a = rg.IntSlice(n, 1, 5)
		return
	}

	testutil.CompareInf(t, inputGenerator, totalSteps, totalStepsWA)
}

type seg []struct {
	l, r int
	val  int
}

func (t seg) set(o int, val int) {
	t[o].val = val
}

func (t seg) op(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.op(lo.val, ro.val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i int, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) (res int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	lv := t.query(o<<1, l, r)
	rv := t.query(o<<1|1, l, r)
	return t.op(lv, rv)
}

func (t seg) queryAll() int { return t[1].val }

func newSegmentTree(a []int) seg {
	t := make(seg, len(a)*4)
	t.build(a, 1, 1, len(a))
	return t
}

func totalStepsWA(a []int) (ans int) {
	for i, n := 0, len(a); i < n; {
		st := i
		v := a[st]
		for i++; i < n && a[i] < v; i++ {
		}

		b := a[st:i]
		n := len(b)

		left := make([]int, n)
		const border int = 2e9
		type pair struct{ v, i int }
		stack := []pair{{border, -1}}
		for i, v := range b {
			for {
				if top := stack[len(stack)-1]; top.v > v {
					left[i] = top.i
					break
				}
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, pair{v, i})
		}
		s := newSegmentTree(make([]int, n))
		mxcnt := 0
		cnt := 1
		for i := 1; i < n; i++ {
			if b[i] < b[i-1] {
				cnt = 1
			} else {
				cnt++
				if left[i] != left[i-1] {
					mx := s.query(1, left[i]+1, i)
					cnt = max(cnt, mx+1)
				}
			}
			s.update(1,i,cnt)
			mxcnt = max(mxcnt, cnt)
		}
		ans = max(ans, mxcnt)

	}
	return
}
