package main

// RUN echo 'hello' > /tmp/lines && echo 'filter' >> /tmp/lines
// before running program

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		uc1 := strings.ToUpper(scanner.Text())
		fmt.Println(uc1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
