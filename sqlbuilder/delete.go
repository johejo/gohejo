package sqlbuilder

type DeleteBuilder interface {
	Builder
	Where(cond string, v ...interface{}) DeleteBuilder
}

type deleteBuilder struct {
	*builder
}

func newDeleteBuilder(from string) *deleteBuilder {
	b := &deleteBuilder{builder: newBuilder()}
	b.appendString("DELETE FROM ")
	b.appendString(from)
	return b
}

func (b *deleteBuilder) Where(cond string, v ...interface{}) DeleteBuilder {
	b.appendCond("WHERE", cond, v...)
	return b
}
