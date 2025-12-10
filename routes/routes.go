package routes
// routes handles url paths and maps them to respective controllers
import(
	"github.com/B-Meghana-Reddy/ecommerce/controllers"
	"github.com/gin-gonic/gin"
	//gin is a web framework (like Flask for Python).
	//It:Creates servers, Handles requests, Manages routes like POST, GET, etc.
)

func UserRoutes(incomingRoutes *gin.Engine){
	// *gin.Engine = the main HTTP server router
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.SearchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery())
}	