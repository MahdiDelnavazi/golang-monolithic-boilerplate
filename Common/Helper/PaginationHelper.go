package Helper

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPaginate struct {
	limit int64
	page  int64
}

func NewMongoPaginate(limit, page int) *MongoPaginate {
	return &MongoPaginate{
		limit: int64(limit),
		page:  int64(page),
	}
}

func (mp *MongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	return &fOpt
}
