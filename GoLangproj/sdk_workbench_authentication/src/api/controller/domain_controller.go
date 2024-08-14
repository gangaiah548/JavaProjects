package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sdk_workbench_authentication/src/config/logger"
	"sdk_workbench_authentication/src/models"
	"sdk_workbench_authentication/src/services"

	"github.com/gin-gonic/gin"
)

type InpCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
	
}
type IDomainController interface {
	CreateDomain(c *gin.Context)
	GetDomain(c *gin.Context)
	GetRoleData(c *gin.Context)
	DeleteDomain(c *gin.Context)
	UpdateDomain(c *gin.Context)
	AuthDomain(c *gin.Context)
	LogoutAuthDomain(c *gin.Context)
	GetUserInfo(c *gin.Context)
	ChangePwd(c *gin.Context)
	ForgotPwd(c *gin.Context)
	CreateRoleData(c *gin.Context)
}
type DomainControllerImpl struct {
	domainService services.IDomainService
}

// LogoutAuthDomain implements IDomainController.
func (*DomainControllerImpl) LogoutAuthDomain(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/auth/logout"

	// Create an HTTP client
	client := &http.Client{}
	// Create a request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Header.Add("Authorization", tokenString)
	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		c.JSON(http.StatusOK, "loged out successfully")
		return
	} else {
		fmt.Println("API request failed with status code:", resp.Status)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please give proper credentials"})
		return
	}

}

func (pdc *DomainControllerImpl) GetUserInfo(c *gin.Context) {
	//middleware.authenticationMiddleware
	//authenticationMiddleware(c)
	//requiredPermissions := secure_App.roleMatrix[""]
	tokenString := c.GetHeader("Authorization")
	var ti []interface{}
	var sc int
	ti, sc = GetUserDetails(tokenString, c)
	var userErrort models.UserError
	var userSuccesst models.UserSuccess
	for _, val := range ti {
		switch v := val.(type) {
		case models.UserError:
			userErrort = v
			fmt.Println("Integer value:", userErrort)
			c.JSON(sc, userErrort)
			return
		case models.UserSuccess:
			userSuccesst = v
			fmt.Println("String value:", userSuccesst)
			c.JSON(sc, userSuccesst)
			return
		}
	}

}

func GetUserDetails(tokenString string, c *gin.Context) (ti []interface{}, sc int) {
	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/user/info"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		de := models.DetailError{
			ErrorCode:        "formingNewequest",
			ErrorDescription: err.Error(),
		}
		var det []models.DetailError
		det = append(det, de)
		tur := models.UserError{
			TraceId: "",
			Type:    "StatusBadRequest",
			Title:   "Post UAM ueerinfo call",
			Status:  "FAILED",
			Detail:  det,
		}
		ti = append(ti, tur)
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return ti, http.StatusBadRequest
	}
	req.Header.Add("Authorization", tokenString)

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		// Handle the error
		de := models.DetailError{
			ErrorCode:        "formingNewequest",
			ErrorDescription: err.Error(),
		}
		var det []models.DetailError
		det = append(det, de)
		tur := models.UserError{
			TraceId: "",
			Type:    "StatusServiceUnavailable",
			Title:   "Post UAM ueerinfo call",
			Status:  "FAILED",
			Detail:  det,
		}
		ti = append(ti, tur)
		//c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Post UAM ueerinfo call " + err.Error()})

		return ti, http.StatusServiceUnavailable

	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		// responseBody, _ := ioutil.ReadAll(resp.Body)
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err.Error())
			de := models.DetailError{
				ErrorCode:        "Reading data in success case",
				ErrorDescription: err.Error(),
			}
			var det []models.DetailError
			det = append(det, de)
			tur := models.UserError{
				TraceId: "",
				Type:    "InternalError",
				Title:   "Post UAM ueerinfo call",
				Status:  "FAILED",
				Detail:  det,
			}
			ti = append(ti, tur)
			//c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Post UAM ueerinfo call " + err.Error()})

			return ti, 500

		}

		// Define a struct to unmarshal the JSON into
		var userSuccessresp models.UserSuccess

		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userSuccessresp); err != nil {
			var userFailuresresp models.UserError

			if err := json.Unmarshal(responseBody, &userFailuresresp); err != nil {
				de := models.DetailError{
					ErrorCode:        "Unmarshaling data case",
					ErrorDescription: err.Error(),
				}
				var det []models.DetailError
				det = append(det, de)
				tur := models.UserError{
					TraceId: "",
					Type:    "InternalError",
					Title:   "unmarshalling error during userFailuresresp return  /Getuserdetails",
					Status:  "FAILED",
					Detail:  det,
				}
				ti = append(ti, tur)
				return ti, 500
			}
			ti = append(ti, userFailuresresp)
			//c.JSON(http.StatusOK, )
			return ti, http.StatusOK

		}

		fmt.Println("response body", userSuccessresp)

		ti = append(ti, userSuccessresp)
		//c.JSON(http.StatusOK, )
		return ti, http.StatusOK
	} else {
		fmt.Println("API request failed with status code:", response.Status)
		de := models.DetailError{
			ErrorCode:        "Unmarshaling data case",
			ErrorDescription: err.Error(),
		}
		var det []models.DetailError
		det = append(det, de)
		tur := models.UserError{
			TraceId: "",
			Type:    "InternalError",
			Title:   "unmarshalling error during userFailuresresp return  /Getuserdetails",
			Status:  "FAILED",
			Detail:  det,
		}
		ti = append(ti, tur)
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "Please give proper token details for getting user details"})
		return ti, http.StatusUnauthorized
	}

}

