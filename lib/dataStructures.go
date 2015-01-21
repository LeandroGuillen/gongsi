package ngsi

type NotifyConditionType int
type UpdateActionType int

const (
	ONTIMEINTERVAL NotifyConditionType = 1
	ONVALUE
	ONCHANGE
)

const (
	UPDATE UpdateActionType = 1
	APPEND
	DELETE
)

// <xs:complexType name="ContextMetadata">
// <xs:sequence>
// <xs:element name="name" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="type" type="xs:anyURI" minOccurs="0" maxOccurs="1"/>
// <xs:element name="value" type="xs:anyType" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextMetadata struct {
	Name  string `json:"name"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value"`
}

// <xs:complexType name="ContextAttribute">
// <xs:sequence>
// <xs:element name="name" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="type" type="xs:anyURI" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextValue" type="xs:anyType" minOccurs="1" maxOccurs="1"/>
// <xs:element name="metadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextAttribute struct {
	Name      string            `json:"name"`
	Type      string            `json:"type,omitempty"`
	Value     string            `json:"value"`
	Metadatas []ContextMetadata `json:"metadatas,omitempty"`
}

// <xs:complexType name="EntityId">
// <xs:sequence>
// <xs:element name="id" type="xs:string" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// <xs:attribute name="type" type="xs:anyURI"/>
// <xs:attribute name="isPattern" type="xs:boolean"/>
// </xs:complexType>

type EntityId struct {
	Name      string `json:"name"`
	Type      string `json:"type,omitempty"`
	IsPattern bool   `json:"isPattern"`
}

// <xs:complexType name="ContextElement">
// <xs:sequence>
// <xs:element name="entityId" type="EntityId" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeDomainName" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextAttributeList" type="ContextAttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="domainMetadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextElement struct {
	Id                  string             `json:"id"`
	Type                string             `json:"type,omitempty"`
	IsPattern           bool               `json:"isPattern"`
	AttributeDomainName string             `json:"attributeDomainName,omitempty"`
	Attributes          []ContextAttribute `json:"attributes,omitempty"`
	Metadatas           []ContextMetadata  `json:"metadatas,omitempty"`
}

// <xs:complexType name="StatusCode">
// <xs:sequence>
// <xs:element name="code" type="xs:int" minOccurs="1" maxOccurs="1"/>
// <xs:element name="reasonPhrase" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="details" type="xs:anyType" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type StatusCode struct {
	Code         int    `json:"name"`
	ReasonPhrase string `json:"type,omitempty"`
	Details      string `json:"details,omitempty"`
}

// <xs:complexType name="SubscribeError">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeError struct {
	SubscriptionId string `json:"subscriptionId,omitempty"`
	ErrorCode      string `json:"errorCode"`
}

// <xs:complexType name="NotifyCondition">
// <xs:sequence>
// <xs:element name="type" type="NotifyConditionType" minOccurs="1" maxOccurs="1"/>
// <xs:element name="condValueList" type="CondValueList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="xs:string" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type NotifyCondition struct {
	Type NotifyConditionType `json:"name"`
	// CondValues  []ValueList         `json:"condValues,omitempty"`
	Restriction string `json:"restriction,omitempty"`
}

// <xs:complexType name="OperationScope">
// <xs:sequence>
// <xs:element name="scopeType" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="scopeValue" type="xs:anyType" maxOccurs="1" minOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type OperationScope struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// <xs:complexType name="Restriction">
// <xs:sequence>
// <xs:element name="attributeExpression" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="scope" type="OperationScopeList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type Restriction struct {
	AttributeExpression string           `json:"attributeExpression"`
	Scope               []OperationScope `json:"scope,omitempty"`
}

// <xs:complexType name="SubscribeResponse">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="throttling" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeResponse struct {
	SubscriptionId string `json:"subscriptionId"`
	Duration       string `json:"duration,omitempty"`
	Throttling     string `json:"throttling,omitempty"`
}

// <xs:complexType name="ContextRegistrationAttribute">
// <xs:sequence>
// <xs:element name="name" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="type" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="isDomain" type="xs:boolean" minOccurs="1" maxOccurs="1"/>
// <xs:element name="metadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextRegistrationAttribute struct {
	Name      string            `json:"name"`
	Type      string            `json:"type,omitempty"`
	IsDomain  bool              `json:"isDomain"`
	Metadatas []ContextMetadata `json:"metadatas,omitempty"`
}

// <xs:complexType name="ContextRegistration">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextRegistrationAttributeList" type="ContextRegistrationAttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="registrationMetadata" type="RegistrationMetadata" minOccurs="0" maxOccurs="1"/>
// <xs:element name="providingApplication" type="xs:anyURI" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextRegistration struct {
	EntityIdList                  []EntityId                     `json:"EntityIdList,omitempty"`
	ContextRegistrationAttributes []ContextRegistrationAttribute `json:"ContextRegistrationAttributes,omitempty"`
	Metadatas                     []ContextMetadata              `json:"metadatas,omitempty"`
	ProvidingApplication          string                         `json:"ProvidingApplication"`
}

// <xs:complexType name="ContextElementResponse">
// <xs:sequence>
// <xs:element name="contextElement" type="ContextElement" minOccurs="1" maxOccurs="1"/>
// <xs:element name="statusCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextElementResponse struct {
	ContextElement ContextElement `json:"contextElement"`
	StatusCode     StatusCode     `json:"statusCode"`
}

// <xs:complexType name="ContextRegistrationResponse">
// <xs:sequence>
// <xs:element name="contextRegistration" type="ContextRegistration" minOccurs="1" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type ContextRegistrationResponse struct {
	ContextRegistration ContextRegistration `json:"contextRegistration"`
	ErrorCode           StatusCode          `json:"errorCode"`
}
