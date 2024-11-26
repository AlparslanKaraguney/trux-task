package filter

type SmartModelFilter struct {
	Limit      int32
	Offset     int32
	Name       string
	Identifier string
	Type       string
	Category   string
	OrderBy    string
}

type SmartFeatureFilter struct {
	Limit         int32
	Offset        int32
	Name          string
	Identifier    string
	Functionality string
	SmartModelId  int32
	OrderBy       string
}
