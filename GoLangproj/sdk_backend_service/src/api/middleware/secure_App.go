package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sdk_backend_service/src/models"

	"github.com/gin-gonic/gin"
)

type InpCred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your authentication logic here
		valid := checkAuthentication(r)
		if !valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}*/

func authenticationMiddleware(c *gin.Context) {
	// Implement your authentication logic here

	var input2 InpCred
	if err := c.ShouldBindJSON(&input2); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	/*if r.Method == http.MethodPost {
		// Read the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// http.Error(w, "Error reading request body", http.StatusBadRequest)
			return false
		}
		defer r.Body.Close()

		// Do something with the request body (body is a byte slice)
		fmt.Printf("Received request body: %s\n", string(body))
	}*/

	url := "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/auth/login"
	// Create an instance of the struct and populate it with data

	requestData := InpCred{
		Username: input2.Username,
		Password: input2.Password,
	}

	// Convert the struct to JSON
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
	if resp.StatusCode == http.StatusOK {
		// Request was successful
		// You can read and process the response here
		// For example, reading the response body:
		// responseBody, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(responseBody))
		c.JSON(http.StatusOK, resp.Body)
		return
	} else {
		fmt.Println("API request failed with status code:", resp.Status)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please give proper credentials"})
	}
	//

}

func checkAuthentication(c *gin.Context) bool {
	//logger.Info().Msg(r.RequestURI)
	//client := &http.Client{}

	return true
}

func checkJWTAuthentication(r *http.Request) bool {
	return true
}

//var jwtSecret = []byte("8Zz5tw0Ionm3XPZZfN0NOml3z9FMfmpgXwovR9fp6ryDIoGRM8EPHAB6iHsc0fb")

func JWTMiddleware(c *gin.Context) gin.HandlerFunc {
	user := c.Query("token")
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tmpreferer := c.GetHeader("Referer")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		req, err := http.NewRequest("POST", "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/user/authenticate?name="+user, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		req.Header.Add("Authorization", tokenString)

		client := &http.Client{}
		response, err := client.Do(req)

		if err != nil {
			// Handle the error
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			c.Abort()
			return

		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			fmt.Println("API request failed with status code:", response.Status)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Please give proper credentials and parameters"})
			c.Abort()
			return
		}
		/*else{  //open code when UAM sending token
					body, err := ioutil.ReadAll(response.Body)
		    if err != nil {
		        // Handle the error
		        fmt.Println(err)
		        return
		    }
		    // Assuming the JWT token is sent in the response body as a string
		    jwtToken := string(body)
		    // Now you have the JWT token from the response
		    fmt.Printf("Received JWT token: %s\n", jwtToken)
				}*/

		// Token is valid, and you can access the claims
		//claims := token.Claims.(jwt.MapClaims)
		//c.Set("user", claims["user"])
		//claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

		if tmpreferer == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No referer key provided"})
			c.Abort()
			return
		}

		req, err = http.NewRequest("POST", "http://a83bc6c0ba96e4b2781d7d6f6f9c0f08-72df3990d92ee63b.elb.us-east-1.amazonaws.com/uam/v1/user/authorize", nil)
		if err != nil {
			fmt.Println("Error creating request for authorization:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		req.Header.Add("Authorization", tokenString)
		req.Header.Add("Referer", tmpreferer)
		client = &http.Client{}
		response, err = client.Do(req)

		if err != nil {
			// Handle the error
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			c.Abort()
			return

		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			fmt.Println("API request failed with status code:", response.Status)
			c.JSON(http.StatusForbidden, gin.H{"error": "Please give proper token and key parameters"})
			c.Abort()
			return
		}
		//signature := c.Param("8Zz5tw0Ionm3XPZZfN0NOml3z9FMfmpgXwovR9fp6ryDIoGRM8EPHAB6iHsc0fb")

		// Decode the Base64-encoded signature

		//var jwtSecret = []byte("8Zz5tw0Ionm3XPZZfN0NOml3z9FMfmpgXwovR9fp6ryDIoGRM8EPHAB6iHsc0fb")
		/*	tokenString = tokenString[len("Bearer "):]
			decodedSignature, err := base64.StdEncoding.DecodeString(tokenString)
			if err != nil {
				logger.Info().Msgf("error in decode string after %s", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			decodedKey := string(decodedSignature)

			fmt.Println("after  token to decode:", decodedKey)
			logger.Info().Msgf("token after %s", decodedKey)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method")
				}
				return decodedKey, nil
			})
			//logger
			fmt.Println("after parse token:", token)
			logger.Info().Msgf("token after %s", token.Raw)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token at parse" + err.Error()})
				c.Abort()
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				logger.Info().Msgf("error token.Valid  %s", err.Error())
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token at claims fetch"})
				c.Abort()
				return
			}

			userRole, ok := claims["sub"].(string)
			fmt.Println("API request role after role parse", userRole)
			if !ok || userRole != userRole {
				logger.Info().Msgf("error in userRole  %s", err.Error())
				c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
				c.Abort()
				return
			}*/

		c.Next()
	}
}

func RoleBasedAuthorization(entitlements ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := "user" //c.GetString("userRole") // Retrieve the user's role from the JWT or elsewhere
		requiredPermissions := models.RoleMatrix[userRole]

		// Check if the user has the required permissions
		for _, requiredPermission := range entitlements {
			if !contains(requiredPermissions, requiredPermission) {
				c.JSON(http.StatusForbidden, gin.H{"message": "Access denied at275"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func contains(slice []string, entitlement string) bool {
	for _, s := range slice {
		if s == entitlement {
			return true
		}
	}
	return false
}
