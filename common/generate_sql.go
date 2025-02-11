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
	FIND_WITH_CREATED_AT
	SOFT_DELETE
	HARD_DELETE
	SELECT_WITHOUT_COUNT
	SELECT_EXIST
)

func (r SQLMethod) String() string {
	return [...]string{"INSERT", "UPDATE", "FIND", "DELETE"}[r]
}

func GenerateSQLQueries(method SQLMethod, table string, fields []string, where *string) string {
	fieldList := strings.Join(fields, ", ")
	mappingList := ":" + strings.Join(fields, ", :")

	switch method {
	case INSERT:
		return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fieldList, mappingList)
	case UPDATE:
		updateList := []string{}
		for _, field := range fields {
			updateList = append(updateList, fmt.Sprintf("%s = :%s", field, field))
		}
		updateString := strings.Join(updateList, ", ")
		return fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, updateString, *where)
	case FIND:
		return fmt.Sprintf("SELECT %s FROM %s WHERE %s", fieldList, table, *where)
	case FIND_WITH_CREATED_AT:
		selectList := fieldList + ", created_at"
		return fmt.Sprintf("SELECT %s FROM %s WHERE %s", selectList, table, *where)
	case SELECT_WITHOUT_COUNT:
		return fmt.Sprintf("SELECT %s FROM %s ", fieldList, table)
	case SOFT_DELETE:
		return fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE %s", table, *where)
	case HARD_DELETE:
		return fmt.Sprintf("DELETE FROM %s WHERE %s", table, *where)
	case SELECT_EXIST:
		return fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s)", table, *where)
	}
	return ""
}
