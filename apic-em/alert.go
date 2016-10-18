package apicem

import (
	"fmt"
	"strings"
)

const alertBasePath = "v1"

// AlarmService is an interface with the Alarm API
type AlarmService service

// AlarmDto is ...
type AlarmDto struct {
	CausedAlarmsURL               string                 `json:"causedAlarmsUrl,omitempty"`
	CausingAlarmURL               string                 `json:"causingAlarmUrl,omitempty"`
	AlarmCreationTime             string                 `json:"alarmCreationTime,omitempty"`
	ApplicationSpecificAlarmID    string                 `json:"applicationSpecificAlarmID,omitempty"`
	EventCount                    int32                  `json:"eventCount,omitempty"`
	IsAcknowledged                bool                   `json:"isAcknowledged,omitempty"`
	LastModifiedTimestamp         string                 `json:"lastModifiedTimestamp,omitempty"`
	MayBeAutoCleared              bool                   `json:"mayBeAutoCleared,omitempty"`
	OwnerID                       string                 `json:"ownerID,omitempty"`
	PreviousSeverity              string                 `json:"previousSeverity,omitempty"`
	SubclassName                  string                 `json:"subclassName,omitempty"`
	CausedAlarms                  []IDRef                `json:"causedAlarms,omitempty"`
	CausingAlarm                  IDRef                  `json:"causingAlarm,omitempty"`
	Description                   string                 `json:"description,omitempty"`
	Severity                      string                 `json:"severity,omitempty"`
	Source                        string                 `json:"source,omitempty"`
	Category                      EventAlarmCategoryEnum `json:"category,omitempty"`
	EventType                     EventTypeEnum          `json:"eventType,omitempty"`
	AlarmDisplayable              bool                   `json:"alarmDisplayable,omitempty"`
	ApplicationCategoryData       string                 `json:"applicationCategoryData,omitempty"`
	DeviceTimestamp               string                 `json:"deviceTimestamp,omitempty"`
	IsDeviceMaster                bool                   `json:"isDeviceMaster,omitempty"`
	NotificationDeliveryMechanism string                 `json:"notificationDeliveryMechanism,omitempty"`
	NotificationState             int32                  `json:"notificationState,omitempty"`
	PendingUntil                  string                 `json:"pendingUntil,omitempty"`
	ReportingEntityAddress        InetAddress            `json:"reportingEntityAddress,omitempty"`
	SourceMacAddress              MacAddress             `json:"sourceMacAddress,omitempty"`
	SrcObjectClassID              int32                  `json:"srcObjectClassId,omitempty"`
	SrcObjectID                   int64                  `json:"srcObjectId,omitempty"`
	Stability                     string                 `json:"stability,omitempty"`
	TransientNameValue            string                 `json:"transientNameValue,omitempty"`
	EventAlarmDetails             []IDRef                `json:"eventAlarmDetails,omitempty"`
	EventAlarmDetailsURL          string                 `json:"eventAlarmDetailsUrl,omitempty"`
	InstanceUUID                  string                 `json:"instanceUuid,omitempty"`
	ID                            string                 `json:"id,omitempty"`
}

// Annotation is ...
type Annotation struct {
}

// CertPath is ...
type CertPath struct {
	Encodings    Iteratorstring `json:"encodings,omitempty"`
	Certificates []Certificate  `json:"certificates,omitempty"`
	Type         string         `json:"type,omitempty"`
	Encoded      []byte         `json:"encoded,omitempty"`
}

// Certificate is ...
type Certificate struct {
	Type      string    `json:"type,omitempty"`
	Encoded   []byte    `json:"encoded,omitempty"`
	PublicKey PublicKey `json:"publicKey,omitempty"`
}

