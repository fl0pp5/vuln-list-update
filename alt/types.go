package alt

import "encoding/xml"

type OVAL struct {
	XMLName     xml.Name    `xml:"oval_definitions"`
	Generator   Generator   `xml:"generator"`
	Definitions Definitions `xml:"definitions"`

	Tests   Tests   `xml:"tests" json:",omitempty"`
	Objects Objects `xml:"objects" json:",omitempty"`
	States  States  `xml:"states" json:",omitempty"`
}

type Generator struct {
	Timestamp     string `xml:"timestamp"`
	ProductName   string `xml:"product_name"`
	SchemaVersion string `xml:"schema_version"`
}

type Definitions struct {
	Definition []Definition `xml:"definition" json:"definition,omitempty"`
}

type Definition struct {
	ID       string   `xml:"id,attr" json:"id,omitempty"`
	Version  string   `xml:"version,attr" json:"version,omitempty"`
	Class    string   `xml:"class,attr" json:"class,omitempty"`
	Metadata Metadata `xml:"metadata" json:"metadata,omitempty"`
	Criteria Criteria `xml:"criteria" json:"criteria,omitempty"`
}

type Metadata struct {
	Title        string      `xml:"title" json:"title,omitempty"`
	AffectedList []Affected  `xml:"affected" json:"affected,omitempty"`
	References   []Reference `xml:"reference" json:"reference,omitempty"`
	Description  string      `xml:"description" json:"description,omitempty"`
	Advisory     Advisory    `xml:"advisory" json:"advisory,omitempty"`
}

type Affected struct {
	Family    string   `xml:"family,attr" json:"family,omitempty"`
	Platforms []string `xml:"platform" json:"platform,omitempty"`
	Products  []string `xml:"product" json:"product,omitempty"`
}

type Reference struct {
	RefID  string `xml:"ref_id,attr" json:"ref_id,omitempty"`
	RefURL string `xml:"ref_url,attr" json:"ref_url,omitempty"`
	Source string `xml:"source,attr" json:"source,omitempty"`
}

type Advisory struct {
	From            string          `xml:"from,attr" json:"from,omitempty"`
	Severity        string          `xml:"severity" json:"severity,omitempty"`
	Rights          string          `xml:"rights" json:"rights,omitempty"`
	Issued          Issued          `xml:"issued" json:"issued,omitempty"`
	Updated         Updated         `xml:"updated" json:"updated,omitempty"`
	BDUs            []CVE           `xml:"bdu" json:"bdu,omitempty"`
	Cves            []CVE           `xml:"cve" json:"cve,omitempty"`
	Bugzilla        []Bugzilla      `xml:"bugzilla" json:"bugzilla,omitempty"`
	AffectedCpeList AffectedCpeList `xml:"affected_cpe_list" json:"affected_cpe_list,omitempty"`
}

type Bugzilla struct {
	Id   string `xml:"id,attr" json:"id,omitempty"`
	Href string `xml:"href,attr" json:"href,omitempty"`
	Data string `xml:",chardata" json:"data:,omitempty"`
}

type Issued struct {
	Date string `xml:"date,attr" json:"date,omitempty"`
}

type Updated struct {
	Date string `xml:"date,attr" json:"date,omitempty"`
}

type CVE struct {
	Cvss   string `xml:"cvss,attr" json:"cvss,omitempty"`
	Cvss3  string `xml:"cvss3,attr" json:"cvss3,omitempty"`
	Cwe    string `xml:"cwe,attr"   json:"cwe,omitempty"`
	Href   string `xml:"href,attr"   json:"href,omitempty"`
	Impact string `xml:"impact,attr" json:"impact,omitempty"`
	Public string `xml:"public,attr" json:"public,omitempty"`
	CveID  string `xml:",chardata"  json:"cve_id,omitempty"`
}

type AffectedCpeList struct {
	Cpe []string `xml:"cpe" json:"cpe,omitempty"`
}

type Criteria struct {
	Operator   string      `xml:"operator,attr" json:"operator,omitempty"`
	Criterions []Criterion `xml:"criterion" json:"criterion,omitempty"`
	Criterias  []Criteria  `xml:"criteria" json:"criteria,omitempty"`
}

type Criterion struct {
	TestRef string `xml:"test_ref,attr" json:"test_ref,omitempty"`
	Comment string `xml:"comment,attr" json:"comment,omitempty"`
}

type Tests struct {
	TextFileContent54Tests []TextFileContent54Test `xml:"textfilecontent54_test" json:"textfilecontent54_test,omitempty"`
	RpmInfoTests           []RpmInfoTest           `xml:"rpminfo_test" json:"rpminfo_test,omitempty"`
}

