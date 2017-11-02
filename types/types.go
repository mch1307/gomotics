package types

import "encoding/json"

//ItemType stores the external to internal item types
type ItemType struct {
	Provider     string
	ProviderType string
	InternalType string
}

// Message generic struct to hold nhc messages
// used to identify the message type before further parsing
type Message struct {
	Cmd   string `json:"cmd"`
	Event string `json:"event"`
	Data  json.RawMessage
}

// Action holds one individual nhc action (equipment)
type Action struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     int    `json:"type"`
	Location int    `json:"location"`
	Value1   int    `json:"value1"`
	Value2   int    `json:"value2"`
	Value3   int    `json:"value3"`
}

// Event holds an individual event
type Event struct {
	ID    int `json:"id"`
	Value int `json:"value1"`
}

// Location holds one nhc location
type Location struct {
	ID   int
	Name string
}

// NHCItem represents a registered item
//
// swagger:model
type NHCItem struct {
	// the provider
	// required: true
	Provider string `json:"provider"`
	// the id of this item
	ID int `json:"id"`
	// the type of this item
	// can be switch, dimmer or blind
	Type string `json:"type"`
	// the name of the item
	Name string `json:"name"`
	// the location of the item
	Location string `json:"location"`
	// the current state of the item
	State int `json:"state"`
	// other value of the item
	Value2 int `json:"value2"`
	// other value of the item
	Value3         int `json:"value3"`
	JeedomID       string
	JeedomUpdState string
	JeedomSubType  string
}

// NHCSystemInfo hold the NHC system information
// swagger:model
type NHCSystemInfo struct {
	// NHC Software version
	Swversion       string `json:"swversion"`
	API             string `json:"api"`
	Time            string `json:"time"`
	Language        string `json:"language"`
	Currency        string `json:"currency"`
	Units           int    `json:"units"`
	DST             int    `json:"DST"`
	TZ              int    `json:"TZ"`
	Lastenergyerase string `json:"lastenergyerase"`
	Lastconfig      string `json:"lastconfig"`
}

// JeedomItem holds one Jeedom item (equipment + location)
type JeedomItem struct {
	ID int
}

type JsonRpcArgs struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
	Params  struct {
		Apikey    string `json:"apikey"`
		Type      string `json:"type"`
		EqLogicID string `json:"eqLogic_id"`
	} `json:"params"`
}

type JeedomObjects struct {
	Jsonrpc string           `json:"jsonrpc"`
	ID      string           `json:"id"`
	Result  []JeedomLocation `json:"result"`
}

type JeedomLocation struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	FatherID      interface{} `json:"father_id"`
	IsVisible     string      `json:"isVisible"`
	Position      string      `json:"position"`
	Configuration struct {
		ParentNumber            int    `json:"parentNumber"`
		TagColor                string `json:"tagColor"`
		TagTextColor            string `json:"tagTextColor"`
		DesktopSummaryTextColor string `json:"desktop::summaryTextColor"`
		MobileSummaryTextColor  string `json:"mobile::summaryTextColor"`
	} `json:"configuration"`
	Display struct {
		Icon         string `json:"icon"`
		TagColor     string `json:"tagColor"`
		TagTextColor string `json:"tagTextColor"`
	} `json:"display"`
}

type JeedomEquipments struct {
	Jsonrpc string            `json:"jsonrpc"`
	ID      string            `json:"id"`
	Result  []JeedomEquipment `json:"result"`
}

