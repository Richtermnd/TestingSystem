package tests

import (
	"reflect"
	"slices"
	"strings"
)

// Question of Test
type Question struct {
	N             int      `json:"n"`             // Number of question
	Text          string   `json:"text"`          // Text of question.
	Variants      []string `json:"variants"`      // Variants of anwers.
	Answers       []string `json:"answers"`       // Correct answers.
	IsMultiAnswer bool     `json:"isMultiAnswer"` // Is question a multianswer. It's usefull for frontend.
}

// Test
type Test struct {
	Title     string     `json:"title"`     // Test title.
	AuthorId  string     `json:"authorID"`  // Author id.
	Questions []Question `json:"questions"` // Questions slice.
}

// Return all answers for test.
func (t *Test) GetAnswers() *QuestionAnswers {
	// Create copy original questions slice.
	questionsCopy := make([]Question, len(t.Questions))
	copy(questionsCopy, t.Questions)

	// Get answers from questions.
	answers := make([]*QuestionAnswer, len(questionsCopy))
	for i, question := range questionsCopy {
		answers[i] = &QuestionAnswer{
			N:       question.N,
			Answers: question.Answers,
		}
	}

	// Convert answers to QuestionAnswers.
	// return QuestionAnswers(answers) not work.
	res := make(QuestionAnswers, len(answers))
	copy(res, answers)
	return &res
}

type QuestionAnswer struct {
	N       int
	Answers []string
}

type QuestionAnswers []*QuestionAnswer

// Sort by nubmer
// !!! Change source
func (a *QuestionAnswers) Sort() *QuestionAnswers {
	slices.SortFunc(*a, func(a, b *QuestionAnswer) int {
		if a.N < b.N {
			return -1
		}
		if a.N > b.N {
			return 1
		}
		return 0
	})
	// Sorting answers
	for _, answer := range *a {
		slices.SortFunc(answer.Answers, strings.Compare)
	}
	return a
}

func (a *QuestionAnswers) Check(test Test) bool {
	a.Sort()
	correctAnswers := test.GetAnswers().Sort()
	return reflect.DeepEqual(a, correctAnswers)
}
