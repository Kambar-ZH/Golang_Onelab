package models

import "fmt"

type SubmissionVerdict struct {
	Passed bool
}

func (sv SubmissionVerdict) String() string {
	if sv.Passed {
		return "Passed all testcases!"
	}
	return "Failed on test!"
}

type User struct {
	Name    string
	Surname string
	Handle  string
	Age     int
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s\nSurname: %s\nHandle: %s\nAge: %d\n", u.Name, u.Surname, u.Handle, u.Age)
}

type Test struct {
	Input  string
	Output string
}

func (t Test) String() string {
	return fmt.Sprintf("Input: %s\nOutput: %s\n", t.Input, t.Output)
}

type Problem struct {
	Id          int
	Name        string
	Description string
	Test        Test
}

func (p Problem) String() string {
	return fmt.Sprintf("Id: %d\nName: %s\nDescription: %s\nTest: %v\n", p.Id, p.Name, p.Description, p.Test)
}

type Program struct {
	Code string
}

func (p Program) String() string {
	return fmt.Sprintf("Code: \n%s\n", p.Code)
}

type Submission struct {
	Id                int
	User              User
	SubmissionVerdict SubmissionVerdict
	Problem           Problem
	Program           Program
}

func (s Submission) String() string {
	return fmt.Sprintf("Id: %d\nUser: %v\nSubmission Verdict: %v\nProblem: %v\nProgram: %v\n", s.Id, s.User, s.SubmissionVerdict, s.Problem, s.Program)
}
