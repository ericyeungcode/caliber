package db_utils

import "gorm.io/gorm"

const (
	DefaultSqlLimit = 500 // default limit for queries
)

type SqlProperties struct {
	Fields string // comma-separated list of fields to select, e.g. "id, name, age"
	Order  string
	Offset int
	Limit  int // default = `DefaultSqlLimit`
}

func BuildStmt(db *gorm.DB, sqlProps *SqlProperties, args ...any) *gorm.DB {
	con := db
	if len(args) > 0 {
		con = con.Where(args[0], args[1:]...)
	}

	sqlLimit := DefaultSqlLimit

	if sqlProps != nil {
		if len(sqlProps.Fields) > 0 {
			con = con.Select(sqlProps.Fields)
		}

		if sqlProps.Order != "" {
			con = con.Order(sqlProps.Order)
		}
		if sqlProps.Offset > 0 {
			con = con.Offset(sqlProps.Offset)
		}

		if sqlProps.Limit > 0 {
			sqlLimit = sqlProps.Limit
		}
	}

	// We always add `limit` clause to prevent accidental large queries
	con = con.Limit(sqlLimit)
	return con
}

/*
	usage example:

fmt.Println("qry 1 count:", len(Must1(AutoQuery[Position](db, nil))))
fmt.Println("qry 2 count:", len(Must1(AutoQuery[Position](db, &SqlLimit{Limit: 1}, "user_id = 476253"))))
fmt.Println("qry 3 count:", len(Must1(AutoQuery[Position](db, nil, "instrument_id = ?", "BTC-USDT-PERPETUAL"))))
*/
func AutoQuery[T any](db *gorm.DB, sqlProps *SqlProperties, args ...any) (result []T, err error) {

	con := BuildStmt(db, sqlProps, args...)
	err = con.Find(&result).Error
	return result, err
}

/*
	usage example:

fmt.Println("sql 1:", GetSql[Position](db, &SqlProperties{Fields: "id"}))
fmt.Println("sql 2:", GetSql[Position](db, &SqlProperties{Limit: 1}, "user_id = 476253"))
fmt.Println("sql 3:", GetSql[Position](db, nil, "instrument_id = ?", "BTC-USDT-PERPETUAL"))

output:
sql 1: SELECT `id` FROM `t_position` LIMIT 500
sql 2: SELECT * FROM `t_position` WHERE user_id = 476253 LIMIT 1
sql 3: SELECT * FROM `t_position` WHERE instrument_id = 'BTC-USDT-PERPETUAL' LIMIT 500
*/
func GetSql[T any](db *gorm.DB, sqlProps *SqlProperties, args ...any) string {
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		con := BuildStmt(tx, sqlProps, args...)
		// T define the `table`
		var v T
		return con.Model(&v).Find(&[]any{})
	})
	return sql
}
