package tweetsave

type View interface {
	DisplayItem(t *tweet)
	//DisplayAll(tweets []*tweet)
	//DisplayResults(results []*tweet)
	//DisplayAddItem()
	//DisplayItemAdded(id int)
	//DisplayEditItem(t *tweet)
	//DisplayItemDeleted()

	DisplayError(err error, code int)
}
