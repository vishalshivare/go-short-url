package configs

import "testing"

func TestGetBaseURL(t *testing.T) {
	DefaultConfigFolder = "../defaults"
	ReadConfig()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "ShouldProvideCorrectBaseURL",
			want: "http://localhost:8080/v1/urlshorter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBaseURL(); got != tt.want {
				t.Errorf("GetBaseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
