package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if true {
		fields := []string{"0", "7"}
		c := NewController(2, 10)
		c.ProcessRequest(fields)

		return
	}

	fmt.Println("Enter total no of lifts and floors:")
	var lifts int
	var floors int
	fmt.Scanf("%d%d", &lifts, &floors)
	fmt.Printf("Total lifts: %d Total floors: %d\n", lifts, floors)

	c := NewController(lifts, floors)

	snr := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter the requests in pairs with space:")
		snr.Scan()
		line := snr.Text()
		if line == "exit" {
			for i, l := range c.lifts {
				fmt.Printf("LIFT %d: %d SECONDS\n", i, l.TotalTime)
			}

			break
		}
		fields := strings.Fields(line)
		//fmt.Printf("Fields: %q\n", fields)
		c.ProcessRequest(fields)
	}
}
