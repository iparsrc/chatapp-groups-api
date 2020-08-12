package domain

import "testing"

func TestConnectDB(t *testing.T) {
	ConnectDB("mongodb://localhost:27017")
}
