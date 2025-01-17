// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package plog

import (
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceLogsSlice logically represents a slice of ResourceLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewResourceLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ResourceLogsSlice internal.ResourceLogsSlice

func newResourceLogsSlice(orig *[]*otlplogs.ResourceLogs) ResourceLogsSlice {
	return ResourceLogsSlice(internal.NewResourceLogsSlice(orig))
}

func (ms ResourceLogsSlice) getOrig() *[]*otlplogs.ResourceLogs {
	return internal.GetOrigResourceLogsSlice(internal.ResourceLogsSlice(ms))
}

// NewResourceLogsSlice creates a ResourceLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewResourceLogsSlice() ResourceLogsSlice {
	orig := []*otlplogs.ResourceLogs(nil)
	return newResourceLogsSlice(&orig)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewResourceLogsSlice()".
func (es ResourceLogsSlice) Len() int {
	return len(*es.getOrig())
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ResourceLogsSlice) At(ix int) ResourceLogs {
	return newResourceLogs((*es.getOrig())[ix])
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ResourceLogsSlice) CopyTo(dest ResourceLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.getOrig())
	if srcLen <= destCap {
		(*dest.getOrig()) = (*dest.getOrig())[:srcLen:destCap]
		for i := range *es.getOrig() {
			newResourceLogs((*es.getOrig())[i]).CopyTo(newResourceLogs((*dest.getOrig())[i]))
		}
		return
	}
	origs := make([]otlplogs.ResourceLogs, srcLen)
	wrappers := make([]*otlplogs.ResourceLogs, srcLen)
	for i := range *es.getOrig() {
		wrappers[i] = &origs[i]
		newResourceLogs((*es.getOrig())[i]).CopyTo(newResourceLogs(wrappers[i]))
	}
	*dest.getOrig() = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ResourceLogsSlice can be initialized:
//
//	es := NewResourceLogsSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ResourceLogsSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.ResourceLogs, len(*es.getOrig()), newCap)
	copy(newOrig, *es.getOrig())
	*es.getOrig() = newOrig
}

// AppendEmpty will append to the end of the slice an empty ResourceLogs.
// It returns the newly added ResourceLogs.
func (es ResourceLogsSlice) AppendEmpty() ResourceLogs {
	*es.getOrig() = append(*es.getOrig(), &otlplogs.ResourceLogs{})
	return es.At(es.Len() - 1)
}

// Sort sorts the ResourceLogs elements within ResourceLogsSlice given the
// provided less function so that two instances of ResourceLogsSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//
//	lessFunc := func(a, b ResourceLogs) bool {
//	  return a.Name() < b.Name() // choose any comparison here
//	}
//	assert.Equal(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es ResourceLogsSlice) Sort(less func(a, b ResourceLogs) bool) ResourceLogsSlice {
	sort.SliceStable(*es.getOrig(), func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ResourceLogsSlice) MoveAndAppendTo(dest ResourceLogsSlice) {
	if *dest.getOrig() == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.getOrig() = *es.getOrig()
	} else {
		*dest.getOrig() = append(*dest.getOrig(), *es.getOrig()...)
	}
	*es.getOrig() = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ResourceLogsSlice) RemoveIf(f func(ResourceLogs) bool) {
	newLen := 0
	for i := 0; i < len(*es.getOrig()); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.getOrig())[newLen] = (*es.getOrig())[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.getOrig() = (*es.getOrig())[:newLen]
}

// ResourceLogs is a collection of logs from a Resource.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewResourceLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.

type ResourceLogs internal.ResourceLogs

func newResourceLogs(orig *otlplogs.ResourceLogs) ResourceLogs {
	return ResourceLogs(internal.NewResourceLogs(orig))
}

func (ms ResourceLogs) getOrig() *otlplogs.ResourceLogs {
	return internal.GetOrigResourceLogs(internal.ResourceLogs(ms))
}

// NewResourceLogs creates a new empty ResourceLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewResourceLogs() ResourceLogs {
	return newResourceLogs(&otlplogs.ResourceLogs{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ResourceLogs) MoveTo(dest ResourceLogs) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlplogs.ResourceLogs{}
}

// Resource returns the resource associated with this ResourceLogs.
func (ms ResourceLogs) Resource() pcommon.Resource {
	return pcommon.Resource(internal.NewResource(&ms.getOrig().Resource))
}

// SchemaUrl returns the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SchemaUrl() string {
	return ms.getOrig().SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ResourceLogs.
func (ms ResourceLogs) SetSchemaUrl(v string) {
	ms.getOrig().SchemaUrl = v
}

