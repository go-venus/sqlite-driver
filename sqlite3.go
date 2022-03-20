package sqlite

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-venus/venus/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dialect.RegisterDialect("sqlite3", &sqlite3{})
}

type sqlite3 struct{}

func (s sqlite3) DataTypeOf(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := t.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", t.Type().Name(), t.Kind()))

}

func (s sqlite3) TableExistSQL(tableName string) (string, []any) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type= 'table' and name = ?", args
}
