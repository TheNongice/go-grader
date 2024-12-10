package utility

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	LF       = 10 // LINEFEED
	EOT      = 3 // END-OF-TEXT like EOF
	SPACEBAR = 20
)

type TestcaseDesc struct {
	ProblemTitle   string `json:"problem_title"`
	MaxTime        int    `json:"max_time"`
	MaxMemory      int    `json:"max_memory"`
	AmountTestcase int    `json:"amount_testcase"`
}

func LookUpMeta(isolateID int) (int, string) {
	err_status := ""
	found_ele := false
	dat, err := os.Open(fmt.Sprintf("%srunner/isolate_logs/%d/meta-log.txt", os.Getenv("DIR_GRADER_PATH"), isolateID))
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
		} else if err_status == "SG" || err_status == "RE" {
			return DIE, "DIE"
		} else if err_status == "XX" {
			return SYS_DIE, "Internal Error"
		}
	} else {
		err_status = "OK"
		return 1, err_status
	}
	return SYS_DIE, "???"
}

func VerifyResult(questID int, testcaseNo int, output string) bool {
	// Lookup output file
	dat, err := os.ReadFile(fmt.Sprintf("%sproblem/%d/%d.out", os.Getenv("DIR_GRADER_PATH"), questID, testcaseNo))
	if err != nil {
		panic(err)
	}
	final_dat := string(dat)
	// Destroy \n as every output
	if output[len(output)-1] == LF {
		output = output[0 : len(output)-1]
	}

	if final_dat[len(final_dat)-1] == LF {
		final_dat = final_dat[0 : len(final_dat)-1]
	}

	return final_dat == output
}

func AutoloadProblem(questID int) (string, int, error) {
	dat, err := os.ReadFile(fmt.Sprintf("%sproblem/%d/desc.json", os.Getenv("DIR_GRADER_PATH"), questID))
	if err != nil {
		panic(err)
	}
	var dat_json TestcaseDesc
	err = json.Unmarshal(dat, &dat_json)
	if err != nil {
		return "", 0, errors.New("can't access json file")
	}

	return dat_json.ProblemTitle, dat_json.AmountTestcase, nil
}
