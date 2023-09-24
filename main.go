package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getPlaylist() playlist {
	reader := bufio.NewReader(os.Stdin)

	id, _ := getInput("Input Id of the public playlist you'd like to see: ", reader)

	sp := getById(id)
	p := createObject(sp)
	fmt.Println("The name of this playlist is: ", p.name, "and the description is: ", p.description)

	return p
}

func main() {

	getPlaylist()
}
