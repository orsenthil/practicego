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
	fmt.Println(err)
	fmt.Println(string(dateOut))
	dateOut, err = exec.Command("date", "-x").Output()
	fmt.Println(err)
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed to execute:", e)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		}
	}

	// `Output` and other methods of `Command` will return
	// `*exec.Error` if there was a problem executing the
	// command (e.g. wrong path), and `*exec.ExitError`
	// if the command ran but exited with a non-zero return
	// code.


	_, err = exec.Command("date", "-x").Output()
	
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed to execute:", e)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		}
	}


	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.

	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.

	// TODO: Create grepIn, _ := grepCmd.StdinPipe()
	grepIn, _ := grepCmd.StdinPipe()
	// TODO: Create grepOut, _ := grepCmd.StdoutPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// Start the process, write some input to it, read the resulting output, and finally wait for the process to exit.

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\n"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))


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