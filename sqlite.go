package sqlite

import (
	"reflect"

	"github.com/go-venus/venus/schema"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	schema.RegisterDialect("sqlite3", &sqlite{})
}

type sqlite struct {
}

func (s sqlite) DataTypeOf(field *schema.Field) string {

	t := field.FieldType
	switch t.Kind() {
	case reflect.Bool:
		return "boolean"
	case reflect.Int, reflect.Uint:
		sqlType := "bigint"

		if field.FieldType.Kind() == reflect.Uint {
			sqlType += " unsigned"
		}

		return sqlType
	case reflect.Float64:
		return "double"
	case reflect.Float32:
		return "float"
	case reflect.String:
		return "longtext"
	case reflect.Struct:
		// field.ValueOf()
		// if _, ok := t.Interface().(time.Time); ok {
		// 	return "datetime"
		// }
	case reflect.Array:
		// if _, ok := t.Interface().([]byte); ok {
		// 	return "longblob"
		// }

	}
	// panic(fmt.Sprintf("invalid sql type %s (%s)", t.Type().Name(), t.Kind()))
	return ""
}

func (s sqlite) TableExistSQL(tableName string) (string, []any) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
