package helper

import (
	"database/sql"
	"strconv"
	"strings"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid: true,
	}
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}

