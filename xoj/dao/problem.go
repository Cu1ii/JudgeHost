package dao

import "time"

type Problem struct {
	Problem  string `gorm:"primaryKey"`
	Author   string
	AddTime  time.Time
	Oj       string
	Title    string
	Des      string
	Input    string
	Output   string
	Sinput   string
	Soutpuy  string
	Source   string
	Time     int
	Memory   int
	Hint     string
	Auth     int
	Template string
}

type ProblemData struct {
	Problem    string `gorm:"primaryKey"`
	Title      string
	Level      int
	Submission int
	Ac         int
	Mle        int
	Tle        int
	Rte        int
	Pe         int
	Wa         int
	Se         int
	Tag        string
	Score      int
	Auth       int
	Oj         string
}
