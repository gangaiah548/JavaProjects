package models

import (
	driver "github.com/arangodb/go-driver"
)

type Smpdata struct {
	Uuid         string  `json:"uuid"`
	Node         string  `json:"node"`
	Houseaddress Address `json:"houseaddress"`
}

type Address struct {
	Hno string `json:"hno"`
	Pin string `json:"pin"`
}

type KnowsEdge struct {
	driver.EdgeDocument
	Relation string `json:"relation"`
}

type JavaClass struct {
	PackageName string `json:"packageName"`
	Type        string `json:"type"`
	ClassName   string `json:"className"`
	Description string `json:"description"`
}

type Domaint struct {
	Uuid        string      `json:"uuid"`
	Domain      string      `json:"domain"`
	CreatedBy   string      `json:"createdBy"`
	Date        string      `json:"date"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	SubDomains  []SubDomain `json:"subdomains"`
	OpType      string      `json:"optype"`
}
type SubDomain struct {
	Subdomain    string        `json:"subdomain"`
	Description  string        `json:"description"`
	Version      string        `json:"version"`
	ValueObjects []ValueObject `json:"valueObjects"`
	Entities     []Entity      `json:"entities"`
	Aggregates   []Aggregate   `json:"aggregates"`
}
type Domain struct {
	Uuid        string      `json:"uuid"`
	Domain      string      `json:"domain"`
	CreatedBy   string      `json:"createdBy"`
	Date        string      `json:"date"`
	Version     string      `json:"version"`
	Description string      `json:"description"`
	SubDomains  []SubDomain `json:"subdomains"`
}

type JavaClasst struct {
	PackageName      string            `json:"packageName"`
	Type             string            `json:"type"`
	ClassName        string            `json:"className"`
	Description      string            `json:"description"`
	Imports          []string          `json:"imports"`
	TextendsClass    []ExtendsClass    `json:"textendsClass"`
	TimplementsClass []ImplementsClass `json:"implementsClass"`
	Fields           []Field           `json:"fields"`
	Methods          []Method          `json:"methods"`
}

type ExtendsClass struct {
	ClassName       string `json:"className"`
	GenericsPattern string `json:"genericsPattern"`
}

type ImplementsClass struct {
	InterfaceClassName string `json:"interfaceClassName"`
	GenericsPattern    string `json:"genericsPattern"`
}

type Annotation struct {
	AnnotationString string            `json:"annotationString"`
	AnnotationParams []AnnotationParam `json:"AnnotationParams"`
}

type AnnotationParam struct {
	ParamKey        string `json:"paramKey"`
	GenericsPattern string `json:"genericsPattern"`
}
type Field struct {
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	DataStructure string       `json:"dataStructure"`
	KeyDataType   string       `json:"KeyDataType"`
	ValueDataType string       `json:"valueDataType"`
	Annotations   []Annotation `json:"annotations"`
}

type Method struct {
	MethodName       string        `json:"methodName"`
	Description      string        `json:"description"`
	ReturnParam      MethodParam   `json:"returnParam"`
	MethodParams     []MethodParam `json:"methodParams"`
	ThrowsExceptions []string      `json:"throwsExceptions"`
	Annotations      []Annotation  `json:"annotations"`
	MethodCalls      []MethodCall  `json:"methodCalls"`
}

type MethodParam struct {
	ParamName   string       `json:"paramName"`
	DataType    string       `json:"dataType"`
	Annotations []Annotation `json:"annotations"`
}
type MethodCall struct {
	MethodName   string        `json:"methodName"`
	MethodParams []MethodParam `json:"methodParams"`
}

type ValueObject struct {
	ValueObject string      `json:"valueObject"`
	Description string      `json:"description"`
	Attributes  []Attribute `json:"attributes"`
}
type Entity struct {
	Entity         string      `json:"entity"`
	Description    string      `json:"description"`
	EntityCRUD     string      `json:"entityCRUD"`
	ControllerCRUD string      `json:"controllerCRUD"`
	Attributes     []Attribute `json:"attributes"`
}

type Aggregate struct {
	Aggregate      string      `json:"aggregate"`
	Description    string      `json:"description"`
	ControllerCRUD string      `json:"controllerCRUD"`
	Attributes     []Attribute `json:"attributes"`
}

type Attribute struct {
	Attribute      string `json:"attribute"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Lov            string `json:"lov"`
	AllowedChars   string `json:"allowedChars"`
	MinLength      string `json:"minLength"`
	MaxLength      string `json:"maxLength"`
	FileSize       string `json:"fileSize"`
	FileType       string `json:"FileType"`
	Validation     string `json:"validation"`
	Unique         bool   `json:"unique"`
	Nullable       bool   `json:"nullable"`
	EtityRead      bool   `json:"entityRead"`
	ControllerRead bool   `json:"controllerRead"`
}

