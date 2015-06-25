// Package fitness provides access to the Fitness.
//
// See https://developers.google.com/fit/rest/
//
// Usage example:
//
//   import "google.golang.org/api/fitness/v1"
//   ...
//   fitnessService, err := fitness.New(oauthHttpClient)
package fitness

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

const apiId = "fitness:v1"
const apiName = "fitness"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/fitness/v1/users/"

// OAuth2 scopes used by this API.
const (
	// View your activity information in Google Fit
	FitnessActivityReadScope = "https://www.googleapis.com/auth/fitness.activity.read"

	// View and store your activity information in Google Fit
	FitnessActivityWriteScope = "https://www.googleapis.com/auth/fitness.activity.write"

	// View body sensor information in Google Fit
	FitnessBodyReadScope = "https://www.googleapis.com/auth/fitness.body.read"

	// View and store body sensor data in Google Fit
	FitnessBodyWriteScope = "https://www.googleapis.com/auth/fitness.body.write"

	// View your stored location data in Google Fit
	FitnessLocationReadScope = "https://www.googleapis.com/auth/fitness.location.read"

	// View and store your location data in Google Fit
	FitnessLocationWriteScope = "https://www.googleapis.com/auth/fitness.location.write"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.DataSources = NewUsersDataSourcesService(s)
	rs.Dataset = NewUsersDatasetService(s)
	rs.Sessions = NewUsersSessionsService(s)
	return rs
}

type UsersService struct {
	s *Service

	DataSources *UsersDataSourcesService

	Dataset *UsersDatasetService

	Sessions *UsersSessionsService
}

func NewUsersDataSourcesService(s *Service) *UsersDataSourcesService {
	rs := &UsersDataSourcesService{s: s}
	rs.Datasets = NewUsersDataSourcesDatasetsService(s)
	return rs
}

type UsersDataSourcesService struct {
	s *Service

	Datasets *UsersDataSourcesDatasetsService
}

func NewUsersDataSourcesDatasetsService(s *Service) *UsersDataSourcesDatasetsService {
	rs := &UsersDataSourcesDatasetsService{s: s}
	return rs
}

type UsersDataSourcesDatasetsService struct {
	s *Service
}

func NewUsersDatasetService(s *Service) *UsersDatasetService {
	rs := &UsersDatasetService{s: s}
	return rs
}

type UsersDatasetService struct {
	s *Service
}

func NewUsersSessionsService(s *Service) *UsersSessionsService {
	rs := &UsersSessionsService{s: s}
	return rs
}

type UsersSessionsService struct {
	s *Service
}

type AggregateBucket struct {
	// Activity: available for Bucket.Type.ACTIVITY_TYPE,
	// Bucket.Type.ACTIVITY_SEGMENT
	Activity int64 `json:"activity,omitempty"`

	// Dataset: There will be one dataset per datatype/datasource
	Dataset []*Dataset `json:"dataset,omitempty"`

	EndTimeMillis int64 `json:"endTimeMillis,omitempty,string"`

	// Session: available for Bucket.Type.SESSION
	Session *Session `json:"session,omitempty"`

	StartTimeMillis int64 `json:"startTimeMillis,omitempty,string"`

	// Type: The type of a bucket signifies how the data aggregation is
	// performed in the bucket.
	//
	// Possible values:
	//   "activitySegment"
	//   "activityType"
	//   "session"
	//   "time"
	//   "unknown"
	Type string `json:"type,omitempty"`
}

type AggregateBy struct {
	DataSourceId string `json:"dataSourceId,omitempty"`

	// DataTypeName: by dataype or by datasource
	DataTypeName string `json:"dataTypeName,omitempty"`

	OutputDataSourceId string `json:"outputDataSourceId,omitempty"`

	OutputDataTypeName string `json:"outputDataTypeName,omitempty"`
}

type AggregateRequest struct {
	AggregateBy []*AggregateBy `json:"aggregateBy,omitempty"`

	BucketByActivitySegment *BucketByActivity `json:"bucketByActivitySegment,omitempty"`

	BucketByActivityType *BucketByActivity `json:"bucketByActivityType,omitempty"`

	BucketBySession *BucketBySession `json:"bucketBySession,omitempty"`

	// BucketByTime: apparently oneof is not supported by reduced_nano_proto
	BucketByTime *BucketByTime `json:"bucketByTime,omitempty"`

	EndTimeMillis int64 `json:"endTimeMillis,omitempty,string"`

	// StartTimeMillis: required time range
	StartTimeMillis int64 `json:"startTimeMillis,omitempty,string"`
}

type AggregateResponse struct {
	Bucket []*AggregateBucket `json:"bucket,omitempty"`
}

