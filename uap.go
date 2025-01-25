package unifi

import (
	"encoding/json"
	"fmt"
	"time"
)

// UAP represents all the data from the Ubiquiti Controller for a Unifi Access Point.
// This was auto generated then edited by hand to get all the data types right.
type UAP struct {
	site                  *Site
	AFTEnabled            FlexBool `json:"atf_enabled"`
	AdoptableWhenUpgraded FlexBool `json:"adoptable_when_upgraded"`
	Adopted               FlexBool `fake:"{constFlexBool:true}"    json:"adopted"`
	AntennaTable          []struct {
		Default   FlexBool `json:"default"`
		ID        FlexInt  `json:"id"`
		Name      string   `fake:"{animal}"   json:"name"`
		Wifi0Gain FlexInt  `json:"wifi0_gain"`
		Wifi1Gain FlexInt  `json:"wifi1_gain"`
	} `fakesize:"5"                       json:"antenna_table"`
	Architecture         string           `json:"architecture"`
	BandsteeringMode     string           `json:"bandsteering_mode,omitempty"`
	BoardRev             int              `json:"board_rev"`
	Bytes                FlexInt          `json:"bytes"`
	BytesD               FlexInt          `json:"bytes-d"`
	BytesR               FlexInt          `json:"bytes-r"`
	Cfgversion           string           `fake:"{appversion}"                json:"cfgversion"`
	ConfigNetwork        *ConfigNetwork   `json:"config_network"`
	ConnectRequestIP     string           `fake:"{ipv4address}"               json:"connect_request_ip"`
	ConnectRequestPort   string           `json:"connect_request_port"`
	ConnectedAt          FlexInt          `json:"connected_at"`
	CountryCode          FlexInt          `json:"country_code"`
	CountrycodeTable     []int            `fakesize:"5"                       json:"countrycode_table"`
	DeviceID             string           `fake:"{uuid}"                      json:"device_id"`
	Dot1XPortctrlEnabled FlexBool         `json:"dot1x_portctrl_enabled"`
	DownlinkTable        []*DownlinkTable `fakesize:"5"                       json:"downlink_table,omitempty"`
	EthernetTable        []*EthernetTable `fakesize:"5"                       json:"ethernet_table"`
	FixedAPAvailable     FlexBool         `json:"fixed_ap_available"`
	FwCaps               int              `json:"fw_caps"`
	GatewayMac           string           `fake:"{macaddress}"                json:"gateway_mac"`
	GuestNumSta          FlexInt          `json:"guest-num_sta"`
	GuestToken           string           `json:"guest_token"`
	GuestWlanNumSta      FlexInt          `json:"guest-wlan-num_sta"`
	HasEth1              FlexBool         `json:"has_eth1"`
	HasFan               FlexBool         `json:"has_fan"`
	HasSpeaker           FlexBool         `json:"has_speaker"`
	HasTemperature       FlexBool         `json:"has_temperature"`
	HwCaps               int              `json:"hw_caps"`
	ID                   string           `fake:"{uuid}"                      json:"_id"`
	IP                   string           `fake:"{ipv4address}"               json:"ip"`
	InformIP             string           `fake:"{ipv4address}"               json:"inform_ip"`
	InformURL            string           `fake:"{url}"                       json:"inform_url"`
	Internet             FlexBool         `json:"internet"`
	Isolated             FlexBool         `json:"isolated"`
	KernelVersion        string           `fake:"{appversion}"                json:"kernel_version"`
	KnownCfgversion      string           `fake:"{appversion}"                json:"known_cfgversion"`
	LastSeen             FlexInt          `json:"last_seen"`
	LastUplink           struct {
		UplinkMac        string `fake:"{macaddress}" json:"uplink_mac"`
		UplinkRemotePort int    `fake:"{port}"       json:"uplink_remote_port"`
	} `json:"last_uplink"`
	LcmBrightnessOverride         FlexBool        `json:"lcm_brightness_override"`
	LcmTimeoutOverride            FlexBool        `json:"lcm_idle_timeout_override"`
	LcmTrackerEnabled             FlexBool        `json:"lcm_tracker_enabled"`
	LcmTrackerSeed                string          `json:"lcm_tracker_seed"`
	LedOverride                   string          `json:"led_override"`
	LedOverrideColor              string          `json:"led_override_color"`
	LedOverrideColorBrightness    FlexInt         `json:"led_override_color_brightness"`
	LicenseState                  string          `json:"license_state"`
	Locating                      FlexBool        `fake:"{constFlexBool:false}"            json:"locating"`
	LteAPN                        string          `json:"lte_apn"`
	LteBand                       string          `json:"lte_band"`
	LteCellID                     string          `json:"lte_cell_id"`
	LteConnected                  FlexBool        `json:"lte_connected"`
	LteDataLimitEnabled           FlexBool        `json:"lte_data_limit_enabled"`
	LteExtAntenna                 FlexBool        `json:"lte_ext_ant"`
	LteFailover                   FlexBool        `json:"lte_failover"`
	LteFailoverMode               string          `json:"lte_failover_mode"`
	LteHardLimit                  FlexInt         `json:"lte_hard_limit"`
	LteICCID                      string          `json:"lte_iccid"`
	LteIMEI                       string          `json:"lte_imei"`
	LteIP                         string          `fake:"{ipv4address}"                    json:"lte_ip"`
	LteMode                       string          `json:"lte_mode"`
	LteNetworkOperator            string          `json:"lte_networkoperator"`
	LtePdpType                    string          `json:"lte_pdptype"`
	LtePoe                        FlexBool        `json:"lte_poe"`
	LteRadio                      string          `json:"lte_radio"`
	LteRadioMode                  string          `json:"lte_radio_mode"`
	LteRat                        string          `json:"lte_rat"`
	LteRsrp                       FlexInt         `json:"lte_rsrp"`
	LteRsrq                       FlexInt         `json:"lte_rsrq"`
	LteRssi                       FlexInt         `json:"lte_rssi"`
	LteRxChannel                  FlexInt         `json:"lte_rx_chan"`
	LteSignal                     string          `json:"lte_signal"`
	LteSoftLimit                  FlexInt         `json:"lte_soft_limit"`
	LteSsoUserUuid                string          `json:"lte_sso_user_uuid"` //nolint:revive
	LteState                      FlexBool        `json:"lte_state"`
	LteSubscriptionApiUnreachable FlexBool        `json:"lte_subscription_api_unreachable"` //nolint:revive
	LteSubscriptionCheckRequired  FlexBool        `json:"lte_subscription_check_required"`
	LteSubscriptionStatus         FlexBool        `json:"lte_subscription_status"`
	LteTxChannel                  FlexInt         `json:"lte_tx_chan"`
	Mac                           string          `fake:"{macaddress}"                     json:"mac"`
	ManufacturerID                FlexInt         `json:"manufacturer_id"`
	MeshStaVapEnabled             FlexBool        `json:"mesh_sta_vap_enabled"`
	Meshv3PeerMac                 string          `json:"meshv3_peer_mac"`
	Model                         string          `json:"model"`
	ModelInEOL                    FlexBool        `json:"model_in_eol"`
	ModelInLTS                    FlexBool        `json:"model_in_lts"`
	ModelIncompatible             FlexBool        `json:"model_incompatible"`
	Name                          string          `fake:"{randomstring:[uap-1,uap-2]}"     json:"name"`
	NumSta                        FlexInt         `json:"num_sta"`
	OutdoorModeOverride           string          `json:"outdoor_mode_override"`
	PortTable                     []Port          `fakesize:"5"                            json:"port_table"`
	ProvisionedAt                 FlexInt         `json:"provisioned_at"`
	RadioTable                    RadioTable      `json:"radio_table"`
	RadioTableStats               RadioTableStats `json:"radio_table_stats"`
	RequiredVersion               string          `fake:"{appversion}"                     json:"required_version"`
	Rollupgrade                   FlexBool        `json:"rollupgrade"`
	RxBytes                       FlexInt         `json:"rx_bytes"`
	RxBytesD                      FlexInt         `json:"rx_bytes-d"`
	SSHSessionTable               []interface{}   `json:"ssh_session_table"`
	Satisfaction                  FlexInt         `json:"satisfaction"`
	ScanRadioTable                []interface{}   `json:"scan_radio_table"`
	Scanning                      FlexBool        `json:"scanning"`
	Serial                        string          `json:"serial"`
	SiteID                        string          `fake:"{uuid}"                           json:"site_id"`
	SiteName                      string          `json:"-"`
	SourceName                    string          `json:"-"`
	SpectrumScanning              FlexBool        `json:"spectrum_scanning"`
	StartConnectedMillis          FlexInt         `json:"start_connected_millis"`
	StartDisconnectedMillis       FlexInt         `json:"start_disconnected_millis"`
	StartupTimestamp              FlexInt         `json:"startup_timestamp"`
	Stat                          UAPStat         `json:"stat"`
	State                         FlexInt         `json:"state"`
	SupportsFingerprintML         FlexBool        `json:"supports_fingerprint_ml"`
	SysErrorCaps                  int             `json:"sys_error_caps"`
	SysStats                      SysStats        `json:"sys_stats"`
	SyslogKey                     string          `json:"syslog_key"`
	SystemStats                   SystemStats     `json:"system-stats"`
	TotalRxBytes                  FlexInt         `json:"total_rx_bytes"`
	TotalTxBytes                  FlexInt         `json:"total_tx_bytes"`
	TwoPhaseAdopt                 FlexBool        `json:"two_phase_adopt,omitempty"`
	TxBytes                       FlexInt         `json:"tx_bytes"`
	TxBytesD                      FlexInt         `json:"tx_bytes-d"`
	Type                          string          `fake:"{lexify:uap}"                     json:"type"`
	UUptime                       FlexInt         `json:"_uptime"`
	Unsupported                   FlexBool        `json:"unsupported"`
	UnsupportedReason             FlexInt         `json:"unsupported_reason"`
	Upgradable                    FlexBool        `json:"upgradable"`
	UpgradeState                  FlexInt         `json:"upgrade_state"`
	Uplink                        struct {
		FullDuplex       FlexBool `json:"full_duplex"`
		IP               string   `fake:"{ipv4address}" json:"ip"`
		Mac              string   `fake:"{macaddress}"  json:"mac"`
		MaxVlan          int      `json:"max_vlan"`
		Name             string   `json:"name"`
		Netmask          string   `json:"netmask"`
		NumPort          int      `fake:"{port}"        json:"num_port"`
		RxBytes          FlexInt  `json:"rx_bytes"`
		RxDropped        FlexInt  `json:"rx_dropped"`
		RxErrors         FlexInt  `json:"rx_errors"`
		RxMulticast      FlexInt  `json:"rx_multicast"`
		RxPackets        FlexInt  `json:"rx_packets"`
		Speed            FlexInt  `json:"speed"`
		TxBytes          FlexInt  `json:"tx_bytes"`
		TxDropped        FlexInt  `json:"tx_dropped"`
		TxErrors         FlexInt  `json:"tx_errors"`
		TxPackets        FlexInt  `json:"tx_packets"`
		Up               FlexBool `json:"up"`
		MaxSpeed         FlexInt  `json:"max_speed"`
		Type             string   `json:"type"`
		TxBytesR         FlexInt  `json:"tx_bytes-r"`
		RxBytesR         FlexInt  `json:"rx_bytes-r"`
		UplinkMac        string   `fake:"{macaddress}"  json:"uplink_mac"`
		UplinkRemotePort int      `fake:"{port}"        json:"uplink_remote_port"`
	} `json:"uplink"`
	UplinkTable    []interface{} `json:"uplink_table"`
	Uptime         FlexInt       `json:"uptime"`
	UserNumSta     FlexInt       `json:"user-num_sta"`
	UserWlanNumSta FlexInt       `json:"user-wlan-num_sta"`
	VapTable       VapTable      `json:"vap_table"`
	Version        string        `fake:"{appversion}"      json:"version"`
	VwireEnabled   FlexBool      `json:"vwireEnabled"`
	VwireTable     []interface{} `json:"vwire_table"`
	VwireVapTable  []interface{} `json:"vwire_vap_table"`
	WifiCaps       int           `json:"wifi_caps"`
	WlangroupIDNa  string        `json:"wlangroup_id_na"`
	WlangroupIDNg  string        `json:"wlangroup_id_ng"`
}

