package parser

import (
	"fmt"
	"hw4/parser/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUsers(t *testing.T) {
	testCases := []struct {
		parseTag string
		json     []byte
		users    []*models.User
	}{
		{
			parseTag: "json",
			json: []byte(`[
	  {
		"id": 1,
		"address": {
		  "city_id": 5,
		  "street": "Satbayev"
		},
		"age": 20
	  },
	  {
		"id": 1,
		"address": {
		  "city_id": "6",
		  "street": "Al-Farabi"
		},
		"age": "32"
	  }
	]`),
			users: []*models.User{
				{
					ID: 1,
					Address: models.Address{
						CityID: 5,
						Street: "Satbayev",
					},
					Age: 20,
				},
				{
					ID: 1,
					Address: models.Address{
						CityID: 6,
						Street: "Al-Farabi",
					},
					Age: 32,
				},
			},
		},
	}
	for _, testCase := range testCases {
		var users []*models.User = ParseUsers(testCase.json, testCase.parseTag)
		assert.Equal(t, testCase.users, users,
			fmt.Sprintf("Incorrect parsing.\nExpected %v.\nGot: %v", testCase.users, users))
	}
}

func TestParsePosts(t *testing.T) {
	testCases := []struct {
		parseTag string
		json     []byte
		posts    []*models.Post
	}{
		{
			parseTag: "xml",
			json: []byte(`[
	  {
		"id": 1,
		"title": "Codeforces Div2 753 Round",
		"content": "Welcome to the best CF round",
		"editorial": {
			"contest_id": "753",
			"authors": [
				{
					"author_id": "202",
					"name": "Yerzhan",
					"surname": "Ismailov",
					"created_problems_count": 22.0
				},
				{
					"author_id": "250",
					"name": "Yerzhan",
					"surname": "Ismailov",
					"created_problems_count": "10"
				}
			]
		}
	  }
	]`),
			posts: []*models.Post{
				{
					Id: 1,
					Title: "Codeforces Div2 753 Round",
					Content: "Welcome to the best CF round",
					Editorial: models.Editorial {
						ContestID: 753,
						Authors: []models.Author{
							{
								AuthorID: 202,
								Name: "Yerzhan",
								Surname: "Ismailov",
								CreatedProblemsCount: 22,
							},
							{
								AuthorID: 250,
								Name: "Yerzhan",
								Surname: "Ismailov",
								CreatedProblemsCount: 10,
							},
						},
					},
				},
			},
		},
	}
	for _, testCase := range testCases {
		var posts []*models.Post = ParsePosts(testCase.json, testCase.parseTag)
		assert.Equal(t, testCase.posts, posts,
			fmt.Sprintf("Incorrect parsing.\nExpected %v.\nGot: %v", testCase.posts, posts))
	}
}