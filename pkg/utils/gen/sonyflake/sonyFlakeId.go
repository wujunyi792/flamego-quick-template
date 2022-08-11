package sonyflake

import (
	"github.com/sony/sonyflake"
	"github.com/wujunyi792/gin-template-new/internal/loging"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenSonyFlakeId() (int64, error) {

	id, err := flake.NextID()
	if err != nil {
		loging.Warning.Println("flake NextID failed: ", err)
		return 0, err
	}

	return int64(id), nil
}
