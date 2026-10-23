package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	value, _ := reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func readInt(prompt string) (int, error) {
	value := readLine(prompt)
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid whole number")
	}
	return number, nil
}

func readFloat(prompt string) (float64, error) {
	value := readLine(prompt)
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid number")
	}
	return number, nil
}

func readItem() (Item, error) {
	id, err := readInt("Enter item ID: ")
	if err != nil {
		return Item{}, err
	}
	name := readLine("Enter item name: ")
	quantity, err := readInt("Enter quantity: ")
	if err != nil {
		return Item{}, err
	}
	price, err := readFloat("Enter price per unit: ")
	if err != nil {
		return Item{}, err
	}

	item := Item{
		ID:       id,
		Name:     strings.TrimSpace(name),
		Quantity: quantity,
		Price:    price,
	}
	if err := validateItem(item); err != nil {
		return Item{}, err
	}
	return item, nil
}

func listItems(inv *Inventory) {
	if len(inv.Items) == 0 {
		fmt.Println("Inventory khatm hai")
		return
	}
	fmt.Println("\nInventory ke items")
	for _, item := range inv.Items {
		fmt.Println(item)
	}
}

func main() {
	inventory := &Inventory{Items: []Item{}}

	fmt.Println("########################################")
	fmt.Println("   INVENTORY MANAGEMENT SYSTEM")
	fmt.Println("########################################")

	for {
		fmt.Println("\n1 Add item")
		fmt.Println("2 Update item")
		fmt.Println("3 Remove item")
		fmt.Println("4 List items")
		fmt.Println("5 Show summary statistics")
		fmt.Println("0 Exit")

		choice, err := readInt("Choose an option: ")
		if err != nil {
			fmt.Println("error", err)
			continue
		}

		switch choice {
		case 1:
			item, err := readItem()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if err := inventory.AddItem(item); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item added successfully.")
			}

		case 2:
			id, err := readInt("enter the id of item to update ")
			if err != nil {
				fmt.Println("error", err)
				continue
			}
			item, err := readItem()
			if err != nil {
				fmt.Println("Errorrrrrrrrrrr", err)
				continue
			}
			if err := inventory.UpdateItem(id, item); err != nil {
				fmt.Println("Errorrrrrrrrrrr", err)
			} else {
				fmt.Println("Item updated successfully.")
			}

		case 3:
			id, err := readInt("Enter the ID of the item to remove: ")
			if err != nil {
				fmt.Println("Errorrrrrrrrrrrrrrr", err)
				continue
			}
			if err := inventory.RemoveItem(id); err != nil {
				fmt.Println("Errorrrrrrrrrrrrrrrr", err)
			} else {
				fmt.Println("Item removed ")
			}

		case 4:
			listItems(inventory)

		case 5:
			summary := inventory.CalculateSummary()
			fmt.Printf("\nNumber of different items: %d\n", summary.ItemCount)
			fmt.Printf("Total units in stock: %d\n", summary.TotalUnits)
			fmt.Printf("Total inventory value: ₹%.2f\n", summary.TotalValue)

		case 0:
			fmt.Println("Exiting........")
			return

		default:
			fmt.Println("Invalid option. Choose  from 0 to 5")
		}
	}
}
