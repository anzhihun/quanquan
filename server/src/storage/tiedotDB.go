package storage

import (
	"errors"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
)

type tiedotDB struct {
	dbDir string
	db    *db.DB
}

var TiedotDB = tiedotDB{dbDir: "db", db: nil}

func (t *tiedotDB) init() error {
	if t.db == nil {
		var err error
		t.db, err = db.OpenDB(t.dbDir)
		if err != nil {
			fmt.Println("init db error: ", err.Error())
			return err
		}
	}

	return nil
}

func (t *tiedotDB) AddCollection(collectionName string) error {
	if t.db == nil {
		if err := t.init(); err != nil {
			return err
		}
	}

	if oldCol := t.db.Use(collectionName); oldCol != nil {
		// already exist
		return nil
	}

	if err := t.db.Create(collectionName); err != nil {
		return err
	}

	return nil
}

func (t *tiedotDB) DeleteCollection(collectionName string) error {
	if t.db == nil {
		if err := t.init(); err != nil {
			return err
		}
	}

	if oldCol := t.db.Use(collectionName); oldCol == nil {
		// not exist
		return nil
	}

	if err := t.db.Drop(collectionName); err != nil {
		return err
	}

	return nil
}

func (t *tiedotDB) AddDocument(collectionName string, docContent map[string]interface{}) (docId int, err error) {
	if t.db == nil {
		if err := t.init(); err != nil {
			return 0, err
		}
	}

	collection := t.db.Use(collectionName)
	if collection == nil {
		return 0, errors.New("collection " + collectionName + " is not exist.")
	}

	docID, err := collection.Insert(docContent)
	if err != nil {
		return 0, err
	}

	return docID, nil
}

func (t *tiedotDB) DeleteDocument(collectionName string, docId int) error {
	if t.db == nil {
		if err := t.init(); err != nil {
			return err
		}
	}

	collection := t.db.Use(collectionName)
	if collection == nil {
		return errors.New("collection " + collectionName + " is not exist.")
	}

	if err := collection.Delete(docId); err != nil {
		return err
	}

	return nil
}

func (t *tiedotDB) UpdateDocument(collectionName string, docId int, docContent map[string]interface{}) error {
	if t.db == nil {
		if err := t.init(); err != nil {
			return err
		}
	}

	collection := t.db.Use(collectionName)
	if collection == nil {
		return errors.New("collection " + collectionName + " is not exist.")
	}

	if err := collection.Update(docId, docContent); err != nil {
		return err
	}

	return nil
}

func (t *tiedotDB) GetAllDocuments(collectionName string) (map[int][]byte, error) {
	if t.db == nil {
		if err := t.init(); err != nil {
			return nil, err
		}
	}

	collection := t.db.Use(collectionName)
	if collection == nil {
		return nil, errors.New("collection " + collectionName + " is not exist.")
	}

	documents := map[int][]byte{}
	collection.ForEachDoc(func(docId int, docContent []byte) (willMoveOn bool) {
		fmt.Println("Document ", docId, " is ", string(docContent))
		documents[docId] = docContent
		// documents = append(documents, docContent)
		return true // move on to the next document OR
	})

	return documents, nil
}
