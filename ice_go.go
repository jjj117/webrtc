// +build !js

package webrtc

import (
	"github.com/pion/sdp/v2"
	"github.com/pion/webrtc/v2/internal/ice"
)

//go:generate go run internal/tools/gen/genaliasdocs.go -pkg "./internal/ice" -build-tags "!js" $GOFILE

const (

	// ICETransportPolicyRelay indicates only media relay candidates such
	// as candidates passing through a TURN server are used.
	ICETransportPolicyRelay = ice.ICETransportPolicyRelay

	// ICETransportPolicyAll indicates any type of candidate is used.
	ICETransportPolicyAll = ice.ICETransportPolicyAll

	// ICETransportStateNew indicates the ICETransport is waiting
	// for remote candidates to be supplied.
	ICETransportStateNew = ice.ICETransportStateNew

	// ICETransportStateChecking indicates the ICETransport has
	// received at least one remote candidate, and a local and remote
	// ICECandidateComplete dictionary was not added as the last candidate.
	ICETransportStateChecking = ice.ICETransportStateChecking

	// ICETransportStateConnected indicates the ICETransport has
	// received a response to an outgoing connectivity check, or has
	// received incoming DTLS/media after a successful response to an
	// incoming connectivity check, but is still checking other candidate
	// pairs to see if there is a better connection.
	ICETransportStateConnected = ice.ICETransportStateConnected

	// ICETransportStateCompleted indicates the ICETransport tested
	// all appropriate candidate pairs and at least one functioning
	// candidate pair has been found.
	ICETransportStateCompleted = ice.ICETransportStateCompleted

	// ICETransportStateFailed indicates the ICETransport the last
	// candidate was added and all appropriate candidate pairs have either
	// failed connectivity checks or have lost consent.
	ICETransportStateFailed = ice.ICETransportStateFailed

	// ICETransportStateDisconnected indicates the ICETransport has received
	// at least one local and remote candidate, but the final candidate was
	// received yet and all appropriate candidate pairs thus far have been
	// tested and failed.
	ICETransportStateDisconnected = ice.ICETransportStateDisconnected

	// ICETransportStateClosed indicates the ICETransport has shut down
	// and is no longer responding to STUN requests.
	ICETransportStateClosed = ice.ICETransportStateClosed

	// ICEConnectionStateNew indicates that any of the ICETransports are
	// in the "new" state and none of them are in the "checking", "disconnected"
	// or "failed" state, or all ICETransports are in the "closed" state, or
	// there are no transports.
	ICEConnectionStateNew = ice.ICEConnectionStateNew

	// ICEConnectionStateChecking indicates that any of the ICETransports
	// are in the "checking" state and none of them are in the "disconnected"
	// or "failed" state.
	ICEConnectionStateChecking = ice.ICEConnectionStateChecking

	// ICEConnectionStateConnected indicates that all ICETransports are
	// in the "connected", "completed" or "closed" state and at least one of
	// them is in the "connected" state.
	ICEConnectionStateConnected = ice.ICEConnectionStateConnected

	// ICEConnectionStateCompleted indicates that all ICETransports are
	// in the "completed" or "closed" state and at least one of them is in the
	// "completed" state.
	ICEConnectionStateCompleted = ice.ICEConnectionStateCompleted

	// ICEConnectionStateDisconnected indicates that any of the
	// ICETransports are in the "disconnected" state and none of them are
	// in the "failed" state.
	ICEConnectionStateDisconnected = ice.ICEConnectionStateDisconnected

	// ICEConnectionStateFailed indicates that any of the ICETransports
	// are in the "failed" state.
	ICEConnectionStateFailed = ice.ICEConnectionStateFailed

	// ICEConnectionStateClosed indicates that the PeerConnection's
	// isClosed is true.
	ICEConnectionStateClosed = ice.ICEConnectionStateClosed

	// ICEGatheringStateNew indicates that any of the ICETransports are
	// in the "new" gathering state and none of the transports are in the
	// "gathering" state, or there are no transports.
	ICEGatheringStateNew = ice.ICEGatheringStateNew

	// ICEGatheringStateGathering indicates that any of the ICETransports
	// are in the "gathering" state.
	ICEGatheringStateGathering = ice.ICEGatheringStateGathering

	// ICEGatheringStateComplete indicates that at least one ICETransport
	// exists, and all ICETransports are in the "completed" gathering state.
	ICEGatheringStateComplete = ice.ICEGatheringStateComplete

	// ICERoleControlled indicates that an ICE agent that waits for the
	// controlling agent to select the final choice of candidate pairs.
	ICERoleControlled = ice.ICERoleControlled

	// ICERoleControlling indicates that the ICE agent that is responsible
	// for selecting the final choice of candidate pairs and signaling them
	// through STUN and an updated offer, if needed. In any session, one agent
	// is always controlling. The other is the controlled agent.
	ICERoleControlling = ice.ICERoleControlling

	// ICECredentialTypePassword describes username and pasword based
	// credentials as described in https://tools.ietf.org/html/rfc5389.
	ICECredentialTypePassword = ice.ICECredentialTypePassword

	// ICECredentialTypeOauth describes token based credential as described
	// in https://tools.ietf.org/html/rfc7635.
	ICECredentialTypeOauth = ice.ICECredentialTypeOauth

	// ICECandidateTypeHost indicates that the candidate is of Host type as
	// described in https://tools.ietf.org/html/rfc8445#section-5.1.1.1. A
	// candidate obtained by binding to a specific port from an IP address on
	// the host. This includes IP addresses on physical interfaces and logical
	// ones, such as ones obtained through VPNs.
	ICECandidateTypeHost = ice.ICECandidateTypeHost

	// ICECandidateTypeSrflx indicates the the candidate is of Server
	// Reflexive type as described
	// https://tools.ietf.org/html/rfc8445#section-5.1.1.2. A candidate type
	// whose IP address and port are a binding allocated by a NAT for an ICE
	// agent after it sends a packet through the NAT to a server, such as a
	// STUN server.
	ICECandidateTypeSrflx = ice.ICECandidateTypeSrflx

	// ICECandidateTypePrflx indicates that the candidate is of Peer
	// Reflexive type. A candidate type whose IP address and port are a binding
	// allocated by a NAT for an ICE agent after it sends a packet through the
	// NAT to its peer.
	ICECandidateTypePrflx = ice.ICECandidateTypePrflx

	// ICECandidateTypeRelay indicates the the candidate is of Relay type as
	// described in https://tools.ietf.org/html/rfc8445#section-5.1.1.2. A
	// candidate type obtained from a relay server, such as a TURN server.
	ICECandidateTypeRelay = ice.ICECandidateTypeRelay

	// NetworkTypeUDP4 indicates UDP over IPv4.
	NetworkTypeUDP4 = ice.NetworkTypeUDP4

	// NetworkTypeUDP6 indicates UDP over IPv6.
	NetworkTypeUDP6 = ice.NetworkTypeUDP6

	// NetworkTypeTCP4 indicates TCP over IPv4.
	NetworkTypeTCP4 = ice.NetworkTypeTCP4

	// NetworkTypeTCP6 indicates TCP over IPv6.
	NetworkTypeTCP6 = ice.NetworkTypeTCP6

	// ICEProtocolUDP indicates the URL uses a UDP transport.
	ICEProtocolUDP = ice.ICEProtocolUDP

	// ICEProtocolTCP indicates the URL uses a TCP transport.
	ICEProtocolTCP = ice.ICEProtocolTCP

	// ICEComponentRTP indicates that the ICE Transport is used for RTP (or
	// RTCP multiplexing), as defined in
	// https://tools.ietf.org/html/rfc5245#section-4.1.1.1. Protocols
	// multiplexed with RTP (e.g. data channel) share its component ID. This
	// represents the component-id value 1 when encoded in candidate-attribute.
	ICEComponentRTP = ice.ICEComponentRTP

	// ICEComponentRTCP indicates that the ICE Transport is used for RTCP as
	// defined by https://tools.ietf.org/html/rfc5245#section-4.1.1.1. This
	// represents the component-id value 2 when encoded in candidate-attribute.
	ICEComponentRTCP = ice.ICEComponentRTCP

	// ICEGathererStateNew indicates object has been created but
	// gather() has not been called.
	ICEGathererStateNew = ice.ICEGathererStateNew

	// ICEGathererStateGathering indicates gather() has been called,
	// and the ICEGatherer is in the process of gathering candidates.
	ICEGathererStateGathering = ice.ICEGathererStateGathering

	// ICEGathererStateComplete indicates the ICEGatherer has completed gathering.
	ICEGathererStateComplete = ice.ICEGathererStateComplete

	// ICEGathererStateClosed indicates the closed state can only be entered
	// when the ICEGatherer has been closed intentionally by calling close().
	ICEGathererStateClosed = ice.ICEGathererStateClosed
)