type JeedomEquipment struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	LogicalID     string      `json:"logicalId"`
	ObjectID      string      `json:"object_id"`
	EqTypeName    string      `json:"eqType_name"`
	EqRealID      interface{} `json:"eqReal_id"`
	IsVisible     string      `json:"isVisible"`
	IsEnable      string      `json:"isEnable"`
	Configuration struct {
		Createtime  string `json:"createtime"`
		Autorefresh string `json:"autorefresh"`
		Updatetime  string `json:"updatetime"`
	} `json:"configuration"`
	Timeout  interface{} `json:"timeout"`
	Category struct {
		Heating    string `json:"heating"`
		Security   string `json:"security"`
		Energy     string `json:"energy"`
		Light      string `json:"light"`
		Automatism string `json:"automatism"`
		Multimedia string `json:"multimedia"`
		Default    string `json:"default"`
	} `json:"category"`
	Display struct {
		ShowOncategory                 int    `json:"showOncategory"`
		ShowObjectNameOncategory       int    `json:"showObjectNameOncategory"`
		ShowNameOncategory             int    `json:"showNameOncategory"`
		ShowOnstyle                    int    `json:"showOnstyle"`
		ShowObjectNameOnstyle          int    `json:"showObjectNameOnstyle"`
		ShowNameOnstyle                int    `json:"showNameOnstyle"`
		ShowObjectNameOnview           int    `json:"showObjectNameOnview"`
		ShowObjectNameOndview          int    `json:"showObjectNameOndview"`
		ShowObjectNameOnmview          int    `json:"showObjectNameOnmview"`
		Height                         string `json:"height"`
		Width                          string `json:"width"`
		LayoutDashboardTableParameters struct {
			Center  int    `json:"center"`
			Styletd string `json:"styletd"`
		} `json:"layout::dashboard::table::parameters"`
		LayoutDashboardTableCmd416Line   int `json:"layout::dashboard::table::cmd::416::line"`
		LayoutDashboardTableCmd416Column int `json:"layout::dashboard::table::cmd::416::column"`
		LayoutDashboardTableCmd415Line   int `json:"layout::dashboard::table::cmd::415::line"`
		LayoutDashboardTableCmd415Column int `json:"layout::dashboard::table::cmd::415::column"`
		LayoutDashboardTableCmd417Line   int `json:"layout::dashboard::table::cmd::417::line"`
		LayoutDashboardTableCmd417Column int `json:"layout::dashboard::table::cmd::417::column"`
		LayoutMobileTableParameters      struct {
			Center  int    `json:"center"`
			Styletd string `json:"styletd"`
		} `json:"layout::mobile::table::parameters"`
		LayoutMobileTableCmd416Line   int `json:"layout::mobile::table::cmd::416::line"`
		LayoutMobileTableCmd416Column int `json:"layout::mobile::table::cmd::416::column"`
		LayoutMobileTableCmd415Line   int `json:"layout::mobile::table::cmd::415::line"`
		LayoutMobileTableCmd415Column int `json:"layout::mobile::table::cmd::415::column"`
		LayoutMobileTableCmd417Line   int `json:"layout::mobile::table::cmd::417::line"`
		LayoutMobileTableCmd417Column int `json:"layout::mobile::table::cmd::417::column"`
	} `json:"display"`
	Order   string      `json:"order"`
	Comment interface{} `json:"comment"`
	Status  string      `json:"status"`
}

/* type JeedomEquipment struct {
	ID            string `json:"id"`
	LogicalID     string `json:"logicalId"`
	EqType        string `json:"eqType"`
	Name          string `json:"name"`
	Order         string `json:"order"`
	Type          string `json:"type"`
	SubType       string `json:"subType"`
	EqLogicID     string `json:"eqLogic_id"`
	IsHistorized  string `json:"isHistorized"`
	Unite         string `json:"unite"`
	Configuration struct {
		VirtualAction    string        `json:"virtualAction"`
		InfoName         string        `json:"infoName"`
		Value            string        `json:"value"`
		MinValue         string        `json:"minValue"`
		MaxValue         string        `json:"maxValue"`
		InfoID           string        `json:"infoId"`
		ActionConfirm    string        `json:"actionConfirm"`
		ActionCodeAccess string        `json:"actionCodeAccess"`
		ActionCheckCmd   []interface{} `json:"actionCheckCmd"`
	} `json:"configuration"`
	Template struct {
		Dashboard string `json:"dashboard"`
		Mobile    string `json:"mobile"`
	} `json:"template"`
	Display struct {
		Icon                  string        `json:"icon"`
		GenericType           string        `json:"generic_type"`
		ShowOndashboard       string        `json:"showOndashboard"`
		ShowOnplan            string        `json:"showOnplan"`
		ShowOnview            string        `json:"showOnview"`
		ShowOnmobile          string        `json:"showOnmobile"`
		ShowNameOndashboard   string        `json:"showNameOndashboard"`
		ShowNameOnplan        string        `json:"showNameOnplan"`
		ShowNameOnview        string        `json:"showNameOnview"`
		ShowNameOnmobile      string        `json:"showNameOnmobile"`
		ForceReturnLineBefore string        `json:"forceReturnLineBefore"`
		ForceReturnLineAfter  string        `json:"forceReturnLineAfter"`
		Parameters            []interface{} `json:"parameters"`
	} `json:"display"`
	HTML struct {
		Enable    string `json:"enable"`
		Dashboard string `json:"dashboard"`
		Dview     string `json:"dview"`
		Dplan     string `json:"dplan"`
		Mobile    string `json:"mobile"`
		Mview     string `json:"mview"`
	} `json:"html"`
	Value     string      `json:"value"`
	IsVisible string      `json:"isVisible"`
	Alert     interface{} `json:"alert"`
} */

type JeedomCMDs struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  []JeedomCMD `json:"result"`
}