// ScopeLogs returns the ScopeLogs associated with this ResourceLogs.
func (ms ResourceLogs) ScopeLogs() ScopeLogsSlice {
	return ScopeLogsSlice(internal.NewScopeLogsSlice(&ms.getOrig().ScopeLogs))
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ResourceLogs) CopyTo(dest ResourceLogs) {
	ms.Resource().CopyTo(dest.Resource())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.ScopeLogs().CopyTo(dest.ScopeLogs())
}

// ScopeLogsSlice logically represents a slice of ScopeLogs.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewScopeLogsSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ScopeLogsSlice internal.ScopeLogsSlice

func newScopeLogsSlice(orig *[]*otlplogs.ScopeLogs) ScopeLogsSlice {
	return ScopeLogsSlice(internal.NewScopeLogsSlice(orig))
}

func (ms ScopeLogsSlice) getOrig() *[]*otlplogs.ScopeLogs {
	return internal.GetOrigScopeLogsSlice(internal.ScopeLogsSlice(ms))
}

// NewScopeLogsSlice creates a ScopeLogsSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewScopeLogsSlice() ScopeLogsSlice {
	orig := []*otlplogs.ScopeLogs(nil)
	return newScopeLogsSlice(&orig)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewScopeLogsSlice()".
func (es ScopeLogsSlice) Len() int {
	return len(*es.getOrig())
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es ScopeLogsSlice) At(ix int) ScopeLogs {
	return newScopeLogs((*es.getOrig())[ix])
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es ScopeLogsSlice) CopyTo(dest ScopeLogsSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.getOrig())
	if srcLen <= destCap {
		(*dest.getOrig()) = (*dest.getOrig())[:srcLen:destCap]
		for i := range *es.getOrig() {
			newScopeLogs((*es.getOrig())[i]).CopyTo(newScopeLogs((*dest.getOrig())[i]))
		}
		return
	}
	origs := make([]otlplogs.ScopeLogs, srcLen)
	wrappers := make([]*otlplogs.ScopeLogs, srcLen)
	for i := range *es.getOrig() {
		wrappers[i] = &origs[i]
		newScopeLogs((*es.getOrig())[i]).CopyTo(newScopeLogs(wrappers[i]))
	}
	*dest.getOrig() = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new ScopeLogsSlice can be initialized:
//
//	es := NewScopeLogsSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es ScopeLogsSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.ScopeLogs, len(*es.getOrig()), newCap)
	copy(newOrig, *es.getOrig())
	*es.getOrig() = newOrig
}

// AppendEmpty will append to the end of the slice an empty ScopeLogs.
// It returns the newly added ScopeLogs.
func (es ScopeLogsSlice) AppendEmpty() ScopeLogs {
	*es.getOrig() = append(*es.getOrig(), &otlplogs.ScopeLogs{})
	return es.At(es.Len() - 1)
}

// Sort sorts the ScopeLogs elements within ScopeLogsSlice given the
// provided less function so that two instances of ScopeLogsSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//
//	lessFunc := func(a, b ScopeLogs) bool {
//	  return a.Name() < b.Name() // choose any comparison here
//	}
//	assert.Equal(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es ScopeLogsSlice) Sort(less func(a, b ScopeLogs) bool) ScopeLogsSlice {
	sort.SliceStable(*es.getOrig(), func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es ScopeLogsSlice) MoveAndAppendTo(dest ScopeLogsSlice) {
	if *dest.getOrig() == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.getOrig() = *es.getOrig()
	} else {
		*dest.getOrig() = append(*dest.getOrig(), *es.getOrig()...)
	}
	*es.getOrig() = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es ScopeLogsSlice) RemoveIf(f func(ScopeLogs) bool) {
	newLen := 0
	for i := 0; i < len(*es.getOrig()); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.getOrig())[newLen] = (*es.getOrig())[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.getOrig() = (*es.getOrig())[:newLen]
}

// ScopeLogs is a collection of logs from a LibraryInstrumentation.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewScopeLogs function to create new instances.
// Important: zero-initialized instance is not valid for use.

type ScopeLogs internal.ScopeLogs

func newScopeLogs(orig *otlplogs.ScopeLogs) ScopeLogs {
	return ScopeLogs(internal.NewScopeLogs(orig))
}

func (ms ScopeLogs) getOrig() *otlplogs.ScopeLogs {
	return internal.GetOrigScopeLogs(internal.ScopeLogs(ms))
}

// NewScopeLogs creates a new empty ScopeLogs.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewScopeLogs() ScopeLogs {
	return newScopeLogs(&otlplogs.ScopeLogs{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ScopeLogs) MoveTo(dest ScopeLogs) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlplogs.ScopeLogs{}
}

// Scope returns the scope associated with this ScopeLogs.
func (ms ScopeLogs) Scope() pcommon.InstrumentationScope {
	return pcommon.InstrumentationScope(internal.NewInstrumentationScope(&ms.getOrig().Scope))
}

