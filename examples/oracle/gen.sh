#!/bin/bash

DBUSER=system
DBPASS=manager
DBHOST=$(docker port orcl 1521)
DBNAME=orcl

SP=$DBUSER/$DBPASS@$DBHOST/$DBNAME
DB=oracle://$DBUSER:$DBPASS@$DBHOST/$DBNAME

EXTRA=$1

SRC=$(realpath $(cd -P "$( dirname "${BASH_SOURCE[0]}" )" && pwd ))

XOBIN=$(which xo)
if [ -e $SRC/../../xo ]; then
  XOBIN=$SRC/../../xo
fi

DEST=$SRC/models

set -x

mkdir -p $DEST
rm -f $DEST/*.go
rm -f $SRC/oracle

sqlplus -S $SP <<< 'DROP INDEX books_title_idx;'
sqlplus -S $SP <<< 'DROP INDEX authors_name_idx;'
sqlplus -S $SP <<< 'DROP TABLE books;'
sqlplus -S $SP <<< 'DROP TABLE authors;'

sqlplus -S $SP << 'ENDSQL'
CREATE TABLE authors (
  author_id integer GENERATED BY DEFAULT ON NULL AS IDENTITY PRIMARY KEY,
  name nvarchar2(255) DEFAULT '' NOT NULL
);

CREATE INDEX authors_name_idx ON authors(name);

CREATE TABLE books (
  book_id integer GENERATED BY DEFAULT ON NULL AS IDENTITY PRIMARY KEY,
  author_id integer REFERENCES authors(author_id) NOT NULL,
  isbn nvarchar2(255) DEFAULT '' UNIQUE NOT NULL,
  booktype book_type DEFAULT 'FICTION' NOT NULL,
  title nvarchar2(255) DEFAULT '' NOT NULL,
  year integer DEFAULT 2000 NOT NULL,
  available timestamp with time zone DEFAULT 'NOW()' NOT NULL,
  tags nvarchar2[] DEFAULT '{}' NOT NULL
);

CREATE INDEX books_title_idx ON books(title, year);

ENDSQL

$XOBIN $DB -o $SRC/models $EXTRA

exit

$XOBIN $DB -N -M -B -T AuthorBookResult --query-type-comment='AuthorBookResult is the result of a search.' -o $SRC/models $EXTRA << ENDSQL
SELECT
  a.author_id::integer AS author_id,
  a.name::text AS author_name,
  b.book_id::integer AS book_id,
  b.isbn::text AS book_isbn,
  b.title::text AS book_title,
  b.tags::text[] AS book_tags
FROM books b
JOIN authors a ON a.author_id = b.author_id
WHERE b.tags && %%tags StringSlice%%::nvarchar2[]
ENDSQL

pushd $SRC &> /dev/null

go build
./oracle $EXTRA

popd &> /dev/null

sqlplus $SP <<< 'select * from books;'
