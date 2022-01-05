package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	// 設置配置中心的地址
	address := host + ":" + strconv.FormatInt(port, 10)

	consulSource := consul.NewSource(
			// 設置配置中心的地址
			consul.WithAddress(address),
			// 設置前綴, 不設置的畫默認為 "micro/config"
			consul.WithPrefix(prefix),
			// 是否移除前綴, 這裡設置為true, 表示可以不帶前綴直接取得對應配置
			consul.StripPrefix(true),
		)

	config, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	// 載入配置
	err = config.Load(consulSource)

	return config, err
}
