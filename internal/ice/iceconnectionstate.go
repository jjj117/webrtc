package ice

// ConnectionState indicates signaling state of the ICE Connection.
type ConnectionState int

const (
	// ConnectionStateNew indicates that any of the ICETransports are
	// in the "new" state and none of them are in the "checking", "disconnected"
	// or "failed" state, or all ICETransports are in the "closed" state, or
	// there are no transports.
	ConnectionStateNew ConnectionState = iota + 1

	// ConnectionStateChecking indicates that any of the ICETransports
	// are in the "checking" state and none of them are in the "disconnected"
	// or "failed" state.
	ConnectionStateChecking

	// ConnectionStateConnected indicates that all ICETransports are
	// in the "connected", "completed" or "closed" state and at least one of
	// them is in the "connected" state.
	ConnectionStateConnected

	// ConnectionStateCompleted indicates that all ICETransports are
	// in the "completed" or "closed" state and at least one of them is in the
	// "completed" state.
	ConnectionStateCompleted

	// ConnectionStateDisconnected indicates that any of the
	// ICETransports are in the "disconnected" state and none of them are
	// in the "failed" state.
	ConnectionStateDisconnected

	// ConnectionStateFailed indicates that any of the ICETransports
	// are in the "failed" state.
	ConnectionStateFailed

	// ConnectionStateClosed indicates that the PeerConnection's
	// isClosed is true.
	ConnectionStateClosed
)

// This is done this way because of a linter.
const (
	iceConnectionStateNewStr          = "new"
	iceConnectionStateCheckingStr     = "checking"
	iceConnectionStateConnectedStr    = "connected"
	iceConnectionStateCompletedStr    = "completed"
	iceConnectionStateDisconnectedStr = "disconnected"
	iceConnectionStateFailedStr       = "failed"
	iceConnectionStateClosedStr       = "closed"
)

// NewConnectionState takes a string and converts it to ConnectionState
func NewConnectionState(raw string) ConnectionState {
	switch raw {
	case iceConnectionStateNewStr:
		return ConnectionStateNew
	case iceConnectionStateCheckingStr:
		return ConnectionStateChecking
	case iceConnectionStateConnectedStr:
		return ConnectionStateConnected
	case iceConnectionStateCompletedStr:
		return ConnectionStateCompleted
	case iceConnectionStateDisconnectedStr:
		return ConnectionStateDisconnected
	case iceConnectionStateFailedStr:
		return ConnectionStateFailed
	case iceConnectionStateClosedStr:
		return ConnectionStateClosed
	default:
		return ConnectionState(Unknown)
	}
}

func (c ConnectionState) String() string {
	switch c {
	case ConnectionStateNew:
		return iceConnectionStateNewStr
	case ConnectionStateChecking:
		return iceConnectionStateCheckingStr
	case ConnectionStateConnected:
		return iceConnectionStateConnectedStr
	case ConnectionStateCompleted:
		return iceConnectionStateCompletedStr
	case ConnectionStateDisconnected:
		return iceConnectionStateDisconnectedStr
	case ConnectionStateFailed:
		return iceConnectionStateFailedStr
	case ConnectionStateClosed:
		return iceConnectionStateClosedStr
	default:
		return ErrUnknownType.Error()
	}
}
