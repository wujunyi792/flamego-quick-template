package utils

import (
	"github.com/bwmarrin/snowflake"
)

func MakeUUID() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}

	return node.Generate().String(), err

}
