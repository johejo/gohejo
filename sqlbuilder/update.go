package sqlbuilder

type UpdateBuilder interface {
	Builder
	Set(column string, v interface{}) UpdateBuilder
	Where(cond string, v interface{}) UpdateBuilder
}

type updateBuilder struct {
	*builder
	set bool
}

func newUpdateBuilder(table string) *updateBuilder {
	b := &updateBuilder{builder: newBuilder()}
	b.appendString("UPDATE ", table)
	return b
}

func (b *updateBuilder) Set(column string, v interface{}) UpdateBuilder {
	if !b.set {
		b.appendString(" SET")
		b.set = true
	} else {
		b.appendString(",")
	}
	b.appendString(" ", column, " = ", "?")
	b.appendArg(v)
	return b
}

func (b *updateBuilder) Where(column string, v interface{}) UpdateBuilder {
	b.appendCond("WHERE", column, v)
	return b
}
