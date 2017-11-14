package types

import "encoding/json"

// jeetypes contains all struct for interacting with Jeedom JRPC API

// JeedomItem holds one Jeedom item (equipment + location)
type JeedomItem struct {
	ID int
}

type JsonRpcCMDArgs struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
	Params  interface{}
}

type BasicRPCParams struct {
	Apikey     string `json:"apikey"`
	Type       string `json:"type"`
	EqLogicID  string `json:"eqLogic_id"`
	Name       string `json:"name"`
	Display    string `json:"display"`
	ObjectID   string `json:"object_id"`
	IsVisible  string `json:"isVisible"`
	IsEnable   string `json:"isEnable"`
	EqTypeName string `json:"eqType_name"`
}

type JsonRpcArgs struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      string      `json:"id"`
	Params  interface{} `json:"params"`
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

type JeedomCMDs struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      string      `json:"id"`
	Result  []JeedomCMD `json:"result"`
}

type JeedomCMDConfig struct {
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
}

type JeedomCMDDisplay struct {
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
	ShowNameOnDashboard      string        `json:"showNameOndashboard"`
	GenericType              string        `json:"generic_type"`
}

type JeedomCMDTemplate struct {
	Dashboard string `json:"dashboard"`
	Mobile    string `json:"mobile"`
}

type JeedomCMD struct {
	Apikey        string            `json:"apikey"`
	ID            string            `json:"id"`
	LogicalID     string            `json:"logicalId"`
	EqType        string            `json:"eqType_name"`
	Name          string            `json:"name"`
	Order         string            `json:"order"`
	Type          string            `json:"type"`
	SubType       string            `json:"subType"`
	EqLogicID     string            `json:"eqLogic_id"`
	IsHistorized  string            `json:"isHistorized"`
	Unite         string            `json:"unite"`
	Configuration JeedomCMDConfig   `json:"configuration"`
	Template      JeedomCMDTemplate `json:"template"`
	Display       JeedomCMDDisplay  `json:"display"`
	HTML          interface{}       `json:"html"`
	IsVisible     string            `json:"isVisible"`
	GenericType   string            `json:"generic_type"`
	CurrentValue  int               `json:"currentValue"`
	Value         string            `json:"value"`
}

type JeedomVersion struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result"`
}

type JeedomPlugins struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Licence      string `json:"licence"`
		Installation string `json:"installation"`
		Author       string `json:"author"`
		Require      string `json:"require"`
		Category     string `json:"category"`
		Data         json.RawMessage
	} `json:"result"`
}

type JeedomCreateRsp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID   string `json:"id"`
		Data json.RawMessage
	} `json:"result"`
}

type JeedomCreateEqRsp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID            string      `json:"id"`
		Name          string      `json:"name"`
		LogicalID     string      `json:"logicalId"`
		ObjectID      string      `json:"object_id"`
		EqTypeName    string      `json:"eqType_name"`
		EqRealID      interface{} `json:"eqReal_id"`
		IsVisible     string      `json:"isVisible"`
		IsEnable      string      `json:"isEnable"`
		Configuration struct {
			Createtime string `json:"createtime"`
			Updatetime string `json:"updatetime"`
		} `json:"configuration"`
		Timeout  interface{} `json:"timeout"`
		Category interface{} `json:"category"`
		Display  struct {
			ShowObjectNameOnview           int    `json:"showObjectNameOnview"`
			ShowObjectNameOndview          int    `json:"showObjectNameOndview"`
			ShowObjectNameOnmview          int    `json:"showObjectNameOnmview"`
			Height                         string `json:"height"`
			Width                          string `json:"width"`
			LayoutDashboardTableParameters struct {
				Center  int    `json:"center"`
				Styletd string `json:"styletd"`
			} `json:"layout::dashboard::table::parameters"`
			LayoutMobileTableParameters struct {
				Center  int    `json:"center"`
				Styletd string `json:"styletd"`
			} `json:"layout::mobile::table::parameters"`
			LayoutDashboardTableCmd21Line   int `json:"layout::dashboard::table::cmd::21::line"`
			LayoutDashboardTableCmd21Column int `json:"layout::dashboard::table::cmd::21::column"`
			LayoutMobileTableCmd21Line      int `json:"layout::mobile::table::cmd::21::line"`
			LayoutMobileTableCmd21Column    int `json:"layout::mobile::table::cmd::21::column"`
		} `json:"display"`
		Order   int         `json:"order"`
		Comment interface{} `json:"comment"`
		Status  string      `json:"status"`
	} `json:"result"`
}

