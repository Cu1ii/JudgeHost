package dao

import "time"

type JudgeStatus struct {
	Id             int64 `gorm:"primaryKey"`
	User           string
	Oj             string
	Problem        string
	Result         int
	Time           int
	Memory         int
	Length         int
	Language       string
	SubmitTime     time.Time
	Judger         string
	Contest        int64
	ContestProblem int
	Code           string
	TestCase       string
	Message        string
	ProblemTitle   string
	Rating         int
	Ip             string
}
