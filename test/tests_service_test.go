package test

import (
	"fmt"
	"log/slog"
	"reflect"
	"testing"

	"github.com/Richtermnd/TestingSystem/internal/service"
	"github.com/Richtermnd/TestingSystem/pkg/tests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var s service.TestService = *service.NewTestService(slog.Default())
var createdIDs []*primitive.ObjectID

func TestCreate(t *testing.T) {
	for i := 0; i < 5; i++ {
		newTest := tests.Test{
			Title:    fmt.Sprintf("%d", i),
			AuthorId: "",
			Questions: []tests.Question{
				{
					N:             1,
					Text:          "Question1",
					Variants:      []string{"a", "b", "c", "d"},
					Answers:       []string{"a"},
					IsMultiAnswer: false,
				},
				{
					N:             2,
					Text:          "Question2",
					Variants:      []string{"a", "b", "c", "d"},
					Answers:       []string{"b"},
					IsMultiAnswer: false,
				},
				{
					N:             3,
					Text:          "Question3",
					Variants:      []string{"a", "b", "c", "d"},
					Answers:       []string{"a", "b"},
					IsMultiAnswer: true,
				},
				{
					N:             4,
					Text:          "Question4",
					Variants:      []string{"a", "b", "c", "d"},
					Answers:       []string{"c"},
					IsMultiAnswer: false,
				},
			},
		}
		id, err := s.Create(newTest)
		if err != nil {
			t.Error(err)
		}
		if _, err := s.ReadOne(id); err != nil {
			t.Error(err)
		}
		createdIDs = append(createdIDs, id)
	}
}

func TestUpdate(t *testing.T) {
	for _, id := range createdIDs {
		t.Logf("Getting test with id %v", id)
		test, err := s.ReadOne(id)
		if err != nil {
			t.Error(err)
			continue
		}
		test.Title += " xd"
		t.Log("Update test")
		updatedId, err := s.Update(id, *test)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log("Get updated test")
		updatedTest, err := s.ReadOne(updatedId)
		if err != nil {
			t.Error(err)
			continue
		}
		if !reflect.DeepEqual(test, updatedTest) {
			t.Error("test not updated")
			continue
		}
	}
}

func TestDelete(t *testing.T) {
	for _, id := range createdIDs {
		res, err := s.Delete(id)
		if err != nil {
			t.Error(err)
		}
		if !res {
			t.Errorf("fail deleting existing test")
		}
	}
	res, err := s.Delete(&primitive.ObjectID{})
	if err != nil {
		t.Error(err)
	}
	if res {
		t.Errorf("succesfull deleting of not existing test")
	}
}
