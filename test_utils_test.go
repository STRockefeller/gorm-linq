package linq

type TestTable struct {
	ID      int    `gorm:"primaryKey"`
	Column1 string `gorm:"column:column1"`
	Column2 string `gorm:"column:column2"`
}

type updateReq struct {
	query         []QueryString
	updatedValues map[string]any
}

func (ur updateReq) Update() map[string]any {
	return ur.updatedValues
}

func (ur updateReq) Where() []QueryString {
	return ur.query
}

func (ur updateReq) whereTest(condition QueryString) updateReq {
	ur.query = append(ur.query, condition)
	return ur
}

func (ur updateReq) updateTest(column string, value any) updateReq {
	ur.updatedValues[column] = value
	return ur
}
