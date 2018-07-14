# take

Simple RethinkDB ORM based on [gorethink](https://github.com/GoRethink/gorethink)

## Component: Connect

### Basic Connection

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
    })
}
```

### Connection Pool

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
        InitialCap: 10,
        MaxOpen:    10,
    })
}
```

### Secure Connection

e.g.: [compose.io](https://compose.io)

```go
package main

import (
	"os"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func init(){
    // Init database
    databaseSession = take.Connect(take.Connection{
        Address:    os.Getenv("RETHINKDB_ADDRESS") + ":" + os.Getenv("RETHINKDB_PORT"),
        Database:   os.Getenv("RETHINKDB_DATABASE"),
        Username:   os.Getenv("RETHINKDB_USERNAME"),
        Password:   os.Getenv("RETHINKDB_PASSWORD"),
        AuthKey:    os.Getenv("RETHINKDB_AUTHKEY"),
        CertPath:   os.Getenv("RETHINKDB_CERT_PATH"),
    })
}
```

## Component: Database

### Create Database

Create database if not exists

```go
package main

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    take.DatabaseCreate("test", databaseSession)
}
```

### Check if Database exists

```go
package main

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    exists := take.DatabaseExists("test", databaseSession)
    
    if exists{
    	log.Println("Database exists")
    	return
    }
    
    log.Println("Database not exists")
}
```

## Component: Index

### Create Index

Create index if not exists

```go
package main

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    take.IndexCreate("table", "index", "field", databaseSession)
}
```

### Check if Index exists

```go
package main

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    exists := take.IndexExists("table", "index", databaseSession)
    
    if exists{
    	log.Println("Index exists")
    	return
    }
    
    log.Println("Index not exists")
}
```

## Sentry support

This package supports [sentry.io](https://sentry.io) real time error reporting.
More informations: [Sentry golang docs](https://docs.sentry.io/clients/go)

```go
package main

import "github.com/getsentry/raven-go"

func init() {
    // Setup sentry
    raven.SetDSN("https://<key>:<secret>@sentry.io/<project>")
    
    // Setup cronjobs ...
}
```