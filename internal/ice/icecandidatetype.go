package ice

import "fmt"

// CandidateType represents the type of the ICE candidate used.
type CandidateType int

const (
	// CandidateTypeHost indicates that the candidate is of Host type as
	// described in https://tools.ietf.org/html/rfc8445#section-5.1.1.1. A
	// candidate obtained by binding to a specific port from an IP address on
	// the host. This includes IP addresses on physical interfaces and logical
	// ones, such as ones obtained through VPNs.
	CandidateTypeHost CandidateType = iota + 1

	// CandidateTypeSrflx indicates the the candidate is of Server
	// Reflexive type as described
	// https://tools.ietf.org/html/rfc8445#section-5.1.1.2. A candidate type
	// whose IP address and port are a binding allocated by a NAT for an ICE
	// agent after it sends a packet through the NAT to a server, such as a
	// STUN server.
	CandidateTypeSrflx

	// CandidateTypePrflx indicates that the candidate is of Peer
	// Reflexive type. A candidate type whose IP address and port are a binding
	// allocated by a NAT for an ICE agent after it sends a packet through the
	// NAT to its peer.
	CandidateTypePrflx

	// CandidateTypeRelay indicates the the candidate is of Relay type as
	// described in https://tools.ietf.org/html/rfc8445#section-5.1.1.2. A
	// candidate type obtained from a relay server, such as a TURN server.
	CandidateTypeRelay
)

// This is done this way because of a linter.
const (
	iceCandidateTypeHostStr  = "host"
	iceCandidateTypeSrflxStr = "srflx"
	iceCandidateTypePrflxStr = "prflx"
	iceCandidateTypeRelayStr = "relay"
)

// NewCandidateType takes a string and converts it into CandidateType
func NewCandidateType(raw string) (CandidateType, error) {
	switch raw {
	case iceCandidateTypeHostStr:
		return CandidateTypeHost, nil
	case iceCandidateTypeSrflxStr:
		return CandidateTypeSrflx, nil
	case iceCandidateTypePrflxStr:
		return CandidateTypePrflx, nil
	case iceCandidateTypeRelayStr:
		return CandidateTypeRelay, nil
	default:
		return CandidateType(Unknown), fmt.Errorf("unknown ICE candidate type: %s", raw)
	}
}

func (t CandidateType) String() string {
	switch t {
	case CandidateTypeHost:
		return iceCandidateTypeHostStr
	case CandidateTypeSrflx:
		return iceCandidateTypeSrflxStr
	case CandidateTypePrflx:
		return iceCandidateTypePrflxStr
	case CandidateTypeRelay:
		return iceCandidateTypeRelayStr
	default:
		return ErrUnknownType.Error()
	}
}