type TextFileContent54Test struct {
	ID      string `xml:"id,attr" json:"id,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
	Check   string `xml:"check,attr" json:"check,omitempty"`
	Comment string `xml:"comment,attr" json:"comment,omitempty"`
	Object  Object `xml:"object" json:"object,omitempty"`
	State   State  `xml:"state" json:"state,omitempty"`
}

type State struct {
	StateRef string `xml:"state_ref,attr" json:"state_ref,omitempty"`
	Text     string `xml:"state" json:"state,omitempty"`
}

type Object struct {
	ObjectRef string `xml:"object_ref,attr" json:"object_ref,omitempty"`
	Text      string `xml:"object" json:"object,omitempty"`
}

type RpmInfoTest struct {
	ID      string `xml:"id,attr" json:"id,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
	Check   string `xml:"check,attr" json:"check,omitempty"`
	Comment string `xml:"comment,attr" json:"comment,omitempty"`
	Object  Object `xml:"object" json:"object,omitempty"`
	State   State  `xml:"state" json:"state,omitempty"`
}

type RpmInfoObject struct {
	ID      string `xml:"id,attr" json:"id,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
	Comment string `xml:"comment,attr" json:"comment,omitempty"`
	Name    string `xml:"name" json:"name,omitempty"`
}

type RpmInfoState struct {
	ID            string        `xml:"id,attr" json:"id,omitempty"`
	Version       string        `xml:"version,attr" json:"version,omitempty"`
	Comment       string        `xml:"comment,attr" json:"comment,omitempty"`
	Arch          Arch          `xml:"arch" json:"arch,omitempty"`
	Evr           Evr           `xml:"evr" json:"evr,omitempty"`
	Subexpression Subexpression `xml:"subexpression" json:"subexpression,omitempty"`
}

type Arch struct {
	Text      string `xml:",chardata" json:"text,omitempty"`
	Datatype  string `xml:"datatype,attr" json:"datatype,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
}

type Evr struct {
	Text      string `xml:",chardata" json:"text,omitempty"`
	Datatype  string `xml:"datatype,attr" json:"datatype,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
}

type Subexpression struct {
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
	Text      string `xml:",chardata" json:"text,omitempty"`
}

type Objects struct {
	TextFileContent54Objects []TextFileContent54Object `xml:"textfilecontent54_object" json:"textfilecontent54_object,omitempty"`
	RpmInfoObjects           []RpmInfoObject           `xml:"rpminfo_object" json:"rpminfo_object,omitempty"`
}

type TextFileContent54Object struct {
	ID       string   `xml:"id,attr" json:"id,omitempty"`
	Version  string   `xml:"version,attr" json:"version,omitempty"`
	Comment  string   `xml:"comment,attr" json:"comment,omitempty"`
	Path     Path     `xml:"path" json:"path,omitempty"`
	Filepath Filepath `xml:"filepath" json:"filepath,omitempty"`
	Pattern  Pattern  `xml:"pattern" json:"pattern,omitempty"`
	Instance Instance `xml:"instance" json:"instance,omitempty"`
}

type Path struct {
	Datatype string `xml:"datatype,attr" json:"datatype,omitempty"`
	Text     string `xml:",chardata" json:"text,omitempty"`
}

type Filepath struct {
	Datatype string `xml:"datatype,attr" json:"datatype,omitempty"`
	Text     string `xml:",chardata" json:"text,omitempty"`
}

type Pattern struct {
	Datatype  string `xml:"datatype,attr" json:"datatype,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
	Text      string `xml:",chardata" json:"text,omitempty"`
}

type Instance struct {
	Datatype string `xml:"datatype,attr" json:"datatype,omitempty"`
	Text     string `xml:",chardata" json:"text,omitempty"`
}

type Name struct {
	Text      string `xml:",chardata" json:"text,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
}

type States struct {
	TextFileContent54State []TextFileContent54State `xml:"textfilecontent54_state" json:"textfilecontent54_state,omitempty"`
	RpmInfoState           []RpmInfoState           `xml:"rpminfo_state" json:"rpminfo_state,omitempty"`
}

type Version struct {
	Text      string `xml:",chardata" json:"text,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
}

type TextFileContent54State struct {
	ID      string `xml:"id,attr" json:"id,omitempty"`
	Version string `xml:"version,attr" json:"version,omitempty"`
	Text    Text   `xml:"text" json:"text,omitempty"`
}

type Text struct {
	Text      string `xml:",chardata" json:"text,omitempty"`
	Operation string `xml:"operation,attr" json:"operation,omitempty"`
}
