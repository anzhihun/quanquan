package adapter

import (
	"encoding/json"
	"entity"
	"fmt"
	"github.com/fatih/structs"
	"storage"
)

const (
	COL_NAME_USER = "users"
	COL_NAME_CHAN = "chans"
	COL_NAME_MSG  = "msgs"
)

func InitStorage() {
	storage.TiedotDB.AddCollection(COL_NAME_USER)
	storage.TiedotDB.AddCollection(COL_NAME_CHAN)
	storage.TiedotDB.AddCollection(COL_NAME_MSG)
}

func UpdateAllUsers(users []*entity.User) {
	fmt.Println(users)
}

func AddUser(user *entity.User) error {
	_, err := storage.TiedotDB.AddDocument(COL_NAME_USER, structs.Map(user))
	if err != nil {
		fmt.Println("add user error!", err.Error())
		return err
	}
	return nil
}

func GetAllUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	userDocuments, err := storage.TiedotDB.GetAllDocuments(COL_NAME_USER)
	if err != nil {
		return users, err
	}

	for _, val := range userDocuments {
		var user entity.User
		if err = json.Unmarshal(val, &user); err == nil {
			users = append(users, &user)
		}
	}

	return users, nil
}

func ClearAllUsers() error {
	err := storage.TiedotDB.DeleteCollection(COL_NAME_USER)
	if err != nil {
		fmt.Println("failed to clear all users.", err.Error())
		return err
	}

	// recreate collection
	storage.TiedotDB.AddCollection(COL_NAME_USER)

	return err
}
