package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Inventory struct {
	Items []Item
}

var (
	ErrInvalidID       = errors.New("ID must be greater than zero")
	ErrInvalidName     = errors.New("name cannot be empty")
	ErrInvalidQuantity = errors.New("quantity cannot be negative")
	ErrInvalidPrice    = errors.New("price cannot be negative")
	ErrDuplicateID     = errors.New("an item with this ID already exists")
	ErrItemNotFound    = errors.New("item not found")
)

func validateItem(item Item) error {
	if item.ID <= 0 {
		return ErrInvalidID
	}
	if item.Name == "" {
		return ErrInvalidName
	}
	if item.Quantity < 0 {
		return ErrInvalidQuantity
	}
	if item.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}

func (inv *Inventory) AddItem(item Item) error {
	if err := validateItem(item); err != nil {
		return err
	}
	for _, existing := range inv.Items {
		if existing.ID == item.ID {
			return ErrDuplicateID
		}
	}
	inv.Items = append(inv.Items, item)
	return nil
}

func (inv *Inventory) UpdateItem(id int, updated Item) error {
	if id <= 0 {
		return ErrInvalidID
	}
	if err := validateItem(updated); err != nil {
		return err
	}
	for i, item := range inv.Items {
		if item.ID == id {
			updated.ID = id
			inv.Items[i] = updated
			return nil
		}
	}
	return ErrItemNotFound
}

func (inv *Inventory) RemoveItem(id int) error {
	if id <= 0 {
		return ErrInvalidID
	}
	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return nil
		}
	}
	return ErrItemNotFound
}

func (inv *Inventory) GetItem(id int) (Item, error) {
	for _, item := range inv.Items {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, ErrItemNotFound
}

type Summary struct {
	ItemCount  int
	TotalUnits int
	TotalValue float64
}

func (inv *Inventory) CalculateSummary() Summary {
	resultChannel := make(chan Summary)

	go func(items []Item) {
		summary := Summary{ItemCount: len(items)}
		for _, item := range items {
			summary.TotalUnits += item.Quantity
			summary.TotalValue += float64(item.Quantity) * item.Price
		}
		resultChannel <- summary
	}(append([]Item(nil), inv.Items...))

	return <-resultChannel
}

func (item Item) String() string {
	return fmt.Sprintf("ID: %d | Name: %s | Quantity: %d | Price: ₹%.2f",
		item.ID, item.Name, item.Quantity, item.Price)
}
