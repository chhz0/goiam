package gormutil

const DefaultLimit = 1000

type LimitAndOffset struct {
	Limit  int
	Offset int
}

func UnpointerLO(limit, offset *int64) LimitAndOffset {
	var l, o int = DefaultLimit, 0

	if limit != nil {
		l = int(*limit)
	}

	if offset != nil {
		o = int(*offset)
	}

	return LimitAndOffset{
		Limit:  l,
		Offset: o,
	}
}