type JeedomCreateCmdRsp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  struct {
		ID            string      `json:"id"`
		LogicalID     interface{} `json:"logicalId"`
		EqType        string      `json:"eqType"`
		Name          string      `json:"name"`
		Order         string      `json:"order"`
		Type          string      `json:"type"`
		SubType       string      `json:"subType"`
		EqLogicID     string      `json:"eqLogic_id"`
		IsHistorized  string      `json:"isHistorized"`
		Unite         string      `json:"unite"`
		Configuration struct {
			RequestType            string `json:"requestType"`
			Request                string `json:"request"`
			NoSslCheck             string `json:"noSslCheck"`
			AllowEmptyResponse     string `json:"allowEmptyResponse"`
			DoNotReportHTTPError   string `json:"doNotReportHttpError"`
			XMLNoSslCheck          string `json:"xmlNoSslCheck"`
			HTMLNoSslCheck         string `json:"htmlNoSslCheck"`
			JSONNoSslCheck         string `json:"jsonNoSslCheck"`
			JeedomCheckCmdOperator string `json:"jeedomCheckCmdOperator"`
		} `json:"configuration"`
		Template struct {
			Dashboard string `json:"dashboard"`
			Mobile    string `json:"mobile"`
		} `json:"template"`
		Display struct {
			InvertBinary          string `json:"invertBinary"`
			ForceReturnLineBefore string `json:"forceReturnLineBefore"`
			ForceReturnLineAfter  string `json:"forceReturnLineAfter"`
			Parameters            string `json:"parameters"`
			ShowOncategory        string `json:"showOncategory"`
			ShowStatsOncategory   string `json:"showStatsOncategory"`
			ShowNameOncategory    string `json:"showNameOncategory"`
			ShowOnstyle           string `json:"showOnstyle"`
			ShowStatsOnstyle      string `json:"showStatsOnstyle"`
			ShowNameOnstyle       string `json:"showNameOnstyle"`
		} `json:"display"`
		HTML struct {
			Enable string `json:"enable"`
		} `json:"html"`
		Value     interface{} `json:"value"`
		IsVisible string      `json:"isVisible"`
		Alert     interface{} `json:"alert"`
	} `json:"result"`
}

