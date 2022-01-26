package sqlbuild

import (
	"fmt"
	"testing"
	"time"
)

func TestSqlBuild(t *testing.T) {

	sb := &SqlBuild{}

	minScore := 0

	sb.From("person").
		Join(LEFTJOIN, "score", "score.student_id = person.id").
		Where("name = ? and age = ?", "linuxea", 12).
		Group("age desc").
		Limit(10).Offset(20).
		Where("created_at > ?", time.Now().Format("2006"))

	if minScore > 0 {
		sb.Where("score > ?", 10)
	}

	sql, args := sb.Sql()
	fmt.Println(sql)
	fmt.Println(args)

}
