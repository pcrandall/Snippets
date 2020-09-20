package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	for {
		refresh("Press Enter to refresh")

		fmt.Println("I'm so refreshed")
	}
}

func refresh(s string) {
	r := bufio.NewReader(os.Stdin)
	for {

		fmt.Printf("%s\n", s)

		_, err := r.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		return

	}

}
