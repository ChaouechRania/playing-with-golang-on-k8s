package server

import (
	"os"
	"playing-with-golang-on-k8s/auth"
	"playing-with-golang-on-k8s/routes"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Config server config
type Config struct {
	Port           string
	AuthMiddleware *jwt.GinJWTMiddleware
}

//NewConfig constructs a new config
func NewConfig(authMiddleware *jwt.GinJWTMiddleware) *Config {
	viper.AutomaticEnv()

	viper.BindEnv("port", "APP_PORT")
	viper.BindEnv("dbHost", "DB_HOST")
	viper.BindEnv("dbPort", "DB_PORT")
	viper.BindEnv("dbUser", "DB_USERNAME")
	viper.BindEnv("dbPassword", "DB_PASSWORD")
	viper.BindEnv("dbName", "DB_NAME")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:           port,
		AuthMiddleware: authMiddleware,
	}
}

//Server is the http layer of the app
type Server struct {
	Config            *Config
	UserActions       *routes.UserActions
	ProdsActions      *routes.ProductActions
	PermissionService *auth.PermissionService
	//SearchActions     *routes.SearchActions
}

//Run setup the app with all dependencies
func (s *Server) Run() error {
	r := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AddAllowHeaders("Authorization")
	corsCfg.AllowAllOrigins = true
	r.Use(cors.New(corsCfg))

	r.GET("/", routes.GetHome)

	r.GET("/status", routes.GetStatus)

	//r.POST("/admin/indexations", s.IndexActions.IndexAll)

	auth := r.Group("/api/auth")
	{
		auth.POST("", s.Config.AuthMiddleware.LoginHandler)
		auth.GET("/refresh_token", s.Config.AuthMiddleware.RefreshHandler)
	}

	auth.Use(s.Config.AuthMiddleware.MiddlewareFunc())

	users := r.Group("/api/users")
	{
		users.GET("", s.Config.AuthMiddleware.MiddlewareFunc(), s.UserActions.GetUsers)
		users.POST("", s.UserActions.CreateUser)
		users.GET("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.UserActions.GetUser)
		users.PATCH("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.UserActions.PatchUser)
		users.DELETE("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.UserActions.DeleteUser)

	}
	//IsOrgOwnerOrAdmin
	prods := r.Group("/api/products")
	{
		prods.POST("", s.Config.AuthMiddleware.MiddlewareFunc(), s.ProdsActions.Create)
		prods.GET("", s.Config.AuthMiddleware.MiddlewareFunc(), s.ProdsActions.List)
		prods.GET("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.ProdsActions.Get)
		prods.DELETE("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.ProdsActions.Delete)
		prods.PUT("/:id", s.Config.AuthMiddleware.MiddlewareFunc(), s.ProdsActions.Update)
	}
	/*refs := r.Group("/api/refs")
	{
		refs.GET("", routes.GetRefs)
	}*/

	return r.Run(":" + s.Config.Port)
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get("id")
	c.JSON(200, gin.H{
		"userID":   claims["ID"],
		"userName": user,
		"claims":   claims,
	})
}