type JeedomDimmerStateCMD struct {
	Apikey        string `json:"apikey"`
	ID            string `json:"id"`
	LogicalID     string `json:"logicalId"`
	EqType        string `json:"eqType_name"`
	Name          string `json:"name"`
	Order         string `json:"order"`
	Type          string `json:"type"`
	SubType       string `json:"subType"`
	EqLogicID     string `json:"eqLogic_id"`
	IsHistorized  string `json:"isHistorized"`
	Unite         string `json:"unite"`
	Configuration struct {
		RequestType                      string        `json:"requestType"`
		Request                          string        `json:"request"`
		NoSslCheck                       string        `json:"noSslCheck"`
		AllowEmptyResponse               string        `json:"allowEmptyResponse"`
		DoNotReportHTTPError             string        `json:"doNotReportHttpError"`
		ReponseMustContain               string        `json:"reponseMustContain"`
		Timeout                          string        `json:"timeout"`
		MaxHTTPRetry                     string        `json:"maxHttpRetry"`
		HTTPUsername                     string        `json:"http_username"`
		HTTPPassword                     string        `json:"http_password"`
		URLXML                           string        `json:"urlXml"`
		XMLNoSslCheck                    string        `json:"xmlNoSslCheck"`
		XMLTimeout                       string        `json:"xmlTimeout"`
		MaxXMLRetry                      string        `json:"maxXmlRetry"`
		XMLUsername                      string        `json:"xml_username"`
		XMLPassword                      string        `json:"xml_password"`
		URLHTML                          string        `json:"urlHtml"`
		HTMLNoSslCheck                   string        `json:"htmlNoSslCheck"`
		HTMLTimeout                      string        `json:"htmlTimeout"`
		MaxHTMLRetry                     string        `json:"maxHtmlRetry"`
		HTMLUsername                     string        `json:"html_username"`
		HTMLPassword                     string        `json:"html_password"`
		URLJSON                          string        `json:"urlJson"`
		JSONNoSslCheck                   string        `json:"jsonNoSslCheck"`
		JSONTimeout                      string        `json:"jsonTimeout"`
		MaxJSONRetry                     string        `json:"maxJsonRetry"`
		JSONUsername                     string        `json:"json_username"`
		JSONPassword                     string        `json:"json_password"`
		MinValue                         string        `json:"minValue"`
		MaxValue                         string        `json:"maxValue"`
		UpdateCmdID                      string        `json:"updateCmdId"`
		UpdateCmdToValue                 string        `json:"updateCmdToValue"`
		CalculValueOffset                string        `json:"calculValueOffset"`
		HistorizeRound                   string        `json:"historizeRound"`
		JeedomCheckCmdOperator           string        `json:"jeedomCheckCmdOperator"`
		JeedomCheckCmdTest               string        `json:"jeedomCheckCmdTest"`
		JeedomCheckCmdTime               string        `json:"jeedomCheckCmdTime"`
		JeedomCheckCmdActionType         string        `json:"jeedomCheckCmdActionType"`
		JeedomCheckCmdCmdActionID        string        `json:"jeedomCheckCmdCmdActionId"`
		JeedomCheckCmdScenarioActionMode string        `json:"jeedomCheckCmdScenarioActionMode"`
		JeedomCheckCmdScenarioActionID   string        `json:"jeedomCheckCmdScenarioActionId"`
		HistorizeMode                    string        `json:"historizeMode"`
		HistoryPurge                     string        `json:"historyPurge"`
		DoNotRepeatEvent                 string        `json:"doNotRepeatEvent"`
		JeedomPushURL                    string        `json:"jeedomPushUrl"`
		Value                            string        `json:"value"`
		TimelineEnable                   string        `json:"timeline::enable"`
		DenyValues                       string        `json:"denyValues"`
		ReturnStateValue                 string        `json:"returnStateValue"`
		ReturnStateTime                  string        `json:"returnStateTime"`
		RepeatEventManagement            string        `json:"repeatEventManagement"`
		ActionCheckCmd                   []interface{} `json:"actionCheckCmd"`
		JeedomPreExecCmd                 []interface{} `json:"jeedomPreExecCmd"`
		JeedomPostExecCmd                []interface{} `json:"jeedomPostExecCmd"`
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
		ShowOncategory           string        `json:"showOncategory"`
		ShowStatsOncategory      string        `json:"showStatsOncategory"`
		ShowNameOncategory       string        `json:"showNameOncategory"`
		ShowOnstyle              string        `json:"showOnstyle"`
		ShowStatsOnstyle         string        `json:"showStatsOnstyle"`
		ShowNameOnstyle          string        `json:"showNameOnstyle"`
		ShowNameOndashboard      string        `json:"showNameOndashboard"`
		GenericType              string        `json:"generic_type"`
		ShowOndashboard          string        `json:"showOndashboard"`
		ShowOnplan               string        `json:"showOnplan"`
		ShowOnview               string        `json:"showOnview"`
		ShowOnmobile             string        `json:"showOnmobile"`
		ShowNameOnplan           string        `json:"showNameOnplan"`
		ShowNameOnview           string        `json:"showNameOnview"`
		ShowNameOnmobile         string        `json:"showNameOnmobile"`
		ShowIconAndNamedashboard string        `json:"showIconAndNamedashboard"`
		ShowIconAndNameplan      string        `json:"showIconAndNameplan"`
		ShowIconAndNameview      string        `json:"showIconAndNameview"`
		ShowIconAndNamemobile    string        `json:"showIconAndNamemobile"`
	} `json:"display"`
	HTML struct {
		Enable    string `json:"enable"`
		Dashboard string `json:"dashboard"`
		Dview     string `json:"dview"`
		Dplan     string `json:"dplan"`
		Mobile    string `json:"mobile"`
		Mview     string `json:"mview"`
	} `json:"html"`
	Value     string `json:"value"`
	IsVisible string `json:"isVisible"`
	Alert     struct {
		Warningif     string `json:"warningif"`
		Warningduring string `json:"warningduring"`
		Dangerif      string `json:"dangerif"`
		Dangerduring  string `json:"dangerduring"`
	} `json:"alert"`
	GenericType  string `json:"generic_type"`
	CurrentValue int    `json:"currentValue"`
}

