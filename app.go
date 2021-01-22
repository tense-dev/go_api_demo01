package demo

import (
	"rest_api/beer"
	_ "rest_api/docs"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer() {
	r := gin.New()
	//r := gin.Default()

	// ===== Middlewares
	// r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// ===== Prometheus
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(r)
	// ===== Route of documents
	beer.NewRoutes(r)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":2000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
