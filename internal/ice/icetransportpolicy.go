package ice

// TransportPolicy defines the ICE candidate policy surface the
// permitted candidates. Only these candidates are used for connectivity checks.
type TransportPolicy int

const (
	// TransportPolicyRelay indicates only media relay candidates such
	// as candidates passing through a TURN server are used.
	TransportPolicyRelay TransportPolicy = iota + 1

	// TransportPolicyAll indicates any type of candidate is used.
	TransportPolicyAll
)

// This is done this way because of a linter.
const (
	iceTransportPolicyRelayStr = "relay"
	iceTransportPolicyAllStr   = "all"
)

// NewTransportPolicy takes a string and converts it to TransportPolicy
func NewTransportPolicy(raw string) TransportPolicy {
	switch raw {
	case iceTransportPolicyRelayStr:
		return TransportPolicyRelay
	case iceTransportPolicyAllStr:
		return TransportPolicyAll
	default:
		return TransportPolicy(Unknown)
	}
}

func (t TransportPolicy) String() string {
	switch t {
	case TransportPolicyRelay:
		return iceTransportPolicyRelayStr
	case TransportPolicyAll:
		return iceTransportPolicyAllStr
	default:
		return ErrUnknownType.Error()
	}
}
