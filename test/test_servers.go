package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var port string
var t float64

type Conn struct {
	qty int
	mu  sync.RWMutex
}

var conn Conn = Conn{
	qty: 0,
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port no. to start test server on: ")

	inp, _ := reader.ReadString('\n')
	port = strings.TrimSpace(inp)

	_, err := strconv.Atoi(port)

	if err != nil {
		fmt.Println(port, "is not a valid port number")

	} else {

		fmt.Print("Enter server duration to respond: ")

		dur, _ := reader.ReadString('\n')
		dur = strings.TrimSpace(dur)
		t, err = strconv.ParseFloat(dur, 32)
		if err != nil {
			fmt.Println(port, "is not a valid port number")

		} else {
			http.HandleFunc("/", getSlash)
			fmt.Println("Starting test server on port", port)
			err := http.ListenAndServe(":"+port, nil)
			if err != nil {
				fmt.Println("Error starting server on:", port)
			}
		}
	}
}

func getSlash(w http.ResponseWriter, r *http.Request) {
	conn.mu.Lock()
	conn.qty++
	fmt.Println(strconv.Itoa(conn.qty), "Got request on:", port)
	conn.mu.Unlock()
	time.Sleep(time.Duration(t) * time.Second)
	w.WriteHeader(http.StatusOK)
}
