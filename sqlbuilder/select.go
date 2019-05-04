package sqlbuilder

type SelectBuilder interface {
	Builder
	Count(column string) SelectBuilder
	From(tables ...string) SelectBuilder
	Distinct(column ...string) SelectBuilder
	OrderBy(column string, order string) SelectBuilder
	GroupBy(column string, order string) SelectBuilder
	Having(cond string, v ...interface{}) SelectBuilder
	Join(table string) SelectBuilder
	LeftJoin(table string) SelectBuilder
	On(left string, right string) SelectBuilder
	Using(columns ...string) SelectBuilder
	Limit(limit int) SelectBuilder
	Offset(offset int) SelectBuilder
	Where(cond string, v ...interface{}) SelectBuilder
}

type selectBuilder struct {
	*builder
	_select bool
}

func newSelectBuilder(columns ...string) *selectBuilder {
	b := &selectBuilder{builder: newBuilder()}
	b.appendString("SELECT")
	b.appendStringWithComma(columns...)
	if len(columns) != 0 {
		b._select = true
	}
	return b
}

func (b *selectBuilder) Count(column string) SelectBuilder {
	if b._select {
		b.appendString(",")
	}
	b.appendString(" COUNT(", column, ")")
	return b
}

func (b *selectBuilder) Distinct(column ...string) SelectBuilder {
	b.appendString(" DISTINCT")
	b.appendStringWithComma(column...)
	return b
}

func (b *selectBuilder) From(tables ...string) SelectBuilder {
	b.appendString(" FROM")
	b.appendStringWithComma(tables...)
	return b
}

func (b *selectBuilder) OrderBy(column string, order string) SelectBuilder {
	b.appendOrder("ORDER BY", column, order)
	return b
}

func (b *selectBuilder) GroupBy(column string, order string) SelectBuilder {
	b.appendOrder("GROUP BY", column, order)
	return b
}

func (b *selectBuilder) Having(cond string, v ...interface{}) SelectBuilder {
	b.appendString(" HAVING ", cond)
	b.appendArg(v...)
	return b
}

func (b *selectBuilder) Join(table string) SelectBuilder {
	b.appendString(" JOIN ", table)
	return b
}

func (b *selectBuilder) LeftJoin(table string) SelectBuilder {
	b.appendString(" LEFT JOIN ", table)
	return b
}

func (b *selectBuilder) On(left string, right string) SelectBuilder {
	b.appendString(" ON ", left, " = ", right)
	return b
}

func (b *selectBuilder) Using(columns ...string) SelectBuilder {
	b.appendString(" USING")
	b.appendStringWithBrackets(columns...)
	return b
}

func (b *selectBuilder) Limit(limit int) SelectBuilder {
	b.appendString(" LIMIT ")
	b.appendPlace(limit)
	return b
}

func (b *selectBuilder) Offset(offset int) SelectBuilder {
	b.appendString(" OFFSET ")
	b.appendPlace(offset)
	return b
}

func (b *selectBuilder) Where(cond string, v ...interface{}) SelectBuilder {
	b.appendCond("WHERE", cond, v...)
	return b
}
