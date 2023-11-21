package gofmt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type GofmtCommand struct {
	Name        string
	Description string
	Required    map[string]struct{}
}

// NewGofmtCommand creates a new instance of the GofmtCommand.
func NewGofmtCommand() *GofmtCommand {
	return &GofmtCommand{
		Name:        "gofmt",
		Description: "gofmt --file [string] | Accepts an *.txt file as input, and before each paragraph in the output, inserts a tab and places a period at the end of sentences.\n",
		Required:    map[string]struct{}{"file": {}},
	}
}

// GetName returns the name of the command.
func (c *GofmtCommand) GetName() string {
	return c.Name
}

// GetDescription returns a description of the command.
func (c *GofmtCommand) GetDescription() string {
	return c.Description
}

// GetRequired returns the list of required arguments for the "gofmt" command.
func (c *GofmtCommand) GetRequired() map[string]struct{} {
	return c.Required
}

// DoAction executes the "gofmt" command's action.
func (c *GofmtCommand) DoAction(args map[string]string) error {
	name, ok := args["file"]
	if !ok {
		return fmt.Errorf("missing 'file' argument")
	}
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("the file '%s' does not exist", name)
		}
		return err
	}
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open the file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var formattedText strings.Builder
	isNewParagraph := true
	for scanner.Scan() {
		line := scanner.Text()
		// Format the line
		formattedLine := formatLine(line, isNewParagraph)
		// Append the formatted line to the result
		formattedText.WriteString(formattedLine)
		isNewParagraph = len(line) == 0
		fmt.Println(formattedText.String())
		formattedText.Reset()
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading the file: %v", err)
	}

	fmt.Println("file successfully formatted")
	return nil
}
func formatLine(line string, isNewParagraph bool) string {
	isNewSentence := false
	formattedLine := line

	for i, char := range line {
		if i > 0 && unicode.IsUpper(char) {
			isNewSentence = true
			break
		}
	}
	if isNewParagraph && len(line) != 0 {
		// Add a tab for the first line of a paragraph
		formattedLine = "\t" + line
	}

	if isNewSentence {
		// Close the previous sentence with a period
		formattedLine = formattedLine + "."
	}

	return formattedLine
}
