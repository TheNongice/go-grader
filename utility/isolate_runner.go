package utility

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"

	compiler "github.com/TheNongice/go-grader/utility/compiler"
)

const (
	OK      = 1
	TIMEOUT = 2
	DIE     = -99
	SYS_DIE = -98
)

func InitalIsolate(isolateID int) int {
	// Isolate inital
	cmd := exec.Command("isolate",
		"--init",
		fmt.Sprintf("--box-id=%d", isolateID),
	)

	// Make logs directory
	cmd_makeLog := exec.Command("mkdir",
		fmt.Sprintf("%srunner/isolate_logs/%d/", os.Getenv("DIR_GRADER_PATH"), isolateID),
	)

	err := cmd.Run()
	err_makeLog := cmd_makeLog.Run()
	if err != nil {
		fmt.Println("ISOLATE DIED! :: Please try again later!")
		return 1
	}
	if err_makeLog != nil {
		fmt.Println("WRONG CONFIG! :: Please check your config!")
		return 1
	}

	fmt.Printf("[Isolate] :: Isolate (ID:%d) has been initalized!\n", isolateID)
	return 0
}

func CompileCode(lang string, isolateID int, codeContent string) (bool, int, string) {
	tempID := rand.IntN(100)
	fileName := ""
	folderName := ""
	if lang == "cpp" {
		fileName = fmt.Sprintf("main%d.cpp", tempID)
		folderName = "cpp"
	} else if lang == "go" {
		fileName = fmt.Sprintf("main%d.go", tempID)
		folderName = "golang"
	}

	file, err := os.Create(fmt.Sprintf("./runner/temp_code/%s/%s", folderName, fileName))
	file.Write([]byte(codeContent))

	if err != nil {
		fmt.Println("Error with create file!")
	}

	compiler := compiler.DefaultCompiler{}
	status, err_msg := compiler.Compile(lang, tempID, isolateID)
	defer file.Close()

	return status, tempID, err_msg
}

func RunnerIsolate(isolateID int, boxFile int, questID int, max_time float32, max_mem int) (bool, int, int, string, error) {
	// Prepare the command
	var note string = ""
	var stats bool = true
	var getScore int = 0
	prob_name, amount_tc, err := AutoloadProblem(questID)
	if err != nil {
		return false, 0, 0, "System Error!", errors.New("system error")
	}

	fmt.Printf("[JUDGING]: %s from box_id:%d has been call!\n", prob_name, isolateID)
	for i := 1; i <= amount_tc; i++ {
		cmd := exec.Command("isolate",
			"--run",
			fmt.Sprintf("--box-id=%d", isolateID),
			fmt.Sprintf("--time=%v", max_time),
			fmt.Sprintf("--mem=%v", max_mem),
			"--processes=1",
			fmt.Sprintf("--dir=%sproblem/%d/:rw", os.Getenv("DIR_GRADER_PATH"), questID),
			fmt.Sprintf("--stdin=%sproblem/%d/%d.in", os.Getenv("DIR_GRADER_PATH"), questID, i),
			fmt.Sprintf("--meta=%srunner/isolate_logs/%d/meta-log.txt", os.Getenv("DIR_GRADER_PATH"), isolateID),
			fmt.Sprintf("./out%d.a", boxFile),
		)
		
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run the command
		err := cmd.Run()

		// Handle errors or output
		if err != nil {
			status, txt := LookUpMeta(isolateID)
			if status == TIMEOUT {
				note = note + "T"
			} else if status == DIE {
				note = note + "X"
			} else if status == SYS_DIE {
				note = note + "!"
			} else if txt == "LOGS_PATH_NOT_EXIST" {
				return false, 0, 0, "Internal Server Error", errors.New("internal server error")
			}
			stats = false
		} else {
			if VerifyResult(questID, i, stdout.String()) {
				note = note + "P"
				getScore++
			} else {
				note = note + "-"
				stats = false
			}
		}
	}

	return stats, getScore, amount_tc, note, nil
}
