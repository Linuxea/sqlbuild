package sqlbuild

import (
	"fmt"
	"strings"
)

type SelectKeyWorld string

const (
	SELECT SelectKeyWorld = "SELECT"
	JOIN   SelectKeyWorld = "JOIN"
	ON     SelectKeyWorld = "ON"
	FROM   SelectKeyWorld = "FROM"
	WHERE  SelectKeyWorld = "WHERE"
	AND    SelectKeyWorld = "AND"
	GROUP  SelectKeyWorld = "GROUP BY"
	LIMIT  SelectKeyWorld = "LIMIT"
	OFFSET SelectKeyWorld = "OFFSET"
)

type sqlArg struct {
	condition string
	args      []interface{}
}

type JoinType int

const (
	INNERJOIN JoinType = iota
	LEFTJOIN
	RIGHTJOIN
	FULLJOIN
)

var joinTypeString = []string{"INNER JOIN", "LEFT JOIN", "RIGHT JOIN"}

type join struct {
	Type  JoinType
	Table string
	On    string
}

type SqlBuild struct {
	selects   []string
	from      string
	join      []*join
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

func (s *SqlBuild) Join(jt JoinType, table, on string) *SqlBuild {
	s.join = append(s.join, &join{Type: jt, Table: table, On: on})
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

	// # select
	sb.WriteString(string(SELECT))
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

	// # from
	sb.WriteString(string(FROM))
	sb.WriteByte(' ')
	sb.WriteString(s.from)
	sb.WriteByte(' ')

	// # join
	if len(s.join) > 0 {
		for index := range s.join {
			sb.WriteString(joinTypeString[s.join[index].Type])
			sb.WriteByte(' ')
			sb.WriteString(s.join[index].Table)
			sb.WriteByte(' ')
			sb.WriteString(string(ON))
			sb.WriteByte(' ')
			sb.WriteString(s.join[index].On)
			sb.WriteByte(' ')
		}
	}

	// # wheres
	if len(s.wheres) > 0 {
		sb.WriteString(string(WHERE))
		sb.WriteByte(' ')

		var count uint
		var allWhere strings.Builder
		for _, temp := range s.wheres {

			if count != 0 {
				allWhere.WriteString(string(AND))
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

	// # group by
	if len(s.groups) > 0 {
		sb.WriteString(string(GROUP))
		sb.WriteByte(' ')
		for _, temp := range s.groups {
			sb.WriteString(temp)
			sb.WriteByte(' ')
		}
	}

	// # limit
	if s.limit > 0 {
		sb.WriteString(string(LIMIT))
		sb.WriteByte(' ')
		sb.WriteString(fmt.Sprintf("%d", s.limit))
		sb.WriteByte(' ')
	}

	// # offset
	if s.offset > 0 {
		sb.WriteString(string(OFFSET))
		sb.WriteByte(' ')
		sb.WriteString(fmt.Sprintf("%d", s.offset))
		sb.WriteByte(' ')
	}

	return sb.String(), args
}
