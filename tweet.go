package tweetsaver

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type tweet struct {
	Id        int
	Text      string
	Author    string
	Link      string
	Timestamp time.Time
	SaveTime  time.Time
}

func NewTweet(r io.Reader) (*tweet, error) {
	t := &tweet{SaveTime: time.Now()}
	err := json.NewDecoder(r).Decode(t)
	return t, err
}

func (t *tweet) String() string {
	return fmt.Sprintf(`"%s" -%s on [%s]`, t.Text, t.Author, t.Timestamp.Format(time.Stamp))
}
