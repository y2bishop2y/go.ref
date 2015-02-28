// +build nacl

package platform

// GetPlatform returns the description of the Platform this process is running on.
// A default value for Platform is provided even if an error is
// returned; nil is never returned for the first return result.
func GetPlatform() (*Platform, error) {
	d := &Platform{
		Vendor:  "google",
		Model:   "generic",
		System:  "nacl",
		Version: "0",
		Release: "0",
		Machine: "0",
		Node:    "0",
	}
	return d, nil
}
