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
			name: "ぼくへ 生活習慣乱れてませんか？",
			text: "ぼくへ 生活習慣乱れてませんか？",
			want: true,
		},
		{
			name: "p1ass さんの 2020/05/23 の contribution 数: 22\n #contributter_report",
			text: "p1ass さんの 2020/05/23 の contribution 数: 22\n #contributter_report",
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