type MenuInfo struct {
	MenuId          int64  `json:"menuId"`
	MenuName        string `json:"menuName"`
	MenuTc          string `json:"MenuTc"`
	Description     string `json:"description"`
	AccessUrl       string `json:"accessUrl"`
	IsLeaf          bool   `json:"isLeaf"`
	ApiModule       string `json:"apiModule"`
	StatusTc        bool   `json:"statusTc"`
	ParentMenuId    int64  `json:"parentMenuId"`
	DisplaySequence int    `json:"displaySequence"`
}

type RoleInfo struct {
	RoleId   int64      `json:"roleId"`
	Menus    []MenuInfo `json:"menus"`
	RoleName string     `json:"roleName"`
}

type TokenResponse struct {
	Token              string   `json:"token"`
	PersonaTc          string   `json:"personaTc"`
	PersonaReferenceId string   `json:"personaReferenceId"`
	TenantCode         string   `json:"tenantCode"`
	StatusTc           string   `json:"statusTc"`
	IsPasswordExpired  bool     `json:"isPasswordExpired"`
	Role               RoleInfo `json:"role"`
}

type DomainSubDomain interface {
	GetCommonField() interface{}
	// Add other common methods if needed
}

// ModelType1 implementation
func (m Domaint) GetCommonField() interface{} {
	return m
}

// ModelType2 implementation
func (m JavaClass) GetCommonField() interface{} {
	return m
}

var RoleMatrix = map[string][]string{
	"admin":   {"read", "write", "delete"},
	"user":    {"read"},
	"cmsuser": {"read", "write", "delete"},
	"guest":   {},
}

type UserInfo struct {
	Token               string   `json:"token"`
	RoleName            string   `json:"roleName"`
	RoleId              int64    `json:"roleId"`
	RequiredPermissions []string `json:"requiredPermissions"`
}

type UserWithRoleInfo struct {
	Token       string      `json:"token"`
	RoleName    string      `json:"roleName"`
	AccessUrl   string      `json:"accessUrl"`
	Permissions []Component `json:"permissions"`
}

type UserSuccess struct {
	TraceId string        `json:"traceId"`
	Type    string        `json:"type"`
	Title   string        `json:"title"`
	Status  string        `json:"status"`
	Detail  DetailSuccess `json:"detail"`
}
type UserChangePwd struct {
	TraceId string `json:"traceId"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  string `json:"status"`
	Detail  string `json:"detail"`
}
type UserError struct {
	TraceId string        `json:"traceId"`
	Type    string        `json:"type"`
	Title   string        `json:"title"`
	Status  string        `json:"status"`
	Detail  []DetailError `json:"detail"`
}

type DetailSuccess struct {
	Token              string   `json:"token"`
	PersonaTc          string   `json:"personaTc"`
	PersonaReferenceId string   `json:"personaReferenceId"`
	TenantCode         string   `json:"tenantCode"`
	StatusTc           string   `json:"statusTc"`
	IsPasswordExpired  bool     `json:"isPasswordExpired"`
	Role               RoleInfo `json:"role"`
}

type DetailError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

type ChangePassword struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type ForgotPassword struct {
	Username string `json:"username"`
}
type Entitlement struct {
	Entitlement string    `json:"entitlement"`
	Component   Component `json:"component"`
}
type Component struct {
	Id         string       `json:"id"`
	Type       string       `json:"type"`
	Visibility bool         `json:"visibility"`
	Children   []UiChildren `json:"children"`
}

type UiChildren struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Visibility bool   `json:"visibility"`
}