// Class is ...
type Class struct {
	Modifiers            int32             `json:"modifiers,omitempty"`
	Interface            bool              `json:"interface,omitempty"`
	Array                bool              `json:"array,omitempty"`
	Primitive            bool              `json:"primitive,omitempty"`
	Superclass           ClassObject       `json:"superclass,omitempty"`
	ComponentType        ClassObject       `json:"componentType,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Annotation           bool              `json:"annotation,omitempty"`
	Synthetic            bool              `json:"synthetic,omitempty"`
	ClassLoader          ClassLoader       `json:"classLoader,omitempty"`
	TypeParameters       []TypeVariable    `json:"typeParameters,omitempty"`
	GenericSuperclass    Type              `json:"genericSuperclass,omitempty"`
	Package              Package           `json:"package,omitempty"`
	Interfaces           []Class           `json:"interfaces,omitempty"`
	GenericInterfaces    []Type            `json:"genericInterfaces,omitempty"`
	Signers              []string          `json:"signers,omitempty"`
	EnclosingMethod      Method            `json:"enclosingMethod,omitempty"`
	EnclosingConstructor ConstructorObject `json:"enclosingConstructor,omitempty"`
	DeclaringClass       ClassObject       `json:"declaringClass,omitempty"`
	EnclosingClass       ClassObject       `json:"enclosingClass,omitempty"`
	SimpleName           string            `json:"simpleName,omitempty"`
	CanonicalName        string            `json:"canonicalName,omitempty"`
	AnonymousClass       bool              `json:"anonymousClass,omitempty"`
	LocalClass           bool              `json:"localClass,omitempty"`
	MemberClass          bool              `json:"memberClass,omitempty"`
	Classes              []Class           `json:"classes,omitempty"`
	Fields               []Field           `json:"fields,omitempty"`
	Methods              []Method          `json:"methods,omitempty"`
	Constructors         []Constructor     `json:"constructors,omitempty"`
	DeclaredClasses      []Class           `json:"declaredClasses,omitempty"`
	DeclaredFields       []Field           `json:"declaredFields,omitempty"`
	DeclaredMethods      []Method          `json:"declaredMethods,omitempty"`
	DeclaredConstructors []Constructor     `json:"declaredConstructors,omitempty"`
	ProtectionDomain     ProtectionDomain  `json:"protectionDomain,omitempty"`
	Enum                 bool              `json:"enum,omitempty"`
	EnumConstants        []string          `json:"enumConstants,omitempty"`
	Annotations          []Annotation      `json:"annotations,omitempty"`
	DeclaredAnnotations  []Annotation      `json:"declaredAnnotations,omitempty"`
}

// ClassLoader is ...
type ClassLoader struct {
	Parent *ClassLoader `json:"parent,omitempty"`
}

// ClassObject is ...
type ClassObject struct {
	Modifiers            int32              `json:"modifiers,omitempty"`
	Interface            bool               `json:"interface,omitempty"`
	Array                bool               `json:"array,omitempty"`
	Primitive            bool               `json:"primitive,omitempty"`
	Superclass           *ClassObject       `json:"superclass,omitempty"`
	ComponentType        *ClassObject       `json:"componentType,omitempty"`
	Name                 string             `json:"name,omitempty"`
	Annotation           bool               `json:"annotation,omitempty"`
	Synthetic            bool               `json:"synthetic,omitempty"`
	ClassLoader          *ClassLoader       `json:"classLoader,omitempty"`
	TypeParameters       []*TypeVariable    `json:"typeParameters,omitempty"`
	GenericSuperclass    *Type              `json:"genericSuperclass,omitempty"`
	Package              *Package           `json:"package,omitempty"`
	Interfaces           []*Class           `json:"interfaces,omitempty"`
	GenericInterfaces    []*Type            `json:"genericInterfaces,omitempty"`
	Signers              []string           `json:"signers,omitempty"`
	EnclosingMethod      *Method            `json:"enclosingMethod,omitempty"`
	EnclosingConstructor *ConstructorObject `json:"enclosingConstructor,omitempty"`
	DeclaringClass       *ClassObject       `json:"declaringClass,omitempty"`
	EnclosingClass       *ClassObject       `json:"enclosingClass,omitempty"`
	SimpleName           string             `json:"simpleName,omitempty"`
	CanonicalName        string             `json:"canonicalName,omitempty"`
	AnonymousClass       bool               `json:"anonymousClass,omitempty"`
	LocalClass           bool               `json:"localClass,omitempty"`
	MemberClass          bool               `json:"memberClass,omitempty"`
	Classes              []*Class           `json:"classes,omitempty"`
	Fields               []*Field           `json:"fields,omitempty"`
	Methods              []*Method          `json:"methods,omitempty"`
	Constructors         []*Constructor     `json:"constructors,omitempty"`
	DeclaredClasses      []*Class           `json:"declaredClasses,omitempty"`
	DeclaredFields       []*Field           `json:"declaredFields,omitempty"`
	DeclaredMethods      []*Method          `json:"declaredMethods,omitempty"`
	DeclaredConstructors []*Constructor     `json:"declaredConstructors,omitempty"`
	ProtectionDomain     *ProtectionDomain  `json:"protectionDomain,omitempty"`
	Enum                 bool               `json:"enum,omitempty"`
	EnumConstants        []string           `json:"enumConstants,omitempty"`
	Annotations          []*Annotation      `json:"annotations,omitempty"`
	DeclaredAnnotations  []*Annotation      `json:"declaredAnnotations,omitempty"`
}

// CodeSigner is ...
type CodeSigner struct {
	SignerCertPath CertPath  `json:"signerCertPath,omitempty"`
	Timestamp      Timestamp `json:"timestamp,omitempty"`
}

// CodeSource is ...
type CodeSource struct {
	Location     URL           `json:"location,omitempty"`
	Certificates []Certificate `json:"certificates,omitempty"`
	CodeSigners  []CodeSigner  `json:"codeSigners,omitempty"`
}

// Constructor is ...
type Constructor struct {
	Modifiers             int32          `json:"modifiers,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Synthetic             bool           `json:"synthetic,omitempty"`
	TypeParameters        []TypeVariable `json:"typeParameters,omitempty"`
	DeclaringClass        ClassObject    `json:"declaringClass,omitempty"`
	DeclaredAnnotations   []Annotation   `json:"declaredAnnotations,omitempty"`
	ParameterTypes        []Class        `json:"parameterTypes,omitempty"`
	GenericParameterTypes []Type         `json:"genericParameterTypes,omitempty"`
	ExceptionTypes        []Class        `json:"exceptionTypes,omitempty"`
	GenericExceptionTypes []Type         `json:"genericExceptionTypes,omitempty"`
	VarArgs               bool           `json:"varArgs,omitempty"`
	ParameterAnnotations  []Annotation   `json:"parameterAnnotations,omitempty"` // Not sure about this type
	Annotations           []Annotation   `json:"annotations,omitempty"`
	Accessible            bool           `json:"accessible,omitempty"`
}

