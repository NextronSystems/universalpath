package unix_test

import (
	"testing"

	"github.com/NextronSystems/universalpath/unix"
)

type stringTest struct {
	input  string
	output string
}

func TestBase(t *testing.T) {
	for _, tc := range []stringTest{
		{input: "/foo/bar/baz.txt", output: "baz.txt"},
		{input: "foo/bar/baz.txt", output: "baz.txt"},
		{input: "baz.txt", output: "baz.txt"},
		{input: "/baz.txt", output: "baz.txt"},
		{input: ".", output: "."},
		{input: "..", output: ".."},
		{input: "/", output: "/"},
		{input: "", output: "."},
	} {
		result := unix.Base(tc.input)
		if result != tc.output {
			t.Errorf("Unix.Base(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestDir(t *testing.T) {
	for _, tc := range []stringTest{
		{input: "/foo/bar/baz.txt", output: "/foo/bar"},
		{input: "foo/bar/baz.txt", output: "foo/bar"},
		{input: "baz.txt", output: "."},
		{input: "/baz.txt", output: "/"},
		{input: ".", output: "."},
		{input: "..", output: "."},
		{input: "/", output: "/"},
		{input: "", output: "."},
	} {
		result := unix.Dir(tc.input)
		if result != tc.output {
			t.Errorf("Unix.Dir(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestJoin(t *testing.T) {
	for _, tc := range []struct {
		parts  []string
		output string
	}{
		{parts: []string{"/foo", "bar", "baz.txt"}, output: "/foo/bar/baz.txt"},
		{parts: []string{"foo", "bar", "baz.txt"}, output: "foo/bar/baz.txt"},
		{parts: []string{"baz.txt"}, output: "baz.txt"},
		{parts: []string{"/", "..", "baz.txt"}, output: "/baz.txt"},
		{parts: []string{"/", "foo", "..", "baz.txt"}, output: "/baz.txt"},
		{parts: []string{"/", "foo", "bar", "..", "baz.txt"}, output: "/foo/baz.txt"},
		{parts: []string{"..", "baz.txt"}, output: "../baz.txt"},
	} {
		result := unix.Join(tc.parts...)
		if result != tc.output {
			t.Errorf("Unix.Join(%q) = %q; want %q", tc.parts, result, tc.output)
		}
	}
}

func TestExt(t *testing.T) {
	for _, tc := range []stringTest{
		{input: "/foo/bar/baz.txt", output: ".txt"},
		{input: "foo/bar/baz.tar.gz", output: ".gz"},
		{input: "baz", output: ""},
		{input: "/baz.", output: "."},
	} {
		result := unix.Ext(tc.input)
		if result != tc.output {
			t.Errorf("Unix.Ext(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		input string
		dir   string
		base  string
	}{
		{input: "/foo/bar/baz.txt", dir: "/foo/bar/", base: "baz.txt"},
		{input: "foo/bar/baz.txt", dir: "foo/bar/", base: "baz.txt"},
		{input: "baz.txt", dir: "", base: "baz.txt"},
		{input: "/baz.txt", dir: "/", base: "baz.txt"},
	} {
		dir, base := unix.Split(tc.input)
		if dir != tc.dir || base != tc.base {
			t.Errorf("Unix.Split(%q) = (%q, %q); want (%q, %q)", tc.input, dir, base, tc.dir, tc.base)
		}
	}
}

func TestClean(t *testing.T) {
	for _, tc := range []stringTest{
		{input: "/foo/./bar//baz.txt", output: "/foo/bar/baz.txt"},
		{input: "foo/bar/../baz.txt", output: "foo/baz.txt"},
		{input: "./baz.txt", output: "baz.txt"},
		{input: "/foo/bar/../../baz.txt", output: "/baz.txt"},
		{input: "/../baz.txt", output: "/baz.txt"},
		{input: "foo//bar///baz.txt", output: "foo/bar/baz.txt"},
	} {
		result := unix.Clean(tc.input)
		if result != tc.output {
			t.Errorf("Unix.Clean(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}
