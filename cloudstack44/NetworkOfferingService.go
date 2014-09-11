//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack44

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type CreateNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *CreateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["conservemode"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("conservemode", vv)
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["egressdefaultpolicy"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("egressdefaultpolicy", vv)
	}
	if v, found := p.p["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := p.p["ispersistent"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("ispersistent", vv)
	}
	if v, found := p.p["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := p.p["maxconnections"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxconnections", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkrate"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("networkrate", vv)
	}
	if v, found := p.p["servicecapabilitylist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].key", i), k)
			u.Set(fmt.Sprintf("servicecapabilitylist[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	if v, found := p.p["serviceproviderlist"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("serviceproviderlist[%d].key", i), k)
			u.Set(fmt.Sprintf("serviceproviderlist[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := p.p["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	return u
}

func (p *CreateNetworkOfferingParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetConservemode(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["conservemode"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetEgressdefaultpolicy(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["egressdefaultpolicy"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetGuestiptype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestiptype"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetIspersistent(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ispersistent"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keepaliveenabled"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetMaxconnections(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxconnections"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetNetworkrate(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkrate"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetServicecapabilitylist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["servicecapabilitylist"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetServiceproviderlist(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceproviderlist"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetSpecifyvlan(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyvlan"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *CreateNetworkOfferingParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

// You should always use this function to get a new CreateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewCreateNetworkOfferingParams(displaytext string, guestiptype string, name string, supportedservices []string, traffictype string) *CreateNetworkOfferingParams {
	p := &CreateNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["displaytext"] = displaytext
	p.p["guestiptype"] = guestiptype
	p.p["name"] = name
	p.p["supportedservices"] = supportedservices
	p.p["traffictype"] = traffictype
	return p
}

// Creates a network offering.
func (s *NetworkOfferingService) CreateNetworkOffering(p *CreateNetworkOfferingParams) (*CreateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("createNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type CreateNetworkOfferingResponse struct {
	Egressdefaultpolicy bool              `json:"egressdefaultpolicy,omitempty"`
	Specifyvlan         bool              `json:"specifyvlan,omitempty"`
	Created             string            `json:"created,omitempty"`
	Availability        string            `json:"availability,omitempty"`
	Networkrate         int               `json:"networkrate,omitempty"`
	Tags                string            `json:"tags,omitempty"`
	State               string            `json:"state,omitempty"`
	Name                string            `json:"name,omitempty"`
	Guestiptype         string            `json:"guestiptype,omitempty"`
	Forvpc              bool              `json:"forvpc,omitempty"`
	Maxconnections      int               `json:"maxconnections,omitempty"`
	Id                  string            `json:"id,omitempty"`
	Ispersistent        bool              `json:"ispersistent,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Isdefault           bool              `json:"isdefault,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Displaytext         string            `json:"displaytext,omitempty"`
	Traffictype         string            `json:"traffictype,omitempty"`
	Service             []struct {
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Value                      string `json:"value,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Provider []struct {
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Name                         string   `json:"name,omitempty"`
		} `json:"provider,omitempty"`
	} `json:"service,omitempty"`
	Specifyipranges bool `json:"specifyipranges,omitempty"`
	Conservemode    bool `json:"conservemode,omitempty"`
}

type UpdateNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *UpdateNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keepaliveenabled"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("keepaliveenabled", vv)
	}
	if v, found := p.p["maxconnections"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("maxconnections", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["sortkey"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("sortkey", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	return u
}

func (p *UpdateNetworkOfferingParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetKeepaliveenabled(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keepaliveenabled"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetMaxconnections(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["maxconnections"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetSortkey(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sortkey"] = v
	return
}

func (p *UpdateNetworkOfferingParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

// You should always use this function to get a new UpdateNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewUpdateNetworkOfferingParams() *UpdateNetworkOfferingParams {
	p := &UpdateNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	return p
}

// Updates a network offering.
func (s *NetworkOfferingService) UpdateNetworkOffering(p *UpdateNetworkOfferingParams) (*UpdateNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("updateNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UpdateNetworkOfferingResponse struct {
	Name           string `json:"name,omitempty"`
	Maxconnections int    `json:"maxconnections,omitempty"`
	Traffictype    string `json:"traffictype,omitempty"`
	Id             string `json:"id,omitempty"`
	Displaytext    string `json:"displaytext,omitempty"`
	Created        string `json:"created,omitempty"`
	Specifyvlan    bool   `json:"specifyvlan,omitempty"`
	State          string `json:"state,omitempty"`
	Service        []struct {
		Provider []struct {
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Name                         string   `json:"name,omitempty"`
			Id                           string   `json:"id,omitempty"`
			State                        string   `json:"state,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
		} `json:"provider,omitempty"`
		Name       string `json:"name,omitempty"`
		Capability []struct {
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
			Value                      string `json:"value,omitempty"`
		} `json:"capability,omitempty"`
	} `json:"service,omitempty"`
	Specifyipranges     bool              `json:"specifyipranges,omitempty"`
	Availability        string            `json:"availability,omitempty"`
	Guestiptype         string            `json:"guestiptype,omitempty"`
	Tags                string            `json:"tags,omitempty"`
	Egressdefaultpolicy bool              `json:"egressdefaultpolicy,omitempty"`
	Conservemode        bool              `json:"conservemode,omitempty"`
	Networkrate         int               `json:"networkrate,omitempty"`
	Ispersistent        bool              `json:"ispersistent,omitempty"`
	Forvpc              bool              `json:"forvpc,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Isdefault           bool              `json:"isdefault,omitempty"`
}

type DeleteNetworkOfferingParams struct {
	p map[string]interface{}
}

func (p *DeleteNetworkOfferingParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteNetworkOfferingParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteNetworkOfferingParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewDeleteNetworkOfferingParams(id string) *DeleteNetworkOfferingParams {
	p := &DeleteNetworkOfferingParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a network offering.
func (s *NetworkOfferingService) DeleteNetworkOffering(p *DeleteNetworkOfferingParams) (*DeleteNetworkOfferingResponse, error) {
	resp, err := s.cs.newRequest("deleteNetworkOffering", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetworkOfferingResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteNetworkOfferingResponse struct {
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListNetworkOfferingsParams struct {
	p map[string]interface{}
}

func (p *ListNetworkOfferingsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["availability"]; found {
		u.Set("availability", v.(string))
	}
	if v, found := p.p["displaytext"]; found {
		u.Set("displaytext", v.(string))
	}
	if v, found := p.p["forvpc"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forvpc", vv)
	}
	if v, found := p.p["guestiptype"]; found {
		u.Set("guestiptype", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isdefault"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isdefault", vv)
	}
	if v, found := p.p["istagged"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("istagged", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["sourcenatsupported"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("sourcenatsupported", vv)
	}
	if v, found := p.p["specifyipranges"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyipranges", vv)
	}
	if v, found := p.p["specifyvlan"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("specifyvlan", vv)
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["supportedservices"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("supportedservices", vv)
	}
	if v, found := p.p["tags"]; found {
		u.Set("tags", v.(string))
	}
	if v, found := p.p["traffictype"]; found {
		u.Set("traffictype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListNetworkOfferingsParams) SetAvailability(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["availability"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetDisplaytext(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["displaytext"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetForvpc(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forvpc"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetGuestiptype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["guestiptype"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetIsdefault(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isdefault"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetIstagged(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["istagged"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetSourcenatsupported(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourcenatsupported"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetSpecifyipranges(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyipranges"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetSpecifyvlan(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["specifyvlan"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetSupportedservices(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["supportedservices"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetTags(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetTraffictype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["traffictype"] = v
	return
}

func (p *ListNetworkOfferingsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListNetworkOfferingsParams instance,
// as then you are sure you have configured all required params
func (s *NetworkOfferingService) NewListNetworkOfferingsParams() *ListNetworkOfferingsParams {
	p := &ListNetworkOfferingsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NetworkOfferingService) GetNetworkOfferingID(name string) (string, error) {
	p := &ListNetworkOfferingsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListNetworkOfferings(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, name, l)
	}
	return l.NetworkOfferings[0].Id, nil
}

// Lists all available network offerings.
func (s *NetworkOfferingService) ListNetworkOfferings(p *ListNetworkOfferingsParams) (*ListNetworkOfferingsResponse, error) {
	resp, err := s.cs.newRequest("listNetworkOfferings", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetworkOfferingsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetworkOfferingsResponse struct {
	Count            int                `json:"count"`
	NetworkOfferings []*NetworkOffering `json:"networkoffering"`
}

type NetworkOffering struct {
	Traffictype         string            `json:"traffictype,omitempty"`
	Specifyipranges     bool              `json:"specifyipranges,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Specifyvlan         bool              `json:"specifyvlan,omitempty"`
	Availability        string            `json:"availability,omitempty"`
	Displaytext         string            `json:"displaytext,omitempty"`
	Guestiptype         string            `json:"guestiptype,omitempty"`
	State               string            `json:"state,omitempty"`
	Egressdefaultpolicy bool              `json:"egressdefaultpolicy,omitempty"`
	Tags                string            `json:"tags,omitempty"`
	Created             string            `json:"created,omitempty"`
	Service             []struct {
		Provider []struct {
			State                        string   `json:"state,omitempty"`
			Id                           string   `json:"id,omitempty"`
			Destinationphysicalnetworkid string   `json:"destinationphysicalnetworkid,omitempty"`
			Physicalnetworkid            string   `json:"physicalnetworkid,omitempty"`
			Canenableindividualservice   bool     `json:"canenableindividualservice,omitempty"`
			Servicelist                  []string `json:"servicelist,omitempty"`
			Name                         string   `json:"name,omitempty"`
		} `json:"provider,omitempty"`
		Capability []struct {
			Value                      string `json:"value,omitempty"`
			Canchooseservicecapability bool   `json:"canchooseservicecapability,omitempty"`
			Name                       string `json:"name,omitempty"`
		} `json:"capability,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"service,omitempty"`
	Id             string `json:"id,omitempty"`
	Conservemode   bool   `json:"conservemode,omitempty"`
	Isdefault      bool   `json:"isdefault,omitempty"`
	Name           string `json:"name,omitempty"`
	Forvpc         bool   `json:"forvpc,omitempty"`
	Networkrate    int    `json:"networkrate,omitempty"`
	Maxconnections int    `json:"maxconnections,omitempty"`
	Ispersistent   bool   `json:"ispersistent,omitempty"`
}
