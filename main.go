package main

import (
	"database/sql"
	"log"
	"member_service_frame/config"
	"member_service_frame/controller"
	"member_service_frame/repo"
	"member_service_frame/repo/psql"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	pb "member_service_frame/grpc"
	r "member_service_frame/repo/redis"
)

var deployTime string = time.Now().UTC().Format(time.RFC3339)

const version string = "1.0.0"
const dbName string = "Member"
const redisPool int = 1

func main() {
	// setting middleWare
	var server *gin.Engine = config.GetEngineWithMiddleWare()
	var psqldb *sql.DB = repo.GetPSQLConnecter(dbName)
	var redisClient *redis.Client = repo.GetRedisConnecter(redisPool)

	var userRepo *psql.PsqlUserRepository = psql.NewUserRepository(psqldb)
	var loginTimeRepo *r.RedisLoginCheckRepository = r.NewLoginCheckRepository(redisClient)
	defer psqldb.Close()

	server.GET("/health_check", healthCheck)
	server.NoRoute(handleNoRoute)

	controller.MemberServiceGroup(server, userRepo, loginTimeRepo)
	controller.CreateTable(server, userRepo, loginTimeRepo)

	// gRPC service
	go openGRPCService(config.Setting.GetMemberServiceGRPCPort(), loginTimeRepo)

	server.Run(config.Setting.GetMemberServiceRESTfulPort())
}

func healthCheck(ctx *gin.Context) {
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

func openGRPCService(grpcPort string, loginTimeRepo *r.RedisLoginCheckRepository) {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizerServer(grpcServer, &controller.Server{RedisClient: loginTimeRepo})
	log.Printf("gRPC service is running on port %s \n", grpcPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
