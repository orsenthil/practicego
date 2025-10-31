// Sometimes our Go programs need to spawn other
// processes.

package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// We'll start with a simple command that takes no
	// arguments or input and just prints something to
	// stdout. The `exec.Command` helper creates an object
	// to represent this external process.

	// TODO: Create dateCmd := exec.Command("date")
	dateCmd := exec.Command("date")

	// The `Output` method runs the command, waits for it
	// to finish and collects its standard output.
	//  If there were no errors, `dateOut` will hold bytes
	// with the date info.

	// TODO: Create dateOut, err := dateCmd.Output()
	// TODO: Print err
	// TODO: Print dateOut
	dateOut, err := dateCmd.Output()
	fmt.Println(err)
	fmt.Println(dateOut)
	

	// `Output` and other methods of `Command` will return
	// `*exec.Error` if there was a problem executing the
	// command (e.g. wrong path), and `*exec.ExitError`
	// if the command ran but exited with a non-zero return
	// code.


	// TODO: Create _, err = exec.Command("date", "-x").Output()
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		fmt.Println(err)
	}
	
	// TODO: Create if err != nil {
	// Create var execErr *exec.Error and var exitErr *exec.ExitError
	// With switch, check if err is an execErr or exitErr
	// If it is, print the error
	// If it is not, panic with the error

	switch err := err.(type) {
	case *exec.Error:
		fmt.Println("exec: ", err)
	case *exec.ExitError:
		fmt.Println("exit: ", err)
	default:
		panic(err)
	}

	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.

	// TODO: Create grepCmd := exec.Command("grep", "hello")
	grepCmd := exec.Command("grep", "hello")
	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.

	// TODO: Create grepIn, _ := grepCmd.StdinPipe()
	// TODO: Create grepOut, _ := grepCmd.StdoutPipe()
	// Start the process, write some input to it, read the resulting output, and finally wait for the process to exit.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\n"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	fmt.Println(string(grepBytes))
	grepCmd.Wait()


	// We omitted error checks in the above example, but
	// you could use the usual `if err != nil` pattern for
	// all of them. We also only collect the `StdoutPipe`
	// results, but you could collect the `StderrPipe` in
	// exactly the same way.

	// TODO: Print "> grep hello" and the result of grepBytes
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	// Note that when spawning commands we need to
	// provide an explicitly delineated command and
	// argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full
	// command with a string, you can use `bash`'s `-c`
	// option:

	// TODO: Create lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	// Call lsCmd.Output() and print the result 
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, _ := lsCmd.Output()
	fmt.Println(string(lsOut))
}