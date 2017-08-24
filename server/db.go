package server

// Equipment holds the global equipment definition
type Equipment struct {
	Provider   string
	ProviderID int
	GlobalId   string
	Name       string
	Type       int
	State      int
	LocationID int
}
