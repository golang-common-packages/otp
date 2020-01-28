package otp

import "github.com/xlzd/gotp"

// IOTP interface for this package
type IOTP interface {
	GetOTP(totp *gotp.TOTP) string
	Default(secret string) *gotp.TOTP
	Custom(secret string, digits, interval int) *gotp.TOTP
	GoogleAuthenticator(secret, accountName, issuerName string) string
	Verify(totp *gotp.TOTP, secret string, timestamp int) bool
}
