package gui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nicumicle/go-docker/pkg/colors"
)

func Clear() {
	//clear the screen
	fmt.Printf("\x1b[2J") //from top

	fmt.Printf("\x1bc")

	fmt.Println("\033[2J") // from bottom
}

func PressEnterToContinue(msg string) {
	defaultMsg := "Press 'Enter' to continue..."
	if msg == "" {
		msg = defaultMsg
	}
	fmt.Print(msg)
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func ReadCommand(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	return strings.TrimSuffix(text, "\n")
}

func DisplayErrorAndWaitForEnter(text string) {
	fmt.Println(colors.WriteRed(text))
	PressEnterToContinue("")
}