func (pdc *DomainControllerImpl) ForgotPwd(c *gin.Context) {
	//middleware.authenticationMiddleware
	//authenticationMiddleware(c)
	var input2 models.ForgotPassword
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/user/forgotpassword"
	requestData := models.ChangePassword{
		Username: input2.Username,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a request object
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the content type and add the bearer token to the request headers
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer YOUR_BEARER_TOKEN")

	// Send the request
	resp, err := client.Do(req)

	//resp, err := client.Get(url)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK && resp.Body != nil {
		// responseBody, _ := ioutil.ReadAll(resp.Body)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Define a struct to unmarshal the JSON into
		var userSuccessresp models.UserChangePwd
		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userSuccessresp); err != nil {
			var userFailuresresp models.UserError
			if err := json.Unmarshal(responseBody, &userFailuresresp); err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "unmarshalling error during login return Jwt code in Forgot pwd"})
				return
			}
			c.JSON(http.StatusOK, userFailuresresp)
			return

		}
		fmt.Println("response body", userSuccessresp)
		//detailsds := userSuccessresp.Detail
		//c.JSON(http.StatusOK, tokenResponse)
		//fmt.Println("tokenResponse at 216:Bearer " + detailsds.Token)
		fmt.Println("response body", userSuccessresp)
		c.JSON(http.StatusOK, userSuccessresp)
		//GetUserDetails("Bearer "+detailsds.Token, c)
		return
	} else {
		fmt.Println("API request failed with status code:", resp.Status)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			c.JSON(http.StatusOK, gin.H{"error": "during reading error reponse in login"})
			return
		}

		// Define a struct to unmarshal the JSON into
		var userErrorResp models.UserError

		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userErrorResp); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "unmarshalling error during login return Jwt code in forgotpwd"})
			return

		}
		return
	}

}

