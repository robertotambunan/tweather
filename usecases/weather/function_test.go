package weather

import "testing"

func Test_composeMessage(t *testing.T) {
	type args struct {
		weather Item
	}
	tests := []struct {
		name      string
		args      args
		wantTweet string
	}{
		{
			name: "test 1",
			args: args{
				weather: Item{
					UpdatedAt:     "Mon 10:41",
					ProjectedTime: "Mon 12:41",
					Weathers: []ItemInfo{
						{
							Place:      "Jakarta",
							Weather:    "Rainy",
							Temprature: "24 C",
						},
						{
							Place:      "Medan",
							Weather:    "Cloudy",
							Temprature: "29 C",
						},
						{
							Place:      "Tangerang",
							Weather:    "Rainy",
							Temprature: "25 C",
						},
					}},
			},
			wantTweet: "Weather at Mon 12:41: Jakarta - Rainy (24 C), Medan - Cloudy (29 C), Tangerang - Rainy (25 C). This tweet updated at Mon 10:41 by Tweather-Bot.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTweet := composeMessage(tt.args.weather); gotTweet != tt.wantTweet {
				t.Errorf("composeMessage() = %v, want %v", gotTweet, tt.wantTweet)
			}
		})
	}
}
