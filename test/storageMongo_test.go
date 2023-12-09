package test

import (
	"log/slog"
	"os"
	"reflect"
	"strconv"
	"testing"

	storage "github.com/Richtermnd/TestingSystem/storage/mongodb"
	"github.com/Richtermnd/TestingSystem/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestStructure struct {
	ID      *primitive.ObjectID `bson:"_id,omitempty"`
	Integer int                 `bson:"integer"`
	String  string              `bson:"string"`
	List    []int               `bson:"list"`
}

var testStorage *storage.Storage[TestStructure]
var createdIDs []string
var createdItems []*TestStructure

// Init storage
func TestStorageInit(t *testing.T) {
	// Change dir
	os.Chdir("..")
	// logger
	log := slog.Default()
	// Setting and loading test enviroment
	t.Setenv("TESTING_SYSTEM_ENV", "test")
	utils.LoadEnviroment(log)

	storage.Init(log)
	testStorage = storage.NewStorage[TestStructure]("test", log)
}

func TestCreate(t *testing.T) {
	for i := 0; i < 5; i++ {
		item := TestStructure{
			Integer: i,
			String:  strconv.Itoa(i),
			List:    []int{i, i, i},
		}
		createdItems = append(createdItems, &item)
		id, err := testStorage.Create(&item)
		if err != nil {
			t.Fatal(err)
		}
		createdIDs = append(createdIDs, id)
	}
}

func TestRead(t *testing.T) {
	for i, id := range createdIDs {
		item, err := testStorage.Read(id)
		if err != nil {
			t.Fatal(err)
		}
		item.ID = nil
		if !reflect.DeepEqual(item, createdItems[i]) {
			t.Fatal("Items not equal")
		}
	}
}

func TestReadAll(t *testing.T) {
	items, err := testStorage.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range items {
		item.ID = nil
	}
	if !reflect.DeepEqual(items, createdItems) {
		t.Fatal("Items not equal")
	}
}

func TestUpdate(t *testing.T) {
	for _, id := range createdIDs {
		oldItem, err := testStorage.Read(id)
		if err != nil {
			t.Fatal(err)
		}
		oldItem.String += "xd"

		newId, err := testStorage.Update(id, oldItem)
		if err != nil {
			t.Fatal(err)
		}
		updatedItem, err := testStorage.Read(newId)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(oldItem, updatedItem) {
			t.Fatal("Item not updated")
		}
	}
}

func TestDeleteCorrectItems(t *testing.T) {
	for _, id := range createdIDs {
		res, err := testStorage.Delete(id)
		if err != nil {
			t.Fatal(err)
		}
		if !res {
			t.Fatal("Item not deleted")
		}
	}
}

func TestDeleteUncorrectitems(t *testing.T) {
	fakeId := primitive.NewObjectID()
	res, err := testStorage.Delete(fakeId.Hex())
	if err != nil {
		t.Fatal(err)
	}
	if res {
		t.Fatal("Succesfull delete of fake item")
	}
}
