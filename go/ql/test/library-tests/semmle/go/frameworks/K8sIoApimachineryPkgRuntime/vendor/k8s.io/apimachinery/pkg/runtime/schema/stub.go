// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for k8s.io/apimachinery/pkg/runtime/schema, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: k8s.io/apimachinery/pkg/runtime/schema (exports: GroupVersionKind,ObjectKind; functions: )

// Package schema is a stub of k8s.io/apimachinery/pkg/runtime/schema, generated by depstubber.
package schema

import ()

type GroupKind struct {
	Group string
	Kind  string
}

func (_ GroupKind) Empty() bool {
	return false
}

func (_ GroupKind) String() string {
	return ""
}

func (_ GroupKind) WithVersion(_ string) GroupVersionKind {
	return GroupVersionKind{}
}

type GroupResource struct {
	Group    string
	Resource string
}

func (_ GroupResource) Empty() bool {
	return false
}

func (_ GroupResource) String() string {
	return ""
}

func (_ GroupResource) WithVersion(_ string) GroupVersionResource {
	return GroupVersionResource{}
}

type GroupVersion struct {
	Group   string
	Version string
}

func (_ GroupVersion) Empty() bool {
	return false
}

func (_ GroupVersion) Identifier() string {
	return ""
}

func (_ GroupVersion) KindForGroupVersionKinds(_ []GroupVersionKind) (GroupVersionKind, bool) {
	return GroupVersionKind{}, false
}

func (_ GroupVersion) String() string {
	return ""
}

func (_ GroupVersion) WithKind(_ string) GroupVersionKind {
	return GroupVersionKind{}
}

func (_ GroupVersion) WithResource(_ string) GroupVersionResource {
	return GroupVersionResource{}
}

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

func (_ GroupVersionKind) Empty() bool {
	return false
}

func (_ GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{}
}

func (_ GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{}
}

func (_ GroupVersionKind) String() string {
	return ""
}

func (_ GroupVersionKind) ToAPIVersionAndKind() (string, string) {
	return "", ""
}

type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}

func (_ GroupVersionResource) Empty() bool {
	return false
}

func (_ GroupVersionResource) GroupResource() GroupResource {
	return GroupResource{}
}

func (_ GroupVersionResource) GroupVersion() GroupVersion {
	return GroupVersion{}
}

func (_ GroupVersionResource) String() string {
	return ""
}

type ObjectKind interface {
	GroupVersionKind() GroupVersionKind
	SetGroupVersionKind(_ GroupVersionKind)
}