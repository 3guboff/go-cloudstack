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
	"strings"
)

type CreateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["cidrlist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("cidrlist", vv)
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["openfirewall"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("openfirewall", vv)
	}
	if v, found := p.p["privateport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("privateport", vv)
	}
	if v, found := p.p["protocol"]; found {
		u.Set("protocol", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["publicport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("publicport", vv)
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *CreateLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetCidrlist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["cidrlist"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetOpenfirewall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["openfirewall"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPrivateport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privateport"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetProtocol(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["protocol"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetPublicport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicport"] = v
	return
}

func (p *CreateLoadBalancerRuleParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new CreateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerRuleParams(algorithm string, name string, privateport int, publicport int) *CreateLoadBalancerRuleParams {
	p := &CreateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["name"] = name
	p.p["privateport"] = privateport
	p.p["publicport"] = publicport
	return p
}

// Creates a load balancer rule
func (s *LoadBalancerService) CreateLoadBalancerRule(p *CreateLoadBalancerRuleParams) (*CreateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerRuleResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Zoneid      string `json:"zoneid,omitempty"`
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	State       string `json:"state,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Privateport string `json:"privateport,omitempty"`
	Publicip    string `json:"publicip,omitempty"`
	Account     string `json:"account,omitempty"`
	Algorithm   string `json:"algorithm,omitempty"`
	Projectid   string `json:"projectid,omitempty"`
	Project     string `json:"project,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Tags        []struct {
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Domain       string `json:"domain,omitempty"`
	} `json:"tags,omitempty"`
	Publicport string `json:"publicport,omitempty"`
	Domain     string `json:"domain,omitempty"`
	Domainid   string `json:"domainid,omitempty"`
	Publicipid string `json:"publicipid,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
}

type DeleteLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerRuleParams(id string) *DeleteLoadBalancerRuleParams {
	p := &DeleteLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer rule.
func (s *LoadBalancerService) DeleteLoadBalancerRule(p *DeleteLoadBalancerRuleParams) (*DeleteLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type RemoveFromLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("virtualmachineids", vv)
	}
	return u
}

func (p *RemoveFromLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *RemoveFromLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
	return
}

// You should always use this function to get a new RemoveFromLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromLoadBalancerRuleParams(id string, virtualmachineids []string) *RemoveFromLoadBalancerRuleParams {
	p := &RemoveFromLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineids"] = virtualmachineids
	return p
}

// Removes a virtual machine or a list of virtual machines from a load balancer rule.
func (s *LoadBalancerService) RemoveFromLoadBalancerRule(p *RemoveFromLoadBalancerRuleParams) (*RemoveFromLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type RemoveFromLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type AssignToLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["virtualmachineids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("virtualmachineids", vv)
	}
	return u
}

func (p *AssignToLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AssignToLoadBalancerRuleParams) SetVirtualmachineids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineids"] = v
	return
}

// You should always use this function to get a new AssignToLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToLoadBalancerRuleParams(id string, virtualmachineids []string) *AssignToLoadBalancerRuleParams {
	p := &AssignToLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["virtualmachineids"] = virtualmachineids
	return p
}

// Assigns virtual machine or a list of virtual machines to a load balancer rule.
func (s *LoadBalancerService) AssignToLoadBalancerRule(p *AssignToLoadBalancerRuleParams) (*AssignToLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AssignToLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type CreateLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["methodname"]; found {
		u.Set("methodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["param"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("param[%d].key", i), k)
			u.Set(fmt.Sprintf("param[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *CreateLBStickinessPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetMethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["methodname"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLBStickinessPolicyParams) SetParam(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["param"] = v
	return
}

// You should always use this function to get a new CreateLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBStickinessPolicyParams(lbruleid string, methodname string, name string) *CreateLBStickinessPolicyParams {
	p := &CreateLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	p.p["methodname"] = methodname
	p.p["name"] = name
	return p
}

// Creates a Load Balancer stickiness policy
func (s *LoadBalancerService) CreateLBStickinessPolicy(p *CreateLBStickinessPolicyParams) (*CreateLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBStickinessPolicyResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateLBStickinessPolicyResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Zoneid           string `json:"zoneid,omitempty"`
	Name             string `json:"name,omitempty"`
	Stickinesspolicy []struct {
		Id          string            `json:"id,omitempty"`
		Description string            `json:"description,omitempty"`
		Params      map[string]string `json:"params,omitempty"`
		Name        string            `json:"name,omitempty"`
		Methodname  string            `json:"methodname,omitempty"`
		State       string            `json:"state,omitempty"`
	} `json:"stickinesspolicy,omitempty"`
	Description string `json:"description,omitempty"`
	State       string `json:"state,omitempty"`
	Lbruleid    string `json:"lbruleid,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Account     string `json:"account,omitempty"`
}

type DeleteLBStickinessPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBStickinessPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBStickinessPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLBStickinessPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBStickinessPolicyParams(id string) *DeleteLBStickinessPolicyParams {
	p := &DeleteLBStickinessPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a LB stickiness policy.
func (s *LoadBalancerService) DeleteLBStickinessPolicy(p *DeleteLBStickinessPolicyParams) (*DeleteLBStickinessPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBStickinessPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBStickinessPolicyResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteLBStickinessPolicyResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRulesParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["publicipid"]; found {
		u.Set("publicipid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetPublicipid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["publicipid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

func (p *ListLoadBalancerRulesParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRulesParams() *ListLoadBalancerRulesParams {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleID(name string) (string, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListLoadBalancerRules(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancerRules {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByName(name string) (*LoadBalancerRule, int, error) {
	id, err := s.GetLoadBalancerRuleID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetLoadBalancerRuleByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerRuleByID(id string) (*LoadBalancerRule, int, error) {
	p := &ListLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListLoadBalancerRules(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListLoadBalancerRules(p *ListLoadBalancerRulesParams) (*ListLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLoadBalancerRulesResponse struct {
	Count             int                 `json:"count"`
	LoadBalancerRules []*LoadBalancerRule `json:"loadbalancerrule"`
}

type LoadBalancerRule struct {
	Projectid   string `json:"projectid,omitempty"`
	Project     string `json:"project,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Account     string `json:"account,omitempty"`
	Algorithm   string `json:"algorithm,omitempty"`
	Publicport  string `json:"publicport,omitempty"`
	Name        string `json:"name,omitempty"`
	Publicipid  string `json:"publicipid,omitempty"`
	Publicip    string `json:"publicip,omitempty"`
	Description string `json:"description,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Id          string `json:"id,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Tags        []struct {
		Customer     string `json:"customer,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Project      string `json:"project,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Key          string `json:"key,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Account      string `json:"account,omitempty"`
	} `json:"tags,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Privateport string `json:"privateport,omitempty"`
	Zoneid      string `json:"zoneid,omitempty"`
	State       string `json:"state,omitempty"`
}

type ListLBStickinessPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBStickinessPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLBStickinessPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLBStickinessPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLBStickinessPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBStickinessPoliciesParams(lbruleid string) *ListLBStickinessPoliciesParams {
	p := &ListLBStickinessPoliciesParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Lists LBStickiness policies.
func (s *LoadBalancerService) ListLBStickinessPolicies(p *ListLBStickinessPoliciesParams) (*ListLBStickinessPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBStickinessPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBStickinessPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLBStickinessPoliciesResponse struct {
	Count                int                   `json:"count"`
	LBStickinessPolicies []*LBStickinessPolicy `json:"lbstickinesspolicy"`
}

type LBStickinessPolicy struct {
	Name             string `json:"name,omitempty"`
	State            string `json:"state,omitempty"`
	Stickinesspolicy []struct {
		Params      map[string]string `json:"params,omitempty"`
		Methodname  string            `json:"methodname,omitempty"`
		Id          string            `json:"id,omitempty"`
		State       string            `json:"state,omitempty"`
		Name        string            `json:"name,omitempty"`
		Description string            `json:"description,omitempty"`
	} `json:"stickinesspolicy,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Zoneid      string `json:"zoneid,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Description string `json:"description,omitempty"`
	Account     string `json:"account,omitempty"`
	Lbruleid    string `json:"lbruleid,omitempty"`
}

type ListLBHealthCheckPoliciesParams struct {
	p map[string]interface{}
}

func (p *ListLBHealthCheckPoliciesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLBHealthCheckPoliciesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLBHealthCheckPoliciesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLBHealthCheckPoliciesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLBHealthCheckPoliciesParams(lbruleid string) *ListLBHealthCheckPoliciesParams {
	p := &ListLBHealthCheckPoliciesParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Lists load balancer HealthCheck policies.
func (s *LoadBalancerService) ListLBHealthCheckPolicies(p *ListLBHealthCheckPoliciesParams) (*ListLBHealthCheckPoliciesResponse, error) {
	resp, err := s.cs.newRequest("listLBHealthCheckPolicies", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLBHealthCheckPoliciesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLBHealthCheckPoliciesResponse struct {
	Count                 int                    `json:"count"`
	LBHealthCheckPolicies []*LBHealthCheckPolicy `json:"lbhealthcheckpolicy"`
}

type LBHealthCheckPolicy struct {
	Domainid          string `json:"domainid,omitempty"`
	Account           string `json:"account,omitempty"`
	Healthcheckpolicy []struct {
		Unhealthcheckthresshold int    `json:"unhealthcheckthresshold,omitempty"`
		State                   string `json:"state,omitempty"`
		Id                      string `json:"id,omitempty"`
		Healthcheckthresshold   int    `json:"healthcheckthresshold,omitempty"`
		Description             string `json:"description,omitempty"`
		Healthcheckinterval     int    `json:"healthcheckinterval,omitempty"`
		Pingpath                string `json:"pingpath,omitempty"`
		Responsetime            int    `json:"responsetime,omitempty"`
	} `json:"healthcheckpolicy,omitempty"`
	Lbruleid string `json:"lbruleid,omitempty"`
	Zoneid   string `json:"zoneid,omitempty"`
	Domain   string `json:"domain,omitempty"`
}

type CreateLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *CreateLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["healthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("healthythreshold", vv)
	}
	if v, found := p.p["intervaltime"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("intervaltime", vv)
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	if v, found := p.p["pingpath"]; found {
		u.Set("pingpath", v.(string))
	}
	if v, found := p.p["responsetimeout"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("responsetimeout", vv)
	}
	if v, found := p.p["unhealthythreshold"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("unhealthythreshold", vv)
	}
	return u
}

func (p *CreateLBHealthCheckPolicyParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetHealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["healthythreshold"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetIntervaltime(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["intervaltime"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetPingpath(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pingpath"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetResponsetimeout(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["responsetimeout"] = v
	return
}

func (p *CreateLBHealthCheckPolicyParams) SetUnhealthythreshold(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["unhealthythreshold"] = v
	return
}

// You should always use this function to get a new CreateLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLBHealthCheckPolicyParams(lbruleid string) *CreateLBHealthCheckPolicyParams {
	p := &CreateLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Creates a Load Balancer healthcheck policy
func (s *LoadBalancerService) CreateLBHealthCheckPolicy(p *CreateLBHealthCheckPolicyParams) (*CreateLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("createLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLBHealthCheckPolicyResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateLBHealthCheckPolicyResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Lbruleid          string `json:"lbruleid,omitempty"`
	Healthcheckpolicy []struct {
		Pingpath                string `json:"pingpath,omitempty"`
		Healthcheckinterval     int    `json:"healthcheckinterval,omitempty"`
		Id                      string `json:"id,omitempty"`
		Description             string `json:"description,omitempty"`
		Responsetime            int    `json:"responsetime,omitempty"`
		Unhealthcheckthresshold int    `json:"unhealthcheckthresshold,omitempty"`
		State                   string `json:"state,omitempty"`
		Healthcheckthresshold   int    `json:"healthcheckthresshold,omitempty"`
	} `json:"healthcheckpolicy,omitempty"`
	Domainid string `json:"domainid,omitempty"`
	Account  string `json:"account,omitempty"`
	Zoneid   string `json:"zoneid,omitempty"`
	Domain   string `json:"domain,omitempty"`
}

type DeleteLBHealthCheckPolicyParams struct {
	p map[string]interface{}
}

func (p *DeleteLBHealthCheckPolicyParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLBHealthCheckPolicyParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLBHealthCheckPolicyParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLBHealthCheckPolicyParams(id string) *DeleteLBHealthCheckPolicyParams {
	p := &DeleteLBHealthCheckPolicyParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer HealthCheck policy.
func (s *LoadBalancerService) DeleteLBHealthCheckPolicy(p *DeleteLBHealthCheckPolicyParams) (*DeleteLBHealthCheckPolicyResponse, error) {
	resp, err := s.cs.newRequest("deleteLBHealthCheckPolicy", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLBHealthCheckPolicyResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteLBHealthCheckPolicyResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListLoadBalancerRuleInstancesParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancerRuleInstancesParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["applied"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("applied", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	return u
}

func (p *ListLoadBalancerRuleInstancesParams) SetApplied(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["applied"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancerRuleInstancesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

// You should always use this function to get a new ListLoadBalancerRuleInstancesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancerRuleInstancesParams(id string) *ListLoadBalancerRuleInstancesParams {
	p := &ListLoadBalancerRuleInstancesParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// List all virtual machine instances that are assigned to a load balancer rule.
func (s *LoadBalancerService) ListLoadBalancerRuleInstances(p *ListLoadBalancerRuleInstancesParams) (*ListLoadBalancerRuleInstancesResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancerRuleInstances", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancerRuleInstancesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLoadBalancerRuleInstancesResponse struct {
	Count                     int                         `json:"count"`
	LoadBalancerRuleInstances []*LoadBalancerRuleInstance `json:"loadbalancerruleinstance"`
}

type LoadBalancerRuleInstance struct {
	Affinitygroup []struct {
		Name              string   `json:"name,omitempty"`
		Id                string   `json:"id,omitempty"`
		Description       string   `json:"description,omitempty"`
		Account           string   `json:"account,omitempty"`
		Domainid          string   `json:"domainid,omitempty"`
		Domain            string   `json:"domain,omitempty"`
		VirtualmachineIds []string `json:"virtualmachineIds,omitempty"`
		Type              string   `json:"type,omitempty"`
	} `json:"affinitygroup,omitempty"`
	Diskiowrite     int    `json:"diskiowrite,omitempty"`
	Rootdevicetype  string `json:"rootdevicetype,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
	Isoid           string `json:"isoid,omitempty"`
	Id              string `json:"id,omitempty"`
	Diskioread      int    `json:"diskioread,omitempty"`
	Diskkbsread     int    `json:"diskkbsread,omitempty"`
	Haenable        bool   `json:"haenable,omitempty"`
	Keypair         string `json:"keypair,omitempty"`
	Groupid         string `json:"groupid,omitempty"`
	Isoname         string `json:"isoname,omitempty"`
	Passwordenabled bool   `json:"passwordenabled,omitempty"`
	Hostid          string `json:"hostid,omitempty"`
	Templatename    string `json:"templatename,omitempty"`
	Project         string `json:"project,omitempty"`
	Isodisplaytext  string `json:"isodisplaytext,omitempty"`
	Displayname     string `json:"displayname,omitempty"`
	Securitygroup   []struct {
		Domainid string `json:"domainid,omitempty"`
		Name     string `json:"name,omitempty"`
		Id       string `json:"id,omitempty"`
		Tags     []struct {
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Project      string `json:"project,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Project     string `json:"project,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Description string `json:"description,omitempty"`
		Egressrule  []struct {
			Ruleid            string `json:"ruleid,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Startport         int    `json:"startport,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
		} `json:"egressrule,omitempty"`
		Ingressrule []struct {
			Startport         int    `json:"startport,omitempty"`
			Cidr              string `json:"cidr,omitempty"`
			Ruleid            string `json:"ruleid,omitempty"`
			Securitygroupname string `json:"securitygroupname,omitempty"`
			Protocol          string `json:"protocol,omitempty"`
			Account           string `json:"account,omitempty"`
			Endport           int    `json:"endport,omitempty"`
			Icmpcode          int    `json:"icmpcode,omitempty"`
			Icmptype          int    `json:"icmptype,omitempty"`
		} `json:"ingressrule,omitempty"`
		Account string `json:"account,omitempty"`
	} `json:"securitygroup,omitempty"`
	Domainid              string `json:"domainid,omitempty"`
	Isdynamicallyscalable bool   `json:"isdynamicallyscalable,omitempty"`
	Tags                  []struct {
		Domainid     string `json:"domainid,omitempty"`
		Value        string `json:"value,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Account      string `json:"account,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Customer     string `json:"customer,omitempty"`
	} `json:"tags,omitempty"`
	Serviceofferingid   string            `json:"serviceofferingid,omitempty"`
	Servicestate        string            `json:"servicestate,omitempty"`
	Password            string            `json:"password,omitempty"`
	Publicipid          string            `json:"publicipid,omitempty"`
	Cpuused             string            `json:"cpuused,omitempty"`
	Hostname            string            `json:"hostname,omitempty"`
	Serviceofferingname string            `json:"serviceofferingname,omitempty"`
	Cpunumber           int               `json:"cpunumber,omitempty"`
	Networkkbswrite     int               `json:"networkkbswrite,omitempty"`
	Zonename            string            `json:"zonename,omitempty"`
	Details             map[string]string `json:"details,omitempty"`
	Rootdeviceid        int               `json:"rootdeviceid,omitempty"`
	Group               string            `json:"group,omitempty"`
	Diskkbswrite        int               `json:"diskkbswrite,omitempty"`
	Templateid          string            `json:"templateid,omitempty"`
	Displayvm           bool              `json:"displayvm,omitempty"`
	Account             string            `json:"account,omitempty"`
	Templatedisplaytext string            `json:"templatedisplaytext,omitempty"`
	Domain              string            `json:"domain,omitempty"`
	Forvirtualnetwork   bool              `json:"forvirtualnetwork,omitempty"`
	Name                string            `json:"name,omitempty"`
	State               string            `json:"state,omitempty"`
	Nic                 []struct {
		Netmask      string   `json:"netmask,omitempty"`
		Ip6cidr      string   `json:"ip6cidr,omitempty"`
		Macaddress   string   `json:"macaddress,omitempty"`
		Broadcasturi string   `json:"broadcasturi,omitempty"`
		Isdefault    bool     `json:"isdefault,omitempty"`
		Ipaddress    string   `json:"ipaddress,omitempty"`
		Id           string   `json:"id,omitempty"`
		Type         string   `json:"type,omitempty"`
		Networkid    string   `json:"networkid,omitempty"`
		Networkname  string   `json:"networkname,omitempty"`
		Ip6gateway   string   `json:"ip6gateway,omitempty"`
		Gateway      string   `json:"gateway,omitempty"`
		Isolationuri string   `json:"isolationuri,omitempty"`
		Traffictype  string   `json:"traffictype,omitempty"`
		Secondaryip  []string `json:"secondaryip,omitempty"`
		Ip6address   string   `json:"ip6address,omitempty"`
	} `json:"nic,omitempty"`
	Guestosid      string `json:"guestosid,omitempty"`
	Hypervisor     string `json:"hypervisor,omitempty"`
	Cpuspeed       int    `json:"cpuspeed,omitempty"`
	Zoneid         string `json:"zoneid,omitempty"`
	Memory         int    `json:"memory,omitempty"`
	Networkkbsread int    `json:"networkkbsread,omitempty"`
	Instancename   string `json:"instancename,omitempty"`
	Publicip       string `json:"publicip,omitempty"`
	Created        string `json:"created,omitempty"`
}

type UpdateLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	return u
}

func (p *UpdateLoadBalancerRuleParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *UpdateLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

// You should always use this function to get a new UpdateLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateLoadBalancerRuleParams(id string) *UpdateLoadBalancerRuleParams {
	p := &UpdateLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Updates load balancer
func (s *LoadBalancerService) UpdateLoadBalancerRule(p *UpdateLoadBalancerRuleParams) (*UpdateLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateLoadBalancerRuleResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UpdateLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Privateport string `json:"privateport,omitempty"`
	Cidrlist    string `json:"cidrlist,omitempty"`
	Account     string `json:"account,omitempty"`
	Projectid   string `json:"projectid,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Publicip    string `json:"publicip,omitempty"`
	Algorithm   string `json:"algorithm,omitempty"`
	State       string `json:"state,omitempty"`
	Id          string `json:"id,omitempty"`
	Networkid   string `json:"networkid,omitempty"`
	Project     string `json:"project,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Domainid    string `json:"domainid,omitempty"`
	Tags        []struct {
		Domain       string `json:"domain,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Key          string `json:"key,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Value        string `json:"value,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Customer     string `json:"customer,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
	} `json:"tags,omitempty"`
	Description string `json:"description,omitempty"`
	Publicipid  string `json:"publicipid,omitempty"`
	Zoneid      string `json:"zoneid,omitempty"`
	Name        string `json:"name,omitempty"`
	Publicport  string `json:"publicport,omitempty"`
}

type UploadSslCertParams struct {
	p map[string]interface{}
}

func (p *UploadSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certchain"]; found {
		u.Set("certchain", v.(string))
	}
	if v, found := p.p["certificate"]; found {
		u.Set("certificate", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["privatekey"]; found {
		u.Set("privatekey", v.(string))
	}
	return u
}

func (p *UploadSslCertParams) SetCertchain(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certchain"] = v
	return
}

func (p *UploadSslCertParams) SetCertificate(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certificate"] = v
	return
}

func (p *UploadSslCertParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *UploadSslCertParams) SetPrivatekey(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["privatekey"] = v
	return
}

// You should always use this function to get a new UploadSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUploadSslCertParams(certificate string, privatekey string) *UploadSslCertParams {
	p := &UploadSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["certificate"] = certificate
	p.p["privatekey"] = privatekey
	return p
}

// Upload a certificate to cloudstack
func (s *LoadBalancerService) UploadSslCert(p *UploadSslCertParams) (*UploadSslCertResponse, error) {
	resp, err := s.cs.newRequest("uploadSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UploadSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type UploadSslCertResponse struct {
	Fingerprint          string   `json:"fingerprint,omitempty"`
	Certchain            string   `json:"certchain,omitempty"`
	Account              string   `json:"account,omitempty"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist,omitempty"`
	Id                   string   `json:"id,omitempty"`
	Certificate          string   `json:"certificate,omitempty"`
	Privatekey           string   `json:"privatekey,omitempty"`
}

type DeleteSslCertParams struct {
	p map[string]interface{}
}

func (p *DeleteSslCertParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteSslCertParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteSslCertParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteSslCertParams(id string) *DeleteSslCertParams {
	p := &DeleteSslCertParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Delete a certificate to cloudstack
func (s *LoadBalancerService) DeleteSslCert(p *DeleteSslCertParams) (*DeleteSslCertResponse, error) {
	resp, err := s.cs.newRequest("deleteSslCert", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteSslCertResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type DeleteSslCertResponse struct {
	Success     string `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ListSslCertsParams struct {
	p map[string]interface{}
}

func (p *ListSslCertsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["accountid"]; found {
		u.Set("accountid", v.(string))
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *ListSslCertsParams) SetAccountid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["accountid"] = v
	return
}

func (p *ListSslCertsParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
	return
}

func (p *ListSslCertsParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

// You should always use this function to get a new ListSslCertsParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListSslCertsParams() *ListSslCertsParams {
	p := &ListSslCertsParams{}
	p.p = make(map[string]interface{})
	return p
}

// Lists SSL certificates
func (s *LoadBalancerService) ListSslCerts(p *ListSslCertsParams) (*ListSslCertsResponse, error) {
	resp, err := s.cs.newRequest("listSslCerts", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSslCertsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSslCertsResponse struct {
	Count    int        `json:"count"`
	SslCerts []*SslCert `json:"sslcert"`
}

type SslCert struct {
	Fingerprint          string   `json:"fingerprint,omitempty"`
	Loadbalancerrulelist []string `json:"loadbalancerrulelist,omitempty"`
	Certchain            string   `json:"certchain,omitempty"`
	Privatekey           string   `json:"privatekey,omitempty"`
	Certificate          string   `json:"certificate,omitempty"`
	Id                   string   `json:"id,omitempty"`
	Account              string   `json:"account,omitempty"`
}

type AssignCertToLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AssignCertToLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["certid"]; found {
		u.Set("certid", v.(string))
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *AssignCertToLoadBalancerParams) SetCertid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["certid"] = v
	return
}

func (p *AssignCertToLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

// You should always use this function to get a new AssignCertToLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignCertToLoadBalancerParams(certid string, lbruleid string) *AssignCertToLoadBalancerParams {
	p := &AssignCertToLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["certid"] = certid
	p.p["lbruleid"] = lbruleid
	return p
}

// Assigns a certificate to a Load Balancer Rule
func (s *LoadBalancerService) AssignCertToLoadBalancer(p *AssignCertToLoadBalancerParams) (*AssignCertToLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("assignCertToLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignCertToLoadBalancerResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AssignCertToLoadBalancerResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type RemoveCertFromLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *RemoveCertFromLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbruleid"]; found {
		u.Set("lbruleid", v.(string))
	}
	return u
}

func (p *RemoveCertFromLoadBalancerParams) SetLbruleid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbruleid"] = v
	return
}

// You should always use this function to get a new RemoveCertFromLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveCertFromLoadBalancerParams(lbruleid string) *RemoveCertFromLoadBalancerParams {
	p := &RemoveCertFromLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbruleid"] = lbruleid
	return p
}

// Removes a certificate from a Load Balancer Rule
func (s *LoadBalancerService) RemoveCertFromLoadBalancer(p *RemoveCertFromLoadBalancerParams) (*RemoveCertFromLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("removeCertFromLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveCertFromLoadBalancerResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type RemoveCertFromLoadBalancerResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type AddF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AddF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddF5LoadBalancerParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddF5LoadBalancerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAddF5LoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddF5LoadBalancerParams {
	p := &AddF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a F5 BigIP load balancer device
func (s *LoadBalancerService) AddF5LoadBalancer(p *AddF5LoadBalancerParams) (*AddF5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddF5LoadBalancerResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AddF5LoadBalancerResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Lbdevicecapacity  int    `json:"lbdevicecapacity,omitempty"`
	Lbdeviceid        string `json:"lbdeviceid,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Lbdevicededicated bool   `json:"lbdevicededicated,omitempty"`
	Lbdevicename      string `json:"lbdevicename,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
	Lbdevicestate     string `json:"lbdevicestate,omitempty"`
}

type ConfigureF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *ConfigureF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdevicecapacity"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *ConfigureF5LoadBalancerParams) SetLbdevicecapacity(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicecapacity"] = v
	return
}

func (p *ConfigureF5LoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new ConfigureF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewConfigureF5LoadBalancerParams(lbdeviceid string) *ConfigureF5LoadBalancerParams {
	p := &ConfigureF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a F5 load balancer device
func (s *LoadBalancerService) ConfigureF5LoadBalancer(p *ConfigureF5LoadBalancerParams) (*ConfigureF5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigureF5LoadBalancerResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type ConfigureF5LoadBalancerResponse struct {
	JobID             string `json:"jobid,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
	Lbdeviceid        string `json:"lbdeviceid,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Lbdevicecapacity  int    `json:"lbdevicecapacity,omitempty"`
	Lbdevicestate     string `json:"lbdevicestate,omitempty"`
	Lbdevicename      string `json:"lbdevicename,omitempty"`
	Lbdevicededicated bool   `json:"lbdevicededicated,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Ipaddress         string `json:"ipaddress,omitempty"`
}

type DeleteF5LoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteF5LoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *DeleteF5LoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new DeleteF5LoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteF5LoadBalancerParams(lbdeviceid string) *DeleteF5LoadBalancerParams {
	p := &DeleteF5LoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

//  delete a F5 load balancer device
func (s *LoadBalancerService) DeleteF5LoadBalancer(p *DeleteF5LoadBalancerParams) (*DeleteF5LoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteF5LoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteF5LoadBalancerResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteF5LoadBalancerResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type ListF5LoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListF5LoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListF5LoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListF5LoadBalancersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListF5LoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListF5LoadBalancersParams() *ListF5LoadBalancersParams {
	p := &ListF5LoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists F5 load balancer devices
func (s *LoadBalancerService) ListF5LoadBalancers(p *ListF5LoadBalancersParams) (*ListF5LoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listF5LoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListF5LoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListF5LoadBalancersResponse struct {
	Count           int               `json:"count"`
	F5LoadBalancers []*F5LoadBalancer `json:"f5loadbalancer"`
}

type F5LoadBalancer struct {
	Ipaddress         string `json:"ipaddress,omitempty"`
	Lbdevicename      string `json:"lbdevicename,omitempty"`
	Provider          string `json:"provider,omitempty"`
	Publicinterface   string `json:"publicinterface,omitempty"`
	Lbdevicecapacity  int    `json:"lbdevicecapacity,omitempty"`
	Lbdevicestate     string `json:"lbdevicestate,omitempty"`
	Lbdeviceid        string `json:"lbdeviceid,omitempty"`
	Lbdevicededicated bool   `json:"lbdevicededicated,omitempty"`
	Privateinterface  string `json:"privateinterface,omitempty"`
	Physicalnetworkid string `json:"physicalnetworkid,omitempty"`
}

type AddNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *AddNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("gslbprovider", vv)
	}
	if v, found := p.p["gslbproviderprivateip"]; found {
		u.Set("gslbproviderprivateip", v.(string))
	}
	if v, found := p.p["gslbproviderpublicip"]; found {
		u.Set("gslbproviderpublicip", v.(string))
	}
	if v, found := p.p["isexclusivegslbprovider"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("isexclusivegslbprovider", vv)
	}
	if v, found := p.p["networkdevicetype"]; found {
		u.Set("networkdevicetype", v.(string))
	}
	if v, found := p.p["password"]; found {
		u.Set("password", v.(string))
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	if v, found := p.p["url"]; found {
		u.Set("url", v.(string))
	}
	if v, found := p.p["username"]; found {
		u.Set("username", v.(string))
	}
	return u
}

func (p *AddNetscalerLoadBalancerParams) SetGslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbprovider"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderprivateip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderprivateip"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetGslbproviderpublicip(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbproviderpublicip"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetIsexclusivegslbprovider(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isexclusivegslbprovider"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetNetworkdevicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkdevicetype"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetPassword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["password"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetUrl(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["url"] = v
	return
}

func (p *AddNetscalerLoadBalancerParams) SetUsername(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["username"] = v
	return
}

// You should always use this function to get a new AddNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAddNetscalerLoadBalancerParams(networkdevicetype string, password string, physicalnetworkid string, url string, username string) *AddNetscalerLoadBalancerParams {
	p := &AddNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["networkdevicetype"] = networkdevicetype
	p.p["password"] = password
	p.p["physicalnetworkid"] = physicalnetworkid
	p.p["url"] = url
	p.p["username"] = username
	return p
}

// Adds a netscaler load balancer device
func (s *LoadBalancerService) AddNetscalerLoadBalancer(p *AddNetscalerLoadBalancerParams) (*AddNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("addNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AddNetscalerLoadBalancerResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AddNetscalerLoadBalancerResponse struct {
	JobID                   string   `json:"jobid,omitempty"`
	Podids                  []string `json:"podids,omitempty"`
	Physicalnetworkid       string   `json:"physicalnetworkid,omitempty"`
	Lbdeviceid              string   `json:"lbdeviceid,omitempty"`
	Lbdevicededicated       bool     `json:"lbdevicededicated,omitempty"`
	Publicinterface         string   `json:"publicinterface,omitempty"`
	Provider                string   `json:"provider,omitempty"`
	Ipaddress               string   `json:"ipaddress,omitempty"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip,omitempty"`
	Lbdevicename            string   `json:"lbdevicename,omitempty"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider,omitempty"`
	Gslbprovider            bool     `json:"gslbprovider,omitempty"`
	Lbdevicecapacity        int      `json:"lbdevicecapacity,omitempty"`
	Lbdevicestate           string   `json:"lbdevicestate,omitempty"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip,omitempty"`
	Privateinterface        string   `json:"privateinterface,omitempty"`
}

type DeleteNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	return u
}

func (p *DeleteNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

// You should always use this function to get a new DeleteNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteNetscalerLoadBalancerParams(lbdeviceid string) *DeleteNetscalerLoadBalancerParams {
	p := &DeleteNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

//  delete a netscaler load balancer device
func (s *LoadBalancerService) DeleteNetscalerLoadBalancer(p *DeleteNetscalerLoadBalancerParams) (*DeleteNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteNetscalerLoadBalancerResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteNetscalerLoadBalancerResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type ConfigureNetscalerLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *ConfigureNetscalerLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["inline"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("inline", vv)
	}
	if v, found := p.p["lbdevicecapacity"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("lbdevicecapacity", vv)
	}
	if v, found := p.p["lbdevicededicated"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("lbdevicededicated", vv)
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["podids"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("podids", vv)
	}
	return u
}

func (p *ConfigureNetscalerLoadBalancerParams) SetInline(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["inline"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicecapacity(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicecapacity"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdevicededicated(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdevicededicated"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ConfigureNetscalerLoadBalancerParams) SetPodids(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podids"] = v
	return
}

// You should always use this function to get a new ConfigureNetscalerLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewConfigureNetscalerLoadBalancerParams(lbdeviceid string) *ConfigureNetscalerLoadBalancerParams {
	p := &ConfigureNetscalerLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["lbdeviceid"] = lbdeviceid
	return p
}

// configures a netscaler load balancer device
func (s *LoadBalancerService) ConfigureNetscalerLoadBalancer(p *ConfigureNetscalerLoadBalancerParams) (*ConfigureNetscalerLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("configureNetscalerLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ConfigureNetscalerLoadBalancerResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type ConfigureNetscalerLoadBalancerResponse struct {
	JobID                   string   `json:"jobid,omitempty"`
	Lbdevicededicated       bool     `json:"lbdevicededicated,omitempty"`
	Ipaddress               string   `json:"ipaddress,omitempty"`
	Publicinterface         string   `json:"publicinterface,omitempty"`
	Podids                  []string `json:"podids,omitempty"`
	Privateinterface        string   `json:"privateinterface,omitempty"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip,omitempty"`
	Gslbprovider            bool     `json:"gslbprovider,omitempty"`
	Lbdeviceid              string   `json:"lbdeviceid,omitempty"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider,omitempty"`
	Lbdevicecapacity        int      `json:"lbdevicecapacity,omitempty"`
	Provider                string   `json:"provider,omitempty"`
	Lbdevicename            string   `json:"lbdevicename,omitempty"`
	Lbdevicestate           string   `json:"lbdevicestate,omitempty"`
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip,omitempty"`
	Physicalnetworkid       string   `json:"physicalnetworkid,omitempty"`
}

type ListNetscalerLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListNetscalerLoadBalancersParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["lbdeviceid"]; found {
		u.Set("lbdeviceid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["physicalnetworkid"]; found {
		u.Set("physicalnetworkid", v.(string))
	}
	return u
}

func (p *ListNetscalerLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetLbdeviceid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["lbdeviceid"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListNetscalerLoadBalancersParams) SetPhysicalnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["physicalnetworkid"] = v
	return
}

// You should always use this function to get a new ListNetscalerLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListNetscalerLoadBalancersParams() *ListNetscalerLoadBalancersParams {
	p := &ListNetscalerLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// lists netscaler load balancer devices
func (s *LoadBalancerService) ListNetscalerLoadBalancers(p *ListNetscalerLoadBalancersParams) (*ListNetscalerLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listNetscalerLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListNetscalerLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListNetscalerLoadBalancersResponse struct {
	Count                  int                      `json:"count"`
	NetscalerLoadBalancers []*NetscalerLoadBalancer `json:"netscalerloadbalancer"`
}

type NetscalerLoadBalancer struct {
	Gslbproviderprivateip   string   `json:"gslbproviderprivateip,omitempty"`
	Gslbproviderpublicip    string   `json:"gslbproviderpublicip,omitempty"`
	Privateinterface        string   `json:"privateinterface,omitempty"`
	Lbdevicestate           string   `json:"lbdevicestate,omitempty"`
	Provider                string   `json:"provider,omitempty"`
	Podids                  []string `json:"podids,omitempty"`
	Gslbprovider            bool     `json:"gslbprovider,omitempty"`
	Lbdevicecapacity        int      `json:"lbdevicecapacity,omitempty"`
	Isexclusivegslbprovider bool     `json:"isexclusivegslbprovider,omitempty"`
	Lbdevicename            string   `json:"lbdevicename,omitempty"`
	Lbdevicededicated       bool     `json:"lbdevicededicated,omitempty"`
	Ipaddress               string   `json:"ipaddress,omitempty"`
	Publicinterface         string   `json:"publicinterface,omitempty"`
	Lbdeviceid              string   `json:"lbdeviceid,omitempty"`
	Physicalnetworkid       string   `json:"physicalnetworkid,omitempty"`
}

type CreateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *CreateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["account"]; found {
		u.Set("account", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["domainid"]; found {
		u.Set("domainid", v.(string))
	}
	if v, found := p.p["gslbdomainname"]; found {
		u.Set("gslbdomainname", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbservicetype"]; found {
		u.Set("gslbservicetype", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	return u
}

func (p *CreateGlobalLoadBalancerRuleParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbdomainname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbdomainname"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbservicetype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbservicetype"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateGlobalLoadBalancerRuleParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
	return
}

// You should always use this function to get a new CreateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateGlobalLoadBalancerRuleParams(gslbdomainname string, gslbservicetype string, name string, regionid int) *CreateGlobalLoadBalancerRuleParams {
	p := &CreateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["gslbdomainname"] = gslbdomainname
	p.p["gslbservicetype"] = gslbservicetype
	p.p["name"] = name
	p.p["regionid"] = regionid
	return p
}

// Creates a global load balancer rule
func (s *LoadBalancerService) CreateGlobalLoadBalancerRule(p *CreateGlobalLoadBalancerRuleParams) (*CreateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("createGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateGlobalLoadBalancerRuleResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateGlobalLoadBalancerRuleResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Project          string `json:"project,omitempty"`
	Domainid         string `json:"domainid,omitempty"`
	Gslbdomainname   string `json:"gslbdomainname,omitempty"`
	Regionid         int    `json:"regionid,omitempty"`
	Gslbservicetype  string `json:"gslbservicetype,omitempty"`
	Account          string `json:"account,omitempty"`
	Loadbalancerrule []struct {
		Algorithm   string `json:"algorithm,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Description string `json:"description,omitempty"`
		Zoneid      string `json:"zoneid,omitempty"`
		Cidrlist    string `json:"cidrlist,omitempty"`
		Protocol    string `json:"protocol,omitempty"`
		Name        string `json:"name,omitempty"`
		Publicport  string `json:"publicport,omitempty"`
		Project     string `json:"project,omitempty"`
		Account     string `json:"account,omitempty"`
		Networkid   string `json:"networkid,omitempty"`
		Publicip    string `json:"publicip,omitempty"`
		Privateport string `json:"privateport,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Publicipid  string `json:"publicipid,omitempty"`
		Tags        []struct {
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Key          string `json:"key,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		State string `json:"state,omitempty"`
		Id    string `json:"id,omitempty"`
	} `json:"loadbalancerrule,omitempty"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname,omitempty"`
	Projectid                   string `json:"projectid,omitempty"`
	Gslblbmethod                string `json:"gslblbmethod,omitempty"`
	Name                        string `json:"name,omitempty"`
	Id                          string `json:"id,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Description                 string `json:"description,omitempty"`
}

type DeleteGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *DeleteGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteGlobalLoadBalancerRuleParams(id string) *DeleteGlobalLoadBalancerRuleParams {
	p := &DeleteGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a global load balancer rule.
func (s *LoadBalancerService) DeleteGlobalLoadBalancerRule(p *DeleteGlobalLoadBalancerRuleParams) (*DeleteGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("deleteGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteGlobalLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteGlobalLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}

type UpdateGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *UpdateGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["gslblbmethod"]; found {
		u.Set("gslblbmethod", v.(string))
	}
	if v, found := p.p["gslbstickysessionmethodname"]; found {
		u.Set("gslbstickysessionmethodname", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslblbmethod(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbmethod"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetGslbstickysessionmethodname(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslbstickysessionmethodname"] = v
	return
}

func (p *UpdateGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new UpdateGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewUpdateGlobalLoadBalancerRuleParams(id string) *UpdateGlobalLoadBalancerRuleParams {
	p := &UpdateGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// update global load balancer rules.
func (s *LoadBalancerService) UpdateGlobalLoadBalancerRule(p *UpdateGlobalLoadBalancerRuleParams) (*UpdateGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("updateGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r UpdateGlobalLoadBalancerRuleResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type UpdateGlobalLoadBalancerRuleResponse struct {
	JobID                       string `json:"jobid,omitempty"`
	Name                        string `json:"name,omitempty"`
	Description                 string `json:"description,omitempty"`
	Domain                      string `json:"domain,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Gslbdomainname              string `json:"gslbdomainname,omitempty"`
	Id                          string `json:"id,omitempty"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname,omitempty"`
	Regionid                    int    `json:"regionid,omitempty"`
	Project                     string `json:"project,omitempty"`
	Loadbalancerrule            []struct {
		Protocol    string `json:"protocol,omitempty"`
		Name        string `json:"name,omitempty"`
		Networkid   string `json:"networkid,omitempty"`
		Description string `json:"description,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Publicipid  string `json:"publicipid,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Tags        []struct {
			Customer     string `json:"customer,omitempty"`
			Project      string `json:"project,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Account      string `json:"account,omitempty"`
			Value        string `json:"value,omitempty"`
			Projectid    string `json:"projectid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"tags,omitempty"`
		Account     string `json:"account,omitempty"`
		Algorithm   string `json:"algorithm,omitempty"`
		Zoneid      string `json:"zoneid,omitempty"`
		Publicip    string `json:"publicip,omitempty"`
		Project     string `json:"project,omitempty"`
		State       string `json:"state,omitempty"`
		Privateport string `json:"privateport,omitempty"`
		Publicport  string `json:"publicport,omitempty"`
		Id          string `json:"id,omitempty"`
		Cidrlist    string `json:"cidrlist,omitempty"`
	} `json:"loadbalancerrule,omitempty"`
	Gslblbmethod    string `json:"gslblbmethod,omitempty"`
	Account         string `json:"account,omitempty"`
	Gslbservicetype string `json:"gslbservicetype,omitempty"`
	Projectid       string `json:"projectid,omitempty"`
}

type ListGlobalLoadBalancerRulesParams struct {
	p map[string]interface{}
}

func (p *ListGlobalLoadBalancerRulesParams) toURLValues() url.Values {
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
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["regionid"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("regionid", vv)
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListGlobalLoadBalancerRulesParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetRegionid(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["regionid"] = v
	return
}

func (p *ListGlobalLoadBalancerRulesParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListGlobalLoadBalancerRulesParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListGlobalLoadBalancerRulesParams() *ListGlobalLoadBalancerRulesParams {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleID(keyword string) (string, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["keyword"] = keyword

	l, err := s.ListGlobalLoadBalancerRules(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", keyword, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.GlobalLoadBalancerRules {
			if v.Name == keyword {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", keyword, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByName(name string) (*GlobalLoadBalancerRule, int, error) {
	id, err := s.GetGlobalLoadBalancerRuleID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetGlobalLoadBalancerRuleByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetGlobalLoadBalancerRuleByID(id string) (*GlobalLoadBalancerRule, int, error) {
	p := &ListGlobalLoadBalancerRulesParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListGlobalLoadBalancerRules(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.GlobalLoadBalancerRules[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for GlobalLoadBalancerRule UUID: %s!", id)
}

// Lists load balancer rules.
func (s *LoadBalancerService) ListGlobalLoadBalancerRules(p *ListGlobalLoadBalancerRulesParams) (*ListGlobalLoadBalancerRulesResponse, error) {
	resp, err := s.cs.newRequest("listGlobalLoadBalancerRules", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListGlobalLoadBalancerRulesResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListGlobalLoadBalancerRulesResponse struct {
	Count                   int                       `json:"count"`
	GlobalLoadBalancerRules []*GlobalLoadBalancerRule `json:"globalloadbalancerrule"`
}

type GlobalLoadBalancerRule struct {
	Description      string `json:"description,omitempty"`
	Gslbservicetype  string `json:"gslbservicetype,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Name             string `json:"name,omitempty"`
	Projectid        string `json:"projectid,omitempty"`
	Loadbalancerrule []struct {
		Publicip    string `json:"publicip,omitempty"`
		Domainid    string `json:"domainid,omitempty"`
		State       string `json:"state,omitempty"`
		Account     string `json:"account,omitempty"`
		Protocol    string `json:"protocol,omitempty"`
		Zoneid      string `json:"zoneid,omitempty"`
		Id          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Projectid   string `json:"projectid,omitempty"`
		Algorithm   string `json:"algorithm,omitempty"`
		Description string `json:"description,omitempty"`
		Privateport string `json:"privateport,omitempty"`
		Cidrlist    string `json:"cidrlist,omitempty"`
		Publicipid  string `json:"publicipid,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Networkid   string `json:"networkid,omitempty"`
		Tags        []struct {
			Projectid    string `json:"projectid,omitempty"`
			Resourcetype string `json:"resourcetype,omitempty"`
			Project      string `json:"project,omitempty"`
			Value        string `json:"value,omitempty"`
			Resourceid   string `json:"resourceid,omitempty"`
			Customer     string `json:"customer,omitempty"`
			Domainid     string `json:"domainid,omitempty"`
			Domain       string `json:"domain,omitempty"`
			Key          string `json:"key,omitempty"`
			Account      string `json:"account,omitempty"`
		} `json:"tags,omitempty"`
		Publicport string `json:"publicport,omitempty"`
		Project    string `json:"project,omitempty"`
	} `json:"loadbalancerrule,omitempty"`
	Project                     string `json:"project,omitempty"`
	Gslbdomainname              string `json:"gslbdomainname,omitempty"`
	Gslblbmethod                string `json:"gslblbmethod,omitempty"`
	Id                          string `json:"id,omitempty"`
	Gslbstickysessionmethodname string `json:"gslbstickysessionmethodname,omitempty"`
	Domainid                    string `json:"domainid,omitempty"`
	Account                     string `json:"account,omitempty"`
	Regionid                    int    `json:"regionid,omitempty"`
}

type AssignToGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *AssignToGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["gslblbruleweightsmap"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].key", i), k)
			u.Set(fmt.Sprintf("gslblbruleweightsmap[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetGslblbruleweightsmap(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["gslblbruleweightsmap"] = v
	return
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *AssignToGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
	return
}

// You should always use this function to get a new AssignToGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewAssignToGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *AssignToGlobalLoadBalancerRuleParams {
	p := &AssignToGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Assign load balancer rule or list of load balancer rules to a global load balancer rules.
func (s *LoadBalancerService) AssignToGlobalLoadBalancerRule(p *AssignToGlobalLoadBalancerRuleParams) (*AssignToGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("assignToGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r AssignToGlobalLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type AssignToGlobalLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type RemoveFromGlobalLoadBalancerRuleParams struct {
	p map[string]interface{}
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["loadbalancerrulelist"]; found {
		vv := strings.Join(v.([]string), ", ")
		u.Set("loadbalancerrulelist", vv)
	}
	return u
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *RemoveFromGlobalLoadBalancerRuleParams) SetLoadbalancerrulelist(v []string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["loadbalancerrulelist"] = v
	return
}

// You should always use this function to get a new RemoveFromGlobalLoadBalancerRuleParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewRemoveFromGlobalLoadBalancerRuleParams(id string, loadbalancerrulelist []string) *RemoveFromGlobalLoadBalancerRuleParams {
	p := &RemoveFromGlobalLoadBalancerRuleParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["loadbalancerrulelist"] = loadbalancerrulelist
	return p
}

// Removes a load balancer rule association with global load balancer rule
func (s *LoadBalancerService) RemoveFromGlobalLoadBalancerRule(p *RemoveFromGlobalLoadBalancerRuleParams) (*RemoveFromGlobalLoadBalancerRuleResponse, error) {
	resp, err := s.cs.newRequest("removeFromGlobalLoadBalancerRule", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RemoveFromGlobalLoadBalancerRuleResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type RemoveFromGlobalLoadBalancerRuleResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
	Success     bool   `json:"success,omitempty"`
}

type CreateLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *CreateLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["algorithm"]; found {
		u.Set("algorithm", v.(string))
	}
	if v, found := p.p["description"]; found {
		u.Set("description", v.(string))
	}
	if v, found := p.p["instanceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("instanceport", vv)
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["sourceport"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("sourceport", vv)
	}
	return u
}

func (p *CreateLoadBalancerParams) SetAlgorithm(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["algorithm"] = v
	return
}

func (p *CreateLoadBalancerParams) SetDescription(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["description"] = v
	return
}

func (p *CreateLoadBalancerParams) SetInstanceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["instanceport"] = v
	return
}

func (p *CreateLoadBalancerParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *CreateLoadBalancerParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *CreateLoadBalancerParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
	return
}

func (p *CreateLoadBalancerParams) SetSourceport(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceport"] = v
	return
}

// You should always use this function to get a new CreateLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewCreateLoadBalancerParams(algorithm string, instanceport int, name string, networkid string, scheme string, sourceipaddressnetworkid string, sourceport int) *CreateLoadBalancerParams {
	p := &CreateLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["algorithm"] = algorithm
	p.p["instanceport"] = instanceport
	p.p["name"] = name
	p.p["networkid"] = networkid
	p.p["scheme"] = scheme
	p.p["sourceipaddressnetworkid"] = sourceipaddressnetworkid
	p.p["sourceport"] = sourceport
	return p
}

// Creates a Load Balancer
func (s *LoadBalancerService) CreateLoadBalancer(p *CreateLoadBalancerParams) (*CreateLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("createLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r CreateLoadBalancerResponse
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

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type CreateLoadBalancerResponse struct {
	JobID            string `json:"jobid,omitempty"`
	Loadbalancerrule []struct {
		Instanceport int    `json:"instanceport,omitempty"`
		Sourceport   int    `json:"sourceport,omitempty"`
		State        string `json:"state,omitempty"`
	} `json:"loadbalancerrule,omitempty"`
	Projectid                string `json:"projectid,omitempty"`
	Description              string `json:"description,omitempty"`
	Domainid                 string `json:"domainid,omitempty"`
	Sourceipaddressnetworkid string `json:"sourceipaddressnetworkid,omitempty"`
	Algorithm                string `json:"algorithm,omitempty"`
	Id                       string `json:"id,omitempty"`
	Sourceipaddress          string `json:"sourceipaddress,omitempty"`
	Account                  string `json:"account,omitempty"`
	Project                  string `json:"project,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	Name                     string `json:"name,omitempty"`
	Tags                     []struct {
		Customer     string `json:"customer,omitempty"`
		Value        string `json:"value,omitempty"`
		Key          string `json:"key,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Account      string `json:"account,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
	} `json:"tags,omitempty"`
	Networkid            string `json:"networkid,omitempty"`
	Loadbalancerinstance []struct {
		State     string `json:"state,omitempty"`
		Ipaddress string `json:"ipaddress,omitempty"`
		Id        string `json:"id,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"loadbalancerinstance,omitempty"`
}

type ListLoadBalancersParams struct {
	p map[string]interface{}
}

func (p *ListLoadBalancersParams) toURLValues() url.Values {
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
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["networkid"]; found {
		u.Set("networkid", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["projectid"]; found {
		u.Set("projectid", v.(string))
	}
	if v, found := p.p["scheme"]; found {
		u.Set("scheme", v.(string))
	}
	if v, found := p.p["sourceipaddress"]; found {
		u.Set("sourceipaddress", v.(string))
	}
	if v, found := p.p["sourceipaddressnetworkid"]; found {
		u.Set("sourceipaddressnetworkid", v.(string))
	}
	if v, found := p.p["tags"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("tags[%d].key", i), k)
			u.Set(fmt.Sprintf("tags[%d].value", i), vv)
			i++
		}
	}
	return u
}

func (p *ListLoadBalancersParams) SetAccount(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["account"] = v
	return
}

func (p *ListLoadBalancersParams) SetDomainid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["domainid"] = v
	return
}

func (p *ListLoadBalancersParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListLoadBalancersParams) SetIsrecursive(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["isrecursive"] = v
	return
}

func (p *ListLoadBalancersParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListLoadBalancersParams) SetListall(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["listall"] = v
	return
}

func (p *ListLoadBalancersParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListLoadBalancersParams) SetNetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["networkid"] = v
	return
}

func (p *ListLoadBalancersParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListLoadBalancersParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListLoadBalancersParams) SetProjectid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["projectid"] = v
	return
}

func (p *ListLoadBalancersParams) SetScheme(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["scheme"] = v
	return
}

func (p *ListLoadBalancersParams) SetSourceipaddress(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddress"] = v
	return
}

func (p *ListLoadBalancersParams) SetSourceipaddressnetworkid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["sourceipaddressnetworkid"] = v
	return
}

func (p *ListLoadBalancersParams) SetTags(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["tags"] = v
	return
}

// You should always use this function to get a new ListLoadBalancersParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewListLoadBalancersParams() *ListLoadBalancersParams {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerID(name string) (string, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListLoadBalancers(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.LoadBalancers {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByName(name string) (*LoadBalancer, int, error) {
	id, err := s.GetLoadBalancerID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetLoadBalancerByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *LoadBalancerService) GetLoadBalancerByID(id string) (*LoadBalancer, int, error) {
	p := &ListLoadBalancersParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListLoadBalancers(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.LoadBalancers[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for LoadBalancer UUID: %s!", id)
}

// Lists Load Balancers
func (s *LoadBalancerService) ListLoadBalancers(p *ListLoadBalancersParams) (*ListLoadBalancersResponse, error) {
	resp, err := s.cs.newRequest("listLoadBalancers", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListLoadBalancersResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListLoadBalancersResponse struct {
	Count         int             `json:"count"`
	LoadBalancers []*LoadBalancer `json:"loadbalancer"`
}

type LoadBalancer struct {
	Project              string `json:"project,omitempty"`
	Domain               string `json:"domain,omitempty"`
	Loadbalancerinstance []struct {
		State     string `json:"state,omitempty"`
		Ipaddress string `json:"ipaddress,omitempty"`
		Name      string `json:"name,omitempty"`
		Id        string `json:"id,omitempty"`
	} `json:"loadbalancerinstance,omitempty"`
	Domainid                 string `json:"domainid,omitempty"`
	Networkid                string `json:"networkid,omitempty"`
	Name                     string `json:"name,omitempty"`
	Projectid                string `json:"projectid,omitempty"`
	Sourceipaddressnetworkid string `json:"sourceipaddressnetworkid,omitempty"`
	Loadbalancerrule         []struct {
		Sourceport   int    `json:"sourceport,omitempty"`
		Instanceport int    `json:"instanceport,omitempty"`
		State        string `json:"state,omitempty"`
	} `json:"loadbalancerrule,omitempty"`
	Id              string `json:"id,omitempty"`
	Description     string `json:"description,omitempty"`
	Sourceipaddress string `json:"sourceipaddress,omitempty"`
	Tags            []struct {
		Customer     string `json:"customer,omitempty"`
		Key          string `json:"key,omitempty"`
		Projectid    string `json:"projectid,omitempty"`
		Domain       string `json:"domain,omitempty"`
		Domainid     string `json:"domainid,omitempty"`
		Resourceid   string `json:"resourceid,omitempty"`
		Account      string `json:"account,omitempty"`
		Project      string `json:"project,omitempty"`
		Resourcetype string `json:"resourcetype,omitempty"`
		Value        string `json:"value,omitempty"`
	} `json:"tags,omitempty"`
	Account   string `json:"account,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
}

type DeleteLoadBalancerParams struct {
	p map[string]interface{}
}

func (p *DeleteLoadBalancerParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DeleteLoadBalancerParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DeleteLoadBalancerParams instance,
// as then you are sure you have configured all required params
func (s *LoadBalancerService) NewDeleteLoadBalancerParams(id string) *DeleteLoadBalancerParams {
	p := &DeleteLoadBalancerParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Deletes a load balancer
func (s *LoadBalancerService) DeleteLoadBalancer(p *DeleteLoadBalancerParams) (*DeleteLoadBalancerResponse, error) {
	resp, err := s.cs.newRequest("deleteLoadBalancer", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DeleteLoadBalancerResponse
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

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DeleteLoadBalancerResponse struct {
	JobID       string `json:"jobid,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Displaytext string `json:"displaytext,omitempty"`
}