type Application struct {
	// DetailsUrl: An optional URI that can be used to link back to the
	// application.
	DetailsUrl string `json:"detailsUrl,omitempty"`

	// Name: The name of this application. This is required for REST
	// clients, but we do not enforce uniqueness of this name. It is
	// provided as a matter of convenience for other developers who would
	// like to identify which REST created an Application or Data Source.
	Name string `json:"name,omitempty"`

	// PackageName: Package name for this application. This is used as a
	// unique identifier when created by Android applications, but cannot be
	// specified by REST clients. REST clients will have their developer
	// project number reflected into the Data Source data stream IDs,
	// instead of the packageName.
	PackageName string `json:"packageName,omitempty"`

	// Version: Version of the application. You should update this field
	// whenever the application changes in a way that affects the
	// computation of the data.
	Version string `json:"version,omitempty"`
}

type BucketByActivity struct {
	// ActivityDataSourceId: default activity stream will be used if not
	// specified
	ActivityDataSourceId string `json:"activityDataSourceId,omitempty"`

	// MinDurationMillis: Only activity segments of duration longer than
	// this is used
	MinDurationMillis int64 `json:"minDurationMillis,omitempty,string"`
}

type BucketBySession struct {
	// MinDurationMillis: Only sessions of duration longer than this is used
	MinDurationMillis int64 `json:"minDurationMillis,omitempty,string"`
}

type BucketByTime struct {
	DurationMillis int64 `json:"durationMillis,omitempty,string"`
}

type DataPoint struct {
	// ComputationTimeMillis: Used for version checking during
	// transformation; that is, a datapoint can only replace another
	// datapoint that has an older computation time stamp.
	ComputationTimeMillis int64 `json:"computationTimeMillis,omitempty,string"`

	// DataTypeName: The data type defining the format of the values in this
	// data point.
	DataTypeName string `json:"dataTypeName,omitempty"`

	// EndTimeNanos: The end time of the interval represented by this data
	// point, in nanoseconds since epoch.
	EndTimeNanos int64 `json:"endTimeNanos,omitempty,string"`

	// ModifiedTimeMillis: Indicates the last time this data point was
	// modified. Useful only in contexts where we are listing the data
	// changes, rather than representing the current state of the data.
	ModifiedTimeMillis int64 `json:"modifiedTimeMillis,omitempty,string"`

	// OriginDataSourceId: If the data point is contained in a dataset for a
	// derived data source, this field will be populated with the data
	// source stream ID that created the data point originally.
	OriginDataSourceId string `json:"originDataSourceId,omitempty"`

	// RawTimestampNanos: The raw timestamp from the original SensorEvent.
	RawTimestampNanos int64 `json:"rawTimestampNanos,omitempty,string"`

	// StartTimeNanos: The start time of the interval represented by this
	// data point, in nanoseconds since epoch.
	StartTimeNanos int64 `json:"startTimeNanos,omitempty,string"`

	// Value: Values of each data type field for the data point. It is
	// expected that each value corresponding to a data type field will
	// occur in the same order that the field is listed with in the data
	// type specified in a data source.
	//
	// Only one of integer and floating point fields will be populated,
	// depending on the format enum value within data source's type field.
	Value []*Value `json:"value,omitempty"`
}

type DataSource struct {
	// Application: Information about an application which feeds sensor data
	// into the platform.
	Application *Application `json:"application,omitempty"`

	// DataStreamId: A unique identifier for the data stream produced by
	// this data source. The identifier includes:
	//
	//
	// - The physical device's manufacturer, model, and serial number (UID).
	//
	// - The application's package name or name. Package name is used when
	// the data source was created by an Android application. The developer
	// project number is used when the data source was created by a REST
	// client.
	// - The data source's type.
	// - The data source's stream name.  Note that not all attributes of the
	// data source are used as part of the stream identifier. In particular,
	// the version of the hardware/the application isn't used. This allows
	// us to preserve the same stream through version updates. This also
	// means that two DataSource objects may represent the same data stream
	// even if they're not equal.
	//
	// The exact format of the data stream ID created by an Android
	// application is:
	// type:dataType.name:application.packageName:device.manufacturer:device.
	// model:device.uid:dataStreamName
	//
	// The exact format of the data stream ID created by a REST client is:
	// type:dataType.name:developer project
	// number:device.manufacturer:device.model:device.uid:dataStreamName
	//
	//
	// When any of the optional fields that comprise of the data stream ID
	// are blank, they will be omitted from the data stream ID. The minnimum
	// viable data stream ID would be: type:dataType.name:developer project
	// number
	//
	// Finally, the developer project number is obfuscated when read by any
	// REST or Android client that did not create the data source. Only the
	// data source creator will see the developer project number in clear
	// and normal form.
	DataStreamId string `json:"dataStreamId,omitempty"`

	// DataStreamName: The stream name uniquely identifies this particular
	// data source among other data sources of the same type from the same
	// underlying producer. Setting the stream name is optional, but should
	// be done whenever an application exposes two streams for the same data
	// type, or when a device has two equivalent sensors.
	DataStreamName string `json:"dataStreamName,omitempty"`

	// DataType: The data type defines the schema for a stream of data being
	// collected by, inserted into, or queried from the Fitness API.
	DataType *DataType `json:"dataType,omitempty"`

	// Device: Representation of an integrated device (such as a phone or a
	// wearable) that can hold sensors.
	Device *Device `json:"device,omitempty"`

	// Name: An end-user visible name for this data source.
	Name string `json:"name,omitempty"`

	// Type: A constant describing the type of this data source. Indicates
	// whether this data source produces raw or derived data.
	//
	// Possible values:
	//   "derived"
	//   "raw"
	Type string `json:"type,omitempty"`
}