// UAPStat holds the "stat" data for an access point.
// This is split out because of a JSON data format change from 5.10 to 5.11.
type UAPStat struct {
	*Ap
}

// Ap is a subtype of UAPStat to make unmarshalling of different controller versions possible.
type Ap struct {
	Ap                       string    `json:"ap"`
	Bytes                    FlexInt   `json:"bytes"`
	Datetime                 time.Time `fake:"{recent_time}"               json:"datetime"`
	Duration                 FlexInt   `json:"duration"`
	GuestMacFilterRejections FlexInt   `json:"guest-mac_filter_rejections"`
	GuestRxBytes             FlexInt   `json:"guest-rx_bytes"`
	GuestRxCrypts            FlexInt   `json:"guest-rx_crypts"`
	GuestRxDropped           FlexInt   `json:"guest-rx_dropped"`
	GuestRxErrors            FlexInt   `json:"guest-rx_errors"`
	GuestRxFrags             FlexInt   `json:"guest-rx_frags"`
	GuestRxPackets           FlexInt   `json:"guest-rx_packets"`
	GuestTxBytes             FlexInt   `json:"guest-tx_bytes"`
	GuestTxDropped           FlexInt   `json:"guest-tx_dropped"`
	GuestTxErrors            FlexInt   `json:"guest-tx_errors"`
	GuestTxPackets           FlexInt   `json:"guest-tx_packets"`
	GuestTxRetries           FlexInt   `json:"guest-tx_retries"`
	GuestWifiTxAttempts      FlexInt   `json:"guest-wifi_tx_attempts"`
	GuestWifiTxDropped       FlexInt   `json:"guest-wifi_tx_dropped"`
	MacFilterRejections      FlexInt   `json:"mac_filter_rejections"`
	O                        string    `json:"o"`
	Oid                      string    `json:"oid"`
	RxBytes                  FlexInt   `json:"rx_bytes"`
	RxCrypts                 FlexInt   `json:"rx_crypts"`
	RxDropped                FlexInt   `json:"rx_dropped"`
	RxErrors                 FlexInt   `json:"rx_errors"`
	RxFrags                  FlexInt   `json:"rx_frags"`
	RxPackets                FlexInt   `json:"rx_packets"`
	SiteID                   string    `fake:"{uuid}"                      json:"site_id"`
	Time                     FlexInt   `json:"time"`
	TxBytes                  FlexInt   `json:"tx_bytes"`
	TxDropped                FlexInt   `json:"tx_dropped"`
	TxErrors                 FlexInt   `json:"tx_errors"`
	TxPackets                FlexInt   `json:"tx_packets"`
	TxRetries                FlexInt   `json:"tx_retries"`
	UserMacFilterRejections  FlexInt   `json:"user-mac_filter_rejections"`
	UserRxBytes              FlexInt   `json:"user-rx_bytes"`
	UserRxCrypts             FlexInt   `json:"user-rx_crypts"`
	UserRxDropped            FlexInt   `json:"user-rx_dropped"`
	UserRxErrors             FlexInt   `json:"user-rx_errors"`
	UserRxFrags              FlexInt   `json:"user-rx_frags"`
	UserRxPackets            FlexInt   `json:"user-rx_packets"`
	UserTxBytes              FlexInt   `json:"user-tx_bytes"`
	UserTxDropped            FlexInt   `json:"user-tx_dropped"`
	UserTxErrors             FlexInt   `json:"user-tx_errors"`
	UserTxPackets            FlexInt   `json:"user-tx_packets"`
	UserTxRetries            FlexInt   `json:"user-tx_retries"`
	UserWifiTxAttempts       FlexInt   `json:"user-wifi_tx_attempts"`
	UserWifiTxDropped        FlexInt   `json:"user-wifi_tx_dropped"`
	WifiTxAttempts           FlexInt   `json:"wifi_tx_attempts"`
	WifiTxDropped            FlexInt   `json:"wifi_tx_dropped"`

	// UAP-AC-PRO names, others may differ.
	/* These are all in VAP TABLE */
	GuestWifi0MacFilterRejections FlexInt `json:"guest-wifi0-mac_filter_rejections"`
	GuestWifi0RxBytes             FlexInt `json:"guest-wifi0-rx_bytes"`
	GuestWifi0RxCrypts            FlexInt `json:"guest-wifi0-rx_crypts"`
	GuestWifi0RxDropped           FlexInt `json:"guest-wifi0-rx_dropped"`
	GuestWifi0RxErrors            FlexInt `json:"guest-wifi0-rx_errors"`
	GuestWifi0RxFrags             FlexInt `json:"guest-wifi0-rx_frags"`
	GuestWifi0RxPackets           FlexInt `json:"guest-wifi0-rx_packets"`
	GuestWifi0TxBytes             FlexInt `json:"guest-wifi0-tx_bytes"`
	GuestWifi0TxDropped           FlexInt `json:"guest-wifi0-tx_dropped"`
	GuestWifi0TxErrors            FlexInt `json:"guest-wifi0-tx_errors"`
	GuestWifi0TxPackets           FlexInt `json:"guest-wifi0-tx_packets"`
	GuestWifi0TxRetries           FlexInt `json:"guest-wifi0-tx_retries"`
	GuestWifi0WifiTxAttempts      FlexInt `json:"guest-wifi0-wifi_tx_attempts"`
	GuestWifi0WifiTxDropped       FlexInt `json:"guest-wifi0-wifi_tx_dropped"`
	GuestWifi1MacFilterRejections FlexInt `json:"guest-wifi1-mac_filter_rejections"`
	GuestWifi1RxBytes             FlexInt `json:"guest-wifi1-rx_bytes"`
	GuestWifi1RxCrypts            FlexInt `json:"guest-wifi1-rx_crypts"`
	GuestWifi1RxDropped           FlexInt `json:"guest-wifi1-rx_dropped"`
	GuestWifi1RxErrors            FlexInt `json:"guest-wifi1-rx_errors"`
	GuestWifi1RxFrags             FlexInt `json:"guest-wifi1-rx_frags"`
	GuestWifi1RxPackets           FlexInt `json:"guest-wifi1-rx_packets"`
	GuestWifi1TxBytes             FlexInt `json:"guest-wifi1-tx_bytes"`
	GuestWifi1TxDropped           FlexInt `json:"guest-wifi1-tx_dropped"`
	GuestWifi1TxErrors            FlexInt `json:"guest-wifi1-tx_errors"`
	GuestWifi1TxPackets           FlexInt `json:"guest-wifi1-tx_packets"`
	GuestWifi1TxRetries           FlexInt `json:"guest-wifi1-tx_retries"`
	GuestWifi1WifiTxAttempts      FlexInt `json:"guest-wifi1-wifi_tx_attempts"`
	GuestWifi1WifiTxDropped       FlexInt `json:"guest-wifi1-wifi_tx_dropped"`
	UserWifi0MacFilterRejections  FlexInt `json:"user-wifi0-mac_filter_rejections"`
	UserWifi0RxBytes              FlexInt `json:"user-wifi0-rx_bytes"`
	UserWifi0RxCrypts             FlexInt `json:"user-wifi0-rx_crypts"`
	UserWifi0RxDropped            FlexInt `json:"user-wifi0-rx_dropped"`
	UserWifi0RxErrors             FlexInt `json:"user-wifi0-rx_errors"`
	UserWifi0RxFrags              FlexInt `json:"user-wifi0-rx_frags"`
	UserWifi0RxPackets            FlexInt `json:"user-wifi0-rx_packets"`
	UserWifi0TxBytes              FlexInt `json:"user-wifi0-tx_bytes"`
	UserWifi0TxDropped            FlexInt `json:"user-wifi0-tx_dropped"`
	UserWifi0TxErrors             FlexInt `json:"user-wifi0-tx_errors"`
	UserWifi0TxPackets            FlexInt `json:"user-wifi0-tx_packets"`
	UserWifi0TxRetries            FlexInt `json:"user-wifi0-tx_retries"`
	UserWifi0WifiTxAttempts       FlexInt `json:"user-wifi0-wifi_tx_attempts"`
	UserWifi0WifiTxDropped        FlexInt `json:"user-wifi0-wifi_tx_dropped"`
	UserWifi1MacFilterRejections  FlexInt `json:"user-wifi1-mac_filter_rejections"`
	UserWifi1RxBytes              FlexInt `json:"user-wifi1-rx_bytes"`
	UserWifi1RxCrypts             FlexInt `json:"user-wifi1-rx_crypts"`
	UserWifi1RxDropped            FlexInt `json:"user-wifi1-rx_dropped"`
	UserWifi1RxErrors             FlexInt `json:"user-wifi1-rx_errors"`
	UserWifi1RxFrags              FlexInt `json:"user-wifi1-rx_frags"`
	UserWifi1RxPackets            FlexInt `json:"user-wifi1-rx_packets"`
	UserWifi1TxBytes              FlexInt `json:"user-wifi1-tx_bytes"`
	UserWifi1TxDropped            FlexInt `json:"user-wifi1-tx_dropped"`
	UserWifi1TxErrors             FlexInt `json:"user-wifi1-tx_errors"`
	UserWifi1TxPackets            FlexInt `json:"user-wifi1-tx_packets"`
	UserWifi1TxRetries            FlexInt `json:"user-wifi1-tx_retries"`
	UserWifi1WifiTxAttempts       FlexInt `json:"user-wifi1-wifi_tx_attempts"`
	UserWifi1WifiTxDropped        FlexInt `json:"user-wifi1-wifi_tx_dropped"`
	Wifi0MacFilterRejections      FlexInt `json:"wifi0-mac_filter_rejections"`
	Wifi0RxBytes                  FlexInt `json:"wifi0-rx_bytes"`
	Wifi0RxCrypts                 FlexInt `json:"wifi0-rx_crypts"`
	Wifi0RxDropped                FlexInt `json:"wifi0-rx_dropped"`
	Wifi0RxErrors                 FlexInt `json:"wifi0-rx_errors"`
	Wifi0RxFrags                  FlexInt `json:"wifi0-rx_frags"`
	Wifi0RxPackets                FlexInt `json:"wifi0-rx_packets"`
	Wifi0TxBytes                  FlexInt `json:"wifi0-tx_bytes"`
	Wifi0TxDropped                FlexInt `json:"wifi0-tx_dropped"`
	Wifi0TxErrors                 FlexInt `json:"wifi0-tx_errors"`
	Wifi0TxPackets                FlexInt `json:"wifi0-tx_packets"`
	Wifi0TxRetries                FlexInt `json:"wifi0-tx_retries"`
	Wifi0WifiTxAttempts           FlexInt `json:"wifi0-wifi_tx_attempts"`
	Wifi0WifiTxDropped            FlexInt `json:"wifi0-wifi_tx_dropped"`
	Wifi1MacFilterRejections      FlexInt `json:"wifi1-mac_filter_rejections"`
	Wifi1RxBytes                  FlexInt `json:"wifi1-rx_bytes"`
	Wifi1RxCrypts                 FlexInt `json:"wifi1-rx_crypts"`
	Wifi1RxDropped                FlexInt `json:"wifi1-rx_dropped"`
	Wifi1RxErrors                 FlexInt `json:"wifi1-rx_errors"`
	Wifi1RxFrags                  FlexInt `json:"wifi1-rx_frags"`
	Wifi1RxPackets                FlexInt `json:"wifi1-rx_packets"`
	Wifi1TxBytes                  FlexInt `json:"wifi1-tx_bytes"`
	Wifi1TxDropped                FlexInt `json:"wifi1-tx_dropped"`
	Wifi1TxErrors                 FlexInt `json:"wifi1-tx_errors"`
	Wifi1TxPackets                FlexInt `json:"wifi1-tx_packets"`
	Wifi1TxRetries                FlexInt `json:"wifi1-tx_retries"`
	Wifi1WifiTxAttempts           FlexInt `json:"wifi1-wifi_tx_attempts"`
	Wifi1WifiTxDropped            FlexInt `json:"wifi1-wifi_tx_dropped"`
	// UDM Names
	GuestRa0MacFilterRejections  FlexInt `json:"guest-ra0-mac_filter_rejections"`
	GuestRa0RxBytes              FlexInt `json:"guest-ra0-rx_bytes"`
	GuestRa0RxCrypts             FlexInt `json:"guest-ra0-rx_crypts"`
	GuestRa0RxDropped            FlexInt `json:"guest-ra0-rx_dropped"`
	GuestRa0RxErrors             FlexInt `json:"guest-ra0-rx_errors"`
	GuestRa0RxFrags              FlexInt `json:"guest-ra0-rx_frags"`
	GuestRa0RxPackets            FlexInt `json:"guest-ra0-rx_packets"`
	GuestRa0TxBytes              FlexInt `json:"guest-ra0-tx_bytes"`
	GuestRa0TxDropped            FlexInt `json:"guest-ra0-tx_dropped"`
	GuestRa0TxErrors             FlexInt `json:"guest-ra0-tx_errors"`
	GuestRa0TxPackets            FlexInt `json:"guest-ra0-tx_packets"`
	GuestRa0TxRetries            FlexInt `json:"guest-ra0-tx_retries"`
	GuestRa0WifiTxAttempts       FlexInt `json:"guest-ra0-wifi_tx_attempts"`
	GuestRa0WifiTxDropped        FlexInt `json:"guest-ra0-wifi_tx_dropped"`
	GuestRai0MacFilterRejections FlexInt `json:"guest-rai0-mac_filter_rejections"`
	GuestRai0RxBytes             FlexInt `json:"guest-rai0-rx_bytes"`
	GuestRai0RxCrypts            FlexInt `json:"guest-rai0-rx_crypts"`
	GuestRai0RxDropped           FlexInt `json:"guest-rai0-rx_dropped"`
	GuestRai0RxErrors            FlexInt `json:"guest-rai0-rx_errors"`
	GuestRai0RxFrags             FlexInt `json:"guest-rai0-rx_frags"`
	GuestRai0RxPackets           FlexInt `json:"guest-rai0-rx_packets"`
	GuestRai0TxBytes             FlexInt `json:"guest-rai0-tx_bytes"`
	GuestRai0TxDropped           FlexInt `json:"guest-rai0-tx_dropped"`
	GuestRai0TxErrors            FlexInt `json:"guest-rai0-tx_errors"`
	GuestRai0TxPackets           FlexInt `json:"guest-rai0-tx_packets"`
	GuestRai0TxRetries           FlexInt `json:"guest-rai0-tx_retries"`
	GuestRai0WifiTxAttempts      FlexInt `json:"guest-rai0-wifi_tx_attempts"`
	GuestRai0WifiTxDropped       FlexInt `json:"guest-rai0-wifi_tx_dropped"`
	Ra0MacFilterRejections       FlexInt `json:"ra0-mac_filter_rejections"`
	Ra0RxBytes                   FlexInt `json:"ra0-rx_bytes"`
	Ra0RxCrypts                  FlexInt `json:"ra0-rx_crypts"`
	Ra0RxDropped                 FlexInt `json:"ra0-rx_dropped"`
	Ra0RxErrors                  FlexInt `json:"ra0-rx_errors"`
	Ra0RxFrags                   FlexInt `json:"ra0-rx_frags"`
	Ra0RxPackets                 FlexInt `json:"ra0-rx_packets"`
	Ra0TxBytes                   FlexInt `json:"ra0-tx_bytes"`
	Ra0TxDropped                 FlexInt `json:"ra0-tx_dropped"`
	Ra0TxErrors                  FlexInt `json:"ra0-tx_errors"`
	Ra0TxPackets                 FlexInt `json:"ra0-tx_packets"`
	Ra0TxRetries                 FlexInt `json:"ra0-tx_retries"`
	Ra0WifiTxAttempts            FlexInt `json:"ra0-wifi_tx_attempts"`
	Ra0WifiTxDropped             FlexInt `json:"ra0-wifi_tx_dropped"`
	Rai0MacFilterRejections      FlexInt `json:"rai0-mac_filter_rejections"`
	Rai0RxBytes                  FlexInt `json:"rai0-rx_bytes"`
	Rai0RxCrypts                 FlexInt `json:"rai0-rx_crypts"`
	Rai0RxDropped                FlexInt `json:"rai0-rx_dropped"`
	Rai0RxErrors                 FlexInt `json:"rai0-rx_errors"`
	Rai0RxFrags                  FlexInt `json:"rai0-rx_frags"`
	Rai0RxPackets                FlexInt `json:"rai0-rx_packets"`
	Rai0TxBytes                  FlexInt `json:"rai0-tx_bytes"`
	Rai0TxDropped                FlexInt `json:"rai0-tx_dropped"`
	Rai0TxErrors                 FlexInt `json:"rai0-tx_errors"`
	Rai0TxPackets                FlexInt `json:"rai0-tx_packets"`
	Rai0TxRetries                FlexInt `json:"rai0-tx_retries"`
	Rai0WifiTxAttempts           FlexInt `json:"rai0-wifi_tx_attempts"`
	Rai0WifiTxDropped            FlexInt `json:"rai0-wifi_tx_dropped"`
	UserRa0MacFilterRejections   FlexInt `json:"user-ra0-mac_filter_rejections"`
	UserRa0RxBytes               FlexInt `json:"user-ra0-rx_bytes"`
	UserRa0RxCrypts              FlexInt `json:"user-ra0-rx_crypts"`
	UserRa0RxDropped             FlexInt `json:"user-ra0-rx_dropped"`
	UserRa0RxErrors              FlexInt `json:"user-ra0-rx_errors"`
	UserRa0RxFrags               FlexInt `json:"user-ra0-rx_frags"`
	UserRa0RxPackets             FlexInt `json:"user-ra0-rx_packets"`
	UserRa0TxBytes               FlexInt `json:"user-ra0-tx_bytes"`
	UserRa0TxDropped             FlexInt `json:"user-ra0-tx_dropped"`
	UserRa0TxErrors              FlexInt `json:"user-ra0-tx_errors"`
	UserRa0TxPackets             FlexInt `json:"user-ra0-tx_packets"`
	UserRa0TxRetries             FlexInt `json:"user-ra0-tx_retries"`
	UserRa0WifiTxAttempts        FlexInt `json:"user-ra0-wifi_tx_attempts"`
	UserRa0WifiTxDropped         FlexInt `json:"user-ra0-wifi_tx_dropped"`
	UserRai0MacFilterRejections  FlexInt `json:"user-rai0-mac_filter_rejections"`
	UserRai0RxBytes              FlexInt `json:"user-rai0-rx_bytes"`
	UserRai0RxCrypts             FlexInt `json:"user-rai0-rx_crypts"`
	UserRai0RxDropped            FlexInt `json:"user-rai0-rx_dropped"`
	UserRai0RxErrors             FlexInt `json:"user-rai0-rx_errors"`
	UserRai0RxFrags              FlexInt `json:"user-rai0-rx_frags"`
	UserRai0RxPackets            FlexInt `json:"user-rai0-rx_packets"`
	UserRai0TxBytes              FlexInt `json:"user-rai0-tx_bytes"`
	UserRai0TxDropped            FlexInt `json:"user-rai0-tx_dropped"`
	UserRai0TxErrors             FlexInt `json:"user-rai0-tx_errors"`
	UserRai0TxPackets            FlexInt `json:"user-rai0-tx_packets"`
	UserRai0TxRetries            FlexInt `json:"user-rai0-tx_retries"`
	UserRai0WifiTxAttempts       FlexInt `json:"user-rai0-wifi_tx_attempts"`
	UserRai0WifiTxDropped        FlexInt `json:"user-rai0-wifi_tx_dropped"`
}

