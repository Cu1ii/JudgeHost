package dao

type CaseStatus struct {
	Id         int `gorm:"primaryKey"`
	StatusId   int
	Username   string
	Problem    string
	Result     string
	Time       int
	Memory     int
	TestCase   string
	CaseData   string
	OutputData string
	UserOutput string
}
