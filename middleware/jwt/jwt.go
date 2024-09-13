package jwt

import (
	"FinancialAssistanceScheme/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

// UserClaims struct, used in JWT token generation
type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"` // For authorization
	jwt.StandardClaims
}

// GenerateJWT generate JWT token
func GenerateJWT(username, role string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	fmt.Println(username, role)

	claims := &UserClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey := []byte(config.GetConfig().JWT.SecretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// JWTAuthMiddleware - Authentication Middleware: Validates JWT token
func JWTAuthMiddleware(ctx iris.Context) {
	tokenString := ctx.GetHeader("Authorization")

	if tokenString == "" {
		log.Println("no token provided")
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "No token provided"})
		return
	}

	jwtKey := []byte(config.GetConfig().JWT.SecretKey)

	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		log.Println("invalid token err", err)
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "Invalid token"})
		return
	}

	ctx.Values().Set("username", claims.Username)
	ctx.Values().Set("role", claims.Role)
	ctx.Next()
}

//
//// Authorization Middleware: Check if the user has the required role
//func RoleAuthorizationMiddleware(role string) iris.Handler {
//	return func(ctx iris.Context) {
//		userRole := ctx.Values().GetString("role")
//		if userRole != role {
//			ctx.StatusCode(iris.StatusForbidden)
//			ctx.JSON(iris.Map{"error": "Access denied"})
//			return
//		}
//		ctx.Next()
//	}
//}
