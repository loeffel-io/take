package take

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/getsentry/raven-go"
)

func Update(table string, id string, data interface{}, session *r.Session) interface{} {
	err := r.Table(table).Get(id).Update(data).Exec(session)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return data
}
