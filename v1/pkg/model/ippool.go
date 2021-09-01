// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package model

// IPPool defines get, list and allocation response body.
type IPPool struct {
	ResourceBase
	Description string

	IPVersion    IPVer
	NetworkID    string
	BaseIP       string
	Netmask      Netmask
	DefaultRoute string
	Sources      []IPSource

	UseRecords []UseRecord

	DNS     []string
	Proxy   string
	NoProxy string
	NTP     []string

	Pool Pool
}

// IPPoolUpdate defines the request body to update an IP pool.
type IPPoolUpdate struct {
	ResourceBase
	Description  string
	DefaultRoute string
	Proxy        string
	NoProxy      string
	NTP          []string
	DNS          []string
}

// NewIPPool defines the request body to create a new IP pool.
type NewIPPool struct {
	Name         string
	Description  string
	IPVersion    IPVer
	BaseIP       string
	Netmask      Netmask
	DefaultRoute string
	Sources      []IPSource
	DNS          []string
	Proxy        string
	NoProxy      string
	NTP          []string
}

// IPVer defines IP versions.
type IPVer string

// Following are possible IP versions.
const (
	IPVerIPv4 IPVer = "IPv4"
	IPVerIPv6 IPVer = "IPv6"
)

// Netmask defines a network subnet mask in CIDR format.
type Netmask string

// Following define possible Netmask values.
const (
	// IPv4
	Netmask8  Netmask = "/8"
	Netmask9  Netmask = "/9"
	Netmask10 Netmask = "/10"
	Netmask11 Netmask = "/11"
	Netmask12 Netmask = "/12"
	Netmask13 Netmask = "/13"
	Netmask14 Netmask = "/14"
	Netmask15 Netmask = "/15"
	Netmask16 Netmask = "/16"
	Netmask17 Netmask = "/17"
	Netmask18 Netmask = "/18"
	Netmask19 Netmask = "/19"
	Netmask20 Netmask = "/20"
	Netmask21 Netmask = "/21"
	Netmask22 Netmask = "/22"
	Netmask23 Netmask = "/23"
	Netmask24 Netmask = "/24"
	Netmask25 Netmask = "/25"
	Netmask26 Netmask = "/26"
	Netmask27 Netmask = "/27"
	Netmask28 Netmask = "/28"
	Netmask29 Netmask = "/29"
	Netmask30 Netmask = "/30"
	Netmask31 Netmask = "/31"

	// IPv6
	Netmask104 Netmask = "/104"
	Netmask105 Netmask = "/105"
	Netmask106 Netmask = "/106"
	Netmask107 Netmask = "/107"
	Netmask108 Netmask = "/108"
	Netmask109 Netmask = "/109"
	Netmask110 Netmask = "/110"
	Netmask111 Netmask = "/111"
	Netmask112 Netmask = "/112"
	Netmask113 Netmask = "/113"
	Netmask114 Netmask = "/114"
	Netmask115 Netmask = "/115"
	Netmask116 Netmask = "/116"
	Netmask117 Netmask = "/117"
	Netmask118 Netmask = "/118"
	Netmask119 Netmask = "/119"
	Netmask120 Netmask = "/120"
	Netmask121 Netmask = "/121"
	Netmask122 Netmask = "/122"
	Netmask123 Netmask = "/123"
	Netmask124 Netmask = "/124"
	Netmask125 Netmask = "/125"
	Netmask126 Netmask = "/126"
	Netmask127 Netmask = "/127"
)

// IPSource defines a structure for an IP pool source.
type IPSource struct {
	Base  string
	Count uint32
}

// UseRecord defines an IP address usage record.
// External is used to indicate that the IP has
// been manually allocated for use outside quake.
type UseRecord struct {
	Base        string
	HostID      string
	RackID      string
	DeviceID    string
	External    bool
	Usage       string
	AllocatedBy string
}

// Pool defines a pool of IP addresses.
type Pool struct {
	Fragments []IPRange
	Stats     IPPoolStats
}

// IPRange defines a range of IP addresses and associated subnet information.
// SourceBase is the Base address of the Source this IPRange was
// allocated from. It is used to prevent adjacent Source ranges from
// being coalesced in the Pool. We don't want to coalesce adjacent
// Source ranges because that would make it very difficult to edit the
// Sources data in the future.
type IPRange struct {
	Base         string
	Count        uint32
	DefaultRoute string
	Netmask      Netmask
	SourceBase   string
}

// IPPoolStats defines IP pool statistics.
type IPPoolStats struct {
	Total     uint32
	Available uint32
	InUse     uint32
}
