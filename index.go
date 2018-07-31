package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

// Creates new index
func IndexCreate(table string, indexName string, indexField string, session *r.Session) {
	if IndexExists(table, indexName, session) {
		return
	}

	err := r.Table(table).IndexCreateFunc(indexName, r.Row.Field(indexField)).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}
}

// Check if index exists
func IndexExists(table string, index string, session *r.Session) bool {
	var exists bool
	res, err := r.Table(table).IndexList().Contains(index).Run(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	if res.One(&exists) != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return exists
}
