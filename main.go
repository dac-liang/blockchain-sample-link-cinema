/*
Copyright 2020 LINE Corporation

LINE Corporation licenses this file to you under the Apache License,
version 2.0 (the "License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at:

  https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
License for the specific language governing permissions and limitations
under the License
*/
package main

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/cookie"
	"link/cinema/config"
	"link/cinema/controller"
	"link/cinema/docs"
	"strings"
)

var authMiddleware *jwt.GinJWTMiddleware

// @title Link Cinema API
// @version 0.1
// @description This is sample dapp to provide trials of LBD service
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v0
func main() {
	r := gin.Default()
	// TODO GetUserProfileFromSession に合わせて実装される想定の形跡があるけど、実装されなかったため、一旦は gin-jwt で進める
	// store := cookie.NewStore([]byte("secret"))
	// r.Use(sessions.Sessions("session", store))

	if configPath := os.Getenv(config.Path); configPath != "" {
		config.LoadAPIConfig(configPath)
	}
	// Google Cloud Run 環境変数上書き
	config.InitAPIConfig()

	host := config.GetAPIConfig().Endpoint
	if strings.HasPrefix(host, "http://") {
		host = host[7:]
	}
	if strings.HasPrefix(host, "https://") {
		host = host[8:]
	}
	docs.SwaggerInfo.Host = host

	// the jwt middleware
	r.Use(InitAuth())
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	auth := r.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	ctr := controller.NewController()

	v0 := r.Group("/api/v0")
	v0.Use(authMiddleware.MiddlewareFunc())
	{
		user := v0.Group("/user")
		{
			//user.GET("/login", ctr.LINELogin)
			//user.GET("/login/callback", ctr.LINELoginCallback)
			user.GET("/proxy", ctr.RequestProxy)
			user.GET("/proxy/commit/:proxyToken", ctr.CommitRequestProxy)
		}

		ticket := v0.Group("/ticket")
		{
			ticket.GET("/", ctr.GetPurchaseInfo)
			ticket.POST("/purchase", ctr.RequestTicketPurchasing)
			ticket.POST("/purchase/extra", ctr.RequestExtraPurchase)
			ticket.POST("/purchase/commit/:baseCoinTransferToken/:movieTokenTransferToken", ctr.CommitPurchasingTicket)
		}

		token := v0.Group("/token")
		{
			token.GET("/balance/base-coin", ctr.GetBaseCoinBalance)
			token.GET("/balance/movie-discount", ctr.GetMovieDiscountBalance)
			token.GET("/balance/movie-ticket", ctr.SearchTicketBalance)
			token.GET("/balance/movie", ctr.GetMovieTokenBalance)
		}
		test := v0.Group("/test")
		{
			test.GET("/init", ctr.InitUser)

			test.GET("/transaction", ctr.GetTransaction)
			test.GET("/config", ctr.ShowConfig)
		}
	}

	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", config.GetAPIConfig().Endpoint)) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}