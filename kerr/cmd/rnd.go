package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		f := get()
		clipboard.WriteAll(f)
		fmt.Print(f)
	}
}

func get() string {
	return fmt.Sprintf("\"%s\", ", randomString(10))
}

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
