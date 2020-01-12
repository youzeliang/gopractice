package model

import (
	"github.com/gin-gonic/gin"
	"github.com/youzeliang/gopractice/helper"
	"time"
)

const (
	Valid   = 1 // 有效
	Invalid = 0 // 无效
)

type FrameWorkOrderRelation struct {
	ID           int64     `gorm:"column:id;primary_key" json:"id"`
	WorkOrderID  string    `gorm:"column:work_order_id" json:"workOrderID"`
	WarehouseID  string    `gorm:"column:warehouse_id" json:"warehouseID"`
	ActionWaveID string    `gorm:"column:action_wave_id" json:"actionWaveID"`
	ContainerID  string    `gorm:"column:container_id" json:"containerID"`
	Status       int       `gorm:"column:status" json:"status"`
	CreateTime   time.Time `gorm:"column:create_time" json:"-"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"-"`
}

func (FrameWorkOrderRelation) TableName() string {
	return "action_wave_repeat"
}

func (p *FrameWorkOrderRelation) Insert(c *gin.Context) (err error) {
	sql := "INSERT INTO `action_wave_repeat` " +
		"(`action_wave_id`, `warehouse_id`, `work_order_id`,`container_id`, `create_time`, `update_time`, `status`) VALUES " +
		"(?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE " +
		"`update_time`=?, status=?"

	dbErr := helper.MysqlClient.Exec(sql, p.ActionWaveID, p.WarehouseID, p.WorkOrderID, p.ContainerID, p.CreateTime, p.UpdateTime, p.Status, p.UpdateTime, p.Status).Error

	if dbErr != nil {

		return err
	}
	return nil
}

func UpdateStatusByActionWaveID(warehouseID, actionWaveID string, status int, c *gin.Context) (err error) {
	dbErr := helper.MysqlClient.Model(&FrameWorkOrderRelation{}).
		Where("warehouse_id = ? and action_wave_id = ?", warehouseID, actionWaveID).
		Updates(map[string]interface{}{
			"status":      status,
			"update_time": time.Now(),
		}).Error
	if dbErr != nil {
		return err
	}
	return nil
}

func QueryRepeatWaveIDList(warehouseID string, c *gin.Context) (frameWorkOrderList []FrameWorkOrderRelation, err error) {
	dbErr := helper.MysqlClient.
		Where(" warehouse_id = ? and  status = ? ", warehouseID, Valid).
		Find(&frameWorkOrderList).Error
	if dbErr != nil {
		return frameWorkOrderList,err
	}

	return frameWorkOrderList, nil
}
