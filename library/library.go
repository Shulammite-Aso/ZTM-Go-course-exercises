//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type CheckinAndCheckoutRecord struct {
	checkedOut                bool
	checkoutTime, checkinTime time.Time
}

type MemberName string
type BookTitle string

type Member struct {
	name         string
	bookBorrowed BookTitle
}

type Library struct {
	allBooks    map[BookTitle]CheckinAndCheckoutRecord
	members     map[MemberName]Member
	lendedBooks uint
}

func checkoutBook(library *Library, book BookTitle, member MemberName) {
	if library.members[member].bookBorrowed != "" {
		fmt.Println("You can't borrow a book until you have returned the last one you borrowed.")
		return
	}
	fmt.Println("You can borrow a book.")
	for title, recordHistory := range library.allBooks {

		if title == book && !recordHistory.checkedOut {
			fmt.Println("You can borrow this book.")
			library.members[member] = Member{name: string(member), bookBorrowed: book}
			library.allBooks[title] = CheckinAndCheckoutRecord{checkedOut: true, checkoutTime: time.Now()}
			library.lendedBooks += 1
		}

	}
}

func checkInBook(library *Library, book BookTitle, member MemberName) {
	if library.members[member].bookBorrowed == "" {
		fmt.Println("You can't checkin a book because you don't have any borrowed book")
		return
	}
	fmt.Println("Checking if this is the book you borrowed...")
	for title, recordHistory := range library.allBooks {

		if title == book && recordHistory.checkedOut {
			fmt.Println("Yes you borrowed this book, thanks for returning it.")
			library.members[member] = Member{name: string(member), bookBorrowed: ""}
			library.allBooks[title] = CheckinAndCheckoutRecord{checkedOut: false, checkinTime: time.Now()}
			library.lendedBooks -= 1
		}

	}
}

func main() {
	mainLibrary := Library{
		allBooks: map[BookTitle]CheckinAndCheckoutRecord{
			"Book One":   {},
			"Book Two":   {},
			"Book three": {},
			"book Four":  {},
		},

		members: map[MemberName]Member{
			"Adaobi john": {name: "Adaobi john"},
			"Cynthia Roy": {name: "Cynthia Roy"},
			"Yul Edochie": {name: "Yul Edochie"},
		},
	}
	// Print Out Initial State
	fmt.Println("Book list:", mainLibrary.allBooks)
	fmt.Println()
	fmt.Println("Member list:", mainLibrary.members)
	fmt.Println()
	fmt.Println("Number of books lended:", mainLibrary.lendedBooks)
	fmt.Println()

	// Checkout one book
	checkoutBook(&mainLibrary, "book Four", "Adaobi john")
	fmt.Println()

	fmt.Println("Book list:", mainLibrary.allBooks)
	fmt.Println()
	fmt.Println("Member list:", mainLibrary.members)
	fmt.Println()
	fmt.Println("Number of books lended:", mainLibrary.lendedBooks)
	fmt.Println()

	// Checkout Second book
	checkoutBook(&mainLibrary, "Book Two", "Yul Edochie")
	fmt.Println()

	fmt.Println("Book list:", mainLibrary.allBooks)
	fmt.Println()
	fmt.Println("Member list:", mainLibrary.members)
	fmt.Println()
	fmt.Println("Number of books lended:", mainLibrary.lendedBooks)
	fmt.Println()

	// Return one checkout book
	checkInBook(&mainLibrary, "book Four", "Adaobi john")
	fmt.Println()

	fmt.Println("Book list:", mainLibrary.allBooks)
	fmt.Println()
	fmt.Println("Member list:", mainLibrary.members)
	fmt.Println()
	fmt.Println("Number of books lended:", mainLibrary.lendedBooks)
	fmt.Println()

}
