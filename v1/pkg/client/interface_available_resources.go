// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// AvailableResourcesAPI defines the client functions provided for AvailableResources.
type AvailableResourcesAPI interface {
	/*
	   List Get lists of available resources for creating hosts and volumes
	   Used to get lists of options that are used when creating hosts and volumes. A get /available-resources will return an object that includes the following arrays: * Images - A list of image service IDs along with their category (Linux),    flavor (ubuntu), and version (18.04)  * MachineSizes - A list of machine size IDs along with the machine size    names and detailed descriptions  * Locations - A list of location IDs along with their country, region,    and data center.  * Networks - A list of available Network IDs along with the network name,   location ID, network kind, and host usage (Required, Default, Optional)  * MachineInventory - Information about the available inventory of machines    based on location ID and machine size ID.  While this information may    change rapidly, it can be used by GUIs and systems to restrict host   creates to locations with the desired machine size.  * StorageInventory - Information about the current available storage capacity    for a specific volume flavor by site.   * VolumeFlavors - A list of volume flavor IDs along with their name and    detailed description.  * Volumes - A list of current, existing volumes.  If the volume is in the   the right state, it could be attached to a new Host.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return AvailableResources
	*/
	List(ctx _context.Context) (AvailableResources, *_nethttp.Response, error)
}
