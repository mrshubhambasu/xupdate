package differentiator

import (
	"reflect"
	"testing"
)

func Test_differentiate_Differentiate(t *testing.T) {
	type args struct {
		oldbs []byte
		newbs []byte
	}
	tests := []struct {
		name    string
		d       differentiate
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := differentiate{}
			got, err := d.Differentiate(tt.args.oldbs, tt.args.newbs)
			if (err != nil) != tt.wantErr {
				t.Errorf("differentiate.Differentiate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("differentiate.Differentiate() = %v, want %v", got, tt.want)
			}
		})
	}
}
