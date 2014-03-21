package tweetsaver

type View interface {
	DisplayItem(t *Tweet)
	DisplayAll(tweets []*Tweet)
	//DisplayResults(results []*Tweet)
	//DisplayAddItem()
	//DisplayItemAdded(id int)
	//DisplayEditItem(t *Tweet)
	//DisplayItemDeleted()

	DisplayError(err error, code int)
}
