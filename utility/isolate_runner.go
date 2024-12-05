package utility

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	OK      = 1
	TIMEOUT = 2
	DIE     = -99
	SYS_DIE = -98
)

func LookUpMeta(isolateID int) (int, string) {
	err_status := ""
	found_ele := false
	dat, err := os.Open(fmt.Sprintf("/home/YOUR_USER/go_grader/runner/isolate_logs/%d/meta-log.txt", isolateID))
	if err != nil {
		return -99, "LOGS_PATH_NOT_EXIST"
	}
	contentScanner := bufio.NewScanner(dat)
	contentScanner.Split(bufio.ScanLines)
	for contentScanner.Scan() {
		if strings.Contains(contentScanner.Text(), "status:") {
			err_status = string(contentScanner.Text()[7:])
			found_ele = true
		}
	}
	dat.Close()

	if found_ele {
		if err_status == "TO" {
			return TIMEOUT, "TIME_OUT"
		} else if err_status == "SG" {
			return DIE, "DIE"
		} else if err_status == "XX" {
			return SYS_DIE, "Internal Error"
		}
	} else {
		err_status = "OK"
		return 1, err_status
	}
	return -98, "???"
}

func InitalIsolate(isolateID int) int {
	// Isolate inital
	cmd := exec.Command("isolate",
		"--init",
		fmt.Sprintf("--box-id=%d", isolateID),
	)

	// Make logs directory
	cmd_makeLog := exec.Command("mkdir",
		fmt.Sprintf("/home/YOUR_USER/go_grader/runner/isolate_logs/%d/", isolateID),
	)

	err := cmd.Run()
	err_makeLog := cmd_makeLog.Run()
	if err != nil {
		fmt.Println("ISOLATE DIED! :: Please try again later!")
		return 1
	}

	if err_makeLog != nil {
		// TODO: Change new
		fmt.Println("WRONG CONFIG! :: Please check your config!")
		return 1
	}

	fmt.Printf("[Isolate] :: Isolate (ID:%d) has been initalized!\n", isolateID)
	return 0
}

func VerifyResult(questID int, testcaseNo int, output string) bool {
	// Lookup output file
	dat, err := os.ReadFile(fmt.Sprintf("/home/YOUR_USER/go_grader/problem/%d/%d.out", questID, testcaseNo))
	if err != nil {
		panic(err)
	}

	// TODO: Work in progess manage string
	return string(dat) == output
}

func RunnerIsolate(isolateID int, questID int) (bool, int, int, string) {
	// Prepare the command
	var note string = ""
	var stats bool = true
	var getScore int = 0
	for i := 1; i <= 3; i++ {
		cmd := exec.Command("isolate",
			"--run",
			fmt.Sprintf("--box-id=%d", isolateID),
			"--time=1",
			"--mem=65536",
			"--processes=1",
			fmt.Sprintf("--dir=/home/YOUR_USER/go_grader/problem/%d/:rw", questID),
			fmt.Sprintf("--stdin=/home/YOUR_USER/go_grader/problem/%d/%d.in", questID, i),
			fmt.Sprintf("--meta=/home/YOUR_USER/go_grader/runner/isolate_logs/%d/meta-log.txt", isolateID),
			"./out.a",
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
				return false, 0, 0, "Internal Server Error"
			}
			stats = false
		} else {
			if VerifyResult(questID, i, stdout.String()) == true {
				note = note + "P"
				getScore++
			} else {
				note = note + "-"
				stats = false
			}
		}
	}
	return stats, getScore, 3, note
}
