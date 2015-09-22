// Package licensing provides access to the Enterprise License Manager API.
//
// See https://developers.google.com/google-apps/licensing/
//
// Usage example:
//
//   import "google.golang.org/api/licensing/v1"
//   ...
//   licensingService, err := licensing.New(oauthHttpClient)
package licensing

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
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

const apiId = "licensing:v1"
const apiName = "licensing"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/apps/licensing/v1/product/"

// OAuth2 scopes used by this API.
const (
	// View and manage Google Apps licenses for your domain
	AppsLicensingScope = "https://www.googleapis.com/auth/apps.licensing"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.LicenseAssignments = NewLicenseAssignmentsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	LicenseAssignments *LicenseAssignmentsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewLicenseAssignmentsService(s *Service) *LicenseAssignmentsService {
	rs := &LicenseAssignmentsService{s: s}
	return rs
}

type LicenseAssignmentsService struct {
	s *Service
}

// LicenseAssignment: Template for LiscenseAssignment Resource
type LicenseAssignment struct {
	// Etags: ETag of the resource.
	Etags string `json:"etags,omitempty"`

	// Kind: Identifies the resource as a LicenseAssignment.
	Kind string `json:"kind,omitempty"`

	// ProductId: Name of the product.
	ProductId string `json:"productId,omitempty"`

	// SelfLink: Link to this page.
	SelfLink string `json:"selfLink,omitempty"`

	// SkuId: Name of the sku of the product.
	SkuId string `json:"skuId,omitempty"`

	// UserId: Email id of the user.
	UserId string `json:"userId,omitempty"`
}

// LicenseAssignmentInsert: Template for LicenseAssignment Insert
// request
type LicenseAssignmentInsert struct {
	// UserId: Email id of the user
	UserId string `json:"userId,omitempty"`
}

// LicenseAssignmentList: LicesnseAssignment List for a given
// product/sku for a customer.
type LicenseAssignmentList struct {
	// Etag: ETag of the resource.
	Etag string `json:"etag,omitempty"`

	// Items: The LicenseAssignments in this page of results.
	Items []*LicenseAssignment `json:"items,omitempty"`

	// Kind: Identifies the resource as a collection of LicenseAssignments.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "licensing.licenseAssignments.delete":

type LicenseAssignmentsDeleteCall struct {
	s         *Service
	productId string
	skuId     string
	userId    string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// Delete: Revoke License.
func (r *LicenseAssignmentsService) Delete(productId string, skuId string, userId string) *LicenseAssignmentsDeleteCall {
	c := &LicenseAssignmentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsDeleteCall) Fields(s ...googleapi.Field) *LicenseAssignmentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsDeleteCall) Ctx(ctx context.Context) *LicenseAssignmentsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/user/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsDeleteCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Revoke License.",
	//   "httpMethod": "DELETE",
	//   "id": "licensing.licenseAssignments.delete",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.get":

type LicenseAssignmentsGetCall struct {
	s         *Service
	productId string
	skuId     string
	userId    string
	opt_      map[string]interface{}
	ctx_      context.Context
}

// Get: Get license assignment of a particular product and sku for a
// user
func (r *LicenseAssignmentsService) Get(productId string, skuId string, userId string) *LicenseAssignmentsGetCall {
	c := &LicenseAssignmentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.userId = userId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsGetCall) Fields(s ...googleapi.Field) *LicenseAssignmentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsGetCall) Ctx(ctx context.Context) *LicenseAssignmentsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/user/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsGetCall) Do() (*LicenseAssignment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get license assignment of a particular product and sku for a user",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.get",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.insert":

type LicenseAssignmentsInsertCall struct {
	s                       *Service
	productId               string
	skuId                   string
	licenseassignmentinsert *LicenseAssignmentInsert
	opt_                    map[string]interface{}
	ctx_                    context.Context
}

// Insert: Assign License.
func (r *LicenseAssignmentsService) Insert(productId string, skuId string, licenseassignmentinsert *LicenseAssignmentInsert) *LicenseAssignmentsInsertCall {
	c := &LicenseAssignmentsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.licenseassignmentinsert = licenseassignmentinsert
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsInsertCall) Fields(s ...googleapi.Field) *LicenseAssignmentsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsInsertCall) Ctx(ctx context.Context) *LicenseAssignmentsInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.licenseassignmentinsert)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/user")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsInsertCall) Do() (*LicenseAssignment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Assign License.",
	//   "httpMethod": "POST",
	//   "id": "licensing.licenseAssignments.insert",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user",
	//   "request": {
	//     "$ref": "LicenseAssignmentInsert"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.listForProduct":

type LicenseAssignmentsListForProductCall struct {
	s          *Service
	productId  string
	customerId string
	opt_       map[string]interface{}
	ctx_       context.Context
}

