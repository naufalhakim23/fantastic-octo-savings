package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"payday-savings/m/internal/app/adapter/repository"
	"payday-savings/m/internal/app/adapter/service"
	"payday-savings/m/internal/app/application/usecase"
	"payday-savings/m/internal/app/domain/valueobject"
)

var (
	bitbank             = service.Bitbank{}
	parameterRepository = repository.Parameter{}
	orderRepository     = repository.Order{}
)

// Controller is a controller
type Controller struct{}

// Router is routing settings
func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}
	// NOTICE: following path is from CURRENT directory, so please run Gin from root directory
	r.LoadHTMLGlob("internal/app/adapter/view/*")
	r.GET("/", ctrl.index)
	r.GET("/ticker", ctrl.ticker)
	r.GET("/candlestick", ctrl.candlestick)
	r.GET("/parameter", ctrl.parameter)
	r.GET("/order", ctrl.order)
	return r
}

func (ctrl Controller) index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello Goilerplate",
	})
}

func (ctrl Controller) ticker(c *gin.Context) {
	pair := valueobject.BtcJpy
	ticker := usecase.Ticker(bitbank, pair) // Dependency Injection
	c.JSON(200, ticker)
}

func (ctrl Controller) candlestick(c *gin.Context) {
	args := usecase.OhlcArgs{
		E: bitbank, // Dependency Injection
		P: valueobject.BtcJpy,
		T: valueobject.OneMin,
	}
	candlestick := usecase.Ohlc(args)
	c.JSON(200, candlestick)
}

func (ctrl Controller) parameter(c *gin.Context) {
	parameter := usecase.Parameter(parameterRepository) // Dependency Injection
	c.JSON(200, parameter)
}

func (ctrl Controller) order(c *gin.Context) {
	order := usecase.AddNewCardAndEatCheese(orderRepository) // Dependency Injection
	c.JSON(200, order)
}
