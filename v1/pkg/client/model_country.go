// (C) Copyright 2021-2024 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// Country the model 'Country'
type Country string

// List of Country
const (
	COUNTRY_USA Country = "USA"
	COUNTRY_AFG Country = "AFG"
	COUNTRY_ALA Country = "ALA"
	COUNTRY_ALB Country = "ALB"
	COUNTRY_DZA Country = "DZA"
	COUNTRY_ASM Country = "ASM"
	COUNTRY_AND Country = "AND"
	COUNTRY_AGO Country = "AGO"
	COUNTRY_AIA Country = "AIA"
	COUNTRY_ATA Country = "ATA"
	COUNTRY_ATG Country = "ATG"
	COUNTRY_ARG Country = "ARG"
	COUNTRY_ARM Country = "ARM"
	COUNTRY_ABW Country = "ABW"
	COUNTRY_AUS Country = "AUS"
	COUNTRY_AUT Country = "AUT"
	COUNTRY_AZE Country = "AZE"
	COUNTRY_BHS Country = "BHS"
	COUNTRY_BHR Country = "BHR"
	COUNTRY_BGD Country = "BGD"
	COUNTRY_BRB Country = "BRB"
	COUNTRY_BLR Country = "BLR"
	COUNTRY_BEL Country = "BEL"
	COUNTRY_BLZ Country = "BLZ"
	COUNTRY_BEN Country = "BEN"
	COUNTRY_BMU Country = "BMU"
	COUNTRY_BTN Country = "BTN"
	COUNTRY_BOL Country = "BOL"
	COUNTRY_BES Country = "BES"
	COUNTRY_BIH Country = "BIH"
	COUNTRY_BWA Country = "BWA"
	COUNTRY_BVT Country = "BVT"
	COUNTRY_BRA Country = "BRA"
	COUNTRY_IOT Country = "IOT"
	COUNTRY_BRN Country = "BRN"
	COUNTRY_BGR Country = "BGR"
	COUNTRY_BFA Country = "BFA"
	COUNTRY_BDI Country = "BDI"
	COUNTRY_KHM Country = "KHM"
	COUNTRY_CMR Country = "CMR"
	COUNTRY_CAN Country = "CAN"
	COUNTRY_CPV Country = "CPV"
	COUNTRY_CYM Country = "CYM"
	COUNTRY_CAF Country = "CAF"
	COUNTRY_TCD Country = "TCD"
	COUNTRY_CHL Country = "CHL"
	COUNTRY_CHN Country = "CHN"
	COUNTRY_CXR Country = "CXR"
	COUNTRY_CCK Country = "CCK"
	COUNTRY_COL Country = "COL"
	COUNTRY_COM Country = "COM"
	COUNTRY_COG Country = "COG"
	COUNTRY_COD Country = "COD"
	COUNTRY_COK Country = "COK"
	COUNTRY_CRI Country = "CRI"
	COUNTRY_CIV Country = "CIV"
	COUNTRY_HRV Country = "HRV"
	COUNTRY_CUB Country = "CUB"
	COUNTRY_CUW Country = "CUW"
	COUNTRY_CYP Country = "CYP"
	COUNTRY_CZE Country = "CZE"
	COUNTRY_DNK Country = "DNK"
	COUNTRY_DJI Country = "DJI"
	COUNTRY_DMA Country = "DMA"
	COUNTRY_DOM Country = "DOM"
	COUNTRY_ECU Country = "ECU"
	COUNTRY_EGY Country = "EGY"
	COUNTRY_SLV Country = "SLV"
	COUNTRY_GNQ Country = "GNQ"
	COUNTRY_ERI Country = "ERI"
	COUNTRY_EST Country = "EST"
	COUNTRY_ETH Country = "ETH"
	COUNTRY_FLK Country = "FLK"
	COUNTRY_FRO Country = "FRO"
	COUNTRY_FJI Country = "FJI"
	COUNTRY_FIN Country = "FIN"
	COUNTRY_FRA Country = "FRA"
	COUNTRY_GUF Country = "GUF"
	COUNTRY_PYF Country = "PYF"
	COUNTRY_ATF Country = "ATF"
	COUNTRY_GAB Country = "GAB"
	COUNTRY_GMB Country = "GMB"
	COUNTRY_GEO Country = "GEO"
	COUNTRY_DEU Country = "DEU"
	COUNTRY_GHA Country = "GHA"
	COUNTRY_GIB Country = "GIB"
	COUNTRY_GRC Country = "GRC"
	COUNTRY_GRL Country = "GRL"
	COUNTRY_GRD Country = "GRD"
	COUNTRY_GLP Country = "GLP"
	COUNTRY_GUM Country = "GUM"
	COUNTRY_GTM Country = "GTM"
	COUNTRY_GGY Country = "GGY"
	COUNTRY_GIN Country = "GIN"
	COUNTRY_GNB Country = "GNB"
	COUNTRY_GUY Country = "GUY"
	COUNTRY_HTI Country = "HTI"
	COUNTRY_HMD Country = "HMD"
	COUNTRY_VAT Country = "VAT"
	COUNTRY_HND Country = "HND"
	COUNTRY_HKG Country = "HKG"
	COUNTRY_HUN Country = "HUN"
	COUNTRY_ISL Country = "ISL"
	COUNTRY_IND Country = "IND"
	COUNTRY_IDN Country = "IDN"
	COUNTRY_IRN Country = "IRN"
	COUNTRY_IRQ Country = "IRQ"
	COUNTRY_IRL Country = "IRL"
	COUNTRY_IMN Country = "IMN"
	COUNTRY_ISR Country = "ISR"
	COUNTRY_ITA Country = "ITA"
	COUNTRY_JAM Country = "JAM"
	COUNTRY_JPN Country = "JPN"
	COUNTRY_JEY Country = "JEY"
	COUNTRY_JOR Country = "JOR"
	COUNTRY_KAZ Country = "KAZ"
	COUNTRY_KEN Country = "KEN"
	COUNTRY_KIR Country = "KIR"
	COUNTRY_PRK Country = "PRK"
	COUNTRY_KOR Country = "KOR"
	COUNTRY_KWT Country = "KWT"
	COUNTRY_KGZ Country = "KGZ"
	COUNTRY_LAO Country = "LAO"
	COUNTRY_LVA Country = "LVA"
	COUNTRY_LBN Country = "LBN"
	COUNTRY_LSO Country = "LSO"
	COUNTRY_LBR Country = "LBR"
	COUNTRY_LBY Country = "LBY"
	COUNTRY_LIE Country = "LIE"
	COUNTRY_LTU Country = "LTU"
	COUNTRY_LUX Country = "LUX"
	COUNTRY_MAC Country = "MAC"
	COUNTRY_MKD Country = "MKD"
	COUNTRY_MDG Country = "MDG"
	COUNTRY_MWI Country = "MWI"
	COUNTRY_MYS Country = "MYS"
	COUNTRY_MDV Country = "MDV"
	COUNTRY_MLI Country = "MLI"
	COUNTRY_MLT Country = "MLT"
	COUNTRY_MHL Country = "MHL"
	COUNTRY_MTQ Country = "MTQ"
	COUNTRY_MRT Country = "MRT"
	COUNTRY_MUS Country = "MUS"
	COUNTRY_MYT Country = "MYT"
	COUNTRY_MEX Country = "MEX"
	COUNTRY_FSM Country = "FSM"
	COUNTRY_MDA Country = "MDA"
	COUNTRY_MCO Country = "MCO"
	COUNTRY_MNG Country = "MNG"
	COUNTRY_MNE Country = "MNE"
	COUNTRY_MSR Country = "MSR"
	COUNTRY_MAR Country = "MAR"
	COUNTRY_MOZ Country = "MOZ"
	COUNTRY_MMR Country = "MMR"
	COUNTRY_NAM Country = "NAM"
	COUNTRY_NRU Country = "NRU"
	COUNTRY_NPL Country = "NPL"
	COUNTRY_NLD Country = "NLD"
	COUNTRY_NCL Country = "NCL"
	COUNTRY_NZL Country = "NZL"
	COUNTRY_NIC Country = "NIC"
	COUNTRY_NER Country = "NER"
	COUNTRY_NGA Country = "NGA"
	COUNTRY_NIU Country = "NIU"
	COUNTRY_NFK Country = "NFK"
	COUNTRY_MNP Country = "MNP"
	COUNTRY_NOR Country = "NOR"
	COUNTRY_OMN Country = "OMN"
	COUNTRY_PAK Country = "PAK"
	COUNTRY_PLW Country = "PLW"
	COUNTRY_PSE Country = "PSE"
	COUNTRY_PAN Country = "PAN"
	COUNTRY_PNG Country = "PNG"
	COUNTRY_PRY Country = "PRY"
	COUNTRY_PER Country = "PER"
	COUNTRY_PHL Country = "PHL"
	COUNTRY_PCN Country = "PCN"
	COUNTRY_POL Country = "POL"
	COUNTRY_PRT Country = "PRT"
	COUNTRY_PRI Country = "PRI"
	COUNTRY_QAT Country = "QAT"
	COUNTRY_REU Country = "REU"
	COUNTRY_ROU Country = "ROU"
	COUNTRY_RUS Country = "RUS"
	COUNTRY_RWA Country = "RWA"
	COUNTRY_BLM Country = "BLM"
	COUNTRY_SHN Country = "SHN"
	COUNTRY_KNA Country = "KNA"
	COUNTRY_LCA Country = "LCA"
	COUNTRY_MAF Country = "MAF"
	COUNTRY_SPM Country = "SPM"
	COUNTRY_VCT Country = "VCT"
	COUNTRY_WSM Country = "WSM"
	COUNTRY_SMR Country = "SMR"
	COUNTRY_STP Country = "STP"
	COUNTRY_SAU Country = "SAU"
	COUNTRY_SEN Country = "SEN"
	COUNTRY_SRB Country = "SRB"
	COUNTRY_SYC Country = "SYC"
	COUNTRY_SLE Country = "SLE"
	COUNTRY_SGP Country = "SGP"
	COUNTRY_SXM Country = "SXM"
	COUNTRY_SVK Country = "SVK"
	COUNTRY_SVN Country = "SVN"
	COUNTRY_SLB Country = "SLB"
	COUNTRY_SOM Country = "SOM"
	COUNTRY_ZAF Country = "ZAF"
	COUNTRY_SGS Country = "SGS"
	COUNTRY_SSD Country = "SSD"
	COUNTRY_ESP Country = "ESP"
	COUNTRY_LKA Country = "LKA"
	COUNTRY_SDN Country = "SDN"
	COUNTRY_SUR Country = "SUR"
	COUNTRY_SJM Country = "SJM"
	COUNTRY_SWZ Country = "SWZ"
	COUNTRY_SWE Country = "SWE"
	COUNTRY_CHE Country = "CHE"
	COUNTRY_SYR Country = "SYR"
	COUNTRY_TWN Country = "TWN"
	COUNTRY_TJK Country = "TJK"
	COUNTRY_TZA Country = "TZA"
	COUNTRY_THA Country = "THA"
	COUNTRY_TLS Country = "TLS"
	COUNTRY_TGO Country = "TGO"
	COUNTRY_TKL Country = "TKL"
	COUNTRY_TON Country = "TON"
	COUNTRY_TTO Country = "TTO"
	COUNTRY_TUN Country = "TUN"
	COUNTRY_TUR Country = "TUR"
	COUNTRY_TKM Country = "TKM"
	COUNTRY_TCA Country = "TCA"
	COUNTRY_TUV Country = "TUV"
	COUNTRY_UGA Country = "UGA"
	COUNTRY_UKR Country = "UKR"
	COUNTRY_ARE Country = "ARE"
	COUNTRY_GBR Country = "GBR"
	COUNTRY_UMI Country = "UMI"
	COUNTRY_URY Country = "URY"
	COUNTRY_UZB Country = "UZB"
	COUNTRY_VUT Country = "VUT"
	COUNTRY_VEN Country = "VEN"
	COUNTRY_VNM Country = "VNM"
	COUNTRY_VGB Country = "VGB"
	COUNTRY_VIR Country = "VIR"
	COUNTRY_WLF Country = "WLF"
	COUNTRY_ESH Country = "ESH"
	COUNTRY_YEM Country = "YEM"
	COUNTRY_ZMB Country = "ZMB"
	COUNTRY_ZWE Country = "ZWE"
)
