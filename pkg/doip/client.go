package doip

// import "net"

// The DoIPClient contains all parameter for a diagnostic connection using DoIP
type DoIPClient struct {
	ECUAddress string
}

// GetECUAddress returns the ECU Address
func (c DoIPClient) GetECUAddress() string {
	return c.ECUAddress
}

// ConnectToECU creates a socket for communication
func (c DoIPClient) ConnectToECU() {
	// do nothing
}