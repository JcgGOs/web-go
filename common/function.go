package common

import "time"

//FmtDate format to date
func FmtDate(t time.Time) string {
	return t.Format("2006-1-2")
}

//FmtTime format to time
func FmtTime(t time.Time) string {
	return t.Format("2006-1-2 15:04:05")
}
