package main

import (
	"log"
	"net"
	"net/http"
	"tiktok/cmd/picture/dal"
	"tiktok/cmd/picture/rpc"
	"tiktok/config"
	picture "tiktok/kitex_gen/picture/pictureservice"
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
	tracer.InitJaegerInKitex(constants.PictureServiceName)
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
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8885")
	if err != nil {
		panic(err)
	}

	svr := picture.NewServer(new(PictureServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "picture"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(opentracing.NewDefaultServerSuite()),
	)

	go func() {
		log.Println(http.ListenAndServe("localhost:6065", nil))
	}()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

//func main() {
//	svr := picture.NewServer(new(PictureServiceImpl))
//
//	err := svr.Run()
//
//	if err != nil {
//		log.Println(err.Error())
//	}
//}
