# go RESTful API

go package
- [go-chi/chi: lightweight, idiomatic and composable router for building Go HTTP services](https://github.com/go-chi/chi)
- [jmoiron/sqlx: general purpose extensions to golang's database/sql](https://github.com/jmoiron/sqlx)

db
- MySQL

## Setup

ENV

```
$ export DB_HOST=127.0.0.1
$ export DB_USER=root
$ export DB_PASS=***
$ export DB_NAME=go_sample
```

or Create an .env file


```sql
mysql> CREATE DATABASE go_sample;
mysql> use go_sample

mysql> CREATE TABLE IF NOT EXISTS `users` (
`id` BIGINT unsigned NOT NULL AUTO_INCREMENT,
`first_name` varchar(50) NOT NULL,
`last_name` varchar(50) NOT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

mysql> DESC users;

mysql> INSERT INTO users (first_name, last_name) VALUES ("tanaka", "taro");
mysql> INSERT INTO users (first_name, last_name) VALUES ("tanaka", "jiro");
mysql> SELECT * FROM users;
```

### Run

```
$ go run main.go
```

```
$ curl -s -XGET localhost:8080/api/v1/users |jq .
```
