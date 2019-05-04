package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	var (
		s            string
		args         []interface{}
		expectedSQL  string
		expectedArgs []interface{}
	)

	expectedSQL = "DELETE FROM foo WHERE bar1 = ? OR bar2 = ?"
	expectedArgs = []interface{}{1, 2}
	s, args = Delete("foo").Where("bar1 = ? OR bar2 = ?", 1, 2).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}
