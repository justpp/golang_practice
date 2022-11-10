package util

import "testing"

func TestIsExists(t *testing.T) {
	type args struct {
		dirName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"../jd/cookies",
			args{"../jd/cookies"},
			true,
			false,
		},
		{
			"./writeFile.go",
			args{"./writeFile.go"},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsExists(tt.args.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}
