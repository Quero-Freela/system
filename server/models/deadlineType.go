package models

type DeadlineType string

const (
	Days   DeadlineType = "days"
	Hours  DeadlineType = "hours"
	Months DeadlineType = "months"
)