// SchemaUrl returns the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SchemaUrl() string {
	return ms.getOrig().SchemaUrl
}

// SetSchemaUrl replaces the schemaurl associated with this ScopeLogs.
func (ms ScopeLogs) SetSchemaUrl(v string) {
	ms.getOrig().SchemaUrl = v
}

// LogRecords returns the LogRecords associated with this ScopeLogs.
func (ms ScopeLogs) LogRecords() LogRecordSlice {
	return LogRecordSlice(internal.NewLogRecordSlice(&ms.getOrig().LogRecords))
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ScopeLogs) CopyTo(dest ScopeLogs) {
	ms.Scope().CopyTo(dest.Scope())
	dest.SetSchemaUrl(ms.SchemaUrl())
	ms.LogRecords().CopyTo(dest.LogRecords())
}

// LogRecordSlice logically represents a slice of LogRecord.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewLogRecordSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type LogRecordSlice internal.LogRecordSlice

func newLogRecordSlice(orig *[]*otlplogs.LogRecord) LogRecordSlice {
	return LogRecordSlice(internal.NewLogRecordSlice(orig))
}

func (ms LogRecordSlice) getOrig() *[]*otlplogs.LogRecord {
	return internal.GetOrigLogRecordSlice(internal.LogRecordSlice(ms))
}

// NewLogRecordSlice creates a LogRecordSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewLogRecordSlice() LogRecordSlice {
	orig := []*otlplogs.LogRecord(nil)
	return newLogRecordSlice(&orig)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewLogRecordSlice()".
func (es LogRecordSlice) Len() int {
	return len(*es.getOrig())
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es LogRecordSlice) At(ix int) LogRecord {
	return newLogRecord((*es.getOrig())[ix])
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es LogRecordSlice) CopyTo(dest LogRecordSlice) {
	srcLen := es.Len()
	destCap := cap(*dest.getOrig())
	if srcLen <= destCap {
		(*dest.getOrig()) = (*dest.getOrig())[:srcLen:destCap]
		for i := range *es.getOrig() {
			newLogRecord((*es.getOrig())[i]).CopyTo(newLogRecord((*dest.getOrig())[i]))
		}
		return
	}
	origs := make([]otlplogs.LogRecord, srcLen)
	wrappers := make([]*otlplogs.LogRecord, srcLen)
	for i := range *es.getOrig() {
		wrappers[i] = &origs[i]
		newLogRecord((*es.getOrig())[i]).CopyTo(newLogRecord(wrappers[i]))
	}
	*dest.getOrig() = wrappers
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new LogRecordSlice can be initialized:
//
//	es := NewLogRecordSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es LogRecordSlice) EnsureCapacity(newCap int) {
	oldCap := cap(*es.getOrig())
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlplogs.LogRecord, len(*es.getOrig()), newCap)
	copy(newOrig, *es.getOrig())
	*es.getOrig() = newOrig
}

// AppendEmpty will append to the end of the slice an empty LogRecord.
// It returns the newly added LogRecord.
func (es LogRecordSlice) AppendEmpty() LogRecord {
	*es.getOrig() = append(*es.getOrig(), &otlplogs.LogRecord{})
	return es.At(es.Len() - 1)
}

// Sort sorts the LogRecord elements within LogRecordSlice given the
// provided less function so that two instances of LogRecordSlice
// can be compared.
//
// Returns the same instance to allow nicer code like:
//
//	lessFunc := func(a, b LogRecord) bool {
//	  return a.Name() < b.Name() // choose any comparison here
//	}
//	assert.Equal(t, expected.Sort(lessFunc), actual.Sort(lessFunc))
func (es LogRecordSlice) Sort(less func(a, b LogRecord) bool) LogRecordSlice {
	sort.SliceStable(*es.getOrig(), func(i, j int) bool { return less(es.At(i), es.At(j)) })
	return es
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es LogRecordSlice) MoveAndAppendTo(dest LogRecordSlice) {
	if *dest.getOrig() == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.getOrig() = *es.getOrig()
	} else {
		*dest.getOrig() = append(*dest.getOrig(), *es.getOrig()...)
	}
	*es.getOrig() = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es LogRecordSlice) RemoveIf(f func(LogRecord) bool) {
	newLen := 0
	for i := 0; i < len(*es.getOrig()); i++ {
		if f(es.At(i)) {
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.getOrig())[newLen] = (*es.getOrig())[i]
		newLen++
	}
	// TODO: Prevent memory leak by erasing truncated values.
	*es.getOrig() = (*es.getOrig())[:newLen]
}

// LogRecord are experimental implementation of OpenTelemetry Log Data Model.

//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewLogRecord function to create new instances.
// Important: zero-initialized instance is not valid for use.

type LogRecord internal.LogRecord

func newLogRecord(orig *otlplogs.LogRecord) LogRecord {
	return LogRecord(internal.NewLogRecord(orig))
}

