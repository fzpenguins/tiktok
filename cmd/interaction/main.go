package main

import (
	"log"
	"net"

	"tiktok/cmd/interaction/dal"
	"tiktok/cmd/interaction/rpc"
	"tiktok/config"
	interaction "tiktok/kitex_gen/interaction/interactionservice"
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
	tracer.InitJaegerInKitex(constants.InteractionServiceName)
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
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8883")
	if err != nil {
		panic(err)
	}

	svr := interaction.NewServer(new(InteractionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "interaction"}),
		server.WithServiceAddr(addr), server.WithRegistry(r),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
	)
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6064", nil))
	}()
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
