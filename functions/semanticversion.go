package main

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

//This is a func
func SemanticVersionString(sv SemanticVersion) string {
	return fmt.Sprintf("%d.%d.%d", sv.minor, sv.minor, sv.patch)
}

//This is a method receiver on a struct
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.minor, sv.minor, sv.patch)
}

func (sv *SemanticVersion) IncrementMajor() {
	sv.major += 1
}

func (sv *SemanticVersion) IncrementMinor() {
	sv.minor += 1
}

func (sv *SemanticVersion) IncrementPatch() {
	sv.patch += 1
}
