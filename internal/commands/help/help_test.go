package help

import (
	"homework_4/internal/commands"
	"testing"
)

type mockCommand struct {
	Name        string
	Description string
	Required    map[string]struct{}
}

func (c *mockCommand) GetName() string {
	return c.Name
}

func (c *mockCommand) GetDescription() string {
	return c.Description
}

func (c *mockCommand) GetRequired() map[string]struct{} {
	return c.Required
}

func (c *mockCommand) DoAction(args map[string]string) error {
	// Implement the action for the mock command
	return nil
}
func TestHelpCommand_DoAction(t *testing.T) {
	availableCommands := make(map[string]commands.CommandInterface) // Create a map of available commands for testing
	availableCommands["command1"] = &mockCommand{Name: "command1", Description: "Description1", Required: map[string]struct{}{}}
	availableCommands["command2"] = &mockCommand{Name: "command2", Description: "Description2", Required: map[string]struct{}{"arg2": {}}}

	type fields struct {
		AvailableCommands map[string]commands.CommandInterface
		Name              string
		Description       string
		Required          map[string]struct{}
	}
	type args struct {
		args map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test with available commands",
			fields: fields{
				AvailableCommands: availableCommands,
				Name:              "help",
				Description:       "Display a list of available commands",
				Required:          map[string]struct{}{},
			},
			args: args{
				args: map[string]string{},
			},
			wantErr: false,
		},
		{
			name: "Test with missing args",
			fields: fields{
				AvailableCommands: availableCommands,
				Name:              "help",
				Description:       "Display a list of available commands",
				Required:          map[string]struct{}{},
			},
			args: args{
				args: map[string]string{"missingArg": "value"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HelpCommand{
				AvailableCommands: tt.fields.AvailableCommands,
				Name:              tt.fields.Name,
				Description:       tt.fields.Description,
				Required:          tt.fields.Required,
			}
			if err := c.DoAction(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DoAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
