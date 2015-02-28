package gongsi

// <xs:complexType name="QueryContextRequest">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeList" type="AttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type QueryContextRequest struct {
	Entities    []ContextElement `json:"entities"`
	Attributes  []string         `json:"attributes,omitempty"`
	Restriction Restriction      `json:"restriction,omitempty"`
}

// <xs:complexType name="QueryContextResponse">
// <xs:sequence>
// <xs:element name="contextResponseList" type="ContextElementResponseList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type QueryContextResponse struct {
	ContextResponses []ContextElementResponse `json:"contextResponses,omitempty"`
	ErrorCode        StatusCode               `json:"errorCode,omitempty"`
}

// <xs:complexType name="SubscribeContextRequest">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeList" type="AttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="reference" type="xs:anyType" minOccurs="1" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// <xs:element name="notifyConditions" type="NotifyConditionList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="throttling" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeContextRequest struct {
	Entities         []EntityId         `json:"entities"`
	Attributes       []ContextAttribute `json:"attributes,omitempty"`
	Reference        string             `json:"reference"`
	Duration         string             `json:"duration,omitempty"`
	Restriction      Restriction        `json:"restriction,omitempty"`
	NotifyConditions []NotifyCondition  `json:"notifyConditions,omitempty"`
	Throttling       string             `json:"throttling,omitempty"`
}

// <xs:complexType name="SubscribeContextResponse">
// <xs:sequence>
// <xs:element name="subscribeResponse" type="SubscribeResponse" minOccurs="0" maxOccurs="1"/>
// <xs:element name="subscribeError" type="SubscribeError" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeContextResponse struct {
	SubscribeResponse `json:"subscribeResponse,omitempty"`
	SubscribeError    `json:"subscribeError,omitempty"`
}

// <xs:complexType name="UpdateContextSubscriptionRequest">
// <xs:sequence>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="notifyConditions" type="NotifyConditionList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="throttling" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextSubscriptionRequest struct {
	SubscriptionId   string            `json:"subscriptionId"`
	Duration         string            `json:"duration,omitempty"`
	Restriction      Restriction       `json:"restriction,omitempty"`
	NotifyConditions []NotifyCondition `json:"notifyConditions,omitempty"`
	Throttling       string            `json:"throttling,omitempty"`
}

// <xs:complexType name="UpdateContextSubscriptionResponse">
// <xs:sequence>
// <xs:element name="subscribeResponse" type="SubscribeResponse" minOccurs="0" maxOccurs="1"/>
// <xs:element name="subscribeError" type="SubscribeError" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextSubscriptionResponse struct {
	SubscribeResponse `json:"subscribeResponse,omitempty"`
	SubscribeError    `json:"subscribeError,omitempty"`
}

// <xs:complexType name="UnsubscribeContextRequest">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UnsubscribeContextRequest struct {
	SubscriptionId string `json:"subscriptionId"`
}

// <xs:complexType name="UnsubscribeContextResponse">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="statusCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UnsubscribeContextResponse struct {
	SubscriptionId string     `json:"subscriptionId"`
	StatusCode     StatusCode `json:"statusCode"`
}

// <xs:complexType name="NotifyContextRequest">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="originator" type="xs:anyURI" minOccurs="1" maxOccurs="1"/>
// <xs:element name="contextResponseList" type="ContextElementResponseList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type NotifyContextRequest struct {
	SubscriptionId   string                   `json:"subscriptionId"`
	Originator       string                   `json:"originator"`
	Restriction      Restriction              `json:"restriction,omitempty"`
	ContextResponses []ContextElementResponse `json:"contextResponses,omitempty"`
}

// <xs:complexType name="NotifyContextResponse">
// <xs:sequence>
// <xs:element name="responseCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type NotifyContextResponse struct {
	ResponseCode StatusCode `json:"responseCode"`
}

// <xs:complexType name="UpdateContextRequest">
// <xs:sequence>
// <xs:element name="contextElementList" type="ContextElementList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="updateAction" type="UpdateActionType" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextRequest struct {
	ContextElements []ContextElement `json:"contextElements"`
	UpdateAction    string           `json:"updateAction"`
}

// <xs:complexType name="UpdateContextResponse">
// <xs:sequence>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextResponseList" type="ContextElementResponseList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextResponse struct {
	ContextResponses []ContextElementResponse `json:"contextResponses"`
	ErrorCode        StatusCode               `json:"errorCode,omitempty"`
}

// <!--
//  _____________________________________ operation message bodies for convenience operations______________________________________________
// -->
// <xs:complexType name="UpdateContextElementRequest">
// <xs:sequence>
// <xs:element name="attributeDomainName" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextAttributeList" type="ContextAttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="domainMetadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <xs:complexType name="UpdateContextElementResponse">
// <xs:sequence>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextResponseList" type="ContextAttributeResponseList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <xs:complexType name="AppendContextElementRequest">
// <xs:sequence>
// <xs:element name="attributeDomainName" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextAttributeList" type="ContextAttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="domainMetadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <xs:complexType name="AppendContextElementResponse">
// <xs:sequence>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextResponseList" type="ContextAttributeResponseList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <xs:complexType name="UpdateContextAttributeRequest">
// <xs:sequence>
// <xs:element name="type" type="xs:string" minOccurs="0" maxOccurs="1"/>
// <xs:element name="contextValue" type="xs:anyType" minOccurs="1" maxOccurs="1"/>
// <xs:element name="metadata" type="ContextMetadataList" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <!--
//  _________________________________ additional complex types used by convenience operations _______________________________________________________________
// -->
// <xs:complexType name="ContextAttributeResponse">
// <xs:sequence>
// <xs:element name="contextAttributeList" type="ContextAttributeList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="statusCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// <xs:complexType name="ContextAttributeResponseList">
// <xs:sequence>
// <xs:element name="contextAttributeResponse" type="ContextAttributeResponse" minOccurs="0" maxOccurs="unbounded"/>
// </xs:sequence>
// </xs:complexType>
// </xs:schema>
