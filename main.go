package main

import (
	"os"
	"flag"
	"fmt"
	"time"
	"strconv"
	"strings"
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

func unixDiffForClaim(diffSecondsStr string) int64 {

	var secondsStr string

	if (! strings.HasPrefix(diffSecondsStr, "+") && ! strings.HasPrefix(diffSecondsStr, "-") ) {
		var str strings.Builder
		str.WriteString("+")
		str.WriteString(diffSecondsStr)
		diffSecondsStr = str.String()
	}
	secondsStr = diffSecondsStr[1:]

	now := time.Now().Unix() 

	secondsInt, err := strconv.Atoi(secondsStr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", "Date Differences (exp/nbf) must be valid integers prefixed with '+' or '-' like '+3600' or '-1000'")
		panic(err)
	}

	seconds := int64(secondsInt)

	var final int64
	if strings.HasPrefix(diffSecondsStr, "+") {
		final =  now + seconds
	} else {
		final = now - seconds
	}

	return final

}

func main() {
	envJwtSecret := os.Getenv("JWT_SECRET")
	if len(envJwtSecret) == 0 {
      	envJwtSecret = ""
  	}
	// jwtbin -secret aaa -c sub:1 -c role:admin -c another:claim 
	var passedClaims passedClaimsSlice

	jwtSecretPtr := flag.String("secret", "none", "JWT Secret (Prefer 'JWT_SECRET' Environment Variable)")	
	expDiffPtr := flag.String("exp-diff", "none", "Expiration Claim (Difference In +/- Seconds from now: +3600, -1000)")
	nbfDiffPtr := flag.String("nbf-diff", "none", "Not Before Claim (Difference In +/- Seconds from now: +3600, -1000)")
	iatDiffPtr := flag.String("iat-diff", "none", "Not Before Claim (Difference In +/- Seconds from now: +3600, -1000)")

	flag.Var(&passedClaims, "c", "List of Additional Claims (Passed in 'key:value' format)")
	flag.Parse()

	var jwtSecret string

	if *jwtSecretPtr == "none" {
		jwtSecret = envJwtSecret
	} else {
		jwtSecret = *jwtSecretPtr
	}

	if len(jwtSecret) < 8 {
		fmt.Fprintf(os.Stderr, "error: %s\n", "Secret too short")
        os.Exit(1)
	}

	claims := jwt.MapClaims{}
	
	if *expDiffPtr != "none" {
		exp := unixDiffForClaim(*expDiffPtr)
		claims["exp"] = exp
	}
	if *nbfDiffPtr != "none" {
		nbf := unixDiffForClaim(*nbfDiffPtr)
		claims["nbf"] = nbf
	}
	if *iatDiffPtr != "none" {
		iat := unixDiffForClaim(*iatDiffPtr)
		claims["iat"] = iat
	}

	for _, passedClaim := range passedClaims {
		keyValStr := strings.Split(passedClaim, ":")
		if len(keyValStr) < 2 {
			fmt.Fprintf(os.Stderr, "error: %s\n", "Key/Value pairs must be passed in 'key:value' format.")
        	os.Exit(1)
		}
		key := keyValStr[0]
		valSlice := keyValStr[1:]
		val := strings.Join(valSlice, ":")
		claims[key] = val
    }


	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecretStrBytes := []byte(jwtSecret)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecretStrBytes)

	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)


}
