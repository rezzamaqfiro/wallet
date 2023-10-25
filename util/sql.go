package util

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func SqlString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func SqlTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func SqlInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func SqlUUID(s string) uuid.NullUUID {
	i, err := uuid.Parse(s)
	return uuid.NullUUID{UUID: i, Valid: err == nil}
}

func SqlBool(s bool) sql.NullBool {
	return sql.NullBool{Bool: s, Valid: true}
}
