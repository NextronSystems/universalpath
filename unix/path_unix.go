// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"strings"
)

func Join(elem ...string) string {
	// If there's a bug here, fix the logic in ./path_plan9.go too.
	for i, e := range elem {
		if e != "" {
			return Clean(strings.Join(elem[i:], string(Separator)))
		}
	}
	return ""
}
