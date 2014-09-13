package modules

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/memcache"
	"github.com/astaxie/beego/context"
)

var AppCache cache.Cache

func init() {
	beego.Info("init cache.")
	// c, _ := cache.NewCache("memory", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120,"interval":60}`)
	c, _ := cache.NewCache("memcache", `{"conn":"127.0.0.1:11211"}`)
	AppCache = c
}

func WriteJson(ctx *context.Context, bytes []byte) {
	ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	ctx.Output.Body(bytes)
}

func WriteInterfaceJson(ctx *context.Context, key string, response interface{}) {
	content, err := json.Marshal(&response)
	if err != nil {
		beego.Error(err)
	}
	AppCache.Put(key, string(content), 3600)
	WriteJson(ctx, content)
}

func JsonRequestCache(ctx *context.Context, key string) bool {

	cacheData := AppCache.Get(key)
	if cacheData == nil {
		return false
	}

	var str string = cacheData.(string)
	WriteJson(ctx, []byte(str))

	return true
}
