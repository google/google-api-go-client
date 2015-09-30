// Package mapsengine provides access to the Google Maps Engine API.
//
// See https://developers.google.com/maps-engine/
//
// Usage example:
//
//   import "google.golang.org/api/mapsengine/exp2"
//   ...
//   mapsengineService, err := mapsengine.New(oauthHttpClient)
package mapsengine // import "google.golang.org/api/mapsengine/exp2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/internal"
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
var _ = internal.MarshalJSON

const apiId = "mapsengine:exp2"
const apiName = "mapsengine"
const apiVersion = "exp2"
const basePath = "https://www.googleapis.com/mapsengine/exp2/"

func urlValues() url.Values { return url.Values{} }

// OAuth2 scopes used by this API.
const (
	// View and manage your Google My Maps data
	MapsengineScope = "https://www.googleapis.com/auth/mapsengine"

	// View your Google My Maps data
	MapsengineReadonlyScope = "https://www.googleapis.com/auth/mapsengine.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Assets = NewAssetsService(s)
	s.Layers = NewLayersService(s)
	s.Maps = NewMapsService(s)
	s.Projects = NewProjectsService(s)
	s.RasterCollections = NewRasterCollectionsService(s)
	s.Rasters = NewRastersService(s)
	s.Tables = NewTablesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Assets *AssetsService

	Layers *LayersService

	Maps *MapsService

	Projects *ProjectsService

	RasterCollections *RasterCollectionsService

	Rasters *RastersService

	Tables *TablesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAssetsService(s *Service) *AssetsService {
	rs := &AssetsService{s: s}
	rs.Parents = NewAssetsParentsService(s)
	rs.Permissions = NewAssetsPermissionsService(s)
	return rs
}

type AssetsService struct {
	s *Service

	Parents *AssetsParentsService

	Permissions *AssetsPermissionsService
}

func NewAssetsParentsService(s *Service) *AssetsParentsService {
	rs := &AssetsParentsService{s: s}
	return rs
}

type AssetsParentsService struct {
	s *Service
}

func NewAssetsPermissionsService(s *Service) *AssetsPermissionsService {
	rs := &AssetsPermissionsService{s: s}
	return rs
}

type AssetsPermissionsService struct {
	s *Service
}

func NewLayersService(s *Service) *LayersService {
	rs := &LayersService{s: s}
	rs.Parents = NewLayersParentsService(s)
	rs.Permissions = NewLayersPermissionsService(s)
	return rs
}

type LayersService struct {
	s *Service

	Parents *LayersParentsService

	Permissions *LayersPermissionsService
}

func NewLayersParentsService(s *Service) *LayersParentsService {
	rs := &LayersParentsService{s: s}
	return rs
}

type LayersParentsService struct {
	s *Service
}

func NewLayersPermissionsService(s *Service) *LayersPermissionsService {
	rs := &LayersPermissionsService{s: s}
	return rs
}

type LayersPermissionsService struct {
	s *Service
}

func NewMapsService(s *Service) *MapsService {
	rs := &MapsService{s: s}
	rs.Permissions = NewMapsPermissionsService(s)
	return rs
}

type MapsService struct {
	s *Service

	Permissions *MapsPermissionsService
}

func NewMapsPermissionsService(s *Service) *MapsPermissionsService {
	rs := &MapsPermissionsService{s: s}
	return rs
}

type MapsPermissionsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Icons = NewProjectsIconsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Icons *ProjectsIconsService
}

func NewProjectsIconsService(s *Service) *ProjectsIconsService {
	rs := &ProjectsIconsService{s: s}
	return rs
}

type ProjectsIconsService struct {
	s *Service
}

func NewRasterCollectionsService(s *Service) *RasterCollectionsService {
	rs := &RasterCollectionsService{s: s}
	rs.Parents = NewRasterCollectionsParentsService(s)
	rs.Permissions = NewRasterCollectionsPermissionsService(s)
	rs.Rasters = NewRasterCollectionsRastersService(s)
	return rs
}

type RasterCollectionsService struct {
	s *Service

	Parents *RasterCollectionsParentsService

	Permissions *RasterCollectionsPermissionsService

	Rasters *RasterCollectionsRastersService
}

func NewRasterCollectionsParentsService(s *Service) *RasterCollectionsParentsService {
	rs := &RasterCollectionsParentsService{s: s}
	return rs
}

type RasterCollectionsParentsService struct {
	s *Service
}

func NewRasterCollectionsPermissionsService(s *Service) *RasterCollectionsPermissionsService {
	rs := &RasterCollectionsPermissionsService{s: s}
	return rs
}

type RasterCollectionsPermissionsService struct {
	s *Service
}

func NewRasterCollectionsRastersService(s *Service) *RasterCollectionsRastersService {
	rs := &RasterCollectionsRastersService{s: s}
	return rs
}

type RasterCollectionsRastersService struct {
	s *Service
}

func NewRastersService(s *Service) *RastersService {
	rs := &RastersService{s: s}
	rs.Files = NewRastersFilesService(s)
	rs.Parents = NewRastersParentsService(s)
	rs.Permissions = NewRastersPermissionsService(s)
	return rs
}

type RastersService struct {
	s *Service

	Files *RastersFilesService

	Parents *RastersParentsService

	Permissions *RastersPermissionsService
}

func NewRastersFilesService(s *Service) *RastersFilesService {
	rs := &RastersFilesService{s: s}
	return rs
}

type RastersFilesService struct {
	s *Service
}

func NewRastersParentsService(s *Service) *RastersParentsService {
	rs := &RastersParentsService{s: s}
	return rs
}

type RastersParentsService struct {
	s *Service
}

func NewRastersPermissionsService(s *Service) *RastersPermissionsService {
	rs := &RastersPermissionsService{s: s}
	return rs
}

type RastersPermissionsService struct {
	s *Service
}

func NewTablesService(s *Service) *TablesService {
	rs := &TablesService{s: s}
	rs.Features = NewTablesFeaturesService(s)
	rs.Files = NewTablesFilesService(s)
	rs.Parents = NewTablesParentsService(s)
	rs.Permissions = NewTablesPermissionsService(s)
	return rs
}

type TablesService struct {
	s *Service

	Features *TablesFeaturesService

	Files *TablesFilesService

	Parents *TablesParentsService

	Permissions *TablesPermissionsService
}

func NewTablesFeaturesService(s *Service) *TablesFeaturesService {
	rs := &TablesFeaturesService{s: s}
	return rs
}

type TablesFeaturesService struct {
	s *Service
}

func NewTablesFilesService(s *Service) *TablesFilesService {
	rs := &TablesFilesService{s: s}
	return rs
}

type TablesFilesService struct {
	s *Service
}

func NewTablesParentsService(s *Service) *TablesParentsService {
	rs := &TablesParentsService{s: s}
	return rs
}

type TablesParentsService struct {
	s *Service
}

func NewTablesPermissionsService(s *Service) *TablesPermissionsService {
	rs := &TablesPermissionsService{s: s}
	return rs
}

type TablesPermissionsService struct {
	s *Service
}

