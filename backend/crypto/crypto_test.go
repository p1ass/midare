package crypto

import "testing"

func TestSecureRandomBase64Encoded(t *testing.T) {
	type args struct {
		entropyByte int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "When 32byte entropy, should generate 43 length string (PKCE spec)",
			args: args{
				entropyByte: 32,
			},
			wantLen: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecureRandomBase64Encoded(tt.args.entropyByte); len(got) != tt.wantLen {
				t.Errorf("SecureRandomBase64Encoded() lenght = %v, want %v", got, tt.wantLen)
			}
		})
	}
}
