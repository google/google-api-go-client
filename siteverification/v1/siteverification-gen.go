// Package siteverification provides access to the Google Site Verification API.
//
// See https://developers.google.com/site-verification/
//
// Usage example:
//
//   import "google.golang.org/api/siteverification/v1"
//   ...
//   siteverificationService, err := siteverification.New(oauthHttpClient)
package siteverification

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

const apiId = "siteVerification:v1"
const apiName = "siteVerification"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/siteVerification/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage the list of sites and domains you control
	SiteverificationScope = "https://www.googleapis.com/auth/siteverification"

	// Manage your new site verifications with Google
	SiteverificationVerifyOnlyScope = "https://www.googleapis.com/auth/siteverification.verify_only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.WebResource = NewWebResourceService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	WebResource *WebResourceService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewWebResourceService(s *Service) *WebResourceService {
	rs := &WebResourceService{s: s}
	return rs
}

type WebResourceService struct {
	s *Service
}

type SiteVerificationWebResourceGettokenRequest struct {
	// Site: The site for which a verification token will be generated.
	Site *SiteVerificationWebResourceGettokenRequestSite `json:"site,omitempty"`

	// VerificationMethod: The verification method that will be used to
	// verify this site. For sites, 'FILE' or 'META' methods may be used.
	// For domains, only 'DNS' may be used.
	VerificationMethod string `json:"verificationMethod,omitempty"`
}

type SiteVerificationWebResourceGettokenRequestSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The type of resource to be verified. Can be SITE or INET_DOMAIN
	// (domain name).
	Type string `json:"type,omitempty"`
}

type SiteVerificationWebResourceGettokenResponse struct {
	// Method: The verification method to use in conjunction with this
	// token. For FILE, the token should be placed in the top-level
	// directory of the site, stored inside a file of the same name. For
	// META, the token should be placed in the HEAD tag of the default page
	// that is loaded for the site. For DNS, the token should be placed in a
	// TXT record of the domain.
	Method string `json:"method,omitempty"`

	// Token: The verification token. The token must be placed appropriately
	// in order for verification to succeed.
	Token string `json:"token,omitempty"`
}

type SiteVerificationWebResourceListResponse struct {
	// Items: The list of sites that are owned by the authenticated user.
	Items []*SiteVerificationWebResourceResource `json:"items,omitempty"`
}

type SiteVerificationWebResourceResource struct {
	// Id: The string used to identify this site. This value should be used
	// in the "id" portion of the REST URL for the Get, Update, and Delete
	// operations.
	Id string `json:"id,omitempty"`

	// Owners: The email addresses of all verified owners.
	Owners []string `json:"owners,omitempty"`

	// Site: The address and type of a site that is verified or will be
	// verified.
	Site *SiteVerificationWebResourceResourceSite `json:"site,omitempty"`
}

type SiteVerificationWebResourceResourceSite struct {
	// Identifier: The site identifier. If the type is set to SITE, the
	// identifier is a URL. If the type is set to INET_DOMAIN, the site
	// identifier is a domain name.
	Identifier string `json:"identifier,omitempty"`

	// Type: The site type. Can be SITE or INET_DOMAIN (domain name).
	Type string `json:"type,omitempty"`
}

// method id "siteVerification.webResource.delete":

type WebResourceDeleteCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Delete: Relinquish ownership of a website or domain.
func (r *WebResourceService) Delete(id string) *WebResourceDeleteCall {
	c := &WebResourceDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceDeleteCall) Fields(s ...googleapi.Field) *WebResourceDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceDeleteCall) IfNoneMatch(ifNoneMatch string) *WebResourceDeleteCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

func (c *WebResourceDeleteCall) DoHeader() (http.Header, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
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
	//   "description": "Relinquish ownership of a website or domain.",
	//   "httpMethod": "DELETE",
	//   "id": "siteVerification.webResource.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.get":

type WebResourceGetCall struct {
	s    *Service
	id   string
	opt_ map[string]interface{}
}

// Get: Get the most current data for a website or domain.
func (r *WebResourceService) Get(id string) *WebResourceGetCall {
	c := &WebResourceGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceGetCall) Fields(s ...googleapi.Field) *WebResourceGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceGetCall) IfNoneMatch(ifNoneMatch string) *WebResourceGetCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceGetCall) Do() (*SiteVerificationWebResourceResource, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourceGetCall) DoHeader() (http.Header, *SiteVerificationWebResourceResource, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceResource
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Get the most current data for a website or domain.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.getToken":

type WebResourceGetTokenCall struct {
	s                                          *Service
	siteverificationwebresourcegettokenrequest *SiteVerificationWebResourceGettokenRequest
	opt_                                       map[string]interface{}
}

// GetToken: Get a verification token for placing on a website or
// domain.
func (r *WebResourceService) GetToken(siteverificationwebresourcegettokenrequest *SiteVerificationWebResourceGettokenRequest) *WebResourceGetTokenCall {
	c := &WebResourceGetTokenCall{s: r.s, opt_: make(map[string]interface{})}
	c.siteverificationwebresourcegettokenrequest = siteverificationwebresourcegettokenrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceGetTokenCall) Fields(s ...googleapi.Field) *WebResourceGetTokenCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceGetTokenCall) IfNoneMatch(ifNoneMatch string) *WebResourceGetTokenCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceGetTokenCall) Do() (*SiteVerificationWebResourceGettokenResponse, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourceGetTokenCall) DoHeader() (http.Header, *SiteVerificationWebResourceGettokenResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourcegettokenrequest)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "token")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceGettokenResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Get a verification token for placing on a website or domain.",
	//   "httpMethod": "POST",
	//   "id": "siteVerification.webResource.getToken",
	//   "path": "token",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceGettokenRequest"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceGettokenResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification",
	//     "https://www.googleapis.com/auth/siteverification.verify_only"
	//   ]
	// }

}