// AcquisitionTime: Acquisition time represents acquired time of a
// raster.
type AcquisitionTime struct {
	// End: The end time if acquisition time is a range. The value is an RFC
	// 3339 formatted date-time value (1970-01-01T00:00:00Z).
	End string `json:"end,omitempty"`

	// Precision: The precision of acquisition time.
	//
	// Possible values:
	//   "day"
	//   "hour"
	//   "minute"
	//   "month"
	//   "second"
	//   "year"
	Precision string `json:"precision,omitempty"`

	// Start: The acquisition time, or start time if acquisition time is a
	// range. The value is an RFC 3339 formatted date-time value
	// (1970-01-01T00:00:00Z).
	Start string `json:"start,omitempty"`

	// ForceSendFields is a list of field names (e.g. "End") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AcquisitionTime) MarshalJSON() ([]byte, error) {
	type noMethod AcquisitionTime
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Asset: An asset is any Google Maps Engine resource that has a
// globally unique ID. Assets include maps, layers, vector tables,
// raster collections, and rasters. Projects and features are not
// considered assets.
//
// More detailed information about an asset can be obtained by querying
// the asset's particular endpoint.
type Asset struct {
	// Bbox: A rectangular bounding box which contains all of the data in
	// this asset. The box is expressed as \"west, south, east, north\". The
	// numbers represent latitude and longitude in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this asset. The value is an RFC
	// 3339-formatted date-time value (for example, 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this asset. This is
	// only returned on GET requests and not LIST requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Description: The asset's description.
	Description string `json:"description,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Id: The asset's globally unique ID.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this asset. The value is
	// an RFC 3339-formatted date-time value (for example,
	// 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// asset. This is only returned on GET requests and not LIST requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// Name: The asset's name.
	Name string `json:"name,omitempty"`

	// ProjectId: The ID of the project to which the asset belongs.
	ProjectId string `json:"projectId,omitempty"`

	// Resource: The URL to query to retrieve the asset's complete object.
	// The assets endpoint only returns high-level information about a
	// resource.
	Resource string `json:"resource,omitempty"`

	// Tags: An array of text strings, with each string representing a tag.
	// More information about tags can be found in the Tagging data article
	// of the Maps Engine help center.
	Tags []string `json:"tags,omitempty"`

	// Type: The type of asset. One of raster, rasterCollection, table, map,
	// or layer.
	//
	// Possible values:
	//   "layer"
	//   "map"
	//   "raster"
	//   "rasterCollection"
	//   "table"
	Type string `json:"type,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bbox") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Asset) MarshalJSON() ([]byte, error) {
	type noMethod Asset
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// AssetsListResponse: The response returned by a call to
// resources.List.
type AssetsListResponse struct {
	// Assets: Assets returned.
	Assets []*Asset `json:"assets,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Assets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AssetsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod AssetsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Border: Border in line style. Both color and width are required.
type Border struct {
	// Color: Color of the border.
	Color string `json:"color,omitempty"`

	// Opacity: Opacity of the border.
	Opacity float64 `json:"opacity,omitempty"`

	// Width: Width of the border, in pixels.
	Width float64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Color") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Border) MarshalJSON() ([]byte, error) {
	type noMethod Border
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Color: Basic color used in styling.
type Color struct {
	// Color: The CSS style color, can be in format of "red" or "#7733EE".
	Color string `json:"color,omitempty"`

	// Opacity: Opacity ranges from 0 to 1, inclusive. If not provided,
	// default to 1.
	Opacity float64 `json:"opacity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Color") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Color) MarshalJSON() ([]byte, error) {
	type noMethod Color
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type Datasource struct {
	// Id: The ID of a datasource.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Datasource) MarshalJSON() ([]byte, error) {
	type noMethod Datasource
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// DisplayRule: A display rule of the vector style.
type DisplayRule struct {
	// Filters: This display rule will only be applied to features that
	// match all of the filters here. If filters is empty, then the rule
	// applies to all features.
	Filters []*Filter `json:"filters,omitempty"`

	// LineOptions: Style applied to lines. Required for LineString
	// Geometry.
	LineOptions *LineStyle `json:"lineOptions,omitempty"`

	// Name: Display rule name. Name is not unique and cannot be used for
	// identification purpose.
	Name string `json:"name,omitempty"`

	// PointOptions: Style applied to points. Required for Point Geometry.
	PointOptions *PointStyle `json:"pointOptions,omitempty"`

	// PolygonOptions: Style applied to polygons. Required for Polygon
	// Geometry.
	PolygonOptions *PolygonStyle `json:"polygonOptions,omitempty"`

	// ZoomLevels: The zoom levels that this display rule apply.
	ZoomLevels *ZoomLevels `json:"zoomLevels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DisplayRule) MarshalJSON() ([]byte, error) {
	type noMethod DisplayRule
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Feature: A feature within a table.
type Feature struct {
	// Geometry: The geometry member of this Feature.
	Geometry GeoJsonGeometry `json:"geometry,omitempty"`

	// Properties: Key/value pairs of this Feature.
	Properties GeoJsonProperties `json:"properties,omitempty"`

	// Type: Identifies this object as a feature.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Geometry") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Feature) MarshalJSON() ([]byte, error) {
	type noMethod Feature
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FeatureInfo: A feature info contains information about individual
// feature.
type FeatureInfo struct {
	// Content: HTML template of the info window. If not provided, a default
	// template with all attributes will be generated.
	Content string `json:"content,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Content") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FeatureInfo) MarshalJSON() ([]byte, error) {
	type noMethod FeatureInfo
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FeaturesBatchDeleteRequest: The request sent to features.BatchDelete.
type FeaturesBatchDeleteRequest struct {
	GxIds []string `json:"gx_ids,omitempty"`

	PrimaryKeys []string `json:"primaryKeys,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GxIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FeaturesBatchDeleteRequest) MarshalJSON() ([]byte, error) {
	type noMethod FeaturesBatchDeleteRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FeaturesBatchInsertRequest: The request sent to features.Insert.
type FeaturesBatchInsertRequest struct {
	Features []*Feature `json:"features,omitempty"`

	// NormalizeGeometries: If true, the server will normalize feature
	// geometries. It is assumed that the South Pole is exterior to any
	// polygons given. See here for a list of normalizations. If false, all
	// feature geometries must be given already normalized. The points in
	// all LinearRings must be listed in counter-clockwise order, and
	// LinearRings may not intersect.
	//
	// Default: true
	NormalizeGeometries *bool `json:"normalizeGeometries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Features") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FeaturesBatchInsertRequest) MarshalJSON() ([]byte, error) {
	type noMethod FeaturesBatchInsertRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FeaturesBatchPatchRequest: The request sent to features.BatchPatch.
type FeaturesBatchPatchRequest struct {
	Features []*Feature `json:"features,omitempty"`

	// NormalizeGeometries: If true, the server will normalize feature
	// geometries. It is assumed that the South Pole is exterior to any
	// polygons given. See here for a list of normalizations. If false, all
	// feature geometries must be given already normalized. The points in
	// all LinearRings must be listed in counter-clockwise order, and
	// LinearRings may not intersect.
	//
	// Default: true
	NormalizeGeometries *bool `json:"normalizeGeometries,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Features") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FeaturesBatchPatchRequest) MarshalJSON() ([]byte, error) {
	type noMethod FeaturesBatchPatchRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// FeaturesListResponse: The response returned by a call to
// features.List.
type FeaturesListResponse struct {
	// AllowedQueriesPerSecond: An indicator of the maximum rate at which
	// queries may be made, if all queries were as expensive as this query.
	AllowedQueriesPerSecond float64 `json:"allowedQueriesPerSecond,omitempty"`

	// Features: Resources returned.
	Features []*Feature `json:"features,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Schema: The feature schema.
	Schema *Schema `json:"schema,omitempty"`

	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AllowedQueriesPerSecond") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FeaturesListResponse) MarshalJSON() ([]byte, error) {
	type noMethod FeaturesListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// File: A single File, which is a component of an Asset.
type File struct {
	// Filename: The name of the file.
	Filename string `json:"filename,omitempty"`

	// Size: The size of the file in bytes.
	Size int64 `json:"size,omitempty,string"`

	// UploadStatus: The upload status of the file.
	//
	// Possible values:
	//   "canceled"
	//   "complete"
	//   "failed"
	//   "inProgress"
	UploadStatus string `json:"uploadStatus,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filename") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *File) MarshalJSON() ([]byte, error) {
	type noMethod File
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Filter: Conditions for filtering features.
type Filter struct {
	// Column: The column name to filter on.
	Column string `json:"column,omitempty"`

	// Operator: Operation used to evaluate the filter.
	//
	// Possible values:
	//   "!="
	//   "<"
	//   "<="
	//   "=="
	//   ">"
	//   ">="
	//   "contains"
	//   "endsWith"
	//   "startsWith"
	Operator string `json:"operator,omitempty"`

	// Value: Value to be evaluated against attribute.
	Value interface{} `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Column") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Filter) MarshalJSON() ([]byte, error) {
	type noMethod Filter
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonGeometry map[string]interface{}

func (t GeoJsonGeometry) Type() string {
	return googleapi.VariantType(t)
}

func (t GeoJsonGeometry) GeometryCollection() (r GeoJsonGeometryCollection, ok bool) {
	if t.Type() != "GeometryCollection" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) LineString() (r GeoJsonLineString, ok bool) {
	if t.Type() != "LineString" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) MultiLineString() (r GeoJsonMultiLineString, ok bool) {
	if t.Type() != "MultiLineString" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) MultiPoint() (r GeoJsonMultiPoint, ok bool) {
	if t.Type() != "MultiPoint" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) MultiPolygon() (r GeoJsonMultiPolygon, ok bool) {
	if t.Type() != "MultiPolygon" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) Point() (r GeoJsonPoint, ok bool) {
	if t.Type() != "Point" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t GeoJsonGeometry) Polygon() (r GeoJsonPolygon, ok bool) {
	if t.Type() != "Polygon" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

// GeoJsonGeometryCollection: A heterogenous collection of
// GeoJsonGeometry objects.
type GeoJsonGeometryCollection struct {
	// Geometries: An array of geometry objects. There must be at least 2
	// different types of geometries in the array.
	Geometries []GeoJsonGeometry `json:"geometries,omitempty"`

	// Type: Identifies this object as a GeoJsonGeometryCollection.
	//
	// Possible values:
	//   "GeometryCollection"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Geometries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonGeometryCollection) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonGeometryCollection
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonLineString struct {
	// Coordinates: An array of two or more positions, representing a line.
	Coordinates [][]float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonLineString.
	//
	// Possible values:
	//   "LineString"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonLineString) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonLineString
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// GeoJsonMultiLineString: Multi Line String
type GeoJsonMultiLineString struct {
	// Coordinates: An array of at least two GeoJsonLineString coordinate
	// arrays.
	Coordinates [][][]float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonMultiLineString.
	//
	// Possible values:
	//   "MultiLineString"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonMultiLineString) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonMultiLineString
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonMultiPoint struct {
	// Coordinates: An array of at least two GeoJsonPoint coordinate arrays.
	Coordinates [][]float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonMultiPoint.
	//
	// Possible values:
	//   "MultiPoint"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonMultiPoint) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonMultiPoint
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonMultiPolygon struct {
	// Coordinates: An array of at least two GeoJsonPolygon coordinate
	// arrays.
	Coordinates [][][][]float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonMultiPolygon.
	//
	// Possible values:
	//   "MultiPolygon"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonMultiPolygon) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonMultiPolygon
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonPoint struct {
	// Coordinates: A single GeoJsonPosition, specifying the location of the
	// point.
	Coordinates []float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonPoint.
	//
	// Possible values:
	//   "Point"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonPoint) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonPoint
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonPolygon struct {
	// Coordinates: An array of LinearRings. A LinearRing is a
	// GeoJsonLineString which is closed (that is, the first and last
	// GeoJsonPositions are equal), and which contains at least four
	// GeoJsonPositions. For polygons with multiple rings, the first
	// LinearRing is the exterior ring, and any subsequent rings are
	// interior rings (that is, holes).
	Coordinates [][][]float64 `json:"coordinates,omitempty"`

	// Type: Identifies this object as a GeoJsonPolygon.
	//
	// Possible values:
	//   "Polygon"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Coordinates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoJsonPolygon) MarshalJSON() ([]byte, error) {
	type noMethod GeoJsonPolygon
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type GeoJsonProperties interface{}

// Icon: An icon is a user-uploaded image that can be used to style
// point geometries.
type Icon struct {
	// Description: The description of this Icon, supplied by the author.
	Description string `json:"description,omitempty"`

	// Id: An ID used to refer to this Icon.
	Id string `json:"id,omitempty"`

	// Name: The name of this Icon, supplied by the author.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Icon) MarshalJSON() ([]byte, error) {
	type noMethod Icon
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// IconStyle: Style for icon, this is part of point style.
type IconStyle struct {
	// Id: Custom icon id.
	Id string `json:"id,omitempty"`

	// Name: Stock icon name. To use a stock icon, prefix it with 'gx_'. See
	// Stock icon names for valid icon names. For example, to specify
	// small_red, set name to 'gx_small_red'.
	Name string `json:"name,omitempty"`

	// ScaledShape: A scalable shape.
	ScaledShape *ScaledShape `json:"scaledShape,omitempty"`

	// ScalingFunction: The function used to scale shapes. Required when a
	// scaledShape is specified.
	ScalingFunction *ScalingFunction `json:"scalingFunction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *IconStyle) MarshalJSON() ([]byte, error) {
	type noMethod IconStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// IconsListResponse: The response returned by a call to icons.List.
type IconsListResponse struct {
	// Icons: Resources returned.
	Icons []*Icon `json:"icons,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Icons") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *IconsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod IconsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// LabelStyle: Text label style.
type LabelStyle struct {
	// Color: Color of the text. If not provided, default to black.
	Color string `json:"color,omitempty"`

	// Column: The column value of the feature to be displayed.
	Column string `json:"column,omitempty"`

	// FontStyle: Font style of the label, defaults to 'normal'.
	//
	// Possible values:
	//   "italic"
	//   "normal"
	FontStyle string `json:"fontStyle,omitempty"`

	// FontWeight: Font weight of the label, defaults to 'normal'.
	//
	// Possible values:
	//   "bold"
	//   "normal"
	FontWeight string `json:"fontWeight,omitempty"`

	// Opacity: Opacity of the text.
	Opacity float64 `json:"opacity,omitempty"`

	// Outline: Outline color of the text.
	Outline *Color `json:"outline,omitempty"`

	// Size: Font size of the label, in pixels. 8 <= size <= 15. If not
	// provided, a default size will be provided.
	Size float64 `json:"size,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Color") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LabelStyle) MarshalJSON() ([]byte, error) {
	type noMethod LabelStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Layer: A Layer combines multiple datasources, with styling
// information, for presentation on a map.
type Layer struct {
	// Bbox: A rectangular bounding box which contains all of the data in
	// this Layer. The box is expressed as \"west, south, east, north\". The
	// numbers represent latitude and longitude in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this layer. The value is an RFC
	// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this layer. This is
	// only returned on GET requests and not LIST requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// DatasourceType: Deprecated: The type of the datasources used to build
	// this Layer. Note: This has been replaced by layerType, but is still
	// available for now to maintain backward compatibility.
	//
	// Possible values:
	//   "image"
	//   "table"
	DatasourceType string `json:"datasourceType,omitempty"`

	// Datasources: An array of datasources used to build this layer. If
	// layerType is "image", or layerType is not specified and
	// datasourceType is "image", then each element in this array is a
	// reference to an Image or RasterCollection. If layerType is "vector",
	// or layerType is not specified and datasourceType is "table" then each
	// element in this array is a reference to a Vector Table. Note:
	// Datasources are returned in response to a get request but not a list
	// request. After requesting a list of layers, you'll need to send a get
	// request to retrieve the datasources for each layer.
	Datasources []*Datasource `json:"datasources,omitempty"`

	// Description: The description of this Layer, supplied by the author.
	Description string `json:"description,omitempty"`

	// DraftAccessList: Deprecated: The name of an access list of the Map
	// Editor type. The user on whose behalf the request is being sent must
	// be an editor on that access list. Note: Google Maps Engine no longer
	// uses access lists. Instead, each asset has its own list of
	// permissions. For backward compatibility, the API still accepts access
	// lists for projects that are already using access lists. If you
	// created a GME account/project after July 14th, 2014, you will not be
	// able to send API requests that include access lists. Note: This is an
	// input field only. It is not returned in response to a list or get
	// request.
	DraftAccessList string `json:"draftAccessList,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Id: A globally unique ID, used to refer to this Layer.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this layer. The value is
	// an RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// layer. This is only returned on GET requests and not LIST requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// LayerType: The type of the datasources used to build this Layer. This
	// should be used instead of datasourceType. At least one of layerType
	// and datasourceType and must be specified, but layerType takes
	// precedence.
	//
	// Possible values:
	//   "image"
	//   "vector"
	LayerType string `json:"layerType,omitempty"`

	// Name: The name of this Layer, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProcessingStatus: The processing status of this layer.
	//
	// Possible values:
	//   "complete"
	//   "failed"
	//   "notReady"
	//   "processing"
	//   "ready"
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// ProjectId: The ID of the project that this Layer is in.
	ProjectId string `json:"projectId,omitempty"`

	// PublishedAccessList: Deprecated: The access list to whom view
	// permissions are granted. The value must be the name of a Maps Engine
	// access list of the Map Viewer type, and the user must be a viewer on
	// that list. Note: Google Maps Engine no longer uses access lists.
	// Instead, each asset has its own list of permissions. For backward
	// compatibility, the API still accepts access lists for projects that
	// are already using access lists. If you created a GME account/project
	// after July 14th, 2014, you will not be able to send API requests that
	// include access lists. Note: This is an input field only. It is not
	// returned in response to a list or get request.
	PublishedAccessList string `json:"publishedAccessList,omitempty"`

	// PublishingStatus: The publishing status of this layer.
	//
	// Possible values:
	//   "notPublished"
	//   "published"
	PublishingStatus string `json:"publishingStatus,omitempty"`

	// Style: The styling information for a vector layer. Note: Style
	// information is returned in response to a get request but not a list
	// request. After requesting a list of layers, you'll need to send a get
	// request to retrieve the VectorStyles for each layer.
	Style *VectorStyle `json:"style,omitempty"`

	// Tags: Tags of this Layer.
	Tags []string `json:"tags,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bbox") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Layer) MarshalJSON() ([]byte, error) {
	type noMethod Layer
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// LayersListResponse: The response returned by a call to layers.List.
// Note: The list response does not include all the fields available in
// a layer. Refer to the layer resource description for details of the
// fields that are not included. You'll need to send a get request to
// retrieve the additional fields for each layer.
type LayersListResponse struct {
	// Layers: Resources returned.
	Layers []*Layer `json:"layers,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Layers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LayersListResponse) MarshalJSON() ([]byte, error) {
	type noMethod LayersListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// LineStyle: Style for lines.
type LineStyle struct {
	// Border: Border of the line. 0 < border.width <= 5.
	Border *Border `json:"border,omitempty"`

	// Dash: Dash defines the pattern of the line, the values are pixel
	// lengths of alternating dash and gap. If dash is not provided, then it
	// means a solid line. Dash can contain up to 10 values and must contain
	// even number of values.
	Dash []float64 `json:"dash,omitempty"`

	// Label: Label style for the line.
	Label *LabelStyle `json:"label,omitempty"`

	// Stroke: Stroke of the line.
	Stroke *LineStyleStroke `json:"stroke,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Border") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LineStyle) MarshalJSON() ([]byte, error) {
	type noMethod LineStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// LineStyleStroke: Stroke of the line.
type LineStyleStroke struct {
	// Color: Color of the line.
	Color string `json:"color,omitempty"`

	// Opacity: Opacity of the line.
	Opacity float64 `json:"opacity,omitempty"`

	// Width: Width of the line, in pixels. 0 <= width <= 10. If width is
	// set to 0, the line will be invisible.
	Width float64 `json:"width,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Color") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LineStyleStroke) MarshalJSON() ([]byte, error) {
	type noMethod LineStyleStroke
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Map: A Map is a collection of Layers, optionally contained within
// folders.
type Map struct {
	// Bbox: A rectangular bounding box which contains all of the data in
	// this Map. The box is expressed as \"west, south, east, north\". The
	// numbers represent latitude and longitude in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// Contents: The contents of this Map.
	Contents []MapItem `json:"contents,omitempty"`

	// CreationTime: The creation time of this map. The value is an RFC 3339
	// formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this map. This is
	// only returned on GET requests and not LIST requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// DefaultViewport: An array of four numbers (west, south, east, north)
	// which defines the rectangular bounding box of the default viewport.
	// The numbers represent latitude and longitude in decimal degrees.
	DefaultViewport []float64 `json:"defaultViewport,omitempty"`

	// Description: The description of this Map, supplied by the author.
	Description string `json:"description,omitempty"`

	// DraftAccessList: Deprecated: The name of an access list of the Map
	// Editor type. The user on whose behalf the request is being sent must
	// be an editor on that access list. Note: Google Maps Engine no longer
	// uses access lists. Instead, each asset has its own list of
	// permissions. For backward compatibility, the API still accepts access
	// lists for projects that are already using access lists. If you
	// created a GME account/project after July 14th, 2014, you will not be
	// able to send API requests that include access lists. Note: This is an
	// input field only. It is not returned in response to a list or get
	// request.
	DraftAccessList string `json:"draftAccessList,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Id: A globally unique ID, used to refer to this Map.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this map. The value is an
	// RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// map. This is only returned on GET requests and not LIST requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// Name: The name of this Map, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProcessingStatus: The processing status of this map. Map processing
	// is automatically started once a map becomes ready for processing.
	//
	// Possible values:
	//   "complete"
	//   "failed"
	//   "notReady"
	//   "processing"
	//   "ready"
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// ProjectId: The ID of the project that this Map is in.
	ProjectId string `json:"projectId,omitempty"`

	// PublishedAccessList: Deprecated: The access list to whom view
	// permissions are granted. The value must be the name of a Maps Engine
	// access list of the Map Viewer type, and the user must be a viewer on
	// that list. Note: Google Maps Engine no longer uses access lists.
	// Instead, each asset has its own list of permissions. For backward
	// compatibility, the API still accepts access lists for projects that
	// are already using access lists. If you created a GME account/project
	// after July 14th, 2014, you will not be able to send API requests that
	// include access lists. This is an input field only. It is not returned
	// in response to a list or get request.
	PublishedAccessList string `json:"publishedAccessList,omitempty"`

	// PublishingStatus: The publishing status of this map.
	//
	// Possible values:
	//   "notPublished"
	//   "published"
	PublishingStatus string `json:"publishingStatus,omitempty"`

	// Tags: Tags of this Map.
	Tags []string `json:"tags,omitempty"`

	// Versions: Deprecated: An array containing the available versions of
	// this Map. Currently may only contain "published". The
	// publishingStatus field should be used instead.
	Versions []string `json:"versions,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bbox") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Map) MarshalJSON() ([]byte, error) {
	type noMethod Map
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type MapFolder struct {
	Contents []MapItem `json:"contents,omitempty"`

	// DefaultViewport: An array of four numbers (west, south, east, north)
	// which defines the rectangular bounding box of the default viewport.
	// The numbers represent latitude and longitude in decimal degrees.
	DefaultViewport []float64 `json:"defaultViewport,omitempty"`

	// Expandable: The expandability setting of this MapFolder. If true, the
	// folder can be expanded.
	Expandable bool `json:"expandable,omitempty"`

	// Key: A user defined alias for this MapFolder, specific to this Map.
	Key string `json:"key,omitempty"`

	// Name: The name of this MapFolder.
	Name string `json:"name,omitempty"`

	// Type: Identifies this object as a MapFolder.
	//
	// Possible values:
	//   "folder"
	Type string `json:"type,omitempty"`

	// Visibility: The visibility setting of this MapFolder. One of
	// "defaultOn" or "defaultOff".
	Visibility string `json:"visibility,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MapFolder) MarshalJSON() ([]byte, error) {
	type noMethod MapFolder
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type MapItem map[string]interface{}

func (t MapItem) Type() string {
	return googleapi.VariantType(t)
}

func (t MapItem) Folder() (r MapFolder, ok bool) {
	if t.Type() != "Folder" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t MapItem) KmlLink() (r MapKmlLink, ok bool) {
	if t.Type() != "KmlLink" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

func (t MapItem) Layer() (r MapLayer, ok bool) {
	if t.Type() != "Layer" {
		return r, false
	}
	ok = googleapi.ConvertVariant(map[string]interface{}(t), &r)
	return r, ok
}

type MapKmlLink struct {
	// DefaultViewport: An array of four numbers (west, south, east, north)
	// which defines the rectangular bounding box of the default viewport.
	// The numbers represent latitude and longitude in decimal degrees.
	DefaultViewport []float64 `json:"defaultViewport,omitempty"`

	// KmlUrl: The URL to the KML file represented by this MapKmlLink.
	KmlUrl string `json:"kmlUrl,omitempty"`

	// Name: The name of this MapKmlLink.
	Name string `json:"name,omitempty"`

	// Type: Identifies this object as a MapKmlLink.
	//
	// Possible values:
	//   "kmlLink"
	Type string `json:"type,omitempty"`

	// Visibility: The visibility setting of this MapKmlLink. One of
	// "defaultOn" or "defaultOff".
	Visibility string `json:"visibility,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultViewport") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MapKmlLink) MarshalJSON() ([]byte, error) {
	type noMethod MapKmlLink
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type MapLayer struct {
	// DefaultViewport: An array of four numbers (west, south, east, north)
	// which defines the rectangular bounding box of the default viewport.
	// The numbers represent latitude and longitude in decimal degrees.
	DefaultViewport []float64 `json:"defaultViewport,omitempty"`

	// Id: The ID of this MapLayer. This ID can be used to request more
	// details about the layer.
	Id string `json:"id,omitempty"`

	// Key: A user defined alias for this MapLayer, specific to this Map.
	Key string `json:"key,omitempty"`

	// Name: The name of this MapLayer.
	Name string `json:"name,omitempty"`

	// Type: Identifies this object as a MapLayer.
	//
	// Possible values:
	//   "layer"
	Type string `json:"type,omitempty"`

	// Visibility: The visibility setting of this MapLayer. One of
	// "defaultOn" or "defaultOff".
	Visibility string `json:"visibility,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DefaultViewport") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MapLayer) MarshalJSON() ([]byte, error) {
	type noMethod MapLayer
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// MapsListResponse: The response returned by a call to maps.List.
type MapsListResponse struct {
	// Maps: Resources returned.
	Maps []*Map `json:"maps,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Maps") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MapsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod MapsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Parent: A list of the parents of an asset.
type Parent struct {
	// Id: The ID of this parent.
	Id string `json:"id,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Parent) MarshalJSON() ([]byte, error) {
	type noMethod Parent
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ParentsListResponse: The response returned by a call to parents.List.
type ParentsListResponse struct {
	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Parents: The parent assets.
	Parents []*Parent `json:"parents,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ParentsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ParentsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Permission: A permission defines the user or group that has access to
// an asset, and the type of access they have.
type Permission struct {
	// Discoverable: Indicates whether a public asset is listed and can be
	// found via a web search (value true), or is visible only to people who
	// have a link to the asset (value false).
	Discoverable bool `json:"discoverable,omitempty"`

	// Id: The unique identifier of the permission. This could be the email
	// address of the user or group this permission refers to, or the string
	// "anyone" for public permissions.
	Id string `json:"id,omitempty"`

	// Role: The type of access granted to this user or group.
	//
	// Possible values:
	//   "owner"
	//   "reader"
	//   "viewer"
	//   "writer"
	Role string `json:"role,omitempty"`

	// Type: The account type.
	//
	// Possible values:
	//   "anyone"
	//   "group"
	//   "user"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Discoverable") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Permission) MarshalJSON() ([]byte, error) {
	type noMethod Permission
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PermissionsBatchDeleteRequest: The request sent to
// mapsengine.permissions.batchDelete.
type PermissionsBatchDeleteRequest struct {
	// Ids: An array of permission ids to be removed. This could be the
	// email address of the user or group this permission refers to, or the
	// string "anyone" for public permissions.
	Ids []string `json:"ids,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Ids") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PermissionsBatchDeleteRequest) MarshalJSON() ([]byte, error) {
	type noMethod PermissionsBatchDeleteRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PermissionsBatchDeleteResponse: The response returned by a call to
// mapsengine.permissions.batchDelete.
type PermissionsBatchDeleteResponse struct {
}

// PermissionsBatchUpdateRequest: The request sent to
// mapsengine.permissions.batchUpdate.
type PermissionsBatchUpdateRequest struct {
	// Permissions: The permissions to be inserted or updated.
	Permissions []*Permission `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PermissionsBatchUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod PermissionsBatchUpdateRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PermissionsBatchUpdateResponse: The response returned by a call to
// mapsengine.permissions.batchUpdate.
type PermissionsBatchUpdateResponse struct {
}

type PermissionsListResponse struct {
	// Permissions: The set of permissions associated with this asset.
	Permissions []*Permission `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PermissionsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod PermissionsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PointStyle: Style for points.
type PointStyle struct {
	// Icon: Icon for the point; if it isn't null, exactly one of 'name',
	// 'id' or 'scaledShape' must be set.
	Icon *IconStyle `json:"icon,omitempty"`

	// Label: Label style for the point.
	Label *LabelStyle `json:"label,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Icon") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PointStyle) MarshalJSON() ([]byte, error) {
	type noMethod PointStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PolygonStyle: Style for polygons.
type PolygonStyle struct {
	// Fill: Fill color of the polygon. If not provided, the polygon will be
	// transparent and not visible if there is no border.
	Fill *Color `json:"fill,omitempty"`

	// Label: Label style for the polygon.
	Label *LabelStyle `json:"label,omitempty"`

	// Stroke: Border of the polygon. 0 < border.width <= 10.
	Stroke *Border `json:"stroke,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fill") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PolygonStyle) MarshalJSON() ([]byte, error) {
	type noMethod PolygonStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ProcessResponse: The response returned by a call to any asset's
// Process method.
type ProcessResponse struct {
}

// Project: A project groups a collection of resources.
type Project struct {
	// Id: An ID used to refer to this project.
	Id string `json:"id,omitempty"`

	// Name: A user provided name for this project.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Project) MarshalJSON() ([]byte, error) {
	type noMethod Project
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ProjectsListResponse: The response returned by a call to
// projects.List.
type ProjectsListResponse struct {
	// Projects: Projects returned.
	Projects []*Project `json:"projects,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Projects") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ProjectsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod ProjectsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PublishResponse: The response returned by a call to any asset's
// Publish method.
type PublishResponse struct {
}

// PublishedLayer: The published version of a layer.
type PublishedLayer struct {
	// Description: The description of this Layer, supplied by the author.
	Description string `json:"description,omitempty"`

	// Id: A globally unique ID, used to refer to this Layer.
	Id string `json:"id,omitempty"`

	// LayerType: The type of the datasources used to build this Layer. This
	// should be used instead of datasourceType. At least one of layerType
	// and datasourceType and must be specified, but layerType takes
	// precedence.
	//
	// Possible values:
	//   "image"
	//   "vector"
	LayerType string `json:"layerType,omitempty"`

	// Name: The name of this Layer, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProjectId: The ID of the project that this Layer is in.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PublishedLayer) MarshalJSON() ([]byte, error) {
	type noMethod PublishedLayer
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PublishedLayersListResponse: The response returned by a call to
// layers.List.published.
type PublishedLayersListResponse struct {
	// Layers: Resources returned.
	Layers []*PublishedLayer `json:"layers,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Layers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PublishedLayersListResponse) MarshalJSON() ([]byte, error) {
	type noMethod PublishedLayersListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PublishedMap: The published version of a map asset.
type PublishedMap struct {
	// Contents: The contents of this Map.
	Contents []MapItem `json:"contents,omitempty"`

	// DefaultViewport: An array of four numbers (west, south, east, north)
	// which defines the rectangular bounding box of the default viewport.
	// The numbers represent latitude and longitude in decimal degrees.
	DefaultViewport []float64 `json:"defaultViewport,omitempty"`

	// Description: The description of this Map, supplied by the author.
	Description string `json:"description,omitempty"`

	// Id: A globally unique ID, used to refer to this Map.
	Id string `json:"id,omitempty"`

	// Name: The name of this Map, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProjectId: The ID of the project that this Map is in.
	ProjectId string `json:"projectId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contents") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PublishedMap) MarshalJSON() ([]byte, error) {
	type noMethod PublishedMap
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// PublishedMapsListResponse: The response returned by a call to
// maps.List.published.
type PublishedMapsListResponse struct {
	// Maps: Resources returned.
	Maps []*PublishedMap `json:"maps,omitempty"`

	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Maps") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PublishedMapsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod PublishedMapsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Raster: A geo-referenced raster.
type Raster struct {
	// AcquisitionTime: The acquisition time of this Raster.
	AcquisitionTime *AcquisitionTime `json:"acquisitionTime,omitempty"`

	// Attribution: The name of the attribution to be used for this Raster.
	Attribution string `json:"attribution,omitempty"`

	// Bbox: A rectangular bounding box which contains all of the data in
	// this Raster. The box is expressed as \"west, south, east, north\".
	// The numbers represent latitudes and longitudes in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this raster. The value is an RFC
	// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this raster. This
	// is only returned on GET requests and not LIST requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Description: The description of this Raster, supplied by the author.
	Description string `json:"description,omitempty"`

	// DraftAccessList: Deprecated: The name of an access list of the Map
	// Editor type. The user on whose behalf the request is being sent must
	// be an editor on that access list. Note: Google Maps Engine no longer
	// uses access lists. Instead, each asset has its own list of
	// permissions. For backward compatibility, the API still accepts access
	// lists for projects that are already using access lists. If you
	// created a GME account/project after July 14th, 2014, you will not be
	// able to send API requests that include access lists. Note: This is an
	// input field only. It is not returned in response to a list or get
	// request.
	DraftAccessList string `json:"draftAccessList,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Files: The files associated with this Raster.
	Files []*File `json:"files,omitempty"`

	// Id: A globally unique ID, used to refer to this Raster.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this raster. The value is
	// an RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// raster. This is only returned on GET requests and not LIST requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// MaskType: The mask processing type of this Raster.
	MaskType string `json:"maskType,omitempty"`

	// Name: The name of this Raster, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProcessingStatus: The processing status of this Raster.
	//
	// Possible values:
	//   "complete"
	//   "failed"
	//   "notReady"
	//   "processing"
	//   "ready"
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// ProjectId: The ID of the project that this Raster is in.
	ProjectId string `json:"projectId,omitempty"`

	// RasterType: The type of this Raster. Always "image" today.
	//
	// Possible values:
	//   "image"
	RasterType string `json:"rasterType,omitempty"`

	// Tags: Tags of this Raster.
	Tags []string `json:"tags,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AcquisitionTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Raster) MarshalJSON() ([]byte, error) {
	type noMethod Raster
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollection: A raster collection groups multiple Raster
// resources for inclusion in a Layer.
type RasterCollection struct {
	// Attribution: The name of the attribution to be used for this
	// RasterCollection. Note: Attribution is returned in response to a get
	// request but not a list request. After requesting a list of raster
	// collections, you'll need to send a get request to retrieve the
	// attribution for each raster collection.
	Attribution string `json:"attribution,omitempty"`

	// Bbox: A rectangular bounding box which contains all of the data in
	// this RasterCollection. The box is expressed as \"west, south, east,
	// north\". The numbers represent latitude and longitude in decimal
	// degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this RasterCollection. The value
	// is an RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this raster
	// collection. This is only returned on GET requests and not LIST
	// requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Description: The description of this RasterCollection, supplied by
	// the author.
	Description string `json:"description,omitempty"`

	// DraftAccessList: Deprecated: The name of an access list of the Map
	// Editor type. The user on whose behalf the request is being sent must
	// be an editor on that access list. Note: Google Maps Engine no longer
	// uses access lists. Instead, each asset has its own list of
	// permissions. For backward compatibility, the API still accepts access
	// lists for projects that are already using access lists. If you
	// created a GME account/project after July 14th, 2014, you will not be
	// able to send API requests that include access lists. Note: This is an
	// input field only. It is not returned in response to a list or get
	// request.
	DraftAccessList string `json:"draftAccessList,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Id: A globally unique ID, used to refer to this RasterCollection.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this RasterCollection.
	// The value is an RFC 3339 formatted date-time value (e.g.
	// 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// raster collection. This is only returned on GET requests and not LIST
	// requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// Mosaic: True if this RasterCollection is a mosaic.
	Mosaic bool `json:"mosaic,omitempty"`

	// Name: The name of this RasterCollection, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProcessingStatus: The processing status of this RasterCollection.
	//
	// Possible values:
	//   "complete"
	//   "failed"
	//   "notReady"
	//   "processing"
	//   "ready"
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// ProjectId: The ID of the project that this RasterCollection is in.
	ProjectId string `json:"projectId,omitempty"`

	// RasterType: The type of rasters contained within this
	// RasterCollection.
	//
	// Possible values:
	//   "image"
	RasterType string `json:"rasterType,omitempty"`

	// Tags: Tags of this RasterCollection.
	Tags []string `json:"tags,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Attribution") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollection) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollection
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollectionsListResponse: The response returned by a call to
// raster_collections.List. Note: The list response does not include all
// the fields available in a raster collection. Refer to the
// RasterCollection resource description for details of the fields that
// are not included. You'll need to send a get request to retrieve the
// additional fields for each raster collection.
type RasterCollectionsListResponse struct {
	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// RasterCollections: Resources returned.
	RasterCollections []*RasterCollection `json:"rasterCollections,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollectionsListResponse) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollectionsListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollectionsRaster: A raster resource.
type RasterCollectionsRaster struct {
	// Bbox: A rectangular bounding box which contains all of the data in
	// this Raster. The box is expressed as \"west, south, east, north\".
	// The numbers represent latitudes and longitudes in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this raster. The value is an RFC
	// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// Description: The description of this Raster, supplied by the author.
	Description string `json:"description,omitempty"`

	// Id: A globally unique ID, used to refer to this Raster.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this raster. The value is
	// an RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// Name: The name of this Raster, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProjectId: The ID of the project that this Raster is in.
	ProjectId string `json:"projectId,omitempty"`

	// RasterType: The type of this Raster. Always "image" today.
	RasterType string `json:"rasterType,omitempty"`

	// Tags: Tags of this Raster.
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bbox") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollectionsRaster) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollectionsRaster
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollectionsRasterBatchDeleteRequest: The request sent to
// rasterCollections.Rasters.BatchDelete.
type RasterCollectionsRasterBatchDeleteRequest struct {
	// Ids: An array of Raster asset IDs to be removed from this
	// RasterCollection.
	Ids []string `json:"ids,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Ids") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollectionsRasterBatchDeleteRequest) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollectionsRasterBatchDeleteRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollectionsRastersBatchDeleteResponse: The response returned by
// a call to rasterCollections.rasters.batchDelete.
type RasterCollectionsRastersBatchDeleteResponse struct {
}

// RasterCollectionsRastersBatchInsertRequest: The request sent to
// rasterCollections.Rasters.BatchInsert.
type RasterCollectionsRastersBatchInsertRequest struct {
	// Ids: An array of Raster asset IDs to be added to this
	// RasterCollection.
	Ids []string `json:"ids,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Ids") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollectionsRastersBatchInsertRequest) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollectionsRastersBatchInsertRequest
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RasterCollectionsRastersBatchInsertResponse: The response returned by
// a call to rasterCollections.rasters.batchInsert.
type RasterCollectionsRastersBatchInsertResponse struct {
}

// RasterCollectionsRastersListResponse: The response returned by a call
// to rasterCollections.rasters.List.
type RasterCollectionsRastersListResponse struct {
	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Rasters: Resources returned.
	Rasters []*RasterCollectionsRaster `json:"rasters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RasterCollectionsRastersListResponse) MarshalJSON() ([]byte, error) {
	type noMethod RasterCollectionsRastersListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// RastersListResponse: The response returned by a call to rasters.List.
type RastersListResponse struct {
	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Rasters: Resources returned.
	Rasters []*Raster `json:"rasters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RastersListResponse) MarshalJSON() ([]byte, error) {
	type noMethod RastersListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ScaledShape: Parameters for styling points as scaled shapes.
type ScaledShape struct {
	// Border: Border color/width of the shape. If not specified the shape
	// won't have a border.
	Border *Border `json:"border,omitempty"`

	// Fill: The fill color of the shape. If not specified the shape will be
	// transparent (although the borders may not be).
	Fill *Color `json:"fill,omitempty"`

	// Shape: Name of the shape.
	//
	// Possible values:
	//   "circle"
	Shape string `json:"shape,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Border") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ScaledShape) MarshalJSON() ([]byte, error) {
	type noMethod ScaledShape
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ScalingFunction: Parameters for scaling scaled shapes.
type ScalingFunction struct {
	// Column: Name of the numeric column used to scale a shape.
	Column string `json:"column,omitempty"`

	// ScalingType: The type of scaling function to use. Defaults to SQRT.
	// Currently only linear and square root scaling are supported.
	//
	// Possible values:
	//   "linear"
	//   "sqrt"
	ScalingType string `json:"scalingType,omitempty"`

	// SizeRange: The range of shape sizes, in pixels. For circles, the size
	// corresponds to the diameter.
	SizeRange *SizeRange `json:"sizeRange,omitempty"`

	// ValueRange: The range of values to display across the size range.
	ValueRange *ValueRange `json:"valueRange,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Column") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ScalingFunction) MarshalJSON() ([]byte, error) {
	type noMethod ScalingFunction
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Schema: A schema indicating the properties which may be associated
// with features within a Table, and the types of those properties.
type Schema struct {
	// Columns: An array of TableColumn objects. The first object in the
	// array must be named geometry and be of type points, lineStrings,
	// polygons, or mixedGeometry.
	Columns []*TableColumn `json:"columns,omitempty"`

	// PrimaryGeometry: The name of the column that contains a feature's
	// geometry. This field can be omitted during table create; Google Maps
	// Engine supports only a single geometry column, which must be named
	// geometry and be the first object in the columns array.
	PrimaryGeometry string `json:"primaryGeometry,omitempty"`

	// PrimaryKey: The name of the column that contains the unique
	// identifier of a Feature.
	PrimaryKey string `json:"primaryKey,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Columns") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Schema) MarshalJSON() ([]byte, error) {
	type noMethod Schema
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// SizeRange: Scaled shape size range in pixels. For circles, size
// corresponds to diameter.
type SizeRange struct {
	// Max: Maximum size, in pixels.
	Max float64 `json:"max,omitempty"`

	// Min: Minimum size, in pixels.
	Min float64 `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SizeRange) MarshalJSON() ([]byte, error) {
	type noMethod SizeRange
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// Table: A collection of geographic features, and associated metadata.
type Table struct {
	// Bbox: A rectangular bounding box which contains all of the data in
	// this Table. The box is expressed as \"west, south, east, north\". The
	// numbers represent latitude and longitude in decimal degrees.
	Bbox []float64 `json:"bbox,omitempty"`

	// CreationTime: The creation time of this table. The value is an RFC
	// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	CreationTime string `json:"creationTime,omitempty"`

	// CreatorEmail: The email address of the creator of this table. This is
	// only returned on GET requests and not LIST requests.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Description: The description of this table, supplied by the author.
	Description string `json:"description,omitempty"`

	// DraftAccessList: Deprecated: The name of an access list of the Map
	// Editor type. The user on whose behalf the request is being sent must
	// be an editor on that access list. Note: Google Maps Engine no longer
	// uses access lists. Instead, each asset has its own list of
	// permissions. For backward compatibility, the API still accepts access
	// lists for projects that are already using access lists. If you
	// created a GME account/project after July 14th, 2014, you will not be
	// able to send API requests that include access lists. Note: This is an
	// input field only. It is not returned in response to a list or get
	// request.
	DraftAccessList string `json:"draftAccessList,omitempty"`

	// Etag: The ETag, used to refer to the current version of the asset.
	Etag string `json:"etag,omitempty"`

	// Files: The files associated with this table.
	Files []*File `json:"files,omitempty"`

	// Id: A globally unique ID, used to refer to this table.
	Id string `json:"id,omitempty"`

	// LastModifiedTime: The last modified time of this table. The value is
	// an RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z).
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// LastModifierEmail: The email address of the last modifier of this
	// table. This is only returned on GET requests and not LIST requests.
	LastModifierEmail string `json:"lastModifierEmail,omitempty"`

	// Name: The name of this table, supplied by the author.
	Name string `json:"name,omitempty"`

	// ProcessingStatus: The processing status of this table.
	//
	// Possible values:
	//   "complete"
	//   "failed"
	//   "notReady"
	//   "processing"
	//   "ready"
	ProcessingStatus string `json:"processingStatus,omitempty"`

	// ProjectId: The ID of the project to which the table belongs.
	ProjectId string `json:"projectId,omitempty"`

	// PublishedAccessList: Deprecated: The access list to whom view
	// permissions are granted. The value must be the name of a Maps Engine
	// access list of the Map Viewer type, and the user must be a viewer on
	// that list. Note: Google Maps Engine no longer uses access lists.
	// Instead, each asset has its own list of permissions. For backward
	// compatibility, the API still accepts access lists for projects that
	// are already using access lists. If you created a GME account/project
	// after July 14th, 2014, you will not be able to send API requests that
	// include access lists. Note: This is an input field only. It is not
	// returned in response to a list or get request.
	PublishedAccessList string `json:"publishedAccessList,omitempty"`

	// Schema: The schema for this table. Note: The schema is returned in
	// response to a get request but not a list request. After requesting a
	// list of tables, you'll need to send a get request to retrieve the
	// schema for each table.
	Schema *Schema `json:"schema,omitempty"`

	// SourceEncoding: Encoding of the uploaded files. Valid values include
	// UTF-8, CP1251, ISO 8859-1, and Shift_JIS.
	SourceEncoding string `json:"sourceEncoding,omitempty"`

	// Tags: An array of text strings, with each string representing a tag.
	// More information about tags can be found in the Tagging data article
	// of the Maps Engine help center.
	Tags []string `json:"tags,omitempty"`

	// WritersCanEditPermissions: If true, WRITERs of the asset are able to
	// edit the asset permissions.
	WritersCanEditPermissions bool `json:"writersCanEditPermissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bbox") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Table) MarshalJSON() ([]byte, error) {
	type noMethod Table
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

type TableColumn struct {
	// Name: The column name.
	Name string `json:"name,omitempty"`

	// Type: The type of data stored in this column.
	//
	// Possible values:
	//   "datetime"
	//   "double"
	//   "integer"
	//   "lineStrings"
	//   "mixedGeometry"
	//   "points"
	//   "polygons"
	//   "string"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TableColumn) MarshalJSON() ([]byte, error) {
	type noMethod TableColumn
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// TablesListResponse: The response returned by a call to tables.List.
// Note: The list response does not include all the fields available in
// a table. Refer to the table resource description for details of the
// fields that are not included. You'll need to send a get request to
// retrieve the additional fields for each table.
type TablesListResponse struct {
	// NextPageToken: Next page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Tables: Resources returned.
	Tables []*Table `json:"tables,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TablesListResponse) MarshalJSON() ([]byte, error) {
	type noMethod TablesListResponse
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ValueRange: Range of values used for scaling shapes. The min/max
// values will be drawn as shapes with the min/max size.
type ValueRange struct {
	// Max: Maximum value.
	Max float64 `json:"max,omitempty"`

	// Min: Minimum value.
	Min float64 `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ValueRange) MarshalJSON() ([]byte, error) {
	type noMethod ValueRange
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// VectorStyle: A vector style contains styling information for vector
// layer.
type VectorStyle struct {
	DisplayRules []*DisplayRule `json:"displayRules,omitempty"`

	// FeatureInfo: Individual feature info, this is called Info Window in
	// Maps Engine UI. If not provided, a default template with all
	// attributes will be generated.
	FeatureInfo *FeatureInfo `json:"featureInfo,omitempty"`

	// Type: The type of the vector style. Currently, only displayRule is
	// supported.
	//
	// Possible values:
	//   "displayRule"
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayRules") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *VectorStyle) MarshalJSON() ([]byte, error) {
	type noMethod VectorStyle
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// ZoomLevels: Zoom level range. Zoom levels are restricted between 0
// and 24, inclusive.
type ZoomLevels struct {
	// Max: Maximum zoom level.
	Max int64 `json:"max,omitempty"`

	// Min: Minimum zoom level.
	Min int64 `json:"min,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Max") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ZoomLevels) MarshalJSON() ([]byte, error) {
	type noMethod ZoomLevels
	raw := noMethod(*s)
	return internal.MarshalJSON(raw, s.ForceSendFields)
}

// method id "mapsengine.assets.get":

type AssetsGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a particular asset.
func (r *AssetsService) Get(id string) *AssetsGetCall {
	c := &AssetsGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsGetCall) Fields(s ...googleapi.Field) *AssetsGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsGetCall) Context(ctx context.Context) *AssetsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *AssetsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *AssetsGetCall) Do() (*Asset, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Asset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a particular asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.assets.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{id}",
	//   "response": {
	//     "$ref": "Asset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.assets.list":

type AssetsListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all assets readable by the current user.
func (r *AssetsService) List() *AssetsListCall {
	c := &AssetsListCall{s: r.s, opt_: urlValues()}
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *AssetsListCall) Bbox(bbox string) *AssetsListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *AssetsListCall) CreatedAfter(createdAfter string) *AssetsListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *AssetsListCall) CreatedBefore(createdBefore string) *AssetsListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *AssetsListCall) CreatorEmail(creatorEmail string) *AssetsListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *AssetsListCall) MaxResults(maxResults int64) *AssetsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *AssetsListCall) ModifiedAfter(modifiedAfter string) *AssetsListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *AssetsListCall) ModifiedBefore(modifiedBefore string) *AssetsListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *AssetsListCall) PageToken(pageToken string) *AssetsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *AssetsListCall) ProjectId(projectId string) *AssetsListCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *AssetsListCall) Role(role string) *AssetsListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *AssetsListCall) Search(search string) *AssetsListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *AssetsListCall) Tags(tags string) *AssetsListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Type sets the optional parameter "type": A comma separated list of
// asset types. Returned assets will have one of the types from the
// provided list. Supported values are 'map', 'layer',
// 'rasterCollection' and 'table'.
func (c *AssetsListCall) Type(type_ string) *AssetsListCall {
	c.opt_.Set("type_", fmt.Sprintf("%v", type_))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsListCall) Fields(s ...googleapi.Field) *AssetsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsListCall) Context(ctx context.Context) *AssetsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AssetsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *AssetsListCall) Do() (*AssetsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *AssetsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all assets readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.assets.list",
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "type": {
	//       "description": "A comma separated list of asset types. Returned assets will have one of the types from the provided list. Supported values are 'map', 'layer', 'rasterCollection' and 'table'.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets",
	//   "response": {
	//     "$ref": "AssetsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.assets.parents.list":

type AssetsParentsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all parent ids of the specified asset.
func (r *AssetsParentsService) List(id string) *AssetsParentsListCall {
	c := &AssetsParentsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *AssetsParentsListCall) MaxResults(maxResults int64) *AssetsParentsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *AssetsParentsListCall) PageToken(pageToken string) *AssetsParentsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsParentsListCall) Fields(s ...googleapi.Field) *AssetsParentsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsParentsListCall) Context(ctx context.Context) *AssetsParentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AssetsParentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{id}/parents")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *AssetsParentsListCall) Do() (*ParentsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ParentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all parent ids of the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.assets.parents.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose parents will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{id}/parents",
	//   "response": {
	//     "$ref": "ParentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.assets.permissions.list":

type AssetsPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *AssetsPermissionsService) List(id string) *AssetsPermissionsListCall {
	c := &AssetsPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsPermissionsListCall) Fields(s ...googleapi.Field) *AssetsPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsPermissionsListCall) Context(ctx context.Context) *AssetsPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *AssetsPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "assets/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *AssetsPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.assets.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "assets/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.cancelProcessing":

type LayersCancelProcessingCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// CancelProcessing: Cancel processing on a layer asset.
func (r *LayersService) CancelProcessing(id string) *LayersCancelProcessingCall {
	c := &LayersCancelProcessingCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersCancelProcessingCall) Fields(s ...googleapi.Field) *LayersCancelProcessingCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersCancelProcessingCall) Context(ctx context.Context) *LayersCancelProcessingCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersCancelProcessingCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/cancelProcessing")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersCancelProcessingCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel processing on a layer asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.cancelProcessing",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/cancelProcessing",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.create":

type LayersCreateCall struct {
	s     *Service
	layer *Layer
	opt_  url.Values
	ctx_  context.Context
}

// Create: Create a layer asset.
func (r *LayersService) Create(layer *Layer) *LayersCreateCall {
	c := &LayersCreateCall{s: r.s, opt_: urlValues()}
	c.layer = layer
	return c
}

// Process sets the optional parameter "process": Whether to queue the
// created layer for processing.
func (c *LayersCreateCall) Process(process bool) *LayersCreateCall {
	c.opt_.Set("process", fmt.Sprintf("%v", process))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersCreateCall) Fields(s ...googleapi.Field) *LayersCreateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersCreateCall) Context(ctx context.Context) *LayersCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.layer)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersCreateCall) Do() (*Layer, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Layer
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a layer asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.create",
	//   "parameters": {
	//     "process": {
	//       "description": "Whether to queue the created layer for processing.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "layers",
	//   "request": {
	//     "$ref": "Layer"
	//   },
	//   "response": {
	//     "$ref": "Layer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.delete":

type LayersDeleteCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Delete: Delete a layer.
func (r *LayersService) Delete(id string) *LayersDeleteCall {
	c := &LayersDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersDeleteCall) Fields(s ...googleapi.Field) *LayersDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersDeleteCall) Context(ctx context.Context) *LayersDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersDeleteCall) Do() error {
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
	//   "description": "Delete a layer.",
	//   "httpMethod": "DELETE",
	//   "id": "mapsengine.layers.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer. Only the layer creator or project owner are permitted to delete. If the layer is published, or included in a map, the request will fail. Unpublish the layer, and remove it from all maps prior to deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.get":

type LayersGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a particular layer.
func (r *LayersService) Get(id string) *LayersGetCall {
	c := &LayersGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Version sets the optional parameter "version": Deprecated: The
// version parameter indicates which version of the layer should be
// returned. When version is set to published, the published version of
// the layer will be returned. Please use the layers.getPublished
// endpoint instead.
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *LayersGetCall) Version(version string) *LayersGetCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersGetCall) Fields(s ...googleapi.Field) *LayersGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersGetCall) Context(ctx context.Context) *LayersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersGetCall) Do() (*Layer, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Layer
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a particular layer.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Deprecated: The version parameter indicates which version of the layer should be returned. When version is set to published, the published version of the layer will be returned. Please use the layers.getPublished endpoint instead.",
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}",
	//   "response": {
	//     "$ref": "Layer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.getPublished":

type LayersGetPublishedCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// GetPublished: Return the published metadata for a particular layer.
func (r *LayersService) GetPublished(id string) *LayersGetPublishedCall {
	c := &LayersGetPublishedCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersGetPublishedCall) Fields(s ...googleapi.Field) *LayersGetPublishedCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersGetPublishedCall) Context(ctx context.Context) *LayersGetPublishedCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersGetPublishedCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/published")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersGetPublishedCall) Do() (*PublishedLayer, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishedLayer
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return the published metadata for a particular layer.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.getPublished",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/published",
	//   "response": {
	//     "$ref": "PublishedLayer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.list":

type LayersListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all layers readable by the current user.
func (r *LayersService) List() *LayersListCall {
	c := &LayersListCall{s: r.s, opt_: urlValues()}
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *LayersListCall) Bbox(bbox string) *LayersListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *LayersListCall) CreatedAfter(createdAfter string) *LayersListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *LayersListCall) CreatedBefore(createdBefore string) *LayersListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *LayersListCall) CreatorEmail(creatorEmail string) *LayersListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *LayersListCall) MaxResults(maxResults int64) *LayersListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *LayersListCall) ModifiedAfter(modifiedAfter string) *LayersListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *LayersListCall) ModifiedBefore(modifiedBefore string) *LayersListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *LayersListCall) PageToken(pageToken string) *LayersListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProcessingStatus sets the optional parameter "processingStatus":
//
// Possible values:
//   "complete" - The layer has completed processing.
//   "failed" - The layer has failed processing.
//   "notReady" - The layer is not ready for processing.
//   "processing" - The layer is processing.
//   "ready" - The layer is ready for processing.
func (c *LayersListCall) ProcessingStatus(processingStatus string) *LayersListCall {
	c.opt_.Set("processingStatus", fmt.Sprintf("%v", processingStatus))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *LayersListCall) ProjectId(projectId string) *LayersListCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *LayersListCall) Role(role string) *LayersListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *LayersListCall) Search(search string) *LayersListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *LayersListCall) Tags(tags string) *LayersListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersListCall) Fields(s ...googleapi.Field) *LayersListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersListCall) Context(ctx context.Context) *LayersListCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersListCall) Do() (*LayersListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *LayersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all layers readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.list",
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "processingStatus": {
	//       "enum": [
	//         "complete",
	//         "failed",
	//         "notReady",
	//         "processing",
	//         "ready"
	//       ],
	//       "enumDescriptions": [
	//         "The layer has completed processing.",
	//         "The layer has failed processing.",
	//         "The layer is not ready for processing.",
	//         "The layer is processing.",
	//         "The layer is ready for processing."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers",
	//   "response": {
	//     "$ref": "LayersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.listPublished":

type LayersListPublishedCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// ListPublished: Return all published layers readable by the current
// user.
func (r *LayersService) ListPublished() *LayersListPublishedCall {
	c := &LayersListPublishedCall{s: r.s, opt_: urlValues()}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *LayersListPublishedCall) MaxResults(maxResults int64) *LayersListPublishedCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *LayersListPublishedCall) PageToken(pageToken string) *LayersListPublishedCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *LayersListPublishedCall) ProjectId(projectId string) *LayersListPublishedCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *LayersListPublishedCall) Search(search string) *LayersListPublishedCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersListPublishedCall) Fields(s ...googleapi.Field) *LayersListPublishedCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersListPublishedCall) Context(ctx context.Context) *LayersListPublishedCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersListPublishedCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/published")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersListPublishedCall) Do() (*PublishedLayersListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishedLayersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all published layers readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.listPublished",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/published",
	//   "response": {
	//     "$ref": "PublishedLayersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.patch":

type LayersPatchCall struct {
	s     *Service
	id    string
	layer *Layer
	opt_  url.Values
	ctx_  context.Context
}

// Patch: Mutate a layer asset.
func (r *LayersService) Patch(id string, layer *Layer) *LayersPatchCall {
	c := &LayersPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.layer = layer
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersPatchCall) Fields(s ...googleapi.Field) *LayersPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersPatchCall) Context(ctx context.Context) *LayersPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.layer)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersPatchCall) Do() error {
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
	//   "description": "Mutate a layer asset.",
	//   "httpMethod": "PATCH",
	//   "id": "mapsengine.layers.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}",
	//   "request": {
	//     "$ref": "Layer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.process":

type LayersProcessCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Process: Process a layer asset.
func (r *LayersService) Process(id string) *LayersProcessCall {
	c := &LayersProcessCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersProcessCall) Fields(s ...googleapi.Field) *LayersProcessCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersProcessCall) Context(ctx context.Context) *LayersProcessCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersProcessCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/process")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersProcessCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Process a layer asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.process",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/process",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.publish":

type LayersPublishCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Publish: Publish a layer asset.
func (r *LayersService) Publish(id string) *LayersPublishCall {
	c := &LayersPublishCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Force sets the optional parameter "force": If set to true, the API
// will allow publication of the layer even if it's out of date. If not
// true, you'll need to reprocess any out-of-date layer before
// publishing.
func (c *LayersPublishCall) Force(force bool) *LayersPublishCall {
	c.opt_.Set("force", fmt.Sprintf("%v", force))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersPublishCall) Fields(s ...googleapi.Field) *LayersPublishCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersPublishCall) Context(ctx context.Context) *LayersPublishCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersPublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/publish")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersPublishCall) Do() (*PublishResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Publish a layer asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.publish",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "force": {
	//       "description": "If set to true, the API will allow publication of the layer even if it's out of date. If not true, you'll need to reprocess any out-of-date layer before publishing.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/publish",
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.unpublish":

type LayersUnpublishCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Unpublish: Unpublish a layer asset.
func (r *LayersService) Unpublish(id string) *LayersUnpublishCall {
	c := &LayersUnpublishCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersUnpublishCall) Fields(s ...googleapi.Field) *LayersUnpublishCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersUnpublishCall) Context(ctx context.Context) *LayersUnpublishCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersUnpublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/unpublish")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersUnpublishCall) Do() (*PublishResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Unpublish a layer asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.unpublish",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/unpublish",
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.parents.list":

type LayersParentsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all parent ids of the specified layer.
func (r *LayersParentsService) List(id string) *LayersParentsListCall {
	c := &LayersParentsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *LayersParentsListCall) MaxResults(maxResults int64) *LayersParentsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *LayersParentsListCall) PageToken(pageToken string) *LayersParentsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersParentsListCall) Fields(s ...googleapi.Field) *LayersParentsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersParentsListCall) Context(ctx context.Context) *LayersParentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersParentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/parents")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersParentsListCall) Do() (*ParentsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ParentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all parent ids of the specified layer.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.parents.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the layer whose parents will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/parents",
	//   "response": {
	//     "$ref": "ParentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.layers.permissions.batchDelete":

type LayersPermissionsBatchDeleteCall struct {
	s                             *Service
	id                            string
	permissionsbatchdeleterequest *PermissionsBatchDeleteRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchDelete: Remove permission entries from an already existing
// asset.
func (r *LayersPermissionsService) BatchDelete(id string, permissionsbatchdeleterequest *PermissionsBatchDeleteRequest) *LayersPermissionsBatchDeleteCall {
	c := &LayersPermissionsBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchdeleterequest = permissionsbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersPermissionsBatchDeleteCall) Fields(s ...googleapi.Field) *LayersPermissionsBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersPermissionsBatchDeleteCall) Context(ctx context.Context) *LayersPermissionsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersPermissionsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/permissions/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersPermissionsBatchDeleteCall) Do() (*PermissionsBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove permission entries from an already existing asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.permissions.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset from which permissions will be removed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/permissions/batchDelete",
	//   "request": {
	//     "$ref": "PermissionsBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.permissions.batchUpdate":

type LayersPermissionsBatchUpdateCall struct {
	s                             *Service
	id                            string
	permissionsbatchupdaterequest *PermissionsBatchUpdateRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchUpdate: Add or update permission entries to an already existing
// asset.
//
// An asset can hold up to 20 different permission entries. Each
// batchInsert request is atomic.
func (r *LayersPermissionsService) BatchUpdate(id string, permissionsbatchupdaterequest *PermissionsBatchUpdateRequest) *LayersPermissionsBatchUpdateCall {
	c := &LayersPermissionsBatchUpdateCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchupdaterequest = permissionsbatchupdaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersPermissionsBatchUpdateCall) Fields(s ...googleapi.Field) *LayersPermissionsBatchUpdateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersPermissionsBatchUpdateCall) Context(ctx context.Context) *LayersPermissionsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersPermissionsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchupdaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/permissions/batchUpdate")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersPermissionsBatchUpdateCall) Do() (*PermissionsBatchUpdateResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add or update permission entries to an already existing asset.\n\nAn asset can hold up to 20 different permission entries. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.layers.permissions.batchUpdate",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset to which permissions will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/permissions/batchUpdate",
	//   "request": {
	//     "$ref": "PermissionsBatchUpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchUpdateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.layers.permissions.list":

type LayersPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *LayersPermissionsService) List(id string) *LayersPermissionsListCall {
	c := &LayersPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LayersPermissionsListCall) Fields(s ...googleapi.Field) *LayersPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LayersPermissionsListCall) Context(ctx context.Context) *LayersPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *LayersPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "layers/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *LayersPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.layers.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "layers/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.maps.create":

type MapsCreateCall struct {
	s    *Service
	map_ *Map
	opt_ url.Values
	ctx_ context.Context
}

// Create: Create a map asset.
func (r *MapsService) Create(map_ *Map) *MapsCreateCall {
	c := &MapsCreateCall{s: r.s, opt_: urlValues()}
	c.map_ = map_
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsCreateCall) Fields(s ...googleapi.Field) *MapsCreateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsCreateCall) Context(ctx context.Context) *MapsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.map_)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsCreateCall) Do() (*Map, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Map
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a map asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.maps.create",
	//   "path": "maps",
	//   "request": {
	//     "$ref": "Map"
	//   },
	//   "response": {
	//     "$ref": "Map"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.delete":

type MapsDeleteCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Delete: Delete a map.
func (r *MapsService) Delete(id string) *MapsDeleteCall {
	c := &MapsDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsDeleteCall) Fields(s ...googleapi.Field) *MapsDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsDeleteCall) Context(ctx context.Context) *MapsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsDeleteCall) Do() error {
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
	//   "description": "Delete a map.",
	//   "httpMethod": "DELETE",
	//   "id": "mapsengine.maps.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the map. Only the map creator or project owner are permitted to delete. If the map is published the request will fail. Unpublish the map prior to deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.get":

type MapsGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a particular map.
func (r *MapsService) Get(id string) *MapsGetCall {
	c := &MapsGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Version sets the optional parameter "version": Deprecated: The
// version parameter indicates which version of the map should be
// returned. When version is set to published, the published version of
// the map will be returned. Please use the maps.getPublished endpoint
// instead.
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *MapsGetCall) Version(version string) *MapsGetCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsGetCall) Fields(s ...googleapi.Field) *MapsGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsGetCall) Context(ctx context.Context) *MapsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsGetCall) Do() (*Map, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Map
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a particular map.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.maps.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the map.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Deprecated: The version parameter indicates which version of the map should be returned. When version is set to published, the published version of the map will be returned. Please use the maps.getPublished endpoint instead.",
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}",
	//   "response": {
	//     "$ref": "Map"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.maps.getPublished":

type MapsGetPublishedCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// GetPublished: Return the published metadata for a particular map.
func (r *MapsService) GetPublished(id string) *MapsGetPublishedCall {
	c := &MapsGetPublishedCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsGetPublishedCall) Fields(s ...googleapi.Field) *MapsGetPublishedCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsGetPublishedCall) Context(ctx context.Context) *MapsGetPublishedCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsGetPublishedCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/published")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsGetPublishedCall) Do() (*PublishedMap, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishedMap
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return the published metadata for a particular map.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.maps.getPublished",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the map.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/published",
	//   "response": {
	//     "$ref": "PublishedMap"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.maps.list":

type MapsListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all maps readable by the current user.
func (r *MapsService) List() *MapsListCall {
	c := &MapsListCall{s: r.s, opt_: urlValues()}
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *MapsListCall) Bbox(bbox string) *MapsListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *MapsListCall) CreatedAfter(createdAfter string) *MapsListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *MapsListCall) CreatedBefore(createdBefore string) *MapsListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *MapsListCall) CreatorEmail(creatorEmail string) *MapsListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *MapsListCall) MaxResults(maxResults int64) *MapsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *MapsListCall) ModifiedAfter(modifiedAfter string) *MapsListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *MapsListCall) ModifiedBefore(modifiedBefore string) *MapsListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *MapsListCall) PageToken(pageToken string) *MapsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProcessingStatus sets the optional parameter "processingStatus":
//
// Possible values:
//   "complete" - The map has completed processing.
//   "failed" - The map has failed processing.
//   "notReady" - The map is not ready for processing.
//   "processing" - The map is processing.
func (c *MapsListCall) ProcessingStatus(processingStatus string) *MapsListCall {
	c.opt_.Set("processingStatus", fmt.Sprintf("%v", processingStatus))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *MapsListCall) ProjectId(projectId string) *MapsListCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *MapsListCall) Role(role string) *MapsListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *MapsListCall) Search(search string) *MapsListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *MapsListCall) Tags(tags string) *MapsListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Version sets the optional parameter "version": Deprecated: The
// version parameter indicates which version of the maps should be
// returned. When version is set to published this parameter will filter
// the result set to include only maps that are published. Please use
// the maps.listPublished endpoint instead.
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *MapsListCall) Version(version string) *MapsListCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsListCall) Fields(s ...googleapi.Field) *MapsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsListCall) Context(ctx context.Context) *MapsListCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsListCall) Do() (*MapsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *MapsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all maps readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.maps.list",
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "processingStatus": {
	//       "enum": [
	//         "complete",
	//         "failed",
	//         "notReady",
	//         "processing"
	//       ],
	//       "enumDescriptions": [
	//         "The map has completed processing.",
	//         "The map has failed processing.",
	//         "The map is not ready for processing.",
	//         "The map is processing."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "Deprecated: The version parameter indicates which version of the maps should be returned. When version is set to published this parameter will filter the result set to include only maps that are published. Please use the maps.listPublished endpoint instead.",
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps",
	//   "response": {
	//     "$ref": "MapsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.maps.listPublished":

type MapsListPublishedCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// ListPublished: Return all published maps readable by the current
// user.
func (r *MapsService) ListPublished() *MapsListPublishedCall {
	c := &MapsListPublishedCall{s: r.s, opt_: urlValues()}
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *MapsListPublishedCall) MaxResults(maxResults int64) *MapsListPublishedCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *MapsListPublishedCall) PageToken(pageToken string) *MapsListPublishedCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *MapsListPublishedCall) ProjectId(projectId string) *MapsListPublishedCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *MapsListPublishedCall) Search(search string) *MapsListPublishedCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsListPublishedCall) Fields(s ...googleapi.Field) *MapsListPublishedCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsListPublishedCall) Context(ctx context.Context) *MapsListPublishedCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsListPublishedCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/published")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsListPublishedCall) Do() (*PublishedMapsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishedMapsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all published maps readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.maps.listPublished",
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/published",
	//   "response": {
	//     "$ref": "PublishedMapsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.maps.patch":

type MapsPatchCall struct {
	s    *Service
	id   string
	map_ *Map
	opt_ url.Values
	ctx_ context.Context
}

// Patch: Mutate a map asset.
func (r *MapsService) Patch(id string, map_ *Map) *MapsPatchCall {
	c := &MapsPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.map_ = map_
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsPatchCall) Fields(s ...googleapi.Field) *MapsPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsPatchCall) Context(ctx context.Context) *MapsPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.map_)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsPatchCall) Do() error {
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
	//   "description": "Mutate a map asset.",
	//   "httpMethod": "PATCH",
	//   "id": "mapsengine.maps.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the map.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}",
	//   "request": {
	//     "$ref": "Map"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.publish":

type MapsPublishCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Publish: Publish a map asset.
func (r *MapsService) Publish(id string) *MapsPublishCall {
	c := &MapsPublishCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Force sets the optional parameter "force": If set to true, the API
// will allow publication of the map even if it's out of date. If false,
// the map must have a processingStatus of complete before publishing.
func (c *MapsPublishCall) Force(force bool) *MapsPublishCall {
	c.opt_.Set("force", fmt.Sprintf("%v", force))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsPublishCall) Fields(s ...googleapi.Field) *MapsPublishCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsPublishCall) Context(ctx context.Context) *MapsPublishCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsPublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/publish")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsPublishCall) Do() (*PublishResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Publish a map asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.maps.publish",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "force": {
	//       "description": "If set to true, the API will allow publication of the map even if it's out of date. If false, the map must have a processingStatus of complete before publishing.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "id": {
	//       "description": "The ID of the map.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/publish",
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.unpublish":

type MapsUnpublishCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Unpublish: Unpublish a map asset.
func (r *MapsService) Unpublish(id string) *MapsUnpublishCall {
	c := &MapsUnpublishCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsUnpublishCall) Fields(s ...googleapi.Field) *MapsUnpublishCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsUnpublishCall) Context(ctx context.Context) *MapsUnpublishCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsUnpublishCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/unpublish")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsUnpublishCall) Do() (*PublishResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Unpublish a map asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.maps.unpublish",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the map.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/unpublish",
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.permissions.batchDelete":

type MapsPermissionsBatchDeleteCall struct {
	s                             *Service
	id                            string
	permissionsbatchdeleterequest *PermissionsBatchDeleteRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchDelete: Remove permission entries from an already existing
// asset.
func (r *MapsPermissionsService) BatchDelete(id string, permissionsbatchdeleterequest *PermissionsBatchDeleteRequest) *MapsPermissionsBatchDeleteCall {
	c := &MapsPermissionsBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchdeleterequest = permissionsbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsPermissionsBatchDeleteCall) Fields(s ...googleapi.Field) *MapsPermissionsBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsPermissionsBatchDeleteCall) Context(ctx context.Context) *MapsPermissionsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsPermissionsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/permissions/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsPermissionsBatchDeleteCall) Do() (*PermissionsBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove permission entries from an already existing asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.maps.permissions.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset from which permissions will be removed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/permissions/batchDelete",
	//   "request": {
	//     "$ref": "PermissionsBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.permissions.batchUpdate":

type MapsPermissionsBatchUpdateCall struct {
	s                             *Service
	id                            string
	permissionsbatchupdaterequest *PermissionsBatchUpdateRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchUpdate: Add or update permission entries to an already existing
// asset.
//
// An asset can hold up to 20 different permission entries. Each
// batchInsert request is atomic.
func (r *MapsPermissionsService) BatchUpdate(id string, permissionsbatchupdaterequest *PermissionsBatchUpdateRequest) *MapsPermissionsBatchUpdateCall {
	c := &MapsPermissionsBatchUpdateCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchupdaterequest = permissionsbatchupdaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsPermissionsBatchUpdateCall) Fields(s ...googleapi.Field) *MapsPermissionsBatchUpdateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsPermissionsBatchUpdateCall) Context(ctx context.Context) *MapsPermissionsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsPermissionsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchupdaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/permissions/batchUpdate")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsPermissionsBatchUpdateCall) Do() (*PermissionsBatchUpdateResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add or update permission entries to an already existing asset.\n\nAn asset can hold up to 20 different permission entries. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.maps.permissions.batchUpdate",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset to which permissions will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/permissions/batchUpdate",
	//   "request": {
	//     "$ref": "PermissionsBatchUpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchUpdateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.maps.permissions.list":

type MapsPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *MapsPermissionsService) List(id string) *MapsPermissionsListCall {
	c := &MapsPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MapsPermissionsListCall) Fields(s ...googleapi.Field) *MapsPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MapsPermissionsListCall) Context(ctx context.Context) *MapsPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *MapsPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "maps/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *MapsPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.maps.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "maps/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.projects.list":

type ProjectsListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all projects readable by the current user.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, opt_: urlValues()}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsListCall) Context(ctx context.Context) *ProjectsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *ProjectsListCall) Do() (*ProjectsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProjectsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all projects readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.projects.list",
	//   "path": "projects",
	//   "response": {
	//     "$ref": "ProjectsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.projects.icons.create":

type ProjectsIconsCreateCall struct {
	s                *Service
	projectId        string
	icon             *Icon
	opt_             url.Values
	media_           io.Reader
	resumable_       googleapi.SizeReaderAt
	mediaType_       string
	protocol_        string
	progressUpdater_ googleapi.ProgressUpdater
	ctx_             context.Context
}

// Create: Create an icon.
func (r *ProjectsIconsService) Create(projectId string, icon *Icon) *ProjectsIconsCreateCall {
	c := &ProjectsIconsCreateCall{s: r.s, opt_: urlValues()}
	c.projectId = projectId
	c.icon = icon
	return c
}

// Media specifies the media to upload in a single chunk. At most one of
// Media and ResumableMedia may be set.
func (c *ProjectsIconsCreateCall) Media(r io.Reader) *ProjectsIconsCreateCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx. At most one of Media and ResumableMedia may be
// set. mediaType identifies the MIME media type of the upload, such as
// "image/png". If mediaType is "", it will be auto-detected. The
// provided ctx will supersede any context previously provided to the
// Context method.
func (c *ProjectsIconsCreateCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *ProjectsIconsCreateCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *ProjectsIconsCreateCall) ProgressUpdater(pu googleapi.ProgressUpdater) *ProjectsIconsCreateCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIconsCreateCall) Fields(s ...googleapi.Field) *ProjectsIconsCreateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *ProjectsIconsCreateCall) Context(ctx context.Context) *ProjectsIconsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIconsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.icon)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/icons")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.opt_.Set("uploadType", c.protocol_)
	}
	urls += "?" + c.opt_.Encode()
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *ProjectsIconsCreateCall) Do() (*Icon, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      c.progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	var ret *Icon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create an icon.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.projects.icons.create",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "100KB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mapsengine/exp2/projects/{projectId}/icons"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mapsengine/exp2/projects/{projectId}/icons"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/icons",
	//   "request": {
	//     "$ref": "Icon"
	//   },
	//   "response": {
	//     "$ref": "Icon"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mapsengine.projects.icons.get":

type ProjectsIconsGetCall struct {
	s         *Service
	projectId string
	id        string
	opt_      url.Values
	ctx_      context.Context
}

// Get: Return an icon or its associated metadata
func (r *ProjectsIconsService) Get(projectId string, id string) *ProjectsIconsGetCall {
	c := &ProjectsIconsGetCall{s: r.s, opt_: urlValues()}
	c.projectId = projectId
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIconsGetCall) Fields(s ...googleapi.Field) *ProjectsIconsGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do and Download
// methods. Any pending HTTP request will be aborted if the provided
// context is canceled.
func (c *ProjectsIconsGetCall) Context(ctx context.Context) *ProjectsIconsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIconsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/icons/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Download fetches the API endpoint's "media" value, instead of the normal
// API response value. If the returned error is nil, the Response is guaranteed to
// have a 2xx status code. Callers must close the Response.Body as usual.
func (c *ProjectsIconsGetCall) Download() (*http.Response, error) {
	res, err := c.doRequest("media")
	if err != nil {
		return nil, err
	}
	if err := googleapi.CheckMediaResponse(res); err != nil {
		res.Body.Close()
		return nil, err
	}
	return res, nil
}

func (c *ProjectsIconsGetCall) Do() (*Icon, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Icon
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return an icon or its associated metadata",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.projects.icons.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the icon.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/icons/{id}",
	//   "response": {
	//     "$ref": "Icon"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ],
	//   "supportsMediaDownload": true
	// }

}

// method id "mapsengine.projects.icons.list":

type ProjectsIconsListCall struct {
	s         *Service
	projectId string
	opt_      url.Values
	ctx_      context.Context
}

// List: Return all icons in the current project
func (r *ProjectsIconsService) List(projectId string) *ProjectsIconsListCall {
	c := &ProjectsIconsListCall{s: r.s, opt_: urlValues()}
	c.projectId = projectId
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *ProjectsIconsListCall) MaxResults(maxResults int64) *ProjectsIconsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *ProjectsIconsListCall) PageToken(pageToken string) *ProjectsIconsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIconsListCall) Fields(s ...googleapi.Field) *ProjectsIconsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsIconsListCall) Context(ctx context.Context) *ProjectsIconsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIconsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{projectId}/icons")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *ProjectsIconsListCall) Do() (*IconsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *IconsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all icons in the current project",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.projects.icons.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{projectId}/icons",
	//   "response": {
	//     "$ref": "IconsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.cancelProcessing":

type RasterCollectionsCancelProcessingCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// CancelProcessing: Cancel processing on a raster collection asset.
func (r *RasterCollectionsService) CancelProcessing(id string) *RasterCollectionsCancelProcessingCall {
	c := &RasterCollectionsCancelProcessingCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsCancelProcessingCall) Fields(s ...googleapi.Field) *RasterCollectionsCancelProcessingCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsCancelProcessingCall) Context(ctx context.Context) *RasterCollectionsCancelProcessingCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsCancelProcessingCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/cancelProcessing")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsCancelProcessingCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancel processing on a raster collection asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.cancelProcessing",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/cancelProcessing",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.create":

type RasterCollectionsCreateCall struct {
	s                *Service
	rastercollection *RasterCollection
	opt_             url.Values
	ctx_             context.Context
}

// Create: Create a raster collection asset.
func (r *RasterCollectionsService) Create(rastercollection *RasterCollection) *RasterCollectionsCreateCall {
	c := &RasterCollectionsCreateCall{s: r.s, opt_: urlValues()}
	c.rastercollection = rastercollection
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsCreateCall) Fields(s ...googleapi.Field) *RasterCollectionsCreateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsCreateCall) Context(ctx context.Context) *RasterCollectionsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rastercollection)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsCreateCall) Do() (*RasterCollection, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a raster collection asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.create",
	//   "path": "rasterCollections",
	//   "request": {
	//     "$ref": "RasterCollection"
	//   },
	//   "response": {
	//     "$ref": "RasterCollection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.delete":

type RasterCollectionsDeleteCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Delete: Delete a raster collection.
func (r *RasterCollectionsService) Delete(id string) *RasterCollectionsDeleteCall {
	c := &RasterCollectionsDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsDeleteCall) Fields(s ...googleapi.Field) *RasterCollectionsDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsDeleteCall) Context(ctx context.Context) *RasterCollectionsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsDeleteCall) Do() error {
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
	//   "description": "Delete a raster collection.",
	//   "httpMethod": "DELETE",
	//   "id": "mapsengine.rasterCollections.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection. Only the raster collection creator or project owner are permitted to delete. If the rastor collection is included in a layer, the request will fail. Remove the raster collection from all layers prior to deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.get":

type RasterCollectionsGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a particular raster collection.
func (r *RasterCollectionsService) Get(id string) *RasterCollectionsGetCall {
	c := &RasterCollectionsGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsGetCall) Fields(s ...googleapi.Field) *RasterCollectionsGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsGetCall) Context(ctx context.Context) *RasterCollectionsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsGetCall) Do() (*RasterCollection, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollection
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a particular raster collection.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasterCollections.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}",
	//   "response": {
	//     "$ref": "RasterCollection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.list":

type RasterCollectionsListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all raster collections readable by the current user.
func (r *RasterCollectionsService) List() *RasterCollectionsListCall {
	c := &RasterCollectionsListCall{s: r.s, opt_: urlValues()}
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *RasterCollectionsListCall) Bbox(bbox string) *RasterCollectionsListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *RasterCollectionsListCall) CreatedAfter(createdAfter string) *RasterCollectionsListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *RasterCollectionsListCall) CreatedBefore(createdBefore string) *RasterCollectionsListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *RasterCollectionsListCall) CreatorEmail(creatorEmail string) *RasterCollectionsListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *RasterCollectionsListCall) MaxResults(maxResults int64) *RasterCollectionsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *RasterCollectionsListCall) ModifiedAfter(modifiedAfter string) *RasterCollectionsListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *RasterCollectionsListCall) ModifiedBefore(modifiedBefore string) *RasterCollectionsListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *RasterCollectionsListCall) PageToken(pageToken string) *RasterCollectionsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProcessingStatus sets the optional parameter "processingStatus":
//
// Possible values:
//   "complete" - The raster collection has completed processing.
//   "failed" - The raster collection has failed processing.
//   "notReady" - The raster collection is not ready for processing.
//   "processing" - The raster collection is processing.
//   "ready" - The raster collection is ready for processing.
func (c *RasterCollectionsListCall) ProcessingStatus(processingStatus string) *RasterCollectionsListCall {
	c.opt_.Set("processingStatus", fmt.Sprintf("%v", processingStatus))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *RasterCollectionsListCall) ProjectId(projectId string) *RasterCollectionsListCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *RasterCollectionsListCall) Role(role string) *RasterCollectionsListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *RasterCollectionsListCall) Search(search string) *RasterCollectionsListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *RasterCollectionsListCall) Tags(tags string) *RasterCollectionsListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsListCall) Fields(s ...googleapi.Field) *RasterCollectionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsListCall) Context(ctx context.Context) *RasterCollectionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsListCall) Do() (*RasterCollectionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollectionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all raster collections readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasterCollections.list",
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "processingStatus": {
	//       "enum": [
	//         "complete",
	//         "failed",
	//         "notReady",
	//         "processing",
	//         "ready"
	//       ],
	//       "enumDescriptions": [
	//         "The raster collection has completed processing.",
	//         "The raster collection has failed processing.",
	//         "The raster collection is not ready for processing.",
	//         "The raster collection is processing.",
	//         "The raster collection is ready for processing."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections",
	//   "response": {
	//     "$ref": "RasterCollectionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.patch":

type RasterCollectionsPatchCall struct {
	s                *Service
	id               string
	rastercollection *RasterCollection
	opt_             url.Values
	ctx_             context.Context
}

// Patch: Mutate a raster collection asset.
func (r *RasterCollectionsService) Patch(id string, rastercollection *RasterCollection) *RasterCollectionsPatchCall {
	c := &RasterCollectionsPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.rastercollection = rastercollection
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsPatchCall) Fields(s ...googleapi.Field) *RasterCollectionsPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsPatchCall) Context(ctx context.Context) *RasterCollectionsPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rastercollection)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsPatchCall) Do() error {
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
	//   "description": "Mutate a raster collection asset.",
	//   "httpMethod": "PATCH",
	//   "id": "mapsengine.rasterCollections.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}",
	//   "request": {
	//     "$ref": "RasterCollection"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.process":

type RasterCollectionsProcessCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Process: Process a raster collection asset.
func (r *RasterCollectionsService) Process(id string) *RasterCollectionsProcessCall {
	c := &RasterCollectionsProcessCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsProcessCall) Fields(s ...googleapi.Field) *RasterCollectionsProcessCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsProcessCall) Context(ctx context.Context) *RasterCollectionsProcessCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsProcessCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/process")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsProcessCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Process a raster collection asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.process",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/process",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.parents.list":

type RasterCollectionsParentsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all parent ids of the specified raster collection.
func (r *RasterCollectionsParentsService) List(id string) *RasterCollectionsParentsListCall {
	c := &RasterCollectionsParentsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *RasterCollectionsParentsListCall) MaxResults(maxResults int64) *RasterCollectionsParentsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *RasterCollectionsParentsListCall) PageToken(pageToken string) *RasterCollectionsParentsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsParentsListCall) Fields(s ...googleapi.Field) *RasterCollectionsParentsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsParentsListCall) Context(ctx context.Context) *RasterCollectionsParentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsParentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/parents")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsParentsListCall) Do() (*ParentsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ParentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all parent ids of the specified raster collection.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasterCollections.parents.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection whose parents will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/parents",
	//   "response": {
	//     "$ref": "ParentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.permissions.batchDelete":

type RasterCollectionsPermissionsBatchDeleteCall struct {
	s                             *Service
	id                            string
	permissionsbatchdeleterequest *PermissionsBatchDeleteRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchDelete: Remove permission entries from an already existing
// asset.
func (r *RasterCollectionsPermissionsService) BatchDelete(id string, permissionsbatchdeleterequest *PermissionsBatchDeleteRequest) *RasterCollectionsPermissionsBatchDeleteCall {
	c := &RasterCollectionsPermissionsBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchdeleterequest = permissionsbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsPermissionsBatchDeleteCall) Fields(s ...googleapi.Field) *RasterCollectionsPermissionsBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsPermissionsBatchDeleteCall) Context(ctx context.Context) *RasterCollectionsPermissionsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsPermissionsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/permissions/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsPermissionsBatchDeleteCall) Do() (*PermissionsBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove permission entries from an already existing asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.permissions.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset from which permissions will be removed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/permissions/batchDelete",
	//   "request": {
	//     "$ref": "PermissionsBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.permissions.batchUpdate":

type RasterCollectionsPermissionsBatchUpdateCall struct {
	s                             *Service
	id                            string
	permissionsbatchupdaterequest *PermissionsBatchUpdateRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchUpdate: Add or update permission entries to an already existing
// asset.
//
// An asset can hold up to 20 different permission entries. Each
// batchInsert request is atomic.
func (r *RasterCollectionsPermissionsService) BatchUpdate(id string, permissionsbatchupdaterequest *PermissionsBatchUpdateRequest) *RasterCollectionsPermissionsBatchUpdateCall {
	c := &RasterCollectionsPermissionsBatchUpdateCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchupdaterequest = permissionsbatchupdaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsPermissionsBatchUpdateCall) Fields(s ...googleapi.Field) *RasterCollectionsPermissionsBatchUpdateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsPermissionsBatchUpdateCall) Context(ctx context.Context) *RasterCollectionsPermissionsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsPermissionsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchupdaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/permissions/batchUpdate")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsPermissionsBatchUpdateCall) Do() (*PermissionsBatchUpdateResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add or update permission entries to an already existing asset.\n\nAn asset can hold up to 20 different permission entries. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.permissions.batchUpdate",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset to which permissions will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/permissions/batchUpdate",
	//   "request": {
	//     "$ref": "PermissionsBatchUpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchUpdateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.permissions.list":

type RasterCollectionsPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *RasterCollectionsPermissionsService) List(id string) *RasterCollectionsPermissionsListCall {
	c := &RasterCollectionsPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsPermissionsListCall) Fields(s ...googleapi.Field) *RasterCollectionsPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsPermissionsListCall) Context(ctx context.Context) *RasterCollectionsPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasterCollections.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.rasters.batchDelete":

type RasterCollectionsRastersBatchDeleteCall struct {
	s                                         *Service
	id                                        string
	rastercollectionsrasterbatchdeleterequest *RasterCollectionsRasterBatchDeleteRequest
	opt_                                      url.Values
	ctx_                                      context.Context
}

// BatchDelete: Remove rasters from an existing raster collection.
//
// Up to 50 rasters can be included in a single batchDelete request.
// Each batchDelete request is atomic.
func (r *RasterCollectionsRastersService) BatchDelete(id string, rastercollectionsrasterbatchdeleterequest *RasterCollectionsRasterBatchDeleteRequest) *RasterCollectionsRastersBatchDeleteCall {
	c := &RasterCollectionsRastersBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.rastercollectionsrasterbatchdeleterequest = rastercollectionsrasterbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsRastersBatchDeleteCall) Fields(s ...googleapi.Field) *RasterCollectionsRastersBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsRastersBatchDeleteCall) Context(ctx context.Context) *RasterCollectionsRastersBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsRastersBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rastercollectionsrasterbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/rasters/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsRastersBatchDeleteCall) Do() (*RasterCollectionsRastersBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollectionsRastersBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove rasters from an existing raster collection.\n\nUp to 50 rasters can be included in a single batchDelete request. Each batchDelete request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.rasters.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection to which these rasters belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/rasters/batchDelete",
	//   "request": {
	//     "$ref": "RasterCollectionsRasterBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "RasterCollectionsRastersBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.rasters.batchInsert":

type RasterCollectionsRastersBatchInsertCall struct {
	s                                          *Service
	id                                         string
	rastercollectionsrastersbatchinsertrequest *RasterCollectionsRastersBatchInsertRequest
	opt_                                       url.Values
	ctx_                                       context.Context
}

// BatchInsert: Add rasters to an existing raster collection. Rasters
// must be successfully processed in order to be added to a raster
// collection.
//
// Up to 50 rasters can be included in a single batchInsert request.
// Each batchInsert request is atomic.
func (r *RasterCollectionsRastersService) BatchInsert(id string, rastercollectionsrastersbatchinsertrequest *RasterCollectionsRastersBatchInsertRequest) *RasterCollectionsRastersBatchInsertCall {
	c := &RasterCollectionsRastersBatchInsertCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.rastercollectionsrastersbatchinsertrequest = rastercollectionsrastersbatchinsertrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsRastersBatchInsertCall) Fields(s ...googleapi.Field) *RasterCollectionsRastersBatchInsertCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsRastersBatchInsertCall) Context(ctx context.Context) *RasterCollectionsRastersBatchInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsRastersBatchInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rastercollectionsrastersbatchinsertrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/rasters/batchInsert")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsRastersBatchInsertCall) Do() (*RasterCollectionsRastersBatchInsertResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollectionsRastersBatchInsertResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add rasters to an existing raster collection. Rasters must be successfully processed in order to be added to a raster collection.\n\nUp to 50 rasters can be included in a single batchInsert request. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasterCollections.rasters.batchInsert",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster collection to which these rasters belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/rasters/batchInsert",
	//   "request": {
	//     "$ref": "RasterCollectionsRastersBatchInsertRequest"
	//   },
	//   "response": {
	//     "$ref": "RasterCollectionsRastersBatchInsertResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasterCollections.rasters.list":

type RasterCollectionsRastersListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all rasters within a raster collection.
func (r *RasterCollectionsRastersService) List(id string) *RasterCollectionsRastersListCall {
	c := &RasterCollectionsRastersListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *RasterCollectionsRastersListCall) Bbox(bbox string) *RasterCollectionsRastersListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *RasterCollectionsRastersListCall) CreatedAfter(createdAfter string) *RasterCollectionsRastersListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *RasterCollectionsRastersListCall) CreatedBefore(createdBefore string) *RasterCollectionsRastersListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *RasterCollectionsRastersListCall) CreatorEmail(creatorEmail string) *RasterCollectionsRastersListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *RasterCollectionsRastersListCall) MaxResults(maxResults int64) *RasterCollectionsRastersListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *RasterCollectionsRastersListCall) ModifiedAfter(modifiedAfter string) *RasterCollectionsRastersListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *RasterCollectionsRastersListCall) ModifiedBefore(modifiedBefore string) *RasterCollectionsRastersListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *RasterCollectionsRastersListCall) PageToken(pageToken string) *RasterCollectionsRastersListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *RasterCollectionsRastersListCall) Role(role string) *RasterCollectionsRastersListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *RasterCollectionsRastersListCall) Search(search string) *RasterCollectionsRastersListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *RasterCollectionsRastersListCall) Tags(tags string) *RasterCollectionsRastersListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RasterCollectionsRastersListCall) Fields(s ...googleapi.Field) *RasterCollectionsRastersListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RasterCollectionsRastersListCall) Context(ctx context.Context) *RasterCollectionsRastersListCall {
	c.ctx_ = ctx
	return c
}

func (c *RasterCollectionsRastersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasterCollections/{id}/rasters")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RasterCollectionsRastersListCall) Do() (*RasterCollectionsRastersListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RasterCollectionsRastersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all rasters within a raster collection.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasterCollections.rasters.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The ID of the raster collection to which these rasters belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasterCollections/{id}/rasters",
	//   "response": {
	//     "$ref": "RasterCollectionsRastersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasters.delete":

type RastersDeleteCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Delete: Delete a raster.
func (r *RastersService) Delete(id string) *RastersDeleteCall {
	c := &RastersDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersDeleteCall) Fields(s ...googleapi.Field) *RastersDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersDeleteCall) Context(ctx context.Context) *RastersDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersDeleteCall) Do() error {
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
	//   "description": "Delete a raster.",
	//   "httpMethod": "DELETE",
	//   "id": "mapsengine.rasters.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster. Only the raster creator or project owner are permitted to delete. If the raster is included in a layer or mosaic, the request will fail. Remove it from all parents prior to deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.get":

type RastersGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a single raster.
func (r *RastersService) Get(id string) *RastersGetCall {
	c := &RastersGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersGetCall) Fields(s ...googleapi.Field) *RastersGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersGetCall) Context(ctx context.Context) *RastersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersGetCall) Do() (*Raster, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Raster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a single raster.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasters.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}",
	//   "response": {
	//     "$ref": "Raster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasters.list":

type RastersListCall struct {
	s         *Service
	projectId string
	opt_      url.Values
	ctx_      context.Context
}

// List: Return all rasters readable by the current user.
func (r *RastersService) List(projectId string) *RastersListCall {
	c := &RastersListCall{s: r.s, opt_: urlValues()}
	c.projectId = projectId
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *RastersListCall) Bbox(bbox string) *RastersListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *RastersListCall) CreatedAfter(createdAfter string) *RastersListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *RastersListCall) CreatedBefore(createdBefore string) *RastersListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *RastersListCall) CreatorEmail(creatorEmail string) *RastersListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *RastersListCall) MaxResults(maxResults int64) *RastersListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *RastersListCall) ModifiedAfter(modifiedAfter string) *RastersListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *RastersListCall) ModifiedBefore(modifiedBefore string) *RastersListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *RastersListCall) PageToken(pageToken string) *RastersListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProcessingStatus sets the optional parameter "processingStatus":
//
// Possible values:
//   "complete" - The raster has completed processing.
//   "failed" - The raster has failed processing.
//   "notReady" - The raster is not ready for processing.
//   "processing" - The raster is processing.
//   "ready" - The raster is ready for processing.
func (c *RastersListCall) ProcessingStatus(processingStatus string) *RastersListCall {
	c.opt_.Set("processingStatus", fmt.Sprintf("%v", processingStatus))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *RastersListCall) Role(role string) *RastersListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *RastersListCall) Search(search string) *RastersListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *RastersListCall) Tags(tags string) *RastersListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersListCall) Fields(s ...googleapi.Field) *RastersListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersListCall) Context(ctx context.Context) *RastersListCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	c.opt_.Set("projectId", fmt.Sprintf("%v", c.projectId))
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersListCall) Do() (*RastersListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *RastersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all rasters readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasters.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "processingStatus": {
	//       "enum": [
	//         "complete",
	//         "failed",
	//         "notReady",
	//         "processing",
	//         "ready"
	//       ],
	//       "enumDescriptions": [
	//         "The raster has completed processing.",
	//         "The raster has failed processing.",
	//         "The raster is not ready for processing.",
	//         "The raster is processing.",
	//         "The raster is ready for processing."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters",
	//   "response": {
	//     "$ref": "RastersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasters.patch":

type RastersPatchCall struct {
	s      *Service
	id     string
	raster *Raster
	opt_   url.Values
	ctx_   context.Context
}

// Patch: Mutate a raster asset.
func (r *RastersService) Patch(id string, raster *Raster) *RastersPatchCall {
	c := &RastersPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.raster = raster
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersPatchCall) Fields(s ...googleapi.Field) *RastersPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersPatchCall) Context(ctx context.Context) *RastersPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.raster)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersPatchCall) Do() error {
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
	//   "description": "Mutate a raster asset.",
	//   "httpMethod": "PATCH",
	//   "id": "mapsengine.rasters.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}",
	//   "request": {
	//     "$ref": "Raster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.process":

type RastersProcessCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Process: Process a raster asset.
func (r *RastersService) Process(id string) *RastersProcessCall {
	c := &RastersProcessCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersProcessCall) Fields(s ...googleapi.Field) *RastersProcessCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersProcessCall) Context(ctx context.Context) *RastersProcessCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersProcessCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/process")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersProcessCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Process a raster asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasters.process",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the raster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/process",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.upload":

type RastersUploadCall struct {
	s      *Service
	raster *Raster
	opt_   url.Values
	ctx_   context.Context
}

// Upload: Create a skeleton raster asset for upload.
func (r *RastersService) Upload(raster *Raster) *RastersUploadCall {
	c := &RastersUploadCall{s: r.s, opt_: urlValues()}
	c.raster = raster
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersUploadCall) Fields(s ...googleapi.Field) *RastersUploadCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersUploadCall) Context(ctx context.Context) *RastersUploadCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersUploadCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.raster)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/upload")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersUploadCall) Do() (*Raster, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Raster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a skeleton raster asset for upload.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasters.upload",
	//   "path": "rasters/upload",
	//   "request": {
	//     "$ref": "Raster"
	//   },
	//   "response": {
	//     "$ref": "Raster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.files.insert":

type RastersFilesInsertCall struct {
	s                *Service
	id               string
	filename         string
	opt_             url.Values
	media_           io.Reader
	resumable_       googleapi.SizeReaderAt
	mediaType_       string
	protocol_        string
	progressUpdater_ googleapi.ProgressUpdater
	ctx_             context.Context
}

// Insert: Upload a file to a raster asset.
func (r *RastersFilesService) Insert(id string, filename string) *RastersFilesInsertCall {
	c := &RastersFilesInsertCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.filename = filename
	return c
}

// Media specifies the media to upload in a single chunk. At most one of
// Media and ResumableMedia may be set.
func (c *RastersFilesInsertCall) Media(r io.Reader) *RastersFilesInsertCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx. At most one of Media and ResumableMedia may be
// set. mediaType identifies the MIME media type of the upload, such as
// "image/png". If mediaType is "", it will be auto-detected. The
// provided ctx will supersede any context previously provided to the
// Context method.
func (c *RastersFilesInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *RastersFilesInsertCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *RastersFilesInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *RastersFilesInsertCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersFilesInsertCall) Fields(s ...googleapi.Field) *RastersFilesInsertCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *RastersFilesInsertCall) Context(ctx context.Context) *RastersFilesInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersFilesInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	c.opt_.Set("filename", fmt.Sprintf("%v", c.filename))
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/files")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.opt_.Set("uploadType", c.protocol_)
	}
	urls += "?" + c.opt_.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersFilesInsertCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      c.progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}
	return nil
	// {
	//   "description": "Upload a file to a raster asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasters.files.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "10GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mapsengine/exp2/rasters/{id}/files"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mapsengine/exp2/rasters/{id}/files"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "id",
	//     "filename"
	//   ],
	//   "parameters": {
	//     "filename": {
	//       "description": "The file name of this uploaded file.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The ID of the raster asset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/files",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mapsengine.rasters.parents.list":

type RastersParentsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all parent ids of the specified rasters.
func (r *RastersParentsService) List(id string) *RastersParentsListCall {
	c := &RastersParentsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *RastersParentsListCall) MaxResults(maxResults int64) *RastersParentsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *RastersParentsListCall) PageToken(pageToken string) *RastersParentsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersParentsListCall) Fields(s ...googleapi.Field) *RastersParentsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersParentsListCall) Context(ctx context.Context) *RastersParentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersParentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/parents")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersParentsListCall) Do() (*ParentsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ParentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all parent ids of the specified rasters.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasters.parents.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the rasters whose parents will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/parents",
	//   "response": {
	//     "$ref": "ParentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.rasters.permissions.batchDelete":

type RastersPermissionsBatchDeleteCall struct {
	s                             *Service
	id                            string
	permissionsbatchdeleterequest *PermissionsBatchDeleteRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchDelete: Remove permission entries from an already existing
// asset.
func (r *RastersPermissionsService) BatchDelete(id string, permissionsbatchdeleterequest *PermissionsBatchDeleteRequest) *RastersPermissionsBatchDeleteCall {
	c := &RastersPermissionsBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchdeleterequest = permissionsbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersPermissionsBatchDeleteCall) Fields(s ...googleapi.Field) *RastersPermissionsBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersPermissionsBatchDeleteCall) Context(ctx context.Context) *RastersPermissionsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersPermissionsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/permissions/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersPermissionsBatchDeleteCall) Do() (*PermissionsBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove permission entries from an already existing asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasters.permissions.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset from which permissions will be removed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/permissions/batchDelete",
	//   "request": {
	//     "$ref": "PermissionsBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.permissions.batchUpdate":

type RastersPermissionsBatchUpdateCall struct {
	s                             *Service
	id                            string
	permissionsbatchupdaterequest *PermissionsBatchUpdateRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchUpdate: Add or update permission entries to an already existing
// asset.
//
// An asset can hold up to 20 different permission entries. Each
// batchInsert request is atomic.
func (r *RastersPermissionsService) BatchUpdate(id string, permissionsbatchupdaterequest *PermissionsBatchUpdateRequest) *RastersPermissionsBatchUpdateCall {
	c := &RastersPermissionsBatchUpdateCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchupdaterequest = permissionsbatchupdaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersPermissionsBatchUpdateCall) Fields(s ...googleapi.Field) *RastersPermissionsBatchUpdateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersPermissionsBatchUpdateCall) Context(ctx context.Context) *RastersPermissionsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersPermissionsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchupdaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/permissions/batchUpdate")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersPermissionsBatchUpdateCall) Do() (*PermissionsBatchUpdateResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add or update permission entries to an already existing asset.\n\nAn asset can hold up to 20 different permission entries. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.rasters.permissions.batchUpdate",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset to which permissions will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/permissions/batchUpdate",
	//   "request": {
	//     "$ref": "PermissionsBatchUpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchUpdateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.rasters.permissions.list":

type RastersPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *RastersPermissionsService) List(id string) *RastersPermissionsListCall {
	c := &RastersPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RastersPermissionsListCall) Fields(s ...googleapi.Field) *RastersPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RastersPermissionsListCall) Context(ctx context.Context) *RastersPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *RastersPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "rasters/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *RastersPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.rasters.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "rasters/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.create":

type TablesCreateCall struct {
	s     *Service
	table *Table
	opt_  url.Values
	ctx_  context.Context
}

// Create: Create a table asset.
func (r *TablesService) Create(table *Table) *TablesCreateCall {
	c := &TablesCreateCall{s: r.s, opt_: urlValues()}
	c.table = table
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesCreateCall) Fields(s ...googleapi.Field) *TablesCreateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesCreateCall) Context(ctx context.Context) *TablesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesCreateCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a table asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.create",
	//   "path": "tables",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.delete":

type TablesDeleteCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Delete: Delete a table.
func (r *TablesService) Delete(id string) *TablesDeleteCall {
	c := &TablesDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesDeleteCall) Fields(s ...googleapi.Field) *TablesDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesDeleteCall) Context(ctx context.Context) *TablesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesDeleteCall) Do() error {
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
	//   "description": "Delete a table.",
	//   "httpMethod": "DELETE",
	//   "id": "mapsengine.tables.delete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table. Only the table creator or project owner are permitted to delete. If the table is included in a layer, the request will fail. Remove it from all layers prior to deleting.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.get":

type TablesGetCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Get: Return metadata for a particular table, including the schema.
func (r *TablesService) Get(id string) *TablesGetCall {
	c := &TablesGetCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Version sets the optional parameter "version":
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *TablesGetCall) Version(version string) *TablesGetCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesGetCall) Fields(s ...googleapi.Field) *TablesGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesGetCall) Context(ctx context.Context) *TablesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesGetCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return metadata for a particular table, including the schema.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.get",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.list":

type TablesListCall struct {
	s    *Service
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all tables readable by the current user.
func (r *TablesService) List() *TablesListCall {
	c := &TablesListCall{s: r.s, opt_: urlValues()}
	return c
}

// Bbox sets the optional parameter "bbox": A bounding box, expressed as
// "west,south,east,north". If set, only assets which intersect this
// bounding box will be returned.
func (c *TablesListCall) Bbox(bbox string) *TablesListCall {
	c.opt_.Set("bbox", fmt.Sprintf("%v", bbox))
	return c
}

// CreatedAfter sets the optional parameter "createdAfter": An RFC 3339
// formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or after this time.
func (c *TablesListCall) CreatedAfter(createdAfter string) *TablesListCall {
	c.opt_.Set("createdAfter", fmt.Sprintf("%v", createdAfter))
	return c
}

// CreatedBefore sets the optional parameter "createdBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been created at or before this time.
func (c *TablesListCall) CreatedBefore(createdBefore string) *TablesListCall {
	c.opt_.Set("createdBefore", fmt.Sprintf("%v", createdBefore))
	return c
}

// CreatorEmail sets the optional parameter "creatorEmail": An email
// address representing a user. Returned assets that have been created
// by the user associated with the provided email address.
func (c *TablesListCall) CreatorEmail(creatorEmail string) *TablesListCall {
	c.opt_.Set("creatorEmail", fmt.Sprintf("%v", creatorEmail))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 100.
func (c *TablesListCall) MaxResults(maxResults int64) *TablesListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// ModifiedAfter sets the optional parameter "modifiedAfter": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or after this time.
func (c *TablesListCall) ModifiedAfter(modifiedAfter string) *TablesListCall {
	c.opt_.Set("modifiedAfter", fmt.Sprintf("%v", modifiedAfter))
	return c
}

// ModifiedBefore sets the optional parameter "modifiedBefore": An RFC
// 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned
// assets will have been modified at or before this time.
func (c *TablesListCall) ModifiedBefore(modifiedBefore string) *TablesListCall {
	c.opt_.Set("modifiedBefore", fmt.Sprintf("%v", modifiedBefore))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *TablesListCall) PageToken(pageToken string) *TablesListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// ProcessingStatus sets the optional parameter "processingStatus":
//
// Possible values:
//   "complete" - The table has completed processing.
//   "failed" - The table has failed processing.
//   "notReady" - The table is not ready for processing.
//   "processing" - The table is processing.
//   "ready" - The table is ready for processing.
func (c *TablesListCall) ProcessingStatus(processingStatus string) *TablesListCall {
	c.opt_.Set("processingStatus", fmt.Sprintf("%v", processingStatus))
	return c
}

// ProjectId sets the optional parameter "projectId": The ID of a Maps
// Engine project, used to filter the response. To list all available
// projects with their IDs, send a Projects: list request. You can also
// find your project ID as the value of the DashboardPlace:cid URL
// parameter when signed in to mapsengine.google.com.
func (c *TablesListCall) ProjectId(projectId string) *TablesListCall {
	c.opt_.Set("projectId", fmt.Sprintf("%v", projectId))
	return c
}

// Role sets the optional parameter "role": The role parameter indicates
// that the response should only contain assets where the current user
// has the specified level of access.
//
// Possible values:
//   "owner" - The user can read, write and administer the asset.
//   "reader" - The user can read the asset.
//   "writer" - The user can read and write the asset.
func (c *TablesListCall) Role(role string) *TablesListCall {
	c.opt_.Set("role", fmt.Sprintf("%v", role))
	return c
}

// Search sets the optional parameter "search": An unstructured search
// string used to filter the set of results based on asset metadata.
func (c *TablesListCall) Search(search string) *TablesListCall {
	c.opt_.Set("search", fmt.Sprintf("%v", search))
	return c
}

// Tags sets the optional parameter "tags": A comma separated list of
// tags. Returned assets will contain all the tags from the list.
func (c *TablesListCall) Tags(tags string) *TablesListCall {
	c.opt_.Set("tags", fmt.Sprintf("%v", tags))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesListCall) Fields(s ...googleapi.Field) *TablesListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesListCall) Context(ctx context.Context) *TablesListCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesListCall) Do() (*TablesListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TablesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all tables readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.list",
	//   "parameters": {
	//     "bbox": {
	//       "description": "A bounding box, expressed as \"west,south,east,north\". If set, only assets which intersect this bounding box will be returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "createdBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been created at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "creatorEmail": {
	//       "description": "An email address representing a user. Returned assets that have been created by the user associated with the provided email address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 100.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "modifiedAfter": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or after this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modifiedBefore": {
	//       "description": "An RFC 3339 formatted date-time value (e.g. 1970-01-01T00:00:00Z). Returned assets will have been modified at or before this time.",
	//       "format": "date-time",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "processingStatus": {
	//       "enum": [
	//         "complete",
	//         "failed",
	//         "notReady",
	//         "processing",
	//         "ready"
	//       ],
	//       "enumDescriptions": [
	//         "The table has completed processing.",
	//         "The table has failed processing.",
	//         "The table is not ready for processing.",
	//         "The table is processing.",
	//         "The table is ready for processing."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of a Maps Engine project, used to filter the response. To list all available projects with their IDs, send a Projects: list request. You can also find your project ID as the value of the DashboardPlace:cid URL parameter when signed in to mapsengine.google.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "role": {
	//       "description": "The role parameter indicates that the response should only contain assets where the current user has the specified level of access.",
	//       "enum": [
	//         "owner",
	//         "reader",
	//         "writer"
	//       ],
	//       "enumDescriptions": [
	//         "The user can read, write and administer the asset.",
	//         "The user can read the asset.",
	//         "The user can read and write the asset."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "search": {
	//       "description": "An unstructured search string used to filter the set of results based on asset metadata.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tags": {
	//       "description": "A comma separated list of tags. Returned assets will contain all the tags from the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables",
	//   "response": {
	//     "$ref": "TablesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.patch":

type TablesPatchCall struct {
	s     *Service
	id    string
	table *Table
	opt_  url.Values
	ctx_  context.Context
}

// Patch: Mutate a table asset.
func (r *TablesService) Patch(id string, table *Table) *TablesPatchCall {
	c := &TablesPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.table = table
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesPatchCall) Fields(s ...googleapi.Field) *TablesPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesPatchCall) Context(ctx context.Context) *TablesPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesPatchCall) Do() error {
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
	//   "description": "Mutate a table asset.",
	//   "httpMethod": "PATCH",
	//   "id": "mapsengine.tables.patch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.process":

type TablesProcessCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// Process: Process a table asset.
func (r *TablesService) Process(id string) *TablesProcessCall {
	c := &TablesProcessCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesProcessCall) Fields(s ...googleapi.Field) *TablesProcessCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesProcessCall) Context(ctx context.Context) *TablesProcessCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesProcessCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/process")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesProcessCall) Do() (*ProcessResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ProcessResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Process a table asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.process",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/process",
	//   "response": {
	//     "$ref": "ProcessResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.upload":

type TablesUploadCall struct {
	s     *Service
	table *Table
	opt_  url.Values
	ctx_  context.Context
}

// Upload: Create a placeholder table asset to which table files can be
// uploaded.
// Once the placeholder has been created, files are uploaded to the
// https://www.googleapis.com/upload/mapsengine/v1/tables/table_id/files
// endpoint.
// See Table Upload in the Developer's Guide or Table.files: insert in
// the reference documentation for more information.
func (r *TablesService) Upload(table *Table) *TablesUploadCall {
	c := &TablesUploadCall{s: r.s, opt_: urlValues()}
	c.table = table
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesUploadCall) Fields(s ...googleapi.Field) *TablesUploadCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesUploadCall) Context(ctx context.Context) *TablesUploadCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesUploadCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.table)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/upload")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesUploadCall) Do() (*Table, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Table
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create a placeholder table asset to which table files can be uploaded.\nOnce the placeholder has been created, files are uploaded to the https://www.googleapis.com/upload/mapsengine/v1/tables/table_id/files endpoint.\nSee Table Upload in the Developer's Guide or Table.files: insert in the reference documentation for more information.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.upload",
	//   "path": "tables/upload",
	//   "request": {
	//     "$ref": "Table"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.features.batchDelete":

type TablesFeaturesBatchDeleteCall struct {
	s                          *Service
	id                         string
	featuresbatchdeleterequest *FeaturesBatchDeleteRequest
	opt_                       url.Values
	ctx_                       context.Context
}

// BatchDelete: Delete all features matching the given IDs.
func (r *TablesFeaturesService) BatchDelete(id string, featuresbatchdeleterequest *FeaturesBatchDeleteRequest) *TablesFeaturesBatchDeleteCall {
	c := &TablesFeaturesBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.featuresbatchdeleterequest = featuresbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFeaturesBatchDeleteCall) Fields(s ...googleapi.Field) *TablesFeaturesBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesFeaturesBatchDeleteCall) Context(ctx context.Context) *TablesFeaturesBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFeaturesBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.featuresbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/features/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFeaturesBatchDeleteCall) Do() error {
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
	//   "description": "Delete all features matching the given IDs.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.features.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table that contains the features to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/features/batchDelete",
	//   "request": {
	//     "$ref": "FeaturesBatchDeleteRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.features.batchInsert":

type TablesFeaturesBatchInsertCall struct {
	s                          *Service
	id                         string
	featuresbatchinsertrequest *FeaturesBatchInsertRequest
	opt_                       url.Values
	ctx_                       context.Context
}

// BatchInsert: Append features to an existing table.
//
// A single batchInsert request can create:
//
// - Up to 50 features.
// - A combined total of 10 000 vertices.
// Feature limits are documented in the Supported data formats and
// limits article of the Google Maps Engine help center. Note that free
// and paid accounts have different limits.
//
// For more information about inserting features, read Creating features
// in the Google Maps Engine developer's guide.
func (r *TablesFeaturesService) BatchInsert(id string, featuresbatchinsertrequest *FeaturesBatchInsertRequest) *TablesFeaturesBatchInsertCall {
	c := &TablesFeaturesBatchInsertCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.featuresbatchinsertrequest = featuresbatchinsertrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFeaturesBatchInsertCall) Fields(s ...googleapi.Field) *TablesFeaturesBatchInsertCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesFeaturesBatchInsertCall) Context(ctx context.Context) *TablesFeaturesBatchInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFeaturesBatchInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.featuresbatchinsertrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/features/batchInsert")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFeaturesBatchInsertCall) Do() error {
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
	//   "description": "Append features to an existing table.\n\nA single batchInsert request can create:\n\n- Up to 50 features.\n- A combined total of 10 000 vertices.\nFeature limits are documented in the Supported data formats and limits article of the Google Maps Engine help center. Note that free and paid accounts have different limits.\n\nFor more information about inserting features, read Creating features in the Google Maps Engine developer's guide.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.features.batchInsert",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table to append the features to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/features/batchInsert",
	//   "request": {
	//     "$ref": "FeaturesBatchInsertRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.features.batchPatch":

type TablesFeaturesBatchPatchCall struct {
	s                         *Service
	id                        string
	featuresbatchpatchrequest *FeaturesBatchPatchRequest
	opt_                      url.Values
	ctx_                      context.Context
}

// BatchPatch: Update the supplied features.
//
// A single batchPatch request can update:
//
// - Up to 50 features.
// - A combined total of 10 000 vertices.
// Feature limits are documented in the Supported data formats and
// limits article of the Google Maps Engine help center. Note that free
// and paid accounts have different limits.
//
// Feature updates use HTTP PATCH semantics:
//
// - A supplied value replaces an existing value (if any) in that
// field.
// - Omitted fields remain unchanged.
// - Complex values in geometries and properties must be replaced as
// atomic units. For example, providing just the coordinates of a
// geometry is not allowed; the complete geometry, including type, must
// be supplied.
// - Setting a property's value to null deletes that property.
// For more information about updating features, read Updating features
// in the Google Maps Engine developer's guide.
func (r *TablesFeaturesService) BatchPatch(id string, featuresbatchpatchrequest *FeaturesBatchPatchRequest) *TablesFeaturesBatchPatchCall {
	c := &TablesFeaturesBatchPatchCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.featuresbatchpatchrequest = featuresbatchpatchrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFeaturesBatchPatchCall) Fields(s ...googleapi.Field) *TablesFeaturesBatchPatchCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesFeaturesBatchPatchCall) Context(ctx context.Context) *TablesFeaturesBatchPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFeaturesBatchPatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.featuresbatchpatchrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/features/batchPatch")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFeaturesBatchPatchCall) Do() error {
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
	//   "description": "Update the supplied features.\n\nA single batchPatch request can update:\n\n- Up to 50 features.\n- A combined total of 10 000 vertices.\nFeature limits are documented in the Supported data formats and limits article of the Google Maps Engine help center. Note that free and paid accounts have different limits.\n\nFeature updates use HTTP PATCH semantics:\n\n- A supplied value replaces an existing value (if any) in that field.\n- Omitted fields remain unchanged.\n- Complex values in geometries and properties must be replaced as atomic units. For example, providing just the coordinates of a geometry is not allowed; the complete geometry, including type, must be supplied.\n- Setting a property's value to null deletes that property.\nFor more information about updating features, read Updating features in the Google Maps Engine developer's guide.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.features.batchPatch",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table containing the features to be patched.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/features/batchPatch",
	//   "request": {
	//     "$ref": "FeaturesBatchPatchRequest"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.features.get":

type TablesFeaturesGetCall struct {
	s       *Service
	tableId string
	id      string
	opt_    url.Values
	ctx_    context.Context
}

// Get: Return a single feature, given its ID.
func (r *TablesFeaturesService) Get(tableId string, id string) *TablesFeaturesGetCall {
	c := &TablesFeaturesGetCall{s: r.s, opt_: urlValues()}
	c.tableId = tableId
	c.id = id
	return c
}

// Select sets the optional parameter "select": A SQL-like projection
// clause used to specify returned properties. If this parameter is not
// included, all properties are returned.
func (c *TablesFeaturesGetCall) Select(select_ string) *TablesFeaturesGetCall {
	c.opt_.Set("select_", fmt.Sprintf("%v", select_))
	return c
}

// Version sets the optional parameter "version": The table version to
// access. See Accessing Public Data for information.
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *TablesFeaturesGetCall) Version(version string) *TablesFeaturesGetCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFeaturesGetCall) Fields(s ...googleapi.Field) *TablesFeaturesGetCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesFeaturesGetCall) Context(ctx context.Context) *TablesFeaturesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFeaturesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{tableId}/features/{id}")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"tableId": c.tableId,
		"id":      c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFeaturesGetCall) Do() (*Feature, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Feature
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return a single feature, given its ID.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.features.get",
	//   "parameterOrder": [
	//     "tableId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the feature to get.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "select": {
	//       "description": "A SQL-like projection clause used to specify returned properties. If this parameter is not included, all properties are returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "tableId": {
	//       "description": "The ID of the table.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "The table version to access. See Accessing Public Data for information.",
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{tableId}/features/{id}",
	//   "response": {
	//     "$ref": "Feature"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.features.list":

type TablesFeaturesListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all features readable by the current user.
func (r *TablesFeaturesService) List(id string) *TablesFeaturesListCall {
	c := &TablesFeaturesListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Include sets the optional parameter "include": A comma separated list
// of optional data to include. Optional data available: schema.
func (c *TablesFeaturesListCall) Include(include string) *TablesFeaturesListCall {
	c.opt_.Set("include", fmt.Sprintf("%v", include))
	return c
}

// Intersects sets the optional parameter "intersects": A geometry
// literal that specifies the spatial restriction of the query.
func (c *TablesFeaturesListCall) Intersects(intersects string) *TablesFeaturesListCall {
	c.opt_.Set("intersects", fmt.Sprintf("%v", intersects))
	return c
}

// Limit sets the optional parameter "limit": The total number of
// features to return from the query, irrespective of the number of
// pages.
func (c *TablesFeaturesListCall) Limit(limit int64) *TablesFeaturesListCall {
	c.opt_.Set("limit", fmt.Sprintf("%v", limit))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in the response, used for paging. The
// maximum supported value is 1000.
func (c *TablesFeaturesListCall) MaxResults(maxResults int64) *TablesFeaturesListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// OrderBy sets the optional parameter "orderBy": An SQL-like order by
// clause used to sort results. If this parameter is not included, the
// order of features is undefined.
func (c *TablesFeaturesListCall) OrderBy(orderBy string) *TablesFeaturesListCall {
	c.opt_.Set("orderBy", fmt.Sprintf("%v", orderBy))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *TablesFeaturesListCall) PageToken(pageToken string) *TablesFeaturesListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Select sets the optional parameter "select": A SQL-like projection
// clause used to specify returned properties. If this parameter is not
// included, all properties are returned.
func (c *TablesFeaturesListCall) Select(select_ string) *TablesFeaturesListCall {
	c.opt_.Set("select_", fmt.Sprintf("%v", select_))
	return c
}

// Version sets the optional parameter "version": The table version to
// access. See Accessing Public Data for information.
//
// Possible values:
//   "draft" - The draft version.
//   "published" - The published version.
func (c *TablesFeaturesListCall) Version(version string) *TablesFeaturesListCall {
	c.opt_.Set("version", fmt.Sprintf("%v", version))
	return c
}

// Where sets the optional parameter "where": An SQL-like predicate used
// to filter results.
func (c *TablesFeaturesListCall) Where(where string) *TablesFeaturesListCall {
	c.opt_.Set("where", fmt.Sprintf("%v", where))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFeaturesListCall) Fields(s ...googleapi.Field) *TablesFeaturesListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesFeaturesListCall) Context(ctx context.Context) *TablesFeaturesListCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFeaturesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/features")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFeaturesListCall) Do() (*FeaturesListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *FeaturesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all features readable by the current user.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.features.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table to which these features belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "include": {
	//       "description": "A comma separated list of optional data to include. Optional data available: schema.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "intersects": {
	//       "description": "A geometry literal that specifies the spatial restriction of the query.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "limit": {
	//       "description": "The total number of features to return from the query, irrespective of the number of pages.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in the response, used for paging. The maximum supported value is 1000.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "orderBy": {
	//       "description": "An SQL-like order by clause used to sort results. If this parameter is not included, the order of features is undefined.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "select": {
	//       "description": "A SQL-like projection clause used to specify returned properties. If this parameter is not included, all properties are returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "version": {
	//       "description": "The table version to access. See Accessing Public Data for information.",
	//       "enum": [
	//         "draft",
	//         "published"
	//       ],
	//       "enumDescriptions": [
	//         "The draft version.",
	//         "The published version."
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "where": {
	//       "description": "An SQL-like predicate used to filter results.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/features",
	//   "response": {
	//     "$ref": "FeaturesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.files.insert":

type TablesFilesInsertCall struct {
	s                *Service
	id               string
	filename         string
	opt_             url.Values
	media_           io.Reader
	resumable_       googleapi.SizeReaderAt
	mediaType_       string
	protocol_        string
	progressUpdater_ googleapi.ProgressUpdater
	ctx_             context.Context
}

// Insert: Upload a file to a placeholder table asset. See Table Upload
// in the Developer's Guide for more information.
// Supported file types are listed in the Supported data formats and
// limits article of the Google Maps Engine help center.
func (r *TablesFilesService) Insert(id string, filename string) *TablesFilesInsertCall {
	c := &TablesFilesInsertCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.filename = filename
	return c
}

// Media specifies the media to upload in a single chunk. At most one of
// Media and ResumableMedia may be set.
func (c *TablesFilesInsertCall) Media(r io.Reader) *TablesFilesInsertCall {
	c.media_ = r
	c.protocol_ = "multipart"
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx. At most one of Media and ResumableMedia may be
// set. mediaType identifies the MIME media type of the upload, such as
// "image/png". If mediaType is "", it will be auto-detected. The
// provided ctx will supersede any context previously provided to the
// Context method.
func (c *TablesFilesInsertCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *TablesFilesInsertCall {
	c.ctx_ = ctx
	c.resumable_ = io.NewSectionReader(r, 0, size)
	c.mediaType_ = mediaType
	c.protocol_ = "resumable"
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *TablesFilesInsertCall) ProgressUpdater(pu googleapi.ProgressUpdater) *TablesFilesInsertCall {
	c.progressUpdater_ = pu
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesFilesInsertCall) Fields(s ...googleapi.Field) *TablesFilesInsertCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *TablesFilesInsertCall) Context(ctx context.Context) *TablesFilesInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesFilesInsertCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	c.opt_.Set("filename", fmt.Sprintf("%v", c.filename))
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/files")
	if c.media_ != nil || c.resumable_ != nil {
		urls = strings.Replace(urls, "https://www.googleapis.com/", "https://www.googleapis.com/upload/", 1)
		c.opt_.Set("uploadType", c.protocol_)
	}
	urls += "?" + c.opt_.Encode()
	body = new(bytes.Buffer)
	ctype := "application/json"
	if c.protocol_ != "resumable" {
		var cancel func()
		cancel, _ = googleapi.ConditionallyIncludeMedia(c.media_, &body, &ctype)
		if cancel != nil {
			defer cancel()
		}
	}
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	if c.protocol_ == "resumable" {
		if c.mediaType_ == "" {
			c.mediaType_ = googleapi.DetectMediaType(c.resumable_)
		}
		req.Header.Set("X-Upload-Content-Type", c.mediaType_)
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	} else {
		req.Header.Set("Content-Type", ctype)
	}
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesFilesInsertCall) Do() error {
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	if c.protocol_ == "resumable" {
		loc := res.Header.Get("Location")
		rx := &googleapi.ResumableUpload{
			Client:        c.s.client,
			UserAgent:     c.s.userAgent(),
			URI:           loc,
			Media:         c.resumable_,
			MediaType:     c.mediaType_,
			ContentLength: c.resumable_.Size(),
			Callback:      c.progressUpdater_,
		}
		res, err = rx.Upload(c.ctx_)
		if err != nil {
			return err
		}
		defer res.Body.Close()
	}
	return nil
	// {
	//   "description": "Upload a file to a placeholder table asset. See Table Upload in the Developer's Guide for more information.\nSupported file types are listed in the Supported data formats and limits article of the Google Maps Engine help center.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.files.insert",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "1GB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/mapsengine/exp2/tables/{id}/files"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/mapsengine/exp2/tables/{id}/files"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "id",
	//     "filename"
	//   ],
	//   "parameters": {
	//     "filename": {
	//       "description": "The file name of this uploaded file.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "id": {
	//       "description": "The ID of the table asset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/files",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

// method id "mapsengine.tables.parents.list":

type TablesParentsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all parent ids of the specified table.
func (r *TablesParentsService) List(id string) *TablesParentsListCall {
	c := &TablesParentsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of items to include in a single response page. The maximum
// supported value is 50.
func (c *TablesParentsListCall) MaxResults(maxResults int64) *TablesParentsListCall {
	c.opt_.Set("maxResults", fmt.Sprintf("%v", maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, used to page through large result sets. To get the next page
// of results, set this parameter to the value of nextPageToken from the
// previous response.
func (c *TablesParentsListCall) PageToken(pageToken string) *TablesParentsListCall {
	c.opt_.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesParentsListCall) Fields(s ...googleapi.Field) *TablesParentsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesParentsListCall) Context(ctx context.Context) *TablesParentsListCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesParentsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/parents")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesParentsListCall) Do() (*ParentsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ParentsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all parent ids of the specified table.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.parents.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the table whose parents will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of items to include in a single response page. The maximum supported value is 50.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, used to page through large result sets. To get the next page of results, set this parameter to the value of nextPageToken from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/parents",
	//   "response": {
	//     "$ref": "ParentsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}

// method id "mapsengine.tables.permissions.batchDelete":

type TablesPermissionsBatchDeleteCall struct {
	s                             *Service
	id                            string
	permissionsbatchdeleterequest *PermissionsBatchDeleteRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchDelete: Remove permission entries from an already existing
// asset.
func (r *TablesPermissionsService) BatchDelete(id string, permissionsbatchdeleterequest *PermissionsBatchDeleteRequest) *TablesPermissionsBatchDeleteCall {
	c := &TablesPermissionsBatchDeleteCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchdeleterequest = permissionsbatchdeleterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesPermissionsBatchDeleteCall) Fields(s ...googleapi.Field) *TablesPermissionsBatchDeleteCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesPermissionsBatchDeleteCall) Context(ctx context.Context) *TablesPermissionsBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesPermissionsBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchdeleterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/permissions/batchDelete")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesPermissionsBatchDeleteCall) Do() (*PermissionsBatchDeleteResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchDeleteResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Remove permission entries from an already existing asset.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.permissions.batchDelete",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset from which permissions will be removed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/permissions/batchDelete",
	//   "request": {
	//     "$ref": "PermissionsBatchDeleteRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchDeleteResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.permissions.batchUpdate":

type TablesPermissionsBatchUpdateCall struct {
	s                             *Service
	id                            string
	permissionsbatchupdaterequest *PermissionsBatchUpdateRequest
	opt_                          url.Values
	ctx_                          context.Context
}

// BatchUpdate: Add or update permission entries to an already existing
// asset.
//
// An asset can hold up to 20 different permission entries. Each
// batchInsert request is atomic.
func (r *TablesPermissionsService) BatchUpdate(id string, permissionsbatchupdaterequest *PermissionsBatchUpdateRequest) *TablesPermissionsBatchUpdateCall {
	c := &TablesPermissionsBatchUpdateCall{s: r.s, opt_: urlValues()}
	c.id = id
	c.permissionsbatchupdaterequest = permissionsbatchupdaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesPermissionsBatchUpdateCall) Fields(s ...googleapi.Field) *TablesPermissionsBatchUpdateCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesPermissionsBatchUpdateCall) Context(ctx context.Context) *TablesPermissionsBatchUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesPermissionsBatchUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.permissionsbatchupdaterequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/permissions/batchUpdate")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesPermissionsBatchUpdateCall) Do() (*PermissionsBatchUpdateResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsBatchUpdateResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add or update permission entries to an already existing asset.\n\nAn asset can hold up to 20 different permission entries. Each batchInsert request is atomic.",
	//   "httpMethod": "POST",
	//   "id": "mapsengine.tables.permissions.batchUpdate",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset to which permissions will be added.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/permissions/batchUpdate",
	//   "request": {
	//     "$ref": "PermissionsBatchUpdateRequest"
	//   },
	//   "response": {
	//     "$ref": "PermissionsBatchUpdateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine"
	//   ]
	// }

}

// method id "mapsengine.tables.permissions.list":

type TablesPermissionsListCall struct {
	s    *Service
	id   string
	opt_ url.Values
	ctx_ context.Context
}

// List: Return all of the permissions for the specified asset.
func (r *TablesPermissionsService) List(id string) *TablesPermissionsListCall {
	c := &TablesPermissionsListCall{s: r.s, opt_: urlValues()}
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TablesPermissionsListCall) Fields(s ...googleapi.Field) *TablesPermissionsListCall {
	c.opt_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *TablesPermissionsListCall) Context(ctx context.Context) *TablesPermissionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *TablesPermissionsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.opt_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "tables/{id}/permissions")
	urls += "?" + c.opt_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"id": c.id,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

func (c *TablesPermissionsListCall) Do() (*PermissionsListResponse, error) {
	res, err := c.doRequest("json")
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *PermissionsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Return all of the permissions for the specified asset.",
	//   "httpMethod": "GET",
	//   "id": "mapsengine.tables.permissions.list",
	//   "parameterOrder": [
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the asset whose permissions will be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "tables/{id}/permissions",
	//   "response": {
	//     "$ref": "PermissionsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/mapsengine",
	//     "https://www.googleapis.com/auth/mapsengine.readonly"
	//   ]
	// }

}
