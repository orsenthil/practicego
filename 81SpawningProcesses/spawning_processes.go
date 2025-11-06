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

	dateOut, err := dateCmd.Output()
	fmt.Println("err:", err)
	fmt.Println("dateOut:", string(dateOut))
	

	// `Output` and other methods of `Command` will return
	// `*exec.Error` if there was a problem executing the
	// command (e.g. wrong path), and `*exec.ExitError`
	// if the command ran but exited with a non-zero return
	// code.


	_, err = exec.Command("date", "-x").Output()
	fmt.Println("err:", err)
	

	// Next we'll look at a slightly more involved case
	// where we pipe data to the external process on its
	// `stdin` and collect the results from its `stdout`.

	grepCmd := exec.Command("grep", "hello")

	// Here we explicitly grab input/output pipes, start
	// the process, write some input to it, read the
	// resulting output, and finally wait for the process
	// to exit.

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\n"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	fmt.Println("grepBytes:", string(grepBytes))
	grepCmd.Wait()


	// We omitted error checks in the above example, but
	// you could use the usual `if err != nil` pattern for
	// all of them. We also only collect the `StdoutPipe`
	// results, but you could collect the `StderrPipe` in
	// exactly the same way.

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Note that when spawning commands we need to
	// provide an explicitly delineated command and
	// argument array, vs. being able to just pass in one
	// command-line string. If you want to spawn a full
	// command with a string, you can use `bash`'s `-c`
	// option:

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	fmt.Println("err:", err)
	fmt.Println("lsOut:", string(lsOut))
}