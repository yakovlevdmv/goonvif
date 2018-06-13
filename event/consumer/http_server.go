package consumer

import (
	"bytes"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/use-go/goonvif/event"
	"github.com/use-go/goonvif/helper"
	"github.com/valyala/fasthttp"
)

//HTTPServer that onvif Device could post Event to
type HTTPServer struct {
	//local address for Consumer using
	Address string
	//transparent response compression or not
	Compress bool
	//connection
	listener net.Listener
	//as a Server end
	server         *fasthttp.Server
	NotifyConsumer func(string, string, *event.Notify, event.NotificationMessage) error
}

// this is default configuration
var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

//StartSrv Serving in One Http Endpoint
func (httpServer *HTTPServer) StartSrv() error {

	if httpServer.Address == "" {
		httpServer.Address = *addr
	}
	if httpServer.Compress {
		fasthttp.CompressHandler(httpServer.ServerRequestHandler)
	} else {
		httpServer.Compress = *compress
	}
	listener, err := net.Listen("tcp", httpServer.Address)
	if err != nil {
		log.Fatal(err.Error() + " : port occupied or ip unavailable")
		os.Exit(1)
		return err
	}
	httpServer.listener = listener
	server := &fasthttp.Server{
		Handler:      httpServer.ServerRequestHandler,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	httpServer.server = server
	log.Fatal(server.Serve(listener))

	return nil
}

//checkNotificationPath by a spec url format
func checkNotificationPath(pathBuf []byte) bool {

	//校验uri格式，需要是系列格式之一,如果格式不符，则为匿名消息
	// 1、/可以包含英文和文的非空白字符的路径/a1340608-a8a1-4495-8132-95303580e588
	// 2、/camera-in-gate-onvif-device/a1340608-a8a1-4495-8132-95303580e588/
	// 3、camera-in-gate-onvif-device/a1340608-a8a1-4495-8132-95303580e588/
	strFormat := `(/?)(?P<DeviceName>\S+?)/(?P<DeviceID>[0-9A-Za-z]{8}(-[0-9A-Za-z]{4}){3}-[0-9A-Za-z]{12})(/?)`
	reg := regexp.MustCompile(strFormat)
	return reg.Match(pathBuf)
}

func getDeviceIDAndName(pathBuf []byte) (deviceName string, deviceID string, err error) {

	if pathBuf == nil {
		return "", "", errors.New("no elemment found and the path empty")
	} else if len(pathBuf) == 1 || checkNotificationPath(pathBuf) == false {
		return "", "", errors.New("no available name and deivice can be found in path context")
	}
	//remove head
	if pathBuf[0] == '/' {
		pathBuf = pathBuf[1:]
	}
	lenth := len(pathBuf)

	if pathBuf[lenth-1] == '/' {
		pathBuf = pathBuf[:lenth-1]
	}
	bytesArray := bytes.SplitN(pathBuf, []byte{'/'}, 2)

	deviceName = string(bytesArray[0])
	deviceID = string(bytesArray[1])

	return deviceName, deviceID, nil
}

//ServerRequestHandler for Received Message
func (httpServer *HTTPServer) ServerRequestHandler(ctx *fasthttp.RequestCtx) {
	//TODO...
	//content := string(ctx.Request.Body())
	notify := event.Notify{}
	//reader := strings.NewReader(ctx.Request.Body())
	bytesReader := bytes.NewReader(ctx.Request.Body())
	err := helper.Unmarshal(bytesReader, "wsnt:Notify", &notify)

	if err != nil {
		log.Print("Error Message Received")
		log.Print(bytesReader)
	} else if len(notify.NotificationMessagesList) > 0 {

		if deviceName, deviceID, err := getDeviceIDAndName(ctx.Path()); err != nil {
			log.Print(err.Error())
		} else {
			for _, notificationMessageItem := range notify.NotificationMessagesList {
				//notify the consumer
				if httpServer.NotifyConsumer != nil {
					httpServer.NotifyConsumer(deviceName, deviceID, &notify, notificationMessageItem)
				}
			}
		}

	} else {
		log.Print("No Available Message Arrived")
	}

	httpServer.replyStatusOK(ctx)

}

//SendResponse HTTP Service
func (httpServer *HTTPServer) replyStatusOK(ctx *fasthttp.RequestCtx) {

	//Send Notifications to consumer
	//Seperate different message
	/*
		HTTP/1.1 200 OK
		Date: Tue, 10 Apr 2018 15:52:16 GMT
		Content-Type: text/plain
		Content-Length: 24
		Accept-Ranges: bytes
		Connection: close
		Notify Message Received!
	*/

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("text/plain")
	ctx.SetBodyString("Notify Message Received!")
}

//StopSrv HTTP Service
func (httpServer *HTTPServer) StopSrv() {
	httpServer.listener.Close()
}
