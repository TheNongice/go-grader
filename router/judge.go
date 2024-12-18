package router

import (
	"fmt"

	"github.com/TheNongice/go-grader/utility"
	"github.com/gofiber/fiber/v2"
)

// Client Recieve Struct
type ResultJudge struct {
	Status     bool   `json:"is_accept"`
	Score      int    `json:"score"`
	Fullscore  int    `json:"full_score"`
	Note       string `json:"note"`
	Additional string `json:"additional,omitempty"`
}

type ResultIsolate struct {
	Status int    `json:"status"`
	Note   string `json:"note"`
}

// Server Recieve Struct
type BodyIsolate struct {
	BoxId int `json:"box_id" form:"box_id"`
}

type BodyJudgeRunner struct {
	BoxId       int    `json:"box_id" form:"box_id"`
	QuestID     int    `json:"question_id" form:"question_id"`
	CodeContent string `json:"code" form:"code"`
}

func JudgeService(router fiber.Router) {
	router.Post("/send", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		p := new(BodyJudgeRunner)
		if err := c.BodyParser(&p); err != nil {
			return err
		}

		status, msg := utility.CompileCode(p.BoxId, p.CodeContent)
		if !status {
			return c.JSON(ResultJudge{
				Status:     false,
				Note:       "Complication Error",
				Additional: msg,
			})
		}

		status, score, full_score, note, err := utility.RunnerIsolate(p.BoxId, p.QuestID)
		if err != nil {
			return c.Status(200).JSON(ResultJudge{
				Status: status,
				Note:   note,
			})
		} else {
			return c.Status(200).JSON(ResultJudge{
				Status:    status,
				Score:     score,
				Fullscore: full_score,
				Note:      note,
			})
		}
	})

	router.Post("/summon_isolate", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		p := new(BodyIsolate)
		if err := c.BodyParser(&p); err != nil {
			return err
		}
		status := utility.InitalIsolate(p.BoxId)
		note := ""
		if status == 0 {
			note = fmt.Sprintf("ISOLATE [ID:%d] INIT", p.BoxId)
		} else {
			note = "CAN'T SUMMON ISOLATE"
		}

		status_http := 200
		if status != 0 {
			status_http = 500
		}

		return c.Status(status_http).JSON(ResultIsolate{
			Status: status,
			Note:   note,
		})
	})
}
