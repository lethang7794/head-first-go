package main

// When responseSize makes the call to http.Get,
// the program has to SIT THERE and WAIT for the remote website to respond.
// It's NOT DOING anything useful while it waits.

// But the program needs to tell the users that it's working.
// The progress bar needs to change, not just stuck there.

// CONCURRENCY allows a program to PAUSE one task and work on other tasks.
// e.g. while waiting for the network request to complete, it's need to update the progress bar.

// If a program is written to support concurrency,
// then it may support PARALLELISM: running tasks SIMULTANEOUSLY.

// In Go, concurrent tasks are called GOROUTINES.
// Other programming languages have a similar concept called THREADS,
// but goroutines require LESS computer MEMORY than threads,
// and LESS TIME to startup and stop,
// meaning you can run MORE GOROUTINES AT ONCE.

// GOROUTINES
// - allow for concurrency: pausing one task to work on others.
// - and in some situations, parallelism: working on multiple tasks simultaneously!

// Every Go program runs at least one goroutine: the main function.
func main() {
}
