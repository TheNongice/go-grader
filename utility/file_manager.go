package utility

import (
	"errors"
	"fmt"
	"os"

	"mime/multipart"

	"github.com/evilsocket/islazy/zip"
)

func RemoveProblemDir(questID int) (bool, error) {
	problem_dir := os.Getenv("DIR_GRADER_PATH") + fmt.Sprintf("problem/%d", questID)

	if _, err := os.ReadDir(problem_dir); err != nil {
		return false, errors.New("directory you're request is already deleted")
	}

	if err := os.RemoveAll(problem_dir); err != nil {
		return false, errors.New("unexpected error")
	}

	return true, nil
}

func NewProblemDir(file *multipart.FileHeader, questID int) (bool, error) {
	originalZip := fmt.Sprintf("%srunner/temp_problem/%s", os.Getenv("DIR_GRADER_PATH"), file.Filename)
	destExtract := fmt.Sprintf("%sproblem/%d", os.Getenv("DIR_GRADER_PATH"), questID)

	_, err := zip.Unzip(originalZip, destExtract)
	if err != nil {
		return false, errors.New("can't extract files successfully")
	}

	os.Remove(originalZip)
	return true, nil
}
