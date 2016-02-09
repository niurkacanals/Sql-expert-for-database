# About xo #

xo is a cli tool to generate [Golang](https://golang.org/project/) types and
funcs based on a database schema or a custom query. xo is designed to vastly
reduce the overhead/redundancy of writing (from scratch) Go types and funcs for
common database tasks.

Currently, xo can generate types for tables, enums, stored procedures, and
custom SQL queries for PostgreSQL databases. Work is also being done to add
support for MySQL, Oracle, and SQLite, which will be released when they become
feature-complete.

Additionally, support for other database abstractions (ie, views, many-to-many
relationships, etc) are in varying states of completion, and will be added as
soon as they are in a usable state.

Please note that xo is **NOT** an ORM, nor does xo generate an ORM. Instead, xo
generates Go code by using database metadata to query the types and
relationships within the database, and then generates representative Go types
and funcs for well-defined database relationships using raw queries.

# Installation #

Install in the usual way for Go:
```sh
go get -u github.com/knq/xo
```

# Quickstart #

The following is a quick working example of how to use xo:

```sh
# make an output directory
mkdir models

# generate code for schema
xo pgsql://user:pass@host/dbname -o models

# generate code for a custom query
cat << ENDSQL | xo pgsql://user:pass@host/dbname -N -M -B -T AuthorResult -o models/
SELECT
  a.name::varchar AS name,
  b.type::integer AS my_type
FROM authors a
  INNER JOIN authortypes b ON a.id = b.author_id
WHERE
  a.id = %%authorID int%%
LIMIT %%limit int%%
ENDSQL

# build generated code
go build ./models
```

# Command Line #

The following are xo's arguments and options:

```
$ xo -h
usage: xo [--schema SCHEMA] [--out OUT] [--suffix SUFFIX] [--single-file] [--package PACKAGE] [--custom-type-package CUSTOM-TYPE-PACKAGE] [--int32-type INT32-TYPE] [--uint32-type UINT32-TYPE] [--query-mode] [--query QUERY] [--query-type QUERY-TYPE] [--query-func QUERY-FUNC] [--query-only-one] [--query-trim] [--query-strip] [--query-type-comment QUERY-TYPE-COMMENT] [--query-func-comment QUERY-FUNC-COMMENT] [--query-delimiter QUERY-DELIMITER] DSN

positional arguments:
  dsn                    data source name

options:
  --schema SCHEMA, -s SCHEMA
                         schema name to generate Go types for [default: public]
  --out OUT, -o OUT      output path or file name
  --suffix SUFFIX, -f SUFFIX
                         output file suffix [default: .xo.go]
  --single-file          toggle single file output
  --package PACKAGE, -p PACKAGE
                         package name used in generated Go code
  --custom-type-package CUSTOM-TYPE-PACKAGE, -C CUSTOM-TYPE-PACKAGE
                         Go package name to use for custom or unknown types
  --int32-type INT32-TYPE, -i INT32-TYPE
                         Go type to assign to integers [default: int]
  --uint32-type UINT32-TYPE, -u UINT32-TYPE
                         Go type to assign to unsigned integers [default: uint]
  --query-mode, -N       enable query mode
  --query QUERY, -Q QUERY
                         query to generate Go type and func from
  --query-type QUERY-TYPE, -T QUERY-TYPE
                         query's generated Go type
  --query-func QUERY-FUNC, -F QUERY-FUNC
                         query's generated Go func name
  --query-only-one, -1   toggle query's generated Go func to return only one result
  --query-trim, -M       toggle trimming of query whitespace in generated Go code
  --query-strip, -B      toggle stripping '::type AS name' from query in generated Go code
  --query-type-comment QUERY-TYPE-COMMENT
                         comment for query's generated Go type
  --query-func-comment QUERY-FUNC-COMMENT
                         comment for query's generated Go func
  --query-delimiter QUERY-DELIMITER, -D QUERY-DELIMITER
                         delimiter for query's embedded Go parameters [default: %%]
  --help, -h             display this help and exit
```

# End-to-End Example #

For example, given the following PostgreSQL schema:
```PLpgSQL
CREATE TABLE authors (
  author_id SERIAL PRIMARY KEY,
  isbn text NOT NULL DEFAULT '' UNIQUE,
  name text NOT NULL DEFAULT '',
  subject text NOT NULL DEFAULT ''
);

CREATE INDEX authors_name_idx ON authors(name);

CREATE TYPE book_type AS ENUM (
  'FICTION',
  'NONFICTION'
);

CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    author_id integer NOT NULL REFERENCES authors(author_id),
    title text NOT NULL DEFAULT '',
    booktype book_type NOT NULL DEFAULT 'FICTION',
    year integer NOT NULL DEFAULT 2000
);

CREATE INDEX books_title_idx ON books(title, year);

CREATE FUNCTION say_hello(text) RETURNS text AS $$
BEGIN
    RETURN CONCAT('hello ' || $1);
END;
$$ LANGUAGE plpgsql;
```

xo will generate the following (note: this is an abbreviated copy of actual
output -- please see the [example](example) directory for how the generated
types and funcs are used (generated via [example/gen.sh](example/gen.sh)), and
see the [example/models](example/models) directory for the full generated
code):
```go
// Author represents a row from public.authors.
type Author struct {
    AuthorID int    // author_id
    Isbn     string // isbn
    Name     string // name
    Subject  string // subject
}

// Exists determines if the Author exists in the database.
func (a *Author) Exists() bool { /* ... */ }

// Deleted provides information if the Author has been deleted from the database.
func (a *Author) Deleted() bool { /* ... */ }

// Insert inserts the Author to the database.
func (a *Author) Insert(db XODB) error { /* ... */ }

// Update updates the Author in the database.
func (a *Author) Update(db XODB) error { /* ... */ }

// Save saves the Author to the database.
func (a *Author) Save(db XODB) error { /* ... */ }

// Upsert performs an upsert for Author.
func (a *Author) Upsert(db XODB) error { /* ... */ }

// Delete deletes the Author from the database.
func (a *Author) Delete(db XODB) error { /* ... */ }

// AuthorByIsbn retrieves a row from public.authors as a Author.
//
// Looks up using index authors_isbn_key.
func AuthorByIsbn(db XODB, isbn string) (*Author, error) { /* ... */ }

// AuthorsByName retrieves rows from public.authors, each as a Author.
//
// Looks up using index authors_name_idx.
func AuthorsByName(db XODB, name string) ([]*Author, error) { /* ... */ }

// AuthorByAuthorID retrieves a row from public.authors as a Author.
//
// Looks up using index authors_pkey.
func AuthorByAuthorID(db XODB, authorID int) (*Author, error) { /* ... */ }

// Book represents a row from public.books.
type Book struct {
    BookID   int      // book_id
    AuthorID int      // author_id
    Title    string   // title
    Booktype BookType // booktype
    Year     int      // year
}

// Exists determines if the Book exists in the database.
func (b *Book) Exists() bool { /* ... */ }

// Deleted provides information if the Book has been deleted from the database.
func (b *Book) Deleted() bool { /* ... */ }

// Insert inserts the Book to the database.
func (b *Book) Insert(db XODB) error { /* ... */ }

// Update updates the Book in the database.
func (b *Book) Update(db XODB) error { /* ... */ }

// Save saves the Book to the database.
func (b *Book) Save(db XODB) error { /* ... */ }

// Upsert performs an upsert for Book.
func (b *Book) Upsert(db XODB) error { /* ... */ }

// Delete deletes the Book from the database.
func (b *Book) Delete(db XODB) error { /* ... */ }

// Book returns the Author associated with the Book's AuthorID (author_id).
func (b *Book) Author(db XODB) (*Author, error) { /* ... */ }

// BookByBookID retrieves a row from public.books as a Book.
//
// Looks up using index books_pkey.
func BookByBookID(db XODB, bookID int) (*Book, error) { /* ... */ }

// BooksByTitle retrieves rows from public.books, each as a Book.
//
// Looks up using index books_title_idx.
func BooksByTitle(db XODB, title string, year int) ([]*Book, error) { /* ... */ }

// BookType is the 'book_type' enum type.
type BookType uint16

const (
    // FictionBookType is the book_type for 'FICTION'.
    FictionBookType = BookType(1)

    // NonfictionBookType is the book_type for 'NONFICTION'.
    NonfictionBookType = BookType(2)
)

// String returns the string value of the BookType.
func (bt BookType) String() string { /* ... */ }

// MarshalText marshals BookType into text.
func (bt BookType) MarshalText() ([]byte, error) { /* ... */ }

// UnmarshalText unmarshals BookType from text.
func (bt *BookType) UnmarshalText(text []byte) error { /* ... */ }

// SayHello calls the stored procedure 'public.say_hello(text) text' on db.
func SayHello(db XODB, v0 string) (string, error) { /* ... */ }

// XODB is the common interface for database operations that can be used with
// types from public.
//
// This should work with database/sql.DB and database/sql.Tx.
type XODB interface {
    Exec(string, ...interface{}) (sql.Result, error)
    Query(string, ...interface{}) (*sql.Rows, error)
    QueryRow(string, ...interface{}) *sql.Row
}
```

# Design, Origin, Philosophy, and History #

xo can likely get you 99% "of the way there" on medium or large database
schemas and 100% of the way there for small or trivial database schemas. In
short, xo is a great launching point for developing standardized packages for
standard database abstractions/relationships, and xo's most common use-case is
indeed in a code generation pipeline, ala ```stringer```.

**NOTE:** While the code generated by xo is production quality, it is not the
goal, nor the intention for xo to be a "silver bullet," nor to completely
eliminate the manual authoring of SQL / Go code.

xo was originally developed while migrating a "large" application written in
PHP to Go. The schema in use in the original app, while well designed, had
become inconsistent over multiple iterations/generations, mainly due to
different naming styles adopted by various developers/database admins over the
preceding years. Additionally, some components had been written in different
languages (Ruby, Java) and had also had drift from the original application and
schema. Simultaneously, a large amount of growth meant that the PHP/Ruby code
could no longer efficiently serve the traffic volumes.

In late 2014/early 2015, a decision was made to unify and strip out certain
backend services and to fully isolate the API from the original application,
allowing the various parts to instead speak to a common API layer instead of
directly to the database, and to build the service layer in Go. 

However, unraveling the old PHP/Ruby/Java code became a relatively large
headache as the code, the database, and the API, had all experienced
significant drift, and thus underlying function names, fields, and API methods
no longer aligned. As such, after a round of standardizing names, dropping
accumulated cruft, and adding a small number of relationship changes to the
schema, the various codebases were fixed to match the schema changes. After
that was determined to be a success, the next target was a rewrite the backend
services in Go.

In order to keep a similar and consistent workflow for the developers, a code
generator similar to what was previously used with PHP was written for Go.
Additionally, at this time, but tangential to the story here, the API
definitions were ported from JSON to Protobuf to make use of its code
generation abilities as well.

xo is part of the fruits of those development efforts, and it is hoped that
others will be able to use and expand xo to support other databases (SQL or
otherwise).

Part of xo's goal is to avoid writing an ORM, or an ORM-like in Go, and to use
type-safe, fast, and idiomatic Go code. Additionally, the xo developers are of
the opinion that relational databases should have proper, well-designed
relationships and all the related definitions should reside within the database
schema itself -- call it "self-documenting" schema. xo is an end to that
pursuit.

# Similar Projects #
The following projects work with similar concepts as xo:

## Go Generators ##
* [ModelQ](https://github.com/mijia/modelq)
* [sqlgen](https://github.com/drone/sqlgen)
* [squirrel](https://github.com/Masterminds/squirrel)

## Go ORM-likes ##
* [sqlc](https://github.com/relops/sqlc)
