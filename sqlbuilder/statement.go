package sqlbuilder

func Select(columns ...string) *selectBuilder {
	return newSelectBuilder(columns...)
}

func Insert(into string) InsertBuilder {
	return newInsertBuilder(into)
}

func Update(table string) UpdateBuilder {
	return newUpdateBuilder(table)
}

func Delete(from string) DeleteBuilder {
	return newDeleteBuilder(from)
}
