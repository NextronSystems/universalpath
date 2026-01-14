Universal Path Handling
=======================

This library provides functions to handle either Windows or Unix style paths,
regardless of the operating system the code is running on.

The code for this is copied from the excellent Golang Standard Library's `path/filepath` package (which, unfortunately,
is restricted via build tags to the native OS path style). We do not claim any ownership of this code. The `patches`
directory contains the modifications made to the original code to make it work here.

There are two subpackages:

- `windows` handles Windows style paths (e.g. `C:\Program Files\app\file.txt`)
- `unix` handles Unix style paths (e.g. `/usr/local/bin/app/file.txt`)

The main package provides a `Style` type that can be set to either `Windows` or `Unix` and uses the appropriate subpackage
to perform path operations.