type (

	// ICEServer describes a single STUN and TURN server that can be used by
	// the ICEAgent to establish a connection with a peer.
	ICEServer = ice.ICEServer

	// ICETransportPolicy defines the ICE candidate policy surface the
	// permitted candidates. Only these candidates are used for connectivity checks.
	ICETransportPolicy = ice.ICETransportPolicy

	// ICETransport allows an application access to information about the ICE
	// transport over which packets are sent and received.
	ICETransport = ice.ICETransport

	// ICEGatheringState describes the state of the candidate gathering process.
	ICEGatheringState = ice.ICEGatheringState

	// ICEConnectionState indicates signaling state of the ICE Connection.
	ICEConnectionState = ice.ICEConnectionState

	// ICECandidate represents a ice candidate
	ICECandidate = ice.ICECandidate

	// ICEGatherer gathers local host, server reflexive and relay
	// candidates, as well as enabling the retrieval of local Interactive
	// Connectivity Establishment (ICE) parameters which can be
	// exchanged in signaling.
	ICEGatherer = ice.ICEGatherer

	// ICEGatherOptions provides options relating to the gathering of ICE candidates.
	ICEGatherOptions = ice.ICEGatherOptions

	// ICETransportState represents the current state of the ICE transport.
	ICETransportState = ice.ICETransportState

	// ICERole describes the role ice.Agent is playing in selecting the
	// preferred the candidate pair.
	ICERole = ice.ICERole

	// ICEParameters includes the ICE username fragment
	// and password and other ICE-related parameters.
	ICEParameters = ice.ICEParameters

	// ICECandidateInit is used to serialize ice candidates
	ICECandidateInit = ice.ICECandidateInit

	// ICECandidateType represents the type of the ICE candidate used.
	ICECandidateType = ice.ICECandidateType

	// ICECredentialType indicates the type of credentials used to connect to
	// an ICE server.
	ICECredentialType = ice.ICECredentialType

	// ICEProtocol indicates the transport protocol type that is used in the
	// ice.URL structure.
	ICEProtocol = ice.ICEProtocol

	// ICECandidatePair represents an ICE Candidate pair
	ICECandidatePair = ice.ICECandidatePair

	// ICEComponent describes if the ice transport is used for RTP
	// (or RTCP multiplexing).
	ICEComponent = ice.ICEComponent

	// ICEGathererState represents the current state of the ICE gatherer.
	ICEGathererState = ice.ICEGathererState

	// NetworkType represents the type of network
	NetworkType = ice.NetworkType

	// OAuthCredential represents OAuth credential information which is used by
	// the STUN/TURN client to connect to an ICE server as defined in
	// https://tools.ietf.org/html/rfc7635. Note that the kid parameter is not
	// located in OAuthCredential, but in ICEServer's username member.
	OAuthCredential = ice.OAuthCredential
)

