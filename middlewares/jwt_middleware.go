package middlewares

import (
	"fmt"
	"hacktiv8_fp_1/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(svc service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rawToken := ctx.GetHeader("Authorization");
		if len(rawToken) == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized);
			return
		}
		fmt.Println(rawToken);
		tokenString := rawToken[len("Bearer "):]// ada spasinya njir
		fmt.Println("str token -> ", tokenString);
		token, err := svc.ValidateToken(tokenString);

		if err != nil {
			log.Println("errersd - ", err);
			ctx.AbortWithStatus(http.StatusUnauthorized);
			return
		}
		if !token.Valid {
			log.Println("token brow-> ", token);
			ctx.AbortWithStatus(http.StatusUnauthorized);
			return
		}

		claims := token.Claims.(jwt.MapClaims);
		log.Println("claimssms ", claims)
		ctx.Set("credential", claims);
		ctx.Next()
	}
}