// run

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test internal print routines that are generated
// by the print builtin.  This test is not exhaustive,
// we're just checking that the formatting is correct.

package main

import "fmt"

var a = 1
var c = 2

func main() {
	fmt.Println("hello world!")
	go func(){
		var b = 1
		_ = a+b+c
	}()
}
