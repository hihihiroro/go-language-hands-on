package hello

import (
	"bufio"
	"fmt"
	"os"
)

func Input(msg string) string {
	scanner := bifop.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
