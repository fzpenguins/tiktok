package main

import (
	"log"
	"net"
	"net/http"
	"tiktok/cmd/follow/dal"
	"tiktok/cmd/follow/rpc"
	"tiktok/config"
	follow "tiktok/kitex_gen/follow/followservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/tracer"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {

	config.ReadConfig()
	dal.Init()
	tracer.InitJaegerInKitex(constants.FollowServiceName)
	rpc.Init()
	log.Println("successfully running...")
	klog.SetLevel(klog.LevelDebug)

}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.EtcdAddr})
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8884")
	if err != nil {
		panic(err)
	}

	svr := follow.NewServer(new(FollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "follow"}),
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
