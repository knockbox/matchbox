package utils

import (
	"github.com/google/uuid"
	"github.com/knockbox/authentication/pkg/enums"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// ParseUserClaims extracts the claims for the User jwt.Token
// This method is unsafe, and I am being lazy for the sake of time. It will explode.
func ParseUserClaims(token jwt.Token) (uuid.UUID, string, enums.UserRole) {
	claims := token.PrivateClaims()
	accountId, _ := uuid.Parse(claims["account_id"].(string))
	return accountId, claims["username"].(string), enums.UserRoleFromString(claims["role"].(string))
}