type JeedomDimmerDimCMD struct {
	Apikey        string `json:"apikey"`
	ID            string `json:"id"`
	LogicalID     string `json:"logicalId"`
	EqType        string `json:"eqType_name"`
	Name          string `json:"name"`
	Order         string `json:"order"`
	Type          string `json:"type"`
	SubType       string `json:"subType"`
	EqLogicID     string `json:"eqLogic_id"`
	IsHistorized  string `json:"isHistorized"`
	Unite         string `json:"unite"`
	Configuration struct {
		RequestType          string        `json:"requestType"`
		Request              string        `json:"request"`
		NoSslCheck           string        `json:"noSslCheck"`
		AllowEmptyResponse   string        `json:"allowEmptyResponse"`
		DoNotReportHTTPError string        `json:"doNotReportHttpError"`
		ReponseMustContain   string        `json:"reponseMustContain"`
		Timeout              string        `json:"timeout"`
		MaxHTTPRetry         string        `json:"maxHttpRetry"`
		HTTPUsername         string        `json:"http_username"`
		HTTPPassword         string        `json:"http_password"`
		URLXML               string        `json:"urlXml"`
		XMLNoSslCheck        string        `json:"xmlNoSslCheck"`
		XMLTimeout           string        `json:"xmlTimeout"`
		MaxXMLRetry          string        `json:"maxXmlRetry"`
		XMLUsername          string        `json:"xml_username"`
		XMLPassword          string        `json:"xml_password"`
		URLHTML              string        `json:"urlHtml"`
		HTMLNoSslCheck       string        `json:"htmlNoSslCheck"`
		HTMLTimeout          string        `json:"htmlTimeout"`
		MaxHTMLRetry         string        `json:"maxHtmlRetry"`
		HTMLUsername         string        `json:"html_username"`
		HTMLPassword         string        `json:"html_password"`
		URLJSON              string        `json:"urlJson"`
		JSONNoSslCheck       string        `json:"jsonNoSslCheck"`
		JSONTimeout          string        `json:"jsonTimeout"`
		MaxJSONRetry         string        `json:"maxJsonRetry"`
		JSONUsername         string        `json:"json_username"`
		JSONPassword         string        `json:"json_password"`
		MinValue             string        `json:"minValue"`
		MaxValue             string        `json:"maxValue"`
		UpdateCmdID          string        `json:"updateCmdId"`
		UpdateCmdToValue     string        `json:"updateCmdToValue"`
		TimelineEnable       string        `json:"timeline::enable"`
		ActionConfirm        string        `json:"actionConfirm"`
		ActionCodeAccess     string        `json:"actionCodeAccess"`
		ActionCheckCmd       []interface{} `json:"actionCheckCmd"`
		JeedomPreExecCmd     []interface{} `json:"jeedomPreExecCmd"`
		JeedomPostExecCmd    []interface{} `json:"jeedomPostExecCmd"`
	} `json:"configuration"`
	Template struct {
		Dashboard string `json:"dashboard"`
		Mobile    string `json:"mobile"`
	} `json:"template"`
	Display struct {
		Icon                     string        `json:"icon"`
		InvertBinary             string        `json:"invertBinary"`
		GenericType              string        `json:"generic_type"`
		ShowOndashboard          string        `json:"showOndashboard"`
		ShowOnplan               string        `json:"showOnplan"`
		ShowOnview               string        `json:"showOnview"`
		ShowOnmobile             string        `json:"showOnmobile"`
		ShowNameOndashboard      string        `json:"showNameOndashboard"`
		ShowNameOnplan           string        `json:"showNameOnplan"`
		ShowNameOnview           string        `json:"showNameOnview"`
		ShowNameOnmobile         string        `json:"showNameOnmobile"`
		ShowIconAndNamedashboard string        `json:"showIconAndNamedashboard"`
		ShowIconAndNameplan      string        `json:"showIconAndNameplan"`
		ShowIconAndNameview      string        `json:"showIconAndNameview"`
		ShowIconAndNamemobile    string        `json:"showIconAndNamemobile"`
		ForceReturnLineBefore    string        `json:"forceReturnLineBefore"`
		ForceReturnLineAfter     string        `json:"forceReturnLineAfter"`
		Parameters               []interface{} `json:"parameters"`
	} `json:"display"`
	HTML struct {
		Enable    string `json:"enable"`
		Dashboard string `json:"dashboard"`
		Dview     string `json:"dview"`
		Dplan     string `json:"dplan"`
		Mobile    string `json:"mobile"`
		Mview     string `json:"mview"`
	} `json:"html"`
	Value        string      `json:"value"`
	IsVisible    string      `json:"isVisible"`
	Alert        interface{} `json:"alert"`
	GenericType  string      `json:"generic_type"`
	CurrentValue interface{} `json:"currentValue"`
}

