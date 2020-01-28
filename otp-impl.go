package otp

import (
	"github.com/xlzd/gotp"
)

// Client manage all OTP action
type Client struct{}

// GetOTP function get OTP code from OTP object
func (c *Client) GetOTP(totp *gotp.TOTP) string {
	return totp.Now()
}

// Default function will generate an OTP object based on your secret
func (c *Client) Default(secret string) *gotp.TOTP {
	return gotp.NewDefaultTOTP(secret)
}

// Custom function will generate an OTP object based on your secret, digits and time interval
func (c *Client) Custom(secret string, digits, interval int) *gotp.TOTP {
	return gotp.NewTOTP(secret, digits, interval, nil)
}

// GoogleAuthenticator works with the Google Authenticator iPhone and Android app, as well as other OTP apps like Authy
func (c *Client) GoogleAuthenticator(secret, accountName, issuerName string) string {
	return gotp.NewDefaultTOTP(secret).ProvisioningUri(accountName, issuerName)
}

// Verify your OTP with timestamp
func (c *Client) Verify(totp *gotp.TOTP, secret string, timestamp int) bool {
	return totp.Verify(secret, timestamp)
}
