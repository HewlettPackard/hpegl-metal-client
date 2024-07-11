// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
)

// VolumeAttachmentsAPI defines the client functions provided for VolumeAttachments.
type VolumeAttachmentsAPI interface {
	/*
	   Add Create a new VolumeAttachment
	   Adds a new VolumeAttachment which enables a machine to use a volume. Note that this API is for creation of a VolumeAttachment for non-Hosts. A VolumeAttachment resource must be deleted using the &#39;DELETE /volume-attachments/{attachmentId}&#39; API. For regular hosts, a VolumeAttachment is created via the &#39;POST /volumes/{volumeId}/attach&#39; API. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param newVolumeAttachment NewVolumeAttachement parameters to create a new VolumeAttachment.
	     - @param optional nil or *VolumeAttachmentsApiAddOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role name
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return VolumeAttachment
	*/
	Add(ctx _context.Context, newVolumeAttachment NewVolumeAttachment, localVarOptionals *VolumeAttachmentsApiAddOpts) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   Delete Delete a VolumeAttachment
	   Deletes the VolumeAttachment with the matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param attachmentId ID of VolumeAttachment to delete
	     - @param optional nil or *VolumeAttachmentsApiDeleteOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role name
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID
	*/
	Delete(ctx _context.Context, attachmentId string, localVarOptionals *VolumeAttachmentsApiDeleteOpts) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve volume attachment by ID
	   Returns a single volume attachment with matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param attachmentId ID of volume attachment to return
	     - @param optional nil or *VolumeAttachmentsApiGetByIDOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role name
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return VolumeAttachment
	*/
	GetByID(ctx _context.Context, attachmentId string, localVarOptionals *VolumeAttachmentsApiGetByIDOpts) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   List List all volume attachments in project
	   Returns an array of all VolumeAttachments defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *VolumeAttachmentsApiListOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role name
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return []VolumeAttachment
	*/
	List(ctx _context.Context, localVarOptionals *VolumeAttachmentsApiListOpts) ([]VolumeAttachment, *_nethttp.Response, error)
}
