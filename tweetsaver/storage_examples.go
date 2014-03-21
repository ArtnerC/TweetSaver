package tweetsaver

import (
	"time"
)

var ExampleTweets = []*Tweet{
	&Tweet{
		Id:        0,
		Text:      "I don't like to wake up!",
		Author:    "NormalTweetGuy",
		Link:      "https://twitter.com/NormalTweetGuy/status/436488438657712130",
		Timestamp: time.Now().AddDate(0, 0, -15),
		SaveTime:  time.Now(),
	},
	&Tweet{
		Id:        1,
		Text:      "Run the marathon, not the sprint.",
		Author:    "Frank_Underwood",
		Link:      "https://twitter.com/Frank_Underwood/status/436479390500016128",
		Timestamp: time.Now().AddDate(0, -1, 0),
		SaveTime:  time.Now(),
	},
	&Tweet{
		Id:        2,
		Text:      "the Earth is 30% land and .000001% Vitamin Water",
		Author:    "lawblob",
		Link:      "https://twitter.com/lawblob/status/436308475694837760",
		Timestamp: time.Now().AddDate(0, -1, -15),
		SaveTime:  time.Now(),
	},
	&Tweet{
		Id:        3,
		Text:      "i got 1100011 problems but knowimg binary aint 1",
		Author:    "jonnysun",
		Link:      "https://twitter.com/jonnysun/status/436222727851356160",
		Timestamp: time.Now().AddDate(0, -2, 0),
		SaveTime:  time.Now(),
	},
	&Tweet{
		Id:        4,
		Text:      "I'm the president of the French Chapter of the Beyoncé fan club. We're called C'est My Name",
		Author:    "mattytalks",
		Link:      "https://twitter.com/mattytalks/status/436351075927347201",
		Timestamp: time.Now().AddDate(0, -2, -15),
		SaveTime:  time.Now(),
	},
	&Tweet{
		Id:        5,
		Text:      "Hi I'm the guy you were avoiding eye contact with. I googled your name, would you like to go out wi..yes that was me who added u on facebook",
		Author:    "briangaar",
		Link:      "https://twitter.com/briangaar/status/442044703190421504",
		Timestamp: time.Now().AddDate(0, -3, 0),
		SaveTime:  time.Now(),
	},
}
