package main

//uses of context package

//cancelations: Stops long running applications when client disconnects or time runs out
//timeouts & deadlines: Automatically stops after a certain time

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context.withTimeout (enforcing timeouts for HTTP requests)
func contHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	result := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		result <- "Process completed"
	}()

	select {
	case res := <-result:
		fmt.Fprintln(w, res)
	case <-ctx.Done():
		http.Error(w, "Request timed out", http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/", contHandler)
	http.ListenAndServe(":8000", nil)

}
