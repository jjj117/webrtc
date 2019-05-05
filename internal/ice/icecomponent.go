package ice

// Component describes if the ice transport is used for RTP
// (or RTCP multiplexing).
type Component int

const (
	// ComponentRTP indicates that the ICE Transport is used for RTP (or
	// RTCP multiplexing), as defined in
	// https://tools.ietf.org/html/rfc5245#section-4.1.1.1. Protocols
	// multiplexed with RTP (e.g. data channel) share its component ID. This
	// represents the component-id value 1 when encoded in candidate-attribute.
	ComponentRTP Component = iota + 1

	// ComponentRTCP indicates that the ICE Transport is used for RTCP as
	// defined by https://tools.ietf.org/html/rfc5245#section-4.1.1.1. This
	// represents the component-id value 2 when encoded in candidate-attribute.
	ComponentRTCP
)

// This is done this way because of a linter.
const (
	iceComponentRTPStr  = "rtp"
	iceComponentRTCPStr = "rtcp"
)

func newComponent(raw string) Component {
	switch raw {
	case iceComponentRTPStr:
		return ComponentRTP
	case iceComponentRTCPStr:
		return ComponentRTCP
	default:
		return Component(Unknown)
	}
}

func (t Component) String() string {
	switch t {
	case ComponentRTP:
		return iceComponentRTPStr
	case ComponentRTCP:
		return iceComponentRTCPStr
	default:
		return ErrUnknownType.Error()
	}
}
