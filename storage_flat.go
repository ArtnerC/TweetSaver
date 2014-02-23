package tweetsave

import (
	"encoding/json"
	"os"
	"time"
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
	Tweets       []*tweet
}

// open handles opening or creating a new file to store tweets as JSON
func (fs *FileStorage) open() {
	var err error
	fs.file, err = os.OpenFile(StorageFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
}

// Get will get a tweet given the unique identifier as id
func (fs *FileStorage) Get(id int) (ret *tweet) {
	for _, t := range fs.GetAll() {
		if t.Id == id {
			ret = t
			break
		}
	}
	return
}

// GetAll will fetch and return all stored tweets
func (fs *FileStorage) GetAll() []*tweet {
	fs.open()
	defer fs.file.Close()

	if err := json.NewDecoder(fs.file).Decode(&fs.fileState); err != nil {
		panic(err)
	}
	fs.fileState.Items = len(fs.fileState.Tweets)
	return fs.fileState.Tweets
}

// Find returns an array with all tweets of the specified author
func (fs *FileStorage) Find(author string) (found []*tweet) {
	for _, v := range fs.GetAll() {
		if v.Author == author {
			found = append(found, v)
		}
	}
	return
}

// Add will save a new tweet and return the index or an error
func (fs *FileStorage) Add(t *tweet) (int, error) {
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
func (fs *FileStorage) Update(t *tweet) error {
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
func (fs *FileStorage) saveAll(tweets []*tweet) error {
	fs.open()
	defer fs.file.Close()
	fs.file.Truncate(0)

	fs.fileState.Tweets = tweets
	fs.fileState.Items = len(tweets)
	fs.fileState.LastModified = time.Now()
	defer func() { fs.fileState.Tweets = nil }()

	if StorageFormatNice == true {
		if err := EncodePretty(fs.file, &fs.fileState); err != nil {
			return err
		}
	} else {
		if err := json.NewEncoder(fs.file).Encode(&fs.fileState); err != nil {
			return err
		}
	}
	return nil
}
