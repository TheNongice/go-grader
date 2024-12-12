package router

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TheNongice/go-grader/utility"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

type BodyProblemDelete struct {
	QuestID int `json:"question_id" form:"question_id"`
}

type ReturnStatus struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason,omitempty"`
}

func ProblemSetService(router fiber.Router) {
	router.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("MONIT_USER"): os.Getenv("MONIT_PASS"),
		},
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It's work!")
	})

	router.Delete("/", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		p := new(BodyProblemDelete)
		if err := c.BodyParser(&p); err != nil {
			return c.Status(400).JSON(ReturnStatus{
				Status:  false,
				Message: "Failed",
				Reason:  "Wrong Body! Please send body with JSON.",
			})
		}

		if _, err := utility.RemoveProblemDir(p.QuestID); err != nil {
			return c.Status(500).JSON(ReturnStatus{
				Status:  false,
				Message: "Failed",
				Reason:  err.Error(),
			})
		}

		return c.Status(200).JSON(ReturnStatus{
			Status:  true,
			Message: "Success",
		})
	})

	router.Post("/", func(c *fiber.Ctx) error {
		// Check requests from api
		file, err := c.FormFile("problem_file")
		if err != nil {
			return c.Status(500).JSON(ReturnStatus{
				Status:  false,
				Message: "File not found.",
			})
		}

		file_id := c.FormValue("question_id")
		if file_id == "" {
			return c.Status(400).JSON(ReturnStatus{
				Status:  false,
				Message: "Question ID not found",
			})
		}

		fileID, err := strconv.Atoi(file_id)
		if err != nil {
			return c.Status(400).JSON(ReturnStatus{
				Status:  false,
				Message: "Question ID is wrong requirement!",
				Reason:  "Please use integer!",
			})
		}

		c.SaveFile(file, fmt.Sprintf("%srunner/temp_problem/%s", os.Getenv("DIR_GRADER_PATH"), file.Filename))
		if _, err = utility.NewProblemDir(file, fileID); err != nil {
			return c.Status(500).JSON(ReturnStatus{
				Status:  true,
				Message: "Upload failed",
				Reason:  err.Error(),
			})
		}
		return c.Status(200).JSON(ReturnStatus{
			Status:  true,
			Message: "Upload Success",
		})
	})
}