// JeedomDimmerUpdStateCMD
type JeedomDimmerUpdStateCMD struct {
	Apikey        string      `json:"apikey"`
	ID            string      `json:"id"`
	LogicalID     interface{} `json:"logicalId"`
	EqType        string      `json:"eqType_name"`
	Name          string      `json:"name"`
	Order         string      `json:"order"`
	Type          string      `json:"type"`
	SubType       string      `json:"subType"`
	EqLogicID     string      `json:"eqLogic_id"`
	IsHistorized  string      `json:"isHistorized"`
	Unite         string      `json:"unite"`
	Configuration struct {
		RequestType          string `json:"requestType"`
		Request              string `json:"request"`
		NoSslCheck           string `json:"noSslCheck"`
		AllowEmptyResponse   string `json:"allowEmptyResponse"`
		DoNotReportHTTPError string `json:"doNotReportHttpError"`
		ReponseMustContain   string `json:"reponseMustContain"`
		Timeout              string `json:"timeout"`
		MaxHTTPRetry         string `json:"maxHttpRetry"`
		HTTPUsername         string `json:"http_username"`
		HTTPPassword         string `json:"http_password"`
		URLXML               string `json:"urlXml"`
		XMLNoSslCheck        string `json:"xmlNoSslCheck"`
		XMLTimeout           string `json:"xmlTimeout"`
		MaxXMLRetry          string `json:"maxXmlRetry"`
		XMLUsername          string `json:"xml_username"`
		XMLPassword          string `json:"xml_password"`
		URLHTML              string `json:"urlHtml"`
		HTMLNoSslCheck       string `json:"htmlNoSslCheck"`
		HTMLTimeout          string `json:"htmlTimeout"`
		MaxHTMLRetry         string `json:"maxHtmlRetry"`
		HTMLUsername         string `json:"html_username"`
		HTMLPassword         string `json:"html_password"`
		URLJSON              string `json:"urlJson"`
		JSONNoSslCheck       string `json:"jsonNoSslCheck"`
		JSONTimeout          string `json:"jsonTimeout"`
		MaxJSONRetry         string `json:"maxJsonRetry"`
		JSONUsername         string `json:"json_username"`
		JSONPassword         string `json:"json_password"`
		MinValue             string `json:"minValue"`
		MaxValue             string `json:"maxValue"`
		UpdateCmdID          string `json:"updateCmdId"`
		UpdateCmdToValue     string `json:"updateCmdToValue"`
	} `json:"configuration"`
	Template interface{} `json:"template"`
	Display  struct {
		Icon         string `json:"icon"`
		InvertBinary string `json:"invertBinary"`
	} `json:"display"`
	HTML         interface{} `json:"html"`
	Value        string      `json:"value"`
	IsVisible    string      `json:"isVisible"`
	Alert        interface{} `json:"alert"`
	GenericType  string      `json:"generic_type"`
	CurrentValue interface{} `json:"currentValue"`
}

