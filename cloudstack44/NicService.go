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
)

type AddIpToNicParams struct {
	p map[string]interface{}
}

func (p *AddIpToNicParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["ipaddress"]; found {
		u.Set("ipaddress", v.(string))
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	return u
}

func (p *AddIpToNicParams) SetIpaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipaddress"] = v
	return
}

func (p *AddIpToNicParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
	return
}

// You should always use this function to get a new AddIpToNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewAddIpToNicParams(nicid string) *AddIpToNicParams {
	p := &AddIpToNicParams{}
	p.p = make(map[string]interface{})
	p.p["nicid"] = nicid
	return p
}

// Assigns secondary IP to NIC
func (s *NicService) AddIpToNic(p *AddIpToNicParams) (*AddIpToNicResponse, error) {
	resp, err := s.cs.newRequest("addIpToNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddIpToNicResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r AddIpToNicResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddIpToNicResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Virtualmachineid string `json:"virtualmachineid,omitempty"`
	Id               string `json:"id,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Networkid        string `json:"networkid,omitempty"`
	Nicid            string `json:"nicid,omitempty"`
}

type RemoveIpFromNicParams struct {
	p map[string]interface{}
}

func (p *RemoveIpFromNicParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RemoveIpFromNicParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RemoveIpFromNicParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewRemoveIpFromNicParams(id string) *RemoveIpFromNicParams {
	p := &RemoveIpFromNicParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Removes secondary IP from the NIC.
func (s *NicService) RemoveIpFromNic(p *RemoveIpFromNicParams) (*RemoveIpFromNicResponse, error) {
	resp, err := s.cs.newRequest("removeIpFromNic", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveIpFromNicResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		var r RemoveIpFromNicResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RemoveIpFromNicResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListNicsParams struct {
	p map[string]interface{}
}

func (p *ListNicsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["nicid"]; found {
		u.Set("nicid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *ListNicsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNicsParams) SetNicid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["nicid"] = v
	return
}

func (p *ListNicsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNicsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNicsParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new ListNicsParams instance,
// as then you are sure you have configured all required params
func (s *NicService) NewListNicsParams(virtualmachineid string) *ListNicsParams {
	p := &ListNicsParams{}
	p.p = make(map[string]interface{})
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *NicService) GetNicID(keyword string, virtualmachineid string) (string, error) {
	p := &ListNicsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword
	p.p["virtualmachineid"] = virtualmachineid

	l, err := s.ListNics(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.Nics[0].Id, nil
}

// list the vm nics  IP to NIC
func (s *NicService) ListNics(p *ListNicsParams) (*ListNicsResponse, error) {
	resp, err := s.cs.newRequest("listNics", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNicsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNicsResponse struct {
	Count int    `json:"count"`
	Nics  []*Nic `json:"nic"`
}

type Nic struct {
	Isdefault    bool     `json:"isdefault,omitempty"`
	Netmask      string   `json:"netmask,omitempty"`
	Traffictype  string   `json:"traffictype,omitempty"`
	Ip6address   string   `json:"ip6address,omitempty"`
	Ipaddress    string   `json:"ipaddress,omitempty"`
	Ip6gateway   string   `json:"ip6gateway,omitempty"`
	Secondaryip  []string `json:"secondaryip,omitempty"`
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Macaddress   string   `json:"macaddress,omitempty"`
	Networkname  string   `json:"networkname,omitempty"`
	Ip6cidr      string   `json:"ip6cidr,omitempty"`
	Gateway      string   `json:"gateway,omitempty"`
	Broadcasturi string   `json:"broadcasturi,omitempty"`
	Networkid    string   `json:"networkid,omitempty"`
	Isolationuri string   `json:"isolationuri,omitempty"`
}
