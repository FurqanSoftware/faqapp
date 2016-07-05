package core

import "fmt"

type ValidationError struct {
	Action  string
	Element string
	Issue   Issue
}

func (e ValidationError) Error() string {
	return fmt.Sprint(e.Action + ": " + e.Element + " is " + string(e.Issue))
}

type Issue string

const (
	IssueTooLong  = "too long"
	IssueTooShort = "too short"
	IssueMissing  = "missing"
	IssueInvalid  = "invalid"
)

type DatabaseError struct {
	Action string
	Base   error
}

func (e DatabaseError) Error() string {
	return fmt.Sprint(e.Action + ": " + e.Base.Error())
}
