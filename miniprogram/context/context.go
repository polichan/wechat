package context

import (
	"github.com/polichan/wechat/v2/credential"
	"github.com/polichan/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
