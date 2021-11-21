package app

import (
	"erajaya/config"
	"fmt"

	"github.com/gin-gonic/gin"

	productHttp "erajaya/handler/product"
	productRepository "erajaya/repository/product"
	productUsecase "erajaya/usecase/product"
)

func NewApp() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// connection DB
	db := config.NewConnection()

	// open redis connection
	// redisClient := cache.NewClient()
	// result, err := cache.Ping(redisClient)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// repository
	productRepository := productRepository.NewProductRepository(db)

	// usecase
	productUsecase := productUsecase.NewProductUsecase(productRepository)

	// handler
	productHandler := productHttp.NewProductHandler(productUsecase)

	// router
	productHandler.Route(engine)

	fmt.Println("Running on port", config.ENV.AppPort)
	return engine
}
