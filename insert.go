package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

// Insert single data
func Insert(table string, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

// InsertMany inserts single or multiple data
func InsertMany(table string, data []interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

// InsertOrUpdate inserts or updates single data
func InsertOrUpdate(table string, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data, r.InsertOpts{Conflict: "update"}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

// InsertOrUpdateMany inserts or updates single or multiple data
func InsertOrUpdateMany(table string, data []interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data, r.InsertOpts{Conflict: "update"}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}
