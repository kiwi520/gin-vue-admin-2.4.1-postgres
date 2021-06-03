package core

import (
	"fmt"
	"gin-vue-admin/global"
	_ "gin-vue-admin/packfile"
	"gin-vue-admin/utils"
	"github.com/spf13/pflag"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {

	var config string
	if len(path) == 0 {
		if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {

			var DBIp = pflag.String("ip", "127.0.0.1", "数据库ip地址")
			var DBPort = pflag.Int64("port", 8080, "数据库端口号")
			var DBName = pflag.String("db", "gva", "请输入数据库名称")
			var DBUsername = pflag.String("username", "gva", "请输入数据库用户名")
			var DBPassword = pflag.String("password", "gva", "请输入数据库用户密码")

			// 解析命令行选项。
			pflag.Parse()

			// 绑定选项到 viper 中。
			if err := viper.BindPFlags(pflag.CommandLine); err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			config = utils.ConfigFile
			vn := viper.New()
			//设置配置文件的名字
			vn.SetConfigFile(config)
			if err := vn.ReadInConfig(); err != nil {
				if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					// Config file not found; ignore error if desired
					log.Println("no such config file")
				} else {
					// Config file was found but another error was produced
					log.Println("read config error")
				}
				log.Fatal(err) // 读取配置文件失败致命错误
			}
			vn.Set("postgresql.host", DBIp)
			vn.Set("postgresql.port", DBPort)
			vn.Set("postgresql.db-name", DBName)
			vn.Set("postgresql.username", DBUsername)
			vn.Set("postgresql.password", DBPassword)

			if err := vn.WriteConfig(); err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
		} else {
			config = configEnv
			fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