type JeedomSwitchStateCMD struct {
	Apikey        string `json:"apikey"`
	ID            string `json:"id"`
	LogicalID     string `json:"logicalId"`
	EqType        string `json:"eqType_name"`
	Name          string `json:"name"`
	Order         string `json:"order"`
	Type          string `json:"type"`
	SubType       string `json:"subType"`
	EqLogicID     string `json:"eqLogic_id"`
	IsHistorized  string `json:"isHistorized"`
	Unite         string `json:"unite"`
	Configuration struct {
		RequestType            string        `json:"requestType"`
		Request                string        `json:"request"`
		NoSslCheck             string        `json:"noSslCheck"`
		AllowEmptyResponse     string        `json:"allowEmptyResponse"`
		DoNotReportHTTPError   string        `json:"doNotReportHttpError"`
		ReponseMustContain     string        `json:"reponseMustContain"`
		Timeout                string        `json:"timeout"`
		MaxHTTPRetry           string        `json:"maxHttpRetry"`
		HTTPUsername           string        `json:"http_username"`
		HTTPPassword           string        `json:"http_password"`
		URLXML                 string        `json:"urlXml"`
		XMLNoSslCheck          string        `json:"xmlNoSslCheck"`
		XMLTimeout             string        `json:"xmlTimeout"`
		MaxXMLRetry            string        `json:"maxXmlRetry"`
		XMLUsername            string        `json:"xml_username"`
		XMLPassword            string        `json:"xml_password"`
		URLHTML                string        `json:"urlHtml"`
		HTMLNoSslCheck         string        `json:"htmlNoSslCheck"`
		HTMLTimeout            string        `json:"htmlTimeout"`
		MaxHTMLRetry           string        `json:"maxHtmlRetry"`
		HTMLUsername           string        `json:"html_username"`
		HTMLPassword           string        `json:"html_password"`
		URLJSON                string        `json:"urlJson"`
		JSONNoSslCheck         string        `json:"jsonNoSslCheck"`
		JSONTimeout            string        `json:"jsonTimeout"`
		MaxJSONRetry           string        `json:"maxJsonRetry"`
		JSONUsername           string        `json:"json_username"`
		JSONPassword           string        `json:"json_password"`
		MinValue               string        `json:"minValue"`
		MaxValue               string        `json:"maxValue"`
		UpdateCmdID            string        `json:"updateCmdId"`
		UpdateCmdToValue       string        `json:"updateCmdToValue"`
		TimelineEnable         string        `json:"timeline::enable"`
		CalculValueOffset      string        `json:"calculValueOffset"`
		JeedomCheckCmdOperator string        `json:"jeedomCheckCmdOperator"`
		JeedomCheckCmdTest     string        `json:"jeedomCheckCmdTest"`
		JeedomCheckCmdTime     string        `json:"jeedomCheckCmdTime"`
		HistoryPurge           string        `json:"historyPurge"`
		DenyValues             string        `json:"denyValues"`
		ReturnStateValue       string        `json:"returnStateValue"`
		ReturnStateTime        string        `json:"returnStateTime"`
		RepeatEventManagement  string        `json:"repeatEventManagement"`
		JeedomPushURL          string        `json:"jeedomPushUrl"`
		ActionCheckCmd         []interface{} `json:"actionCheckCmd"`
		JeedomPreExecCmd       []interface{} `json:"jeedomPreExecCmd"`
		JeedomPostExecCmd      []interface{} `json:"jeedomPostExecCmd"`
	} `json:"configuration"`
	Template struct {
		Dashboard string `json:"dashboard"`
		Mobile    string `json:"mobile"`
	} `json:"template"`
	Display struct {
		Icon                     string        `json:"icon"`
		InvertBinary             string        `json:"invertBinary"`
		GenericType              string        `json:"generic_type"`
		ShowOndashboard          string        `json:"showOndashboard"`
		ShowOnplan               string        `json:"showOnplan"`
		ShowOnview               string        `json:"showOnview"`
		ShowOnmobile             string        `json:"showOnmobile"`
		ShowNameOndashboard      string        `json:"showNameOndashboard"`
		ShowNameOnplan           string        `json:"showNameOnplan"`
		ShowNameOnview           string        `json:"showNameOnview"`
		ShowNameOnmobile         string        `json:"showNameOnmobile"`
		ShowIconAndNamedashboard string        `json:"showIconAndNamedashboard"`
		ShowIconAndNameplan      string        `json:"showIconAndNameplan"`
		ShowIconAndNameview      string        `json:"showIconAndNameview"`
		ShowIconAndNamemobile    string        `json:"showIconAndNamemobile"`
		ForceReturnLineBefore    string        `json:"forceReturnLineBefore"`
		ForceReturnLineAfter     string        `json:"forceReturnLineAfter"`
		Parameters               []interface{} `json:"parameters"`
	} `json:"display"`
	HTML struct {
		Enable    string `json:"enable"`
		Dashboard string `json:"dashboard"`
		Dview     string `json:"dview"`
		Dplan     string `json:"dplan"`
		Mobile    string `json:"mobile"`
		Mview     string `json:"mview"`
	} `json:"html"`
	Value     string `json:"value"`
	IsVisible string `json:"isVisible"`
	Alert     struct {
		Warningif     string `json:"warningif"`
		Warningduring string `json:"warningduring"`
		Dangerif      string `json:"dangerif"`
		Dangerduring  string `json:"dangerduring"`
	} `json:"alert"`
	GenericType  string `json:"generic_type"`
	CurrentValue int    `json:"currentValue"`
}

