// Package autoscaler provides access to the Google Compute Engine Autoscaler API.
//
// See http://developers.google.com/compute/docs/autoscaler
//
// Usage example:
//
//   import "google.golang.org/api/autoscaler/v1beta2"
//   ...
//   autoscalerService, err := autoscaler.New(oauthHttpClient)
package autoscaler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "autoscaler:v1beta2"
const apiName = "autoscaler"
const apiVersion = "v1beta2"
const basePath = "https://www.googleapis.com/autoscaler/v1beta2/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Compute Engine resources
	ComputeScope = "https://www.googleapis.com/auth/compute"

	// View your Google Compute Engine resources
	ComputeReadonlyScope = "https://www.googleapis.com/auth/compute.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Autoscalers = NewAutoscalersService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	s.Zones = NewZonesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Autoscalers *AutoscalersService

	ZoneOperations *ZoneOperationsService

	Zones *ZonesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAutoscalersService(s *Service) *AutoscalersService {
	rs := &AutoscalersService{s: s}
	return rs
}

type AutoscalersService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

func NewZonesService(s *Service) *ZonesService {
	rs := &ZonesService{s: s}
	return rs
}

type ZonesService struct {
	s *Service
}

type Autoscaler struct {
	// AutoscalingPolicy: Configuration parameters for autoscaling
	// algorithm.
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource provided
	// by the client.
	Description string `json:"description,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// Name: Name of the Autoscaler resource. Must be unique per project and
	// zone.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output Only] A self-link to the Autoscaler configuration
	// resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Target: URL to the entity which will be autoscaled. Currently the
	// only supported value is ReplicaPool?s URL. Note: it is illegal to
	// specify multiple Autoscalers for the same target.
	Target string `json:"target,omitempty"`
}

type AutoscalerListResponse struct {
	// Items: Autoscaler resources.
	Items []*Autoscaler `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output only] A token used to continue a truncated
	// list request.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type AutoscalingPolicy struct {
	// CoolDownPeriodSec: The number of seconds that the Autoscaler should
	// wait between two succeeding changes to the number of virtual
	// machines. You should define an interval that is at least as long as
	// the initialization time of a virtual machine and the time it may take
	// for replica pool to create the virtual machine. The default is 60
	// seconds.
	CoolDownPeriodSec int64 `json:"coolDownPeriodSec,omitempty"`

	// CpuUtilization: Exactly one utilization policy should be provided.
	// Configuration parameters of CPU based autoscaling policy.
	CpuUtilization *AutoscalingPolicyCpuUtilization `json:"cpuUtilization,omitempty"`

	// CustomMetricUtilizations: Configuration parameters of autoscaling
	// based on custom metric.
	CustomMetricUtilizations []*AutoscalingPolicyCustomMetricUtilization `json:"customMetricUtilizations,omitempty"`

	// LoadBalancingUtilization: Configuration parameters of autoscaling
	// based on load balancer.
	LoadBalancingUtilization *AutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty"`

	// MaxNumReplicas: The maximum number of replicas that the Autoscaler
	// can scale up to.
	MaxNumReplicas int64 `json:"maxNumReplicas,omitempty"`

	// MinNumReplicas: The minimum number of replicas that the Autoscaler
	// can scale down to.
	MinNumReplicas int64 `json:"minNumReplicas,omitempty"`
}

type AutoscalingPolicyCpuUtilization struct {
	// UtilizationTarget: The target utilization that the Autoscaler should
	// maintain. It is represented as a fraction of used cores. For example:
	// 6 cores used in 8-core VM are represented here as 0.75. Must be a
	// float value between (0, 1]. If not defined, the default is 0.8.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`
}

type AutoscalingPolicyCustomMetricUtilization struct {
	// Metric: Identifier of the metric. It should be a Cloud Monitoring
	// metric. The metric can not have negative values. The metric should be
	// an utilization metric (increasing number of VMs handling requests x
	// times should reduce average value of the metric roughly x times). For
	// example you could use:
	// compute.googleapis.com/instance/network/received_bytes_count.
	Metric string `json:"metric,omitempty"`

	// UtilizationTarget: Target value of the metric which Autoscaler should
	// maintain. Must be a positive value.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`

	// UtilizationTargetType: Defines type in which utilization_target is
	// expressed.
	UtilizationTargetType string `json:"utilizationTargetType,omitempty"`
}

