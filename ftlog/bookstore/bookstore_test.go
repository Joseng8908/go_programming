package bookstore_test

import (
	"bookstore"
	"sort"
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
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Title: "For the Love of Go", ID: 1},
		{Title: "The Power of Go: Tools", ID: 2},
	}

	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go",
			ID: 1},
		2: {Title: "The Power of Go: Tools",
			ID: 2},
	}

	want := bookstore.Book{
		Title: "For the Love of Go",
		ID:    1,
	}

	got, err := catalog.GetBook(1)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}

	_, err := catalog.GetBook(999)

	if err == nil {
		t.Fatal("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:           "For the Love of Go",
		Author:          "John Arundel",
		Copies:          4,
		ID:              5,
		PriceCents:      100,
		DiscountPercent: 20,
	}
	want := 80
	got := book.NetPriceCents()
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
