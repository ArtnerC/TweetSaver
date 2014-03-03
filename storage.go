package tweetsaver

type Persistence interface {
	Get(id int) *tweet
	GetAll() []*tweet
	Find(author string) []*tweet
	Add(t *tweet) (int, error)
	Update(t *tweet) error
	Delete(id int)
}
