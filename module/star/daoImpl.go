package star

import "davinci/client"

var DaoGetter Dao

type DaoGetterImpl struct {
}

func init() {
	DaoGetter = NewDaoGetterImpl()
}

func NewDaoGetterImpl() *DaoGetterImpl {
	return &DaoGetterImpl{}
}

func (d DaoGetterImpl) Select(userId int64, targetId int64, target string) (*Star, error) {
	var star Star
	err := client.Orm.
		Table("star").
		Where("target = ? and target_id = ? and user_id = ?", target, targetId, userId).
		First(&star).
		Error

	if err != nil {
		return nil, err
	}
	return &star, nil
}
