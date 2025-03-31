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