var (

	// NewICEGatherer creates a new NewICEGatherer.
	// This constructor is part of the ORTC API. It is not
	// meant to be used together with the basic WebRTC API.
	NewICEGatherer = ice.NewICEGatherer

	// NewICETransport creates a new NewICETransport.
	// This constructor is part of the ORTC API. It is not
	// meant to be used together with the basic WebRTC API.
	NewICETransport = ice.NewICETransport

	// NewICECandidatePair returns an initialized *ICECandidatePair
	// for the given pair of ICECandidate instances
	NewICECandidatePair = ice.NewICECandidatePair

	// NewICECandidateType takes a string and converts it into ICECandidateType
	NewICECandidateType = ice.NewICECandidateType

	// NewICEProtocol takes a string and converts it to ICEProtocol
	NewICEProtocol = ice.NewICEProtocol

	// NewICEConnectionState takes a string and converts it to ICEConnectionState
	NewICEConnectionState = ice.NewICEConnectionState

	// NewICEGatheringState takes a string and converts it to ICEGatheringState
	NewICEGatheringState = ice.NewICEGatheringState

	// NewICETransportPolicy takes a string and converts it to ICETransportPolicy
	NewICETransportPolicy = ice.NewICETransportPolicy

	// ErrNoTurnCredencials indicates that a TURN server URL was provided
	// without required credentials.
	ErrNoTurnCredencials = ice.ErrNoTurnCredencials

	// ErrTurnCredencials indicates that provided TURN credentials are partial
	// or malformed.
	ErrTurnCredencials = ice.ErrTurnCredencials
)