// ListForProduct: List license assignments for given product of the
// customer.
func (r *LicenseAssignmentsService) ListForProduct(productId string, customerId string) *LicenseAssignmentsListForProductCall {
	c := &LicenseAssignmentsListForProductCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.customerId = customerId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of campaigns to return at one time. Must be positive.  Default value
// is 100.
func (c *LicenseAssignmentsListForProductCall) MaxResults(maxResults int64) *LicenseAssignmentsListForProductCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Token to fetch the
// next page. By default server will return first page
func (c *LicenseAssignmentsListForProductCall) PageToken(pageToken string) *LicenseAssignmentsListForProductCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsListForProductCall) Fields(s ...googleapi.Field) *LicenseAssignmentsListForProductCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsListForProductCall) Ctx(ctx context.Context) *LicenseAssignmentsListForProductCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsListForProductCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("customerId", fmt.Sprintf("%v", c.customerId))
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsListForProductCall) Do() (*LicenseAssignmentList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignmentList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List license assignments for given product of the customer.",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.listForProduct",
	//   "parameterOrder": [
	//     "productId",
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "CustomerId represents the customer for whom licenseassignments are queried",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "100",
	//       "description": "Maximum number of campaigns to return at one time. Must be positive. Optional. Default value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "default": "",
	//       "description": "Token to fetch the next page.Optional. By default server will return first page",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/users",
	//   "response": {
	//     "$ref": "LicenseAssignmentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.listForProductAndSku":

type LicenseAssignmentsListForProductAndSkuCall struct {
	s          *Service
	productId  string
	skuId      string
	customerId string
	opt_       map[string]interface{}
	ctx_       context.Context
}

// ListForProductAndSku: List license assignments for given product and
// sku of the customer.
func (r *LicenseAssignmentsService) ListForProductAndSku(productId string, skuId string, customerId string) *LicenseAssignmentsListForProductAndSkuCall {
	c := &LicenseAssignmentsListForProductAndSkuCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.customerId = customerId
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of campaigns to return at one time. Must be positive.  Default value
// is 100.
func (c *LicenseAssignmentsListForProductAndSkuCall) MaxResults(maxResults int64) *LicenseAssignmentsListForProductAndSkuCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": Token to fetch the
// next page. By default server will return first page
func (c *LicenseAssignmentsListForProductAndSkuCall) PageToken(pageToken string) *LicenseAssignmentsListForProductAndSkuCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsListForProductAndSkuCall) Fields(s ...googleapi.Field) *LicenseAssignmentsListForProductAndSkuCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsListForProductAndSkuCall) Ctx(ctx context.Context) *LicenseAssignmentsListForProductAndSkuCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsListForProductAndSkuCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", alt)
	params.Set("customerId", fmt.Sprintf("%v", c.customerId))
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsListForProductAndSkuCall) Do() (*LicenseAssignmentList, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignmentList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List license assignments for given product and sku of the customer.",
	//   "httpMethod": "GET",
	//   "id": "licensing.licenseAssignments.listForProductAndSku",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "CustomerId represents the customer for whom licenseassignments are queried",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "100",
	//       "description": "Maximum number of campaigns to return at one time. Must be positive. Optional. Default value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "1000",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "default": "",
	//       "description": "Token to fetch the next page.Optional. By default server will return first page",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/users",
	//   "response": {
	//     "$ref": "LicenseAssignmentList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.patch":

type LicenseAssignmentsPatchCall struct {
	s                 *Service
	productId         string
	skuId             string
	userId            string
	licenseassignment *LicenseAssignment
	opt_              map[string]interface{}
	ctx_              context.Context
}

// Patch: Assign License. This method supports patch semantics.
func (r *LicenseAssignmentsService) Patch(productId string, skuId string, userId string, licenseassignment *LicenseAssignment) *LicenseAssignmentsPatchCall {
	c := &LicenseAssignmentsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.userId = userId
	c.licenseassignment = licenseassignment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsPatchCall) Fields(s ...googleapi.Field) *LicenseAssignmentsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsPatchCall) Ctx(ctx context.Context) *LicenseAssignmentsPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.licenseassignment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/user/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsPatchCall) Do() (*LicenseAssignment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Assign License. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "licensing.licenseAssignments.patch",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku for which license would be revoked",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "request": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}

// method id "licensing.licenseAssignments.update":

type LicenseAssignmentsUpdateCall struct {
	s                 *Service
	productId         string
	skuId             string
	userId            string
	licenseassignment *LicenseAssignment
	opt_              map[string]interface{}
	ctx_              context.Context
}

// Update: Assign License.
func (r *LicenseAssignmentsService) Update(productId string, skuId string, userId string, licenseassignment *LicenseAssignment) *LicenseAssignmentsUpdateCall {
	c := &LicenseAssignmentsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.productId = productId
	c.skuId = skuId
	c.userId = userId
	c.licenseassignment = licenseassignment
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LicenseAssignmentsUpdateCall) Fields(s ...googleapi.Field) *LicenseAssignmentsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// Ctx sets the context to be used in this call's Do method.
// Any pending HTTP request will be aborted if the provided context
// is cancelled.
func (c *LicenseAssignmentsUpdateCall) Ctx(ctx context.Context) *LicenseAssignmentsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *LicenseAssignmentsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.licenseassignment)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", alt)
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{productId}/sku/{skuId}/user/{userId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"productId": c.productId,
		"skuId":     c.skuId,
		"userId":    c.userId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LicenseAssignmentsUpdateCall) Do() (*LicenseAssignment, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LicenseAssignment
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Assign License.",
	//   "httpMethod": "PUT",
	//   "id": "licensing.licenseAssignments.update",
	//   "parameterOrder": [
	//     "productId",
	//     "skuId",
	//     "userId"
	//   ],
	//   "parameters": {
	//     "productId": {
	//       "description": "Name for product",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "skuId": {
	//       "description": "Name for sku for which license would be revoked",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "email id or unique Id of the user",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{productId}/sku/{skuId}/user/{userId}",
	//   "request": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "response": {
	//     "$ref": "LicenseAssignment"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.licensing"
	//   ]
	// }

}
