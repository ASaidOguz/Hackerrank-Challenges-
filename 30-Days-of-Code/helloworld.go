package main

import (
	"bufio"
	"fmt"

	//"strconv"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	input := reader.Text()

	fmt.Println("Hello, World.")
	fmt.Println(input)
}
