package integrator

import "testing"

func Test_integrate_Integrate(t *testing.T) {
	type args struct {
		oldfile   string
		newfile   string
		patchfile string
	}
	tests := []struct {
		name    string
		i       integrate
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := integrate{}
			if err := i.Integrate(tt.args.oldfile, tt.args.newfile, tt.args.patchfile); (err != nil) != tt.wantErr {
				t.Errorf("integrate.Integrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
