/*
 * Red Hat Security Data API
 *
 * Unofficial OpenAPI definitions for Red Hat Security Data API
 *
 * API version: 1.0
 * Contact: mustafa@ctrliq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rhsecurity

import (
	"encoding/json"
)

// CVE CVE model used in listing
type CVE struct {
	CVE                 string   `json:"CVE"`
	Severity            string   `json:"severity"`
	PublicDate          string   `json:"public_date"`
	Advisories          []string `json:"advisories"`
	Bugzilla            string   `json:"bugzilla"`
	BugzillaDescription string   `json:"bugzilla_description"`
	CvssScore           *float32 `json:"cvss_score,omitempty"`
	CvssScoringVector   *string  `json:"cvss_scoring_vector,omitempty"`
	CWE                 string   `json:"CWE"`
	AffectedPackages    []string `json:"affected_packages"`
	ResourceUrl         string   `json:"resource_url"`
	Cvss3ScoringVector  string   `json:"cvss3_scoring_vector"`
	Cvss3Score          string   `json:"cvss3_score"`
}

// NewCVE instantiates a new CVE object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCVE(cVE string, severity string, publicDate string, advisories []string, bugzilla string, bugzillaDescription string, cWE string, affectedPackages []string, resourceUrl string, cvss3ScoringVector string, cvss3Score string) *CVE {
	this := CVE{}
	this.CVE = cVE
	this.Severity = severity
	this.PublicDate = publicDate
	this.Advisories = advisories
	this.Bugzilla = bugzilla
	this.BugzillaDescription = bugzillaDescription
	this.CWE = cWE
	this.AffectedPackages = affectedPackages
	this.ResourceUrl = resourceUrl
	this.Cvss3ScoringVector = cvss3ScoringVector
	this.Cvss3Score = cvss3Score
	return &this
}

// NewCVEWithDefaults instantiates a new CVE object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCVEWithDefaults() *CVE {
	this := CVE{}
	return &this
}

// GetCVE returns the CVE field value
func (o *CVE) GetCVE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CVE
}

// GetCVEOk returns a tuple with the CVE field value
// and a boolean to check if the value has been set.
func (o *CVE) GetCVEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CVE, true
}

// SetCVE sets field value
func (o *CVE) SetCVE(v string) {
	o.CVE = v
}

// GetSeverity returns the Severity field value
func (o *CVE) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *CVE) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *CVE) SetSeverity(v string) {
	o.Severity = v
}

// GetPublicDate returns the PublicDate field value
func (o *CVE) GetPublicDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicDate
}

// GetPublicDateOk returns a tuple with the PublicDate field value
// and a boolean to check if the value has been set.
func (o *CVE) GetPublicDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicDate, true
}

// SetPublicDate sets field value
func (o *CVE) SetPublicDate(v string) {
	o.PublicDate = v
}

// GetAdvisories returns the Advisories field value
func (o *CVE) GetAdvisories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Advisories
}

// GetAdvisoriesOk returns a tuple with the Advisories field value
// and a boolean to check if the value has been set.
func (o *CVE) GetAdvisoriesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advisories, true
}

// SetAdvisories sets field value
func (o *CVE) SetAdvisories(v []string) {
	o.Advisories = v
}

// GetBugzilla returns the Bugzilla field value
func (o *CVE) GetBugzilla() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bugzilla
}

// GetBugzillaOk returns a tuple with the Bugzilla field value
// and a boolean to check if the value has been set.
func (o *CVE) GetBugzillaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bugzilla, true
}

// SetBugzilla sets field value
func (o *CVE) SetBugzilla(v string) {
	o.Bugzilla = v
}

// GetBugzillaDescription returns the BugzillaDescription field value
func (o *CVE) GetBugzillaDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BugzillaDescription
}

// GetBugzillaDescriptionOk returns a tuple with the BugzillaDescription field value
// and a boolean to check if the value has been set.
func (o *CVE) GetBugzillaDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BugzillaDescription, true
}

// SetBugzillaDescription sets field value
func (o *CVE) SetBugzillaDescription(v string) {
	o.BugzillaDescription = v
}

// GetCvssScore returns the CvssScore field value if set, zero value otherwise.
func (o *CVE) GetCvssScore() float32 {
	if o == nil || o.CvssScore == nil {
		var ret float32
		return ret
	}
	return *o.CvssScore
}

// GetCvssScoreOk returns a tuple with the CvssScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CVE) GetCvssScoreOk() (*float32, bool) {
	if o == nil || o.CvssScore == nil {
		return nil, false
	}
	return o.CvssScore, true
}

// HasCvssScore returns a boolean if a field has been set.
func (o *CVE) HasCvssScore() bool {
	if o != nil && o.CvssScore != nil {
		return true
	}

	return false
}

// SetCvssScore gets a reference to the given float32 and assigns it to the CvssScore field.
func (o *CVE) SetCvssScore(v float32) {
	o.CvssScore = &v
}

// GetCvssScoringVector returns the CvssScoringVector field value if set, zero value otherwise.
func (o *CVE) GetCvssScoringVector() string {
	if o == nil || o.CvssScoringVector == nil {
		var ret string
		return ret
	}
	return *o.CvssScoringVector
}

// GetCvssScoringVectorOk returns a tuple with the CvssScoringVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CVE) GetCvssScoringVectorOk() (*string, bool) {
	if o == nil || o.CvssScoringVector == nil {
		return nil, false
	}
	return o.CvssScoringVector, true
}

// HasCvssScoringVector returns a boolean if a field has been set.
func (o *CVE) HasCvssScoringVector() bool {
	if o != nil && o.CvssScoringVector != nil {
		return true
	}

	return false
}

// SetCvssScoringVector gets a reference to the given string and assigns it to the CvssScoringVector field.
func (o *CVE) SetCvssScoringVector(v string) {
	o.CvssScoringVector = &v
}

// GetCWE returns the CWE field value
func (o *CVE) GetCWE() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CWE
}

// GetCWEOk returns a tuple with the CWE field value
// and a boolean to check if the value has been set.
func (o *CVE) GetCWEOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CWE, true
}

// SetCWE sets field value
func (o *CVE) SetCWE(v string) {
	o.CWE = v
}

// GetAffectedPackages returns the AffectedPackages field value
func (o *CVE) GetAffectedPackages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AffectedPackages
}

// GetAffectedPackagesOk returns a tuple with the AffectedPackages field value
// and a boolean to check if the value has been set.
func (o *CVE) GetAffectedPackagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AffectedPackages, true
}

// SetAffectedPackages sets field value
func (o *CVE) SetAffectedPackages(v []string) {
	o.AffectedPackages = v
}

// GetResourceUrl returns the ResourceUrl field value
func (o *CVE) GetResourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceUrl
}

// GetResourceUrlOk returns a tuple with the ResourceUrl field value
// and a boolean to check if the value has been set.
func (o *CVE) GetResourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceUrl, true
}

// SetResourceUrl sets field value
func (o *CVE) SetResourceUrl(v string) {
	o.ResourceUrl = v
}

// GetCvss3ScoringVector returns the Cvss3ScoringVector field value
func (o *CVE) GetCvss3ScoringVector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvss3ScoringVector
}

// GetCvss3ScoringVectorOk returns a tuple with the Cvss3ScoringVector field value
// and a boolean to check if the value has been set.
func (o *CVE) GetCvss3ScoringVectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvss3ScoringVector, true
}

// SetCvss3ScoringVector sets field value
func (o *CVE) SetCvss3ScoringVector(v string) {
	o.Cvss3ScoringVector = v
}

// GetCvss3Score returns the Cvss3Score field value
func (o *CVE) GetCvss3Score() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvss3Score
}

// GetCvss3ScoreOk returns a tuple with the Cvss3Score field value
// and a boolean to check if the value has been set.
func (o *CVE) GetCvss3ScoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvss3Score, true
}

// SetCvss3Score sets field value
func (o *CVE) SetCvss3Score(v string) {
	o.Cvss3Score = v
}

func (o CVE) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CVE"] = o.CVE
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["public_date"] = o.PublicDate
	}
	if true {
		toSerialize["advisories"] = o.Advisories
	}
	if true {
		toSerialize["bugzilla"] = o.Bugzilla
	}
	if true {
		toSerialize["bugzilla_description"] = o.BugzillaDescription
	}
	if o.CvssScore != nil {
		toSerialize["cvss_score"] = o.CvssScore
	}
	if o.CvssScoringVector != nil {
		toSerialize["cvss_scoring_vector"] = o.CvssScoringVector
	}
	if true {
		toSerialize["CWE"] = o.CWE
	}
	if true {
		toSerialize["affected_packages"] = o.AffectedPackages
	}
	if true {
		toSerialize["resource_url"] = o.ResourceUrl
	}
	if true {
		toSerialize["cvss3_scoring_vector"] = o.Cvss3ScoringVector
	}
	if true {
		toSerialize["cvss3_score"] = o.Cvss3Score
	}
	return json.Marshal(toSerialize)
}

type NullableCVE struct {
	value *CVE
	isSet bool
}

func (v NullableCVE) Get() *CVE {
	return v.value
}

func (v *NullableCVE) Set(val *CVE) {
	v.value = val
	v.isSet = true
}

func (v NullableCVE) IsSet() bool {
	return v.isSet
}

func (v *NullableCVE) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCVE(val *CVE) *NullableCVE {
	return &NullableCVE{value: val, isSet: true}
}

func (v NullableCVE) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCVE) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
