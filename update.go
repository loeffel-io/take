package take

import (
	"github.com/getsentry/raven-go"
	r "gopkg.in/gorethink/gorethink.v4"
	"log"
)

func Update(table string, id interface{}, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Get(id).Update(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}
