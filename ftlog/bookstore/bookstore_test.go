package bookstore_test

import (
	"bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
		ID:     1,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	book1 := bookstore.Book{
		Title:  "Vegetarian",
		Author: "Hangang",
		Copies: 5,
	}
	origin := book1.Copies
	want := 4
	result, err := bookstore.Buy(book1)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf(`
		decreased copies(want) are %d
		real decreased cpoies(got) is %d 
		present copies is %d`, want, got, origin)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}

	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := bookstore.GetAllBooks(catalog)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := []bookstore.Book{
		{Title: "For the Love of Go",
			ID: 1},
		{Title: "The Power of Go: Tools",
			ID: 2},
	}

	want := bookstore.Book{
		Title: "For the Love of Go",
		ID:    1,
	}

	got := bookstore.GetBook(catalog, 1)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
