# hpegl-metal-client
HPE GL Bare-Metal-as-a-Service Client APIs (project-level)

This repository provides the API documentation and the GO SDK for the HPE GreenLake Metal Project Client APIs. The APIs are documented in a standard OpenAPI 3 specification format, and the GO SDK has generated automatically from the OpenAPI generator tools.  

Navigate to version-specific API Client directories (\<version\>/pkg/client) to view the documentation and get an orientation to the APIs.

## Usage
The project API is intended for use by services and users accessing and provisioning project resources assigned specifically to a project. Projects are created and managed by a project administrator. Project resource limits for number of hosts and volumes and total storage capacity and project members are managed by the project administrator for the project. The scope of resource access and management by project members is limited to hosts and storage volumes created within the project limits by team members with project credentials. The project API provides access to available inventory which can be consumed for creating hosts, sshkeys, custom networks, and persistent storage volumes. Inventory usage is constrained by the project limits definitions. 

The APIs use bearer authentication (also called token authentication) that requires the user to provide a bearer token in the authorization header. Two token methods are supported, the HPE GreenLake IAM token and Metal classic tokens.

## Code generation
Client code is generated from the OpenAPI files in \<version\>/api/swagger/ using openapi-generator. In order to generate go client source code from the OpenAPI files, clone the repository https://github.com/hewlettpackard/hpegl-metal-client.git and install openapi-generator. Development and build environment requires Ubuntu 18.04. It is recommended to use this for working with hpegl-metal-client source generation also. Instructions on installing the openapi-generator on Ubuntu are below.

_NOTE: At this time openapi-generator version 4.2.3 only is supported. The latest version is 4.3 and has issues with this api spec._

_NOTE: java is required for openapi-generator-cli to work.  Install default JRE if you don't already have java installed on your development environment._

_NOTE: All files under **\<version\>/pkg/client** are generated code.  Do not modify these files directly._


### Installing openapi-generator and ifacemaker

```bash
$ sudo apt update
$ sudo apt install default-jre

$ sudo npm install @openapitools/openapi-generator-cli@cli-4.2.3 -g
$ go get github.com/vburenin/ifacemaker@v1.1.0

```
### Using openapi-generator for Go
With the generator installed and repository cloned, navigate to the hgegl-metal-client/\<version\> directory. Generate the Go client using

```bash
$ make gen

```

The generated client code will be created in **/\<version\>/pkg/client** with package name "client". 

Interfaces are created using **ifacemaker** (`go get github.com/vburenin/ifacemaker`). This is primarily intended to allow clients to create mocks for the APIs. 

Below commands create interfaces for each resource object.

```bash
ifacemaker --struct=AvailableResourcesApiService --file=api_available_resources.go --output=interface_available_resources.go -p client --iface=AvailableResourcesAPI
ifacemaker --struct=HostsApiService --file=api_hosts.go --output=interface_hosts.go -p client --iface=HostsAPI
ifacemaker --struct=NetworksApiService --file=api_networks.go --output=interface_networks.go -p client --iface=NetworksAPI
ifacemaker --struct=ProjectsApiService --file=api_projects.go --output=interface_projects.go -p client --iface=ProjectsAPI
ifacemaker --struct=ProjectsInfoApiService --file=api_projects_info.go --output=interface_projects_info.go -p client --iface=ProjectsInfoAPI
ifacemaker --struct=SshkeysApiService --file=api_sshkeys.go --output=interface_sshkeys.go -p client --iface=SshkeysAPI
ifacemaker --struct=UsageReportsApiService --file=api_usage_reports.go --output=interface_usage_reports.go -p client --iface=UsageReportsAPI
ifacemaker --struct=VersionApiService --file=api_version.go --output=interface_version.go -p client --iface=VersionAPI
ifacemaker --struct=VolumeAttachmentsApiService --file=api_volume_attachments.go --output=interface_volume_attachments.go -p client --iface=VolumeAttachmentsAPI
ifacemaker --struct=VolumesApiService --file=api_volumes.go --output=interface_volumes.go -p client --iface=VolumesAPI
ifacemaker --struct=IppoolsApiService --file=api_ippools.go --output=interface_ippools.go -p client --iface=IPPoolsAPI
```