// ConstructorObject is ...
type ConstructorObject struct {
	Modifiers             int32          `json:"modifiers,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Synthetic             bool           `json:"synthetic,omitempty"`
	TypeParameters        []TypeVariable `json:"typeParameters,omitempty"`
	DeclaringClass        ClassObject    `json:"declaringClass,omitempty"`
	DeclaredAnnotations   []Annotation   `json:"declaredAnnotations,omitempty"`
	ParameterTypes        []Class        `json:"parameterTypes,omitempty"`
	GenericParameterTypes []Type         `json:"genericParameterTypes,omitempty"`
	ExceptionTypes        []Class        `json:"exceptionTypes,omitempty"`
	GenericExceptionTypes []Type         `json:"genericExceptionTypes,omitempty"`
	VarArgs               bool           `json:"varArgs,omitempty"`
	ParameterAnnotations  []Annotation   `json:"parameterAnnotations,omitempty"` // Not sure about this type
	Annotations           []Annotation   `json:"annotations,omitempty"`
	Accessible            bool           `json:"accessible,omitempty"`
}

// EventAlarmCategoryEnum is ...
type EventAlarmCategoryEnum struct {
	Value   string `json:"value,omitempty"`
	Ordinal int32  `json:"ordinal,omitempty"`
}

// EventTypeEnum is ...
type EventTypeEnum struct {
	Value   string `json:"value,omitempty"`
	Ordinal int32  `json:"ordinal,omitempty"`
}

// Field is ...
type Field struct {
	Modifiers           int32        `json:"modifiers,omitempty"`
	Name                string       `json:"name,omitempty"`
	Synthetic           bool         `json:"synthetic,omitempty"`
	DeclaringClass      ClassObject  `json:"declaringClass,omitempty"`
	DeclaredAnnotations []Annotation `json:"declaredAnnotations,omitempty"`
	EnumConstant        bool         `json:"enumConstant,omitempty"`
	Type                ClassObject  `json:"type,omitempty"`
	GenericType         Type         `json:"genericType,omitempty"`
	Annotations         []Annotation `json:"annotations,omitempty"`
	Accessible          bool         `json:"accessible,omitempty"`
}

// GenericDeclaration is ...
type GenericDeclaration struct {
	TypeParameters []TypeVariable `json:"typeParameters,omitempty"`
}

// IDRef is ...
type IDRef struct {
	URL      string      `json:"url,omitempty"`
	ID       int64       `json:"id,omitempty"`
	Type     string      `json:"type,omitempty"`
	LongType ClassObject `json:"longType,omitempty"`
}

// InetAddress is ...
type InetAddress struct {
	Address       string `json:"address,omitempty"`
	Octets        []byte `json:"octets,omitempty"`
	PaddedAddress string `json:"paddedAddress,omitempty"`
	AddressType   string `json:"addressType,omitempty"`
	AllZeros      bool   `json:"allZeros,omitempty"`
	AllOnes       bool   `json:"allOnes,omitempty"`
}

// Iteratorstring is ...
type Iteratorstring struct {
}

// ListAlarmDtoResponse is ...
type ListAlarmDtoResponse struct {
	Version  string     `json:"version,omitempty"`
	Response []AlarmDto `json:"response,omitempty"`
}

// MacAddress is ...
type MacAddress struct {
	AllZeros bool   `json:"allZeros,omitempty"`
	AllOnes  bool   `json:"allOnes,omitempty"`
	Octets   []byte `json:"octets,omitempty"`
}

// Method is ...
type Method struct {
	Modifiers             int32          `json:"modifiers,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Synthetic             bool           `json:"synthetic,omitempty"`
	TypeParameters        []TypeVariable `json:"typeParameters,omitempty"`
	DeclaringClass        *ClassObject   `json:"declaringClass,omitempty"`
	DeclaredAnnotations   []Annotation   `json:"declaredAnnotations,omitempty"`
	ReturnType            *ClassObject   `json:"returnType,omitempty"`
	ParameterTypes        []Class        `json:"parameterTypes,omitempty"`
	GenericReturnType     Type           `json:"genericReturnType,omitempty"`
	GenericParameterTypes []Type         `json:"genericParameterTypes,omitempty"`
	ExceptionTypes        []Class        `json:"exceptionTypes,omitempty"`
	GenericExceptionTypes []Type         `json:"genericExceptionTypes,omitempty"`
	Bridge                bool           `json:"bridge,omitempty"`
	VarArgs               bool           `json:"varArgs,omitempty"`
	DefaultValue          string         `json:"defaultValue,omitempty"`
	ParameterAnnotations  []Annotation   `json:"parameterAnnotations,omitempty"` // Not sure about this type
	Annotations           []Annotation   `json:"annotations,omitempty"`
	Accessible            bool           `json:"accessible,omitempty"`
}

