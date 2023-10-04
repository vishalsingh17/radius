//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

const (
	moduleName = "v20231001preview"
	moduleVersion = "v0.0.1"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{	
		ActionTypeInternal,
	}
}

// CertificateFormats - Represents certificate formats
type CertificateFormats string

const (
	// CertificateFormatsPem - PEM Certificate format
	CertificateFormatsPem CertificateFormats = "pem"
	// CertificateFormatsPfx - PFX Certificate format
	CertificateFormatsPfx CertificateFormats = "pfx"
)

// PossibleCertificateFormatsValues returns the possible values for the CertificateFormats const type.
func PossibleCertificateFormatsValues() []CertificateFormats {
	return []CertificateFormats{	
		CertificateFormatsPem,
		CertificateFormatsPfx,
	}
}

// CertificateTypes - Represents certificate types
type CertificateTypes string

const (
	// CertificateTypesCertificate - Certificate type
	CertificateTypesCertificate CertificateTypes = "certificate"
	// CertificateTypesPrivatekey - Private Key type
	CertificateTypesPrivatekey CertificateTypes = "privatekey"
	// CertificateTypesPublickey - Public Key type
	CertificateTypesPublickey CertificateTypes = "publickey"
)

// PossibleCertificateTypesValues returns the possible values for the CertificateTypes const type.
func PossibleCertificateTypesValues() []CertificateTypes {
	return []CertificateTypes{	
		CertificateTypesCertificate,
		CertificateTypesPrivatekey,
		CertificateTypesPublickey,
	}
}

// ContainerResourceProvisioning - Specifies how the underlying service/resource is provisioned and managed. Available values
// are 'internal', where Radius manages the lifecycle of the resource internally, and 'manual', where a user
// manages the resource.
type ContainerResourceProvisioning string

const (
	// ContainerResourceProvisioningInternal - The resource lifecycle will be managed internally by Radius
	ContainerResourceProvisioningInternal ContainerResourceProvisioning = "internal"
	// ContainerResourceProvisioningManual - The resource lifecycle will be managed by the user
	ContainerResourceProvisioningManual ContainerResourceProvisioning = "manual"
)

