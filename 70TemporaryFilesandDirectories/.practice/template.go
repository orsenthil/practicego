// Throughout program execution, we often want to create
// data that isn't needed after the program exits.
// *Temporary files and directories* are useful for this
// purpose since they don't pollute the file system over
// time.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)


// TODO: Create function check(e error) that checks if e is not nil, panic with e

func main() {

	// The easiest way to create a temporary file is by
	// calling `os.CreateTemp`. It creates a file *and*
	// opens it for reading and writing. We provide `""`
	// as the first argument, so `os.CreateTemp` will
	// create the file in the default location for our OS.


	// TODO: Create f, err := os.CreateTemp("", "sample")
	// TODO: Print err

	// Display the name of the temporary file. On
	// Unix-based OSes the directory will likely be `/tmp`.
	// The file name starts with the prefix given as the
	// second argument to `os.CreateTemp` and the rest
	// is chosen automatically to ensure that concurrent
	// calls will always create different file names.

	// TODO: Print fmt.Println("Temp file name:", f.Name())

	// Clean up the file after we're done. The OS is
	// likely to clean up temporary files by itself after
	// some time, but it's good practice to do this
	// explicitly.

	// TODO: Defer the removal of the file with defer os.Remove(f.Name())

	// We can write some data to the file.

	// TODO: Create _, err = f.Write([]byte{1, 2, 3, 4})
	// TODO: Print err

	// If we intend to write many temporary files, we may
	// prefer to create a temporary *directory*.
	// `os.MkdirTemp`'s arguments are the same as
	// `CreateTemp`'s, but it returns a directory *name*
	// rather than an open file.

	// TODO: Create dname, err := os.MkdirTemp("", "sampledir")
	// TODO: Print err

	// TODO: Defer the removal of the directory with defer os.RemoveAll(dname)

	// Now we can synthesize temporary file names by
	// prefixing them with our temporary directory.

	// TODO: Create fname := filepath.Join(dname, "file1")
	// TODO: Create err = os.WriteFile(fname, []byte{1, 2}, 0666)
	// TODO: Print err
}