type JeedomSwitchOnOffCMD struct {
	Apikey        string      `json:"apikey"`
	ID            string      `json:"id"`
	LogicalID     interface{} `json:"logicalId"`
	EqType        string      `json:"eqType_name"`
	Name          string      `json:"name"`
	Order         string      `json:"order"`
	Type          string      `json:"type"`
	SubType       string      `json:"subType"`
	EqLogicID     string      `json:"eqLogic_id"`
	IsHistorized  string      `json:"isHistorized"`
	Unite         string      `json:"unite"`
	Configuration struct {
		RequestType          string `json:"requestType"`
		Request              string `json:"request"`
		NoSslCheck           string `json:"noSslCheck"`
		AllowEmptyResponse   string `json:"allowEmptyResponse"`
		DoNotReportHTTPError string `json:"doNotReportHttpError"`
		ReponseMustContain   string `json:"reponseMustContain"`
		Timeout              string `json:"timeout"`
		MaxHTTPRetry         string `json:"maxHttpRetry"`
		HTTPUsername         string `json:"http_username"`
		HTTPPassword         string `json:"http_password"`
		URLXML               string `json:"urlXml"`
		XMLNoSslCheck        string `json:"xmlNoSslCheck"`
		XMLTimeout           string `json:"xmlTimeout"`
		MaxXMLRetry          string `json:"maxXmlRetry"`
		XMLUsername          string `json:"xml_username"`
		XMLPassword          string `json:"xml_password"`
		URLHTML              string `json:"urlHtml"`
		HTMLNoSslCheck       string `json:"htmlNoSslCheck"`
		HTMLTimeout          string `json:"htmlTimeout"`
		MaxHTMLRetry         string `json:"maxHtmlRetry"`
		HTMLUsername         string `json:"html_username"`
		HTMLPassword         string `json:"html_password"`
		URLJSON              string `json:"urlJson"`
		JSONNoSslCheck       string `json:"jsonNoSslCheck"`
		JSONTimeout          string `json:"jsonTimeout"`
		MaxJSONRetry         string `json:"maxJsonRetry"`
		JSONUsername         string `json:"json_username"`
		JSONPassword         string `json:"json_password"`
		MinValue             string `json:"minValue"`
		MaxValue             string `json:"maxValue"`
		UpdateCmdID          string `json:"updateCmdId"`
		UpdateCmdToValue     string `json:"updateCmdToValue"`
	} `json:"configuration"`
	Template interface{} `json:"template"`
	Display  struct {
		Icon         string `json:"icon"`
		InvertBinary string `json:"invertBinary"`
	} `json:"display"`
	HTML         interface{} `json:"html"`
	Value        string      `json:"value"`
	IsVisible    string      `json:"isVisible"`
	Alert        interface{} `json:"alert"`
	GenericType  string      `json:"generic_type"`
	CurrentValue interface{} `json:"currentValue"`
}

