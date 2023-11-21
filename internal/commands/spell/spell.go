package spell

import (
	"fmt"
	"os"
)

type SpellCommand struct {
	Name        string
	Description string
	Required    map[string]struct{}
}

func NewSpell() *SpellCommand {
	return &SpellCommand{
		Name:        "spell",
		Description: "spell --word [string] | Accepts a word as input and outputs all the letters of that word separated by spaces in the console based on the results of the operation.\n",
		Required:    map[string]struct{}{"word": {}},
	}
}

func (cmd *SpellCommand) GetName() string {
	return cmd.Name
}

func (cmd *SpellCommand) GetDescription() string {
	return cmd.Description
}

func (cmd *SpellCommand) GetRequired() map[string]struct{} {
	return cmd.Required
}

func (cmd *SpellCommand) DoAction(args map[string]string) error {
	word, ok := args["word"]
	if !ok {
		return fmt.Errorf("missing 'word' argument")
	}
	formattedWord := formatStringWithSpaces(word)
	if _, err := fmt.Fprintln(os.Stderr, formattedWord); err != nil {
		return err
	}
	return nil
}
func formatStringWithSpaces(word string) string {
	formattedWord := ""
	for _, value := range word {
		formattedWord += string(value) + " "
	}
	return formattedWord
}
