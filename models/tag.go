package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created-by"`
	ModifiedBy string `json:"modified-by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