// RadioTable is part of the data for UAPs and UDMs.
type RadioTable []struct {
	AntennaGain        FlexInt  `json:"antenna_gain"`
	BuiltinAntGain     FlexInt  `json:"builtin_ant_gain"`
	BuiltinAntenna     FlexBool `json:"builtin_antenna"`
	Channel            FlexInt  `json:"channel"`
	CurrentAntennaGain FlexInt  `json:"current_antenna_gain"`
	HasDfs             FlexBool `json:"has_dfs"`
	HasFccdfs          FlexBool `json:"has_fccdfs"`
	HasHt160           FlexBool `json:"has_ht160"`
	Ht                 FlexInt  `json:"ht"`
	Is11Ac             FlexBool `json:"is_11ac"`
	MaxTxpower         FlexInt  `json:"max_txpower"`
	MinRssi            FlexInt  `json:"min_rssi,omitempty"`
	MinRssiEnabled     FlexBool `json:"min_rssi_enabled"`
	MinTxpower         FlexInt  `json:"min_txpower"`
	Name               string   `fake:"{lexify:mock-radio}"  json:"name"`
	Nss                FlexInt  `json:"nss"`
	Radio              string   `json:"radio"`
	RadioCaps          FlexInt  `json:"radio_caps"`
	SensLevelEnabled   FlexBool `json:"sens_level_enabled"`
	TxPower            FlexInt  `json:"tx_power"`
	TxPowerMode        string   `json:"tx_power_mode"`
	VwireEnabled       FlexBool `json:"vwire_enabled"`
	WlangroupID        string   `fake:"{uuid}"               json:"wlangroup_id"`
}

