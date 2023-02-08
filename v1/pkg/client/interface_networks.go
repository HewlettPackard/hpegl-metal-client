// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

// Code generated by ifacemaker; DO NOT EDIT.

package client

import (
	_context "context"
	_nethttp "net/http"
)

// NetworksAPI defines the client functions provided for Networks.
type NetworksAPI interface {
	/*
	   Add Add a new network
	   Adds a new network that can be referenced when creating a Host
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param newNetwork Network that is to be added to the project

	   @return Network
	*/
	Add(ctx _context.Context, newNetwork NewNetwork) (Network, *_nethttp.Response, error)
	/*
	   Delete Delete a network
	   Deletes the network with the matching ID
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param networkId ID of network to delete
	*/
	Delete(ctx _context.Context, networkId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve network by ID
	   Returns a single network with matching ID
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param networkId ID of network to return

	   @return Network
	*/
	GetByID(ctx _context.Context, networkId string) (Network, *_nethttp.Response, error)
	/*
	   List List all networks in project
	   Returns an array of all network objects defined within the project.
	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

	   @return []Network
	*/
	List(ctx _context.Context) ([]Network, *_nethttp.Response, error)
	/*
	   Update Update an existing network.

	     - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	     - @param networkId ID of network to update
	     - @param network Updated network

	   @return Network
	*/
	Update(ctx _context.Context, networkId string, network Network) (Network, *_nethttp.Response, error)
}
