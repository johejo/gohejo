package sqlbuilder

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	var (
		s            string
		args         []interface{}
		expectedSQL  string
		expectedArgs []interface{}
	)

	expectedSQL = "UPDATE foo SET bar1 = ?"
	expectedArgs = []interface{}{1}
	s, args = Update("foo").Set("bar1", 1).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}

	expectedSQL = "UPDATE foo SET bar1 = ? WHERE baz = ?"
	expectedArgs = []interface{}{1, 2}
	s, args = Update("foo").Set("bar1", 1).Where("baz = ?", 2).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}

	expectedSQL = "UPDATE foo SET bar1 = ?, bar2 = ? WHERE baz = ?"
	expectedArgs = []interface{}{1, 2, 3}
	s, args = Update("foo").Set("bar1", 1).Set("bar2", 2).Where("baz = ?", 3).ToSQL()
	fmt.Println(s)
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}
