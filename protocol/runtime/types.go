// Code generated by cdpgen. DO NOT EDIT.

package runtime

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ScriptID Unique script identifier.
type ScriptID string

// RemoteObjectID Unique object identifier.
type RemoteObjectID string

// UnserializableValue Primitive value which cannot be JSON-stringified.
type UnserializableValue int

// UnserializableValue as enums.
const (
	UnserializableValueNotSet UnserializableValue = iota
	UnserializableValueInfinity
	UnserializableValueNaN
	UnserializableValueNegativeInfinity
	UnserializableValueNegative0
)

// Valid returns true if enum is set.
func (e UnserializableValue) Valid() bool {
	return e >= 1 && e <= 4
}

func (e UnserializableValue) String() string {
	switch e {
	case 0:
		return "UnserializableValueNotSet"
	case 1:
		return "Infinity"
	case 2:
		return "NaN"
	case 3:
		return "-Infinity"
	case 4:
		return "-0"
	}
	return fmt.Sprintf("UnserializableValue(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e UnserializableValue) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("runtime.UnserializableValue: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *UnserializableValue) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"Infinity\"":
		*e = 1
	case "\"NaN\"":
		*e = 2
	case "\"-Infinity\"":
		*e = 3
	case "\"-0\"":
		*e = 4
	default:
		return fmt.Errorf("runtime.UnserializableValue: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}

// RemoteObject Mirror object referencing original JavaScript object.
type RemoteObject struct {
	// Type Object type.
	//
	// Values: "object", "function", "undefined", "string", "number", "boolean", "symbol".
	Type string `json:"type"`
	// Subtype Object subtype hint. Specified for object type values only.
	//
	// Values: "array", "null", "node", "regexp", "date", "map", "set", "weakmap", "weakset", "iterator", "generator", "error", "proxy", "promise", "typedarray".
	Subtype             *string             `json:"subtype,omitempty"`
	ClassName           *string             `json:"className,omitempty"`           // Object class (constructor) name. Specified for object type values only.
	Value               json.RawMessage     `json:"value,omitempty"`               // Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified does not have value, but gets this property.
	Description         *string             `json:"description,omitempty"`         // String representation of the object.
	ObjectID            *RemoteObjectID     `json:"objectId,omitempty"`            // Unique object identifier (for non-primitive values).
	// Preview Preview containing abbreviated property values. Specified for object type values only.
	//
	// Note: This property is experimental.
	Preview *ObjectPreview `json:"preview,omitempty"`
	// CustomPreview
	//
	// Note: This property is experimental.
	CustomPreview *CustomPreview `json:"customPreview,omitempty"`
}

// CustomPreview
//
// Note: This type is experimental.
type CustomPreview struct {
	Header                     string          `json:"header"`                     // No description.
	HasBody                    bool            `json:"hasBody"`                    // No description.
	FormatterObjectID          RemoteObjectID  `json:"formatterObjectId"`          // No description.
	BindRemoteObjectFunctionID RemoteObjectID  `json:"bindRemoteObjectFunctionId"` // No description.
	ConfigObjectID             *RemoteObjectID `json:"configObjectId,omitempty"`   // No description.
}

// ObjectPreview Object containing abbreviated remote object value.
//
// Note: This type is experimental.
type ObjectPreview struct {
	// Type Object type.
	//
	// Values: "object", "function", "undefined", "string", "number", "boolean", "symbol".
	Type string `json:"type"`
	// Subtype Object subtype hint. Specified for object type values only.
	//
	// Values: "array", "null", "node", "regexp", "date", "map", "set", "weakmap", "weakset", "iterator", "generator", "error".
	Subtype     *string           `json:"subtype,omitempty"`
	Description *string           `json:"description,omitempty"` // String representation of the object.
	Overflow    bool              `json:"overflow"`              // True iff some of the properties or entries of the original object did not fit.
	Properties  []PropertyPreview `json:"properties"`            // List of the properties.
	Entries     []EntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for map and set subtype values only.
}

// PropertyPreview
//
// Note: This type is experimental.
type PropertyPreview struct {
	Name string `json:"name"` // Property name.
	// Type Object type. Accessor means that the property itself is an accessor property.
	//
	// Values: "object", "function", "undefined", "string", "number", "boolean", "symbol", "accessor".
	Type         string         `json:"type"`
	Value        *string        `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *ObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	// Subtype Object subtype hint. Specified for object type values only.
	//
	// Values: "array", "null", "node", "regexp", "date", "map", "set", "weakmap", "weakset", "iterator", "generator", "error".
	Subtype *string `json:"subtype,omitempty"`
}

// EntryPreview
//
// Note: This type is experimental.
type EntryPreview struct {
	Key   *ObjectPreview `json:"key,omitempty"` // Preview of the key. Specified for map-like collection entries.
	Value ObjectPreview  `json:"value"`         // Preview of the value.
}

// PropertyDescriptor Object property descriptor.
type PropertyDescriptor struct {
	Name         string        `json:"name"`                // Property name or symbol description.
	Value        *RemoteObject `json:"value,omitempty"`     // The value associated with the property.
	Writable     *bool         `json:"writable,omitempty"`  // True if the value associated with the property may be changed (data descriptors only).
	Get          *RemoteObject `json:"get,omitempty"`       // A function which serves as a getter for the property, or undefined if there is no getter (accessor descriptors only).
	Set          *RemoteObject `json:"set,omitempty"`       // A function which serves as a setter for the property, or undefined if there is no setter (accessor descriptors only).
	Configurable bool          `json:"configurable"`        // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool          `json:"enumerable"`          // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    *bool         `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn        *bool         `json:"isOwn,omitempty"`     // True if the property is owned for the object.
	Symbol       *RemoteObject `json:"symbol,omitempty"`    // Property symbol object, if the property is of the symbol type.
}

// InternalPropertyDescriptor Object internal property descriptor. This property isn't normally visible in JavaScript code.
type InternalPropertyDescriptor struct {
	Name  string        `json:"name"`            // Conventional property name.
	Value *RemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// CallArgument Represents function call argument. Either remote object id objectId, primitive value, unserializable primitive value or neither of (for undefined) them should be specified.
type CallArgument struct {
	Value               json.RawMessage     `json:"value,omitempty"`               // Primitive value.
	UnserializableValue UnserializableValue `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified.
	ObjectID            *RemoteObjectID     `json:"objectId,omitempty"`            // Remote object handle.
}

// ExecutionContextID Id of an execution context.
type ExecutionContextID int

// ExecutionContextDescription Description of an isolated world.
type ExecutionContextDescription struct {
	ID      ExecutionContextID `json:"id"`                // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin  string             `json:"origin"`            // Execution context origin.
	Name    string             `json:"name"`              // Human readable name describing given context.
	AuxData json.RawMessage    `json:"auxData,omitempty"` // Embedder-specific auxiliary data.
}

// ExceptionDetails Detailed information about exception (or error) that was thrown during script compilation or execution.
type ExceptionDetails struct {
	ExceptionID        int                 `json:"exceptionId"`                  // Exception id.
	Text               string              `json:"text"`                         // Exception text, which should be used together with exception object when available.
	LineNumber         int                 `json:"lineNumber"`                   // Line number of the exception location (0-based).
	ColumnNumber       int                 `json:"columnNumber"`                 // Column number of the exception location (0-based).
	ScriptID           *ScriptID           `json:"scriptId,omitempty"`           // Script ID of the exception location.
	URL                *string             `json:"url,omitempty"`                // URL of the exception location, to be used when the script was not reported.
	StackTrace         *StackTrace         `json:"stackTrace,omitempty"`         // JavaScript stack trace if available.
	Exception          *RemoteObject       `json:"exception,omitempty"`          // Exception object if available.
	ExecutionContextID *ExecutionContextID `json:"executionContextId,omitempty"` // Identifier of the context where exception happened.
}

// Timestamp Number of milliseconds since epoch.
type Timestamp float64

// String calls (time.Time).String().
func (t Timestamp) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t Timestamp) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("runtime.Timestamp: " + err.Error())
	}
	*t = Timestamp(f)
	return nil
}

var _ json.Marshaler = (*Timestamp)(nil)
var _ json.Unmarshaler = (*Timestamp)(nil)

// CallFrame Stack entry for runtime errors and assertions.
type CallFrame struct {
	FunctionName string   `json:"functionName"` // JavaScript function name.
	ScriptID     ScriptID `json:"scriptId"`     // JavaScript script id.
	URL          string   `json:"url"`          // JavaScript script name or url.
	LineNumber   int      `json:"lineNumber"`   // JavaScript script line number (0-based).
	ColumnNumber int      `json:"columnNumber"` // JavaScript script column number (0-based).
}

// StackTrace Call frames for assertions or error messages.
type StackTrace struct {
	Description *string     `json:"description,omitempty"` // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames  []CallFrame `json:"callFrames"`            // JavaScript function name.
	Parent      *StackTrace `json:"parent,omitempty"`      // Asynchronous JavaScript stack trace that preceded this stack, if available.
	// PromiseCreationFrame Creation frame of the Promise which produced the next synchronous trace when resolved, if available.
	//
	// Note: This property is experimental.
	PromiseCreationFrame *CallFrame `json:"promiseCreationFrame,omitempty"`
}
