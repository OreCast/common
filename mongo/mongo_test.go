package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// TestMongoInsert
func TestMongoInsert(t *testing.T) {
	// our db attributes
	dbname := "chess"
	collname := "test"
	InitMongoDB("mongodb://localhost:8230")

	// remove all records in test collection
	Remove(dbname, collname, bson.M{})

	// insert one record
	var records []Record
	dataset := "/a/b/c"
	rec := Record{"dataset": dataset}
	records = append(records, rec)
	Insert(dbname, collname, records)

	// look-up one record
	spec := bson.M{"dataset": dataset}
	idx := 0
	limit := 1
	records = Get(dbname, collname, spec, idx, limit)
	if len(records) != 1 {
		t.Errorf("unable to find records using spec '%s', records %+v", spec, records)
	}

	// modify our record
	rec = Record{"dataset": dataset, "test": 1}
	records = []Record{}
	records = append(records, rec)
	err := Upsert(dbname, collname, "dataset", records)
	if err != nil {
		t.Error(err)
	}
	spec = bson.M{"test": 1}
	records = Get(dbname, collname, spec, idx, limit)
	if len(records) != 1 {
		t.Errorf("unable to find records using spec '%s', records %+v", spec, records)
	}
}
