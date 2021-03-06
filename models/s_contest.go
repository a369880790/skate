package models

type SContest struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(3)"`
	Name        string `json:"name" xorm:"not null default '' VARCHAR(45)"`
	Station     string `json:"station" xorm:"not null default '' VARCHAR(45)"`
	Type        string `json:"type" xorm:"not null default '' VARCHAR(45)"`
	ContestTime string `json:"contest_time" xorm:"not null default '' VARCHAR(45)"`
	CreateTime  string `json:"create_time" xorm:"not null default '' VARCHAR(45)"`
	UpdateTime  string `json:"update_time" xorm:"not null default '' VARCHAR(45)"`
	DeleteTime  string `json:"delete_time" xorm:"not null default '' VARCHAR(45)"`
}


func NewContest() *SContest {
	return new(SContest)
}

func MoreContest() []*SContest {
	return make([]*SContest, 0)
}