type JeedomSwitchUpdStateCMD struct {
	Apikey        string      `json:"apikey"`
	ID            string      `json:"id"`
	LogicalID     interface{} `json:"logicalId"`
	EqType        string      `json:"eqType_name"`
	Name          string      `json:"name"`
	Order         string      `json:"order"`
	Type          string      `json:"type"`
	SubType       string      `json:"subType"`
	EqLogicID     string      `json:"eqLogic_id"`
	IsHistorized  string      `json:"isHistorized"`
	Unite         string      `json:"unite"`
	Configuration struct {
		RequestType          string `json:"requestType"`
		Request              string `json:"request"`
		NoSslCheck           string `json:"noSslCheck"`
		AllowEmptyResponse   string `json:"allowEmptyResponse"`
		DoNotReportHTTPError string `json:"doNotReportHttpError"`
		ReponseMustContain   string `json:"reponseMustContain"`
		Timeout              string `json:"timeout"`
		MaxHTTPRetry         string `json:"maxHttpRetry"`
		HTTPUsername         string `json:"http_username"`
		HTTPPassword         string `json:"http_password"`
		URLXML               string `json:"urlXml"`
		XMLNoSslCheck        string `json:"xmlNoSslCheck"`
		XMLTimeout           string `json:"xmlTimeout"`
		MaxXMLRetry          string `json:"maxXmlRetry"`
		XMLUsername          string `json:"xml_username"`
		XMLPassword          string `json:"xml_password"`
		URLHTML              string `json:"urlHtml"`
		HTMLNoSslCheck       string `json:"htmlNoSslCheck"`
		HTMLTimeout          string `json:"htmlTimeout"`
		MaxHTMLRetry         string `json:"maxHtmlRetry"`
		HTMLUsername         string `json:"html_username"`
		HTMLPassword         string `json:"html_password"`
		URLJSON              string `json:"urlJson"`
		JSONNoSslCheck       string `json:"jsonNoSslCheck"`
		JSONTimeout          string `json:"jsonTimeout"`
		MaxJSONRetry         string `json:"maxJsonRetry"`
		JSONUsername         string `json:"json_username"`
		JSONPassword         string `json:"json_password"`
		MinValue             string `json:"minValue"`
		MaxValue             string `json:"maxValue"`
		UpdateCmdID          string `json:"updateCmdId"`
		UpdateCmdToValue     string `json:"updateCmdToValue"`
	} `json:"configuration"`
	Template interface{} `json:"template"`
	Display  struct {
		Icon         string `json:"icon"`
		InvertBinary string `json:"invertBinary"`
	} `json:"display"`
	HTML         interface{} `json:"html"`
	Value        string      `json:"value"`
	IsVisible    string      `json:"isVisible"`
	Alert        interface{} `json:"alert"`
	GenericType  string      `json:"generic_type"`
	CurrentValue interface{} `json:"currentValue"`
}
