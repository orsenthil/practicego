// The `filepath` package provides functions to parse
// and construct *file paths* in a way that is portable
// between operating systems; `dir/file` on Linux vs.
// `dirile` on Windows, for example.
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` should be used to construct paths in a
	// portable way. It takes any number of arguments
	// and constructs a hierarchical path from them.

	// TODO: Create p := filepath.Join("dir1", "dir2", "filename")
	// TODO: Print p

	// You should always use `Join` instead of
	// concatenating `/`s or `\`s manually. In addition
	// to providing portability, `Join` will also
	// normalize paths by removing superfluous separators
	// and directory changes.

	// TODO: Create fmt.Println(filepath.Join("dir1//", "filename"))
	// TODO: Create fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` and `Base` can be used to split a path to the
	// directory and the file. Alternatively, `Split` will
	// return both in the same call.

	// TODO: Create fmt.Println("Dir(p):", filepath.Dir(p))
	// TODO: Create fmt.Println("Base(p):", filepath.Base(p))

	// We can check whether a path is absolute.

	// TODO: Create fmt.Println(filepath.IsAbs("dir/file"))
	// TODO: Create fmt.Println(filepath.IsAbs("/dir/file"))

	// TODO: Create filename := "config.json"
	// TODO: Print filename


	// Some file names have extensions following a dot. We
	// can split the extension out of such names with `Ext`.


	// TODO: Create ext := filepath.Ext(filename)
	// TODO: Print ext

	// To find the file's name with the extension removed,
	// use `strings.TrimSuffix`.

	// TODO: Create fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` finds a relative path between a *base* and a
	// *target*. It returns an error if the target cannot
	// be made relative to base.

	// TODO: Create rel, err := filepath.Rel("a/b", "a/b/t/file")
	// TODO: Print err
	// TODO: Create rel, err = filepath.Rel("a/b", "a/c/t/file")
	// TODO: Print err
	// TODO: Print rel

}