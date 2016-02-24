package loaders

import (
	"net/url"
	"path"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/knq/xo/internal"
	"github.com/knq/xo/models"
)

func init() {
	internal.SchemaLoaders["sqlite3"] = internal.TypeLoader{
		Schemes:        []string{"sqlite3", "sqlite", "file"},
		ProcessDSN:     SqProcessDSN,
		ProcessRelkind: SqRelkind,
		ParamN:         func(int) string { return "?" },
		MaskFunc:       func() string { return "?" },
		ParseType:      SqParseType,
		TableList: func(db models.XODB, schema string, relkind string) ([]*models.Table, error) {
			return models.SqTables(db, relkind)
		},
		ColumnList: func(db models.XODB, schema string, table string) ([]*models.Column, error) {
			return models.SqTableColumns(db, table)
		},
		ForeignKeyList: func(db models.XODB, schema string, table string) ([]*models.ForeignKey, error) {
			return models.SqTableForeignKeys(db, table)
		},
		IndexList: func(db models.XODB, schema string, table string) ([]*models.Index, error) {
			return models.SqTableIndexes(db, table)
		},
		IndexColumnList: func(db models.XODB, schema string, table string, index string) ([]*models.IndexColumn, error) {
			return models.SqIndexColumns(db, index)
		},
		QueryColumnList: SqQueryColumns,
	}
}

// SqProcessDSN processes a sqlite DSN.
func SqProcessDSN(u *url.URL, protocol string) string {
	p := u.Opaque
	if u.Path != "" {
		p = u.Path
	}

	if u.Host != "" && u.Host != "localhost" {
		p = path.Join(u.Host, p)
	}

	return p + u.Query().Encode()
}

// SqRelkind returns the sqlite string representation for RelType.
func SqRelkind(relType internal.RelType) string {
	var s string
	switch relType {
	case internal.Table:
		s = "table"
	case internal.View:
		s = "view"
	default:
		panic("unsupported RelType")
	}
	return s
}

// SqParseType parse a postgres type into a Go type based on the column
// definition.
func SqParseType(args *internal.ArgType, dt string, nullable bool) (int, string, string) {
	precision := 0
	nilType := "nil"

	// extract length
	if loc := internal.LenRE.FindStringIndex(dt); len(loc) != 0 {
		precision, _ = strconv.Atoi(dt[loc[0]+1 : loc[1]-1])
		dt = dt[:loc[0]]
	}

	var typ string
	switch dt {
	case "boolean":
		nilType = "false"
		typ = "bool"
		if nullable {
			nilType = "sql.NullBool{}"
			typ = "sql.NullBool"
		}

	case "text":
		nilType = `""`
		typ = "string"
		if nullable {
			nilType = "sql.NullString{}"
			typ = "sql.NullString"
		}

	case "integer":
		nilType = "0"
		typ = args.Int32Type
		if nullable {
			nilType = "sql.NullInt64{}"
			typ = "sql.NullInt64"
		}

	case "numeric", "real":
		nilType = "0.0"
		typ = "float64"
		if nullable {
			nilType = "sql.NullFloat64{}"
			typ = "sql.NullFloat64"
		}

	case "blob":
		typ = "[]byte"
	}

	return precision, nilType, typ
}

// SqQueryColumns parses a sqlite query and generates a type for it.
func SqQueryColumns(args *internal.ArgType, inspect []string) ([]*models.Column, error) {
	var err error

	// create temporary view xoid
	xoid := "_xo_" + internal.GenRandomID()
	viewq := `CREATE TEMPORARY VIEW ` + xoid + ` AS ` + strings.Join(inspect, "\n")
	models.XOLog(viewq)
	_, err = args.DB.Exec(viewq)
	if err != nil {
		return nil, err
	}

	// load column information
	return models.SqTableColumns(args.DB, xoid)
}
