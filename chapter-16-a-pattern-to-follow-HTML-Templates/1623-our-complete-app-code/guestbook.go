package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// GuestBook is a struct used in rendering view.html
type GuestBook struct {
	SignatureCount int      // Hold the total number of signatures.
	Signatures     []string // Hold the signatures themselves.
}

// check calls log.Fatal on any non-nill error.
func check(err error) { // Call this when we need to check an error value returned from a function or a method.
	if err != nil { // Most of the time the value will be nill, but if not ...
		log.Fatal(err) // ... output the error and exit the program.
	}
}

// viewHandler reads guestbook signatures and displays them together
// with a count of all signatures.
func viewHandler(writer http.ResponseWriter, request *http.Request) { // Accept a http.ResponseWrite and a *http.Request.
	// Read signatures from a file.
	signatures := getStrings("signatures.txt")

	// Create a template based on the content of view.html.
	html, err := template.ParseFiles("view.html")
	check(err)

	// Store the signatures data in a struct
	guestBook := GuestBook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	// Insert the Guestbook struct data into the template and write the result out to the ResponseWriter.
	err = html.Execute(writer, guestBook)
	check(err)
}

// newHandler displays a form to enter a signature.
func newHandler(writer http.ResponseWriter, request *http.Request) {
	// Load the HTML form in from a template.
	html, err := template.ParseFiles("new.html")
	check(err)

	// Write the template to the ResponseWriter (there is no data to insert).
	err = html.Execute(writer, nil)
	check(err)
}

// createHandler takes a POST request with a signature to add, and
// appends it to the signature file.
func createHandler(writer http.ResponseWriter, request *http.Request) {
	// Get the value of the "signature" form field.
	signature := request.FormValue("signature")

	// Open the file for writing. If it exists, append to it. If not, create it.
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)

	// Add the form field contents to the file.
	_, err = fmt.Fprintln(file, signature)
	check(err)

	// Close the file.
	err = file.Close()
	check(err)

	// Redirect the browser to the main guestbook page.
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

// getStrings returns a slice of strings read from fileName,
// one string per line.
func getStrings(fileName string) []string {
	// Each line of the file will be appended to this slice as a string.
	var lines []string

	// Open the file
	file, err := os.Open(fileName)
	// If we get an error indicating the file doesn't exist...
	if os.IsNotExist(err) {
		return nil // ... return nil instead of a slice.
	}
	// All other errors should be checked and reported normally.
	check(err)
	// After the function exits, ensure the file is closed.
	defer file.Close()

	// Create a scanner for the file contents.
	scanner := bufio.NewScanner(file)
	// For each line of the file...
	for scanner.Scan() {
		// ... append its texts to the slice.
		lines = append(lines, scanner.Text())
	}
	// Report any errors encountered while scanning.
	check(scanner.Err())
	// Return the slice of strings.
	return lines
}

func main() {
	// Requests to view the signature list will be handled by the viewHandler function.
	http.HandleFunc("/guestbook", viewHandler)

	// Requests to get the HTML fom will be handled by newHandler.
	http.HandleFunc("/guestbook/new", newHandler)

	// Requests to submit the form will be handled by createHandler.
	http.HandleFunc("/guestbook/create", createHandler)

	// Loop forever, passing HTTP requests to the appropriate function for handling.
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