func (pdc *DomainControllerImpl) ChangePwd(c *gin.Context) {
	//middleware.authenticationMiddleware
	//authenticationMiddleware(c)
	var input2 models.ChangePassword
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/user/changepassword"
	requestData := models.ChangePassword{
		Username:    input2.Username,
		Password:    input2.Password,
		NewPassword: input2.NewPassword,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a request object
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the content type and add the bearer token to the request headers
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer YOUR_BEARER_TOKEN")

	// Send the request
	resp, err := client.Do(req)

	//resp, err := client.Get(url)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK && resp.Body != nil {
		// responseBody, _ := ioutil.ReadAll(resp.Body)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Define a struct to unmarshal the JSON into
		var userSuccessresp models.UserChangePwd
		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userSuccessresp); err != nil {
			var userFailuresresp models.UserError
			if err := json.Unmarshal(responseBody, &userFailuresresp); err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "310unmarshalling error during login return Jwt code in Changed pwd"})
				return
			}
			c.JSON(http.StatusOK, userFailuresresp)
			return

		}
		fmt.Println("response body", userSuccessresp)
		//detailsds := userSuccessresp.Detail
		//c.JSON(http.StatusOK, tokenResponse)
		//fmt.Println("tokenResponse at 216:Bearer " + detailsds.Token)
		fmt.Println("response body", userSuccessresp)
		c.JSON(http.StatusOK, userSuccessresp)
		//GetUserDetails("Bearer "+detailsds.Token, c)
		return
	} else {
		fmt.Println("API request failed with status code:", resp.Status)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			c.JSON(http.StatusOK, gin.H{"error": "during reading error reponse in login"})
			return
		}

		// Define a struct to unmarshal the JSON into
		var userErrorResp models.UserError

		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userErrorResp); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "337unmarshalling error during login return Jwt code in changepwd error"})
			return

		}
		return
	}

}

// AuthDomain implements IDomainController.
func (pdc *DomainControllerImpl) AuthDomain(c *gin.Context) {
	//middleware.authenticationMiddleware
	//authenticationMiddleware(c)
	var input2 InpCred
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/auth/login"
	requestData := InpCred{
		Username: input2.Username,
		Password: input2.Password,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a request object
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the content type and add the bearer token to the request headers
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer YOUR_BEARER_TOKEN")

	// Send the request
	resp, err := client.Do(req)

	//resp, err := client.Get(url)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK && resp.Body != nil {
		// responseBody, _ := ioutil.ReadAll(resp.Body)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Define a struct to unmarshal the JSON into
		var userSuccessresp models.UserSuccess
		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userSuccessresp); err != nil {
			var userFailuresresp models.UserError
			if err := json.Unmarshal(responseBody, &userFailuresresp); err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "unmarshalling error during login return Jwt code in Auth module"})
				return
			}
			c.JSON(http.StatusOK, userFailuresresp)
			return

		}
		fmt.Println("response body", userSuccessresp)
		detailsds := userSuccessresp.Detail
		//c.JSON(http.StatusOK, tokenResponse)
		fmt.Println("tokenResponse at 216:Bearer " + detailsds.Token + "Role " + detailsds.Role.RoleName)
		fmt.Println("response body", userSuccessresp)
		fmt.Println("tokenResponse at 216:Bearer " + detailsds.Token + "Role " + detailsds.Role.RoleName)
		//tokenString = tokenString[len("Bearer "):]
		accessurl := detailsds.Role.Menus[0].AccessUrl
		entitlements, count, err := pdc.domainService.GetAllRoleData(c.Writer, c.Request, detailsds.Token, detailsds.Role.RoleName, accessurl)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"err": err,
					"objectValue": map[string]interface{}{"arrayValue": "Error while fetching associate role permission matrix"}},
			)
			return
		}

		logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

		c.JSON(http.StatusOK,
			entitlements)
		return
	} else {
		fmt.Println("API request failed with status code:", resp.Status)
		responseBody, err := ioutil.ReadAll(resp.Body)
		if (err != nil) || responseBody == nil {
			c.JSON(http.StatusOK, gin.H{"error": "during reading error reponse in login"})
			return
		}

		// Define a struct to unmarshal the JSON into
		var userErrorResp models.UserError

		// Unmarshal the JSON data into the struct
		if err := json.Unmarshal(responseBody, &userErrorResp); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "unmarshalling error during login return Jwt code in changepwd"})
			return

		}
		return
	}

}

