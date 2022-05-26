Gorm Adapter
====

[![Go Report Card](https://goreportcard.com/badge/github.com/casbin/gorm-adapter)](https://goreportcard.com/report/github.com/casbin/gorm-adapter)
[![Build Status](https://travis-ci.com/casbin/gorm-adapter.svg?branch=master)](https://travis-ci.com/casbin/gorm-adapter)
[![Coverage Status](https://coveralls.io/repos/github/casbin/gorm-adapter/badge.svg?branch=master)](https://coveralls.io/github/casbin/gorm-adapter?branch=master)
[![Godoc](https://godoc.org/github.com/casbin/gorm-adapter?status.svg)](https://godoc.org/github.com/casbin/gorm-adapter)
[![Release](https://img.shields.io/github/release/casbin/gorm-adapter.svg)](https://github.com/casbin/gorm-adapter/releases/latest)
[![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/casbin/lobby)
[![Sourcegraph](https://sourcegraph.com/github.com/casbin/gorm-adapter/-/badge.svg)](https://sourcegraph.com/github.com/casbin/gorm-adapter?badge)

Gorm Adapter is the [Gorm](https://gorm.io/gorm) adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from Gorm supported database or save policy to it.

Based on [Officially Supported Databases](http://jinzhu.me/gorm/database.html), The current supported databases are:

- MySQL
- PostgreSQL
- Sqlite3
- SQL Server

You may find other 3rd-party supported DBs in Gorm website or other places.

## Installation

    go get github.com/casbin/gorm-adapter/v3

## Simple Example

```go
package main

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	a, _ := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)
	
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()
	
	// Check the permission.
	e.Enforce("alice", "data1", "read")
	
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	
	// Save the policy back to DB.
	e.SavePolicy()
}
```

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under Apache 2.0 License. See the [LICENSE](LICENSE) file for the full license text.
