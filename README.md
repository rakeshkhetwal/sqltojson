# sqltojson
Golang package to extract the sql data to json format.

## Features

- Perform the DB queries without the struct creation
- Result is in JSON format
- It supports Mysql and Postgresql Databases


## Tech

Uses a number of open source projects to work properly:

- Golang 
- postgresql/mysql

## Prerequisite

This requires [Golang] v1.0+ to run.

## Usage
Install the sqltojson package
```sh
go get github.com/rakeshkhetwal/sqltojson
```

Call the package, while passing the database name and db query
```sh
import "github.com/rakeshkhetwal/sqltojson"

DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

db, err := sql.Open(DB_DRIVER, DBURL)

query:="select * from table"

//calling sqltojson package
err,queryData:=sqltojson.SqlToJson(db,query)
```

## License

MIT

**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [golang]: <https://go.dev/dl/>