func (pdc *DomainControllerImpl) GetDomain(c *gin.Context) {
	active, count, err := pdc.domainService.GetDomains(c.Writer, c.Request)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"err": err,
				"objectValue": map[string]interface{}{"arrayValue": active}},
		)
		return
	}

	logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

	c.JSON(
		http.StatusOK,
		active,
	)
	//return
}
func (pdc *DomainControllerImpl) GetRoleData(c *gin.Context) {

	//tokenString[len("Bearer "):],
	tokenString := c.GetHeader("Authorization")
	var ti []interface{}
	var sc int
	ti, sc = GetUserDetails(tokenString, c)
	var userErrort models.UserError
	var userSuccesst models.UserSuccess
	for _, val := range ti {
		switch v := val.(type) {
		case models.UserError:
			userErrort = v
			fmt.Println("Integer value:", userErrort)
			c.JSON(sc, userErrort)
			return
		case models.UserSuccess:
			userSuccesst = v
			fmt.Println("String value:", userSuccesst)
			detailsds := userSuccesst.Detail
			//c.JSON(http.StatusOK, tokenResponse)
			fmt.Println("tokenResponse at 216:Bearer " + detailsds.Token + "Role " + detailsds.Role.RoleName)
			tokenString = tokenString[len("Bearer "):]
			entitlements, count, err := pdc.domainService.GetAllRoleData(c.Writer, c.Request, tokenString, detailsds.Role.RoleName, detailsds.Role.Menus[0].AccessUrl)

			if err != nil {
				c.JSON(
					http.StatusInternalServerError,
					map[string]interface{}{"err": err,
						"objectValue": map[string]interface{}{"arrayValue": "Error while fetching associate role permission matrix"}},
				)
				return
			}

			logger.Info().Msgf("There are %s active processes", fmt.Sprint(count))

			c.JSON(http.StatusOK,
				entitlements)
			return
		}
	}

}
func (pdc *DomainControllerImpl) DeleteDomain(c *gin.Context) {

	// Create book
	key := c.Query("key")
	dproduct := models.Domain{}

	domainresponse, err := pdc.domainService.DeleteDomain(c.Writer, c.Request, key, dproduct)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"Acknowledgment": err.Error()},
		)
		return
	}

	// returning empty string since process definition as return is not required as it will unnecessarily increase the return payload size
	//deploymentResponse.Definition = ""
	c.JSON(
		http.StatusOK,
		domainresponse)
	return
}

func (pdc *DomainControllerImpl) UpdateDomain(c *gin.Context) {
	key := c.Query("key")
	if key == "" || len(key) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"key is not empty or proper": key})
		return
	}
	var input2 models.Domain
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ddomain := models.Domain{
		Uuid:        input2.Uuid,
		Domain:      input2.Domain,
		CreatedBy:   input2.CreatedBy,
		Date:        input2.Date,
		Version:     input2.Version,
		Description: input2.Description,
		SubDomains:  input2.SubDomains,
	}
	domainresponse, err := pdc.domainService.UpdateDomain(key, ddomain)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"description": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		domainresponse)
	return
}

// CreateDomain implements IProductController.
func (pdc *DomainControllerImpl) CreateDomain(c *gin.Context) {
	//var input models.JavaClass
	var input2 models.Domain
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ddomain := models.Domain{
		Domain:      input2.Domain,
		CreatedBy:   input2.CreatedBy,
		Date:        input2.Date,
		Version:     input2.Version,
		Description: input2.Description,
		SubDomains:  input2.SubDomains,
	}

	productResponse, err := pdc.domainService.CreateDomain(c.Writer, c.Request, ddomain)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		productResponse)
	return
}

func (pdc *DomainControllerImpl) CreateRoleData(c *gin.Context) {
	//var input models.JavaClass
	var input2 models.Entitlement
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ddomain := models.Entitlement{
		Entitlement: input2.Entitlement,
		Component:   input2.Component,
	}

	productResponse, err := pdc.domainService.CreateRoleData(c.Writer, c.Request, ddomain)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{"data": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		productResponse)
	return
}

func NewDomaincreation() (IDomainController, error) {
	var DomainDBconnectionService, vrr = services.NewDomaindepService()

	if vrr != nil {
		return nil, vrr
	}
	return &DomainControllerImpl{
		domainService: DomainDBconnectionService,
	}, nil

}

func authenticationMiddleware(c *gin.Context) {
	// Implement your authentication logic here

	//

}
