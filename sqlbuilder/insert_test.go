package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Parallel()
	var (
		s            string
		args         []interface{}
		expectedSQL  string
		expectedArgs []interface{}
	)
	expectedSQL = "INSERT INTO foo (bar) VALUES (?)"
	expectedArgs = []interface{}{1}
	s, args = Insert("foo").Columns("bar").Values(1).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}

	expectedSQL = "INSERT INTO foo (bar1, bar2) VALUES (?, ?)"
	expectedArgs = []interface{}{1, 2}
	s, args = Insert("foo").Columns("bar1", "bar2").Values(1, 2).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}
