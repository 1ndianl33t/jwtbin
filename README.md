# JWTBin

A simple binary written in golang to generate JWT tokens for a secret.

## Usage

You can set the secret as the environment variable JWT_SECRET (Preferred) or use the -secret.

### Claims

Claims can be passed in useing the '-c key:value" option. The key and value must be separated by a ":" and not contain more than one semi-colon (for now).

PRs are welcome..