// Package is ...
type Package struct {
	Name                  string       `json:"name,omitempty"`
	Annotations           []Annotation `json:"annotations,omitempty"`
	DeclaredAnnotations   []Annotation `json:"declaredAnnotations,omitempty"`
	Sealed                bool         `json:"sealed,omitempty"`
	SpecificationTitle    string       `json:"specificationTitle,omitempty"`
	SpecificationVersion  string       `json:"specificationVersion,omitempty"`
	SpecificationVendor   string       `json:"specificationVendor,omitempty"`
	ImplementationTitle   string       `json:"implementationTitle,omitempty"`
	ImplementationVersion string       `json:"implementationVersion,omitempty"`
	ImplementationVendor  string       `json:"implementationVendor,omitempty"`
}

// PermissionCollection is ...
type PermissionCollection struct {
	ReadOnly bool `json:"readOnly,omitempty"`
}

// Principal is ...
type Principal struct {
	Name string `json:"name,omitempty"`
}

// ProtectionDomain is ...
type ProtectionDomain struct {
	ClassLoader ClassLoader          `json:"classLoader,omitempty"`
	CodeSource  CodeSource           `json:"codeSource,omitempty"`
	Principals  []Principal          `json:"principals,omitempty"`
	Permissions PermissionCollection `json:"permissions,omitempty"`
}

