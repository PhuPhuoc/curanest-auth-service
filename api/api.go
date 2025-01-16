package api

import (
	"log"
	"net/http"

	"github.com/PhuPhuoc/curanest-auth-service/builder"
	"github.com/PhuPhuoc/curanest-auth-service/common"
	"github.com/PhuPhuoc/curanest-auth-service/docs"
	"github.com/PhuPhuoc/curanest-auth-service/middleware"
	accounthttpservice "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/httpservice"
	accountrpcservice "github.com/PhuPhuoc/curanest-auth-service/module/account/infars/rpcservice"
	accountcommands "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/commands"
	accountqueries "github.com/PhuPhuoc/curanest-auth-service/module/account/usecase/queries"
	rolehttpservice "github.com/PhuPhuoc/curanest-auth-service/module/role/infars/httpservice"
	rolequeries "github.com/PhuPhuoc/curanest-auth-service/module/role/usecase/queries"
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

// @Summary		ping server
// @Description	ping server
// @Tags			ping
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]any	"message success"
// @Failure		400	{object}	error			"Bad request error"
// @Router			/ping [get]
func (sv *server) RunApp() error {
	docs.SwaggerInfo.BasePath = "/"
	// gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middleware.SkipSwaggerLog(),
		gin.Recovery(),
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	/* ping - test */
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "curanest-auth-service - pong x777"})
	})

	tokenProvider := common.NewJWTx("hehe-haha-hihi-kkkk-huhu-hichic", 60*60*24*7)

	/*
	* usecase (commandes - queries)
	* */
	// role
	role_query_builder := rolequeries.NewRoleQueryWithBuilder(
		builder.NewRoleBuilder(sv.db),
	)

	// account
	acc_query_builder := accountqueries.NewAccountQueryWithBuilder(
		builder.NewAccountBuilder(sv.db).AddTokenProvider(tokenProvider),
	)
	acc_cmd_builder := accountcommands.NewAccountCmdWithBuilder(
		builder.NewAccountBuilder(sv.db),
	)

	// http vs rpc
	api := router.Group("/api/v1")
	{
		rolehttpservice.NewCategoryHTTPService(role_query_builder).Routes(api)
		accounthttpservice.NewAccountHTTPService(acc_query_builder).Routes(api)
	}

	rpc := router.Group("/internal/rpc")
	{
		accountrpcservice.NewAccountRPCService(
			acc_cmd_builder,
			acc_query_builder,
		).Routes(rpc)
	}

	log.Println("server start listening at port: ", sv.port)
	return router.Run(sv.port)
}
