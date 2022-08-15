package user

import (
	"log"
	"time"

	"github.com/tanweiquan/webmake/dao/db"
)

func init() {
	db.SyncAppend(new(User))
}

type User struct {
	Id         string    `json:"id" xorm:"varchar(36) pk"`
	Name       string    `json:"name" xorm:"varchar(20) default ''"`
	Age        int       `json:"age" xorm:"int default 0 "`
	CreateDate time.Time `json:"create_date" xorm:"created notnull"`
	UpdateDate time.Time `json:"update_date" xorm:"updated notnull"`
	Birthday   time.Time `json:"birthday" xorm:"datetime null default null"`
}

func (t *User) ListAll() []User {
	record := make([]User, 0)
	err := db.GetDbengine().Find(&record)
	if err != nil {
		log.Println(err)
	}
	return record
}
