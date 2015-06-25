// Package prediction provides access to the Prediction API.
//
// See https://developers.google.com/prediction/docs/developer-guide
//
// Usage example:
//
//   import "google.golang.org/api/prediction/v1.3"
//   ...
//   predictionService, err := prediction.New(oauthHttpClient)
package prediction

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

const apiId = "prediction:v1.3"
const apiName = "prediction"
const apiVersion = "v1.3"
const basePath = "https://www.googleapis.com/prediction/v1.3/"

// OAuth2 scopes used by this API.
const (
	// Manage your data and permissions in Google Cloud Storage
	DevstorageFullControlScope = "https://www.googleapis.com/auth/devstorage.full_control"

	// View your data in Google Cloud Storage
	DevstorageReadOnlyScope = "https://www.googleapis.com/auth/devstorage.read_only"

	// Manage your data in Google Cloud Storage
	DevstorageReadWriteScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// Manage your data in the Google Prediction API
	PredictionScope = "https://www.googleapis.com/auth/prediction"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Hostedmodels = NewHostedmodelsService(s)
	s.Training = NewTrainingService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Hostedmodels *HostedmodelsService

	Training *TrainingService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewHostedmodelsService(s *Service) *HostedmodelsService {
	rs := &HostedmodelsService{s: s}
	return rs
}

type HostedmodelsService struct {
	s *Service
}

func NewTrainingService(s *Service) *TrainingService {
	rs := &TrainingService{s: s}
	return rs
}

type TrainingService struct {
	s *Service
}

type Input struct {
	// Input: Input to the model for a prediction
	Input *InputInput `json:"input,omitempty"`
}

// InputInput: Input to the model for a prediction
type InputInput struct {
	// CsvInstance: A list of input features, these can be strings or
	// doubles.
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

type Output struct {
	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// OutputLabel: The most likely class [Categorical models only].
	OutputLabel string `json:"outputLabel,omitempty"`

	// OutputMulti: A list of classes with their estimated probabilities
	// [Categorical models only].
	OutputMulti []*OutputOutputMulti `json:"outputMulti,omitempty"`

	// OutputValue: The estimated regression value [Regression models only].
	OutputValue float64 `json:"outputValue,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type OutputOutputMulti struct {
	// Label: The class label.
	Label string `json:"label,omitempty"`

	// Score: The probability of the class.
	Score float64 `json:"score,omitempty"`
}

type Training struct {
	// Id: The unique name for the predictive model.
	Id string `json:"id,omitempty"`

	// Kind: What kind of resource this is.
	Kind string `json:"kind,omitempty"`

	// ModelInfo: Model metadata.
	ModelInfo *TrainingModelInfo `json:"modelInfo,omitempty"`

	// SelfLink: A URL to re-request this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// TrainingStatus: The current status of the training job. This can be
	// one of following: RUNNING; DONE; ERROR; ERROR: TRAINING JOB NOT FOUND
	TrainingStatus string `json:"trainingStatus,omitempty"`

	// Utility: A class weighting function, which allows the importance
	// weights for classes to be specified [Categorical models only].
	Utility []*TrainingUtility `json:"utility,omitempty"`
}

// TrainingModelInfo: Model metadata.
type TrainingModelInfo struct {
	// ClassWeightedAccuracy: Estimated accuracy of model taking utility
	// weights into account [Categorical models only].
	ClassWeightedAccuracy float64 `json:"classWeightedAccuracy,omitempty"`

	// ClassificationAccuracy: A number between 0.0 and 1.0, where 1.0 is
	// 100% accurate. This is an estimate, based on the amount and quality
	// of the training data, of the estimated prediction accuracy. You can
	// use this is a guide to decide whether the results are accurate enough
	// for your needs. This estimate will be more reliable if your real
	// input data is similar to your training data [Categorical models
	// only].
	ClassificationAccuracy float64 `json:"classificationAccuracy,omitempty"`

	// ConfusionMatrix: An output confusion matrix. This shows an estimate
	// for how this model will do in predictions. This is first indexed by
	// the true class label. For each true class label, this provides a pair
	// {predicted_label, count}, where count is the estimated number of
	// times the model will predict the predicted label given the true
	// label. Will not output if more then 100 classes [Categorical models
	// only].
	ConfusionMatrix *TrainingModelInfoConfusionMatrix `json:"confusionMatrix,omitempty"`

	// ConfusionMatrixRowTotals: A list of the confusion matrix row totals
	ConfusionMatrixRowTotals *TrainingModelInfoConfusionMatrixRowTotals `json:"confusionMatrixRowTotals,omitempty"`

	// MeanSquaredError: An estimated mean squared error. The can be used to
	// measure the quality of the predicted model [Regression models only].
	MeanSquaredError float64 `json:"meanSquaredError,omitempty"`

	// ModelType: Type of predictive model (CLASSIFICATION or REGRESSION)
	ModelType string `json:"modelType,omitempty"`

	// NumberClasses: Number of classes in the trained model [Categorical
	// models only].
	NumberClasses int64 `json:"numberClasses,omitempty,string"`

	// NumberInstances: Number of valid data instances used in the trained
	// model.
	NumberInstances int64 `json:"numberInstances,omitempty,string"`
}

// TrainingModelInfoConfusionMatrix: An output confusion matrix. This
// shows an estimate for how this model will do in predictions. This is
// first indexed by the true class label. For each true class label,
// this provides a pair {predicted_label, count}, where count is the
// estimated number of times the model will predict the predicted label
// given the true label. Will not output if more then 100 classes
// [Categorical models only].
type TrainingModelInfoConfusionMatrix struct {
}

// TrainingModelInfoConfusionMatrixRowTotals: A list of the confusion
// matrix row totals
type TrainingModelInfoConfusionMatrixRowTotals struct {
}

// TrainingUtility: Class label (string).
type TrainingUtility struct {
}

type Update struct {
	// ClassLabel: The true class label of this instance
	ClassLabel string `json:"classLabel,omitempty"`

	// CsvInstance: The input features for this instance
	CsvInstance []interface{} `json:"csvInstance,omitempty"`
}

// method id "prediction.hostedmodels.predict":

type HostedmodelsPredictCall struct {
	s               *Service
	hostedModelName string
	input           *Input
	opt_            map[string]interface{}
}

// Predict: Submit input and request an output against a hosted model
func (r *HostedmodelsService) Predict(hostedModelName string, input *Input) *HostedmodelsPredictCall {
	c := &HostedmodelsPredictCall{s: r.s, opt_: make(map[string]interface{})}
	c.hostedModelName = hostedModelName
	c.input = input
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *HostedmodelsPredictCall) Fields(s ...googleapi.Field) *HostedmodelsPredictCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *HostedmodelsPredictCall) IfNoneMatch(entityTag string) *HostedmodelsPredictCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.hostedmodels.predict" call.
// Exactly one of the return values is non-nil.
func (c *HostedmodelsPredictCall) Do() (*Output, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "prediction.hostedmodels.predict" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *HostedmodelsPredictCall) DoHeader() (ret *Output, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.input)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "hostedmodels/{hostedModelName}/predict")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"hostedModelName": c.hostedModelName,
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
	//   "description": "Submit input and request an output against a hosted model",
	//   "httpMethod": "POST",
	//   "id": "prediction.hostedmodels.predict",
	//   "parameterOrder": [
	//     "hostedModelName"
	//   ],
	//   "parameters": {
	//     "hostedModelName": {
	//       "description": "The name of a hosted model",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "hostedmodels/{hostedModelName}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.delete":

type TrainingDeleteCall struct {
	s    *Service
	data string
	opt_ map[string]interface{}
}

// Delete: Delete a trained model
func (r *TrainingService) Delete(data string) *TrainingDeleteCall {
	c := &TrainingDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.data = data
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingDeleteCall) Fields(s ...googleapi.Field) *TrainingDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *TrainingDeleteCall) IfNoneMatch(entityTag string) *TrainingDeleteCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.training.delete" call.
func (c *TrainingDeleteCall) Do() error {
	_, err := c.DoHeader()
	return err
}

// DoHeader executes the "prediction.training.delete" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *TrainingDeleteCall) DoHeader() (resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "training/{data}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"data": c.data,
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
	//   "description": "Delete a trained model",
	//   "httpMethod": "DELETE",
	//   "id": "prediction.training.delete",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.get":

type TrainingGetCall struct {
	s    *Service
	data string
	opt_ map[string]interface{}
}

// Get: Check training status of your model
func (r *TrainingService) Get(data string) *TrainingGetCall {
	c := &TrainingGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.data = data
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingGetCall) Fields(s ...googleapi.Field) *TrainingGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *TrainingGetCall) IfNoneMatch(entityTag string) *TrainingGetCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.training.get" call.
// Exactly one of the return values is non-nil.
func (c *TrainingGetCall) Do() (*Training, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "prediction.training.get" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *TrainingGetCall) DoHeader() (ret *Training, resHeader http.Header, err error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "training/{data}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"data": c.data,
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
	//   "description": "Check training status of your model",
	//   "httpMethod": "GET",
	//   "id": "prediction.training.get",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.insert":

type TrainingInsertCall struct {
	s        *Service
	training *Training
	opt_     map[string]interface{}
}

// Insert: Begin training your model
func (r *TrainingService) Insert(training *Training) *TrainingInsertCall {
	c := &TrainingInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.training = training
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingInsertCall) Fields(s ...googleapi.Field) *TrainingInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *TrainingInsertCall) IfNoneMatch(entityTag string) *TrainingInsertCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.training.insert" call.
// Exactly one of the return values is non-nil.
func (c *TrainingInsertCall) Do() (*Training, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "prediction.training.insert" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *TrainingInsertCall) DoHeader() (ret *Training, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.training)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "training")
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
		return nil, res.Header, err
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, res.Header, err
	}
	return ret, res.Header, nil
	// {
	//   "description": "Begin training your model",
	//   "httpMethod": "POST",
	//   "id": "prediction.training.insert",
	//   "path": "training",
	//   "request": {
	//     "$ref": "Training"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/devstorage.full_control",
	//     "https://www.googleapis.com/auth/devstorage.read_only",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.predict":

type TrainingPredictCall struct {
	s     *Service
	data  string
	input *Input
	opt_  map[string]interface{}
}

// Predict: Submit data and request a prediction
func (r *TrainingService) Predict(data string, input *Input) *TrainingPredictCall {
	c := &TrainingPredictCall{s: r.s, opt_: make(map[string]interface{})}
	c.data = data
	c.input = input
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingPredictCall) Fields(s ...googleapi.Field) *TrainingPredictCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *TrainingPredictCall) IfNoneMatch(entityTag string) *TrainingPredictCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.training.predict" call.
// Exactly one of the return values is non-nil.
func (c *TrainingPredictCall) Do() (*Output, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "prediction.training.predict" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *TrainingPredictCall) DoHeader() (ret *Output, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.input)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "training/{data}/predict")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"data": c.data,
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
	//   "description": "Submit data and request a prediction",
	//   "httpMethod": "POST",
	//   "id": "prediction.training.predict",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}/predict",
	//   "request": {
	//     "$ref": "Input"
	//   },
	//   "response": {
	//     "$ref": "Output"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}

// method id "prediction.training.update":

type TrainingUpdateCall struct {
	s      *Service
	data   string
	update *Update
	opt_   map[string]interface{}
}

// Update: Add new data to a trained model
func (r *TrainingService) Update(data string, update *Update) *TrainingUpdateCall {
	c := &TrainingUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.data = data
	c.update = update
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TrainingUpdateCall) Fields(s ...googleapi.Field) *TrainingUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation fail if
// the object's Etag matches the given value. This is useful for getting updates
// only after the object has changed since the last request.
func (c *TrainingUpdateCall) IfNoneMatch(entityTag string) *TrainingUpdateCall {
	c.opt_["ifNoneMatch"] = entityTag
	return c
}

// Do executes the "prediction.training.update" call.
// Exactly one of the return values is non-nil.
func (c *TrainingUpdateCall) Do() (*Training, error) {
	v, _, err := c.DoHeader()
	return v, err
}

// DoHeader executes the "prediction.training.update" call.
// resHeader is populated with the response header when a response is received,
// regardless of the status code returned. This can be useful for checking for
// header values such as "Etag" even when http.StatusNotModified is returned.
func (c *TrainingUpdateCall) DoHeader() (ret *Training, resHeader http.Header, err error) {
	var body io.Reader = nil
	body, err = googleapi.WithoutDataWrapper.JSONReader(c.update)
	if err != nil {
		return nil, nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "training/{data}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"data": c.data,
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
	//   "description": "Add new data to a trained model",
	//   "httpMethod": "PUT",
	//   "id": "prediction.training.update",
	//   "parameterOrder": [
	//     "data"
	//   ],
	//   "parameters": {
	//     "data": {
	//       "description": "mybucket/mydata resource in Google Storage",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "training/{data}",
	//   "request": {
	//     "$ref": "Update"
	//   },
	//   "response": {
	//     "$ref": "Training"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/prediction"
	//   ]
	// }

}
