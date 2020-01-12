package service

import (
	"github.com/youzeliang/gopractice/model"
	"github.com/youzeliang/gopractice/utils"
	"time"
)

func Test() {
	time := time.Now()
	var relation = model.FrameWorkOrderRelation{
		WorkOrderID:  utils.GetUuidString(),
		WarehouseID: utils.GetUuidString(),
		ActionWaveID: utils.GetUuidString(),
		ContainerID:  utils.GetUuidString(),
		Status:       model.Valid,
		CreateTime:   time,
		UpdateTime:   time,
	}

	err := relation.Insert(nil)
	if err != nil {

	}
}
