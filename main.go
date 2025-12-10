package main
import(
	"github.com/B-Meghana-Reddy/ecommerce/controllers"
	"github.com/B-Meghana-Reddy/ecommerce/database"
	"github.com/B-Meghana-Reddy/ecommerce/middleware"
	"github.com/B-Meghana-Reddy/ecommerce/routes"

	"github.com/gin-gonic/gin"

	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(
		database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users"),
	)

	router := gin.New()
	router.Use(gin.Logger())

	// 👇 PUBLIC ROUTES (NO TOKEN NEEDED)
	routes.UserRoutes(router) // /users/signup, /users/login, /users/productview, /users/search

	// 👇 PROTECTED ROUTES (TOKEN REQUIRED)
	auth := router.Group("/")
	auth.Use(middleware.Authentication())

	auth.GET("/addtocart", app.AddToCart())
	auth.GET("/removeitem", app.RemoveItem())
	auth.GET("/listcart", controllers.GetItemFromCart())
	auth.POST("/addaddress", controllers.AddAddress())
	auth.PUT("/edithomeaddress", controllers.EditHomeAddress())
	auth.PUT("/editworkaddress", controllers.EditWorkAddress())
	auth.GET("/deleteaddresses", controllers.DeleteAddress())
	auth.GET("/cartcheckout", app.BuyFromCart())
	auth.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
