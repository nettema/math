package math

import (
	"reflect"
	"testing"

	"github.com/GDG-Cloud-Innopolis/Go-begginners/l3/test"
)

func TestCalc(t *testing.T) {
	type args struct {
		operation test.CalcOperation
	}
	tests := []struct {
		name string
		args args
		want func(a, b float64) float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calc(tt.args.operation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
