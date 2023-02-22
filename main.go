package main

import (
	"douyin/dao"
	"douyin/global"
	"douyin/pkg/logger"
	"douyin/pkg/setting"
	"douyin/router"
	"fmt"
	"log"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupMysqlEngine()
	if err != nil {
		log.Fatalf("init.setupMysqlEngine err: %v", err)
	}
	err = setupRedisEngine()
	if err != nil {
		log.Fatalf("init.setupRedisEngine err: %v", err)
	}
	err = setupOSSEngine()
	if err != nil {
		log.Fatalf("init.setupOSSEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func main() {
	r := router.NewRouter()
	fmt.Println("redis settings:", global.RedisSetting)
	fmt.Println("server settings:", global.ServerSetting)
	fmt.Println("mysql settings:", global.MysqlSetting)
	fmt.Println("logger settings", global.LoggerSetting)

	err := r.Run(fmt.Sprintf(":%d", global.ServerSetting.HttpPort))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}

// 初始化数据库配置
func setupMysqlEngine() error {
	var err error
	global.MysqlEngine, err = dao.NewMysqlEngine(global.MysqlSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化redis配置
func setupRedisEngine() error {
	var err error
	global.RedisEngine, err = dao.NewRedisEngine(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化OSS配置
func setupOSSEngine() error {
	var err error
	global.OSSEngine, err = dao.NewOSSEngine(global.OSSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志配置
func setupLogger() error {
	var err error
	global.Logger, err = logger.NewLogger(global.LoggerSetting, global.ServerSetting.RunMode)
	if err != nil {
		return err
	}
	return nil
}

// 初始化环境配置
func setupSetting() error {
	// 使用viper进行配置
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("redis", &global.RedisSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("mysql", &global.MysqlSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("log", &global.LoggerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("oss", &global.OSSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
