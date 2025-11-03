package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			fmt.Print("Enter Book ID: ")
			id, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Title: ")
			title := readLine(reader)
			fmt.Print("Enter Author: ")
			author := readLine(reader)
			library.AddBook(models.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added successfully.")

		case 2:
			fmt.Print("Enter Book ID to remove: ")
			id, _ := strconv.Atoi(readLine(reader))
			err := library.RemoveBook(id)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book removed successfully.")
			}

		case 3:
			fmt.Print("Enter Book ID: ")
			bookID, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}

		case 4:
			fmt.Print("Enter Book ID: ")
			bookID, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned successfully.")
			}

		case 5:
			books := library.ListAvailableBooks()
			fmt.Println("\nAvailable Books:")
			for _, b := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}

		case 6:
			fmt.Print("Enter Member ID: ")
			memberID, _ := strconv.Atoi(readLine(reader))
			books, err := library.ListBorrowedBooks(memberID)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("\nBorrowed Books:")
			for _, b := range books {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}

		case 7:
			fmt.Print("Enter Member ID: ")
			id, _ := strconv.Atoi(readLine(reader))
			fmt.Print("Enter Member Name: ")
			name := readLine(reader)
			library.Members[id] = models.Member{ID: id, Name: name}
			fmt.Println("Member added successfully.")

		case 0:
			fmt.Println("Exiting... Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