// RadioTableStats is part of the data shared between UAP and UDM.
type RadioTableStats []struct {
	AstBeXmit    FlexInt     `json:"ast_be_xmit"`
	AstCst       interface{} `json:"ast_cst"`
	AstTxto      interface{} `json:"ast_txto"`
	Channel      FlexInt     `json:"channel"`
	CuSelfRx     FlexInt     `json:"cu_self_rx"`
	CuSelfTx     FlexInt     `json:"cu_self_tx"`
	CuTotal      FlexInt     `json:"cu_total"`
	Extchannel   FlexInt     `json:"extchannel"`
	Gain         FlexInt     `json:"gain"`
	GuestNumSta  FlexInt     `json:"guest-num_sta"`
	Name         string      `fake:"{lexify:mock-radio}" json:"name"`
	NumSta       FlexInt     `json:"num_sta"`
	Radio        string      `json:"radio"`
	Satisfaction FlexInt     `json:"satisfaction"`
	State        string      `json:"state"`
	TxPackets    FlexInt     `json:"tx_packets"`
	TxPower      FlexInt     `json:"tx_power"`
	TxRetries    FlexInt     `json:"tx_retries"`
	UserNumSta   FlexInt     `json:"user-num_sta"`
}

// VapTable holds much of the UAP wireless data. Shared by UDM.
type VapTable []struct {
	AnomaliesBarChart struct {
		HighDNSLatency    FlexInt `json:"high_dns_latency"`
		HighTCPLatency    FlexInt `json:"high_tcp_latency"`
		HighTCPPacketLoss FlexInt `json:"high_tcp_packet_loss"`
		HighWifiLatency   FlexInt `json:"high_wifi_latency"`
		HighWifiRetries   FlexInt `json:"high_wifi_retries"`
		LowPhyRate        FlexInt `json:"low_phy_rate"`
		PoorStreamEff     FlexInt `json:"poor_stream_eff"`
		SleepyClient      FlexInt `json:"sleepy_client"`
		StaArpTimeout     FlexInt `json:"sta_arp_timeout"`
		StaDNSTimeout     FlexInt `json:"sta_dns_timeout"`
		StaIPTimeout      FlexInt `json:"sta_ip_timeout"`
		WeakSignal        FlexInt `json:"weak_signal"`
	} `json:"anomalies_bar_chart"`
	AnomaliesBarChartNow struct {
		HighDNSLatency    FlexInt `json:"high_dns_latency"`
		HighTCPLatency    FlexInt `json:"high_tcp_latency"`
		HighTCPPacketLoss FlexInt `json:"high_tcp_packet_loss"`
		HighWifiLatency   FlexInt `json:"high_wifi_latency"`
		HighWifiRetries   FlexInt `json:"high_wifi_retries"`
		LowPhyRate        FlexInt `json:"low_phy_rate"`
		PoorStreamEff     FlexInt `json:"poor_stream_eff"`
		SleepyClient      FlexInt `json:"sleepy_client"`
		StaArpTimeout     FlexInt `json:"sta_arp_timeout"`
		StaDNSTimeout     FlexInt `json:"sta_dns_timeout"`
		StaIPTimeout      FlexInt `json:"sta_ip_timeout"`
		WeakSignal        FlexInt `json:"weak_signal"`
	} `json:"anomalies_bar_chart_now"`
	ApMac               string      `fake:"{macaddress}"             json:"ap_mac"`
	AvgClientSignal     FlexInt     `json:"avg_client_signal"`
	Bssid               string      `fake:"{macaddress}"             json:"bssid"`
	Ccq                 int         `json:"ccq"`
	Channel             FlexInt     `json:"channel"`
	DNSAvgLatency       FlexInt     `json:"dns_avg_latency"`
	Essid               string      `fake:"{macaddress}"             json:"essid"`
	Extchannel          int         `json:"extchannel"`
	ID                  string      `fake:"{uuid}"                   json:"id"`
	IsGuest             FlexBool    `json:"is_guest"`
	IsWep               FlexBool    `json:"is_wep"`
	MacFilterRejections int         `json:"mac_filter_rejections"`
	MapID               interface{} `json:"map_id"`
	Name                string      `fake:"{randomstring:[foo,bar]}" json:"name"`
	NumSatisfactionSta  FlexInt     `json:"num_satisfaction_sta"`
	NumSta              int         `json:"num_sta"`
	Radio               string      `json:"radio"`
	RadioName           string      `json:"radio_name"`
	ReasonsBarChart     struct {
		PhyRate       FlexInt `json:"phy_rate"`
		Signal        FlexInt `json:"signal"`
		SleepyClient  FlexInt `json:"sleepy_client"`
		StaArpTimeout FlexInt `json:"sta_arp_timeout"`
		StaDNSLatency FlexInt `json:"sta_dns_latency"`
		StaDNSTimeout FlexInt `json:"sta_dns_timeout"`
		StaIPTimeout  FlexInt `json:"sta_ip_timeout"`
		StreamEff     FlexInt `json:"stream_eff"`
		TCPLatency    FlexInt `json:"tcp_latency"`
		TCPPacketLoss FlexInt `json:"tcp_packet_loss"`
		WifiLatency   FlexInt `json:"wifi_latency"`
		WifiRetries   FlexInt `json:"wifi_retries"`
	} `json:"reasons_bar_chart"`
	ReasonsBarChartNow struct {
		PhyRate       FlexInt `json:"phy_rate"`
		Signal        FlexInt `json:"signal"`
		SleepyClient  FlexInt `json:"sleepy_client"`
		StaArpTimeout FlexInt `json:"sta_arp_timeout"`
		StaDNSLatency FlexInt `json:"sta_dns_latency"`
		StaDNSTimeout FlexInt `json:"sta_dns_timeout"`
		StaIPTimeout  FlexInt `json:"sta_ip_timeout"`
		StreamEff     FlexInt `json:"stream_eff"`
		TCPLatency    FlexInt `json:"tcp_latency"`
		TCPPacketLoss FlexInt `json:"tcp_packet_loss"`
		WifiLatency   FlexInt `json:"wifi_latency"`
		WifiRetries   FlexInt `json:"wifi_retries"`
	} `json:"reasons_bar_chart_now"`
	RxBytes    FlexInt `json:"rx_bytes"`
	RxCrypts   FlexInt `json:"rx_crypts"`
	RxDropped  FlexInt `json:"rx_dropped"`
	RxErrors   FlexInt `json:"rx_errors"`
	RxFrags    FlexInt `json:"rx_frags"`
	RxNwids    FlexInt `json:"rx_nwids"`
	RxPackets  FlexInt `json:"rx_packets"`
	RxTCPStats struct {
		Goodbytes FlexInt `json:"goodbytes"`
		LatAvg    FlexInt `json:"lat_avg"`
		LatMax    FlexInt `json:"lat_max"`
		LatMin    FlexInt `json:"lat_min"`
		Stalls    FlexInt `json:"stalls"`
	} `json:"rx_tcp_stats"`
	Satisfaction      FlexInt `json:"satisfaction"`
	SatisfactionNow   FlexInt `json:"satisfaction_now"`
	SiteID            string  `fake:"{uuid}"              json:"site_id"`
	State             string  `json:"state"`
	T                 string  `json:"t"`
	TxBytes           FlexInt `json:"tx_bytes"`
	TxCombinedRetries FlexInt `json:"tx_combined_retries"`
	TxDataMpduBytes   FlexInt `json:"tx_data_mpdu_bytes"`
	TxDropped         FlexInt `json:"tx_dropped"`
	TxErrors          FlexInt `json:"tx_errors"`
	TxPackets         FlexInt `json:"tx_packets"`
	TxPower           FlexInt `json:"tx_power"`
	TxRetries         FlexInt `json:"tx_retries"`
	TxRtsRetries      FlexInt `json:"tx_rts_retries"`
	TxSuccess         FlexInt `json:"tx_success"`
	TxTCPStats        struct {
		Goodbytes FlexInt `json:"goodbytes"`
		LatAvg    FlexInt `json:"lat_avg"`
		LatMax    FlexInt `json:"lat_max"`
		LatMin    FlexInt `json:"lat_min"`
		Stalls    FlexInt `json:"stalls"`
	} `json:"tx_tcp_stats"`
	TxTotal          FlexInt  `json:"tx_total"`
	Up               FlexBool `json:"up"`
	Usage            string   `json:"usage"`
	WifiTxAttempts   FlexInt  `json:"wifi_tx_attempts"`
	WifiTxDropped    FlexInt  `json:"wifi_tx_dropped"`
	WifiTxLatencyMov struct {
		Avg        FlexInt `json:"avg"`
		Max        FlexInt `json:"max"`
		Min        FlexInt `json:"min"`
		Total      FlexInt `json:"total"`
		TotalCount FlexInt `json:"total_count"`
	} `json:"wifi_tx_latency_mov"`
	WlanconfID string `json:"wlanconf_id"`
}

