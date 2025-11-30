package figi

import "errors"

var (
	ErrMissingID     = errors.New("missing ID")
	ErrMissingIDType = errors.New("missing ID type")
)

type MappingRequest struct {
	IDType       IDType `json:"idType"`
	IDValue      string `json:"idValue"`
	ExchangeCode string `json:"exchCode,omitempty"`
	MIC          string `json:"micCode,omitempty"`
	Currency     string `json:"currency,omitempty"`
	MarketSector string `json:"marketSecDes,omitempty"`
}

func (m *MappingRequest) Validate() error {
	switch {
	case m.IDType == 0 || !m.IDType.IsAIDType():
		return ErrMissingIDType
	case m.IDValue == "":
		return ErrMissingID
	default:
		return nil
	}
}

//go:generate enumer -type IDType -json
type IDType uint8

const (
	IDTypeUnspecified              IDType = iota
	ID_ISIN                               //	ISIN - International Securities Identification Number.
	ID_BB_UNIQUE                          //	Unique Bloomberg Identifier - A legacy, internal Bloomberg identifier.
	ID_SEDOL                              //	Sedol Number - Stock Exchange Daily Official List.
	ID_COMMON                             //	Common Code - A nine digit identification number.
	ID_WERTPAPIER                         //	Wertpapierkennnummer/WKN - German securities identification code.
	ID_CUSIP                              //	CUSIP - Committee on Uniform Securities Identification Procedures.
	ID_CINS                               //	CINS - CUSIP International Numbering System.
	ID_BB                                 //	A legacy Bloomberg identifier.
	ID_BB_8_CHR                           //	A legacy Bloomberg identifier (8 characters only).
	ID_TRACE                              //	Trace eligible bond identifier issued by FINRA.
	ID_ITALY                              //	Italian Identifier Number - The Italian Identification number consisting of five or six digits.
	ID_EXCH_SYMBOL                        //	Common Code - A nine digit identification number.
	ID_FULL_EXCHANGE_SYMBOL               //	Full Exchange Symbol - Contains the exchange symbol for futures, options, indices inclusive of base symbol and other security elements.
	COMPOSITE_ID_BB_GLOBAL                //	Composite Financial Instrument Global Identifier - The Composite Financial Instrument Global Identifier (FIGI) enables users to link multiple FIGIs at the trading venue level within the same country or market in order to obtain an aggregated view for an instrument within that country or market.
	ID_BB_GLOBAL_SHARE_CLASS_LEVEL        //	Share Class Financial Instrument Global Identifier - A Share Class level Financial Instrument Global Identifier is assigned to an instrument that is traded in more than one country. This enables users to link multiple Composite FIGIs for the same instrument in order to obtain an aggregated view for that instrument across all countries (globally).
	ID_BB_GLOBAL                          //	Financial Instrument Global Identifier (FIGI) - An identifier that is assigned to instruments of all asset classes and is unique to an individual instrument. Once issued, the FIGI assigned to an instrument will not change.
	ID_BB_SEC_NUM_DES                     //	Security ID Number Description - Descriptor for a financial instrument. Similar to the ticker field, but will provide additional metadata data.
	TICKER                                //	Ticker - Ticker is a specific identifier for a financial instrument that reflects common usage.
	BASE_TICKER                           //	An indistinct identifier which may be linked to multiple instruments. May need to be combined with other values to identify a unique instrument.
	ID_CUSIP_8_CHR                        //	CUSIP (8 Characters Only) - Committee on Uniform Securities Identification Procedures.
	OCC_SYMBOL                            //	OCC Symbol - A twenty-one character option symbol standardized by the Options Clearing Corporation (OCC) to identify a U.S. option.
	UNIQUE_ID_FUT_OPT                     //	Unique Identifier for Future Option - Bloomberg unique ticker with logic for index, currency, single stock futures, commodities and commodity options.
	OPRA_SYMBOL                           //	OPRA Symbol - Option symbol standardized by the Options Price Reporting Authority (OPRA) to identify a U.S. option.
	TRADING_SYSTEM_IDENTIFIER             //	Trading System Identifier - Unique identifier for the instrument as used on the source trading system.
	ID_SHORT_CODE                         //	An exchange venue specific code to identify fixed income instruments primarily traded in Asia.
	VENDOR_INDEX_CODE                     //	Index code assigned by the index provider for the purpose of identifying the security.
)

type MappingResponse struct {
	Error string
	Data  []struct {
		FIGI                 string `json:"figi"`
		SecurityType         string `json:"securityType"`
		MarketSector         string `json:"marketSector"`
		Ticker               string `json:"ticker"`
		Name                 string `json:"name"`
		UniqueID             string `json:"uniqueID"`
		ExchangeCode         string `json:"exchCode"`
		ShareClassFIGI       string `json:"shareClassFIGI"`
		CompositeFIGI        string `json:"compositeFIGI"`
		SecurityType2        string `json:"securityType2"`
		SecurityDescription  string `json:"securityDescription"`
		UniqueIDFutureOption string `json:"uniqueIDFutOpt"`
	}
}
