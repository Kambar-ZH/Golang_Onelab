package models

import "fmt"

type Post struct {
	Id        int       `xml:"id"`
	Title     string    `xml:"title"`
	Content   string    `xml:"content"`
	Editorial Editorial `xml:"editorial"`
}

func (p *Post) String() string {
	return fmt.Sprintf("ID: %d, Title: %v, Content: %s, Editorial %v.", p.Id, p.Title, p.Content, p.Editorial)
}

type Editorial struct {
	ContestID int      `xml:"contest_id"`
	Authors   []Author `xml:"authors"`
}

func (e Editorial) String() string {
	return fmt.Sprintf("ContestID: %d, Authors: %v", e.ContestID, e.Authors)
}

type Author struct {
	AuthorID             int    `xml:"author_id"`
	Name                 string `xml:"name"`
	Surname              string `xml:"surname"`
	CreatedProblemsCount int    `xml:"created_problems_count"`
}

func (a Author) String() string {
	return fmt.Sprintf("AuthorID: %d, Name: %v, Surname: %s, CreatedProblemsCount: %d", a.AuthorID, a.Name, a.Surname, a.CreatedProblemsCount)
}