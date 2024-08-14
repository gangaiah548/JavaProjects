package routes

import (
	"net/http"
	"time"

	"sdk_workbench_authentication/src/api/controller"
	"sdk_workbench_authentication/src/api/middleware"
	"sdk_workbench_authentication/src/config/logger"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func SmpdataRouter(timeout time.Duration, group *gin.RouterGroup) error {

	// use timeout if required at the controller level
	logger.Info().Msg("Before Db in Product router")
	router := gin.Default()
	router.Use(CORS)
	var createSmpDataControllerObject, err = controller.NewDomaincreation()

	if err != nil {
		return err
	}
	logger.Info().Msg("Before calling POSt router")
	router.POST("/login", createSmpDataControllerObject.AuthDomain)
	router.POST("/create", createSmpDataControllerObject.CreateDomain)
	router.POST("/changepwd", createSmpDataControllerObject.ChangePwd)
	router.POST("/forgotpwd", createSmpDataControllerObject.ForgotPwd)
	router.GET("/getAll", createSmpDataControllerObject.GetDomain)
	router.GET("/getUserInfo", createSmpDataControllerObject.GetUserInfo)
	router.GET("/logout", createSmpDataControllerObject.LogoutAuthDomain)
	router.PUT("/update", createSmpDataControllerObject.UpdateDomain)
	router.DELETE("/delete", createSmpDataControllerObject.DeleteDomain)
	router.POST("/createroledata", createSmpDataControllerObject.CreateRoleData)
	//router.GET("/getAllRoleData", createSmpDataControllerObject.GetRoleData)
	authorized := router.Group("/api")
	authorized.Use(middleware.JWTMiddleware(&gin.Context{})) // Middleware applies only to /api routes
	//authorized.Use(middleware.RoleBasedAuthorization("read"))
	//authorized.Any()
	authorized.GET("/private", func(c *gin.Context) {
		//middleware.JWTMiddleware(); here sample request authorized.POST("private/create", createSmpDataControllerObject.CreateDomain)
		c.JSON(http.StatusOK, gin.H{"message": "Private endpoint"})
	})
	authorized.POST("/private/create", createSmpDataControllerObject.CreateDomain)
	authorized.GET("/private/getAllRoleData", createSmpDataControllerObject.GetRoleData)
	/* r.GET("/admin", authAndAuthorizeMiddleware("admin"), func(c *gin.Context) { //we can cross verify role also aacording to resource access
	    c.JSON(http.StatusOK, gin.H{"message": "Admin route"})
	})*/
	router.Run("localhost:8080")
	return nil
}

/*func ProductRouter(timeout time.Duration, group *gin.RouterGroup) error {

	// use timeout if required at the controller level
	logger.Info().Msg("Before Db in Product router")
	router := gin.Default()
	var createProductControllerObject, err = controller.Newcreation()

	if err != nil {
		return err
	}
	logger.Info().Msg("Before calling POSt router")
	router.POST("/create", createProductControllerObject.Create)
	router.PUT("/update", createProductControllerObject.UpdateProduct)
	router.GET("/getAll", createProductControllerObject.GetProducts)
	router.DELETE("/delete", createProductControllerObject.DeleteProduct)
	//group.POST("/create", createProductControllerObject.Create)
	//router.Run("localhost:8080")
	return nil
}*/

/*func SmpdataRouter(timeout time.Duration, group *gin.RouterGroup) error {

	// use timeout if required at the controller level
	logger.Info().Msg("Before Db in Product router")
	router := gin.Default()
	var createSmpDataControllerObject, err = controller.Newsmpcreation()

	if err != nil {
		return err
	}
	logger.Info().Msg("Before calling POSt router")
	router.POST("/createsmpdata", createSmpDataControllerObject.Createsmpdata)
	router.DELETE("/deletesmpdata", createSmpDataControllerObject.DeleteSmpdata)
	router.GET("/getsmpdata", createSmpDataControllerObject.GetSmpdata)
	router.PUT("/updatesmpdata", createSmpDataControllerObject.UpdateSmpData)
	router.Run("localhost:8080")
	return nil
}*/
