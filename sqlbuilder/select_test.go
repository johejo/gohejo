package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT foo FROM bar"
	s, _ = Select("foo").From("bar").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}

	expectedSQL = "SELECT foo1, foo2 FROM bar"
	s, _ = Select("foo1", "foo2").From("bar").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}

	expectedSQL = "SELECT foo1, foo2, foo3 FROM bar"
	s, _ = Select("foo1", "foo2", "foo3").From("bar").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_Where(t *testing.T) {
	t.Parallel()
	var (
		s            string
		args         []interface{}
		expectedSQL  string
		expectedArgs []interface{}
	)

	expectedSQL = "SELECT foo1, foo2 FROM bar WHERE baz = ?"
	expectedArgs = []interface{}{1}
	s, args = Select("foo1", "foo2").From("bar").Where("baz = ?", 1).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}

func TestSelectBuilder_OrderBy(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT foo1, foo2 FROM bar ORDER BY id DESC"
	s, _ = Select("foo1", "foo2").From("bar").OrderBy("id", "DESC").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_GroupBy(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT foo1, foo2 FROM bar GROUP BY id DESC"
	s, _ = Select("foo1", "foo2").From("bar").GroupBy("id", "DESC").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_Count(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT COUNT(id) FROM foo"
	s, _ = Select().Count("id").From("foo").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}

	expectedSQL = "SELECT id, COUNT(id) FROM foo"
	s, _ = Select("id").Count("id").From("foo").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_Distinct(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT DISTINCT id FROM foo"
	s, _ = Select().Distinct("id").From("foo").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}

	expectedSQL = "SELECT DISTINCT id, name FROM foo"
	s, _ = Select().Distinct("id", "name").From("foo").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_Having(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT id, name, COUNT(id) FROM foo GROUP BY bar ASC HAVING name = ?"
	s, _ = Select("id", "name").Count("id").From("foo").GroupBy("bar", "ASC").Having("name = ?", 1).ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, s=%s", expectedSQL, s)
	}
}

func TestSelectBuilder_Limit(t *testing.T) {
	t.Parallel()
	var (
		s            string
		expectedSQL  string
		args         []interface{}
		expectedArgs []interface{}
	)

	expectedSQL = "SELECT id FROM foo LIMIT ?"
	expectedArgs = []interface{}{10}
	s, args = Select("id").From("foo").Limit(10).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}

func TestSelectBuilder_Offset(t *testing.T) {
	t.Parallel()
	var (
		s            string
		expectedSQL  string
		args         []interface{}
		expectedArgs []interface{}
	)

	expectedSQL = "SELECT id FROM foo OFFSET ?"
	expectedArgs = []interface{}{10}
	s, args = Select("id").From("foo").Offset(10).ToSQL()
	if s != expectedSQL || !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("failed to build s: expected=%s, %s, s=%s, args=%s", expectedSQL, expectedArgs, s, args)
	}
}

func TestSelectBuilder_Join(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT id FROM foo JOIN bar ON foo.id = bar.id"
	s, _ = Select("id").From("foo").Join("bar").On("foo.id", "bar.id").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}

	expectedSQL = "SELECT id FROM foo JOIN bar USING(id)"
	s, _ = Select("id").From("foo").Join("bar").Using("id").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}
}

func TestSelectBuilder_LeftJoin(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT id FROM foo LEFT JOIN bar ON foo.id = bar.id"
	s, _ = Select("id").From("foo").LeftJoin("bar").On("foo.id", "bar.id").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}

	expectedSQL = "SELECT id FROM foo LEFT JOIN bar USING(id)"
	s, _ = Select("id").From("foo").LeftJoin("bar").Using("id").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}
}

func TestSelectBuilder_Using(t *testing.T) {
	t.Parallel()
	var (
		s           string
		expectedSQL string
	)

	expectedSQL = "SELECT id FROM foo JOIN bar USING(id)"
	s, _ = Select("id").From("foo").Join("bar").Using("id").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}

	expectedSQL = "SELECT id FROM foo JOIN bar USING(id, name)"
	s, _ = Select("id").From("foo").Join("bar").Using("id", "name").ToSQL()
	if s != expectedSQL {
		t.Errorf("failed to build s: expected=%s, %s", expectedSQL, s)
	}
}

func BenchmarkSelect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Select("foo1", "foo2").From("bar").Where("baz = ?", 1).OrderBy("id", "DESC").ToSQL()
	}
}
