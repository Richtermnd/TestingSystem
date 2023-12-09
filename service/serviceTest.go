package service

import (
	"log/slog"

	storage "github.com/Richtermnd/TestingSystem/storage/mongodb"
	tests "github.com/Richtermnd/TestingSystem/testingSystem"
)

type TestService struct {
	log     *slog.Logger
	storage *storage.Storage[tests.Test]
}

func (t *TestService) Create(newTest tests.Test) (string, error) {
	pk, err := t.storage.Create(&newTest)
	// TODO: Error Analyzing to create correct ErrorResponse
	if err != nil {
		t.log.Error(err.Error())
		return "", err
	}
	return pk, nil

}

func (t *TestService) Read(pk string) (*tests.Test, error) {
	test, err := t.storage.Read(pk)
	// TODO: Error Analyzing to create correct ErrorResponse
	if err != nil {
		t.log.Error(err.Error())
		return nil, err
	}
	return test, nil
}

func (t *TestService) ReadAll() ([]*tests.Test, error) {
	allTests, err := t.storage.ReadAll()
	// TODO: Error Analyzing to create correct ErrorResponse
	if err != nil {
		t.log.Error(err.Error())
		return nil, err
	}
	return allTests, nil
}

func (t *TestService) Update(pk string, updatedTest *tests.Test) (string, error) {
	updatedPk, err := t.storage.Update(pk, updatedTest)
	// TODO: Error Analyzing to create correct ErrorResponse
	if err != nil {
		t.log.Error(err.Error())
		return "", err
	}
	return updatedPk, nil
}

func (t *TestService) Delete(pk string) (bool, error) {
	res, err := t.storage.Delete(pk)
	// TODO: Error Analyzing to create correct ErrorResponse
	if err != nil {
		t.log.Error(err.Error())
		return false, err
	}
	return res, nil
}

func NewTestService(log *slog.Logger) *TestService {
	return &TestService{log, storage.NewStorage[tests.Test]("tests", log)}
}
