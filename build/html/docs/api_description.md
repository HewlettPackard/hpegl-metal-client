 
# Introduction
Welcome to HPE GreenLake for bare metal APIs. 

The HPE GreenLake for bare metal is included as part of the HPE GreenLake for Private Cloud
Enterprise service, and it is sometimes also called Bare Metal as a Service or BMaaS.

This document describes the HPE GreenLake for bare metal APIs protocol and available endpoints.

The BMaaS API is built on HTTP and is RESTful. It has predictable resource URLs. It returns HTTP 
response codes to indicate errors. It also accepts and returns JSON in the HTTP body. You can 
use your favorite HTTP/REST library for your programming language to use HPE GreenLake for Private 
Cloud Enterprise bare metal APIs.

The APIs provides access to bare metal service within a single project context. Clients can create 
fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated 
project environment. Clients can also access information about available services and resources 
through the available-resources API. This API provides detailed information about OS imaging options, 
machine size options, storage volume options, and data center locations which are needed when creating 
hosts and volumes.

The API specification references a few resources by the deprecated name, which differs from the name 
displayed in HPE GreenLake Central UI. The table below maps the resource's old name in the API 
specification with the corresponding name in HPE GreenLake Central UI.

| **SN** | **Name in API Spec** | **Name in GreenLake Central** |
|--------|----------------------|-------------------------------|
| 1      | Project              | Compute Group                 |
| 2      | Host                 | Compute Instance              |
| 3      | Machine Size         | Compute Instance Type         |


# Authentication

## Bearer Auth

The HPE GreenLake for bare metal APIs use Bearer Authentication that requires the users to provide 
a bearer token in the **Authorization** header. The service supports two types of tokens, the HPE 
GreenLake IAM token and the Metal classic token. The Metal classic token is for HPE internal use only.
The clients are required to provide HPE GreenLake IAM token to authenticate with the HPE GreenLake 
for bare metal APIs. The access token can be obtained using any of the following ways.

<h2>Option 1: API client(recommended)</h2>

An API client allows nonhuman entities (an application service account, for instance) programmatic access to a resource on a space.

  + 1.
<a
    href="https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=greenlakecentral-create-api-client.html"
    target="_blank">
    Create an API client
    <img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg"
    height="12">
</a>

**Note:** Make sure to save the **Issuer Url**, **Client ID**, and **Client Secret**

  + 2.
<a
    href="https://support.hpe.com/hpesc/public/docDisplay?docId=a00092451en_us&page=GUID-1CEA233B-C4B0-41B7-9A25-7A36D9FC0312.html"
    target="_blank">
    Creating an assignment for an API client
    <img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg"
    height="12">
</a>

**Note:** Make sure to assign the appropriate HPE GreenLake for bare metal user roles in HPE GreenLake Central.

  + 3. Make a REST call to generate the API access token:

```
curl -i -X POST \
  '{issuerUrl}/v1/token' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d 'client_id={clientID}' \
  -d 'client_secret={clientSecret}' \
  -d grant_type=client_credentials \
  -d scope=hpe-tenant
```
Obtain the `access_token` from the response.

<h2>Option 2: Getting the token directly from UI</h2>

To authenticate with the BMaaS API you need to obtain the access token from
<a
    href="https://client.greenlake.hpe.com"
    target="_blank">
    HPE GreenLake Central
    <img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg"
    height="12">
</a>.
Once logged in to **HPE GreenLake Central**,
Click the
<img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg">
icon in the top right corner and then click
<input
    type="button"
    value="API Access"
    style="color: #000;
        background-color: #fff;
        border: 2px solid #00b388;
        border-radius: 4px;"/>
. In the API Access page, click the
<img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/clipboard.svg">
icon to copy the personal access token.

<br>

<img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/circle-information.svg"> **INFO**
For both the options, the access token lease expires in **15** minutes.

## Project API Key
The specific Project ID must be provided with **Project** header when accessing service within a single project context.
Used only when using HPE GreenLake IAM issued access tokens.

## Membership API Key
HPE Internal use only. Required only when using classic Metal access tokens.

# Getting a Space name or ID
Space name or ID is required in header for a few Project level APIs. You can obtain the Space name and the corresponding ID from 
<a
    href="https://client.greenlake.hpe.com"
    target="_blank">
    HPE GreenLake Central
    <img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/share.svg"
    height="12">
</a>.   
Once logged in to **HPE GreenLake Central**,
Click the
<img src="https://raw.githubusercontent.com/grommet/grommet-icons/stable/img/user.svg">
icon in the top right corner and copy the value of **Space** field.  

The corresponding ID for the Space can be extracted from the current URL in the address bar.  
It is the value of **spaceid** query parameter at the end of a URL after a '?' symbol.

# HPE GreenLake for bare metal APIs
