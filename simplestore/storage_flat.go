package simplestore

import (
	"encoding/json"
	"os"
	"sort"
	"time"

	ts "github.com/ArtnerC/TweetSaver/tweetsaver"
)

var StorageFileName = "Tweets.json" // Target file
var StorageFormatNice = true        // Pretty Print output (uses more space)

// FileStorage implements the Persistence interface as a flat file with JSON
// encoded data. Operations are expensive and require reading/writing the
// entire file. This is an alternative/placeholder for a database.
type FileStorage struct {
	file      *os.File
	fileState FileStruct
}

// FileStruct represents the structure of the file the data is stored in
type FileStruct struct {
	CurrentID    int
	Items        int
	LastModified time.Time
	Tweets       []*ts.Tweet
}

type TweetSorter []*ts.Tweet

func (tsort TweetSorter) Len() int           { return len(tsort) }
func (tsort TweetSorter) Swap(i, j int)      { tsort[i], tsort[j] = tsort[j], tsort[i] }
func (tsort TweetSorter) Less(i, j int) bool { return tsort[i].Timestamp.Before(tsort[j].Timestamp) }

// open handles opening or creating a new file to store tweets as JSON
func (fs *FileStorage) open() {
	var err error
	fs.file, err = os.OpenFile(StorageFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
}

// Get will get a tweet given the unique identifier as id
func (fs *FileStorage) Get(id int) (ret *ts.Tweet) {
	for _, t := range fs.GetAll() {
		if t.Id == id {
			ret = t
			break
		}
	}
	return
}

// GetAt will return a slice of tweets at the given position with the selected
// limit of returned results
func (fs *FileStorage) GetAt(pos, limit int) ([]*ts.Tweet, int) {
	tweets := fs.GetAll()
	if pos >= len(tweets) {
		return nil, 0
	}
	end := pos + limit
	if end > len(tweets) {
		end = len(tweets)
	}

	sort.Sort(TweetSorter(tweets))

	return tweets[pos:end], len(tweets)
}

// GetAll will fetch and return all stored tweets
func (fs *FileStorage) GetAll() []*ts.Tweet {
	fs.open()
	defer fs.file.Close()

	if err := json.NewDecoder(fs.file).Decode(&fs.fileState); err != nil {
		panic(err)
	}
	fs.fileState.Items = len(fs.fileState.Tweets)
	return fs.fileState.Tweets
}

// Find returns an array with all tweets of the specified author
func (fs *FileStorage) Find(author string) (found []*ts.Tweet) {
	for _, v := range fs.GetAll() {
		if v.Author == author {
			found = append(found, v)
		}
	}
	return
}

// Add will save a new tweet and return the index or an error
func (fs *FileStorage) Add(t *ts.Tweet) (int, error) {
	tweets := fs.GetAll()
	id := fs.fileState.CurrentID
	fs.fileState.CurrentID++

	t.Id = id
	tweets = append(tweets, t)

	if err := fs.saveAll(tweets); err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates the provided tweet and returns an error or nil
func (fs *FileStorage) Update(t *ts.Tweet) error {
	tweets := fs.GetAll()
	for i, v := range tweets {
		if v.Id == t.Id {
			tweets[i] = t
			break
		}
	}
	return fs.saveAll(tweets)
}

// Delete will remove a tweet at the specified id
func (fs *FileStorage) Delete(id int) {
	tweets := fs.GetAll()
	for i, v := range tweets {
		if v.Id == id {
			tweets[i] = nil
			tweets = append(tweets[:i], tweets[i+1:]...)
			break
		}
	}
	fs.saveAll(tweets)
}

// saveAll is a private helper to facilitate re-saving the tweets file
func (fs *FileStorage) saveAll(tweets []*ts.Tweet) error {
	fs.open()
	defer fs.file.Close()
	fs.file.Truncate(0)

	fs.fileState.Tweets = tweets
	fs.fileState.Items = len(tweets)
	fs.fileState.LastModified = time.Now()
	defer func() { fs.fileState.Tweets = nil }()

	if StorageFormatNice == true {
		if err := ts.EncodePretty(fs.file, &fs.fileState); err != nil {
			return err
		}
	} else {
		if err := json.NewEncoder(fs.file).Encode(&fs.fileState); err != nil {
			return err
		}
	}
	return nil
}
