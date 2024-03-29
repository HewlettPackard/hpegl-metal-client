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
	   Adds a new volume to the project.  Volumes may be created separately and then referenced in the create Host call; or volumes may be created directly within the create Host call.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param newVolume Volume that is to be added to the project

	   @return Volume
	*/
	Add(ctx _context.Context, newVolume NewVolume) (Volume, *_nethttp.Response, error)
	/*
	   Attach Attach existing volume to Host
	   Attaches the indicated volume to a host identified in the requestBody.   This attachment will create a VolumeAttachment object that contains  details about the connection of the volume and will update the Host  with iSCSI configuration information.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to attach
	     - @param volumeAttachHostUuid Unique ID of the Host to which the volume will be attached

	   @return VolumeAttachment
	*/
	Attach(ctx _context.Context, volumeId string, volumeAttachHostUuid VolumeAttachHostUuid) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   Delete Delete a volume
	   Deletes the volume with the matching ID
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to delete
	*/
	Delete(ctx _context.Context, volumeId string) (*_nethttp.Response, error)
	/*
	   Detach Detach existing volume from Host
	   Detaches the indicated volume from the host identified in the requestBody.   This detachment will delete the VolumeAttachment object that contains  details about the connection of the volume and will update the Host  to remove selected iSCSI configuration information. Note that the HostID is required in the body of the request to ensure that the operation is well understood and that a volume is not accidently being removed from the wrong host.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to attach
	     - @param volumeAttachHostUuid Unique ID of the Host from which a volume will be detached
	*/
	Detach(ctx _context.Context, volumeId string, volumeAttachHostUuid VolumeAttachHostUuid) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve volume by ID
	   Returns a single volume with matching ID
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to return

	   @return Volume
	*/
	GetByID(ctx _context.Context, volumeId string) (Volume, *_nethttp.Response, error)
	/*
	   List List all volumes in project
	   Returns an array of all volumes defined within the project.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

	   @return []Volume
	*/
	List(ctx _context.Context) ([]Volume, *_nethttp.Response, error)
	/*
	   Update Update an existing volume
	   Updates volume with matching ID. Update is permitted only when volume is in &#39;Allocated&#39; or &#39;Visible&#39; state. Only the Volume &#39;Capacity&#39; can be updated with a value greater than the existing one to expand the volume.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param volumeId ID of volume to return
	     - @param updateVolume Volume object with its ID and Capacity in GiB indicating the expanded size to be specified.

	   @return Volume
	*/
	Update(ctx _context.Context, volumeId string, updateVolume UpdateVolume) (Volume, *_nethttp.Response, error)
}
