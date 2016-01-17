package main // import "kego.io/kerr/cmd/kerr"

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"kego.io/kerr/util"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for {
		f := util.GetRandomId()
		clipboard.WriteAll(fmt.Sprintf("\"%s\", ", f))
		fmt.Print(fmt.Sprintf("%s (copied to clipboard)", f))
		scanner.Scan()
	}
}
