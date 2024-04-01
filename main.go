package main

import (
	"database/sql"
	"member_service_frame/config"
	"member_service_frame/controller"
	"member_service_frame/repo"
	"member_service_frame/repo/psql"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	r "member_service_frame/repo/redis"
)

var deployTime string = time.Now().UTC().Format(time.RFC3339)

const version string = "1.0.0"
const dbName string = "Member"
const redisPool int = 0

func main() {
	// setting middleWare
	var server *gin.Engine = config.GetEngineWithMiddleWare()
	var psqldb *sql.DB = repo.GetPSQLConnecter(dbName)
	var reddisdb *redis.Client = repo.GetRedisConnecter(redisPool)

	var userRepo *psql.PsqlUserRepository = psql.NewUserRepository(psqldb)
	var loginTimeRepo *r.RedisLoginCheckRepository = r.NewLoginCheckRepository(reddisdb)
	defer psqldb.Close()

	server.GET("/health_check", healthCheck)
	server.NoRoute(handleNoRoute)

	controller.MemberServiceGroup(server, userRepo, loginTimeRepo)
	controller.CreateTable(server, userRepo)

	server.Run(config.Setting.GetMemberServiceGRPCPort())
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

// func openGRPCService(grpcPort string) {
// var redisClient *redis.Client = r.Connect(redisPool)
// var loginRepo repository.LoginTimeInterface = redisLogincheck.NewRedisLoginCheckerRepository(redisClient)

// lis, err := net.Listen("tcp", grpcPort)
// if err != nil {
// 	logger.HandleSevereCrashed(err)
// 	return
// }
// s := grpc.NewServer()
// pb.RegisterAuthorizerServer(s, &controller.Server{RedisClient: loginRepo})
// fmt.Printf("gRPC Server is listening on port %v \n", grpcPort)
// if err := s.Serve(lis); err != nil {
// 	fmt.Printf("Failed to serve: %v\n", err)
// }
// }
