package tweetsave

import (
	"strings"
	"testing"
	"time"
)

func init() {
	StorageFileName = "Test" + StorageFileName
	StorageFormatNice = true

	baseTestStorage.fileState.CurrentID = len(ExampleTweets)
	baseTestStorage.saveAll(ExampleTweets)
}

var baseTestStorage = new(FileStorage)

var testStorage Persistence = baseTestStorage

//var testStorage Persistence = NewStorageCache(baseTestStorage)

func TestGet(t *testing.T) {
	v := testStorage.Get(0)
	if v.Text != ExampleTweets[0].Text {
		t.Fail()
	}
	if v.Author != ExampleTweets[0].Author {
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	res := testStorage.GetAll()
	if len(res) != len(ExampleTweets) {
		t.Fail()
	}
}

func TestFind(t *testing.T) {
	results := testStorage.Find(ExampleTweets[2].Author)
	if len(results) != 1 {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	tweets := testStorage.Find(ExampleTweets[3].Author)
	for _, v := range tweets {
		testStorage.Delete(v.Id)
	}

	newTweet := *ExampleTweets[3]
	newTweet.SaveTime = time.Now()

	id, err := testStorage.Add(&newTweet)
	if err != nil {
		t.Fail()
	}
	if v := testStorage.Get(id); v.Author != ExampleTweets[3].Author {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	tweet := *ExampleTweets[1]
	tweet.Author = strings.Replace(tweet.Author, "_", " ", -1)
	testStorage.Update(&tweet)
	comp := testStorage.Get(tweet.Id)
	if comp.Author != tweet.Author {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	tweets := testStorage.GetAll()
	length := len(tweets) - 1
	last := tweets[length]
	tweet := testStorage.Get(last.Id)
	defer testStorage.Add(tweet)

	testStorage.Delete(last.Id)
	if len(testStorage.GetAll()) != length {
		t.Fail()
	}
}
