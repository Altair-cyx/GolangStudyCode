package split

import (
	"reflect"
	"testing"
)

// go test
// go test -cover
// go test -cover -coverprofile=...
// go tool cover -html=...

// 测试
// func TestSplit(t *testing.T) {
// 	ret := Split("abcdef", "e")
// 	want := []string{"abcd", "f"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Errorf("want:%v but got:%v\n", want, ret)
// 	}
// }

// 测试组
// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}

// 	testGroup := []testCase{
// 		testCase{"abcdef", "e", []string{"abcd", "f"}},
// 		testCase{"abcef", "bc", []string{"a", "ef"}},
// 	}

// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Fatalf("want:%v but got:%v\n", tc.want, got)
// 		}
// 	}
// }

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{
		"case_1": testCase{"abcdef", "e", []string{"abcd", "f"}},
		"case_2": testCase{"abcef", "bc", []string{"a", "ef"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v but got:%#v\n", tc.want, got)
			}
		})
	}
}

// BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