type AutoscalingPolicyLoadBalancingUtilization struct {
	// UtilizationTarget: Fraction of backend capacity utilization (set in
	// HTTP load balancing configuration) that Autoscaler should maintain.
	// Must be a positive float value. If not defined, the default is 0.8.
	// For example if your maxRatePerInstance capacity (in HTTP Load
	// Balancing configuration) is set at 10 and you would like to keep
	// number of instances such that each instance receives 7 QPS on
	// average, set this to 0.7.
	UtilizationTarget float64 `json:"utilizationTarget,omitempty"`
}

type DeprecationStatus struct {
	Deleted string `json:"deleted,omitempty"`

	Deprecated string `json:"deprecated,omitempty"`

	Obsolete string `json:"obsolete,omitempty"`

	Replacement string `json:"replacement,omitempty"`

	State string `json:"state,omitempty"`
}

type Operation struct {
	ClientOperationId string `json:"clientOperationId,omitempty"`

	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	Error *OperationError `json:"error,omitempty"`

	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	Id uint64 `json:"id,omitempty,string"`

	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always compute#Operation
	// for Operation resources.
	Kind string `json:"kind,omitempty"`

	Name string `json:"name,omitempty"`

	OperationType string `json:"operationType,omitempty"`

	Progress int64 `json:"progress,omitempty"`

	Region string `json:"region,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`

	StartTime string `json:"startTime,omitempty"`

	Status string `json:"status,omitempty"`

	StatusMessage string `json:"statusMessage,omitempty"`

	TargetId uint64 `json:"targetId,omitempty,string"`

	TargetLink string `json:"targetLink,omitempty"`

	User string `json:"user,omitempty"`

	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	Zone string `json:"zone,omitempty"`
}

type OperationError struct {
	Errors []*OperationErrorErrors `json:"errors,omitempty"`
}

type OperationErrorErrors struct {
	Code string `json:"code,omitempty"`

	Location string `json:"location,omitempty"`

	Message string `json:"message,omitempty"`
}

type OperationWarnings struct {
	Code string `json:"code,omitempty"`

	Data []*OperationWarningsData `json:"data,omitempty"`

	Message string `json:"message,omitempty"`
}

type OperationWarningsData struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type OperationList struct {
	Id string `json:"id,omitempty"`

	Items []*Operation `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always compute#operations for
	// Operations resource.
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`
}

type Zone struct {
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	Deprecated *DeprecationStatus `json:"deprecated,omitempty"`

	Description string `json:"description,omitempty"`

	Id uint64 `json:"id,omitempty,string"`

	// Kind: [Output Only] Type of the resource. Always kind#zone for zones.
	Kind string `json:"kind,omitempty"`

	MaintenanceWindows []*ZoneMaintenanceWindows `json:"maintenanceWindows,omitempty"`

	Name string `json:"name,omitempty"`

	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] Server defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	Status string `json:"status,omitempty"`
}

type ZoneMaintenanceWindows struct {
	BeginTime string `json:"beginTime,omitempty"`

	Description string `json:"description,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	Name string `json:"name,omitempty"`
}

type ZoneList struct {
	Id string `json:"id,omitempty"`

	Items []*Zone `json:"items,omitempty"`

	// Kind: Type of resource.
	Kind string `json:"kind,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: Server defined URL for this resource (output only).
	SelfLink string `json:"selfLink,omitempty"`
}

// method id "autoscaler.autoscalers.delete":

type AutoscalersDeleteCall struct {
	s          *Service
	project    string
	zone       string
	autoscaler string
	opt_       map[string]interface{}
}

// Delete: Deletes the specified Autoscaler resource.
func (r *AutoscalersService) Delete(project string, zone string, autoscaler string) *AutoscalersDeleteCall {
	c := &AutoscalersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.autoscaler = autoscaler
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersDeleteCall) Fields(s ...googleapi.Field) *AutoscalersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersDeleteCall) IfNoneMatch(entityTag string) *AutoscalersDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.delete" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersDeleteCall) Do() (*Operation, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersDeleteCall) DoHeader() (ret *Operation, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers/{autoscaler}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Deletes the specified Autoscaler resource.",
	//   "httpMethod": "DELETE",
	//   "id": "autoscaler.autoscalers.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.get":

type AutoscalersGetCall struct {
	s          *Service
	project    string
	zone       string
	autoscaler string
	opt_       map[string]interface{}
}

