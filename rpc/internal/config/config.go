package config

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 14:37:45
 * @LastEditTime: 2024-03-08 14:37:46
 */

import (
	"gitee.com/study-helper/common-go/db/dbtools"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf dbtools.DatabaseConf
}
