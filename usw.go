package unifi

import (
	"encoding/json"
	"time"
)

// USW represents all the data from the Ubiquiti Controller for a Unifi Switch.
type USW struct {
	site                  *Site
	AdoptableWhenUpgraded FlexBool         `json:"adoptable_when_upgraded,omitempty"`
	Adopted               FlexBool         `fake:"{constFlexBool:true}"              json:"adopted"`
	AdoptionCompleted     FlexBool         `json:"adoption_completed"`
	Anomalies             FlexInt          `json:"anomalies"`
	Architecture          string           `json:"architecture"`
	BoardRev              FlexInt          `json:"board_rev"`
	Bytes                 FlexInt          `json:"bytes"`
	Cfgversion            string           `json:"cfgversion"`
	ConfigNetwork         *ConfigNetwork   `json:"config_network"`
	ConnectRequestIP      string           `json:"connect_request_ip"`
	ConnectRequestPort    string           `json:"connect_request_port"`
	ConnectedAt           FlexInt          `json:"connected_at"`
	ConnectionNetworkName string           `json:"connection_network_name"`
	DeviceID              string           `json:"device_id"`
	DisplayableVersion    string           `json:"displayable_version"`
	Dot1XPortctrlEnabled  FlexBool         `json:"dot1x_portctrl_enabled"`
	DownlinkTable         []*DownlinkTable `json:"downlink_table"`
	EthernetTable         []*EthernetTable `json:"ethernet_table"`
	FanLevel              FlexInt          `json:"fan_level"`
	FlowctrlEnabled       FlexBool         `json:"flowctrl_enabled"`
	FwCaps                FlexInt          `json:"fw_caps"`
	GatewayMac            string           `json:"gateway_mac"`
	GeneralTemperature    FlexInt          `json:"general_temperature"`
	GuestNumSta           FlexInt          `json:"guest-num_sta"`
	HasFan                FlexBool         `json:"has_fan"`
	HasTemperature        FlexBool         `json:"has_temperature"`
	HwCaps                FlexInt          `json:"hw_caps"`
	ID                    string           `json:"_id"`
	IP                    string           `json:"ip"`
	InformIP              string           `json:"inform_ip"`
	InformURL             string           `json:"inform_url"`
	IsAccessPoint         FlexBool         `json:"is_access_point"`
	JumboframeEnabled     FlexBool         `json:"jumboframe_enabled"`
	KernelVersion         string           `json:"kernel_version"`
	KnownCfgversion       string           `json:"known_cfgversion"`
	LCMTrackerEnabled     FlexBool         `json:"lcm_tracker_enabled"`
	LastSeen              FlexInt          `json:"last_seen"`
	LastUplink            struct {
		UplinkMac string `json:"uplink_mac"`
	} `json:"last_uplink"`
	LedOverride              string   `json:"led_override"`
	LicenseState             string   `json:"license_state"`
	Locating                 FlexBool `fake:"{constFlexBool:false}"       json:"locating"`
	Mac                      string   `json:"mac"`
	ManufacturerID           FlexInt  `json:"manufacturer_id"`
	MinInformIntervalSeconds FlexInt  `json:"min_inform_interval_seconds"`
	Model                    string   `json:"model"`
	ModelInEOL               FlexBool `json:"model_in_eol"`
	ModelInLTS               FlexBool `json:"model_in_lts"`
	ModelIncompatible        FlexBool `json:"model_incompatible"`
	Name                     string   `fake:"{animal}"                    json:"name"`
	NextInterval             FlexInt  `json:"next_interval"`
	NumSta                   FlexInt  `json:"num_sta"`
	OutdoorModeOverride      string   `json:"outdoor_mode_override"`
	Overheating              FlexBool `json:"overheating"`
	PortOverrides            []struct {
		Name       string  `fake:"{randomstring:[override-1,override-2]}" json:"name,omitempty"`
		PoeMode    string  `json:"poe_mode,omitempty"`
		PortIdx    FlexInt `json:"port_idx"`
		PortconfID string  `json:"portconf_id"`
	} `json:"port_overrides"`
	PortTable               []Port      `json:"port_table"`
	PowerSource             string      `json:"power_source"`
	PowerSourceCtrlEnabled  FlexBool    `json:"power_source_ctrl_enabled"`
	PowerSourceVoltage      string      `json:"power_source_voltage"`
	PreviousNonBusyState    FlexInt     `json:"prev_non_busy_state"`
	ProvisionedAt           FlexInt     `json:"provisioned_at"`
	RequiredVersion         string      `json:"required_version"`
	Rollupgrade             FlexBool    `json:"rollupgrade,omitempty"`
	RxBytes                 FlexInt     `json:"rx_bytes"`
	Satisfaction            FlexInt     `json:"satisfaction"`
	Serial                  string      `json:"serial"`
	SetupID                 string      `json:"setup_id"`
	SiteID                  string      `json:"site_id"`
	SiteName                string      `json:"-"`
	SourceName              string      `json:"-"`
	StartConnectedMillis    FlexInt     `json:"start_connected_millis"`
	StartDisconnectedMillis FlexInt     `json:"start_disconnected_millis"`
	StartupTimestamp        FlexInt     `json:"startup_timestamp"`
	Stat                    USWStat     `json:"stat"`
	State                   FlexInt     `json:"state"`
	StpPriority             FlexInt     `json:"stp_priority"`
	StpVersion              string      `json:"stp_version"`
	SwitchCaps              *SwitchCaps `json:"switch_caps"`
	SysErrorCaps            FlexInt     `json:"sys_error_caps"`
	SysStats                SysStats    `json:"sys_stats"`
	SystemStats             SystemStats `json:"system-stats"`
	TotalMaxPower           FlexInt     `json:"total_max_power"`
	TwoPhaseAdopt           FlexBool    `json:"two_phase_adopt"`
	TxBytes                 FlexInt     `json:"tx_bytes"`
	Type                    string      `fake:"{randomstring:[usg,pdu]}"  json:"type"`
	Unsupported             FlexBool    `json:"unsupported"`
	UnsupportedReason       FlexInt     `json:"unsupported_reason"`
	Upgradable              FlexBool    `json:"upgradable,omitempty"`
	Upgradeable             FlexBool    `json:"upgradeable"`
	Uplink                  Uplink      `json:"uplink"`
	UplinkDepth             FlexInt     `json:"uplink_depth"`
	Uptime                  FlexInt     `json:"uptime"`
	UserNumSta              FlexInt     `json:"user-num_sta"`
	Version                 string      `json:"version"`
}

