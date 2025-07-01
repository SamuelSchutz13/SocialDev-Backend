package utils

import "database/sql"

func SqlNullToString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}

	return s.String
}
