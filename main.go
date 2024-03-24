package main

import (
	"fmt"
	"member_service_frame/config"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis"
	"google.golang.org/grpc"
)

var deployTime string = time.Now().UTC().Format("2006-01-02 15:04:05") + " UTC"

const version string = "1.0.0"
const dbName string = "Member"
const redisPool int = 0

func main() {
	// setting middleWare
	var server *gin.Engine = config.GetEngineWithMiddleWare()

}

func testerURL(ctx *gin.Context) {
	hostName, _ := os.Hostname()
	ctx.JSON(http.StatusOK, gin.H{
		"Deployment Time": deployTime,
		"Version":         version,
		"Host":            hostName,
	})
}

func handleNoRoute(ctx *gin.Context) {
	ctx.File("no_route.html")
}

func openGRPCService(grpcPort string) {
	var redisClient *redis.Client = r.Connect(redisPool)
	var loginRepo repository.LoginTimeInterface = redisLogincheck.NewRedisLoginCheckerRepository(redisClient)

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.HandleSevereCrashed(err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterAuthorizerServer(s, &controller.Server{RedisClient: loginRepo})
	fmt.Printf("gRPC Server is listening on port %v \n", grpcPort)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
