package revocation

import (
	"time"
)

func NewMockRevocationManager() RevocationManager {
	revocationDB = &mockDatabase{make(map[string][]byte), make(map[string]*time.Time)}
	return &revocationManager{}
}

type mockDatabase struct {
	tpCavIDToRevCavID   map[string][]byte
	revCavIDToTimestamp map[string]*time.Time
}

func (m *mockDatabase) InsertCaveat(thirdPartyCaveatID string, revocationCaveatID []byte) error {
	m.tpCavIDToRevCavID[thirdPartyCaveatID] = revocationCaveatID
	return nil
}

func (m *mockDatabase) Revoke(thirdPartyCaveatID string) error {
	timestamp := time.Now()
	m.revCavIDToTimestamp[string(m.tpCavIDToRevCavID[thirdPartyCaveatID])] = &timestamp
	return nil
}

func (m *mockDatabase) IsRevoked(revocationCaveatID []byte) (bool, error) {
	_, exists := m.revCavIDToTimestamp[string(revocationCaveatID)]
	return exists, nil
}

func (m *mockDatabase) RevocationTime(thirdPartyCaveatID string) (*time.Time, error) {
	return m.revCavIDToTimestamp[string(m.tpCavIDToRevCavID[thirdPartyCaveatID])], nil
}
