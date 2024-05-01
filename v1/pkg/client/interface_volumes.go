// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
)

// VolumesAPI defines the client functions provided for Volumes.
type VolumesAPI interface {
	/*
	   Add Add a new volume
	   Adds a new volume to the project. Volumes may be created separately and then referenced in the create Host call; or volumes may be created directly within the create Host call. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param newVolume Volume that is to be added to the project
	     - @param optional nil or *VolumesApiAddOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return Volume
	*/
	Add(ctx _context.Context, newVolume NewVolume, localVarOptionals *VolumesApiAddOpts) (Volume, *_nethttp.Response, error)
	/*
	   Attach Attach existing volume to Host
	   Attaches the indicated volume to a host identified in the requestBody.   This attachment will create a VolumeAttachment object that contains  details about the connection of the volume and will update the Host  with iSCSI configuration information. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to attach
	     - @param volumeAttachHostUuid Unique ID of the Host to which the volume will be attached
	     - @param optional nil or *VolumesApiAttachOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return VolumeAttachment
	*/
	Attach(ctx _context.Context, volumeId string, volumeAttachHostUuid VolumeAttachHostUuid, localVarOptionals *VolumesApiAttachOpts) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   Delete Delete a volume
	   Deletes the volume with the matching ID. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to delete
	     - @param optional nil or *VolumesApiDeleteOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID
	*/
	Delete(ctx _context.Context, volumeId string, localVarOptionals *VolumesApiDeleteOpts) (*_nethttp.Response, error)
	/*
	   Detach Detach existing volume from Host
	   Detaches the indicated volume from the host identified in the requestBody.   This detachment will delete the VolumeAttachment object that contains  details about the connection of the volume and will update the Host  to remove selected iSCSI configuration information. Note that the HostID is required in the body of the request to ensure that the operation is well understood and that a volume is not accidently being removed from the wrong host. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to attach
	     - @param volumeAttachHostUuid Unique ID of the Host from which a volume will be detached
	     - @param optional nil or *VolumesApiDetachOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID
	*/
	Detach(ctx _context.Context, volumeId string, volumeAttachHostUuid VolumeAttachHostUuid, localVarOptionals *VolumesApiDetachOpts) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve volume by ID
	   Returns a single volume with matching imaged. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to return
	     - @param optional nil or *VolumesApiGetByIDOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return Volume
	*/
	GetByID(ctx _context.Context, volumeId string, localVarOptionals *VolumesApiGetByIDOpts) (Volume, *_nethttp.Response, error)
	/*
	   List List all volumes in project
	   Returns an array of all volumes defined within the project. If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param optional nil or *VolumesApiListOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return []Volume
	*/
	List(ctx _context.Context, localVarOptionals *VolumesApiListOpts) ([]Volume, *_nethttp.Response, error)
	/*
	   Update Update an existing volume
	   Updates volume with matching ID. Update is permitted only when volume is in &#39;Allocated&#39; or &#39;Visible&#39; state. Only the Volume &#39;Capacity&#39; can be updated with a value greater than the existing one to expand the volume.  If GreenLake Platform IAM issued token is used for authentication, then it is required to pass  &#39;X-Role&#39; and &#39;X-Workspaceid&#39; headers.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to return
	     - @param updateVolume Volume object with its ID and Capacity in GiB indicating the expanded size to be specified.
	     - @param optional nil or *VolumesApiUpdateOpts - Optional Parameters:
	     - @param "XRole" (optional.String) -  GreenLake Platform role
	     - @param "XWorkspaceid" (optional.String) -  GreenLake Platform workspace ID

	   @return Volume
	*/
	Update(ctx _context.Context, volumeId string, updateVolume UpdateVolume, localVarOptionals *VolumesApiUpdateOpts) (Volume, *_nethttp.Response, error)
}
