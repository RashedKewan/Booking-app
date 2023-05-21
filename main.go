/*
	Initiate our go app inside  project/module
	> go mod init < project name>

	Explain:
	Create a new module
	Module path can be correspond to a repository you plan to publish your module to( e.g. github.com/booking-app)
	Initialize a go.mod file
	Where the go.mod file describes the module: with name/module path and go version used in program
*/

/*
	To Run a specific  Program:
	> go run <file name>.go

	or multiple prograns:
	> go run .
*/

/*
All our code must belong to a package
Go programs are organized into packages
A package is a collection of Go files
*/
package main

// fmt stands for Format package that provides us with functions that we can use
import (
	"Booking-app/validator"
	"fmt"
	"time"
	"sync"
)

// Note       : Go compile Errors to enforce better code quality
//	            Variable names must be used
//			    to avoid possible dead code, code that is never used in the code
//
// For summary: if you declare a variable without using it, it would be an error
// Important  : we can only define a variable without specifing its type just in
//
//	case we assign a value to it so Go can imply the data type based on the value
//
// var conferenceName = "Go Conference"  equals to =>  conferenceName := "Go Conference"
// this thing just for var not available for const
var conferenceName string = "Go Conference"
var remainingTickets uint = 50
const conferenceTickets = 50

// Mixed data type
// type < struct name > struct{ ... }
type UserData struct{
	firstName string
	lastName  string
	email     string
	tickets   uint
}


var wg = sync.WaitGroup{}

/*
Go gonna asks it self where to start the program? where is the entrypoint(main)
main func is the entrypoint of go program
*/
func main() {
	// Array
	// var bookings [50]string
	// Slices <- Dynamic array , adding elements via append() method which adds to the end of the array
	// if we want to define an empty slice we can do as following:
	//     bookings := []map[string]string{}
	// var bookings  = []map[string]string{}
	//     bookings := make([]map[string]string, 0)
	// var bookings []map[string]string -> this is a slice of maps as a type
	// but now we want to use the struct data type
	bookings := make( []UserData, 0)
	
	
	greetUsers()

	var firstName   string
	var lastName    string
	var email       string
	var userTickets uint

	// pointers in Go like in C, C++ => (&) for address , (*) for the value we are pointing on
	var firstName_pointr 		 = &firstName
	var lastName_pointr			 = &lastName
	var email_pointr 			 = &email
	var userTickets_pointr       = &userTickets
	var remainingTickets_pointer = &remainingTickets

	var dataSaved = false

	for {
		if !dataSaved {
			getUserInput(firstName_pointr, lastName_pointr, email_pointr, userTickets_pointr)
		}

		isValisName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		isValidInput := isValisName && isValidEmail && isValidTicketNumber
		if isValidInput {

			updateRemainingTickets(remainingTickets_pointer, userTickets_pointr)
			bookTicket(&bookings, firstName, lastName, userTickets, email, remainingTickets)
			wg.Add(1) // one new thread to add over the main thread

			go sendTickets(firstName , lastName , userTickets , email )
			printAllBookings(bookings)
			if helper.IsConferenceBookedOut(remainingTickets) {
				break
			}
			dataSaved = false

		} else {

			dataSaved = true
			helper.HandleInValidUserInput(firstName_pointr, lastName_pointr, email_pointr, userTickets_pointr)

		}
	}

	wg.Wait()
}



func getUserInput(firstName_pointr *string, lastName_pointr *string, email_pointr *string, userTickets_pointr *uint) {
	helper.Input(firstName_pointr  , "Enter your first name    :")
	helper.Input(lastName_pointr   , "Enter your last name     :")
	helper.Input(email_pointr      , "Enter your email address :")
	helper.Input(userTickets_pointr, "Enter number of tickets  :")
}


// // For bookings as a slice of map data type
// func printAllBookings(bookings []map[string]string) {
// 	for _,booking := range bookings{
		
// 		fmt.Println("firstName :", booking["firstName"])
// 		fmt.Println("lastName  :", booking["lastName"])
// 		fmt.Println("email     :", booking["email"])
// 		fmt.Println("tickets   :", booking["tickets"])

// 	}
// }

func printAllBookings(bookings []UserData) {
	for _,booking := range bookings{
		
		fmt.Println("firstName :", booking.firstName)
		fmt.Println("lastName  :", booking.lastName)
		fmt.Println("email     :", booking.email)
		fmt.Println("tickets   :", booking.tickets)

	}
}



func sendConfirmationOutpot(firstName string, lastName string, userTickets uint, email string, remainingTickets uint) {
	fmt.Printf("\nThank you %v %v  for booking %v tickets. You will receive a confirmation email at: %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
}



func sendTickets(firstName string, lastName string, userTickets uint, email string){
	fmt.Println("Sending ...")
	time.Sleep(5 * time.Second)
	tickets := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("*******************************")
	fmt.Println("Sending tickets :", tickets     )
	fmt.Println("To email address:", email       )
	fmt.Println("*******************************")
	wg.Done() // removes the thread from the waiting list , counter?
}


func updateRemainingTickets(remainingTickets_pointer *uint, userTickets_pointr *uint) {
	*remainingTickets_pointer = *remainingTickets_pointer - *userTickets_pointr
}



func bookTicket(bookings *[]UserData, firstName string, lastName string, userTickets uint, email string, remainingTickets uint) {
	userData := UserData{
		
		firstName:firstName,
		lastName:lastName ,
		email:email,
		tickets:userTickets,
	}

	*bookings = append(*bookings, userData)
	sendConfirmationOutpot(firstName, lastName, userTickets, email, remainingTickets)
}

// func bookTicket(bookings *[]map[string]string, firstName string, lastName string, userTickets uint, email string, remainingTickets uint) {
// 	userData := make(map[string]string)

// 	userData["firstName"] = firstName
// 	userData["lastName" ] = lastName 
// 	userData["email"    ] = email
// 	userData["tickets"  ] = strconv.FormatUint( uint64(userTickets), 10)
	 
// 	*bookings = append(*bookings, userData)
// 	sendConfirmationOutpot(firstName, lastName, userTickets, email, remainingTickets)
// }



func greetUsers() {

	/*	printing using println	*/
	// fmt.Println("Wellcome to", conferenceName ,"booking application")
	// fmt.Println("We have total of", conferenceTickets,"tickits and", remainingTickets,"are still available")
	// fmt.Println("Get your tickets here to attend")

	/*	printing using printf	*/
	fmt.Printf("Wellcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickits and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")
}
