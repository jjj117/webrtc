package ice

import (
	"net"
	"testing"

	"github.com/pion/ice"
	"github.com/stretchr/testify/assert"
)

func TestCandidate_Convert(t *testing.T) {
	testCases := []struct {
		native Candidate
		ice    *ice.Candidate
	}{
		{
			Candidate{
				Foundation: "foundation",
				Priority:   128,
				IP:         "1.0.0.1",
				Protocol:   ProtocolUDP,
				Port:       1234,
				Typ:        CandidateTypeHost,
				Component:  1,
			}, &ice.Candidate{
				IP:              net.ParseIP("1.0.0.1"),
				NetworkType:     ice.NetworkTypeUDP4,
				Port:            1234,
				Type:            ice.CandidateTypeHost,
				Component:       1,
				LocalPreference: 65535,
			},
		},
		{
			Candidate{
				Foundation:     "foundation",
				Priority:       128,
				IP:             "::1",
				Protocol:       ProtocolUDP,
				Port:           1234,
				Typ:            CandidateTypeSrflx,
				Component:      1,
				RelatedAddress: "1.0.0.1",
				RelatedPort:    4321,
			}, &ice.Candidate{
				IP:              net.ParseIP("::1"),
				NetworkType:     ice.NetworkTypeUDP6,
				Port:            1234,
				Type:            ice.CandidateTypeServerReflexive,
				Component:       1,
				LocalPreference: 65535,
				RelatedAddress: &ice.CandidateRelatedAddress{
					Address: "1.0.0.1",
					Port:    4321,
				},
			},
		},
		{
			Candidate{
				Foundation:     "foundation",
				Priority:       128,
				IP:             "::1",
				Protocol:       ProtocolUDP,
				Port:           1234,
				Typ:            CandidateTypePrflx,
				Component:      1,
				RelatedAddress: "1.0.0.1",
				RelatedPort:    4321,
			}, &ice.Candidate{
				IP:              net.ParseIP("::1"),
				NetworkType:     ice.NetworkTypeUDP6,
				Port:            1234,
				Type:            ice.CandidateTypePeerReflexive,
				Component:       1,
				LocalPreference: 65535,
				RelatedAddress: &ice.CandidateRelatedAddress{
					Address: "1.0.0.1",
					Port:    4321,
				},
			},
		},
	}

	for i, testCase := range testCases {
		actualICE, err := testCase.native.toICE()
		assert.Nil(t, err)
		assert.Equal(t,
			testCase.ice,
			actualICE,
			"testCase: %d ice not equal %v", i, actualICE,
		)
	}
}

func TestConvertTypeFromICE(t *testing.T) {
	t.Run("host", func(t *testing.T) {
		ct, err := convertTypeFromICE(ice.CandidateTypeHost)
		if err != nil {
			t.Fatal("failed coverting ice.CandidateTypeHost")
		}
		if ct != CandidateTypeHost {
			t.Fatal("should be coverted to CandidateTypeHost")
		}
	})
	t.Run("srflx", func(t *testing.T) {
		ct, err := convertTypeFromICE(ice.CandidateTypeServerReflexive)
		if err != nil {
			t.Fatal("failed coverting ice.CandidateTypeServerReflexive")
		}
		if ct != CandidateTypeSrflx {
			t.Fatal("should be coverted to CandidateTypeSrflx")
		}
	})
	t.Run("prflx", func(t *testing.T) {
		ct, err := convertTypeFromICE(ice.CandidateTypePeerReflexive)
		if err != nil {
			t.Fatal("failed coverting ice.CandidateTypePeerReflexive")
		}
		if ct != CandidateTypePrflx {
			t.Fatal("should be coverted to CandidateTypePrflx")
		}
	})
}