// RogueAP are your neighbors access points.
type RogueAP struct {
	Age        FlexInt  `json:"age"`
	ApMac      string   `fake:"{macaddress}" json:"ap_mac"`
	Band       string   `json:"band"`
	Bssid      string   `fake:"{macaddress}" json:"bssid"`
	Bw         FlexInt  `json:"bw"`
	CenterFreq FlexInt  `json:"center_freq"`
	Channel    int      `json:"channel"`
	Essid      string   `fake:"{macaddress}" json:"essid"`
	Freq       FlexInt  `json:"freq"`
	ID         string   `fake:"{uuid}"       json:"_id"`
	IsAdhoc    FlexBool `json:"is_adhoc"`
	IsRogue    FlexBool `json:"is_rogue"`
	IsUbnt     FlexBool `json:"is_ubnt"`
	LastSeen   FlexInt  `json:"last_seen"`
	Noise      FlexInt  `json:"noise"`
	Oui        string   `json:"oui"`
	Radio      string   `json:"radio"`
	RadioName  string   `json:"radio_name"`
	ReportTime FlexInt  `json:"report_time"`
	Rssi       FlexInt  `json:"rssi"`
	RssiAge    FlexInt  `json:"rssi_age"`
	Security   string   `json:"security"`
	Signal     FlexInt  `json:"signal"`
	SiteID     string   `fake:"{uuid}"       json:"site_id"`
	SiteName   string   `json:"-"`
	SourceName string   `json:"-"`
}

