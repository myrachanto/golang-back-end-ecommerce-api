package routes

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/myrachanto/asokomonolith/controllers"
)

func ApiMicroservice() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}
	PORT := os.Getenv("PORT")
	key := os.Getenv("EncryptionKey")
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) 
	e.Use(middleware.CORS())

	JWTgroup := e.Group("/api/")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(key),
	}))
	// admin := e.Group("admin/")
	// admin.Use(isAdmin)

	// var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningMethod: "HS256",
	// 	SigningKey: []byte(key),
	// })
	//JwtG := e.Group("/users")
	// JwtG.Use(middleware.JWT([]byte(key)))
	// Routes
	//e.GET("/is-loggedin", h.private, IsLoggedIn)
	//e.POST("/register", IsLoggedIn,isAdmin,isEmployee,isSupervisor, controllers.UserController.Create)
	// Routes
	////////////////////////////////////////////////////////
	/////////////users//////////////////////////////////
	e.POST("/register", controllers.UserController.Create)
	e.POST("/login", controllers.UserController.Login)
	e.GET("/logout", controllers.UserController.Logout)
	e.GET("/users", controllers.UserController.GetAll)
	e.GET("/users/:id", controllers.UserController.GetOne)
	e.PUT("/users/:id", controllers.UserController.Update)
	e.PUT("/usersA/:id", controllers.UserController.AUpdate)
	e.DELETE("/users/:id", controllers.UserController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/categorys", controllers.CategoryController.Create)
	e.GET("/categorys", controllers.CategoryController.GetAll)
	e.GET("/categorys/:id", controllers.CategoryController.GetOne)
	e.PUT("/categorys/:id", controllers.CategoryController.Update)
	e.DELETE("/categorys/:id", controllers.CategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/countys", controllers.CountyController.Create)
	e.GET("/countys", controllers.CountyController.GetAll)
	e.GET("/countys/:id", controllers.CountyController.GetOne)
	e.PUT("/countys/:id", controllers.CountyController.Update)
	e.DELETE("/countys/:id", controllers.CountyController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/customers", controllers.CustomerController.Create)
	e.GET("/customers", controllers.CustomerController.GetAll)
	e.GET("/customers/:id", controllers.CustomerController.GetOne)
	e.PUT("/customers/:id", controllers.CustomerController.Update)
	e.DELETE("/customers/:id", controllers.CustomerController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/divisions", controllers.DivisionController.Create)
	e.GET("/divisions", controllers.DivisionController.GetAll)
	e.GET("/divisions/:id", controllers.DivisionController.GetOne)
	e.PUT("/divisions/:id", controllers.DivisionController.Update)
	e.DELETE("/divisions/:id", controllers.DivisionController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/industrys", controllers.IndustryController.Create)
	e.GET("/industrys", controllers.IndustryController.GetAll)
	e.GET("/industrys/:id", controllers.IndustryController.GetOne)
	e.PUT("/industrys/:id", controllers.IndustryController.Update)
	e.DELETE("/industrys/:id", controllers.IndustryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/majorcategorys", controllers.MajorcategoryController.Create)
	e.GET("/majorcategorys", controllers.MajorcategoryController.GetAll)
	e.GET("/majorcategorys/:id", controllers.MajorcategoryController.GetOne)
	e.PUT("/majorcategorys/:id", controllers.MajorcategoryController.Update)
	e.DELETE("/majorcategorys/:id", controllers.MajorcategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/streets", controllers.StreetController.Create)
	e.GET("/streets", controllers.StreetController.GetAll)
	e.GET("/streets/:id", controllers.StreetController.GetOne)
	e.PUT("/streets/:id", controllers.StreetController.Update)
	e.DELETE("/streets/:id", controllers.StreetController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/subcategorys", controllers.SubcategoryController.Create)
	e.GET("/subcategorys", controllers.SubcategoryController.GetAll)
	e.GET("/subcategorys/:id", controllers.SubcategoryController.GetOne)
	e.PUT("/subcategorys/:id", controllers.SubcategoryController.Update)
	e.DELETE("/subcategorys/:id", controllers.SubcategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////category//////////////////////////////////
	e.POST("/categorys", controllers.CategoryController.Create)
	e.GET("/categorys", controllers.CategoryController.GetAll)
	e.GET("/categorys/:id", controllers.CategoryController.GetOne)
	e.PUT("/categorys/:id", controllers.CategoryController.Update)
	e.DELETE("/categorys/:id", controllers.CategoryController.Delete)
	////////////////////////////////////////////////////////
	/////////////town//////////////////////////////////
	e.POST("/towns", controllers.TownController.Create)
	e.GET("/towns", controllers.TownController.GetAll)
	e.GET("/towns/:id", controllers.TownController.GetOne)
	e.PUT("/towns/:id", controllers.TownController.Update)
	e.DELETE("/towns/:id", controllers.TownController.Delete)
	////////////////////////////////////////////////////////
	/////////////Shop//////////////////////////////////
	e.POST("/shops", controllers.ShopController.Create)
	e.GET("/shops", controllers.ShopController.GetAll)
	e.GET("/shops/:id", controllers.ShopController.GetOne)
	e.PUT("/shops/:id", controllers.ShopController.Update)
	e.DELETE("/shops/:id", controllers.ShopController.Delete)
	////////////////////////////////////////////////////////
	/////////////products//////////////////////////////////
	e.POST("/products", controllers.ProductController.Create)
	e.GET("/products", controllers.ProductController.GetAll)
	e.GET("/products/:id", controllers.ProductController.GetOne)
	e.PUT("/products/:id", controllers.ProductController.Update)
	e.DELETE("/products/:id", controllers.ProductController.Delete)
	////////////////////////////////////////////////////////
	/////////////tags//////////////////////////////////
	e.POST("/tags", controllers.TagController.Create)
	e.GET("/tags", controllers.TagController.GetAll)
	e.GET("/tags/:id", controllers.TagController.GetOne)
	e.PUT("/tags/:id", controllers.TagController.Update)
	e.DELETE("/tags/:id", controllers.TagController.Delete)
	////////////////////////////////////////////////////////
	/////////////ratings//////////////////////////////////
	e.POST("/ratings", controllers.RatingController.Create)
	e.GET("/ratings", controllers.RatingController.GetAll)
	e.GET("/ratings/:id", controllers.RatingController.GetOne)
	e.PUT("/ratings/:id", controllers.RatingController.Update)
	e.DELETE("/ratings/:id", controllers.RatingController.Delete)
	////////////////////////////////////////////////////////
	/////////////nortificatrions//////////////////////////////////
	e.POST("/nortificatrions", controllers.NortificationController.Create)
	e.GET("/nortificatrions", controllers.NortificationController.GetAll)
	e.GET("/nortificatrions/:id", controllers.NortificationController.GetOne)
	e.PUT("/nortificatrions/:id", controllers.NortificationController.Update)
	e.DELETE("/nortificatrions/:id", controllers.NortificationController.Delete)
	////////////////////////////////////////////////////////
	/////////////verify//////////////////////////////////
	e.POST("/verify", controllers.VerifyController.Create)
	e.GET("/verify", controllers.VerifyController.GetAll)
	e.GET("/verify/:id", controllers.VerifyController.GetOne)
	e.PUT("/verify/:id", controllers.VerifyController.Update)
	e.DELETE("/verify/:id", controllers.VerifyController.Delete)
	////////////////////////////////////////////////////////
	/////////////invoice//////////////////////////////////
	e.POST("/invoice", controllers.InvoiceController.Create)
	e.GET("/invoice", controllers.InvoiceController.GetAll)
	e.GET("/invoice/:id", controllers.InvoiceController.GetOne)
	// e.PUT("/invoice/:id", controllers.InvoiceController.Update)
	// e.DELETE("/invoice/:id", controllers.InvoiceController.Delete)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}
