// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package client

import (
	_context "context"
	_nethttp "net/http"
)

// VolumeAttachmentsAPI defines the functions provided by the client for
// VolumeAttachments.
type VolumeAttachmentsAPI interface {
	/*
	   GetByID Retrieve volume attachment by ID
	   Returns a single volume attachment with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param attachmentId ID of volume attachment to return
	   @return VolumeAttachment
	*/
	GetByID(ctx _context.Context, attachmentId string) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   List List all volume attachments in project
	   Returns an array of all VolumeAttachments defined within the project.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return []VolumeAttachment
	*/
	List(ctx _context.Context) ([]VolumeAttachment, *_nethttp.Response, error)
}
