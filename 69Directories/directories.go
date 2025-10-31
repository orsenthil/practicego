// Go has several useful functions for working with
// *directories* in the file system.

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// TODO: Create function check(e error) that checks if e is not nil, panic with e

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Create a new sub-directory in the current working
	// directory.

	// TODO: Create err := os.Mkdir("subdir", 0755)
	// TODO: Print err
	err := os.Mkdir("subdir", 0755)
	check(err)

	// When creating temporary directories, it's good
	// practice to `defer` their removal. `os.RemoveAll`
	// will delete a whole directory tree (similarly to
	// `rm -rf`).

	// TODO: Defer the removal of the directory with defer os.RemoveAll("subdir")
	defer os.RemoveAll("subdir")

	// Helper function to create a new empty file.

	// TODO: Create createEmptyFile := func(name string) {
	// That gets d := []byte("")
	// and calls check(os.WriteFile(name, d, 0644))

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}
	// TODO: call createEmptyFile("subdir/file1")
	createEmptyFile("subdir/file1")
	// We can create a hierarchy of directories, including
	// parents with `MkdirAll`. This is similar to the
	// command-line `mkdir -p`.

	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	// TODO: Create err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)
	fmt.Println("Created subdir/parent/child directory.")
	// TODO: Print err
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` lists directory contents, returning a
	// slice of `os.DirEntry` objects.

	// TODO: Create c, err := os.ReadDir("subdir/parent")
	// TODO: Print err
	c, err := os.ReadDir("subdir/parent")
	check(err)
	// TODO: For _, entry := range c, print " ", entry.Name(), entry.IsDir()
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` lets us change the current working directory,
	// similarly to `cd`.

	// TODO: Create err = os.Chdir("subdir/parent/child")
	err = os.Chdir("subdir/parent/child")
	check(err)
	// TODO: Print err
	// Now we'll see the contents of `subdir/parent/child`
	// when listing the *current* directory.

	// TODO: Create c, err := os.ReadDir(".")
	// TODO: Print err
	c, err = os.ReadDir(".")
	check(err)
	// TODO: For _, entry := range c, print " ", entry.Name(), entry.IsDir()
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
	fmt.Println()

	// `cd` back to where we started.

	// TODO: err = os.Chdir("../../..")
	err = os.Chdir("../../..")
	check(err)
	// We can also visit a directory *recursively*,
	// including all its sub-directories. `WalkDir` accepts
	// a callback function to handle every file or
	// directory visited.
	
	// TODO: err = filepath.WalkDir("subdir", visit)
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// `visit` is called for every file or directory found
// recursively by `filepath.WalkDir`.

// TODO: Create function visit(path string, d fs.DirEntry, err error) error
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}