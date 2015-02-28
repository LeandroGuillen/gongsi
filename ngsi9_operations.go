package gongsi

// <xs:complexType name="RegisterContextRequest">
// <xs:sequence>
// <xs:element name="contextRegistrationList" type="ContextRegistrationList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="registrationId" type="xs:string" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type RegisterContextRequest struct {
	ContextRegistrations []ContextRegistration `json:"contextRegistrations,omitempty"`
	Duration             string                `json:"duration,omitempty"`
	RegistrationId       string                `json:"registrationId,omitempty"`
}

// <xs:complexType name="RegisterContextResponse">
// <xs:sequence>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="registrationId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type RegisterContextResponse struct {
	Duration       string     `json:"duration,omitempty"`
	RegistrationId string     `json:"registrationId"`
	ErrorCode      StatusCode `json:"errorCode,omitempty"`
}

// <xs:complexType name="DiscoveryContextAvailabilityRequest">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeList" type="AttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type DiscoverContextAvailabilityRequest struct {
	Entities    []EntityId         `json:"entities"`
	Attributes  []ContextAttribute `json:"attributes,omitempty"`
	Restriction Restriction        `json:"restriction,omitempty"`
}

// <xs:complexType name="DiscoveryContextAvailabilityResponse">
// <xs:sequence>
// <xs:element name="contextRegistrationResponseList" type="ContextRegistrationResponseList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type DiscoverContextAvailabilityResponse struct {
	ContextRegistrationResponses []ContextRegistrationResponse `json:"contextRegistrationResponses,omitempty"`
	ErrorCode                    StatusCode                    `json:"errorCode,omitempty"`
}

// <xs:complexType name="SubscribeContextAvailabilityRequest">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeList" type="AttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="reference" type="xs:anyURI" minOccurs="1" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeContextAvailabilityRequest struct {
	Entities       []EntityId         `json:"entities"`
	Attributes     []ContextAttribute `json:"attributes,omitempty"`
	Reference      string             `json:"reference"`
	Duration       string             `json:"duration,omitempty"`
	Restriction    Restriction        `json:"restriction,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
}

// <xs:complexType name="SubscribeContextAvailabilityResponse">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type SubscribeContextAvailabilityResponse struct {
	SubscriptionId string     `json:"subscriptionId"`
	Duration       string     `json:"duration,omitempty"`
	ErrorCode      StatusCode `json:"errorCode,omitempty"`
}

// <xs:complexType name="UpdateContextAvailabilitySubscriptionRequest">
// <xs:sequence>
// <xs:element name="entityIdList" type="EntityIdList" minOccurs="1" maxOccurs="1"/>
// <xs:element name="attributeList" type="AttributeList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="restriction" type="Restriction" minOccurs="0" maxOccurs="1"/>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextAvailabilitySubscriptionRequest struct {
	Entities       []EntityId         `json:"entities"`
	Attributes     []ContextAttribute `json:"attributes,omitempty"`
	Duration       string             `json:"duration,omitempty"`
	Restriction    Restriction        `json:"restriction,omitempty"`
	SubscriptionId string             `json:"subscriptionId,omitempty"`
}

// <xs:complexType name="UpdateContextAvailabilitySubscriptionResponse">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="duration" type="xs:duration" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UpdateContextAvailabilitySubscriptionResponse struct {
	SubscriptionId string     `json:"subscriptionId"`
	Duration       string     `json:"duration,omitempty"`
	ErrorCode      StatusCode `json:"errorCode,omitempty"`
}

// <xs:complexType name="UnsubscribeContextAvailabilityRequest">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UnsubscribeContextAvailabilityRequest struct {
	SubscriptionId string `json:"subscriptionId"`
}

// <xs:complexType name="UnsubscribeContextAvailabilityResponse">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="statusCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type UnsubscribeContextAvailabilityResponse struct {
	SubscriptionId string     `json:"subscriptionId"`
	StatusCode     StatusCode `json:"statusCode"`
}

// <xs:complexType name="NotifyContextAvailabilityRequest">
// <xs:sequence>
// <xs:element name="subscriptionId" type="xs:string" minOccurs="1" maxOccurs="1"/>
// <xs:element name="contextRegistrationResponseList" type="ContextRegistrationResponseList" minOccurs="0" maxOccurs="1"/>
// <xs:element name="errorCode" type="StatusCode" minOccurs="0" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>

type NotifyContextAvailabilityRequest struct {
	SubscriptionId       string                `json:"subscriptionId"`
	ContextRegistrations []ContextRegistration `json:"contextRegistrations,omitempty"`
	ErrorCode            StatusCode            `json:"errorCode,omitempty"`
}

// <xs:complexType name="NotifyContextAvailabilityResponse">
// <xs:sequence>
// <xs:element name="responseCode" type="StatusCode" minOccurs="1" maxOccurs="1"/>
// </xs:sequence>
// </xs:complexType>
// </xs:schema>

type NotifyContextAvailabilityResponse struct {
	ResponseCode StatusCode `json:"responseCode"`
}
