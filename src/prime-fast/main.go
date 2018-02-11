package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
	"math"
)

type Response struct {
	Host string `json:"host"`
	Duration int64 `json:"duration"`
	Result int64 `json:"result"`
}

func nextPrime(prime int64) int64 {
	if prime > 2 {
		i := int64(-1)
		q := int64(0)

		for i <= q {
			i = 3
			prime += 2
			q = int64(math.Floor(math.Sqrt(float64(prime))))

			for i <= q && prime % i != 0 {
				i+= 2
			}
		}

		return prime;
	}

	if prime == 2 {
		return 3
	}

	return 2
}

func findNthPrime(n int64) int64 {
	if n == 0 {
		return 0
	}

	last := int64(1)
	for ;n > 0; n-- {
		last = nextPrime(last)
	}

	return last
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve parameters
	n, paramErr := strconv.ParseInt(r.URL.Query().Get("n"), 10, 64);

	// Handle parameter errors
	if paramErr != nil {
		fmt.Println("Error: Could not parse parameters")
		w.Write([]byte(paramErr.Error()))
		return;
	}

	// Calculate nth prime
	start := time.Now()
	prime := findNthPrime(n)
	duration := time.Now().Sub(start) / time.Millisecond;

	// Create response
	res := Response{
		Host: os.Getenv("HOSTNAME"), // The kubernetes pod name is available in the env
		Duration: int64(duration),
		Result: prime,
	}

	// Convert to json
	jsonRes, jsonErr := json.Marshal(res)

	// Handle json errors
	if jsonErr != nil {
		fmt.Println("Error: Could not convert response to json")
		w.Write([]byte(jsonErr.Error()))
		return;
	}

	// Send response
	fmt.Printf("Success: %d = %s", n, string(jsonRes))
	fmt.Println()

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Write(jsonRes)
}

func main() {
	// Run a http server on port 8080
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}