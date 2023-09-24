package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sknutsen/Zdk/internal/config"
)

type JWTMiddleware struct {
	Config *config.Config
}

func NewJWTMiddleware(config *config.Config) *JWTMiddleware {
	return &JWTMiddleware{Config: config}
}

func (jwt *JWTMiddleware) GetAuthMiddleware() fiber.Handler {
	return adaptor.HTTPMiddleware(EnsureValidToken)
}

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func EnsureValidToken(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return middleware.CheckJWT(next)
}

// NOTE: we are copying the HTTPMiddleware adaptor because the request context doesn't get propegated to fiber context
func HTTPMiddleware(c *fiber.Ctx) error {
	// var next bool
	// nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	next = true
	// 	// Convert again in case request may modify by middleware
	// 	c.Request().Header.SetMethod(r.Method)
	// 	c.Request().SetRequestURI(r.RequestURI)
	// 	c.Request().SetHost(r.Host)
	// 	for key, val := range r.Header {
	// 		for _, v := range val {
	// 			c.Request().Header.Set(key, v)
	// 		}
	// 	}

	// 	reqCtx := r.Context().Value(jwtmiddleware.ContextKey{})
	// 	fmt.Println(reqCtx)
	// 	c.SetUserContext(r.Context())
	// 	c.Locals("user", reqCtx)
	// 	// more code that relates to ad-hoc custom claims)
	// })
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	authHeader := c.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid authorization header",
		})
	}

	// Validate the token
	tokenInfo, err := jwtValidator.ValidateToken(c.Context(), authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	fmt.Println(tokenInfo)

	return c.Next()

	// errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
	// 	log.Printf("Encountered error while validating JWT: %v", err)

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	// }

	// middleware := jwtmiddleware.New(
	// 	jwtValidator.ValidateToken,
	// 	jwtmiddleware.WithErrorHandler(errorHandler),
	// )
	// err = adaptor.HTTPHandler(middleware.CheckJWT)(c)
	// if err == nil {
	// 	return c.Next()
	// }
	// return nil
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}
