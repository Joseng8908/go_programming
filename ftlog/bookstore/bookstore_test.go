package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
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
	result := bookstore.Buy(book1)
	got := result.Copies
	if want != got {
		t.Errorf(`
		decreased copies(want) are %d
		real decreased cpoies(got) is %d 
		present copies is %d`, want, got, origin)
	}
}
