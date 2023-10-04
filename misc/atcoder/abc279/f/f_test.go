// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc279/tasks/abc279_f
// 提交：https://atcoder.jp/contests/abc279/submit?taskScreenName=abc279_f
// 对拍：https://atcoder.jp/contests/abc279/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc279_f&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [f]")
	testCases := [][2]string{
		{
			`5 10
3 5
1 1 4
2 1
2 4
3 7
1 3 1
3 4
1 1 4
3 7
3 6`,
			`5
4
3
1
3`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