type DataType struct {
	// Field: A field represents one dimension of a data type.
	Field []*DataTypeField `json:"field,omitempty"`

	// Name: Each data type has a unique, namespaced, name. All data types
	// in the com.google namespace are shared as part of the platform.
	Name string `json:"name,omitempty"`
}

type DataTypeField struct {
	// Format: The different supported formats for each field in a data
	// type.
	//
	// Possible values:
	//   "floatList"
	//   "floatPoint"
	//   "integer"
	//   "integerList"
	//   "map"
	//   "string"
	Format string `json:"format,omitempty"`

	// Name: Defines the name and format of data. Unlike data type names,
	// field names are not namespaced, and only need to be unique within the
	// data type.
	Name string `json:"name,omitempty"`

	Optional bool `json:"optional,omitempty"`
}

type Dataset struct {
	// DataSourceId: The data stream ID of the data source that created the
	// points in this dataset.
	DataSourceId string `json:"dataSourceId,omitempty"`

	// MaxEndTimeNs: The largest end time of all data points in this
	// possibly partial representation of the dataset. Time is in
	// nanoseconds from epoch. This should also match the first part of the
	// dataset identifier.
	MaxEndTimeNs int64 `json:"maxEndTimeNs,omitempty,string"`

	// MinStartTimeNs: The smallest start time of all data points in this
	// possibly partial representation of the dataset. Time is in
	// nanoseconds from epoch. This should also match the first part of the
	// dataset identifier.
	MinStartTimeNs int64 `json:"minStartTimeNs,omitempty,string"`

	// NextPageToken: This token will be set when a dataset is received in
	// response to a GET request and the dataset is too large to be included
	// in a single response. Provide this value in a subsequent GET request
	// to return the next page of data points within this dataset.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Point: A partial list of data points contained in the dataset,
	// ordered by largest endTimeNanos first. This list is considered
	// complete when retrieving a small dataset and partial when patching a
	// dataset or retrieving a dataset that is too large to include in a
	// single response.
	Point []*DataPoint `json:"point,omitempty"`
}

type Device struct {
	// Manufacturer: Manufacturer of the product/hardware.
	Manufacturer string `json:"manufacturer,omitempty"`

	// Model: End-user visible model name for the device.
	Model string `json:"model,omitempty"`

	// Type: A constant representing the type of the device.
	//
	// Possible values:
	//   "chestStrap"
	//   "phone"
	//   "scale"
	//   "tablet"
	//   "unknown"
	//   "watch"
	Type string `json:"type,omitempty"`

	// Uid: The serial number or other unique ID for the hardware. This
	// field is obfuscated when read by any REST or Android client that did
	// not create the data source. Only the data source creator will see the
	// uid field in clear and normal form.
	Uid string `json:"uid,omitempty"`

	// Version: Version string for the device hardware/software.
	Version string `json:"version,omitempty"`
}

type ListDataSourcesResponse struct {
	// DataSource: A previously created data source.
	DataSource []*DataSource `json:"dataSource,omitempty"`
}

