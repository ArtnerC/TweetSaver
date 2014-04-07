package tsapp

import (
	"fmt"
	"log"
	"time"
	"github.com/ArtnerC/TweetSaver/tweetsaver"

	"appengine"
	"appengine/datastore"
)

const tweet_kind = "Tweet"

type DataStore struct {
	c appengine.Context
}

func NewDataStore(context appengine.Context) *DataStore {
	return &DataStore{c: context}
}

func (ds *DataStore) Get(id int) *tweetsaver.Tweet {
	key := datastore.NewKey(ds.c, tweet_kind, "", id, tweetsaverKey(ds.c))
	t := new(tweetsaver.Tweet)

	err := datastore.Get(ds.c, key, t)
	if err != nil {
		return nil
	}

	t.Id = key.IntID()
	return t
}

func (ds *DataStore) GetAll() []*tweetsaver.Tweet {
	q := datastore.NewQuery(tweet_kind).Ancestor(tweetsaverKey(ds.c)).Order("-SaveTime")

	num, err := q.Count(ds.c)
	if err != nil {
		log.Printf("Datastore GetAll (count): %s", err.Error())
		return nil
	}

	tweets := make([]*tweetsaver.Tweet, 0, num)
	if _, err := q.GetAll(ds.c, &tweets); err != nil {
		log.Printf("Datastore GetAll (getall): %s", err.Error())
		return nil
	}

	return tweets
}

func (ds *DataStore) Find(author string) []*tweetsaver.Tweet {
	panic("Not Implemented: Find")
	return nil
}

func (ds *DataStore) Add(t *tweetsaver.Tweet) (int, error) {
	t.SaveTime = time.Now()

	key := datastore.NewIncompleteKey(ds.c, tweet_kind, tweetsaverKey(ds.c))
	t.Id = key.IntID()

	k, err := datastore.Put(ds.c, key, t)
	if err != nil {
		return 0, fmt.Errorf("Datastore Add: %s", err.Error())
	}

	return k.IntID(), nil
}

func (ds *DataStore) Update(t *tweetsaver.Tweet) error {
	panic("Not Implemented: Update")
	return nil
}

func (ds *DataStore) Delete(id int) {
	panic("Not Implemented: Delete")
}

func tweetsaverKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Tweets", "all_tweets", 0, nil)
}