// PossibleContainerResourceProvisioningValues returns the possible values for the ContainerResourceProvisioning const type.
func PossibleContainerResourceProvisioningValues() []ContainerResourceProvisioning {
	return []ContainerResourceProvisioning{	
		ContainerResourceProvisioningInternal,
		ContainerResourceProvisioningManual,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication CreatedByType = "Application"
	CreatedByTypeKey CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{	
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DaprSidecarExtensionProtocol - The Dapr sidecar extension protocol
type DaprSidecarExtensionProtocol string

const (
	// DaprSidecarExtensionProtocolGrpc - gRPC protocol
	DaprSidecarExtensionProtocolGrpc DaprSidecarExtensionProtocol = "grpc"
	// DaprSidecarExtensionProtocolHTTP - HTTP protocol
	DaprSidecarExtensionProtocolHTTP DaprSidecarExtensionProtocol = "http"
)

// PossibleDaprSidecarExtensionProtocolValues returns the possible values for the DaprSidecarExtensionProtocol const type.
func PossibleDaprSidecarExtensionProtocolValues() []DaprSidecarExtensionProtocol {
	return []DaprSidecarExtensionProtocol{	
		DaprSidecarExtensionProtocolGrpc,
		DaprSidecarExtensionProtocolHTTP,
	}
}

// IAMKind - The kind of IAM provider to configure
type IAMKind string

const (
	// IAMKindAzure - Azure Active Directory
	IAMKindAzure IAMKind = "azure"
)

// PossibleIAMKindValues returns the possible values for the IAMKind const type.
func PossibleIAMKindValues() []IAMKind {
	return []IAMKind{	
		IAMKindAzure,
	}
}

// IdentitySettingKind - IdentitySettingKind is the kind of supported external identity setting
type IdentitySettingKind string

const (
	// IdentitySettingKindAzureComWorkload - azure ad workload identity
	IdentitySettingKindAzureComWorkload IdentitySettingKind = "azure.com.workload"
	// IdentitySettingKindUndefined - undefined identity
	IdentitySettingKindUndefined IdentitySettingKind = "undefined"
)

// PossibleIdentitySettingKindValues returns the possible values for the IdentitySettingKind const type.
func PossibleIdentitySettingKindValues() []IdentitySettingKind {
	return []IdentitySettingKind{	
		IdentitySettingKindAzureComWorkload,
		IdentitySettingKindUndefined,
	}
}

// ImagePullPolicy - The image pull policy for the container
type ImagePullPolicy string

const (
	// ImagePullPolicyAlways - Always
	ImagePullPolicyAlways ImagePullPolicy = "Always"
	// ImagePullPolicyIfNotPresent - IfNotPresent
	ImagePullPolicyIfNotPresent ImagePullPolicy = "IfNotPresent"
	// ImagePullPolicyNever - Never
	ImagePullPolicyNever ImagePullPolicy = "Never"
)

// PossibleImagePullPolicyValues returns the possible values for the ImagePullPolicy const type.
func PossibleImagePullPolicyValues() []ImagePullPolicy {
	return []ImagePullPolicy{	
		ImagePullPolicyAlways,
		ImagePullPolicyIfNotPresent,
		ImagePullPolicyNever,
	}
}

// ManagedStore - The managed store for the ephemeral volume
type ManagedStore string

const (
	// ManagedStoreDisk - Disk store
	ManagedStoreDisk ManagedStore = "disk"
	// ManagedStoreMemory - Memory store
	ManagedStoreMemory ManagedStore = "memory"
)

// PossibleManagedStoreValues returns the possible values for the ManagedStore const type.
func PossibleManagedStoreValues() []ManagedStore {
	return []ManagedStore{	
		ManagedStoreDisk,
		ManagedStoreMemory,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem Origin = "system"
	OriginUser Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{	
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PortProtocol - The protocol in use by the port
type PortProtocol string

const (
	// PortProtocolTCP - TCP protocol
	PortProtocolTCP PortProtocol = "TCP"
	// PortProtocolUDP - UDP protocol
	PortProtocolUDP PortProtocol = "UDP"
)

// PossiblePortProtocolValues returns the possible values for the PortProtocol const type.
func PossiblePortProtocolValues() []PortProtocol {
	return []PortProtocol{	
		PortProtocolTCP,
		PortProtocolUDP,
	}
}

// ProvisioningState - Provisioning state of the portable resource at the time the operation was called
type ProvisioningState string

const (
	// ProvisioningStateAccepted - The resource create request has been accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - The resource is being deleted
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The resource is being provisioned
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The resource is updating
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{	
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceProvisioning - Specifies how the underlying service/resource is provisioned and managed. Available values are 'recipe',
// where Radius manages the lifecycle of the resource through a Recipe, and 'manual', where a user
// manages the resource and provides the values.
type ResourceProvisioning string

const (
	// ResourceProvisioningManual - The resource lifecycle will be managed by the user
	ResourceProvisioningManual ResourceProvisioning = "manual"
	// ResourceProvisioningRecipe - The resource lifecycle will be managed by Radius
	ResourceProvisioningRecipe ResourceProvisioning = "recipe"
)

// PossibleResourceProvisioningValues returns the possible values for the ResourceProvisioning const type.
func PossibleResourceProvisioningValues() []ResourceProvisioning {
	return []ResourceProvisioning{	
		ResourceProvisioningManual,
		ResourceProvisioningRecipe,
	}
}

// SecretStoreDataType - The type of SecretStore data
type SecretStoreDataType string

const (
	// SecretStoreDataTypeCertificate - Certificate secret data type
	SecretStoreDataTypeCertificate SecretStoreDataType = "certificate"
	// SecretStoreDataTypeGeneric - Generic secret data type
	SecretStoreDataTypeGeneric SecretStoreDataType = "generic"
)

// PossibleSecretStoreDataTypeValues returns the possible values for the SecretStoreDataType const type.
func PossibleSecretStoreDataTypeValues() []SecretStoreDataType {
	return []SecretStoreDataType{	
		SecretStoreDataTypeCertificate,
		SecretStoreDataTypeGeneric,
	}
}

// SecretValueEncoding - The type of SecretValue Encoding
type SecretValueEncoding string

const (
	// SecretValueEncodingBase64 - The base64-encoded secret value
	SecretValueEncodingBase64 SecretValueEncoding = "base64"
	// SecretValueEncodingRaw - The raw secret value
	SecretValueEncodingRaw SecretValueEncoding = "raw"
)

// PossibleSecretValueEncodingValues returns the possible values for the SecretValueEncoding const type.
func PossibleSecretValueEncodingValues() []SecretValueEncoding {
	return []SecretValueEncoding{	
		SecretValueEncodingBase64,
		SecretValueEncodingRaw,
	}
}

// TLSMinVersion - Tls Minimum versions for Gateway resource.
type TLSMinVersion string

const (
	// TLSMinVersionTls12 - TLS Version 1.2
	TLSMinVersionTls12 TLSMinVersion = "1.2"
	// TLSMinVersionTls13 - TLS Version 1.3
	TLSMinVersionTls13 TLSMinVersion = "1.3"
)

// PossibleTLSMinVersionValues returns the possible values for the TLSMinVersion const type.
func PossibleTLSMinVersionValues() []TLSMinVersion {
	return []TLSMinVersion{	
		TLSMinVersionTls12,
		TLSMinVersionTls13,
	}
}

// Versions - Supported API versions for the Applications.Core resource provider.
type Versions string

const (
	// VersionsV20231001Preview - 2023-10-01-preview
	VersionsV20231001Preview Versions = "2023-10-01-preview"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{	
		VersionsV20231001Preview,
	}
}

// VolumePermission - The persistent volume permission
type VolumePermission string

const (
	// VolumePermissionRead - Read only
	VolumePermissionRead VolumePermission = "read"
	// VolumePermissionWrite - Read and write
	VolumePermissionWrite VolumePermission = "write"
)

// PossibleVolumePermissionValues returns the possible values for the VolumePermission const type.
func PossibleVolumePermissionValues() []VolumePermission {
	return []VolumePermission{	
		VolumePermissionRead,
		VolumePermissionWrite,
	}
}

// VolumeSecretEncodings - Represents secret encodings
type VolumeSecretEncodings string

const (
	// VolumeSecretEncodingsBase64 - Base64 encoding
	VolumeSecretEncodingsBase64 VolumeSecretEncodings = "base64"
	// VolumeSecretEncodingsHex - Hex encoding
	VolumeSecretEncodingsHex VolumeSecretEncodings = "hex"
	// VolumeSecretEncodingsUTF8 - UTF-8 encoding
	VolumeSecretEncodingsUTF8 VolumeSecretEncodings = "utf-8"
)

// PossibleVolumeSecretEncodingsValues returns the possible values for the VolumeSecretEncodings const type.
func PossibleVolumeSecretEncodingsValues() []VolumeSecretEncodings {
	return []VolumeSecretEncodings{	
		VolumeSecretEncodingsBase64,
		VolumeSecretEncodingsHex,
		VolumeSecretEncodingsUTF8,
	}
}

