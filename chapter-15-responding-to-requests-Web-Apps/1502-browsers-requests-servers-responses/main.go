package main

import "fmt"

func main() {
	fmt.Println("What happen when we type a URL into the browser?")
	// In the EARLY DAYS of the web, the server reads the content of a HTML file on the server's hard drive,
	// then send that HTML back to the browser:
	// 1. We give the browser the URL with includes the server's address and the name of the thing we're looking for.
	// 2. With that, the browser sends a REQUEST to the server.
	// 3. The SERVER looks up the file the user requested.
	// 4. The SERVER reads the file.
	// 5. The SERVER sends the file's contents back to the browser.
	// 6. The browser displays the RESPONSE.

	// But TODAY, it's much more common for the server to
	// communicate with a PROGRAM to fulfill the requests
	// instead of reading from a file:
	// 1. ... URL ...
	// 2. ... REQUEST ...
	// 3. The server passes the request to a program.
	// 4. The PROGRAM generates an appropriate response.
	// 5. The server sends the response back to the browser.
	// 6. ... display ...
}
