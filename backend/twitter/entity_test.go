package twitter

import "testing"

func TestTweet_ContainExcludeWord(t *testing.T) {
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
			tw := &Tweet{
				Text: tt.text,
			}
			if got := tw.ContainExcludedWord(); got != tt.want {
				t.Errorf("ContainExcludedWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
