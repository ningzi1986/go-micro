package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"go-micro/go-micro-part03/plugins/breaker"
	"net"
	"net/http"

	"go-micro/go-micro-part03/basic"

	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro/go-micro-part03/basic/config"
	_ "go-micro/go-micro-part03/plugins/zap"
	"go-micro/go-micro-part03/web/handler"
)

func main() {

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)

	// 初始化配置
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.employee"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8086"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// 员工信息
	service.HandleFunc("/employee/employees", handler.GetEmployees)
	// 登录
	//service.HandleFunc("/employee/login", handler.Login)
	handlerLogin := http.HandlerFunc(handler.Login)
	service.Handle("/employee/login", breaker.BreakerWrapper(handlerLogin))


	// 退出登录
	service.HandleFunc("/employee/logout", handler.Logout)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	host := config.AppConfig.App.Etcd.Host
	port := config.AppConfig.App.Etcd.Port
	ops.Addrs = []string{fmt.Sprintf("%s:%d", host, port)}
}