// method id "siteVerification.webResource.insert":

type WebResourceInsertCall struct {
	s                                   *Service
	verificationMethod                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Insert: Attempt verification of a website or domain.
func (r *WebResourceService) Insert(verificationMethod string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceInsertCall {
	c := &WebResourceInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.verificationMethod = verificationMethod
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceInsertCall) Fields(s ...googleapi.Field) *WebResourceInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceInsertCall) IfNoneMatch(ifNoneMatch string) *WebResourceInsertCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceInsertCall) Do() (*SiteVerificationWebResourceResource, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourceInsertCall) DoHeader() (http.Header, *SiteVerificationWebResourceResource, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("verificationMethod", fmt.Sprintf("%v", c.verificationMethod))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceResource
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Attempt verification of a website or domain.",
	//   "httpMethod": "POST",
	//   "id": "siteVerification.webResource.insert",
	//   "parameterOrder": [
	//     "verificationMethod"
	//   ],
	//   "parameters": {
	//     "verificationMethod": {
	//       "description": "The method to use for verifying a site or domain.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification",
	//     "https://www.googleapis.com/auth/siteverification.verify_only"
	//   ]
	// }

}

// method id "siteVerification.webResource.list":

type WebResourceListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Get the list of your verified websites and domains.
func (r *WebResourceService) List() *WebResourceListCall {
	c := &WebResourceListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceListCall) Fields(s ...googleapi.Field) *WebResourceListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceListCall) IfNoneMatch(ifNoneMatch string) *WebResourceListCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceListCall) Do() (*SiteVerificationWebResourceListResponse, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourceListCall) DoHeader() (http.Header, *SiteVerificationWebResourceListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Get the list of your verified websites and domains.",
	//   "httpMethod": "GET",
	//   "id": "siteVerification.webResource.list",
	//   "path": "webResource",
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.patch":

type WebResourcePatchCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Patch: Modify the list of owners for your website or domain. This
// method supports patch semantics.
func (r *WebResourceService) Patch(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourcePatchCall {
	c := &WebResourcePatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourcePatchCall) Fields(s ...googleapi.Field) *WebResourcePatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourcePatchCall) IfNoneMatch(ifNoneMatch string) *WebResourcePatchCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourcePatchCall) Do() (*SiteVerificationWebResourceResource, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourcePatchCall) DoHeader() (http.Header, *SiteVerificationWebResourceResource, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceResource
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Modify the list of owners for your website or domain. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "siteVerification.webResource.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}

// method id "siteVerification.webResource.update":

type WebResourceUpdateCall struct {
	s                                   *Service
	id                                  string
	siteverificationwebresourceresource *SiteVerificationWebResourceResource
	opt_                                map[string]interface{}
}

// Update: Modify the list of owners for your website or domain.
func (r *WebResourceService) Update(id string, siteverificationwebresourceresource *SiteVerificationWebResourceResource) *WebResourceUpdateCall {
	c := &WebResourceUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.id = id
	c.siteverificationwebresourceresource = siteverificationwebresourceresource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *WebResourceUpdateCall) Fields(s ...googleapi.Field) *WebResourceUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter
// "ifNoneMatch": Makes the operation conditional on whether
// the object's Etag does not match the given value.
func (c *WebResourceUpdateCall) IfNoneMatch(ifNoneMatch string) *WebResourceUpdateCall {
	c.opt_["ifNoneMatch"] = ifNoneMatch
	return c
}

func (c *WebResourceUpdateCall) Do() (*SiteVerificationWebResourceResource, error) {
	_, v, err := c.DoHeader()
	return v, err
}

func (c *WebResourceUpdateCall) DoHeader() (http.Header, *SiteVerificationWebResourceResource, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.siteverificationwebresourceresource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "webResource/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
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
		return res.Header, nil, err
	}
	var ret *SiteVerificationWebResourceResource
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return res.Header, nil, err
	}
	return res.Header, ret, nil
	// {
	//   "description": "Modify the list of owners for your website or domain.",
	//   "httpMethod": "PUT",
	//   "id": "siteVerification.webResource.update",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The id of a verified site or domain.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "webResource/{id}",
	//   "request": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "response": {
	//     "$ref": "SiteVerificationWebResourceResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/siteverification"
	//   ]
	// }

}
