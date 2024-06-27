package command

import (
	"fmt"
	"os/exec"
)

func ApplyResource(content []byte) {
	// prepare command
	cmd := exec.Command("kubectl", "apply", "-f", "-")
	stdin, _ := cmd.StdinPipe()
	cmd.Start()

	// append file content to command
	// TODO: checken op command injection
	stdin.Write(content)
	stdin.Close()

	// print eventual errors or responses
	//var out bytes.Buffer
	//var stderr bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &stderr

	err := cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
