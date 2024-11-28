package utils

import "go.uber.org/zap"

var Log *zap.Logger

func InitLogger() {
	var err error
	Log, err = zap.NewProduction() // 使用生产模式日志
	if err != nil {
		panic(err)
	}
}
