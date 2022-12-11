package dao

import "time"

type ConstInfo struct {
	Id        int
	Creator   string
	Oj        string
	Title     string
	Level     int
	Des       string
	Note      string
	BeginTime time.Time
	LastTime  int
	Type      string
	Auth      int
	CloneFrom int
	Classes   string
	IpRange   string
	LockBoard int
	LockTime  int
}

type ContestProblem struct {
	Id           int
	ContestId    int
	ProblemId    string
	ProblemTitle string
	Rank         int
}
