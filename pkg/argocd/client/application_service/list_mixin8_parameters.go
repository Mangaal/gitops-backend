// Code generated by go-swagger; DO NOT EDIT.

package application_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListMixin8Params creates a new ListMixin8Params object
// with the default values initialized.
func NewListMixin8Params() *ListMixin8Params {
	var ()
	return &ListMixin8Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListMixin8ParamsWithTimeout creates a new ListMixin8Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListMixin8ParamsWithTimeout(timeout time.Duration) *ListMixin8Params {
	var ()
	return &ListMixin8Params{

		timeout: timeout,
	}
}

// NewListMixin8ParamsWithContext creates a new ListMixin8Params object
// with the default values initialized, and the ability to set a context for a request
func NewListMixin8ParamsWithContext(ctx context.Context) *ListMixin8Params {
	var ()
	return &ListMixin8Params{

		Context: ctx,
	}
}

// NewListMixin8ParamsWithHTTPClient creates a new ListMixin8Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListMixin8ParamsWithHTTPClient(client *http.Client) *ListMixin8Params {
	var ()
	return &ListMixin8Params{
		HTTPClient: client,
	}
}

/*ListMixin8Params contains all the parameters to send to the API endpoint
for the list mixin8 operation typically these are written to a http.Request
*/
type ListMixin8Params struct {

	/*Name
	  the application's name.

	*/
	Name *string
	/*Project
	  the project names to restrict returned list applications.

	*/
	Project []string
	/*Refresh
	  forces application reconciliation if set to true.

	*/
	Refresh *string
	/*ResourceVersion
	  when specified with a watch call, shows changes that occur after that particular version of a resource.

	*/
	ResourceVersion *string
	/*Selector
	  the selector to to restrict returned list to applications only with matched labels.

	*/
	Selector *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list mixin8 params
func (o *ListMixin8Params) WithTimeout(timeout time.Duration) *ListMixin8Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list mixin8 params
func (o *ListMixin8Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list mixin8 params
func (o *ListMixin8Params) WithContext(ctx context.Context) *ListMixin8Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list mixin8 params
func (o *ListMixin8Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list mixin8 params
func (o *ListMixin8Params) WithHTTPClient(client *http.Client) *ListMixin8Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list mixin8 params
func (o *ListMixin8Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the list mixin8 params
func (o *ListMixin8Params) WithName(name *string) *ListMixin8Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the list mixin8 params
func (o *ListMixin8Params) SetName(name *string) {
	o.Name = name
}

// WithProject adds the project to the list mixin8 params
func (o *ListMixin8Params) WithProject(project []string) *ListMixin8Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the list mixin8 params
func (o *ListMixin8Params) SetProject(project []string) {
	o.Project = project
}

// WithRefresh adds the refresh to the list mixin8 params
func (o *ListMixin8Params) WithRefresh(refresh *string) *ListMixin8Params {
	o.SetRefresh(refresh)
	return o
}

// SetRefresh adds the refresh to the list mixin8 params
func (o *ListMixin8Params) SetRefresh(refresh *string) {
	o.Refresh = refresh
}

// WithResourceVersion adds the resourceVersion to the list mixin8 params
func (o *ListMixin8Params) WithResourceVersion(resourceVersion *string) *ListMixin8Params {
	o.SetResourceVersion(resourceVersion)
	return o
}

// SetResourceVersion adds the resourceVersion to the list mixin8 params
func (o *ListMixin8Params) SetResourceVersion(resourceVersion *string) {
	o.ResourceVersion = resourceVersion
}

// WithSelector adds the selector to the list mixin8 params
func (o *ListMixin8Params) WithSelector(selector *string) *ListMixin8Params {
	o.SetSelector(selector)
	return o
}

// SetSelector adds the selector to the list mixin8 params
func (o *ListMixin8Params) SetSelector(selector *string) {
	o.Selector = selector
}

// WriteToRequest writes these params to a swagger request
func (o *ListMixin8Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	valuesProject := o.Project

	joinedProject := swag.JoinByFormat(valuesProject, "")
	// query array param project
	if err := r.SetQueryParam("project", joinedProject...); err != nil {
		return err
	}

	if o.Refresh != nil {

		// query param refresh
		var qrRefresh string
		if o.Refresh != nil {
			qrRefresh = *o.Refresh
		}
		qRefresh := qrRefresh
		if qRefresh != "" {
			if err := r.SetQueryParam("refresh", qRefresh); err != nil {
				return err
			}
		}

	}

	if o.ResourceVersion != nil {

		// query param resourceVersion
		var qrResourceVersion string
		if o.ResourceVersion != nil {
			qrResourceVersion = *o.ResourceVersion
		}
		qResourceVersion := qrResourceVersion
		if qResourceVersion != "" {
			if err := r.SetQueryParam("resourceVersion", qResourceVersion); err != nil {
				return err
			}
		}

	}

	if o.Selector != nil {

		// query param selector
		var qrSelector string
		if o.Selector != nil {
			qrSelector = *o.Selector
		}
		qSelector := qrSelector
		if qSelector != "" {
			if err := r.SetQueryParam("selector", qSelector); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}