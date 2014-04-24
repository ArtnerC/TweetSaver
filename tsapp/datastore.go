package tsapp

import (
	"fmt"
	"log"
	"time"
	"github.com/ArtnerC/TweetSaver/tweetsaver"

	"appengine"
	"appengine/datastore"
)

const tweet_kind = "SavedTweet"

type DataStore struct {
	c appengine.Context
}

func NewDataStore(context appengine.Context) *DataStore {
	return &DataStore{c: context}
}

func (ds *DataStore) Get(id int) *tweetsaver.Tweet {
	key := datastore.NewKey(ds.c, tweet_kind, "", int64(id), tweetsaverKey(ds.c))
	t := new(tweetsaver.Tweet)

	err := datastore.Get(ds.c, key, t)
	if err != nil {
		log.Printf("Get error: %s", err.Error())
		return nil
	}

	t.Id = id
	return t
}

func (ds *DataStore) GetAt(pos, limit int) []*tweetsaver.Tweet {
	q := datastore.NewQuery(tweet_kind).Ancestor(tweetsaverKey(ds.c)).Order("-SaveTime").Offset(pos).Limit(limit)

	num, err := q.Count(ds.c)
	if err != nil {
		log.Printf("Datastore GetAt (count): %s", err.Error())
		return nil
	}

	tweets := make([]*tweetsaver.Tweet, 0, num)
	keys, err := q.GetAll(ds.c, &tweets)
	if err != nil {
		log.Printf("Datastore GetAt (getall): %s", err.Error())
		return nil
	}

	for i := range tweets {
		tweets[i].Id = int(keys[i].IntID())
	}

	return tweets
}

func (ds *DataStore) GetAll() []*tweetsaver.Tweet {
	q := datastore.NewQuery(tweet_kind).Ancestor(tweetsaverKey(ds.c)).Order("-SaveTime")

	num, err := q.Count(ds.c)
	if err != nil {
		log.Printf("Datastore GetAll (count): %s", err.Error())
		return nil
	}

	tweets := make([]*tweetsaver.Tweet, 0, num)
	keys, err := q.GetAll(ds.c, &tweets)
	if err != nil {
		log.Printf("Datastore GetAll (getall): %s", err.Error())
		return nil
	}

	for i := range tweets {
		tweets[i].Id = int(keys[i].IntID())
	}

	return tweets
}

func (ds *DataStore) Find(author string) []*tweetsaver.Tweet {
	panic("Not Implemented: Find")
	return nil
}

func (ds *DataStore) Add(t *tweetsaver.Tweet) (int, error) {
	t.SaveTime = time.Now()

	// id, _, err := datastore.AllocateIDs(ds.c, tweet_kind, tweetsaverKey(ds.c), 1)
	// if err != nil {
	// 	return 0, fmt.Errorf("Datastore Add: %s", err.Error())
	// }
	// t.Id = int(id)

	// key := datastore.NewKey(ds.c, tweet_kind, "", id, tweetsaverKey(ds.c))
	key := datastore.NewIncompleteKey(ds.c, tweet_kind, tweetsaverKey(ds.c))

	k, err := datastore.Put(ds.c, key, t)
	if err != nil {
		return 0, fmt.Errorf("Datastore Add: %s", err.Error())
	}

	return int(k.IntID()), nil
}

func (ds *DataStore) Update(t *tweetsaver.Tweet) error {
	t.SaveTime = time.Now()
	key := datastore.NewKey(ds.c, tweet_kind, "", int64(t.Id), tweetsaverKey(ds.c))
	if _, err := datastore.Put(ds.c, key, t); err != nil {
		return fmt.Errorf("Datastore Update: %s", err.Error())
	}
	return nil
}

func (ds *DataStore) Delete(id int) {
	key := datastore.NewKey(ds.c, tweet_kind, "", int64(id), tweetsaverKey(ds.c))
	if err := datastore.Delete(ds.c, key); err != nil {
		log.Printf("Delete error: %s", err.Error())
	}
}

func tweetsaverKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Tweets", "all_tweets", 0, nil)
}
