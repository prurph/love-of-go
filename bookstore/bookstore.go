package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        Category
}

type Category int

const (
	CategoryAutobiography Category = iota
	CategoryLargePrintRomance
	CategoryParticlePhysics
)

type Catalog map[int]Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No Copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
}

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("Negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

func (b *Book) SetCategory(category Category) error {
	validCategory := map[Category]bool{
		CategoryAutobiography:     true,
		CategoryLargePrintRomance: true,
		CategoryParticlePhysics:   true,
	}
	if !validCategory[category] {
		return fmt.Errorf("Unknown category %v", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() Category {
	return b.category
}
