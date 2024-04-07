package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/iamsad5566/member_service_frame/config"
	"github.com/iamsad5566/member_service_frame/controller"
	_ "github.com/iamsad5566/member_service_frame/docs"
	"github.com/iamsad5566/member_service_frame/repo"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	pb "github.com/iamsad5566/member_service_frame/grpc"
	"github.com/iamsad5566/member_service_frame/repo/psql"
	r "github.com/iamsad5566/member_service_frame/repo/redis"
)

var deployTime string = time.Now().UTC().Format(time.RFC3339)

const version string = "1.0.1"
const dbName string = "Member"
const redisPool int = 1

// @title           Member Service API
// @version         1.0.1
// @description     This is a RESTful API service for member service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
	controller.OAuth2Group(server, userRepo, loginTimeRepo)
	controller.CreateTable(server, userRepo, loginTimeRepo)

	// gRPC service
	go openGRPCService(config.Setting.GetMemberServiceGRPCPort(), loginTimeRepo)

	// Swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(config.Setting.GetMemberServiceRESTfulPort())
}

// healthCheck is a function to check the health of the service
func healthCheck(ctx *gin.Context) {
	hostName, _ := os.Hostname()
	ctx.JSON(http.StatusOK, gin.H{
		"Deployment Time": deployTime,
		"Version":         version,
		"Host":            hostName,
	})
}

// handleNoRoute is a function to handle the no route situation
func handleNoRoute(ctx *gin.Context) {
	ctx.File("no_route.html")
}

// openGRPCService is a function to open the gRPC service
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
