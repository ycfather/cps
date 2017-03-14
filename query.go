package main

type Query interface {
	BuildSql() (sql string)
}