type SwitchCaps struct {
	FeatureCaps          FlexInt `json:"feature_caps"`
	MaxMirrorSessions    FlexInt `json:"max_mirror_sessions"`
	MaxAggregateSessions FlexInt `json:"max_aggregate_sessions"`
}

// MacTable is a newer feature on some switched ports.
type MacTable struct {
	Age           int64    `json:"age"`
	Authorized    FlexBool `json:"authorized"`
	Hostname      string   `json:"hostname"`
	IP            string   `json:"ip"`
	LastReachable int64    `json:"lastReachable"`
	Mac           string   `json:"mac"`
}

// PortDelta is part of a Port.
type PortDelta struct {
	TimeDelta         FlexInt `json:"time_delta"`
	TimeDeltaActivity FlexInt `json:"time_delta_activity"`
}

// USWStat holds the "stat" data for a switch.
// This is split out because of a JSON data format change from 5.10 to 5.11.
type USWStat struct {
	*Sw
}

// Sw is a subtype of USWStat to make unmarshalling of different controller versions possible.
type Sw struct {
	Bytes       FlexInt   `json:"bytes"`
	Datetime    time.Time `fake:"{recent_time}" json:"datetime"`
	Duration    FlexInt   `json:"duration"`
	O           string    `json:"o"`
	Oid         string    `json:"oid"`
	RxBroadcast FlexInt   `json:"rx_broadcast"`
	RxBytes     FlexInt   `json:"rx_bytes"`
	RxCrypts    FlexInt   `json:"rx_crypts"`
	RxDropped   FlexInt   `json:"rx_dropped"`
	RxErrors    FlexInt   `json:"rx_errors"`
	RxFrags     FlexInt   `json:"rx_frags"`
	RxMulticast FlexInt   `json:"rx_multicast"`
	RxPackets   FlexInt   `json:"rx_packets"`
	SiteID      string    `fake:"{uuid}"        json:"site_id"`
	Sw          string    `json:"sw"`
	Time        FlexInt   `json:"time"`
	TxBroadcast FlexInt   `json:"tx_broadcast"`
	TxBytes     FlexInt   `json:"tx_bytes"`
	TxDropped   FlexInt   `json:"tx_dropped"`
	TxErrors    FlexInt   `json:"tx_errors"`
	TxMulticast FlexInt   `json:"tx_multicast"`
	TxPackets   FlexInt   `json:"tx_packets"`
	TxRetries   FlexInt   `json:"tx_retries"`
	/* These are all in port table */
	/*
		Port1RxPackets    FlexInt   `json:"port_1-rx_packets,omitempty"`
		Port1RxBytes      FlexInt   `json:"port_1-rx_bytes,omitempty"`
		Port1TxPackets    FlexInt   `json:"port_1-tx_packets,omitempty"`
		Port1TxBytes      FlexInt   `json:"port_1-tx_bytes,omitempty"`
		Port1TxMulticast  FlexInt   `json:"port_1-tx_multicast"`
		Port1TxBroadcast  FlexInt   `json:"port_1-tx_broadcast"`
		Port3RxPackets    FlexInt   `json:"port_3-rx_packets,omitempty"`
		Port3RxBytes      FlexInt   `json:"port_3-rx_bytes,omitempty"`
		Port3TxPackets    FlexInt   `json:"port_3-tx_packets,omitempty"`
		Port3TxBytes      FlexInt   `json:"port_3-tx_bytes,omitempty"`
		Port3RxBroadcast  FlexInt   `json:"port_3-rx_broadcast"`
		Port3TxMulticast  FlexInt   `json:"port_3-tx_multicast"`
		Port3TxBroadcast  FlexInt   `json:"port_3-tx_broadcast"`
		Port6RxPackets    FlexInt   `json:"port_6-rx_packets,omitempty"`
		Port6RxBytes      FlexInt   `json:"port_6-rx_bytes,omitempty"`
		Port6TxPackets    FlexInt   `json:"port_6-tx_packets,omitempty"`
		Port6TxBytes      FlexInt   `json:"port_6-tx_bytes,omitempty"`
		Port6RxMulticast  FlexInt   `json:"port_6-rx_multicast"`
		Port6TxMulticast  FlexInt   `json:"port_6-tx_multicast"`
		Port6TxBroadcast  FlexInt   `json:"port_6-tx_broadcast"`
		Port7RxPackets    FlexInt   `json:"port_7-rx_packets,omitempty"`
		Port7RxBytes      FlexInt   `json:"port_7-rx_bytes,omitempty"`
		Port7TxPackets    FlexInt   `json:"port_7-tx_packets,omitempty"`
		Port7TxBytes      FlexInt   `json:"port_7-tx_bytes,omitempty"`
		Port7TxMulticast  FlexInt   `json:"port_7-tx_multicast"`
		Port7TxBroadcast  FlexInt   `json:"port_7-tx_broadcast"`
		Port9RxPackets    FlexInt   `json:"port_9-rx_packets,omitempty"`
		Port9RxBytes      FlexInt   `json:"port_9-rx_bytes,omitempty"`
		Port9TxPackets    FlexInt   `json:"port_9-tx_packets,omitempty"`
		Port9TxBytes      FlexInt   `json:"port_9-tx_bytes,omitempty"`
		Port9TxMulticast  FlexInt   `json:"port_9-tx_multicast"`
		Port9TxBroadcast  FlexInt   `json:"port_9-tx_broadcast"`
		Port10RxPackets   FlexInt   `json:"port_10-rx_packets,omitempty"`
		Port10RxBytes     FlexInt   `json:"port_10-rx_bytes,omitempty"`
		Port10TxPackets   FlexInt   `json:"port_10-tx_packets,omitempty"`
		Port10TxBytes     FlexInt   `json:"port_10-tx_bytes,omitempty"`
		Port10RxMulticast FlexInt   `json:"port_10-rx_multicast"`
		Port10TxMulticast FlexInt   `json:"port_10-tx_multicast"`
		Port10TxBroadcast FlexInt   `json:"port_10-tx_broadcast"`
		Port11RxPackets   FlexInt   `json:"port_11-rx_packets,omitempty"`
		Port11RxBytes     FlexInt   `json:"port_11-rx_bytes,omitempty"`
		Port11TxPackets   FlexInt   `json:"port_11-tx_packets,omitempty"`
		Port11TxBytes     FlexInt   `json:"port_11-tx_bytes,omitempty"`
		Port11TxMulticast FlexInt   `json:"port_11-tx_multicast"`
		Port11TxBroadcast FlexInt   `json:"port_11-tx_broadcast"`
		Port12RxPackets   FlexInt   `json:"port_12-rx_packets,omitempty"`
		Port12RxBytes     FlexInt   `json:"port_12-rx_bytes,omitempty"`
		Port12TxPackets   FlexInt   `json:"port_12-tx_packets,omitempty"`
		Port12TxBytes     FlexInt   `json:"port_12-tx_bytes,omitempty"`
		Port12TxMulticast FlexInt   `json:"port_12-tx_multicast"`
		Port12TxBroadcast FlexInt   `json:"port_12-tx_broadcast"`
		Port13RxPackets   FlexInt   `json:"port_13-rx_packets,omitempty"`
		Port13RxBytes     FlexInt   `json:"port_13-rx_bytes,omitempty"`
		Port13TxPackets   FlexInt   `json:"port_13-tx_packets,omitempty"`
		Port13TxBytes     FlexInt   `json:"port_13-tx_bytes,omitempty"`
		Port13RxMulticast FlexInt   `json:"port_13-rx_multicast"`
		Port13RxBroadcast FlexInt   `json:"port_13-rx_broadcast"`
		Port13TxMulticast FlexInt   `json:"port_13-tx_multicast"`
		Port13TxBroadcast FlexInt   `json:"port_13-tx_broadcast"`
		Port15RxPackets   FlexInt   `json:"port_15-rx_packets,omitempty"`
		Port15RxBytes     FlexInt   `json:"port_15-rx_bytes,omitempty"`
		Port15TxPackets   FlexInt   `json:"port_15-tx_packets,omitempty"`
		Port15TxBytes     FlexInt   `json:"port_15-tx_bytes,omitempty"`
		Port15RxBroadcast FlexInt   `json:"port_15-rx_broadcast"`
		Port15TxMulticast FlexInt   `json:"port_15-tx_multicast"`
		Port15TxBroadcast FlexInt   `json:"port_15-tx_broadcast"`
		Port16RxPackets   FlexInt   `json:"port_16-rx_packets,omitempty"`
		Port16RxBytes     FlexInt   `json:"port_16-rx_bytes,omitempty"`
		Port16TxPackets   FlexInt   `json:"port_16-tx_packets,omitempty"`
		Port16TxBytes     FlexInt   `json:"port_16-tx_bytes,omitempty"`
		Port16TxMulticast FlexInt   `json:"port_16-tx_multicast"`
		Port16TxBroadcast FlexInt   `json:"port_16-tx_broadcast"`
		Port17RxPackets   FlexInt   `json:"port_17-rx_packets,omitempty"`
		Port17RxBytes     FlexInt   `json:"port_17-rx_bytes,omitempty"`
		Port17TxPackets   FlexInt   `json:"port_17-tx_packets,omitempty"`
		Port17TxBytes     FlexInt   `json:"port_17-tx_bytes,omitempty"`
		Port17TxMulticast FlexInt   `json:"port_17-tx_multicast"`
		Port17TxBroadcast FlexInt   `json:"port_17-tx_broadcast"`
		Port18RxPackets   FlexInt   `json:"port_18-rx_packets,omitempty"`
		Port18RxBytes     FlexInt   `json:"port_18-rx_bytes,omitempty"`
		Port18TxPackets   FlexInt   `json:"port_18-tx_packets,omitempty"`
		Port18TxBytes     FlexInt   `json:"port_18-tx_bytes,omitempty"`
		Port18RxMulticast FlexInt   `json:"port_18-rx_multicast"`
		Port18TxMulticast FlexInt   `json:"port_18-tx_multicast"`
		Port18TxBroadcast FlexInt   `json:"port_18-tx_broadcast"`
		Port19RxPackets   FlexInt   `json:"port_19-rx_packets,omitempty"`
		Port19RxBytes     FlexInt   `json:"port_19-rx_bytes,omitempty"`
		Port19TxPackets   FlexInt   `json:"port_19-tx_packets,omitempty"`
		Port19TxBytes     FlexInt   `json:"port_19-tx_bytes,omitempty"`
		Port19TxMulticast FlexInt   `json:"port_19-tx_multicast"`
		Port19TxBroadcast FlexInt   `json:"port_19-tx_broadcast"`
		Port21RxPackets   FlexInt   `json:"port_21-rx_packets,omitempty"`
		Port21RxBytes     FlexInt   `json:"port_21-rx_bytes,omitempty"`
		Port21TxPackets   FlexInt   `json:"port_21-tx_packets,omitempty"`
		Port21TxBytes     FlexInt   `json:"port_21-tx_bytes,omitempty"`
		Port21RxBroadcast FlexInt   `json:"port_21-rx_broadcast"`
		Port21TxMulticast FlexInt   `json:"port_21-tx_multicast"`
		Port21TxBroadcast FlexInt   `json:"port_21-tx_broadcast"`
		Port22RxPackets   FlexInt   `json:"port_22-rx_packets,omitempty"`
		Port22RxBytes     FlexInt   `json:"port_22-rx_bytes,omitempty"`
		Port22TxPackets   FlexInt   `json:"port_22-tx_packets,omitempty"`
		Port22TxBytes     FlexInt   `json:"port_22-tx_bytes,omitempty"`
		Port22RxMulticast FlexInt   `json:"port_22-rx_multicast"`
		Port22TxMulticast FlexInt   `json:"port_22-tx_multicast"`
		Port22TxBroadcast FlexInt   `json:"port_22-tx_broadcast"`
		Port23RxPackets   FlexInt   `json:"port_23-rx_packets,omitempty"`
		Port23RxBytes     FlexInt   `json:"port_23-rx_bytes,omitempty"`
		Port23RxDropped   FlexInt   `json:"port_23-rx_dropped"`
		Port23TxPackets   FlexInt   `json:"port_23-tx_packets,omitempty"`
		Port23TxBytes     FlexInt   `json:"port_23-tx_bytes,omitempty"`
		Port23RxMulticast FlexInt   `json:"port_23-rx_multicast"`
		Port23RxBroadcast FlexInt   `json:"port_23-rx_broadcast"`
		Port23TxMulticast FlexInt   `json:"port_23-tx_multicast"`
		Port23TxBroadcast FlexInt   `json:"port_23-tx_broadcast"`
		Port24RxPackets   FlexInt   `json:"port_24-rx_packets,omitempty"`
		Port24RxBytes     FlexInt   `json:"port_24-rx_bytes,omitempty"`
		Port24TxPackets   FlexInt   `json:"port_24-tx_packets,omitempty"`
		Port24TxBytes     FlexInt   `json:"port_24-tx_bytes,omitempty"`
		Port24RxMulticast FlexInt   `json:"port_24-rx_multicast"`
		Port24TxMulticast FlexInt   `json:"port_24-tx_multicast"`
		Port24TxBroadcast FlexInt   `json:"port_24-tx_broadcast"`
		Port1RxMulticast  FlexInt   `json:"port_1-rx_multicast"`
		Port3RxDropped    FlexInt   `json:"port_3-rx_dropped"`
		Port3RxMulticast  FlexInt   `json:"port_3-rx_multicast"`
		Port6RxDropped    FlexInt   `json:"port_6-rx_dropped"`
		Port7RxDropped    FlexInt   `json:"port_7-rx_dropped"`
		Port7RxMulticast  FlexInt   `json:"port_7-rx_multicast"`
		Port9RxDropped    FlexInt   `json:"port_9-rx_dropped"`
		Port9RxMulticast  FlexInt   `json:"port_9-rx_multicast"`
		Port9RxBroadcast  FlexInt   `json:"port_9-rx_broadcast"`
		Port10RxBroadcast FlexInt   `json:"port_10-rx_broadcast"`
		Port12RxDropped   FlexInt   `json:"port_12-rx_dropped"`
		Port12RxMulticast FlexInt   `json:"port_12-rx_multicast"`
		Port13RxDropped   FlexInt   `json:"port_13-rx_dropped"`
		Port17RxDropped   FlexInt   `json:"port_17-rx_dropped"`
		Port17RxMulticast FlexInt   `json:"port_17-rx_multicast"`
		Port17RxBroadcast FlexInt   `json:"port_17-rx_broadcast"`
		Port19RxDropped   FlexInt   `json:"port_19-rx_dropped"`
		Port19RxMulticast FlexInt   `json:"port_19-rx_multicast"`
		Port19RxBroadcast FlexInt   `json:"port_19-rx_broadcast"`
		Port21RxDropped   FlexInt   `json:"port_21-rx_dropped"`
		Port21RxMulticast FlexInt   `json:"port_21-rx_multicast"`
		Port7RxBroadcast  FlexInt   `json:"port_7-rx_broadcast"`
		Port18RxBroadcast FlexInt   `json:"port_18-rx_broadcast"`
		Port16RxMulticast FlexInt   `json:"port_16-rx_multicast"`
		Port15RxDropped   FlexInt   `json:"port_15-rx_dropped"`
		Port15RxMulticast FlexInt   `json:"port_15-rx_multicast"`
		Port16RxBroadcast FlexInt   `json:"port_16-rx_broadcast"`
		Port11RxBroadcast FlexInt   `json:"port_11-rx_broadcast"`
		Port12RxBroadcast FlexInt   `json:"port_12-rx_broadcast"`
		Port6RxBroadcast  FlexInt   `json:"port_6-rx_broadcast"`
		Port24RxBroadcast FlexInt   `json:"port_24-rx_broadcast"`
		Port22RxBroadcast FlexInt   `json:"port_22-rx_broadcast"`
		Port10TxDropped   FlexInt   `json:"port_10-tx_dropped"`
		Port16TxDropped   FlexInt   `json:"port_16-tx_dropped"`
		Port1RxBroadcast  FlexInt   `json:"port_1-rx_broadcast"`
		Port4RxPackets    FlexInt   `json:"port_4-rx_packets,omitempty"`
		Port4RxBytes      FlexInt   `json:"port_4-rx_bytes,omitempty"`
		Port4RxDropped    FlexInt   `json:"port_4-rx_dropped"`
		Port4TxPackets    FlexInt   `json:"port_4-tx_packets,omitempty"`
		Port4TxBytes      FlexInt   `json:"port_4-tx_bytes,omitempty"`
		Port4TxDropped    FlexInt   `json:"port_4-tx_dropped"`
		Port4RxMulticast  FlexInt   `json:"port_4-rx_multicast"`
		Port4RxBroadcast  FlexInt   `json:"port_4-rx_broadcast"`
		Port4TxMulticast  FlexInt   `json:"port_4-tx_multicast"`
		Port4TxBroadcast  FlexInt   `json:"port_4-tx_broadcast"`
	*/
}

// UnmarshalJSON unmarshalls 5.10 or 5.11 formatted Switch Stat data.
func (v *USWStat) UnmarshalJSON(data []byte) error {
	var n struct {
		Sw `json:"sw"`
	}

	v.Sw = &n.Sw

	err := json.Unmarshal(data, v.Sw) // controller version 5.10.
	if err != nil {
		return json.Unmarshal(data, &n) // controller version 5.11.
	}

	return nil
}
