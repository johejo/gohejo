package sqlbuilder

type InsertBuilder interface {
	Builder
	Columns(columns ...string) InsertBuilder
	Values(v ...interface{}) InsertBuilder
}

type insertBuilder struct {
	*builder
}

func newInsertBuilder(into string) *insertBuilder {
	b := &insertBuilder{builder: newBuilder()}
	b.appendString("INSERT INTO ", into, " ")
	return b
}

func (b *insertBuilder) Columns(columns ...string) InsertBuilder {
	b.appendStringWithBrackets(columns...)
	return b
}

func (b *insertBuilder) Values(v ...interface{}) InsertBuilder {
	b.appendString(" VALUES ")
	b.appendPlaceWithBrackets(v...)
	return b
}