func (ms LogRecord) getOrig() *otlplogs.LogRecord {
	return internal.GetOrigLogRecord(internal.LogRecord(ms))
}

// NewLogRecord creates a new empty LogRecord.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewLogRecord() LogRecord {
	return newLogRecord(&otlplogs.LogRecord{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms LogRecord) MoveTo(dest LogRecord) {
	*dest.getOrig() = *ms.getOrig()
	*ms.getOrig() = otlplogs.LogRecord{}
}

// ObservedTimestamp returns the observedtimestamp associated with this LogRecord.
func (ms LogRecord) ObservedTimestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.getOrig().ObservedTimeUnixNano)
}

// SetObservedTimestamp replaces the observedtimestamp associated with this LogRecord.
func (ms LogRecord) SetObservedTimestamp(v pcommon.Timestamp) {
	ms.getOrig().ObservedTimeUnixNano = uint64(v)
}

// Timestamp returns the timestamp associated with this LogRecord.
func (ms LogRecord) Timestamp() pcommon.Timestamp {
	return pcommon.Timestamp(ms.getOrig().TimeUnixNano)
}

// SetTimestamp replaces the timestamp associated with this LogRecord.
func (ms LogRecord) SetTimestamp(v pcommon.Timestamp) {
	ms.getOrig().TimeUnixNano = uint64(v)
}

// TraceID returns the traceid associated with this LogRecord.
func (ms LogRecord) TraceID() pcommon.TraceID {
	return pcommon.TraceID(ms.getOrig().TraceId)
}

// SetTraceID replaces the traceid associated with this LogRecord.
func (ms LogRecord) SetTraceID(v pcommon.TraceID) {
	ms.getOrig().TraceId = data.TraceID(v)
}

// SpanID returns the spanid associated with this LogRecord.
func (ms LogRecord) SpanID() pcommon.SpanID {
	return pcommon.SpanID(ms.getOrig().SpanId)
}

// SetSpanID replaces the spanid associated with this LogRecord.
func (ms LogRecord) SetSpanID(v pcommon.SpanID) {
	ms.getOrig().SpanId = data.SpanID(v)
}

// Flags returns the flags associated with this LogRecord.
func (ms LogRecord) Flags() LogRecordFlags {
	return LogRecordFlags(ms.getOrig().Flags)
}

// SetFlags replaces the flags associated with this LogRecord.
func (ms LogRecord) SetFlags(v LogRecordFlags) {
	ms.getOrig().Flags = uint32(v)
}

// SeverityText returns the severitytext associated with this LogRecord.
func (ms LogRecord) SeverityText() string {
	return ms.getOrig().SeverityText
}

// SetSeverityText replaces the severitytext associated with this LogRecord.
func (ms LogRecord) SetSeverityText(v string) {
	ms.getOrig().SeverityText = v
}

// SeverityNumber returns the severitynumber associated with this LogRecord.
func (ms LogRecord) SeverityNumber() SeverityNumber {
	return SeverityNumber(ms.getOrig().SeverityNumber)
}

// SetSeverityNumber replaces the severitynumber associated with this LogRecord.
func (ms LogRecord) SetSeverityNumber(v SeverityNumber) {
	ms.getOrig().SeverityNumber = otlplogs.SeverityNumber(v)
}

// Body returns the body associated with this LogRecord.
func (ms LogRecord) Body() pcommon.Value {
	return pcommon.Value(internal.NewValue(&ms.getOrig().Body))
}

// Attributes returns the Attributes associated with this LogRecord.
func (ms LogRecord) Attributes() pcommon.Map {
	return pcommon.Map(internal.NewMap(&ms.getOrig().Attributes))
}

// DroppedAttributesCount returns the droppedattributescount associated with this LogRecord.
func (ms LogRecord) DroppedAttributesCount() uint32 {
	return ms.getOrig().DroppedAttributesCount
}

// SetDroppedAttributesCount replaces the droppedattributescount associated with this LogRecord.
func (ms LogRecord) SetDroppedAttributesCount(v uint32) {
	ms.getOrig().DroppedAttributesCount = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms LogRecord) CopyTo(dest LogRecord) {
	dest.SetObservedTimestamp(ms.ObservedTimestamp())
	dest.SetTimestamp(ms.Timestamp())
	dest.SetTraceID(ms.TraceID())
	dest.SetSpanID(ms.SpanID())
	dest.SetFlags(ms.Flags())
	dest.SetSeverityText(ms.SeverityText())
	dest.SetSeverityNumber(ms.SeverityNumber())
	ms.Body().CopyTo(dest.Body())
	ms.Attributes().CopyTo(dest.Attributes())
	dest.SetDroppedAttributesCount(ms.DroppedAttributesCount())
}
