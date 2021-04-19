// Application entities.
//
// @author TSS

package domain

type Keys struct {
	DerivedKey  []byte
	DerivedMac  []byte
	MasterKey   []byte
	MasterMac   []byte
	OverviewKey []byte
	OverviewMac []byte
}

func NewKeys(derivedKey, derivedMac, masterKey, masterMac, overviewKey, overviewMac []byte) *Keys {
	return &Keys{
		DerivedKey:  derivedKey,
		DerivedMac:  derivedMac,
		MasterKey:   masterKey,
		MasterMac:   masterMac,
		OverviewKey: overviewKey,
		OverviewMac: overviewMac,
	}
}
