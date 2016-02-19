{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Type }} represents a row from {{ schema .TableSchema .TableName }}.
{{- end }}
type {{ .Type }} struct {
{{- range .Fields }}
	{{ .Field }} {{ retype .Type }}{{ if .Tag }} `{{ .Tag }}`{{ end }} // {{ .ColumnName }}
{{- end }}
{{- if .PrimaryKeyField }}

	// xo fields
	_exists, _deleted bool
{{ end }}
}

{{ if .PrimaryKeyField }}
// Exists determines if the {{ .Type }} exists in the database.
func ({{ shortname .Type }} *{{ .Type }}) Exists() bool {
	return {{ shortname .Type }}._exists
}

// Deleted provides information if the {{ .Type }} has been deleted from the database.
func ({{ shortname .Type }} *{{ .Type }}) Deleted() bool {
	return {{ shortname .Type }}._deleted
}

// Insert inserts the {{ .Type }} to the database.
func ({{ shortname .Type }} *{{ .Type }}) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if {{ shortname .Type }}._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO {{ schema .TableSchema .TableName }} (` +
		`{{ colnames .Fields .PrimaryKeyField }}` +
		`) VALUES (` +
		`{{ colvals .Fields .PrimaryKeyField }}` +
		`) RETURNING {{ .PrimaryKey }}`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields (shortname .Type) .PrimaryKeyField }})
	err = db.QueryRow(sqlstr, {{ fieldnames .Fields (shortname .Type) .PrimaryKeyField }}).Scan(&{{ shortname .Type }}.{{ .PrimaryKeyField }})
	if err != nil {
		return err
	}

	// set existence
	{{ shortname .Type }}._exists = true

	return nil
}

// Update updates the {{ .Type }} in the database.
func ({{ shortname .Type }} *{{ .Type }}) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !{{ shortname .Type }}._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if {{ shortname .Type }}._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE {{ schema .TableSchema .TableName }} SET (` +
		`{{ colnames .Fields .PrimaryKeyField }}` +
		`) = ( ` +
		`{{ colvals .Fields .PrimaryKeyField }}` +
		`) WHERE {{ .PrimaryKey }} = ${{ colcount .Fields .PrimaryKeyField }}`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields (shortname .Type) .PrimaryKeyField }}, {{ shortname .Type }}.{{ .PrimaryKeyField }})
	_, err = db.Exec(sqlstr, {{ fieldnames .Fields (shortname .Type) .PrimaryKeyField }}, {{ shortname .Type }}.{{ .PrimaryKeyField }})
	return err
}

// Save saves the {{ .Type }} to the database.
func ({{ shortname .Type }} *{{ .Type }}) Save(db XODB) error {
	if {{ shortname .Type }}.Exists() {
		return {{ shortname .Type }}.Update(db)
	}

	return {{ shortname .Type }}.Insert(db)
}

// Upsert performs an upsert for {{ .Type }}.
//
// NOTE: PostgreSQL 9.5+ only
func ({{ shortname .Type }} *{{ .Type }}) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if {{ shortname .Type }}._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO {{ schema .TableSchema .TableName }} (` +
		`{{ colnames .Fields }}` +
		`) VALUES (` +
		`{{ colvals .Fields }}` +
		`) ON CONFLICT ({{ .PrimaryKey }}) DO UPDATE SET (` +
		`{{ colnames .Fields }}` +
		`) = (` +
		`{{ colprefixnames .Fields "EXCLUDED" }}` +
		`)`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields (shortname .Type) }})
	_, err = db.Exec(sqlstr, {{ fieldnames .Fields (shortname .Type) }})
	if err != nil {
		return err
	}

	// set existence
	{{ shortname .Type }}._exists = true

	return nil
}

// Delete deletes the {{ .Type }} from the database.
func ({{ shortname .Type }} *{{ .Type }}) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !{{ shortname .Type }}._exists {
		return nil
	}

	// if deleted, bail
	if {{ shortname .Type }}._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM {{ schema .TableSchema .TableName }} WHERE {{ .PrimaryKey }} = $1`

	// run query
	XOLog(sqlstr, {{ shortname .Type }}.{{ .PrimaryKeyField }})
	_, err = db.Exec(sqlstr, {{ shortname .Type }}.{{ .PrimaryKeyField }})
	if err != nil {
		return err
	}

	// set deleted
	{{ shortname .Type }}._deleted = true

	return nil
}
{{- end }}

