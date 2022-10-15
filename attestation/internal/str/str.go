package str

import (
	"attestation/internal/bill"
	"attestation/internal/inc"
	"attestation/internal/mail"
	"attestation/internal/mms"
	"attestation/internal/sms"
	"attestation/internal/supp"
	"attestation/internal/vc"
	"fmt"
)

var (
	status   = true
	resError = ""
)

type ResultT struct {
	Status bool       `json:"status"` // true if all data collection steps were successful, false in all other cases
	Data   ResultSetT `json:"data"`   // filled if all data collection steps were successful, nil in all other cases
	Error  string     `json:"error"`  // empty string if all stages of data collection were successful, in case of an error it is filled with the error text
}

func NewResultT() *ResultT {
	return &ResultT{}

}

type ResultSetT struct {
	SMS       [][]sms.SMSData               `json:"sms"`
	MMS       [][]mms.MMSData               `json:"mms"`
	VoiceCall []vc.VoiceCallData            `json:"voice_call"`
	Email     map[string][][]mail.EmailData `json:"email"`
	Billing   bill.BillingData              `json:"billing"`
	Support   []int                         `json:"support"`
	Incidents []inc.IncidentData            `json:"incident"`
}

func NewResultSetT() *ResultSetT {
	return &ResultSetT{}
}

// GetResultData - bypass all data collection steps and return a ready-to-send dataset
func GetResultData(alphaCodes map[string]string, host, simulatorAddr string) *ResultT {
	res := NewResultT()
	res.Data = *NewResultSetT()

	// Get structured sms-data
	smsSortedByProvider, smsSortedByCountry, err := sms.GetStructuredData(alphaCodes)
	checkSMSData(smsSortedByProvider, smsSortedByCountry, err)

	// Add sms-data to the result-struct
	res.Data.SMS = append(res.Data.SMS, smsSortedByProvider)
	res.Data.SMS = append(res.Data.SMS, smsSortedByCountry)

	// Get structured mms-data
	mmsSortedByProvider, mmsSortedByCountry, err := mms.GetStructuredData(alphaCodes, host, simulatorAddr)
	checkMMSData(mmsSortedByProvider, mmsSortedByCountry, err)

	// Add mms-data to the result-struct
	res.Data.MMS = append(res.Data.MMS, mmsSortedByProvider)
	res.Data.MMS = append(res.Data.MMS, mmsSortedByCountry)

	// Get structured voice-call-data
	vcData, err := vc.GetStatus(alphaCodes)
	checkVcData(vcData, err)

	// Add voice-call-data to the result-struct
	res.Data.VoiceCall = vcData

	// Get structured email-data
	emailData, err := mail.GetStructuredData(alphaCodes)
	checkEmailData(emailData, err)

	// Add email-data to the result-struct
	res.Data.Email = emailData

	// Get structured billing-data
	billData, err := bill.GetStatus()
	checkBillData(billData, err)

	// Add billing-data to the result-struct
	res.Data.Billing = billData

	// Get structured support-data
	load, waitingTime, err := supp.GetStructuredData(host, simulatorAddr)
	checkSuppData(load, waitingTime, err)

	// Add support-data to the result-struct
	res.Data.Support = append(res.Data.Support, load)
	res.Data.Support = append(res.Data.Support, waitingTime)

	// Get structured accident-data
	incData, err := inc.GetStructuredData(host, simulatorAddr)
	checkIncData(incData, err)

	// Add accident data to the result-struct
	res.Data.Incidents = incData

	res.Status = status
	res.Error = resError

	return res
}

// checkSMSData - checks for sms-data collection errors
func checkSMSData(smsSortedByProvider, smsSortedByCountry []sms.SMSData, err error) {
	if smsSortedByProvider == nil || smsSortedByCountry == nil ||
		len(smsSortedByProvider) == 0 || len(smsSortedByCountry) == 0 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// checkMMSData - checks for sms-data collection errors
func checkMMSData(mmsSortedByProvider, mmsSortedByCountry []mms.MMSData, err error) {
	if mmsSortedByProvider == nil || mmsSortedByCountry == nil ||
		len(mmsSortedByProvider) == 0 || len(mmsSortedByCountry) == 0 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// checkVcData - check for voice-call-data collection errors
func checkVcData(vcData []vc.VoiceCallData, err error) {
	if vcData == nil || len(vcData) == 0 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// Check for email-data collection errors
func checkEmailData(emailData map[string][][]mail.EmailData, err error) {
	if emailData == nil || len(emailData) == 0 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// Check for billing-data collection errors
func checkBillData(billData bill.BillingData, err error) {
	if billData == *bill.NewBillingData() {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// Check for support-data collection errors
func checkSuppData(load, waitingTime int, err error) {
	if load == -1 || waitingTime == -1 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}

// Check for accident-data collection errors
func checkIncData(incData []inc.IncidentData, err error) {
	if incData == nil || len(incData) == 0 {
		status = false
	}
	if err != nil {
		resError = fmt.Sprintf("%s\n%s", resError, err)
	}
}