type ListSessionsResponse struct {
	// DeletedSession: If includeDeleted is set to true in the request, this
	// list will contain sessions deleted with original end times that are
	// within the startTime and endTime frame.
	DeletedSession []*Session `json:"deletedSession,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Session: Sessions with an end time that is between startTime and
	// endTime of the request.
	Session []*Session `json:"session,omitempty"`
}

type Session struct {
	// ActiveTimeMillis: Session active time. While start_time_millis and
	// end_time_millis define the full session time, the active time can be
	// shorter and specified by active_time_millis. If the inactive time
	// during the session is known, it should also be inserted via a
	// com.google.activity.segment data point with a STILL activity value
	ActiveTimeMillis int64 `json:"activeTimeMillis,omitempty,string"`

	// ActivityType: The type of activity this session represents.
	ActivityType int64 `json:"activityType,omitempty"`

	// Application: The application that created the session.
	Application *Application `json:"application,omitempty"`

	// Description: A description for this session.
	Description string `json:"description,omitempty"`

	// EndTimeMillis: An end time, in milliseconds since epoch, inclusive.
	EndTimeMillis int64 `json:"endTimeMillis,omitempty,string"`

	// Id: A client-generated identifier that is unique across all sessions
	// owned by this particular user.
	Id string `json:"id,omitempty"`

	// ModifiedTimeMillis: A timestamp that indicates when the session was
	// last modified.
	ModifiedTimeMillis int64 `json:"modifiedTimeMillis,omitempty,string"`

	// Name: A human readable name of the session.
	Name string `json:"name,omitempty"`

	// StartTimeMillis: A start time, in milliseconds since epoch,
	// inclusive.
	StartTimeMillis int64 `json:"startTimeMillis,omitempty,string"`
}

type Value struct {
	// FpVal: Floating point value. When this is set, intVal must not be
	// set.
	FpVal float64 `json:"fpVal,omitempty"`

	// IntVal: Integer value. When this is set, fpVal must not be set.
	IntVal int64 `json:"intVal,omitempty"`
}

// method id "fitness.users.dataSources.create":

type UsersDataSourcesCreateCall struct {
	s          *Service
	userId     string
	datasource *DataSource
	opt_       map[string]interface{}
}

// Create: Creates a new data source that is unique across all data
// sources belonging to this user. The data stream ID field can be
// omitted and will be generated by the server with the correct format.
// The data stream ID is an ordered combination of some fields from the
// data source. In addition to the data source fields reflected into the
// data source ID, the developer project number that is authenticated
// when creating the data source is included. This developer project
// number is obfuscated when read by any other developer reading public
// data types.
func (r *UsersDataSourcesService) Create(userId string, datasource *DataSource) *UsersDataSourcesCreateCall {
	c := &UsersDataSourcesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.datasource = datasource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesCreateCall) Fields(s ...googleapi.Field) *UsersDataSourcesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesCreateCall) IfNoneMatch(entityTag string) *UsersDataSourcesCreateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.create" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesCreateCall) Do() (*DataSource, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.create" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesCreateCall) DoHeader() (ret *DataSource, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.datasource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
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
	//   "description": "Creates a new data source that is unique across all data sources belonging to this user. The data stream ID field can be omitted and will be generated by the server with the correct format. The data stream ID is an ordered combination of some fields from the data source. In addition to the data source fields reflected into the data source ID, the developer project number that is authenticated when creating the data source is included. This developer project number is obfuscated when read by any other developer reading public data types.",
	//   "httpMethod": "POST",
	//   "id": "fitness.users.dataSources.create",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "description": "Create the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.delete":

type UsersDataSourcesDeleteCall struct {
	s            *Service
	userId       string
	dataSourceId string
	opt_         map[string]interface{}
}

// Delete: Delete the data source if there are no datapoints associated
// with it
func (r *UsersDataSourcesService) Delete(userId string, dataSourceId string) *UsersDataSourcesDeleteCall {
	c := &UsersDataSourcesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDeleteCall) Fields(s ...googleapi.Field) *UsersDataSourcesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesDeleteCall) IfNoneMatch(entityTag string) *UsersDataSourcesDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.delete" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesDeleteCall) Do() (*DataSource, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesDeleteCall) DoHeader() (ret *DataSource, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
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
	//   "description": "Delete the data source if there are no datapoints associated with it",
	//   "httpMethod": "DELETE",
	//   "id": "fitness.users.dataSources.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Retrieve a data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.get":

type UsersDataSourcesGetCall struct {
	s            *Service
	userId       string
	dataSourceId string
	opt_         map[string]interface{}
}

// Get: Returns a data source identified by a data stream ID.
func (r *UsersDataSourcesService) Get(userId string, dataSourceId string) *UsersDataSourcesGetCall {
	c := &UsersDataSourcesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesGetCall) Fields(s ...googleapi.Field) *UsersDataSourcesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesGetCall) IfNoneMatch(entityTag string) *UsersDataSourcesGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.get" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesGetCall) Do() (*DataSource, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesGetCall) DoHeader() (ret *DataSource, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
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
	//   "description": "Returns a data source identified by a data stream ID.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.get",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Retrieve a data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.list":

type UsersDataSourcesListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Lists all data sources that are visible to the developer, using
// the OAuth scopes provided. The list is not exhaustive: the user may
// have private data sources that are only visible to other developers
// or calls using other scopes.
func (r *UsersDataSourcesService) List(userId string) *UsersDataSourcesListCall {
	c := &UsersDataSourcesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// DataTypeName sets the optional parameter "dataTypeName": The names of
// data types to include in the list. If not specified, all data sources
// will be returned.
func (c *UsersDataSourcesListCall) DataTypeName(dataTypeName string) *UsersDataSourcesListCall {
	c.opt_["dataTypeName"] = dataTypeName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesListCall) Fields(s ...googleapi.Field) *UsersDataSourcesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesListCall) IfNoneMatch(entityTag string) *UsersDataSourcesListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.list" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesListCall) Do() (*ListDataSourcesResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesListCall) DoHeader() (ret *ListDataSourcesResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["dataTypeName"]; ok {
		params.Set("dataTypeName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
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
	//   "description": "Lists all data sources that are visible to the developer, using the OAuth scopes provided. The list is not exhaustive: the user may have private data sources that are only visible to other developers or calls using other scopes.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "dataTypeName": {
	//       "description": "The names of data types to include in the list. If not specified, all data sources will be returned.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "List data sources for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources",
	//   "response": {
	//     "$ref": "ListDataSourcesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.patch":

type UsersDataSourcesPatchCall struct {
	s            *Service
	userId       string
	dataSourceId string
	datasource   *DataSource
	opt_         map[string]interface{}
}

// Patch: Updates a given data source. It is an error to modify the data
// source's data stream ID, data type, type, stream name or device
// information apart from the device version. Changing these fields
// would require a new unique data stream ID and separate data
// source.
//
// Data sources are identified by their data stream ID. This method
// supports patch semantics.
func (r *UsersDataSourcesService) Patch(userId string, dataSourceId string, datasource *DataSource) *UsersDataSourcesPatchCall {
	c := &UsersDataSourcesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	c.datasource = datasource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesPatchCall) Fields(s ...googleapi.Field) *UsersDataSourcesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesPatchCall) IfNoneMatch(entityTag string) *UsersDataSourcesPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.patch" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesPatchCall) Do() (*DataSource, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesPatchCall) DoHeader() (ret *DataSource, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.datasource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
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
	//   "description": "Updates a given data source. It is an error to modify the data source's data stream ID, data type, type, stream name or device information apart from the device version. Changing these fields would require a new unique data stream ID and separate data source.\n\nData sources are identified by their data stream ID. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fitness.users.dataSources.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Update the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.update":

type UsersDataSourcesUpdateCall struct {
	s            *Service
	userId       string
	dataSourceId string
	datasource   *DataSource
	opt_         map[string]interface{}
}

// Update: Updates a given data source. It is an error to modify the
// data source's data stream ID, data type, type, stream name or device
// information apart from the device version. Changing these fields
// would require a new unique data stream ID and separate data
// source.
//
// Data sources are identified by their data stream ID.
func (r *UsersDataSourcesService) Update(userId string, dataSourceId string, datasource *DataSource) *UsersDataSourcesUpdateCall {
	c := &UsersDataSourcesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	c.datasource = datasource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesUpdateCall) Fields(s ...googleapi.Field) *UsersDataSourcesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesUpdateCall) IfNoneMatch(entityTag string) *UsersDataSourcesUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.update" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesUpdateCall) Do() (*DataSource, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesUpdateCall) DoHeader() (ret *DataSource, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.datasource)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
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
	//   "description": "Updates a given data source. It is an error to modify the data source's data stream ID, data type, type, stream name or device information apart from the device version. Changing these fields would require a new unique data stream ID and separate data source.\n\nData sources are identified by their data stream ID.",
	//   "httpMethod": "PUT",
	//   "id": "fitness.users.dataSources.update",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Update the data source for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}",
	//   "request": {
	//     "$ref": "DataSource"
	//   },
	//   "response": {
	//     "$ref": "DataSource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.delete":

type UsersDataSourcesDatasetsDeleteCall struct {
	s            *Service
	userId       string
	dataSourceId string
	datasetId    string
	opt_         map[string]interface{}
}

// Delete: Performs an inclusive delete of all data points whose start
// and end times have any overlap with the time range specified by the
// dataset ID. For most data types, the entire data point will be
// deleted. For data types where the time span represents a consistent
// value (such as com.google.activity.segment), and a data point
// straddles either end point of the dataset, only the overlapping
// portion of the data point will be deleted.
func (r *UsersDataSourcesDatasetsService) Delete(userId string, dataSourceId string, datasetId string) *UsersDataSourcesDatasetsDeleteCall {
	c := &UsersDataSourcesDatasetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	c.datasetId = datasetId
	return c
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersDataSourcesDatasetsDeleteCall) CurrentTimeMillis(currentTimeMillis int64) *UsersDataSourcesDatasetsDeleteCall {
	c.opt_["currentTimeMillis"] = currentTimeMillis
	return c
}

// ModifiedTimeMillis sets the optional parameter "modifiedTimeMillis":
// When the operation was performed on the client.
func (c *UsersDataSourcesDatasetsDeleteCall) ModifiedTimeMillis(modifiedTimeMillis int64) *UsersDataSourcesDatasetsDeleteCall {
	c.opt_["modifiedTimeMillis"] = modifiedTimeMillis
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDatasetsDeleteCall) Fields(s ...googleapi.Field) *UsersDataSourcesDatasetsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesDatasetsDeleteCall) IfNoneMatch(entityTag string) *UsersDataSourcesDatasetsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.datasets.delete" call.
func (c *UsersDataSourcesDatasetsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "fitness.users.dataSources.datasets.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesDatasetsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["currentTimeMillis"]; ok {
		params.Set("currentTimeMillis", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["modifiedTimeMillis"]; ok {
		params.Set("modifiedTimeMillis", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
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
	//   "description": "Performs an inclusive delete of all data points whose start and end times have any overlap with the time range specified by the dataset ID. For most data types, the entire data point will be deleted. For data types where the time span represents a consistent value (such as com.google.activity.segment), and a data point straddles either end point of the dataset, only the overlapping portion of the data point will be deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "fitness.users.dataSources.datasets.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "modifiedTimeMillis": {
	//       "description": "When the operation was performed on the client.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Delete a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.get":

type UsersDataSourcesDatasetsGetCall struct {
	s            *Service
	userId       string
	dataSourceId string
	datasetId    string
	opt_         map[string]interface{}
}

// Get: Returns a dataset containing all data points whose start and end
// times overlap with the specified range of the dataset minimum start
// time and maximum end time. Specifically, any data point whose start
// time is less than or equal to the dataset end time and whose end time
// is greater than or equal to the dataset start time.
func (r *UsersDataSourcesDatasetsService) Get(userId string, dataSourceId string, datasetId string) *UsersDataSourcesDatasetsGetCall {
	c := &UsersDataSourcesDatasetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	c.datasetId = datasetId
	return c
}

// Limit sets the optional parameter "limit": If specified, no more than
// this many data points will be included in the dataset. If the there
// are more data points in the dataset, nextPageToken will be set in the
// dataset response.
func (c *UsersDataSourcesDatasetsGetCall) Limit(limit int64) *UsersDataSourcesDatasetsGetCall {
	c.opt_["limit"] = limit
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large datasets. To get the next
// page of a dataset, set this parameter to the value of nextPageToken
// from the previous response. Each subsequent call will yield a partial
// dataset with data point end timestamps that are strictly smaller than
// those in the previous partial response.
func (c *UsersDataSourcesDatasetsGetCall) PageToken(pageToken string) *UsersDataSourcesDatasetsGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDatasetsGetCall) Fields(s ...googleapi.Field) *UsersDataSourcesDatasetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesDatasetsGetCall) IfNoneMatch(entityTag string) *UsersDataSourcesDatasetsGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.datasets.get" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesDatasetsGetCall) Do() (*Dataset, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.datasets.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesDatasetsGetCall) DoHeader() (ret *Dataset, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["limit"]; ok {
		params.Set("limit", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
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
	//   "description": "Returns a dataset containing all data points whose start and end times overlap with the specified range of the dataset minimum start time and maximum end time. Specifically, any data point whose start time is less than or equal to the dataset end time and whose end time is greater than or equal to the dataset start time.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.dataSources.datasets.get",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "limit": {
	//       "description": "If specified, no more than this many data points will be included in the dataset. If the there are more data points in the dataset, nextPageToken will be set in the dataset response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large datasets. To get the next page of a dataset, set this parameter to the value of nextPageToken from the previous response. Each subsequent call will yield a partial dataset with data point end timestamps that are strictly smaller than those in the previous partial response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Retrieve a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataSources.datasets.patch":

type UsersDataSourcesDatasetsPatchCall struct {
	s            *Service
	userId       string
	dataSourceId string
	datasetId    string
	dataset      *Dataset
	opt_         map[string]interface{}
}

// Patch: Adds data points to a dataset. The dataset need not be
// previously created. All points within the given dataset will be
// returned with subsquent calls to retrieve this dataset. Data points
// can belong to more than one dataset. This method does not use patch
// semantics.
func (r *UsersDataSourcesDatasetsService) Patch(userId string, dataSourceId string, datasetId string, dataset *Dataset) *UsersDataSourcesDatasetsPatchCall {
	c := &UsersDataSourcesDatasetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.dataSourceId = dataSourceId
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch. Note that the
// minStartTimeNs and maxEndTimeNs properties in the request body are in
// nanoseconds instead of milliseconds.
func (c *UsersDataSourcesDatasetsPatchCall) CurrentTimeMillis(currentTimeMillis int64) *UsersDataSourcesDatasetsPatchCall {
	c.opt_["currentTimeMillis"] = currentTimeMillis
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDataSourcesDatasetsPatchCall) Fields(s ...googleapi.Field) *UsersDataSourcesDatasetsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDataSourcesDatasetsPatchCall) IfNoneMatch(entityTag string) *UsersDataSourcesDatasetsPatchCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataSources.datasets.patch" call.
// Exactly one of the return values is non-nil.
func (c *UsersDataSourcesDatasetsPatchCall) Do() (*Dataset, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataSources.datasets.patch" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDataSourcesDatasetsPatchCall) DoHeader() (ret *Dataset, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["currentTimeMillis"]; ok {
		params.Set("currentTimeMillis", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":       c.userId,
		"dataSourceId": c.dataSourceId,
		"datasetId":    c.datasetId,
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
	//   "description": "Adds data points to a dataset. The dataset need not be previously created. All points within the given dataset will be returned with subsquent calls to retrieve this dataset. Data points can belong to more than one dataset. This method does not use patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "fitness.users.dataSources.datasets.patch",
	//   "parameterOrder": [
	//     "userId",
	//     "dataSourceId",
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch. Note that the minStartTimeNs and maxEndTimeNs properties in the request body are in nanoseconds instead of milliseconds.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dataSourceId": {
	//       "description": "The data stream ID of the data source that created the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "datasetId": {
	//       "description": "Dataset identifier that is a composite of the minimum data point start time and maximum data point end time represented as nanoseconds from the epoch. The ID is formatted like: \"startTime-endTime\" where startTime and endTime are 64 bit integers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Patch a dataset for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataSources/{dataSourceId}/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.dataset.aggregate":

type UsersDatasetAggregateCall struct {
	s                *Service
	userId           string
	aggregaterequest *AggregateRequest
	opt_             map[string]interface{}
}

// Aggregate:
func (r *UsersDatasetService) Aggregate(userId string, aggregaterequest *AggregateRequest) *UsersDatasetAggregateCall {
	c := &UsersDatasetAggregateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.aggregaterequest = aggregaterequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDatasetAggregateCall) Fields(s ...googleapi.Field) *UsersDatasetAggregateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersDatasetAggregateCall) IfNoneMatch(entityTag string) *UsersDatasetAggregateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.dataset.aggregate" call.
// Exactly one of the return values is non-nil.
func (c *UsersDatasetAggregateCall) Do() (*AggregateResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.dataset.aggregate" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersDatasetAggregateCall) DoHeader() (ret *AggregateResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.aggregaterequest)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/dataset:aggregate")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
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
	//   "httpMethod": "POST",
	//   "id": "fitness.users.dataset.aggregate",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "userId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/dataset:aggregate",
	//   "request": {
	//     "$ref": "AggregateRequest"
	//   },
	//   "response": {
	//     "$ref": "AggregateResponse"
	//   }
	// }

}

// method id "fitness.users.sessions.delete":

type UsersSessionsDeleteCall struct {
	s         *Service
	userId    string
	sessionId string
	opt_      map[string]interface{}
}

// Delete: Deletes a session specified by the given session ID.
func (r *UsersSessionsService) Delete(userId string, sessionId string) *UsersSessionsDeleteCall {
	c := &UsersSessionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.sessionId = sessionId
	return c
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersSessionsDeleteCall) CurrentTimeMillis(currentTimeMillis int64) *UsersSessionsDeleteCall {
	c.opt_["currentTimeMillis"] = currentTimeMillis
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersSessionsDeleteCall) Fields(s ...googleapi.Field) *UsersSessionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersSessionsDeleteCall) IfNoneMatch(entityTag string) *UsersSessionsDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.sessions.delete" call.
func (c *UsersSessionsDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "fitness.users.sessions.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersSessionsDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["currentTimeMillis"]; ok {
		params.Set("currentTimeMillis", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/sessions/{sessionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":    c.userId,
		"sessionId": c.sessionId,
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
	//   "description": "Deletes a session specified by the given session ID.",
	//   "httpMethod": "DELETE",
	//   "id": "fitness.users.sessions.delete",
	//   "parameterOrder": [
	//     "userId",
	//     "sessionId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sessionId": {
	//       "description": "The ID of the session to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Delete a session for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions/{sessionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write"
	//   ]
	// }

}

// method id "fitness.users.sessions.list":

type UsersSessionsListCall struct {
	s      *Service
	userId string
	opt_   map[string]interface{}
}

// List: Lists sessions previously created.
func (r *UsersSessionsService) List(userId string) *UsersSessionsListCall {
	c := &UsersSessionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	return c
}

// EndTime sets the optional parameter "endTime": An RFC3339 timestamp.
// Only sessions ending between the start and end times will be included
// in the response.
func (c *UsersSessionsListCall) EndTime(endTime string) *UsersSessionsListCall {
	c.opt_["endTime"] = endTime
	return c
}

// IncludeDeleted sets the optional parameter "includeDeleted": If true,
// deleted sessions will be returned. When set to true, sessions
// returned in this response will only have an ID and will not have any
// other fields.
func (c *UsersSessionsListCall) IncludeDeleted(includeDeleted bool) *UsersSessionsListCall {
	c.opt_["includeDeleted"] = includeDeleted
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// nextPageToken from the previous response.
func (c *UsersSessionsListCall) PageToken(pageToken string) *UsersSessionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartTime sets the optional parameter "startTime": An RFC3339
// timestamp. Only sessions ending between the start and end times will
// be included in the response.
func (c *UsersSessionsListCall) StartTime(startTime string) *UsersSessionsListCall {
	c.opt_["startTime"] = startTime
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersSessionsListCall) Fields(s ...googleapi.Field) *UsersSessionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersSessionsListCall) IfNoneMatch(entityTag string) *UsersSessionsListCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.sessions.list" call.
// Exactly one of the return values is non-nil.
func (c *UsersSessionsListCall) Do() (*ListSessionsResponse, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.sessions.list" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersSessionsListCall) DoHeader() (ret *ListSessionsResponse, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["endTime"]; ok {
		params.Set("endTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["includeDeleted"]; ok {
		params.Set("includeDeleted", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/sessions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
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
	//   "description": "Lists sessions previously created.",
	//   "httpMethod": "GET",
	//   "id": "fitness.users.sessions.list",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "An RFC3339 timestamp. Only sessions ending between the start and end times will be included in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "includeDeleted": {
	//       "description": "If true, deleted sessions will be returned. When set to true, sessions returned in this response will only have an ID and will not have any other fields.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "An RFC3339 timestamp. Only sessions ending between the start and end times will be included in the response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "List sessions for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions",
	//   "response": {
	//     "$ref": "ListSessionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.read",
	//     "https://www.googleapis.com/auth/fitness.activity.write",
	//     "https://www.googleapis.com/auth/fitness.body.read",
	//     "https://www.googleapis.com/auth/fitness.body.write",
	//     "https://www.googleapis.com/auth/fitness.location.read",
	//     "https://www.googleapis.com/auth/fitness.location.write"
	//   ]
	// }

}

// method id "fitness.users.sessions.update":

type UsersSessionsUpdateCall struct {
	s         *Service
	userId    string
	sessionId string
	session   *Session
	opt_      map[string]interface{}
}

// Update: Updates or insert a given session.
func (r *UsersSessionsService) Update(userId string, sessionId string, session *Session) *UsersSessionsUpdateCall {
	c := &UsersSessionsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.userId = userId
	c.sessionId = sessionId
	c.session = session
	return c
}

// CurrentTimeMillis sets the optional parameter "currentTimeMillis":
// The client's current time in milliseconds since epoch.
func (c *UsersSessionsUpdateCall) CurrentTimeMillis(currentTimeMillis int64) *UsersSessionsUpdateCall {
	c.opt_["currentTimeMillis"] = currentTimeMillis
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersSessionsUpdateCall) Fields(s ...googleapi.Field) *UsersSessionsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *UsersSessionsUpdateCall) IfNoneMatch(entityTag string) *UsersSessionsUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "fitness.users.sessions.update" call.
// Exactly one of the return values is non-nil.
func (c *UsersSessionsUpdateCall) Do() (*Session, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "fitness.users.sessions.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *UsersSessionsUpdateCall) DoHeader() (ret *Session, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.session)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["currentTimeMillis"]; ok {
		params.Set("currentTimeMillis", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{userId}/sessions/{sessionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"userId":    c.userId,
		"sessionId": c.sessionId,
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
	//   "description": "Updates or insert a given session.",
	//   "httpMethod": "PUT",
	//   "id": "fitness.users.sessions.update",
	//   "parameterOrder": [
	//     "userId",
	//     "sessionId"
	//   ],
	//   "parameters": {
	//     "currentTimeMillis": {
	//       "description": "The client's current time in milliseconds since epoch.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sessionId": {
	//       "description": "The ID of the session to be created.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Create sessions for the person identified. Use me to indicate the authenticated user. Only me is supported at this time.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{userId}/sessions/{sessionId}",
	//   "request": {
	//     "$ref": "Session"
	//   },
	//   "response": {
	//     "$ref": "Session"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/fitness.activity.write"
	//   ]
	// }

}
