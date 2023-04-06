package routers

import (
	"giao/tour/blog/global"
	"giao/tour/blog/internal/middleware"
	"giao/tour/blog/internal/routers/api"
	v1 "giao/tour/blog/internal/routers/api/v1"
	"giao/tour/blog/pkg/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var methodLimiter = limiter.NewMethodLimiter().AddBucket(limiter.BucketRule{
	Key:          "/auth",
	FillInternal: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.ApiInfo())
	r.Use(middleware.RateLimiter(methodLimiter))
	r.Use(middleware.ContextTimeout(time.Second * global.AppSetting.DefaultContextTimeout))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", api.GetAuth)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles", article.Get)
	}

	return r
}