// NewICEGatherer creates a new NewICEGatherer.
// This constructor is part of the ORTC API. It is not
// meant to be used together with the basic WebRTC API.
func (api *API) NewICEGatherer(opts ICEGatherOptions) (*ICEGatherer, error) {
	return NewICEGatherer(
		api.settingEngine.ephemeralUDP.PortMin,
		api.settingEngine.ephemeralUDP.PortMax,
		api.settingEngine.timeout.ICEConnection,
		api.settingEngine.timeout.ICEKeepalive,
		api.settingEngine.LoggerFactory,
		api.settingEngine.candidates.ICENetworkTypes,
		opts,
	)
}

// NewICETransport creates a new NewICETransport.
// This constructor is part of the ORTC API. It is not
// meant to be used together with the basic WebRTC API.
func (api *API) NewICETransport(gatherer *ICEGatherer) *ICETransport {
	return NewICETransport(gatherer, api.settingEngine.LoggerFactory)
}

// Conversion for package sdp
func newICECandidateFromSDP(c sdp.ICECandidate) (ICECandidate, error) {
	typ, err := NewICECandidateType(c.Typ)
	if err != nil {
		return ICECandidate{}, err
	}
	protocol, err := NewICEProtocol(c.Protocol)
	if err != nil {
		return ICECandidate{}, err
	}
	return ICECandidate{
		Foundation:     c.Foundation,
		Priority:       c.Priority,
		IP:             c.IP,
		Protocol:       protocol,
		Port:           c.Port,
		Component:      c.Component,
		Typ:            typ,
		RelatedAddress: c.RelatedAddress,
		RelatedPort:    c.RelatedPort,
	}, nil
}

func iceCandidateToSDP(c ICECandidate) sdp.ICECandidate {
	return sdp.ICECandidate{
		Foundation:     c.Foundation,
		Priority:       c.Priority,
		IP:             c.IP,
		Protocol:       c.Protocol.String(),
		Port:           c.Port,
		Component:      c.Component,
		Typ:            c.Typ.String(),
		RelatedAddress: c.RelatedAddress,
		RelatedPort:    c.RelatedPort,
	}
}
