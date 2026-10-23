package main

import (
	"errors"
	"testing"
)

func TestAddItem(t *testing.T) {
	inv := &Inventory{}
	item := Item{ID: 1, Name: "Keyboard", Quantity: 4, Price: 750}

	if err := inv.AddItem(item); err != nil {
		t.Fatalf("AddItem returned unexpected error: %v", err)
	}
	if len(inv.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(inv.Items))
	}
	if inv.Items[0].Name != "Keyboard" {
		t.Errorf("expected Keyboard, got %q", inv.Items[0].Name)
	}
}

func TestAddDuplicateID(t *testing.T) {
	inv := &Inventory{}
	item := Item{ID: 1, Name: "Mouse", Quantity: 2, Price: 300}
	if err := inv.AddItem(item); err != nil {
		t.Fatal(err)
	}
	err := inv.AddItem(Item{ID: 1, Name: "Monitor", Quantity: 1, Price: 8000})
	if !errors.Is(err, ErrDuplicateID) {
		t.Errorf("expected ErrDuplicateID, got %v", err)
	}
}

func TestUpdateItem(t *testing.T) {
	inv := &Inventory{}
	_ = inv.AddItem(Item{ID: 1, Name: "Mouse", Quantity: 2, Price: 300})

	err := inv.UpdateItem(1, Item{ID: 99, Name: "Wireless Mouse", Quantity: 5, Price: 600})
	if err != nil {
		t.Fatalf("UpdateItem returned unexpected error: %v", err)
	}
	got, err := inv.GetItem(1)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "Wireless Mouse" || got.Quantity != 5 {
		t.Errorf("item was not updated correctly: %+v", got)
	}
	if got.ID != 1 {
		t.Errorf("update should preserve original ID 1, got %d", got.ID)
	}
}

func TestRemoveItem(t *testing.T) {
	inv := &Inventory{}
	_ = inv.AddItem(Item{ID: 1, Name: "Cable", Quantity: 3, Price: 100})

	if err := inv.RemoveItem(1); err != nil {
		t.Fatalf("RemoveItem returned unexpected error: %v", err)
	}
	if len(inv.Items) != 0 {
		t.Errorf("expected empty inventory, got %d items", len(inv.Items))
	}
}

func TestCalculateSummary(t *testing.T) {
	inv := &Inventory{}
	_ = inv.AddItem(Item{ID: 1, Name: "Notebook", Quantity: 10, Price: 40})
	_ = inv.AddItem(Item{ID: 2, Name: "Pen", Quantity: 5, Price: 10})

	got := inv.CalculateSummary()
	if got.ItemCount != 2 {
		t.Errorf("expected 2 item types, got %d", got.ItemCount)
	}
	if got.TotalUnits != 15 {
		t.Errorf("expected 15 total units, got %d", got.TotalUnits)
	}
	if got.TotalValue != 450 {
		t.Errorf("expected total value 450, got %.2f", got.TotalValue)
	}
}
