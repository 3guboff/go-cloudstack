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

package cloudstack43

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type CreateRemoteAccessVpnParams struct {
	p map[string]interface{}
}

func (p *CreateRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["iprange"]; found {
		u.Set("iprange", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *CreateRemoteAccessVpnParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateRemoteAccessVpnParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateRemoteAccessVpnParams) SetIprange(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["iprange"] = v
	return
}

func (p *CreateRemoteAccessVpnParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
	return
}

func (p *CreateRemoteAccessVpnParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

// You should always use this function to get a new CreateRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateRemoteAccessVpnParams(publicipid string) *CreateRemoteAccessVpnParams {
	p := &CreateRemoteAccessVpnParams{}
	p.p = make(map[string]interface{})
	p.p["publicipid"] = publicipid
	return p
}

// Creates a l2tp/ipsec remote access vpn
func (s *VPNService) CreateRemoteAccessVpn(p *CreateRemoteAccessVpnParams) (*CreateRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newRequest("createRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateRemoteAccessVpnResponse
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

		var r CreateRemoteAccessVpnResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateRemoteAccessVpnResponse struct {
	JobID        string `json:"jobid,omitempty"`
	Id           string `json:"id,omitempty"`
	State        string `json:"state,omitempty"`
	Publicipid   string `json:"publicipid,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Publicip     string `json:"publicip,omitempty"`
	Presharedkey string `json:"presharedkey,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Account      string `json:"account,omitempty"`
	Project      string `json:"project,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Iprange      string `json:"iprange,omitempty"`
}

type DeleteRemoteAccessVpnParams struct {
	p map[string]interface{}
}

func (p *DeleteRemoteAccessVpnParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *DeleteRemoteAccessVpnParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

// You should always use this function to get a new DeleteRemoteAccessVpnParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteRemoteAccessVpnParams(publicipid string) *DeleteRemoteAccessVpnParams {
	p := &DeleteRemoteAccessVpnParams{}
	p.p = make(map[string]interface{})
	p.p["publicipid"] = publicipid
	return p
}

// Destroys a l2tp/ipsec remote access vpn
func (s *VPNService) DeleteRemoteAccessVpn(p *DeleteRemoteAccessVpnParams) (*DeleteRemoteAccessVpnResponse, error) {
	resp, err := s.cs.newRequest("deleteRemoteAccessVpn", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteRemoteAccessVpnResponse
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

		var r DeleteRemoteAccessVpnResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteRemoteAccessVpnResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListRemoteAccessVpnsParams struct {
	p map[string]interface{}
}

func (p *ListRemoteAccessVpnsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
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
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	return u
}

func (p *ListRemoteAccessVpnsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListRemoteAccessVpnsParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

// You should always use this function to get a new ListRemoteAccessVpnsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListRemoteAccessVpnsParams() *ListRemoteAccessVpnsParams {
	p := &ListRemoteAccessVpnsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetRemoteAccessVpnID(keyword string) (string, error) {
	p := &ListRemoteAccessVpnsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListRemoteAccessVpns(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.RemoteAccessVpns[0].Id, nil
}

// Lists remote access vpns
func (s *VPNService) ListRemoteAccessVpns(p *ListRemoteAccessVpnsParams) (*ListRemoteAccessVpnsResponse, error) {
	resp, err := s.cs.newRequest("listRemoteAccessVpns", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListRemoteAccessVpnsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListRemoteAccessVpnsResponse struct {
	Count            int                `json:"count"`
	RemoteAccessVpns []*RemoteAccessVpn `json:"remoteaccessvpn"`
}

type RemoteAccessVpn struct {
	Id           string `json:"id,omitempty"`
	Publicip     string `json:"publicip,omitempty"`
	State        string `json:"state,omitempty"`
	Iprange      string `json:"iprange,omitempty"`
	Account      string `json:"account,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Project      string `json:"project,omitempty"`
	Publicipid   string `json:"publicipid,omitempty"`
	Domainid     string `json:"domainid,omitempty"`
	Projectid    string `json:"projectid,omitempty"`
	Presharedkey string `json:"presharedkey,omitempty"`
}

type AddVpnUserParams struct {
	p map[string]interface{}
}

func (p *AddVpnUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddVpnUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *AddVpnUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *AddVpnUserParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddVpnUserParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *AddVpnUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewAddVpnUserParams(password string, username string) *AddVpnUserParams {
	p := &AddVpnUserParams{}
	p.p = make(map[string]interface{})
	p.p["password"] = password
	p.p["username"] = username
	return p
}

// Adds vpn users
func (s *VPNService) AddVpnUser(p *AddVpnUserParams) (*AddVpnUserResponse, error) {
	resp, err := s.cs.newRequest("addVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddVpnUserResponse
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

		var r AddVpnUserResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type AddVpnUserResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Account   string `json:"account,omitempty"`
	Id        string `json:"id,omitempty"`
	State     string `json:"state,omitempty"`
	Project   string `json:"project,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Username  string `json:"username,omitempty"`
}

type RemoveVpnUserParams struct {
	p map[string]interface{}
}

func (p *RemoveVpnUserParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *RemoveVpnUserParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *RemoveVpnUserParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *RemoveVpnUserParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *RemoveVpnUserParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new RemoveVpnUserParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewRemoveVpnUserParams(username string) *RemoveVpnUserParams {
	p := &RemoveVpnUserParams{}
	p.p = make(map[string]interface{})
	p.p["username"] = username
	return p
}

// Removes vpn user
func (s *VPNService) RemoveVpnUser(p *RemoveVpnUserParams) (*RemoveVpnUserResponse, error) {
	resp, err := s.cs.newRequest("removeVpnUser", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveVpnUserResponse
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

		var r RemoveVpnUserResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type RemoveVpnUserResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListVpnUsersParams struct {
	p map[string]interface{}
}

func (p *ListVpnUsersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *ListVpnUsersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVpnUsersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVpnUsersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVpnUsersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVpnUsersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVpnUsersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVpnUsersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVpnUsersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVpnUsersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVpnUsersParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new ListVpnUsersParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnUsersParams() *ListVpnUsersParams {
	p := &ListVpnUsersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnUserID(keyword string) (string, error) {
	p := &ListVpnUsersParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListVpnUsers(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.VpnUsers[0].Id, nil
}

// Lists vpn users
func (s *VPNService) ListVpnUsers(p *ListVpnUsersParams) (*ListVpnUsersResponse, error) {
	resp, err := s.cs.newRequest("listVpnUsers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnUsersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVpnUsersResponse struct {
	Count    int        `json:"count"`
	VpnUsers []*VpnUser `json:"vpnuser"`
}

type VpnUser struct {
	Domainid  string `json:"domainid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Username  string `json:"username,omitempty"`
	Account   string `json:"account,omitempty"`
	Id        string `json:"id,omitempty"`
	Project   string `json:"project,omitempty"`
	Domain    string `json:"domain,omitempty"`
	State     string `json:"state,omitempty"`
}

type CreateVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *CreateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := p.p["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := p.p["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := p.p["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := p.p["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *CreateVpnCustomerGatewayParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetDpd(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dpd"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetEsplifetime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esplifetime"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esppolicy"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetIkelifetime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikelifetime"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikepolicy"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipsecpsk"] = v
	return
}

func (p *CreateVpnCustomerGatewayParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new CreateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, ikepolicy string, ipsecpsk string) *CreateVpnCustomerGatewayParams {
	p := &CreateVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["cidrlist"] = cidrlist
	p.p["esppolicy"] = esppolicy
	p.p["gateway"] = gateway
	p.p["ikepolicy"] = ikepolicy
	p.p["ipsecpsk"] = ipsecpsk
	return p
}

// Creates site to site vpn customer gateway
func (s *VPNService) CreateVpnCustomerGateway(p *CreateVpnCustomerGatewayParams) (*CreateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("createVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnCustomerGatewayResponse
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

		var r CreateVpnCustomerGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVpnCustomerGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Esplifetime int    `json:"esplifetime,omitempty"`
	Ipsecpsk    string `json:"ipsecpsk,omitempty"`
	Ikepolicy   string `json:"ikepolicy,omitempty"`
	Esppolicy   string `json:"esppolicy,omitempty"`
	Removed     string `json:"removed,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Projectid   string `json:"projectid,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Project     string `json:"project,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Account     string `json:"account,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Ikelifetime int    `json:"ikelifetime,omitempty"`
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Dpd         bool   `json:"dpd,omitempty"`
}

type CreateVpnGatewayParams struct {
	p map[string]interface{}
}

func (p *CreateVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *CreateVpnGatewayParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new CreateVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnGatewayParams(vpcid string) *CreateVpnGatewayParams {
	p := &CreateVpnGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["vpcid"] = vpcid
	return p
}

// Creates site to site vpn local gateway
func (s *VPNService) CreateVpnGateway(p *CreateVpnGatewayParams) (*CreateVpnGatewayResponse, error) {
	resp, err := s.cs.newRequest("createVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnGatewayResponse
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

		var r CreateVpnGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVpnGatewayResponse struct {
	JobID     string `json:"jobid,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Id        string `json:"id,omitempty"`
	Removed   string `json:"removed,omitempty"`
	Account   string `json:"account,omitempty"`
	Vpcid     string `json:"vpcid,omitempty"`
	Project   string `json:"project,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
	Publicip  string `json:"publicip,omitempty"`
}

type CreateVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *CreateVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["passive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("passive", vv)
	}
	if v, found := p.p["s2scustomergatewayid"]; found {
		u.Set("s2scustomergatewayid", v.(string))
	}
	if v, found := p.p["s2svpngatewayid"]; found {
		u.Set("s2svpngatewayid", v.(string))
	}
	return u
}

func (p *CreateVpnConnectionParams) SetPassive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["passive"] = v
	return
}

func (p *CreateVpnConnectionParams) SetS2scustomergatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s2scustomergatewayid"] = v
	return
}

func (p *CreateVpnConnectionParams) SetS2svpngatewayid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["s2svpngatewayid"] = v
	return
}

// You should always use this function to get a new CreateVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewCreateVpnConnectionParams(s2scustomergatewayid string, s2svpngatewayid string) *CreateVpnConnectionParams {
	p := &CreateVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["s2scustomergatewayid"] = s2scustomergatewayid
	p.p["s2svpngatewayid"] = s2svpngatewayid
	return p
}

// Create site to site vpn connection
func (s *VPNService) CreateVpnConnection(p *CreateVpnConnectionParams) (*CreateVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("createVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateVpnConnectionResponse
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

		var r CreateVpnConnectionResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type CreateVpnConnectionResponse struct {
	JobID                string `json:"jobid,omitempty"`
	S2scustomergatewayid string `json:"s2scustomergatewayid,omitempty"`
	Created              string `json:"created,omitempty"`
	S2svpngatewayid      string `json:"s2svpngatewayid,omitempty"`
	Ipsecpsk             string `json:"ipsecpsk,omitempty"`
	Account              string `json:"account,omitempty"`
	Esplifetime          int    `json:"esplifetime,omitempty"`
	Removed              string `json:"removed,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Cidrlist             string `json:"cidrlist,omitempty"`
	State                string `json:"state,omitempty"`
	Projectid            string `json:"projectid,omitempty"`
	Dpd                  bool   `json:"dpd,omitempty"`
	Passive              bool   `json:"passive,omitempty"`
	Domain               string `json:"domain,omitempty"`
	Id                   string `json:"id,omitempty"`
	Ikelifetime          int    `json:"ikelifetime,omitempty"`
	Ikepolicy            string `json:"ikepolicy,omitempty"`
	Esppolicy            string `json:"esppolicy,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Project              string `json:"project,omitempty"`
}

type DeleteVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnCustomerGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnCustomerGatewayParams(id string) *DeleteVpnCustomerGatewayParams {
	p := &DeleteVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn customer gateway
func (s *VPNService) DeleteVpnCustomerGateway(p *DeleteVpnCustomerGatewayParams) (*DeleteVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnCustomerGatewayResponse
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

		var r DeleteVpnCustomerGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteVpnCustomerGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type DeleteVpnGatewayParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVpnGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnGatewayParams(id string) *DeleteVpnGatewayParams {
	p := &DeleteVpnGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn gateway
func (s *VPNService) DeleteVpnGateway(p *DeleteVpnGatewayParams) (*DeleteVpnGatewayResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnGatewayResponse
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

		var r DeleteVpnGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteVpnGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type DeleteVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *DeleteVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteVpnConnectionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewDeleteVpnConnectionParams(id string) *DeleteVpnConnectionParams {
	p := &DeleteVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete site to site vpn connection
func (s *VPNService) DeleteVpnConnection(p *DeleteVpnConnectionParams) (*DeleteVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("deleteVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteVpnConnectionResponse
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

		var r DeleteVpnConnectionResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type DeleteVpnConnectionResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type UpdateVpnCustomerGatewayParams struct {
	p map[string]interface{}
}

func (p *UpdateVpnCustomerGatewayParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		u.Set("cidrlist", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["dpd"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("dpd", vv)
	}
	if v, found := p.p["esplifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("esplifetime", vv)
	}
	if v, found := p.p["esppolicy"]; found {
		u.Set("esppolicy", v.(string))
	}
	if v, found := p.p["gateway"]; found {
		u.Set("gateway", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["ikelifetime"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("ikelifetime", vv)
	}
	if v, found := p.p["ikepolicy"]; found {
		u.Set("ikepolicy", v.(string))
	}
	if v, found := p.p["ipsecpsk"]; found {
		u.Set("ipsecpsk", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateVpnCustomerGatewayParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetCidrlist(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetDpd(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["dpd"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetEsplifetime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esplifetime"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetEsppolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["esppolicy"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetGateway(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gateway"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetIkelifetime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikelifetime"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetIkepolicy(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ikepolicy"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetIpsecpsk(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["ipsecpsk"] = v
	return
}

func (p *UpdateVpnCustomerGatewayParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new UpdateVpnCustomerGatewayParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewUpdateVpnCustomerGatewayParams(cidrlist string, esppolicy string, gateway string, id string, ikepolicy string, ipsecpsk string) *UpdateVpnCustomerGatewayParams {
	p := &UpdateVpnCustomerGatewayParams{}
	p.p = make(map[string]interface{})
	p.p["cidrlist"] = cidrlist
	p.p["esppolicy"] = esppolicy
	p.p["gateway"] = gateway
	p.p["id"] = id
	p.p["ikepolicy"] = ikepolicy
	p.p["ipsecpsk"] = ipsecpsk
	return p
}

// Update site to site vpn customer gateway
func (s *VPNService) UpdateVpnCustomerGateway(p *UpdateVpnCustomerGatewayParams) (*UpdateVpnCustomerGatewayResponse, error) {
	resp, err := s.cs.newRequest("updateVpnCustomerGateway", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateVpnCustomerGatewayResponse
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

		var r UpdateVpnCustomerGatewayResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type UpdateVpnCustomerGatewayResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Dpd         bool   `json:"dpd,omitempty"`
	Projectid   string `json:"projectid,omitempty"`
	Name        string `json:"name,omitempty"`
	Project     string `json:"project,omitempty"`
	Removed     string `json:"removed,omitempty"`
	Ikelifetime int    `json:"ikelifetime,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Ipsecpsk    string `json:"ipsecpsk,omitempty"`
	Account     string `json:"account,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Ikepolicy   string `json:"ikepolicy,omitempty"`
	Esppolicy   string `json:"esppolicy,omitempty"`
	Id          string `json:"id,omitempty"`
	Esplifetime int    `json:"esplifetime,omitempty"`
}

type ResetVpnConnectionParams struct {
	p map[string]interface{}
}

func (p *ResetVpnConnectionParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *ResetVpnConnectionParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ResetVpnConnectionParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ResetVpnConnectionParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new ResetVpnConnectionParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewResetVpnConnectionParams(id string) *ResetVpnConnectionParams {
	p := &ResetVpnConnectionParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reset site to site vpn connection
func (s *VPNService) ResetVpnConnection(p *ResetVpnConnectionParams) (*ResetVpnConnectionResponse, error) {
	resp, err := s.cs.newRequest("resetVpnConnection", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ResetVpnConnectionResponse
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

		var r ResetVpnConnectionResponse
		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
		return &r, nil
	}
	return &r, nil
}

type ResetVpnConnectionResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Removed              string `json:"removed,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Created              string `json:"created,omitempty"`
	Project              string `json:"project,omitempty"`
	Domain               string `json:"domain,omitempty"`
	S2scustomergatewayid string `json:"s2scustomergatewayid,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Esppolicy            string `json:"esppolicy,omitempty"`
	Ikelifetime          int    `json:"ikelifetime,omitempty"`
	State                string `json:"state,omitempty"`
	Dpd                  bool   `json:"dpd,omitempty"`
	Ipsecpsk             string `json:"ipsecpsk,omitempty"`
	Cidrlist             string `json:"cidrlist,omitempty"`
	Ikepolicy            string `json:"ikepolicy,omitempty"`
	Esplifetime          int    `json:"esplifetime,omitempty"`
	Projectid            string `json:"projectid,omitempty"`
	Passive              bool   `json:"passive,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Id                   string `json:"id,omitempty"`
	S2svpngatewayid      string `json:"s2svpngatewayid,omitempty"`
	Account              string `json:"account,omitempty"`
}

type ListVpnCustomerGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListVpnCustomerGatewaysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	return u
}

func (p *ListVpnCustomerGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVpnCustomerGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

// You should always use this function to get a new ListVpnCustomerGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnCustomerGatewaysParams() *ListVpnCustomerGatewaysParams {
	p := &ListVpnCustomerGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnCustomerGatewayID(keyword string) (string, error) {
	p := &ListVpnCustomerGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListVpnCustomerGateways(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.VpnCustomerGateways[0].Id, nil
}

// Lists site to site vpn customer gateways
func (s *VPNService) ListVpnCustomerGateways(p *ListVpnCustomerGatewaysParams) (*ListVpnCustomerGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnCustomerGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnCustomerGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVpnCustomerGatewaysResponse struct {
	Count               int                   `json:"count"`
	VpnCustomerGateways []*VpnCustomerGateway `json:"vpncustomergateway"`
}

type VpnCustomerGateway struct {
	Removed     string `json:"removed,omitempty"`
	Project     string `json:"project,omitempty"`
	Dpd         bool   `json:"dpd,omitempty"`
	Ikelifetime int    `json:"ikelifetime,omitempty"`
	Esplifetime int    `json:"esplifetime,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Name        string `json:"name,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Projectid   string `json:"projectid,omitempty"`
	Ipsecpsk    string `json:"ipsecpsk,omitempty"`
	Account     string `json:"account,omitempty"`
	Ikepolicy   string `json:"ikepolicy,omitempty"`
	Gateway     string `json:"gateway,omitempty"`
	Ipaddress   string `json:"ipaddress,omitempty"`
	Id          string `json:"id,omitempty"`
	Esppolicy   string `json:"esppolicy,omitempty"`
}

type ListVpnGatewaysParams struct {
	p map[string]interface{}
}

func (p *ListVpnGatewaysParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListVpnGatewaysParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVpnGatewaysParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVpnGatewaysParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVpnGatewaysParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVpnGatewaysParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVpnGatewaysParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVpnGatewaysParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVpnGatewaysParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVpnGatewaysParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVpnGatewaysParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new ListVpnGatewaysParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnGatewaysParams() *ListVpnGatewaysParams {
	p := &ListVpnGatewaysParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnGatewayID(keyword string) (string, error) {
	p := &ListVpnGatewaysParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListVpnGateways(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.VpnGateways[0].Id, nil
}

// Lists site 2 site vpn gateways
func (s *VPNService) ListVpnGateways(p *ListVpnGatewaysParams) (*ListVpnGatewaysResponse, error) {
	resp, err := s.cs.newRequest("listVpnGateways", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnGatewaysResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVpnGatewaysResponse struct {
	Count       int           `json:"count"`
	VpnGateways []*VpnGateway `json:"vpngateway"`
}

type VpnGateway struct {
	Project   string `json:"project,omitempty"`
	Removed   string `json:"removed,omitempty"`
	Projectid string `json:"projectid,omitempty"`
	Vpcid     string `json:"vpcid,omitempty"`
	Account   string `json:"account,omitempty"`
	Id        string `json:"id,omitempty"`
	Publicip  string `json:"publicip,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Domainid  string `json:"domainid,omitempty"`
}

type ListVpnConnectionsParams struct {
	p map[string]interface{}
}

func (p *ListVpnConnectionsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["isrecursive"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isrecursive", vv)
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["listall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("listall", vv)
	}
	if v, found := p.p["page"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.FormatInt(v.(int64), 10)
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["vpcid"]; found {
		u.Set("vpcid", v.(string))
	}
	return u
}

func (p *ListVpnConnectionsParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListVpnConnectionsParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListVpnConnectionsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListVpnConnectionsParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListVpnConnectionsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListVpnConnectionsParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListVpnConnectionsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListVpnConnectionsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListVpnConnectionsParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListVpnConnectionsParams) SetVpcid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["vpcid"] = v
	return
}

// You should always use this function to get a new ListVpnConnectionsParams instance,
// as then you are sure you have configured all required params
func (s *VPNService) NewListVpnConnectionsParams() *ListVpnConnectionsParams {
	p := &ListVpnConnectionsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *VPNService) GetVpnConnectionID(keyword string) (string, error) {
	p := &ListVpnConnectionsParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListVpnConnections(p)
	if err != nil {
		return "", err
	}
	if l.Count != 1 {
		return "", fmt.Errorf("%d matches found for %s: %+v", l.Count, keyword, l)
	}
	return l.VpnConnections[0].Id, nil
}

// Lists site to site vpn connection gateways
func (s *VPNService) ListVpnConnections(p *ListVpnConnectionsParams) (*ListVpnConnectionsResponse, error) {
	resp, err := s.cs.newRequest("listVpnConnections", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListVpnConnectionsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListVpnConnectionsResponse struct {
	Count          int              `json:"count"`
	VpnConnections []*VpnConnection `json:"vpnconnection"`
}

type VpnConnection struct {
	S2svpngatewayid      string `json:"s2svpngatewayid,omitempty"`
	State                string `json:"state,omitempty"`
	Projectid            string `json:"projectid,omitempty"`
	Project              string `json:"project,omitempty"`
	Id                   string `json:"id,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Domainid             string `json:"domainid,omitempty"`
	Account              string `json:"account,omitempty"`
	Created              string `json:"created,omitempty"`
	Ikelifetime          int    `json:"ikelifetime,omitempty"`
	Dpd                  bool   `json:"dpd,omitempty"`
	Passive              bool   `json:"passive,omitempty"`
	Esplifetime          int    `json:"esplifetime,omitempty"`
	S2scustomergatewayid string `json:"s2scustomergatewayid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Domain               string `json:"domain,omitempty"`
	Ikepolicy            string `json:"ikepolicy,omitempty"`
	Removed              string `json:"removed,omitempty"`
	Esppolicy            string `json:"esppolicy,omitempty"`
	Ipsecpsk             string `json:"ipsecpsk,omitempty"`
	Cidrlist             string `json:"cidrlist,omitempty"`
}
