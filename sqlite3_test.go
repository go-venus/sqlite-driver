package sqlite

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestSqlite_DataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Tom", "text"},
		{123, "integer"},
		{1.2, "real"},
		{[]int{1, 2, 3}, "blob"},
		{time.Now(), "datetime"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.Value), func(t *testing.T) {
			if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type {
				t.Fatalf("expect %s, but got %s", c.Type, typ)
			}
		})
	}
}
