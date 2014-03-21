package tweetsaver

import (
	"strings"
	"testing"
	"time"
)

var testTweetJSON string = `{"Text":"This is a funny tweet","Author":"Spry"}`

func TestNewTweet(t *testing.T) {
	tweet, err := NewTweet(strings.NewReader(testTweetJSON))
	if err != nil {
		t.Errorf("New Tweet failed: %v Error: %s", tweet, err.Error())
	}
	if tweet.Text != "This is a funny tweet" {
		t.Error("Text mismatch:", tweet.Text)
	}
	if tweet.Author != "Spry" {
		t.Error("Author mismatch:", tweet.Author)
	}
	if tweet.SaveTime.IsZero() {
		t.Error("Save time uninitialized")
	}
	if !tweet.Timestamp.IsZero() {
		t.Error("Timestamp non-zero")
	}
}

func TestString(t *testing.T) {
	a := &Tweet{Text: "This is a tweet", Author: "Person", Timestamp: time.Now()}
	s := `"` + a.Text + `" -` + a.Author + ` on [` + a.Timestamp.Format(time.Stamp) + `]`
	if as := a.String(); as != s {
		t.Error("Improper string created:", as)
	}
}
