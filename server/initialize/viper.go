package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"server/global"
)

func InitViper() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	config.AddConfigPath(path + "/")
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	err = config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听配置文件变化
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = config.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置文件内容解析到config结构体中
	if err = config.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

}

// 单体应用这里就用viper，后期微服改用nacos的远程配置管理