// PublicKey is ...
type PublicKey struct {
	Format    string `json:"format,omitempty"`
	Encoded   []byte `json:"encoded,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
}

// Timestamp is ...
type Timestamp struct {
	SignerCertPath CertPath `json:"signerCertPath,omitempty"`
	Timestamp      string   `json:"timestamp,omitempty"`
}

// Type is ...
type Type struct {
}

// TypeVariable is ...
type TypeVariable struct {
	Bounds             []Type             `json:"bounds,omitempty"`
	GenericDeclaration GenericDeclaration `json:"genericDeclaration,omitempty"`
	Name               string             `json:"name,omitempty"`
}

// URL is ...
type URL struct {
	Path        string `json:"path,omitempty"`
	Authority   string `json:"authority,omitempty"`
	Query       string `json:"query,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	File        string `json:"file,omitempty"`
	Host        string `json:"host,omitempty"`
	Ref         string `json:"ref,omitempty"`
	UserInfo    string `json:"userInfo,omitempty"`
	Port        int32  `json:"port,omitempty"`
	DefaultPort int32  `json:"defaultPort,omitempty"`
	Content     string `json:"content,omitempty"`
}

// AcknowledgeAlarm is ...
// This method is used to acknowledge an Alarm by UUID
//
//
//
//  * @param alarmDTO alarmDTO
// * @return *SuccessResult
func (s *AlarmService) AcknowledgeAlarm(alertUUID string, alarmDTO AlarmDto) (*SuccessResult, *Response, error) {

	path := alertBasePath + "/alarm/{alertUUID}"
	path = strings.Replace(path, "{"+"alertUUID"+"}", fmt.Sprintf("%v", alertUUID), -1)
	req, err := s.client.NewRequest("PUT", path, alarmDTO)
	if err != nil {
		return nil, nil, err
	}

	root := new(SuccessResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAlarmWithFilter is ...
// getAlarmWithFilter
//
//
//
//
// * @return *ListAlarmDtoResponse
func (s *AlarmService) GetAlarmWithFilter() (*ListAlarmDtoResponse, *Response, error) {

	path := alertBasePath + "/alarm"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ListAlarmDtoResponse)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

// GetAuditCountWithFilter is ...
// getAuditCountWithFilter
//
//
//
//
// * @return *int32
func (s *AlarmService) GetAuditCountWithFilter() (*int32, *Response, error) {

	path := alertBasePath + "/alarm/count"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(int32)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
