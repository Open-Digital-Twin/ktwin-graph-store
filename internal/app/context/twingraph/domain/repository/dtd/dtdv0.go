package dtd

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TwinInstancePhase string

const (
	TwinInstancePhasePending TwinInstancePhase = "Pending"
	TwinInstancePhaseUnknown TwinInstancePhase = "Unknown"
	TwinInstancePhaseRunning TwinInstancePhase = "Running"
	TwinInstancePhaseFailed  TwinInstancePhase = "Failed"
)

// TwinInstanceSpec defines the desired state of TwinInstance
type TwinInstanceSpec struct {
	Interface                 string                        `json:"interface,omitempty"`
	Events                    []TwinInstanceEvents          `json:"events,omitempty"`
	EndpointSettings          *TwinInstanceEndpointSettings `json:"endpointSettings,omitempty"`
	Data                      *TwinInstanceDataSpec         `json:"data,omitempty"`
	TwinInstanceRelationships []TwinInstanceRelationship    `json:"twinInstanceRelationships,omitempty"`
}

type TwinInstanceDataSpec struct {
	Properties  []TwinInstancePropertyData  `json:"properties,omitempty"`  // TODO: read-only
	Telemetries []TwinInstanceTelemetryData `json:"telemetries,omitempty"` // TODO: read-only
}

// TODO: read-only
type TwinInstancePropertyData struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value"`
}

// TODO: read-only
type TwinInstanceTelemetryData struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value"`
}

type TwinInstanceRelationship struct {
	// The TwinInstance Relationship name
	Name string `json:"name"`
	// The Target TwinInterface of the Relationship
	Interface string `json:"interface"`
	// The Target TwinInstance of the Relationship
	Instance string `json:"instance"`
}

// TODO: Configure as read-only
type TwinInstanceEndpointSettings struct {
	HttpEndpoint *TwinInstanceHttpEndpointSettings `json:"httpEndpoint,omitempty"`
	MqttEndpoint *TwinInstanceMqttEndpointSettings `json:"mqttEndpoint,omitempty"`
	AmqpEndpoint *TwinInstanceAmqpEndpointSettings `json:"amqpEndpoint,omitempty"`
}

// TODO: Configure as read-only
type TwinInstanceHttpEndpointSettings struct {
	Url string `json:"url,omitempty"`
}

// TODO: Configure as read-only
type TwinInstanceMqttEndpointSettings struct {
	Url             string `json:"url,omitempty"`
	PublisherTopic  string `json:"publisherTopic,omitempty"`
	SubscriberTopic string `json:"subscriberTopic,omitempty"`
}

type TwinInstanceAmqpEndpointSettings struct {
	Url             string `json:"url,omitempty"`
	PublisherTopic  string `json:"publisherTopic,omitempty"`
	SubscriberTopic string `json:"subscriberTopic,omitempty"`
}

type TwinInstanceEvents struct {
	Filters TwinInstanceEventsFilters `json:"filters,omitempty"`
	Sink    TwinInterfaceEventsSink   `json:"sink,omitempty"`
}

// Based on CN Cloud Event Filters definitions: https://github.com/cloudevents/spec/blob/main/subscriptions/spec.md#324-filters
type TwinInstanceEventsFilters struct {
	Exact TwinInstanceEventsFiltersAttributes `json:"exact,omitempty"`
	// Prefix TwinInstanceEventsFiltersProperties `json:"prefix,omitempty"` // Unsupported
	// Suffix TwinInstanceEventsFiltersProperties `json:"suffix,omitempty"` // Unsupported
	// All    TwinInstanceEventsFiltersProperties `json:"all,omitempty"` // Unsupported
	// Any    TwinInstanceEventsFiltersProperties `json:"any,omitempty"` // Unsupported
	// Not    TwinInstanceEventsFiltersProperties `json:"not,omitempty"` // Unsupported
}

type TwinInstanceEventsFiltersAttributes map[string]string

type TwinInterfaceEventsSink struct {
	InstanceId string `json:"instanceId,omitempty"`
}

// TwinInstanceStatus defines the observed state of TwinInstance
type TwinInstanceStatus struct {
	Status TwinInstancePhase `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TwinInstance is the Schema for the twininstances API
type TwinInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TwinInstanceSpec   `json:"spec,omitempty"`
	Status TwinInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TwinInstanceList contains a list of TwinInstance
type TwinInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TwinInstance `json:"items"`
}
