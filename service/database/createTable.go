package database

/* hold for future develop

type Propertis struct {
	Ftype      string
	PrimaryKey bool
	Unique     bool
	NotNull    bool
	Check      string
}

type Field map[string]Propertis

type Table struct {
	NameTable string
	Fields    Field
	Ref       []References
}

type References struct {
	From       Table
	FromFields Field
	To         Table
	ToFields   Field
}

const (
	SQLITE_INTEGER = "INTEGER"
	SQLITE_TEXT    = "TEXT"
	SQLITE_BLOB    = "BLOB"
	SQLITE_REAL    = "REAL"
	SQLITE_NUMERIC = "NUMERIC"
	SQLITE_TIME    = "TIME"
	SQLITE_BOOL    = "BOOL"
	SQLITE_NULL    = "NULL"
)

*/

/*
create a new table for users, if it exists the table will not be created.
The new table will have the fields specified as parameters.

The function will return an error if the table exists or cannot be created. If table will be created the function
will return nil.
*/
/* func createTable(db *sql.DB, table Table) (_error error) {
	var query strings.Builder

	query.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.NameTable))

	for key, property := range table.Fields {
		query.WriteString(key + " ")
		query.WriteString(property.Ftype + " ")

		if property.Check != "" {
			query.WriteString(property.Check + " ")

		}
		if property.NotNull {
			query.WriteString("NOT NULL ")
		}
		if property.PrimaryKey {
			query.WriteString("PRIMARY KEY")
		} else if property.Unique {
			query.WriteString("UNIQUE")
		}

		query.WriteString(", ")

	}

	buff := strings.TrimRight(query.String(), ", ")
	query.Reset()
	query.WriteString(buff)

	for _, constraint := range table.Ref {

		query.WriteString("FOREING KEY(")
		for name := range constraint.FromFields {
			query.WriteString(name + ", ")
		}
		query.WriteString(") REFERENCES " + constraint.To.NameTable + "(")

		for name := range constraint.ToFields {
			query.WriteString(name + ", ")
		}

		query.WriteString(")")

	}
	query.WriteString(");")
	fmt.Println(query.String())
	_, _error = db.Exec(query.String())
	return _error

}
*/
