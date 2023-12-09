package test

import (
	"testing"

	"github.com/Richtermnd/TestingSystem/testingSystem"
)

func TestCheck(t *testing.T) {
	// Most cursed thing that I write in Go.
	testCases := []struct {
		Name         string
		Test         tests.Test
		InputAnswers tests.QuestionAnswers
		Expected     bool
	}{
		{
			Name: "Correct test",
			Test: tests.Test{
				Title:    "Test1",
				AuthorId: "",
				Questions: []tests.Question{
					{
						N:             1,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a"},
						IsMultiAnswer: false,
					},
					{
						N:             2,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a", "b"},
						IsMultiAnswer: true,
					},
				},
			},
			InputAnswers: tests.QuestionAnswers{
				&tests.QuestionAnswer{
					N:       1,
					Answers: []string{"a"},
				},
				&tests.QuestionAnswer{
					N:       2,
					Answers: []string{"a", "b"},
				},
			},
			Expected: true,
		},
		{
			Name: "Uncorrect test",
			Test: tests.Test{
				Title:    "Test1",
				AuthorId: "",
				Questions: []tests.Question{
					{
						N:             1,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a"},
						IsMultiAnswer: false,
					},
					{
						N:             2,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a", "b"},
						IsMultiAnswer: true,
					},
				},
			},
			InputAnswers: tests.QuestionAnswers{
				&tests.QuestionAnswer{
					N:       1,
					Answers: []string{"b"},
				},
				&tests.QuestionAnswer{
					N:       2,
					Answers: []string{"a", "c"},
				},
			},
			Expected: false,
		},
		{
			Name: "Unordered answers test",
			Test: tests.Test{
				Title:    "Test1",
				AuthorId: "",
				Questions: []tests.Question{
					{
						N:             1,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a"},
						IsMultiAnswer: false,
					},
					{
						N:             2,
						Text:          "",
						Variants:      []string{"a", "b", "c", "d"},
						Answers:       []string{"a", "b"},
						IsMultiAnswer: true,
					},
				},
			},
			InputAnswers: tests.QuestionAnswers{
				&tests.QuestionAnswer{
					N:       2,
					Answers: []string{"b", "a"},
				},
				&tests.QuestionAnswer{
					N:       1,
					Answers: []string{"a"},
				},
			},
			Expected: true,
		},
	}
	for _, test := range testCases {
		res := test.InputAnswers.Check(&test.Test)
		if res != test.Expected {
			t.Errorf("%s failed", test.Name)
		}
	}
}
