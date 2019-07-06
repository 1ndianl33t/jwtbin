package main

import (
	"os"
	"flag"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go" 
)

type passedClaimsSlice []string
 
func (claims *passedClaimsSlice) String() string {
    return fmt.Sprintf("%d", *claims)
}
 
// The second method is Set(value string) error
func (claims *passedClaimsSlice) Set(value string) error {
	*claims = append(*claims, value)
    return nil
}

func main() {
	envJwtSecret := os.Getenv("JWT_SECRET")
	if len(envJwtSecret) == 0 {
      	envJwtSecret = ""
  	}
	// jwtbin -secret aaa -c sub 1 -c role admin -c another claim 
	var passedClaims passedClaimsSlice

	flag.Var(&passedClaims, "c", "List of Claims Key/Values")

	jwtSecretPtr := flag.String("secret", envJwtSecret, "JWT Secret (Prefer 'JWT_SECRET' Environment Variable)")
	
	flag.Parse()
	fmt.Printf("%+v\n", passedClaims)
	jwtSecret := *jwtSecretPtr

	// if len(jwtSecret) < 10 {
	// 	fmt.Fprintf(os.Stderr, "error: %s\n", "Secret too short")
    //     os.Exit(1)
	// }


	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)

	fmt.Println(tokenString, err)


}
