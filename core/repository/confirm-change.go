package repository

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskForConfirmation(newPCName string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Nuevo Nombre del PC:", newPCName)
	fmt.Println("Estas seguro de que deseas cambiar el nombre del PC? (s/n): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	return response == "s" || response == "S"
}
