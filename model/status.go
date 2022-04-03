package model

// UserStatus 用户状态信息
type UserStatus struct {
	Model `xorm:"extends"`
	//StatusId        uint      `json:"id" xorm:"pk"`
	//CreatedTime     time.Time `json:"created_time" xorm:"datetime created"`
	//UpdatedTime     time.Time `json:"updated_time" xorm:"datetime updated"`
	DrivingStatus   string  `json:"drivingStatus" xorm:"varchar(255)"`
	HeartRate       float64 `json:"heartRate" xorm:"Double"`       //心跳
	Temperature     float64 `json:"temperature" xorm:"Double"`     //体温
	AlcoholStrength float64 `json:"alcoholStrength" xorm:"Double"` //酒精浓度
}
