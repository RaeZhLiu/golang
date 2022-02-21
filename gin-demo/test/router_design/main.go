package main

import (
	"fmt"
	"goyard/gin-demo/test/router_design/router/app/blog"
	"goyard/gin-demo/test/router_design/router/app/shop"
	"goyard/gin-demo/test/router_design/router/routers"
)

func main() {
	//r := gin.Default()
	//
	//blog.LoadBlog(r)
	//shop.LoadShop(r)
	//_ = r.Run(":9090")
	//加载多个APP的路由配置
	routers.Include(shop.Routers, blog.Routers)
	//初始化路由
	r := routers.Init()
	if err := r.Run(":9090"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
