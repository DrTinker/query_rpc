package config

import (
	"errors"
	"fmt"
	"query_rpc/conf"
	"query_rpc/models"
	"query_rpc/pkg/helper"

	"gopkg.in/ini.v1"
)

type ConfigClientImpl struct {
	DB     *models.DBConfig
	Rpc    *models.RpcConfig
	source *ini.File
}

func NewConfigClientImpl() *ConfigClientImpl {
	return &ConfigClientImpl{
		DB:  &models.DBConfig{},
		Rpc: &models.RpcConfig{},
	}
}

func (c *ConfigClientImpl) Load(path string) error {
	var err error
	//判断配置文件是否存在
	exists, _ := helper.PathExists(path)
	if !exists {
		return errors.New("Config path not exists!")
	}
	c.source, err = ini.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func (c *ConfigClientImpl) GetRPCConfig() (*models.RpcConfig, error) {
	//判断配置是否加载成功
	if c.source == nil {
		return nil, errors.New("empty rpc config")
	}
	section := c.source.Section("RpcServer")
	c.Rpc = &models.RpcConfig{}
	c.Rpc.Address = section.Key("address").MustString("127.0.0.1")
	c.Rpc.Port = section.Key("port").MustInt(50052)
	c.Rpc.ClientPoolConnsSizeCap = section.Key("clientPoolConnsSizeCap").MustInt(conf.DefaultClientPoolConnsSizeCap)
	c.Rpc.DialTimeout = section.Key("dialTimeout").MustInt(int(conf.DefaultDialTimeout))
	c.Rpc.KeepAlive = section.Key("keepAlive").MustInt(int(conf.DefaultKeepAlive))
	c.Rpc.KeepAliveTimeout = section.Key("keepAliveTimeout").MustInt(int(conf.DefaultKeepAliveTimeout))
	return c.Rpc, nil
}

func (c *ConfigClientImpl) GetDBConfig() (driver, source string, err error) {
	//判断配置是否加载成功
	if c.source == nil {
		return "", "", errors.New("Empty DB config!")
	}
	c.DB.Type = c.source.Section("DB").Key("type").MustString("mysql")
	c.DB.User = c.source.Section("DB").Key("user").MustString("root")
	c.DB.IP = c.source.Section("DB").Key("ip").MustString("127.0.0.1")
	c.DB.Pwd = c.source.Section("DB").Key("pwd").MustString("xiaonajia123")
	c.DB.DB = c.source.Section("DB").Key("db").MustString("query_system")
	c.DB.Port = c.source.Section("DB").Key("port").MustInt(3306)

	driver = c.DB.Type
	source = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		c.DB.User, c.DB.Pwd, c.DB.IP, c.DB.Port, c.DB.DB)

	return driver, source, nil
}
