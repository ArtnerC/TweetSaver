package tweetsaver

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Tweet struct {
	Id        int
	Text      string
	Author    string
	Link      string
	Timestamp time.Time
	SaveTime  time.Time
}

func NewTweet(r io.Reader) (*Tweet, error) {
	t := &Tweet{SaveTime: time.Now()}
	err := json.NewDecoder(r).Decode(t)
	return t, err
}

func (t *Tweet) String() string {
	return fmt.Sprintf(`"%s" -%s on [%s]`, t.Text, t.Author, t.Timestamp.Format(time.Stamp))
}
