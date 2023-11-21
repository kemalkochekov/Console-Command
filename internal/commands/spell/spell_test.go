package spell

import "testing"

func TestSpellCommand_DoAction(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Required    map[string]struct{}
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
			name: "ValidInput",
			fields: fields{
				Name:        "spell",
				Description: "Description for the spell command",
				Required:    map[string]struct{}{"word": {}},
			},
			args: args{
				args: map[string]string{"word": "hello"},
			},
			wantErr: false,
		},
		{
			name: "MissingRequiredArg",
			fields: fields{
				Name:        "spell",
				Description: "Description for the spell command",
				Required:    map[string]struct{}{"word": {}},
			},
			args: args{
				args: map[string]string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &SpellCommand{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Required:    tt.fields.Required,
			}
			if err := cmd.DoAction(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DoAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formatStringWithSpaces(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{
			name: "Single character",
			word: "A",
			want: "A ",
		},
		{
			name: "Word with symbols",
			word: "Hello-%@!World",
			want: "H e l l o - % @ ! W o r l d ",
		},
		{
			name: "Empty string",
			word: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStringWithSpaces(tt.word); got != tt.want {
				t.Errorf("formatStringWithSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
