package main

import (
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Price    float64
	Category string
	Date     int
}

type Transaction struct {
	ID         int
	Items      []Item
	TotalPrice float64
	Date       int
}

type Store struct {
	Items         []Item
	Transactions  []Transaction
	Capital       float64
	MostSoldItems []Item
}

func main() {
	var s Store

	for {
		fmt.Println("================================")
		fmt.Println("===       Final Project      ===")
		fmt.Println("===        Programming       ===")
		fmt.Println("===         Algorithm        ===")
		fmt.Println("===                          ===")
		fmt.Println("================================")
		fmt.Println("")
		fmt.Println("===============================")
		fmt.Println("=== Store Management System ===")
		fmt.Println("===============================")
		fmt.Println("")
		fmt.Println("1. Add Item")
		fmt.Println("2. Remove Item")
		fmt.Println("3. Edit Item")
		fmt.Println("4. Add Transaction")
		fmt.Println("5. Remove Transaction")
		fmt.Println("6. Edit Transaction")
		fmt.Println("7. Display Items")
		fmt.Println("8. Display Transactions")
		fmt.Println("9. Display Total Sales")
		fmt.Println("10. Display Most Often Sold Items")
		fmt.Println("11. Display Items by Category")
		fmt.Println("12. Find Items by Keyword")
		fmt.Println("13. Display Total Capital")
		fmt.Println("14. Search Trasaction by Date")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("===============================")
			fmt.Println("===      Add Item Menu      ===")
			fmt.Println("===============================")
			item := getItemDetailsFromUser()
			AddItem(&s, item)
			fmt.Println("Item added successfully.")
		} else if choice == 2 {
			fmt.Println("===============================")
			fmt.Println("===    Remove Item Menu     ===")
			fmt.Println("===============================")
			itemID := getItemIDFromUser()
			RemoveItem(&s, itemID)
			fmt.Println("Item removed successfully.")
		} else if choice == 3 {
			fmt.Println("===============================")
			fmt.Println("===      Edit Item Menu     ===")
			fmt.Println("===============================")
			itemID := getItemIDFromUser()
			newItem := getItemDetailsFromUser()
			EditItem(&s, itemID, newItem)
			fmt.Println("Item edited successfully.")
		} else if choice == 4 {
			fmt.Println("===============================")
			fmt.Println("===     Add Transaction     ===")
			fmt.Println("===============================")
			transaction := getTransactionDetailsFromUser(&s, s.Items)
			AddTransaction(&s, transaction)
			fmt.Println("Transaction added successfully.")
		} else if choice == 5 {
			fmt.Println("===============================")
			fmt.Println("===    Remove Transaction   ===")
			fmt.Println("===============================")
			transactionID := getTransactionIDFromUser()
			RemoveTransaction(&s, transactionID)
			fmt.Println("Transaction removed successfully.")
		} else if choice == 6 {
			fmt.Println("===============================")
			fmt.Println("===     Edit Transaction    ===")
			fmt.Println("===============================")
			transactionID := getTransactionIDFromUser()
			newTransaction := getTransactionDetailsFromUser(&s, s.Items)
			EditTransaction(&s, transactionID, newTransaction)
			fmt.Println("Transaction edited successfully.")
		} else if choice == 7 {
			fmt.Println("===============================")
			fmt.Println("===   Display Transaction   ===")
			fmt.Println("===============================")
			fmt.Println("Items:")
			DisplayItems(&s)
		} else if choice == 8 {
			fmt.Println("===============================")
			fmt.Println("===   Display Transaction   ===")
			fmt.Println("===============================")
			fmt.Println("Transactions:")
			sortByDate(&s, s.Transactions)
			DisplayTransactions(&s)
		} else if choice == 9 {
			fmt.Println("===============================")
			fmt.Println("===   Display Total Sales   ===")
			fmt.Println("===============================")
			fmt.Println("Total Capital:", s.Capital)
		} else if choice == 10 {
			fmt.Println("===============================")
			fmt.Println("=Display Most Often Sold Items=")
			fmt.Println("===============================")
			fmt.Println("Most Often Sold Items:")
			DisplayMostSoldItems(&s)
		} else if choice == 11 {
			fmt.Println("===============================")
			fmt.Println("==  Display Item By Category ==")
			fmt.Println("===============================")
			category := getCategoryFromUser()
			fmt.Println("Items in Category:", category)
			DisplayItemsByCategory(&s, category)
		} else if choice == 12 {
			fmt.Println("===============================")
			fmt.Println("===   Find Item By Keyword  ===")
			fmt.Println("===============================")
			keyword := getKeywordFromUser()
			fmt.Println("Items matching keyword:", keyword)
			FindItemsByKeyword(&s, keyword)
		} else if choice == 13 {
			fmt.Println("===============================")
			fmt.Println("===  Display Total Capital  ===")
			fmt.Println("===============================")
			capital := getCapitalFromUser()
			fmt.Println("Total Capital: ", capital)
		} else if choice == 14 {
			fmt.Println("===============================")
			fmt.Println("= Search Transaction By Date  =")
			fmt.Println("===============================")
			nDate := getDateFromUser()
			fmt.Println("")
			fmt.Println("Chosen Date: ", nDate)
			// SearchTransactionByDate(&s, nDate)
			DisplayItemsByDate(&s, nDate)
		} else if choice == 0 {
			fmt.Println("Exiting...")
			return
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}

func getItemDetailsFromUser() Item {
	var item Item

	fmt.Print("Enter item ID: ")
	fmt.Scanln(&item.ID)

	fmt.Print("Enter item name: ")
	fmt.Scanln(&item.Name)

	fmt.Print("Enter item price: ")
	fmt.Scanln(&item.Price)

	fmt.Print("Enter item category: ")
	fmt.Scanln(&item.Category)

	return item
}

func getItemIDFromUser() int {
	var itemID int
	fmt.Print("Enter item ID: ")
	fmt.Scanln(&itemID)
	return itemID
}

func getTransactionDetailsFromUser(store *Store, items []Item) Transaction {
	var transaction Transaction
	var total float64 = 0

	fmt.Print("Enter transaction date: ")
	fmt.Scanln(&transaction.Date)

	fmt.Print("Enter transaction ID: ")
	fmt.Scanln(&transaction.ID)

	fmt.Print("Enter number of items in the transaction: ")
	var numItems int
	fmt.Scanln(&numItems)

	i := 0
	for i < numItems {
		fmt.Printf("Enter item %d ID: ", i+1)
		var itemID int
		fmt.Scanln(&itemID)

		item := findItemByID(store, items, itemID)
		if item == nil {
			fmt.Printf("Item with ID %d does not exist.\n", itemID)
		} else {
			total += item.Price
			transaction.Items = append(transaction.Items, *item)
			i++
		}
	}
	transaction.TotalPrice = total

	fmt.Println("Total price of the transaction:", transaction.TotalPrice)
	return transaction
}

func sortByDate(store *Store, transactions []Transaction) {
	var j, idxmin int

	i := 0
	N := len(transactions)
	for i < N {
		idxmin = i
		j = i + 1
		for j < N {
			if transactions[idxmin].Date > transactions[j].Date {
				idxmin = j
			}
			j++
		}
		transactions[i], transactions[idxmin] = transactions[idxmin], transactions[i]
		i++
	}
}

func getTransactionIDFromUser() int {
	var transactionID int
	fmt.Print("Enter transaction ID: ")
	fmt.Scanln(&transactionID)
	return transactionID
}

func getCategoryFromUser() string {
	var category string
	fmt.Print("Enter category: ")
	fmt.Scanln(&category)
	return category
}

func getKeywordFromUser() string {
	var keyword string
	fmt.Print("Enter keyword: ")
	fmt.Scanln(&keyword)
	return keyword
}

func getCapitalFromUser() int {
	var capital, totalCap int

	totalCap = 0
	capital = 0
	fmt.Println("Enter capital: ")
	for capital != -1 {
		totalCap += capital
		fmt.Scanln(&capital)
	}
	return totalCap
}

func findItemByID(store *Store, items []Item, itemID int) *Item {
	var result *Item
	for i := 0; i < len(items); i++ {
		if items[i].ID == itemID {
			result = &items[i]
		}
	}
	return result
}

func AddItem(store *Store, item Item) {
	store.Items = append(store.Items, item)
}

func RemoveItem(store *Store, itemID int) {
	var i int
	i = 0
	for i < len(store.Items) {
		if store.Items[i].ID == itemID {
			store.Items = append(store.Items[:i], store.Items[i+1:]...)
		}
		i++
	}
}

func EditItem(store *Store, itemID int, newItem Item) {
	i := 0
	for i < len(store.Items) && store.Items[i].ID != itemID {
		store.Items[i] = newItem
		i++
	}
}

func AddTransaction(store *Store, transaction Transaction) {
	store.Transactions = append(store.Transactions, transaction)
	UpdateCapital(&*store, transaction.TotalPrice)
	UpdateMostSoldItems(&*store, transaction.Items)
}

func RemoveTransaction(store *Store, transactionID int) {
	i := 0
	found := false

	for i < len(store.Transactions) && !found {
		if store.Transactions[i].ID == transactionID {
			UpdateCapital(&*store, store.Transactions[i].TotalPrice)
			store.Transactions = append(store.Transactions[:i], store.Transactions[i+1:]...)
			found = true
		}
		i++
	}
}

func EditTransaction(store *Store, transactionID int, newTransaction Transaction) {
	i := 0
	found := false

	for i < len(store.Transactions) && !found {
		if store.Transactions[i].ID == transactionID {
			UpdateCapital(&*store, store.Transactions[i].TotalPrice)
			store.Transactions[i] = newTransaction
			UpdateCapital(&*store, newTransaction.TotalPrice)
			UpdateMostSoldItems(&*store, newTransaction.Items)
			found = true
		}
		i++
	}
}

func UpdateCapital(store *Store, amount float64) {
	store.Capital += amount
}

func UpdateMostSoldItems(store *Store, items []Item) {
	for i := 0; i < len(items); i++ {
		existingItem := findItemByID(&*store, store.MostSoldItems, items[i].ID)
		if existingItem == nil {
			store.MostSoldItems = append(store.MostSoldItems, items[i])
		} else {
			existingItem.ID = items[i].ID
			existingItem.Name = items[i].Name
			existingItem.Price = items[i].Price
			existingItem.Category = items[i].Category
		}
	}
}

func DisplayItems(store *Store) {
	for i := 0; i < len(store.Items); i++ {
		fmt.Println("Item ID:", store.Items[i].ID)
		fmt.Println("Name:", store.Items[i].Name)
		fmt.Println("Price:", store.Items[i].Price)
		fmt.Println("Category:", store.Items[i].Category)
		fmt.Println()
	}
}

func DisplayTransactions(store *Store) {
	for i := 0; i < len(store.Transactions); i++ {
		fmt.Println("Transaction ID:", store.Transactions[i].ID)
		fmt.Println("Total Price:", store.Transactions[i].TotalPrice)
		fmt.Println("Items:")
		for j := 0; j < len(store.Transactions[i].Items); j++ {
			fmt.Println("Date: ", store.Transactions[i].Date)
			fmt.Println("Item ID:", store.Transactions[i].Items[j].ID)
			fmt.Println("Name:", store.Transactions[i].Items[j].Name)
			fmt.Println("Price:", store.Transactions[i].Items[j].Price)
			fmt.Println("Category:", store.Transactions[i].Items[j].Category)
			fmt.Println()
		}
		fmt.Println()
	}
}

func DisplayMostSoldItems(store *Store) {
	for i := 0; i < len(store.MostSoldItems); i++ {
		fmt.Println("Item ID:", store.MostSoldItems[i].ID)
		fmt.Println("Name:", store.MostSoldItems[i].Name)
		fmt.Println("Price:", store.MostSoldItems[i].Price)
		fmt.Println("Category:", store.MostSoldItems[i].Category)
		fmt.Println()
	}
}

func DisplayItemsByCategory(store *Store, category string) {
	for i := 0; i < len(store.Items); i++ {
		if store.Items[i].Category == category {
			fmt.Println("Item ID:", store.Items[i].ID)
			fmt.Println("Name:", store.Items[i].Name)
			fmt.Println("Price:", store.Items[i].Price)
			fmt.Println("Category:", store.Items[i].Category)
			fmt.Println()
		}
	}
}

func DisplayItemsByDate(store *Store, nDate int) {

	left := 0
	right := len(store.Transactions) - 1
	found := false

	for left <= right && !found {
		mid := (left + right) / 2

		if store.Transactions[mid].Date < nDate {
			left = mid + 1
		} else if store.Transactions[mid].Date > nDate {
			right = mid - 1
		} else {
			// Found a match, display the item details

			fmt.Println("Item ID:", store.Items[med].ID)
			fmt.Println("Name:", store.Items[med].Name)
			fmt.Println("Price:", store.Items[med].Price)
			fmt.Println("Category:", store.Items[med].Category)
			fmt.Println()

			found = true
		}
	}

	if !found {
		fmt.Println("No items found for the specified date.")
	}
}

func FindItemsByKeyword(store *Store, keyword string) {
	for i := 0; i < len(store.Items); i++ {
		if containsKeyword(store.Items[i], keyword) {
			fmt.Println("Item ID:", store.Items[i].ID)
			fmt.Println("Name:", store.Items[i].Name)
			fmt.Println("Price:", store.Items[i].Price)
			fmt.Println("Category:", store.Items[i].Category)
			fmt.Println()
		}
	}
}

func containsKeyword(item Item, keyword string) bool {
	return containsSubstring(item.Name, keyword) || containsSubstring(item.Category, keyword)
}

func containsSubstring(str string, substr string) bool {
	N := len(str)
	M := len(substr)

	found := false
	k := 0

	for !found && k < N-M+1 {
		found = str[k:k+M] == substr
		k = k + 1
	}

	return found
}

func getDateFromUser() int {
	var nDate int

	fmt.Print("Enter the date to search: ")
	fmt.Scanln(&nDate)
	return nDate
}
