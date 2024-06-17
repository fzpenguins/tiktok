package main

import (
	"log"
	"net"
	"tiktok/cmd/user/dal"
	"tiktok/config"
	user "tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/tracer"

	"net/http"
	_ "net/http/pprof"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	config.ReadConfig()

	dal.Init()

	tracer.InitJaegerInKitex(constants.UserServiceName)

	log.Println("successfully running...")
	klog.SetLevel(klog.LevelDebug)

}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8881")
	if err != nil {

		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
	)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
