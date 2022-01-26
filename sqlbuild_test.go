package sqlbuild

import (
	"fmt"
	"testing"
	"time"
)

func TestSqlBuild(t *testing.T) {

	sb := &SqlBuild{}

	sb.From("persons").Where("name = ? and age = ?", "linuxea", 12).
		Where("score > ?", 18).
		Group("age desc").
		Limit(10).Offset(20).
		Where("created_at > ?", time.Now())

	fmt.Println(sb.Sql())
}
