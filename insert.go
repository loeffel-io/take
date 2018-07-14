package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

func Insert(table string, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

func InsertMany(table string, data []interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

func InsertOrUpdate(table string, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data, r.InsertOpts{Conflict: "update"}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}

func InsertOrUpdateMany(table string, data []interface{}, session *r.Session) interface{} {
	err := r.Table(table).Insert(data, r.InsertOpts{Conflict: "update"}).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}
