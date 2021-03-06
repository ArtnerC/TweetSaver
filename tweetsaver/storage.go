package tweetsaver

type Persistence interface {
	Get(id int) *Tweet
	GetAt(pos, limit int) ([]*Tweet, int)
	GetAll() []*Tweet
	Find(author string) []*Tweet
	Add(t *Tweet) (int, error)
	Update(t *Tweet) error
	Delete(id int)
}
