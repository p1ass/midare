package twitter

import (
	"context"
	"time"
)

// FakeTwitterClient is a fake implementation of TwitterClient.
// It is used for ONLY testing.
type FakeTwitterClient struct {
}

var _ Client = &FakeTwitterClient{}

func (c *FakeTwitterClient) GetMe(ctx context.Context) (*User, error) {
	return &User{
		ID:         "1032935958964973568",
		Name:       "ぷらす",
		ScreenName: "p1ass",
		ImageURL:   "https://pbs.twimg.com/profile_images/1401046091227811842/AOffsP6w_normal.jpg",
	}, nil
}

func (c *FakeTwitterClient) GetTweets(ctx context.Context, userID string) ([]*Tweet, error) {
	// 存在しないユーザの場合は空の配列を返す (APIと同様の挙動)
	if userID != "1032935958964973568" {
		return []*Tweet{}, nil
	}
	return []*Tweet{
		{
			ID:      "1512031854010187780",
			Text:    "dummy text (2022-04-07 20:37:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T20:37:40+09:00"),
		},
		{
			ID:      "1512014281222803460",
			Text:    "dummy text (2022-04-07 19:27:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T19:27:51+09:00"),
		},
		{
			ID:      "1512003402171305984",
			Text:    "dummy text (2022-04-07 18:44:37 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T18:44:37+09:00"),
		},
		{
			ID:      "1511999498473844741",
			Text:    "dummy text (2022-04-07 18:29:06 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T18:29:06+09:00"),
		},
		{
			ID:      "1511979067486392324",
			Text:    "dummy text (2022-04-07 17:07:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T17:07:55+09:00"),
		},
		{
			ID:      "1511979018752786432",
			Text:    "dummy text (2022-04-07 17:07:43 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T17:07:43+09:00"),
		},
		{
			ID:      "1511975965827633153",
			Text:    "dummy text (2022-04-07 16:55:36 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T16:55:36+09:00"),
		},
		{
			ID:      "1511720764197777411",
			Text:    "dummy text (2022-04-07 00:01:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-07T00:01:31+09:00"),
		},
		{
			ID:      "1511691214776836102",
			Text:    "dummy text (2022-04-06 22:04:06 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T22:04:06+09:00"),
		},
		{
			ID:      "1511667395525746691",
			Text:    "dummy text (2022-04-06 20:29:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T20:29:27+09:00"),
		},
		{
			ID:      "1511621845090385925",
			Text:    "dummy text (2022-04-06 17:28:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T17:28:27+09:00"),
		},
		{
			ID:      "1511359356839038977",
			Text:    "dummy text (2022-04-06 00:05:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T00:05:25+09:00"),
		},
		{
			ID:      "1511359161984249860",
			Text:    "dummy text (2022-04-06 00:04:38 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T00:04:38+09:00"),
		},
		{
			ID:      "1511358348905570304",
			Text:    "dummy text (2022-04-06 00:01:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-06T00:01:24+09:00"),
		},
		{
			ID:      "1511354181906735105",
			Text:    "dummy text (2022-04-05 23:44:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T23:44:51+09:00"),
		},
		{
			ID:      "1511354078491983872",
			Text:    "dummy text (2022-04-05 23:44:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T23:44:26+09:00"),
		},
		{
			ID:      "1511307579040813060",
			Text:    "dummy text (2022-04-05 20:39:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T20:39:40+09:00"),
		},
		{
			ID:      "1511307535126773763",
			Text:    "dummy text (2022-04-05 20:39:29 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T20:39:29+09:00"),
		},
		{
			ID:      "1511295628953124873",
			Text:    "dummy text (2022-04-05 19:52:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T19:52:11+09:00"),
		},
		{
			ID:      "1511295128178393092",
			Text:    "dummy text (2022-04-05 19:50:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T19:50:11+09:00"),
		},
		{
			ID:      "1511254361401597952",
			Text:    "dummy text (2022-04-05 17:08:12 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T17:08:12+09:00"),
		},
		{
			ID:      "1511200622002913288",
			Text:    "dummy text (2022-04-05 13:34:39 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T13:34:39+09:00"),
		},
		{
			ID:      "1511199659913490432",
			Text:    "dummy text (2022-04-05 13:30:50 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T13:30:50+09:00"),
		},
		{
			ID:      "1510995926931881994",
			Text:    "dummy text (2022-04-05 00:01:16 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-05T00:01:16+09:00"),
		},
		{
			ID:      "1510975965144313864",
			Text:    "dummy text (2022-04-04 22:41:57 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:41:57+09:00"),
		},
		{
			ID:      "1510973594201108488",
			Text:    "dummy text (2022-04-04 22:32:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:32:32+09:00"),
		},
		{
			ID:      "1510973496381292545",
			Text:    "dummy text (2022-04-04 22:32:08 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:32:08+09:00"),
		},
		{
			ID:      "1510973363149500422",
			Text:    "dummy text (2022-04-04 22:31:36 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:31:36+09:00"),
		},
		{
			ID:      "1510973291326226440",
			Text:    "dummy text (2022-04-04 22:31:19 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:31:19+09:00"),
		},
		{
			ID:      "1510973139815387145",
			Text:    "dummy text (2022-04-04 22:30:43 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:30:43+09:00"),
		},
		{
			ID:      "1510973058634641416",
			Text:    "dummy text (2022-04-04 22:30:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:30:24+09:00"),
		},
		{
			ID:      "1510973004632961030",
			Text:    "dummy text (2022-04-04 22:30:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T22:30:11+09:00"),
		},
		{
			ID:      "1510958493721047048",
			Text:    "dummy text (2022-04-04 21:32:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T21:32:31+09:00"),
		},
		{
			ID:      "1510953107483947010",
			Text:    "dummy text (2022-04-04 21:11:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T21:11:07+09:00"),
		},
		{
			ID:      "1510948776319811584",
			Text:    "dummy text (2022-04-04 20:53:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T20:53:55+09:00"),
		},
		{
			ID:      "1510944624655089669",
			Text:    "dummy text (2022-04-04 20:37:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T20:37:25+09:00"),
		},
		{
			ID:      "1510927878141083657",
			Text:    "dummy text (2022-04-04 19:30:52 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T19:30:52+09:00"),
		},
		{
			ID:      "1510927723643867139",
			Text:    "dummy text (2022-04-04 19:30:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T19:30:15+09:00"),
		},
		{
			ID:      "1510855588758319106",
			Text:    "dummy text (2022-04-04 14:43:37 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T14:43:37+09:00"),
		},
		{
			ID:      "1510855268837761028",
			Text:    "dummy text (2022-04-04 14:42:21 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T14:42:21+09:00"),
		},
		{
			ID:      "1510852444787720195",
			Text:    "dummy text (2022-04-04 14:31:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T14:31:07+09:00"),
		},
		{
			ID:      "1510851545960632322",
			Text:    "dummy text (2022-04-04 14:27:33 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T14:27:33+09:00"),
		},
		{
			ID:      "1510633530111709185",
			Text:    "dummy text (2022-04-04 00:01:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-04T00:01:14+09:00"),
		},
		{
			ID:      "1510633137423847429",
			Text:    "dummy text (2022-04-03 23:59:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T23:59:40+09:00"),
		},
		{
			ID:      "1510629836595011586",
			Text:    "dummy text (2022-04-03 23:46:33 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T23:46:33+09:00"),
		},
		{
			ID:      "1510626655458324486",
			Text:    "dummy text (2022-04-03 23:33:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T23:33:55+09:00"),
		},
		{
			ID:      "1510509624725995522",
			Text:    "dummy text (2022-04-03 15:48:53 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T15:48:53+09:00"),
		},
		{
			ID:      "1510508716038451208",
			Text:    "dummy text (2022-04-03 15:45:16 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T15:45:16+09:00"),
		},
		{
			ID:      "1510271141935472645",
			Text:    "dummy text (2022-04-03 00:01:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-03T00:01:14+09:00"),
		},
		{
			ID:      "1510223299892355072",
			Text:    "dummy text (2022-04-02 20:51:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T20:51:07+09:00"),
		},
		{
			ID:      "1510197054165118981",
			Text:    "dummy text (2022-04-02 19:06:50 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T19:06:50+09:00"),
		},
		{
			ID:      "1510196699150811136",
			Text:    "dummy text (2022-04-02 19:05:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T19:05:25+09:00"),
		},
		{
			ID:      "1510195193219145729",
			Text:    "dummy text (2022-04-02 18:59:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T18:59:26+09:00"),
		},
		{
			ID:      "1510152468654878720",
			Text:    "dummy text (2022-04-02 16:09:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T16:09:40+09:00"),
		},
		{
			ID:      "1510110972995723274",
			Text:    "dummy text (2022-04-02 13:24:47 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T13:24:47+09:00"),
		},
		{
			ID:      "1510110618480951299",
			Text:    "dummy text (2022-04-02 13:23:22 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T13:23:22+09:00"),
		},
		{
			ID:      "1509931683566727169",
			Text:    "dummy text (2022-04-02 01:32:21 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T01:32:21+09:00"),
		},
		{
			ID:      "1509908766199795718",
			Text:    "dummy text (2022-04-02 00:01:17 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-02T00:01:17+09:00"),
		},
		{
			ID:      "1509900993776926735",
			Text:    "dummy text (2022-04-01 23:30:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T23:30:24+09:00"),
		},
		{
			ID:      "1509849315610861570",
			Text:    "dummy text (2022-04-01 20:05:03 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T20:05:03+09:00"),
		},
		{
			ID:      "1509839387735261186",
			Text:    "dummy text (2022-04-01 19:25:36 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T19:25:36+09:00"),
		},
		{
			ID:      "1509818091089412141",
			Text:    "dummy text (2022-04-01 18:00:58 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T18:00:58+09:00"),
		},
		{
			ID:      "1509775869018980357",
			Text:    "dummy text (2022-04-01 15:13:12 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T15:13:12+09:00"),
		},
		{
			ID:      "1509747333202087936",
			Text:    "dummy text (2022-04-01 13:19:48 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T13:19:48+09:00"),
		},
		{
			ID:      "1509702983092633602",
			Text:    "dummy text (2022-04-01 10:23:34 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T10:23:34+09:00"),
		},
		{
			ID:      "1509554460724719626",
			Text:    "dummy text (2022-04-01 00:33:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T00:33:24+09:00"),
		},
		{
			ID:      "1509553063698534404",
			Text:    "dummy text (2022-04-01 00:27:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T00:27:51+09:00"),
		},
		{
			ID:      "1509550757233631236",
			Text:    "dummy text (2022-04-01 00:18:41 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T00:18:41+09:00"),
		},
		{
			ID:      "1509547082822332427",
			Text:    "dummy text (2022-04-01 00:04:05 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T00:04:05+09:00"),
		},
		{
			ID:      "1509546403508465670",
			Text:    "dummy text (2022-04-01 00:01:23 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-04-01T00:01:23+09:00"),
		},
		{
			ID:      "1509542500692029447",
			Text:    "dummy text (2022-03-31 23:45:52 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T23:45:52+09:00"),
		},
		{
			ID:      "1509517710077882371",
			Text:    "dummy text (2022-03-31 22:07:22 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T22:07:22+09:00"),
		},
		{
			ID:      "1509516563401310212",
			Text:    "dummy text (2022-03-31 22:02:48 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T22:02:48+09:00"),
		},
		{
			ID:      "1509514020529008642",
			Text:    "dummy text (2022-03-31 21:52:42 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T21:52:42+09:00"),
		},
		{
			ID:      "1509504664651599882",
			Text:    "dummy text (2022-03-31 21:15:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T21:15:31+09:00"),
		},
		{
			ID:      "1509501120586330113",
			Text:    "dummy text (2022-03-31 21:01:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T21:01:26+09:00"),
		},
		{
			ID:      "1509499360354381835",
			Text:    "dummy text (2022-03-31 20:54:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T20:54:27+09:00"),
		},
		{
			ID:      "1509499069018013700",
			Text:    "dummy text (2022-03-31 20:53:17 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T20:53:17+09:00"),
		},
		{
			ID:      "1509498985186480136",
			Text:    "dummy text (2022-03-31 20:52:57 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T20:52:57+09:00"),
		},
		{
			ID:      "1509498936197025805",
			Text:    "dummy text (2022-03-31 20:52:46 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T20:52:46+09:00"),
		},
		{
			ID:      "1509498893033414673",
			Text:    "dummy text (2022-03-31 20:52:35 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T20:52:35+09:00"),
		},
		{
			ID:      "1509484226739339264",
			Text:    "dummy text (2022-03-31 19:54:19 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T19:54:19+09:00"),
		},
		{
			ID:      "1509482360622174212",
			Text:    "dummy text (2022-03-31 19:46:54 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T19:46:54+09:00"),
		},
		{
			ID:      "1509416042317053952",
			Text:    "dummy text (2022-03-31 15:23:22 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T15:23:22+09:00"),
		},
		{
			ID:      "1509409703969042432",
			Text:    "dummy text (2022-03-31 14:58:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T14:58:11+09:00"),
		},
		{
			ID:      "1509385713774899200",
			Text:    "dummy text (2022-03-31 13:22:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T13:22:51+09:00"),
		},
		{
			ID:      "1509383287936286721",
			Text:    "dummy text (2022-03-31 13:13:13 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T13:13:13+09:00"),
		},
		{
			ID:      "1509348361518092289",
			Text:    "dummy text (2022-03-31 10:54:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T10:54:26+09:00"),
		},
		{
			ID:      "1509196986201051136",
			Text:    "dummy text (2022-03-31 00:52:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:52:55+09:00"),
		},
		{
			ID:      "1509193718523633665",
			Text:    "dummy text (2022-03-31 00:39:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:39:56+09:00"),
		},
		{
			ID:      "1509193677549469704",
			Text:    "dummy text (2022-03-31 00:39:46 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:39:46+09:00"),
		},
		{
			ID:      "1509191912733155329",
			Text:    "dummy text (2022-03-31 00:32:46 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:32:46+09:00"),
		},
		{
			ID:      "1509188747791077385",
			Text:    "dummy text (2022-03-31 00:20:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:20:11+09:00"),
		},
		{
			ID:      "1509188109149569024",
			Text:    "dummy text (2022-03-31 00:17:39 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:17:39+09:00"),
		},
		{
			ID:      "1509187074620268552",
			Text:    "dummy text (2022-03-31 00:13:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:13:32+09:00"),
		},
		{
			ID:      "1509183980196794376",
			Text:    "dummy text (2022-03-31 00:01:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-31T00:01:14+09:00"),
		},
		{
			ID:      "1509153737574871043",
			Text:    "dummy text (2022-03-30 22:01:04 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T22:01:04+09:00"),
		},
		{
			ID:      "1509149528179736581",
			Text:    "dummy text (2022-03-30 21:44:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T21:44:20+09:00"),
		},
		{
			ID:      "1509147311615922191",
			Text:    "dummy text (2022-03-30 21:35:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T21:35:32+09:00"),
		},
		{
			ID:      "1509028881378209796",
			Text:    "dummy text (2022-03-30 13:44:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T13:44:56+09:00"),
		},
		{
			ID:      "1508824290405531648",
			Text:    "dummy text (2022-03-30 00:11:58 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T00:11:58+09:00"),
		},
		{
			ID:      "1508823762715963405",
			Text:    "dummy text (2022-03-30 00:09:52 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T00:09:52+09:00"),
		},
		{
			ID:      "1508822980633460741",
			Text:    "dummy text (2022-03-30 00:06:45 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T00:06:45+09:00"),
		},
		{
			ID:      "1508821631933820930",
			Text:    "dummy text (2022-03-30 00:01:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-30T00:01:24+09:00"),
		},
		{
			ID:      "1508789649862725635",
			Text:    "dummy text (2022-03-29 21:54:19 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T21:54:19+09:00"),
		},
		{
			ID:      "1508789558078742531",
			Text:    "dummy text (2022-03-29 21:53:57 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T21:53:57+09:00"),
		},
		{
			ID:      "1508785843271761925",
			Text:    "dummy text (2022-03-29 21:39:11 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T21:39:11+09:00"),
		},
		{
			ID:      "1508776699298992137",
			Text:    "dummy text (2022-03-29 21:02:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T21:02:51+09:00"),
		},
		{
			ID:      "1508771601470222336",
			Text:    "dummy text (2022-03-29 20:42:36 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T20:42:36+09:00"),
		},
		{
			ID:      "1508769477709950984",
			Text:    "dummy text (2022-03-29 20:34:09 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T20:34:09+09:00"),
		},
		{
			ID:      "1508768216084287492",
			Text:    "dummy text (2022-03-29 20:29:08 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T20:29:08+09:00"),
		},
		{
			ID:      "1508763761272057857",
			Text:    "dummy text (2022-03-29 20:11:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T20:11:26+09:00"),
		},
		{
			ID:      "1508763196504809473",
			Text:    "dummy text (2022-03-29 20:09:12 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T20:09:12+09:00"),
		},
		{
			ID:      "1508701612218322945",
			Text:    "dummy text (2022-03-29 16:04:29 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T16:04:29+09:00"),
		},
		{
			ID:      "1508459228083761153",
			Text:    "dummy text (2022-03-29 00:01:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-29T00:01:20+09:00"),
		},
		{
			ID:      "1508424445228363777",
			Text:    "dummy text (2022-03-28 21:43:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T21:43:07+09:00"),
		},
		{
			ID:      "1508416339769454593",
			Text:    "dummy text (2022-03-28 21:10:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T21:10:55+09:00"),
		},
		{
			ID:      "1508412648488013826",
			Text:    "dummy text (2022-03-28 20:56:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T20:56:15+09:00"),
		},
		{
			ID:      "1508411726387064834",
			Text:    "dummy text (2022-03-28 20:52:35 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T20:52:35+09:00"),
		},
		{
			ID:      "1508384632638939136",
			Text:    "dummy text (2022-03-28 19:04:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T19:04:55+09:00"),
		},
		{
			ID:      "1508313544135426050",
			Text:    "dummy text (2022-03-28 14:22:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T14:22:26+09:00"),
		},
		{
			ID:      "1508305243624587264",
			Text:    "dummy text (2022-03-28 13:49:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T13:49:27+09:00"),
		},
		{
			ID:      "1508302722231316483",
			Text:    "dummy text (2022-03-28 13:39:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T13:39:26+09:00"),
		},
		{
			ID:      "1508300200061808648",
			Text:    "dummy text (2022-03-28 13:29:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T13:29:25+09:00"),
		},
		{
			ID:      "1508291027211517954",
			Text:    "dummy text (2022-03-28 12:52:58 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T12:52:58+09:00"),
		},
		{
			ID:      "1508283866829959169",
			Text:    "dummy text (2022-03-28 12:24:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T12:24:31+09:00"),
		},
		{
			ID:      "1508105827068280833",
			Text:    "dummy text (2022-03-28 00:37:03 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T00:37:03+09:00"),
		},
		{
			ID:      "1508096854118088717",
			Text:    "dummy text (2022-03-28 00:01:23 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-28T00:01:23+09:00"),
		},
		{
			ID:      "1508094843096305669",
			Text:    "dummy text (2022-03-27 23:53:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T23:53:24+09:00"),
		},
		{
			ID:      "1508065914872557569",
			Text:    "dummy text (2022-03-27 21:58:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:58:27+09:00"),
		},
		{
			ID:      "1508064311218143236",
			Text:    "dummy text (2022-03-27 21:52:04 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:52:04+09:00"),
		},
		{
			ID:      "1508056271852740620",
			Text:    "dummy text (2022-03-27 21:20:08 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:20:08+09:00"),
		},
		{
			ID:      "1508055909150658561",
			Text:    "dummy text (2022-03-27 21:18:41 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:18:41+09:00"),
		},
		{
			ID:      "1508055066112327684",
			Text:    "dummy text (2022-03-27 21:15:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:15:20+09:00"),
		},
		{
			ID:      "1508054944741732352",
			Text:    "dummy text (2022-03-27 21:14:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T21:14:51+09:00"),
		},
		{
			ID:      "1508046683435057152",
			Text:    "dummy text (2022-03-27 20:42:02 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T20:42:02+09:00"),
		},
		{
			ID:      "1508046669421903876",
			Text:    "dummy text (2022-03-27 20:41:58 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T20:41:58+09:00"),
		},
		{
			ID:      "1507978802839752706",
			Text:    "dummy text (2022-03-27 16:12:18 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T16:12:18+09:00"),
		},
		{
			ID:      "1507975571774799880",
			Text:    "dummy text (2022-03-27 15:59:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:59:27+09:00"),
		},
		{
			ID:      "1507973595636826128",
			Text:    "dummy text (2022-03-27 15:51:36 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:51:36+09:00"),
		},
		{
			ID:      "1507971564511559688",
			Text:    "dummy text (2022-03-27 15:43:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:43:32+09:00"),
		},
		{
			ID:      "1507971288891277312",
			Text:    "dummy text (2022-03-27 15:42:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:42:26+09:00"),
		},
		{
			ID:      "1507970876431794177",
			Text:    "dummy text (2022-03-27 15:40:48 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:40:48+09:00"),
		},
		{
			ID:      "1507963191820972034",
			Text:    "dummy text (2022-03-27 15:10:16 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T15:10:16+09:00"),
		},
		{
			ID:      "1507949089123868676",
			Text:    "dummy text (2022-03-27 14:14:13 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T14:14:13+09:00"),
		},
		{
			ID:      "1507939654653607939",
			Text:    "dummy text (2022-03-27 13:36:44 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T13:36:44+09:00"),
		},
		{
			ID:      "1507746026224259073",
			Text:    "dummy text (2022-03-27 00:47:19 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T00:47:19+09:00"),
		},
		{
			ID:      "1507742932253933575",
			Text:    "dummy text (2022-03-27 00:35:02 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T00:35:02+09:00"),
		},
		{
			ID:      "1507742757959630849",
			Text:    "dummy text (2022-03-27 00:34:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T00:34:20+09:00"),
		},
		{
			ID:      "1507741455372423169",
			Text:    "dummy text (2022-03-27 00:29:10 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T00:29:10+09:00"),
		},
		{
			ID:      "1507734432110817280",
			Text:    "dummy text (2022-03-27 00:01:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-27T00:01:15+09:00"),
		},
		{
			ID:      "1507727604115439618",
			Text:    "dummy text (2022-03-26 23:34:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T23:34:07+09:00"),
		},
		{
			ID:      "1507709156588875779",
			Text:    "dummy text (2022-03-26 22:20:49 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T22:20:49+09:00"),
		},
		{
			ID:      "1507682026182541319",
			Text:    "dummy text (2022-03-26 20:33:01 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T20:33:01+09:00"),
		},
		{
			ID:      "1507596190892044289",
			Text:    "dummy text (2022-03-26 14:51:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T14:51:56+09:00"),
		},
		{
			ID:      "1507585663977541634",
			Text:    "dummy text (2022-03-26 14:10:06 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T14:10:06+09:00"),
		},
		{
			ID:      "1507547927765389312",
			Text:    "dummy text (2022-03-26 11:40:09 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T11:40:09+09:00"),
		},
		{
			ID:      "1507547460243111937",
			Text:    "dummy text (2022-03-26 11:38:18 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T11:38:18+09:00"),
		},
		{
			ID:      "1507381157134213128",
			Text:    "dummy text (2022-03-26 00:37:28 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T00:37:28+09:00"),
		},
		{
			ID:      "1507372056912044046",
			Text:    "dummy text (2022-03-26 00:01:18 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-26T00:01:18+09:00"),
		},
		{
			ID:      "1507361083123646476",
			Text:    "dummy text (2022-03-25 23:17:42 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T23:17:42+09:00"),
		},
		{
			ID:      "1507359725117394953",
			Text:    "dummy text (2022-03-25 23:12:18 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T23:12:18+09:00"),
		},
		{
			ID:      "1507356235485437958",
			Text:    "dummy text (2022-03-25 22:58:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T22:58:26+09:00"),
		},
		{
			ID:      "1507353685646782464",
			Text:    "dummy text (2022-03-25 22:48:18 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T22:48:18+09:00"),
		},
		{
			ID:      "1507328325253214208",
			Text:    "dummy text (2022-03-25 21:07:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T21:07:32+09:00"),
		},
		{
			ID:      "1507322502049308680",
			Text:    "dummy text (2022-03-25 20:44:23 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T20:44:23+09:00"),
		},
		{
			ID:      "1507281354794450991",
			Text:    "dummy text (2022-03-25 18:00:53 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T18:00:53+09:00"),
		},
		{
			ID:      "1507250936045268994",
			Text:    "dummy text (2022-03-25 16:00:01 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T16:00:01+09:00"),
		},
		{
			ID:      "1507250317255389193",
			Text:    "dummy text (2022-03-25 15:57:33 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T15:57:33+09:00"),
		},
		{
			ID:      "1507230934957162499",
			Text:    "dummy text (2022-03-25 14:40:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T14:40:32+09:00"),
		},
		{
			ID:      "1507210918626992138",
			Text:    "dummy text (2022-03-25 13:21:00 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T13:21:00+09:00"),
		},
		{
			ID:      "1507192926123700237",
			Text:    "dummy text (2022-03-25 12:09:30 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T12:09:30+09:00"),
		},
		{
			ID:      "1507027627068243971",
			Text:    "dummy text (2022-03-25 01:12:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T01:12:40+09:00"),
		},
		{
			ID:      "1507009649626652679",
			Text:    "dummy text (2022-03-25 00:01:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-25T00:01:14+09:00"),
		},
		{
			ID:      "1506998262552993798",
			Text:    "dummy text (2022-03-24 23:15:59 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T23:15:59+09:00"),
		},
		{
			ID:      "1506994353486852098",
			Text:    "dummy text (2022-03-24 23:00:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T23:00:27+09:00"),
		},
		{
			ID:      "1506983812932444171",
			Text:    "dummy text (2022-03-24 22:18:34 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T22:18:34+09:00"),
		},
		{
			ID:      "1506982249694384133",
			Text:    "dummy text (2022-03-24 22:12:21 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T22:12:21+09:00"),
		},
		{
			ID:      "1506981577469100035",
			Text:    "dummy text (2022-03-24 22:09:41 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T22:09:41+09:00"),
		},
		{
			ID:      "1506973033663459330",
			Text:    "dummy text (2022-03-24 21:35:44 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T21:35:44+09:00"),
		},
		{
			ID:      "1506961058778718209",
			Text:    "dummy text (2022-03-24 20:48:09 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T20:48:09+09:00"),
		},
		{
			ID:      "1506959313407213569",
			Text:    "dummy text (2022-03-24 20:41:12 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T20:41:12+09:00"),
		},
		{
			ID:      "1506955958148546569",
			Text:    "dummy text (2022-03-24 20:27:52 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T20:27:52+09:00"),
		},
		{
			ID:      "1506918532952788992",
			Text:    "dummy text (2022-03-24 17:59:10 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:59:10+09:00"),
		},
		{
			ID:      "1506917857749528579",
			Text:    "dummy text (2022-03-24 17:56:29 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:56:29+09:00"),
		},
		{
			ID:      "1506915289446825987",
			Text:    "dummy text (2022-03-24 17:46:16 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:46:16+09:00"),
		},
		{
			ID:      "1506914525203013633",
			Text:    "dummy text (2022-03-24 17:43:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:43:14+09:00"),
		},
		{
			ID:      "1506912377996201990",
			Text:    "dummy text (2022-03-24 17:34:42 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:34:42+09:00"),
		},
		{
			ID:      "1506912283624017922",
			Text:    "dummy text (2022-03-24 17:34:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:34:20+09:00"),
		},
		{
			ID:      "1506908820945387521",
			Text:    "dummy text (2022-03-24 17:20:34 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T17:20:34+09:00"),
		},
		{
			ID:      "1506889495978799106",
			Text:    "dummy text (2022-03-24 16:03:47 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T16:03:47+09:00"),
		},
		{
			ID:      "1506860320203362304",
			Text:    "dummy text (2022-03-24 14:07:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T14:07:51+09:00"),
		},
		{
			ID:      "1506850815826624512",
			Text:    "dummy text (2022-03-24 13:30:05 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T13:30:05+09:00"),
		},
		{
			ID:      "1506848428755927040",
			Text:    "dummy text (2022-03-24 13:20:35 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T13:20:35+09:00"),
		},
		{
			ID:      "1506844940621725699",
			Text:    "dummy text (2022-03-24 13:06:44 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T13:06:44+09:00"),
		},
		{
			ID:      "1506647267947724803",
			Text:    "dummy text (2022-03-24 00:01:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-24T00:01:15+09:00"),
		},
		{
			ID:      "1506641656971673603",
			Text:    "dummy text (2022-03-23 23:38:57 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T23:38:57+09:00"),
		},
		{
			ID:      "1506622742166409220",
			Text:    "dummy text (2022-03-23 22:23:48 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T22:23:48+09:00"),
		},
		{
			ID:      "1506608413718958093",
			Text:    "dummy text (2022-03-23 21:26:51 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T21:26:51+09:00"),
		},
		{
			ID:      "1506510066324676615",
			Text:    "dummy text (2022-03-23 14:56:04 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T14:56:04+09:00"),
		},
		{
			ID:      "1506489644279427073",
			Text:    "dummy text (2022-03-23 13:34:55 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T13:34:55+09:00"),
		},
		{
			ID:      "1506449152582234113",
			Text:    "dummy text (2022-03-23 10:54:01 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T10:54:01+09:00"),
		},
		{
			ID:      "1506442618024398849",
			Text:    "dummy text (2022-03-23 10:28:03 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T10:28:03+09:00"),
		},
		{
			ID:      "1506442391217389573",
			Text:    "dummy text (2022-03-23 10:27:09 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T10:27:09+09:00"),
		},
		{
			ID:      "1506441852400320513",
			Text:    "dummy text (2022-03-23 10:25:00 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T10:25:00+09:00"),
		},
		{
			ID:      "1506436402707595265",
			Text:    "dummy text (2022-03-23 10:03:21 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T10:03:21+09:00"),
		},
		{
			ID:      "1506302933495074816",
			Text:    "dummy text (2022-03-23 01:12:59 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T01:12:59+09:00"),
		},
		{
			ID:      "1506302889350041600",
			Text:    "dummy text (2022-03-23 01:12:49 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T01:12:49+09:00"),
		},
		{
			ID:      "1506302666615693316",
			Text:    "dummy text (2022-03-23 01:11:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T01:11:56+09:00"),
		},
		{
			ID:      "1506302285340897282",
			Text:    "dummy text (2022-03-23 01:10:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T01:10:25+09:00"),
		},
		{
			ID:      "1506299383641669634",
			Text:    "dummy text (2022-03-23 00:58:53 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:58:53+09:00"),
		},
		{
			ID:      "1506298266191003649",
			Text:    "dummy text (2022-03-23 00:54:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:54:26+09:00"),
		},
		{
			ID:      "1506298190429290499",
			Text:    "dummy text (2022-03-23 00:54:08 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:54:08+09:00"),
		},
		{
			ID:      "1506298139934093320",
			Text:    "dummy text (2022-03-23 00:53:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:53:56+09:00"),
		},
		{
			ID:      "1506297029332717568",
			Text:    "dummy text (2022-03-23 00:49:32 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:49:32+09:00"),
		},
		{
			ID:      "1506296920800894978",
			Text:    "dummy text (2022-03-23 00:49:06 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:49:06+09:00"),
		},
		{
			ID:      "1506288647993270279",
			Text:    "dummy text (2022-03-23 00:16:13 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:16:13+09:00"),
		},
		{
			ID:      "1506284902655926274",
			Text:    "dummy text (2022-03-23 00:01:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-23T00:01:20+09:00"),
		},
		{
			ID:      "1506252393276588036",
			Text:    "dummy text (2022-03-22 21:52:10 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T21:52:10+09:00"),
		},
		{
			ID:      "1506237820603875328",
			Text:    "dummy text (2022-03-22 20:54:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T20:54:15+09:00"),
		},
		{
			ID:      "1506236280450945024",
			Text:    "dummy text (2022-03-22 20:48:08 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T20:48:08+09:00"),
		},
		{
			ID:      "1506236036094967810",
			Text:    "dummy text (2022-03-22 20:47:10 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T20:47:10+09:00"),
		},
		{
			ID:      "1506235978050322434",
			Text:    "dummy text (2022-03-22 20:46:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T20:46:56+09:00"),
		},
		{
			ID:      "1506131552337494021",
			Text:    "dummy text (2022-03-22 13:51:59 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T13:51:59+09:00"),
		},
		{
			ID:      "1506121608641277953",
			Text:    "dummy text (2022-03-22 13:12:28 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T13:12:28+09:00"),
		},
		{
			ID:      "1506079342849568769",
			Text:    "dummy text (2022-03-22 10:24:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T10:24:31+09:00"),
		},
		{
			ID:      "1506077802176847874",
			Text:    "dummy text (2022-03-22 10:18:24 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T10:18:24+09:00"),
		},
		{
			ID:      "1505930585277370369",
			Text:    "dummy text (2022-03-22 00:33:25 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T00:33:25+09:00"),
		},
		{
			ID:      "1505929950586871816",
			Text:    "dummy text (2022-03-22 00:30:53 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T00:30:53+09:00"),
		},
		{
			ID:      "1505923879864512512",
			Text:    "dummy text (2022-03-22 00:06:46 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T00:06:46+09:00"),
		},
		{
			ID:      "1505922490719211524",
			Text:    "dummy text (2022-03-22 00:01:15 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-22T00:01:15+09:00"),
		},
		{
			ID:      "1505874962472583173",
			Text:    "dummy text (2022-03-21 20:52:23 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T20:52:23+09:00"),
		},
		{
			ID:      "1505829650286915588",
			Text:    "dummy text (2022-03-21 17:52:20 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T17:52:20+09:00"),
		},
		{
			ID:      "1505805894382075905",
			Text:    "dummy text (2022-03-21 16:17:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T16:17:56+09:00"),
		},
		{
			ID:      "1505805537056747528",
			Text:    "dummy text (2022-03-21 16:16:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T16:16:31+09:00"),
		},
		{
			ID:      "1505790474648387585",
			Text:    "dummy text (2022-03-21 15:16:40 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T15:16:40+09:00"),
		},
		{
			ID:      "1505786458531004418",
			Text:    "dummy text (2022-03-21 15:00:42 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T15:00:42+09:00"),
		},
		{
			ID:      "1505560100618092549",
			Text:    "dummy text (2022-03-21 00:01:14 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-21T00:01:14+09:00"),
		},
		{
			ID:      "1505541601552048137",
			Text:    "dummy text (2022-03-20 22:47:44 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T22:47:44+09:00"),
		},
		{
			ID:      "1505534342579179523",
			Text:    "dummy text (2022-03-20 22:18:53 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T22:18:53+09:00"),
		},
		{
			ID:      "1505413688324345856",
			Text:    "dummy text (2022-03-20 14:19:27 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T14:19:27+09:00"),
		},
		{
			ID:      "1505238045146349569",
			Text:    "dummy text (2022-03-20 02:41:30 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:41:30+09:00"),
		},
		{
			ID:      "1505235896647680005",
			Text:    "dummy text (2022-03-20 02:32:58 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:32:58+09:00"),
		},
		{
			ID:      "1505235186472341504",
			Text:    "dummy text (2022-03-20 02:30:09 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:30:09+09:00"),
		},
		{
			ID:      "1505229973916770304",
			Text:    "dummy text (2022-03-20 02:09:26 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:09:26+09:00"),
		},
		{
			ID:      "1505229832451289089",
			Text:    "dummy text (2022-03-20 02:08:52 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:08:52+09:00"),
		},
		{
			ID:      "1505229509867409410",
			Text:    "dummy text (2022-03-20 02:07:35 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T02:07:35+09:00"),
		},
		{
			ID:      "1505197745321254915",
			Text:    "dummy text (2022-03-20 00:01:22 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-20T00:01:22+09:00"),
		},
		{
			ID:      "1505162759515770885",
			Text:    "dummy text (2022-03-19 21:42:21 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-19T21:42:21+09:00"),
		},
		{
			ID:      "1505105335538790400",
			Text:    "dummy text (2022-03-19 17:54:10 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-19T17:54:10+09:00"),
		},
		{
			ID:      "1504852004853919746",
			Text:    "dummy text (2022-03-19 01:07:31 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-19T01:07:31+09:00"),
		},
		{
			ID:      "1504835335234244610",
			Text:    "dummy text (2022-03-19 00:01:17 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-19T00:01:17+09:00"),
		},
		{
			ID:      "1504830183848955908",
			Text:    "dummy text (2022-03-18 23:40:48 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T23:40:48+09:00"),
		},
		{
			ID:      "1504829493206151170",
			Text:    "dummy text (2022-03-18 23:38:04 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T23:38:04+09:00"),
		},
		{
			ID:      "1504821294004842497",
			Text:    "dummy text (2022-03-18 23:05:29 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T23:05:29+09:00"),
		},
		{
			ID:      "1504819566194884609",
			Text:    "dummy text (2022-03-18 22:58:37 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T22:58:37+09:00"),
		},
		{
			ID:      "1504761852324048898",
			Text:    "dummy text (2022-03-18 19:09:17 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T19:09:17+09:00"),
		},
		{
			ID:      "1504744532734365709",
			Text:    "dummy text (2022-03-18 18:00:28 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T18:00:28+09:00"),
		},
		{
			ID:      "1504705584016556032",
			Text:    "dummy text (2022-03-18 15:25:41 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T15:25:41+09:00"),
		},
		{
			ID:      "1504681552378138626",
			Text:    "dummy text (2022-03-18 13:50:12 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T13:50:12+09:00"),
		},
		{
			ID:      "1504673731343691779",
			Text:    "dummy text (2022-03-18 13:19:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T13:19:07+09:00"),
		},
		{
			ID:      "1504635431761375239",
			Text:    "dummy text (2022-03-18 10:46:56 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T10:46:56+09:00"),
		},
		{
			ID:      "1504632206463242240",
			Text:    "dummy text (2022-03-18 10:34:07 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T10:34:07+09:00"),
		},
		{
			ID:      "1504472975633903616",
			Text:    "dummy text (2022-03-18 00:01:23 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-18T00:01:23+09:00"),
		},
		{
			ID:      "1504468790918397953",
			Text:    "dummy text (2022-03-17 23:44:46 +0900 Asia/Tokyo)",
			Created: mustToTime("2022-03-17T23:44:46+09:00"),
		},
	}, nil
}

func mustToTime(formatted string) time.Time {
	t, err := time.Parse(time.RFC3339, formatted)
	if err != nil {
		panic(err)
	}
	return t
}
