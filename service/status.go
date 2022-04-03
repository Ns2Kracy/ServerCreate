package service

//TODO:完成数值得返回HeartRate，AlcoholStren， Temperature

import (
	"ServerCreate/model"
	"xorm.io/xorm"
)

type UserStatusService struct {
	db *xorm.Engine
}

// 检查Id返回所有状态信息
func (uss *UserStatusService) CheckStatusId(statusId uint) (model.UserStatus, error) {
	var status model.UserStatus
	if _, err := uss.db.Table("users").Where(" user_id = ? ", statusId).Get(&status); err != nil {
		return model.UserStatus{}, err
	}
	return status, nil
}

func (uss *UserStatusService) StatusInserter(status model.UserStatus) bool {
	//TODO:获取从硬件设备得到的信息并保存至数据库以供调用
	//设备->MQTT->Kafaka->Go Server->App
	return true
}

func (uss *UserStatusService) StatusUpdater(status model.UserStatus) error {
	//更新用户状态信息
	if _, err := uss.db.Table("user_status").Update(status); err != nil {
		return err
	}
	return nil
}

func (uss *UserStatusService) StatusDeleter(status model.UserStatus) error {
	//将用户状态信息进行软删除
	if _, err := uss.db.Table("user_status").AllCols().Delete(status); err != nil {
		return err
	}
	return nil
}

/*
 * 驾驶状态判断相关
 */

//判断驾驶状态并根据此进行报警...
func (uss *UserStatusService) CheckDrivingStatus(DrivingStatus string) error {
	//TODO:实现判断

	// 状态: 抽烟； 分心； 瞌睡； 哈欠；...
	return nil
}

/*
 * 心率相关服务
 */

// 区间段最大心跳
func (uss *UserStatusService) MaxHeartRateReport(heartRate float64) interface{} {
	//TODO:获取一段时间区间内心跳的最大值

	//var Exec = "SELECT MAX(heart_rate) from user_status;"

	MaxHeartRate, _ := uss.db.Table("statusreport").Exec(" SELECT MAX(heart_rate) `heart_rate` FORM user_status ")

	return MaxHeartRate
}

// 区间段最小心跳
func (uss *UserStatusService) MinHeartRateReport(heartRate float64) interface{} {
	//TODO:获取一段时间区间内心跳的最小值
	//Exec := "select max(sum) from user_status where logdate between '03:03:03' and '03:07:03';"
	MinHeartRate, _ := uss.db.Table("statusreport").Exec(" SELECT MIN(heart_rate) `heart_rate` FORM user_status ")
	return MinHeartRate
}

// 区间段平均心跳
func (uss *UserStatusService) AverageHeartReport(StatusType interface{}) interface{} {
	//TODO:获取一段时间区间内的平均

	AvgHeartRate, _ := uss.db.Table("statusreport").Exec(" SELECT AVG(heart_rate) `heart_rate` FORM user_status ")
	return AvgHeartRate
}

/*
 * 温度相关服务
 */

// 区间段最大体温
func (uss *UserStatusService) MaxTemperatureReport(temperature float64) interface{} {
	//TODO:获取一段时间区间内的最大值

	MaxTemperature, _ := uss.db.Table("statusreport").Exec(" SELECT MAX(temperature) `temperature` FORM user_status ")
	return MaxTemperature
}

// 区间段最小体温
func (uss *UserStatusService) MinTemperatureReport(temperature float64) error {
	//TODO:获取一段时间区间内的最小值
	var status model.UserStatus
	uss.db.Table("statusreport").Where(" temperature = ?").Get(&status)
	return nil
}

// 区间段平均体温
func (uss *UserStatusService) AverageTemperatureReport(temperature float64) interface{} {
	//TODO:获取一段时间区间内的平均
	AverageTemperature, _ := uss.db.Table("statusreport").Exec(" SELECT AVG(temperature) `temperature` FORM user_status ")
	return AverageTemperature
}
