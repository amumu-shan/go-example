package proc

import (
	"fmt"
	"os"
	"os/exec"
)

func getProcess1() {
	env := os.Environ()
	attr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	process, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, attr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("the process id is &v", process)

}

func getProcess2() {
	cmd := exec.Command("ls", "-l")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!\n", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v\n", cmd)
}
