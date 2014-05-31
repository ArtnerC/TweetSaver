package tweetsaver

type View interface {
	DisplayItem(t *Tweet)
	DisplayAll(tweets []*Tweet)
	DisplayResults(tweets []*Tweet, pos, total int)
	DisplayAddItem()
	DisplayItemAdded(id int)
	//DisplayEditItem(t *Tweet)
	//DisplayItemDeleted()

	DisplayError(err error, code int)
}
