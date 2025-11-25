package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-project_junior/webhook/internal/repository"
	"go-project_junior/webhook/internal/repository/dao"
	"go-project_junior/webhook/internal/service"
	"go-project_junior/webhook/internal/web"
	"go-project_junior/webhook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDb()
	server := initWebServer()
	u := initUserHanl(db)
	u.RegisterRouters(server)
	server.Run(":8080")
}

func initDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webhook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://localhost:3000"},
		//表示运行 携带cookie 的内容
		AllowCredentials: true,

		AllowHeaders: []string{"Content-Type", "Authorization"},
		//AllowHeaders: []string{"content-type"},
		//AllowMethods: []string{"POST"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是我的 Middleware")
	})
	// 这个用法其实和redis 一样的 cookie 也可以改成redis
	//里面的具体存储的键值对 mysession 是一个浏览器的key
	// 会出现mysession 然后对应的value 在浏览量里
	stroe := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mysession", stroe))
	server.Use(middleware.NewLoginMiddlewareBuilder().CheckLogin())
	return server
}

func initUserHanl(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDao(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}
