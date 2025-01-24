package api

import (
	"log"
	"net/http"
	"time"

	"github.com/PhuPhuoc/curanest-auth-service/builder"
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/PhuPhuoc/curanest-auth-service/config"
	"github.com/PhuPhuoc/curanest-auth-service/docs"
	"github.com/PhuPhuoc/curanest-auth-service/middleware"
	accounthttpservice "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/httpservice"
	accountrpcservice "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/rpcservice"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	rolehttpservice "github.com/PhuPhuoc/curanest-auth-service/module/role/infars/httpservice"
	rolequeries "github.com/PhuPhuoc/curanest-auth-service/module/role/usecase/queries"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type server struct {
	port string
	db   *sqlx.DB
}

func InitServer(port string, db *sqlx.DB) *server {
	return &server{
		port: port,
		db:   db,
	}
}

const (
	env_local = "local"
	env_vps   = "vps"
)

// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	envDevlopment := config.AppConfig.EnvDev
	if envDevlopment == env_local {
		// gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/"
	}

	if envDevlopment == env_vps {
		gin.SetMode(gin.ReleaseMode)
		docs.SwaggerInfo.BasePath = "/auth"
	}

	router := gin.New()

	configcors := cors.DefaultConfig()
	configcors.AllowAllOrigins = true
	configcors.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	configcors.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	configcors.ExposeHeaders = []string{"Content-Length"}
	configcors.AllowCredentials = true
	configcors.MaxAge = 12 * time.Hour

	router.Use(cors.New(configcors))
	router.Use(middleware.SkipSwaggerLog(), gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "curanest-auth-service - pong"}) })

	tokenProvider := common.NewJWTx(config.AppConfig.Key, 65*60*24*7)
	role_query_builder := rolequeries.NewRoleQueryWithBuilder(builder.NewRoleBuilder(sv.db))
	acc_query_builder := accountqueries.NewAccountQueryWithBuilder(builder.NewAccountBuilder(sv.db).AddTokenProvider(tokenProvider))
	acc_cmd_builder := accountcommands.NewAccountCmdWithBuilder(builder.NewAccountBuilder(sv.db))

	api := router.Group("/api/v1")
	{
		rolehttpservice.NewCategoryHTTPService(role_query_builder).Routes(api)
		accounthttpservice.NewAccountHTTPService(acc_query_builder).Routes(api)
	}

	rpc := router.Group("/external/rpc")
	{
		accountrpcservice.NewAccountRPCService(
			acc_cmd_builder,
			acc_query_builder,
		).AddAuth(tokenProvider).Routes(rpc)
	}

	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
