package svc

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:38:21
 * @LastEditTime: 2024-03-08 14:38:21
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	_ "github.com/xxcheng123/primary-school-system/ent/runtime"
	"github.com/xxcheng123/primary-school-system/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	err := db.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
