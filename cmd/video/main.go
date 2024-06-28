package main

import (
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"tiktok/cmd/video/dal"
	"tiktok/cmd/video/rpc"
	"tiktok/config"
	video "tiktok/kitex_gen/video/videoservice"
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
	tracer.InitJaegerInKitex(constants.VideoServiceName)
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
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8882")
	if err != nil {
		panic(err)
	}
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "video"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
	)

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6061", nil))
	}()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
