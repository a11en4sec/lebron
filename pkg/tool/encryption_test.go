package tool

import "testing"

func TestMd5ByString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "realm",
			args:	args{
				str: "reaml",
			},
			want:    "b94b7ef7f17d2394d6fbdf458dadc7b0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Md5ByString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5ByString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5ByString() = %v, want %v", got, tt.want)
			}
		})
	}
}