// Get: Gets the specified Autoscaler resource.
func (r *AutoscalersService) Get(project string, zone string, autoscaler string) *AutoscalersGetCall {
	c := &AutoscalersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.autoscaler = autoscaler
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersGetCall) Fields(s ...googleapi.Field) *AutoscalersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersGetCall) IfNoneMatch(entityTag string) *AutoscalersGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.get" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersGetCall) Do() (*Autoscaler, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersGetCall) DoHeader() (ret *Autoscaler, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers/{autoscaler}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Gets the specified Autoscaler resource.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.autoscalers.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "response": {
	//     "$ref": "Autoscaler"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.insert":

type AutoscalersInsertCall struct {
	s          *Service
	project    string
	zone       string
	autoscaler *Autoscaler
	opt_       map[string]interface{}
}

// Insert: Adds new Autoscaler resource.
func (r *AutoscalersService) Insert(project string, zone string, autoscaler *Autoscaler) *AutoscalersInsertCall {
	c := &AutoscalersInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.autoscaler = autoscaler
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersInsertCall) Fields(s ...googleapi.Field) *AutoscalersInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersInsertCall) IfNoneMatch(entityTag string) *AutoscalersInsertCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.insert" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersInsertCall) Do() (*Operation, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.insert" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersInsertCall) DoHeader() (ret *Operation, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.autoscaler)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Adds new Autoscaler resource.",
	//   "httpMethod": "POST",
	//   "id": "autoscaler.autoscalers.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.list":

type AutoscalersListCall struct {
	s       *Service
	project string
	zone    string
	opt_    map[string]interface{}
}

// List: Lists all Autoscaler resources in this zone.
func (r *AutoscalersService) List(project string, zone string) *AutoscalersListCall {
	c := &AutoscalersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	return c
}

// Filter sets the optional parameter "filter":
func (c *AutoscalersListCall) Filter(filter string) *AutoscalersListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *AutoscalersListCall) MaxResults(maxResults int64) *AutoscalersListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *AutoscalersListCall) PageToken(pageToken string) *AutoscalersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersListCall) Fields(s ...googleapi.Field) *AutoscalersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersListCall) IfNoneMatch(entityTag string) *AutoscalersListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.list" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersListCall) Do() (*AutoscalerListResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersListCall) DoHeader() (ret *AutoscalerListResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Lists all Autoscaler resources in this zone.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.autoscalers.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers",
	//   "response": {
	//     "$ref": "AutoscalerListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.patch":

type AutoscalersPatchCall struct {
	s           *Service
	project     string
	zone        string
	autoscaler  string
	autoscaler2 *Autoscaler
	opt_        map[string]interface{}
}

// Patch: Update the entire content of the Autoscaler resource. This
// method supports patch semantics.
func (r *AutoscalersService) Patch(project string, zone string, autoscaler string, autoscaler2 *Autoscaler) *AutoscalersPatchCall {
	c := &AutoscalersPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.autoscaler = autoscaler
	c.autoscaler2 = autoscaler2
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersPatchCall) Fields(s ...googleapi.Field) *AutoscalersPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersPatchCall) IfNoneMatch(entityTag string) *AutoscalersPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.patch" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersPatchCall) Do() (*Operation, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersPatchCall) DoHeader() (ret *Operation, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.autoscaler2)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers/{autoscaler}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Update the entire content of the Autoscaler resource. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "autoscaler.autoscalers.patch",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.autoscalers.update":

type AutoscalersUpdateCall struct {
	s           *Service
	project     string
	zone        string
	autoscaler  string
	autoscaler2 *Autoscaler
	opt_        map[string]interface{}
}

// Update: Update the entire content of the Autoscaler resource.
func (r *AutoscalersService) Update(project string, zone string, autoscaler string, autoscaler2 *Autoscaler) *AutoscalersUpdateCall {
	c := &AutoscalersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.autoscaler = autoscaler
	c.autoscaler2 = autoscaler2
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AutoscalersUpdateCall) Fields(s ...googleapi.Field) *AutoscalersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *AutoscalersUpdateCall) IfNoneMatch(entityTag string) *AutoscalersUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.autoscalers.update" call.
// Exactly one of the return values is non-nil.
func (c *AutoscalersUpdateCall) Do() (*Operation, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.autoscalers.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *AutoscalersUpdateCall) DoHeader() (ret *Operation, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.autoscaler2)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/zones/{zone}/autoscalers/{autoscaler}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":    c.project,
		"zone":       c.zone,
		"autoscaler": c.autoscaler,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Update the entire content of the Autoscaler resource.",
	//   "httpMethod": "PUT",
	//   "id": "autoscaler.autoscalers.update",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "autoscaler"
	//   ],
	//   "parameters": {
	//     "autoscaler": {
	//       "description": "Name of the Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Zone name of Autoscaler resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/zones/{zone}/autoscalers/{autoscaler}",
	//   "request": {
	//     "$ref": "Autoscaler"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.delete":

