package simplestore

import (
	ts "github.com/ArtnerC/TweetSaver/tweetsaver"
)

// CacheStorage implements the Persistence interface as an in-memory
// storage layer. Use NewStorageCache and NewMemoryStorage to initialize.
// This is an alternative/placeholder for MEMCACHED.
type CacheStorage struct {
	tweets    map[int]*ts.Tweet
	finds     map[string][]*ts.Tweet
	store     ts.Persistence
	allLoaded bool
}

// NewStorageCache creates a CacheStorage object that applies
// a caching layer to the provided p which implements the
// Persistence interface.
func NewStorageCache(p ts.Persistence) ts.Persistence {
	return &CacheStorage{
		tweets:    make(map[int]*ts.Tweet),
		finds:     make(map[string][]*ts.Tweet),
		store:     p,
		allLoaded: false,
	}
}

// NewMemoryStorage creates a CacheStorage object that can
// be used to store data in volatile RAM without the need
// for an underlying Persistence object.
func NewMemoryStorage() ts.Persistence {
	return NewStorageCache(nil)
}

// Get retrieves an item by ID
func (cs *CacheStorage) Get(id int) *ts.Tweet {
	if v, ok := cs.tweets[id]; ok == true {
		return v
	}
	if cs.store != nil {
		if v := cs.store.Get(id); v != nil {
			cs.tweets[id] = v
			return v
		}
	}
	return nil
}

func (cs *CacheStorage) GetAt(pos, limit int) ([]*ts.Tweet, int) {
	if cs.store == nil {
		panic("Cannot GetAt, implement GetAt")
	}

	return cs.store.GetAt(pos, limit)
}

// GetAll returns all stored elements. It also caches everything for future
// quick access.
func (cs *CacheStorage) GetAll() []*ts.Tweet {
	if cs.store != nil && cs.allLoaded == false {
		all := cs.store.GetAll()
		for _, t := range all {
			if _, ok := cs.tweets[t.Id]; ok == false {
				cs.tweets[t.Id] = t
			}
		}
		cs.allLoaded = true
		return all
	} else {
		tweets := make([]*ts.Tweet, 0, len(cs.tweets))
		for _, v := range cs.tweets {
			tweets = append(tweets, v)
		}
		if len(tweets) > 0 {
			return tweets
		}
	}

	return nil
}

// Find does a search for all items by author. Individual finds are cached.
func (cs *CacheStorage) Find(author string) []*ts.Tweet {
	if v, ok := cs.finds[author]; ok == true {
		return v
	}
	if cs.store != nil {
		if v := cs.store.Find(author); v != nil {
			cs.finds[author] = v
			return v
		} else {
			return nil
		}
	}

	found := make([]*ts.Tweet, 0, 10)
	for _, v := range cs.tweets {
		if v.Author == author {
			found = append(found, v)
		}
	}
	if len(found) > 0 {
		cs.finds[author] = found
		return found
	}
	return nil
}

// Add will save a new item to underlying storage, cache, and cached finds.
func (cs *CacheStorage) Add(t *ts.Tweet) (int, error) {
	id := len(cs.tweets)
	var err error
	if cs.store != nil {
		if id, err = cs.store.Add(t); err != nil {
			return 0, err
		}
	}

	t.Id = id
	nt := *t
	cs.tweets[id] = &nt

	if a, ok := cs.finds[nt.Author]; ok == true {
		cs.finds[nt.Author] = append(a, &nt)
	}
	return id, nil
}

// Update modifies an existing item in underlying storage, cache and finds.
func (cs *CacheStorage) Update(t *ts.Tweet) error {
	if cs.store != nil {
		if err := cs.store.Update(t); err != nil {
			return err
		}
	}

	cs.tweets[t.Id] = t
	if a, ok := cs.finds[t.Author]; ok == true {
		for i := range a {
			if a[i].Id == t.Id {
				a[i] = t
			}
		}
	}
	return nil
}

// Delete removes an item by id from underlying storage, cache, and finds.
func (cs *CacheStorage) Delete(id int) {
	if cs.store != nil {
		cs.store.Delete(id)
	}

	if t, ok := cs.tweets[id]; ok == true {
		if a, ok := cs.finds[t.Author]; ok == true {
			for i := range a {
				if a[i].Id == id {
					cs.finds[t.Author] = append(a[:i], a[i+1:]...)
					break
				}
			}
		}
	}
	delete(cs.tweets, id)
}
