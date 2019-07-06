# JWTBin

A simple binary written in golang to generate JWT tokens for a secret.

## Usage

You can set the secret as the environment variable JWT_SECRET (Preferred) or use the -secret.

### Claims

Claims can be passed in useing the '-c key:value" option. The key and value must be separated by a ":".

PRs are welcome.

```
Usage of ./jwtbin:
  -c value
    	List of Additional Claims (Passed in 'key:value' format)
  -exp-diff string
    	Expiration Claim (Difference In +/- Seconds from now: +3600, -1000) (default "none")
  -iat-diff string
    	Not Before Claim (Difference In +/- Seconds from now: +3600, -1000) (default "none")
  -nbf-diff string
    	Not Before Claim (Difference In +/- Seconds from now: +3600, -1000) (default "none")
  -secret string
    	JWT Secret (Prefer 'JWT_SECRET' Environment Variable) (default "none")

```
