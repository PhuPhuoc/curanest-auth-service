package common

import (
	"fmt"
	"strings"
)

type SQLMethod int

const (
	INSERT SQLMethod = iota
	UPDATE
	FIND
	DELETE
)

func (r SQLMethod) String() string {
	return [...]string{"INSERT", "UPDATE", "FIND", "DELETE"}[r]
}

func GenerateSQLQueries(method SQLMethod, table string, fields []string, where *string) string {
	fieldList := strings.Join(fields, ", ")
	mappingList := ":" + strings.Join(fields, ", :")

	updateList := []string{}
	for _, field := range fields {
		updateList = append(updateList, fmt.Sprintf("%s = :%s", field, field))
	}
	updateString := strings.Join(updateList, ", ")
	selectList := fieldList + ", created_at"

	switch method {
	case INSERT:
		return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fieldList, mappingList)
	case UPDATE:
		return fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, updateString, *where)
	case FIND:
		return fmt.Sprintf("SELECT %s FROM %s WHERE %s", selectList, table, *where)
	case DELETE:
		return fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE %s", table, *where)
	}
	return ""
}
