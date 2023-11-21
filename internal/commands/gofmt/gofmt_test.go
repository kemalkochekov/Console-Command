package gofmt

import (
	"testing"
)

func TestGofmtCommand_DoAction(t *testing.T) {

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
			name: "Test with no arguments",
			fields: fields{
				Name:        "gofmt",
				Description: "Format Go source code files",
				Required:    map[string]struct{}{},
			},
			args: args{
				args: map[string]string{},
			},
			wantErr: true,
		},
		{
			name: "Test with non-file argument",
			fields: fields{
				Name:        "gofmt",
				Description: "Format Go source code files",
				Required:    map[string]struct{}{"file": {}},
			},
			args: args{
				args: map[string]string{"somefile": "example.txt"}, // Using a non-Go source code file
			},
			wantErr: true, // There should be an error when a non-file argument is provided
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GofmtCommand{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Required:    tt.fields.Required,
			}
			if err := c.DoAction(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DoAction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_formatLine(t *testing.T) {
	tests := []struct {
		name           string
		line           string
		isNewParagraph bool
		want           string
	}{
		{
			name:           "Regular sentence in a new paragraph",
			line:           "This is a sentence.",
			isNewParagraph: true,
			want:           "\tThis is a sentence.",
		},
		{
			name:           "Regular sentence in the same paragraph",
			line:           "This is another sentence.",
			isNewParagraph: false,
			want:           "This is another sentence.",
		},
		{
			name:           "Uppercase start in a new paragraph",
			line:           "Starts with a capital letter.",
			isNewParagraph: true,
			want:           "\tStarts with a capital letter.",
		},
		{
			name:           "Empty line in a new paragraph",
			line:           "",
			isNewParagraph: true,
			want:           "",
		},
		{
			name:           "Empty line in the same paragraph",
			line:           "",
			isNewParagraph: false,
			want:           "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatLine(tt.line, tt.isNewParagraph); got != tt.want {
				t.Errorf("formatLine() = %v, want :%v", got, tt.want)
			}
		})
	}
}
