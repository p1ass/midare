package web

import (
	"testing"
)

func TestHandler_containExcludeWord(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want bool
	}{
		{
			name: "みんなへ 生活習慣乱れてませんか？",
			text: "みんなへ 生活習慣乱れてませんか？",
			want: true,
		},
		{
			name: "p1ass さんの 2020/05/23 の contribution 数: 22\n #contributter_report",
			text: "p1ass さんの 2020/05/23 の contribution 数: 22\n #contributter_report",
			want: true,
		},
		{
			name: "@uzimaru0000 05-24のポスト数：24 (うちRT：0)",
			text: "@uzimaru0000 05-24のポスト数：24 (うちRT：0)",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{}
			if got := h.containExcludeWord(tt.text); got != tt.want {
				t.Errorf("containExcludeWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
