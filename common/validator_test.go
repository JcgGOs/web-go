package common

import (
	"reflect"
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		field string
		value string
		rules []string
	}
	tests := []struct {
		name string
		args args
		want []Err
	}{
		{"email", args{"email", "email@email.com", []string{"email"}}, []Err{}},
		{"error email",
			args{"error_email", "email", []string{"email"}}, []Err{
				Err{Field: "error_email", Message: "invalid-email"},
			},
		},
		{"required",
			args{"name", "", []string{"required"}}, []Err{
				Err{Field: "name", Message: "field-required"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.args.field, tt.args.value, tt.args.rules...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
