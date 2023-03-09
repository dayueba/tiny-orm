package schema

import (
	"testing"
	"tiny-orm/dialect"
)

type User struct {
	Name string `db:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.TableName != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	if schema.GetField("TableName").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse primary key")
	}
}
