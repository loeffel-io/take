# take

[WIP] Simple RethinkDB ORM based on [gorethink](https://github.com/GoRethink/gorethink)

- Components
    - [Connect](https://github.com/loeffel-io/take#component-connect)
        - [Basic Connection](https://github.com/loeffel-io/take#basic-connection)
        - [Connection Pool](https://github.com/loeffel-io/take#connection-pool)
        - [Secure Connection](https://github.com/loeffel-io/take#secure-connection)
    - [Database](https://github.com/loeffel-io/take#component-database)
        - [Create Database](https://github.com/loeffel-io/take#create-database)
        - [Check if Database exists](https://github.com/loeffel-io/take#check-if-database-exists)
    - [Index](https://github.com/loeffel-io/take#component-index)
        - [Create Index](https://github.com/loeffel-io/take#create-index)
        - [Check if Index exists](https://github.com/loeffel-io/take#check-if-index-exists)
    - [Table](https://github.com/loeffel-io/take#component-table)
        - [Create Tables](https://github.com/loeffel-io/take#create-tables)
        - [Check if Table exists](https://github.com/loeffel-io/take#check-if-table-exists)
    - [User](https://github.com/loeffel-io/take#component-user)
        - [Setup user](https://github.com/loeffel-io/take#setup-user)                

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

## Component: Table

### Create Tables

Create tables if not exists

```go
package main

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    take.TablesCreate([]string{"table1", "table2"}, databaseSession)
}
```

### Check if Table exists

```go
package main

import (
	"log"
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
    exists := take.TableExists("table", databaseSession)
    
    if exists{
    	log.Println("Table exists")
    	return
    }
    
    log.Println("Table not exists")
}
```

## Component: User

### Setup user

Create or update user and set user permissions for the current database 

```go
package main

import (
	r "gopkg.in/gorethink/gorethink.v4"
	"github.com/loeffel-io/take"
)

var databaseSession *r.Session

func main(){
	user := take.User{
        ID:             "username",
        Password:       "password",
        Permissions:    take.UserPermissions{
            Read:   true,
            Write:  false,
            Config: false,
        },
    }
	
    take.SetupUser(user, databaseSession)
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