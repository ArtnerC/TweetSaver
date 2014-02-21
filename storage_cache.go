package tweetsave

type CacheStorage struct {
	tweets    map[int]*tweet
	finds     map[string][]*tweet
	store     Persistence
	allLoaded bool
}

// NewStorageCache creates a CacheStorage object that applies
// a caching layer to the provided p which implements the
// Persistence interface.
func NewStorageCache(p Persistence) *CacheStorage {
	return &CacheStorage{
		tweets:    make(map[int]*tweet),
		finds:     make(map[string][]*tweet),
		store:     p,
		allLoaded: false,
	}
}

// NewMemoryStorage creates a CacheStorage object that can
// be used to store data in volatile RAM without the need
// for an underlying Persistence object.
func NewMemoryStorage() *CacheStorage {
	return NewStorageCache(nil)
}

func (cs *CacheStorage) Get(id int) *tweet {
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

func (cs *CacheStorage) GetAll() []*tweet {
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
		tweets := make([]*tweet, 0, len(cs.tweets))
		for _, v := range cs.tweets {
			tweets = append(tweets, v)
		}
		if len(tweets) > 0 {
			return tweets
		}
	}

	return nil
}

func (cs *CacheStorage) Find(author string) []*tweet {
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

	found := make([]*tweet, 0, 10)
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

func (cs *CacheStorage) Add(t *tweet) (int, error) {
	var id int
	if cs.store != nil {
		if id, err := cs.store.Add(t); err != nil {
			return id, err
		}
	} else {
		id = len(cs.tweets)
	}

	cs.tweets[id] = t
	if a, ok := cs.finds[t.Author]; ok == true {
		cs.finds[t.Author] = append(a, t)
	}
	return id, nil
}

func (cs *CacheStorage) Update(t *tweet) error {
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