type JeedomCMD struct {
	ID            string `json:"id"`
	LogicalID     string `json:"logicalId"`
	EqType        string `json:"eqType"`
	Name          string `json:"name"`
	Order         string `json:"order"`
	Type          string `json:"type"`
	SubType       string `json:"subType"`
	EqLogicID     string `json:"eqLogic_id"`
	IsHistorized  string `json:"isHistorized"`
	Unite         string `json:"unite"`
	Configuration struct {
		RequestType                      string `json:"requestType"`
		Request                          string `json:"request"`
		NoSslCheck                       string `json:"noSslCheck"`
		AllowEmptyResponse               string `json:"allowEmptyResponse"`
		DoNotReportHTTPError             string `json:"doNotReportHttpError"`
		ReponseMustContain               string `json:"reponseMustContain"`
		Timeout                          string `json:"timeout"`
		MaxHTTPRetry                     string `json:"maxHttpRetry"`
		HTTPUsername                     string `json:"http_username"`
		HTTPPassword                     string `json:"http_password"`
		URLXML                           string `json:"urlXml"`
		XMLNoSslCheck                    string `json:"xmlNoSslCheck"`
		XMLTimeout                       string `json:"xmlTimeout"`
		MaxXMLRetry                      string `json:"maxXmlRetry"`
		XMLUsername                      string `json:"xml_username"`
		XMLPassword                      string `json:"xml_password"`
		URLHTML                          string `json:"urlHtml"`
		HTMLNoSslCheck                   string `json:"htmlNoSslCheck"`
		HTMLTimeout                      string `json:"htmlTimeout"`
		MaxHTMLRetry                     string `json:"maxHtmlRetry"`
		HTMLUsername                     string `json:"html_username"`
		HTMLPassword                     string `json:"html_password"`
		URLJSON                          string `json:"urlJson"`
		JSONNoSslCheck                   string `json:"jsonNoSslCheck"`
		JSONTimeout                      string `json:"jsonTimeout"`
		MaxJSONRetry                     string `json:"maxJsonRetry"`
		JSONUsername                     string `json:"json_username"`
		JSONPassword                     string `json:"json_password"`
		MinValue                         string `json:"minValue"`
		MaxValue                         string `json:"maxValue"`
		UpdateCmdID                      string `json:"updateCmdId"`
		UpdateCmdToValue                 string `json:"updateCmdToValue"`
		CalculValueOffset                string `json:"calculValueOffset"`
		HistorizeRound                   string `json:"historizeRound"`
		JeedomCheckCmdOperator           string `json:"jeedomCheckCmdOperator"`
		JeedomCheckCmdTest               string `json:"jeedomCheckCmdTest"`
		JeedomCheckCmdTime               string `json:"jeedomCheckCmdTime"`
		JeedomCheckCmdActionType         string `json:"jeedomCheckCmdActionType"`
		JeedomCheckCmdCmdActionID        string `json:"jeedomCheckCmdCmdActionId"`
		JeedomCheckCmdScenarioActionMode string `json:"jeedomCheckCmdScenarioActionMode"`
		JeedomCheckCmdScenarioActionID   string `json:"jeedomCheckCmdScenarioActionId"`
		HistorizeMode                    string `json:"historizeMode"`
		HistoryPurge                     string `json:"historyPurge"`
		DoNotRepeatEvent                 string `json:"doNotRepeatEvent"`
		JeedomPushURL                    string `json:"jeedomPushUrl"`
		Value                            string `json:"value"`
	} `json:"configuration"`
	Template struct {
		Dashboard string `json:"dashboard"`
		Mobile    string `json:"mobile"`
	} `json:"template"`
	Display struct {
		Icon                     string        `json:"icon"`
		InvertBinary             string        `json:"invertBinary"`
		HideOndashboard          string        `json:"hideOndashboard"`
		HideOnmobile             string        `json:"hideOnmobile"`
		DoNotShowNameOnDashboard string        `json:"doNotShowNameOnDashboard"`
		DoNotShowNameOnView      string        `json:"doNotShowNameOnView"`
		DoNotShowStatOnDashboard string        `json:"doNotShowStatOnDashboard"`
		DoNotShowStatOnView      string        `json:"doNotShowStatOnView"`
		DoNotShowStatOnMobile    string        `json:"doNotShowStatOnMobile"`
		ForceReturnLineBefore    string        `json:"forceReturnLineBefore"`
		ForceReturnLineAfter     string        `json:"forceReturnLineAfter"`
		Parameters               []interface{} `json:"parameters"`
		ShowOncategory           int           `json:"showOncategory"`
		ShowStatsOncategory      int           `json:"showStatsOncategory"`
		ShowNameOncategory       int           `json:"showNameOncategory"`
		ShowOnstyle              int           `json:"showOnstyle"`
		ShowStatsOnstyle         int           `json:"showStatsOnstyle"`
		ShowNameOnstyle          int           `json:"showNameOnstyle"`
		GenericType              string        `json:"generic_type"`
	} `json:"display"`
	HTML         interface{} `json:"html"`
	Value        string      `json:"value"`
	IsVisible    string      `json:"isVisible"`
	Alert        interface{} `json:"alert"`
	GenericType  string      `json:"generic_type"`
	CurrentValue int         `json:"currentValue"`
}
