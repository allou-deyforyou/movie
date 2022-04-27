package firebase

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
)

func GetAll[T interface{}](collection firestore.Query) []T {
	documentIterator := collection.Documents(context.Background())
	documentList, err := documentIterator.GetAll()
	if err != nil {
		panic(err)
	}
	results := make([]T, 0)
	for _, document := range documentList {
		data := new(T)
		b, _ := json.Marshal(document.Data())
		json.Unmarshal(b, data)
		results = append(results, *data)
	}
	return results
}

func Snapshots[T interface{}](collection firestore.Query, callback func([]T)) {
	snapshotIterator := collection.Snapshots(context.Background())
	for {
		querySnapshot, err := snapshotIterator.Next()
		if err != nil {
			break
		}
		documentList, err := querySnapshot.Documents.GetAll()
		if err != nil {
			panic(err)
		}
		results := make([]T, 0)
		for _, document := range documentList {
			data := new(T)
			b, _ := json.Marshal(document.Data())
			json.Unmarshal(b, data)
			results = append(results, *data)
		}
		callback(results)
	}
}
