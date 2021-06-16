package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	models "github.com/xo/xo/_examples/booktest/postgres"
)

func runPostgres(ctx context.Context, db *sql.DB) error {
	// create an author
	a := models.Author{
		Name: "Unknown Master",
	}
	// save author to database
	if err := a.Save(ctx, db); err != nil {
		return err
	}
	// create transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// save first book
	now := time.Now()
	b0 := models.Book{
		AuthorID:  a.AuthorID,
		Isbn:      "1",
		Title:     "my book title",
		BookType:  models.BookTypeFiction,
		Year:      2016,
		Available: now,
	}
	if err := b0.Save(ctx, tx); err != nil {
		return err
	}
	// save second book
	b1 := models.Book{
		AuthorID:  a.AuthorID,
		Isbn:      "2",
		Title:     "the second book",
		BookType:  models.BookTypeFiction,
		Year:      2016,
		Available: now,
		Tags:      models.StringSlice{"cool", "unique"},
	}
	if err := b1.Save(ctx, tx); err != nil {
		return err
	}
	// update the title and tags
	b1.Title = "changed second title"
	b1.Tags = models.StringSlice{"cool", "disastor"}
	if err := b1.Update(ctx, tx); err != nil {
		return err
	}
	// save third book
	b2 := models.Book{
		AuthorID:  a.AuthorID,
		Isbn:      "3",
		Title:     "the third book",
		BookType:  models.BookTypeFiction,
		Year:      2001,
		Available: now,
		Tags:      models.StringSlice{"cool"},
	}
	if err := b2.Save(ctx, tx); err != nil {
		return err
	}
	// save fourth book
	b3 := models.Book{
		AuthorID:  a.AuthorID,
		Isbn:      "4",
		Title:     "4th place finisher",
		BookType:  models.BookTypeNonfiction,
		Year:      2011,
		Available: now,
		Tags:      models.StringSlice{"other"},
	}
	if err := b3.Save(ctx, tx); err != nil {
		return err
	}
	// tx commit
	if err := tx.Commit(); err != nil {
		return err
	}
	// upsert, changing ISBN and title
	b4 := models.Book{
		BookID:    b3.BookID,
		AuthorID:  a.AuthorID,
		Isbn:      "NEW ISBN",
		BookType:  b3.BookType,
		Title:     "never ever gonna finish, a quatrain",
		Year:      b3.Year,
		Available: b3.Available,
		Tags:      models.StringSlice{"someother"},
	}
	if err := b4.Upsert(ctx, db); err != nil {
		return err
	}
	// retrieve first book
	books0, err := models.BooksByTitleYear(ctx, db, "my book title", 2016)
	if err != nil {
		return err
	}
	for _, book := range books0 {
		fmt.Printf("Book %d %q: %q available: %q\n", book.BookID, book.BookType, book.Title, book.Available.Format(time.RFC822Z))
		author, err := book.Author(ctx, db)
		if err != nil {
			return err
		}
		fmt.Printf("Book %d author: %q\n", book.BookID, author.Name)
	}
	// find a book with either "cool" or "other" tag
	fmt.Printf("---------\nTag search results:\n")
	res, err := models.AuthorBookResultsByTags(ctx, db, models.StringSlice{"cool", "other", "someother"})
	if err != nil {
		return err
	}
	for _, ab := range res {
		fmt.Printf("Book %d: %q, Author: %q, ISBN: %q Tags: %q\n", ab.BookID, ab.BookTitle, ab.AuthorName, ab.BookIsbn, ab.BookTags)
	}
	// call say_hello(varchar)
	str, err := models.SayHello(ctx, db, "john")
	if err != nil {
		return err
	}
	fmt.Printf("SayHello response: %q\n", str)
	// get book 4 and delete
	b5, err := models.BookByBookID(ctx, db, 4)
	if err != nil {
		return err
	}
	if err := b5.Delete(ctx, db); err != nil {
		return err
	}
	return nil
}
