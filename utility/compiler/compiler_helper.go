package utility

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Compiler interface {
	Compile(language string, tempID int, isolateID int) (bool, string)
}

type DefaultCompiler struct{}

func (c DefaultCompiler) Compile(language string, tempID int, isolateID int) (bool, string) {
	switch strings.ToLower(language) {
	case "cpp":
		return c.compileCPP(tempID, isolateID)
	default:
		return false, fmt.Sprintf("Unsupported language: %s", language)
	}
}

func (c DefaultCompiler) compileCPP(tempID int, isolateID int) (bool, string) {
	cmd := exec.Command("g++",
		fmt.Sprintf("./runner/temp_code/cpp/main%d.cpp", tempID),
		"-o",
		fmt.Sprintf("./runner/temp_code/cpp/output/out%d.a", tempID),
	)

	cmd_copy := exec.Command("cp",
		fmt.Sprintf("./runner/temp_code/cpp/output/out%d.a", tempID),
		fmt.Sprintf("%s/%d/box/out%d.a", os.Getenv("ISOLATE_PATH"), isolateID, tempID),
	)

	fmt.Println("Command is: ", cmd.Args)
	var _, stderr bytes.Buffer
	cmd.Stderr = &stderr

	err_compile := cmd.Run()
	err_copy := cmd_copy.Run()

	if err_compile != nil || err_copy != nil {
		return false, stderr.String()
	}

	return true, ""
}