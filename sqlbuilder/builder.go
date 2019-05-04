package sqlbuilder

import (
	"strings"
)

type Builder interface {
	ToSQL() (string, []interface{})
}

type builder struct {
	sb strings.Builder
	a  []interface{}
}

func newBuilder() *builder {
	var (
		sb strings.Builder
		a  []interface{}
	)
	return &builder{sb: sb, a: a}
}

func (b *builder) ToSQL() (string, []interface{}) {
	return b.sb.String(), b.a
}

func (b *builder) appendString(v ...string) {
	for _, s := range v {
		b.sb.WriteString(s)
	}
}

func (b *builder) appendArg(v ...interface{}) {
	for _, s := range v {
		b.a = append(b.a, s)
	}
}

func (b *builder) appendPlace(v interface{}) {
	b.appendString("?")
	b.appendArg(v)
}

func (b *builder) appendOrder(by string, column string, order string) {
	if order == "" {
		order = "DESC"
	}
	b.appendString(" ", by, " ", column, " ", order)
}

func (b *builder) appendStringWithComma(ss ...string) {
	for i, s := range ss {
		b.appendString(" ", s)
		if i != len(ss)-1 {
			b.appendString(",")
		}
	}
}

func (b *builder) appendStringWithBrackets(ss ...string) {
	b.appendString("(")
	for i, s := range ss {
		b.appendString(s)
		if i != len(ss)-1 {
			b.appendString(", ")
		}
	}
	b.appendString(")")
}

func (b *builder) appendPlaceWithBrackets(v ...interface{}) {
	b.appendString("(")
	for i, s := range v {
		b.appendPlace(s)
		if i != len(v)-1 {
			b.appendString(", ")
		}
	}
	b.appendString(")")
}

func (b *builder) appendCond(s string, cond string, v ...interface{}) {
	b.appendString(" ", s, " ", cond)
	b.appendArg(v...)
}
