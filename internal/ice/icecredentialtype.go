package ice

// CredentialType indicates the type of credentials used to connect to
// an ICE server.
type CredentialType int

const (
	// CredentialTypePassword describes username and pasword based
	// credentials as described in https://tools.ietf.org/html/rfc5389.
	CredentialTypePassword CredentialType = iota + 1

	// CredentialTypeOauth describes token based credential as described
	// in https://tools.ietf.org/html/rfc7635.
	CredentialTypeOauth
)

// This is done this way because of a linter.
const (
	iceCredentialTypePasswordStr = "password"
	iceCredentialTypeOauthStr    = "oauth"
)

func newCredentialType(raw string) CredentialType {
	switch raw {
	case iceCredentialTypePasswordStr:
		return CredentialTypePassword
	case iceCredentialTypeOauthStr:
		return CredentialTypeOauth
	default:
		return CredentialType(Unknown)
	}
}

func (t CredentialType) String() string {
	switch t {
	case CredentialTypePassword:
		return iceCredentialTypePasswordStr
	case CredentialTypeOauth:
		return iceCredentialTypeOauthStr
	default:
		return ErrUnknownType.Error()
	}
}
