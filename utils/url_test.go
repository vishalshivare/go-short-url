package utils

import "testing"

func TestGetDomain(t *testing.T) {
	type args struct {
		rawUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "ShouldReturnYoutubeDomain",
			args:    args{rawUrl: "https://www.youtube.com/watch?v=k8AObcX8azM"},
			want:    "www.youtube.com",
			wantErr: false,
		},
		{
			name:    "ShouldReturnExampleDomain",
			args:    args{rawUrl: "https://example.com/watch?v=k8AObcX8azM"},
			want:    "example.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDomain(tt.args.rawUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