// GetRogueAPs returns RogueAPs for a list of Sites.
// Use GetRogueAPsSite if you want more control.
func (u *Unifi) GetRogueAPs(sites []*Site) ([]*RogueAP, error) {
	data := []*RogueAP{}

	for _, site := range sites {
		response, err := u.GetRogueAPsSite(site)
		if err != nil {
			return data, err
		}

		data = append(data, response...)
	}

	return data, nil
}

// GetRogueAPsSite returns RogueAPs for a single Site.
func (u *Unifi) GetRogueAPsSite(site *Site) ([]*RogueAP, error) {
	if site == nil || site.Name == "" {
		return nil, ErrNoSiteProvided
	}

	u.DebugLog("Polling Controller for RogueAPs, site %s (%s)", site.SiteName, site.Desc)

	var (
		path     = fmt.Sprintf(APIRogueAP, site.Name)
		rogueaps struct {
			Data []*RogueAP `json:"data"`
		}
	)

	if err := u.GetData(path, &rogueaps, ""); err != nil {
		return rogueaps.Data, err
	}

	for i := range rogueaps.Data {
		// Add special SourceName value.
		rogueaps.Data[i].SourceName = u.URL
		// Add the special "Site Name" to each event. This becomes a Grafana filter somewhere.
		rogueaps.Data[i].SiteName = site.SiteName
	}

	return rogueaps.Data, nil
}

// UnmarshalJSON unmarshalls 5.10 or 5.11 formatted Access Point Stat data.
func (v *UAPStat) UnmarshalJSON(data []byte) error {
	var n struct {
		Ap `json:"ap"`
	}

	v.Ap = &n.Ap

	err := json.Unmarshal(data, v.Ap) // controller version 5.10.
	if err != nil {
		return json.Unmarshal(data, &n) // controller version 5.11.
	}

	return nil
}
