// Code generated by go-swagger; DO NOT EDIT.

package remediation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetRemediationsURL generates an URL for the get remediations operation
type GetRemediationsURL struct {
	ProjectName string
	ServiceName string
	StageName   string

	NextPageKey *string
	PageSize    *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRemediationsURL) WithBasePath(bp string) *GetRemediationsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetRemediationsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetRemediationsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/project/{projectName}/stage/{stageName}/service/{serviceName}/remediation"

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{projectName}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on GetRemediationsURL")
	}

	serviceName := o.ServiceName
	if serviceName != "" {
		_path = strings.Replace(_path, "{serviceName}", serviceName, -1)
	} else {
		return nil, errors.New("serviceName is required on GetRemediationsURL")
	}

	stageName := o.StageName
	if stageName != "" {
		_path = strings.Replace(_path, "{stageName}", stageName, -1)
	} else {
		return nil, errors.New("stageName is required on GetRemediationsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var nextPageKeyQ string
	if o.NextPageKey != nil {
		nextPageKeyQ = *o.NextPageKey
	}
	if nextPageKeyQ != "" {
		qs.Set("nextPageKey", nextPageKeyQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt64(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetRemediationsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetRemediationsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetRemediationsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetRemediationsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetRemediationsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetRemediationsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}