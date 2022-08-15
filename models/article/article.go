package article

import (
	"log"

	"github.com/tanweiquan/webmake/dao/db"
)

func init() {
	db.SyncAppend(new(Article))
}

type Article struct {
	Id      string `json:"id" xorm:"varchar(36) pk"`
	Article string `json:"article" xorm:"text notnull"`
	Author  string `json:"author" xorm:"varchar(50) index notnull"`
}

func (t *Article) ListAll() []Article {
	record := make([]Article, 0)
	err := db.GetDbengine().Find(&record)
	if err != nil {
		log.Println(err)
	}
	return record
}
