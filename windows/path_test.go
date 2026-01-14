package windows_test

import (
	"testing"

	"github.com/NextronSystems/universalpath/windows"
)

type stringTest struct {
	input  string
	output string
}

func TestBase(t *testing.T) {
	for _, tc := range []stringTest{
		{input: `C:\foo\bar\baz.txt`, output: "baz.txt"},
		{input: `foo\bar\baz.txt`, output: "baz.txt"},
		{input: `baz.txt`, output: "baz.txt"},
		{input: `\\.\pipe\baz.txt`, output: "baz.txt"},
		{input: ".", output: "."},
		{input: "..", output: ".."},
		{input: "/", output: "\\"},
		{input: "", output: "."},
	} {
		result := windows.Base(tc.input)
		if result != tc.output {
			t.Errorf("Windows.Base(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestDir(t *testing.T) {
	for _, tc := range []stringTest{
		{input: `C:\foo\bar\baz.txt`, output: `C:\foo\bar`},
		{input: `foo\bar\baz.txt`, output: `foo\bar`},
		{input: `baz.txt`, output: `.`},
		{input: ".", output: "."},
		{input: "..", output: "."},
		{input: "C:\\", output: "C:\\"},
		{input: "", output: "."},
	} {
		result := windows.Dir(tc.input)
		if result != tc.output {
			t.Errorf("Windows.Dir(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestJoin(t *testing.T) {
	for _, tc := range []struct {
		parts  []string
		output string
	}{
		{parts: []string{`C:\foo`, "bar", "baz.txt"}, output: `C:\foo\bar\baz.txt`},
		{parts: []string{`foo`, "bar", "baz.txt"}, output: `foo\bar\baz.txt`},
		{parts: []string{`baz.txt`}, output: `baz.txt`},
		{parts: []string{`C:\`, "foo", "..", "baz.txt"}, output: `C:\baz.txt`},
		{parts: []string{`C:\`, "..", "baz.txt"}, output: `C:\baz.txt`},
		{parts: []string{"..", "baz.txt"}, output: "..\\baz.txt"},
	} {
		result := windows.Join(tc.parts...)
		if result != tc.output {
			t.Errorf("Windows.Join(%q) = %q; want %q", tc.parts, result, tc.output)
		}
	}
}

func TestExt(t *testing.T) {
	for _, tc := range []stringTest{
		{input: `C:\foo\bar\baz.txt`, output: ".txt"},
		{input: `foo\bar\baz.tar.gz`, output: ".gz"},
		{input: `baz`, output: ""},
		{input: `\baz.`, output: "."},
	} {
		result := windows.Ext(tc.input)
		if result != tc.output {
			t.Errorf("Windows.Ext(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		input string
		dir   string
		base  string
	}{
		{input: `C:\foo\bar\baz.txt`, dir: `C:\foo\bar\`, base: "baz.txt"},
		{input: `foo\bar\baz.txt`, dir: `foo\bar\`, base: "baz.txt"},
		{input: `baz.txt`, dir: ``, base: "baz.txt"},
		{input: `\\.\pipe\baz.txt`, dir: `\\.\pipe\`, base: "baz.txt"},
		{input: `\\network\path\baz.txt`, dir: `\\network\path\`, base: "baz.txt"},
	} {
		dir, base := windows.Split(tc.input)
		if dir != tc.dir || base != tc.base {
			t.Errorf("Windows.Split(%q) = (%q, %q); want (%q, %q)", tc.input, dir, base, tc.dir, tc.base)
		}
	}
}

func TestClean(t *testing.T) {
	for _, tc := range []stringTest{
		{input: `C:\foo\..\bar\baz.txt`, output: `C:\bar\baz.txt`},
		{input: `foo\..\bar\baz.txt`, output: `bar\baz.txt`},
		{input: `.\baz.txt`, output: `baz.txt`},
		{input: `C:\foo\.\bar\baz.txt`, output: `C:\foo\bar\baz.txt`},
		{input: `C:\foo\\bar\\baz.txt`, output: `C:\foo\bar\baz.txt`},
		{input: `C:\foo\bar\..\..\baz.txt`, output: `C:\baz.txt`},
		{input: `..\baz.txt`, output: `..\baz.txt`},
		{input: `\\network\path\..\baz.txt`, output: `\\network\path\baz.txt`},
	} {
		result := windows.Clean(tc.input)
		if result != tc.output {
			t.Errorf("Windows.Clean(%q) = %q; want %q", tc.input, result, tc.output)
		}
	}
}
