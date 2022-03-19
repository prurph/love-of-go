package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	book := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := book
	want.Copies--
	got, err := bookstore.Buy(book)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("Buy(%v): want %v, got %v", book, want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(book)
	if err == nil {
		t.Error("want error buying book with zero remaining copies, but got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}
	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}
	want := bookstore.Book{ID: 2, Title: "The Power of Go: Tools"}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(1)
	if err == nil {
		t.Error("want error for non-existent ID, got nil")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "For the Love of Go",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("with price %d, after %d%% discount want %d, got %d", b.PriceCents, b.DiscountPercent, want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the Love of Go"}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the Love of Go"}
	err := b.SetPriceCents(-1000)
	if err == nil {
		t.Error("want error for negative price, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the Love of Go"}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{Title: "For the Love of Go"}
	err := b.SetCategory("asdf1234watermelonyum")
	if err == nil {
		t.Error("want error for invalid category, got nil")
	}
}
