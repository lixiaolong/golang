package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Content struct {
	Id      int64
	Time    string
	Title   string
	Details string
}

type Timeline struct {
	Title    string
	Contents []Content
}

func NewTimeline() *Timeline {
	return &Timeline{Title: "Timeline"}
}

func (this *Timeline) AppendLine(time, title, detail string) {
	line := &Content{Time: time, Title: title, Details: detail}
	this.Contents = append(this.Contents, *line)
}

func (this *Timeline) SetTitle(t string) {
	this.Title = t
}

func (this *Timeline) SaveDB(db string) {
	orm, err := xorm.NewEngine("sqlite3", db)
	if err != nil {
		log.Fatalln(err)
	}
	err = orm.Sync(new(Content))
	if err != nil {
		log.Fatalln(err)
	}
	session := orm.NewSession()
	session.Begin()
	session.InsertMulti(this.Contents)
	session.Commit()
}

func (this *Timeline) LoadDB(db string, offset, within int) {
	orm, err := xorm.NewEngine("sqlite3", db)
	if err != nil {
		log.Fatalln(err)
	}
	this.Contents = make([]Content, 0)

	if -1 == offset || -1 == within {
		orm.Desc("id").Find(&this.Contents)
	} else {
		orm.Desc("id").Limit(within, offset).Find(&this.Contents)
	}

}
