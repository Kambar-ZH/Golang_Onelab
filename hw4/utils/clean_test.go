package utils

import (
	"fmt"
	"hw4/utils/models"
	"testing"
)

func TestRemoveCyrillicFromStruct(t *testing.T) {
	tests := struct {
		submission []models.Submission
	}{
		[]models.Submission{
			{
				Id: 1,
				User: models.User{
					Name:    "Kambar(имя на русском Камбар)",
					Surname: "Zhamauov",
					Handle:  "Kambar_Z(мой телеграм Kambarych)",
					Age:     19,
				},
				SubmissionVerdict: models.SubmissionVerdict{
					Passed: true,
				},
				Problem: models.Problem{
					Id:          10,
					Name:        "Легкая вресия задачи Fibonacci",
					Description: "This problem (эта проблема) is about Fibonacci numbers (о числе фибоначи)",
					Test: models.Test{
						Input:  "4",
						Output: "5",
					},
				},
				Program: models.Program{
					Code: `package main
		func main() { \\ не забудь закрыть файлы
		print(5)
	}`,
				},
			},
		},
	}
	for _, test := range tests.submission {
		RemoveCyrillicFromStruct(&test)
		got := ShowStruct(&test)
		fmt.Println(got)
		for _, x := range got {
			if ('а' <= x && x <= 'я') || ('А' <= x && x <= 'Я') {
				t.Fatal("Found cirillic rune!")
			}
		}
	}
}
