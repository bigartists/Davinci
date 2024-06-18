package star

type Dao interface {
	Select(userId int64, targetId int64, target string) (*Star, error)
}
