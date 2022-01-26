package sqlbuild

import (
	"fmt"
	"strings"
)

type sqlArg struct {
	condition string
	args      []interface{}
}

type SqlBuild struct {
	selects   []string
	from      string
	wheres    []*sqlArg
	groups    []string
	orders    []string
	limit     uint
	offset    uint
	ForUpdate bool
}

func (s *SqlBuild) Select(sql string) *SqlBuild {
	s.selects = append(s.selects, sql)
	return s
}

func (s *SqlBuild) From(from string) *SqlBuild {
	s.from = from
	return s
}

func (s *SqlBuild) Where(sql string, args ...interface{}) *SqlBuild {
	s.wheres = append(s.wheres, &sqlArg{condition: sql, args: args})
	return s
}

func (s *SqlBuild) Group(sql string) *SqlBuild {
	s.groups = append(s.groups, sql)
	return s
}

func (s *SqlBuild) Order(sql string) *SqlBuild {
	s.orders = append(s.orders, sql)
	return s
}

func (s *SqlBuild) Limit(limit uint) *SqlBuild {
	s.limit = limit
	return s
}

func (s *SqlBuild) Offset(offset uint) *SqlBuild {
	s.offset = offset
	return s
}

func (s *SqlBuild) Sql() (string, interface{}) {

	sb := strings.Builder{}
	var args []interface{}

	// 1 select

	sb.WriteString("select")
	sb.WriteByte(' ')
	if len(s.selects) > 0 {
		for _, selectField := range s.selects {
			sb.WriteString(selectField)
			sb.WriteByte(' ')
		}
	} else {
		sb.WriteString("*")
		sb.WriteByte(' ')
	}

	// 2 from
	sb.WriteString("from")
	sb.WriteByte(' ')
	sb.WriteString(s.from)
	sb.WriteByte(' ')

	// 3 wheres
	if len(s.wheres) > 0 {
		sb.WriteString("where")
		sb.WriteByte(' ')

		var count uint
		var allWhere strings.Builder
		for _, temp := range s.wheres {

			if count != 0 {
				allWhere.WriteString("and")
				allWhere.WriteByte(' ')
			}

			allWhere.WriteString(temp.condition)
			allWhere.WriteByte(' ')
			args = append(args, temp.args...)
			count++
		}
		sb.WriteString(allWhere.String())
		sb.WriteByte(' ')
	}

	// 4 group by
	if len(s.groups) > 0 {
		sb.WriteString("group by")
		sb.WriteByte(' ')
		for _, temp := range s.groups {
			sb.WriteString(temp)
			sb.WriteByte(' ')
		}
	}

	// 5 limit
	if s.limit > 0 {
		sb.WriteString("limit")
		sb.WriteByte(' ')
		sb.WriteString(fmt.Sprintf("%d", s.limit))
		sb.WriteByte(' ')
	}

	// 6 offset
	if s.offset > 0 {
		sb.WriteString("offset")
		sb.WriteByte(' ')
		sb.WriteString(fmt.Sprintf("%d", s.offset))
		sb.WriteByte(' ')
	}

	return sb.String(), args
}