type ZoneOperationsDeleteCall struct {
	s         *Service
	project   string
	zone      string
	operation string
	opt_      map[string]interface{}
}

// Delete: Deletes the specified zone-specific operation resource.
func (r *ZoneOperationsService) Delete(project string, zone string, operation string) *ZoneOperationsDeleteCall {
	c := &ZoneOperationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsDeleteCall) Fields(s ...googleapi.Field) *ZoneOperationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ZoneOperationsDeleteCall) IfNoneMatch(entityTag string) *ZoneOperationsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.zoneOperations.delete" call.
func (c *ZoneOperationsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "autoscaler.zoneOperations.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ZoneOperationsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/operations/{operation}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return res.Header, err
	}
	return res.Header, nil
	// {
	//   "description": "Deletes the specified zone-specific operation resource.",
	//   "httpMethod": "DELETE",
	//   "id": "autoscaler.zoneOperations.delete",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.get":

type ZoneOperationsGetCall struct {
	s         *Service
	project   string
	zone      string
	operation string
	opt_      map[string]interface{}
}

// Get: Retrieves the specified zone-specific operation resource.
func (r *ZoneOperationsService) Get(project string, zone string, operation string) *ZoneOperationsGetCall {
	c := &ZoneOperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsGetCall) Fields(s ...googleapi.Field) *ZoneOperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ZoneOperationsGetCall) IfNoneMatch(entityTag string) *ZoneOperationsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.zoneOperations.get" call.
// Exactly one of the return values is non-nil.
func (c *ZoneOperationsGetCall) Do() (*Operation, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.zoneOperations.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ZoneOperationsGetCall) DoHeader() (ret *Operation, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/operations/{operation}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the specified zone-specific operation resource.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.zoneOperations.list":

type ZoneOperationsListCall struct {
	s       *Service
	project string
	zone    string
	opt_    map[string]interface{}
}

// List: Retrieves the list of operation resources contained within the
// specified zone.
func (r *ZoneOperationsService) List(project string, zone string) *ZoneOperationsListCall {
	c := &ZoneOperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.zone = zone
	return c
}

// Filter sets the optional parameter "filter":
func (c *ZoneOperationsListCall) Filter(filter string) *ZoneOperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *ZoneOperationsListCall) MaxResults(maxResults int64) *ZoneOperationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ZoneOperationsListCall) PageToken(pageToken string) *ZoneOperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsListCall) Fields(s ...googleapi.Field) *ZoneOperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ZoneOperationsListCall) IfNoneMatch(entityTag string) *ZoneOperationsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.zoneOperations.list" call.
// Exactly one of the return values is non-nil.
func (c *ZoneOperationsListCall) Do() (*OperationList, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.zoneOperations.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ZoneOperationsListCall) DoHeader() (ret *OperationList, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/operations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Retrieves the list of operation resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zoneOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "location": "path",
	//       "pattern": "[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}

// method id "autoscaler.zones.list":

type ZonesListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List:
func (r *ZonesService) List(project string) *ZonesListCall {
	c := &ZonesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Filter sets the optional parameter "filter":
func (c *ZonesListCall) Filter(filter string) *ZonesListCall {
	c.opt_["filter"] = filter
	return c
}

// MaxResults sets the optional parameter "maxResults":
func (c *ZonesListCall) MaxResults(maxResults int64) *ZonesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken":
func (c *ZonesListCall) PageToken(pageToken string) *ZonesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZonesListCall) Fields(s ...googleapi.Field) *ZonesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *ZonesListCall) IfNoneMatch(entityTag string) *ZonesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "autoscaler.zones.list" call.
// Exactly one of the return values is non-nil.
func (c *ZonesListCall) Do() (*ZoneList, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "autoscaler.zones.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *ZonesListCall) DoHeader() (ret *ZoneList, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if v, ok := c.opt_["ifNoneMatch"]; ok {
		req.Header.Set("If-None-Match", fmt.Sprintf("%v", v))
	}
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "",
	//   "httpMethod": "GET",
	//   "id": "autoscaler.zones.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones",
	//   "response": {
	//     "$ref": "ZoneList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/compute",
	//     "https://www.googleapis.com/auth/compute.readonly"
	//   ]
	// }

}
