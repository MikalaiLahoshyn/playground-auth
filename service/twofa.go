package service

import "github.com/pquerna/otp/totp"

//TODO: store secret in redis with TTL 30 seconds
//Possible options:

//QR
//SMS
//Email
//Push
//Backup Codes

// Generates Time-Based-One-Time password
func GenerateTOTP(username string) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "a3ther-playground",
		AccountName: username,
	})

	return key.Secret(), err
}

// Validates TOTP code
func ValidateTOTP(passcode string) error {
	//valid := totp.Validate(passcode)
	return nil
}
