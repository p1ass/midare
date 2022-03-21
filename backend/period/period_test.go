package period

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/p1ass/midare/twitterv1"
)

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func Test_CalcAwakePeriods(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ts   []*twitterv1.Tweet
		want []*Period
	}{
		{
			name: "ツイートが一つも存在しない場合はperiodは空",
			ts:   nil,
			want: []*Period{},
		},
		{
			name: "1ツイートしか存在しない場合は起きている時間がないのでperiodは空",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{},
		},
		{
			name: "ツイートが2つ存在し、3.5時間以内のツイートであればperiodが1つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst),
					},
				},
			},
		},
		{
			name: "ツイートが2つ存在し、間隔が3.5時間より大きいツイートであればperiodが空",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 3, 31, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{},
		},
		{
			name: "ツイートが3つ存在し、全ての間隔が3.5時間以内のツイートであればperiodが1つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 7, 0, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 7, 0, 0, 0, jst),
					},
				},
			},
		},
		{
			name: "ツイートが3つ存在し、全ての間隔が3.5時間より大きいのツイートであればperiodが0つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 7, 32, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 3, 31, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{},
		},
		{
			name: "ツイートが3つ存在し、最初の2つの間隔が3.5時間以内のツイートであればperiodが1つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 7, 1, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst),
					},
				},
			},
		},
		{
			name: "ツイートが3つ存在し、最後の2つの間隔が3.5時間以内のツイートであればperiodが1つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 7, 1, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 3, 31, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 3, 31, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 7, 1, 0, 0, jst),
					},
				},
			},
		},
		{
			name: "ツイートが4つ存在し、最初の2つと最後の2つがそれぞれ間隔が3.5時間以内のツイートであればperiodが2つ",
			ts: []*twitterv1.Tweet{
				{Created: time.Date(2020, 1, 1, 10, 0, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 7, 1, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst)},
				{Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst)},
			},
			want: []*Period{
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 7, 1, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 10, 0, 0, 0, jst),
					},
				},
				{
					OkiTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 0, 0, 0, 0, jst),
					},
					NeTime: &twitterv1.Tweet{
						Created: time.Date(2020, 1, 1, 3, 30, 0, 0, jst),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcAwakePeriods(tt.ts); !cmp.Equal(got, tt.want) {
				t.Errorf("calcAwakePeriods() = diff=%v", cmp.Diff(tt.want, got))
			}
		})
	}
}
