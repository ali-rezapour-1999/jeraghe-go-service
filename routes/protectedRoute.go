package routes

import (
	"go-server/middleware"
	"regexp"
)

type ProtectionRule = middleware.ProtectionRule

func GetProtectionRules() []ProtectionRule {
	return []ProtectionRule{
		{
			PathPattern:  regexp.MustCompile(`^/api/private/auth/login/`),
			RequiresAuth: false,
		},
		{
			PathPattern:  regexp.MustCompile(`^/api/private/auth/register/`),
			RequiresAuth: false,
		},
		{
			PathPattern:  regexp.MustCompile(`^/api/private/auth/token-verify/`),
			RequiresAuth: false,
		},
		{
			PathPattern:  regexp.MustCompile(`^/api/private/auth/token-refresh/`),
			RequiresAuth: false,
		},
		{
			PathPattern:  regexp.MustCompile(`^/api/private/auth/.*`),
			RequiresAuth: true,
		},
		{
			PathPattern:  regexp.MustCompile(`^/api/private/profile/.*`),
			RequiresAuth: true,
		},
	}
}
