// --------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// --------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: protobom/v1/types.proto

package sbom

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HashAlgorithm represents the hashing algorithms used within the Software Bill of Materials (SBOM) document.
// It enumerates various hash algorithms that can be employed to generate checksums or unique identifiers for files or data.
type HashAlgorithm int32

const (
	// Unknown hash algorithm.
	HashAlgorithm_UNKNOWN HashAlgorithm = 0
	// MD5 hash algorithm.
	HashAlgorithm_MD5 HashAlgorithm = 1
	// SHA-1 hash algorithm.
	HashAlgorithm_SHA1 HashAlgorithm = 2
	// SHA-256 hash algorithm.
	HashAlgorithm_SHA256 HashAlgorithm = 3
	// SHA-384 hash algorithm.
	HashAlgorithm_SHA384 HashAlgorithm = 4
	// SHA-512 hash algorithm.
	HashAlgorithm_SHA512 HashAlgorithm = 5
	// SHA3-256 hash algorithm.
	HashAlgorithm_SHA3_256 HashAlgorithm = 6
	// SHA3-384 hash algorithm.
	HashAlgorithm_SHA3_384 HashAlgorithm = 7
	// SHA3-512 hash algorithm.
	HashAlgorithm_SHA3_512 HashAlgorithm = 8
	// BLAKE2B-256 hash algorithm.
	HashAlgorithm_BLAKE2B_256 HashAlgorithm = 9
	// BLAKE2B-384 hash algorithm.
	HashAlgorithm_BLAKE2B_384 HashAlgorithm = 10
	// BLAKE2B-512 hash algorithm.
	HashAlgorithm_BLAKE2B_512 HashAlgorithm = 11
	// BLAKE3 hash algorithm.
	HashAlgorithm_BLAKE3 HashAlgorithm = 12
	// MD2 hash algorithm, not supported by SPDX formats.
	HashAlgorithm_MD2 HashAlgorithm = 13
	// Adler-32 hash algorithm, not supported by SPDX formats..
	HashAlgorithm_ADLER32 HashAlgorithm = 14
	// MD4 hash algorithm, not supported by SPDX formats..
	HashAlgorithm_MD4 HashAlgorithm = 15
	// MD6 hash algorithm, not supported by SPDX formats..
	HashAlgorithm_MD6 HashAlgorithm = 16
	// SHA-224 hash algorithm, not supported by SPDX formats..
	HashAlgorithm_SHA224 HashAlgorithm = 17
)

// Enum value maps for HashAlgorithm.
var (
	HashAlgorithm_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "MD5",
		2:  "SHA1",
		3:  "SHA256",
		4:  "SHA384",
		5:  "SHA512",
		6:  "SHA3_256",
		7:  "SHA3_384",
		8:  "SHA3_512",
		9:  "BLAKE2B_256",
		10: "BLAKE2B_384",
		11: "BLAKE2B_512",
		12: "BLAKE3",
		13: "MD2",
		14: "ADLER32",
		15: "MD4",
		16: "MD6",
		17: "SHA224",
	}
	HashAlgorithm_value = map[string]int32{
		"UNKNOWN":     0,
		"MD5":         1,
		"SHA1":        2,
		"SHA256":      3,
		"SHA384":      4,
		"SHA512":      5,
		"SHA3_256":    6,
		"SHA3_384":    7,
		"SHA3_512":    8,
		"BLAKE2B_256": 9,
		"BLAKE2B_384": 10,
		"BLAKE2B_512": 11,
		"BLAKE3":      12,
		"MD2":         13,
		"ADLER32":     14,
		"MD4":         15,
		"MD6":         16,
		"SHA224":      17,
	}
)

func (x HashAlgorithm) Enum() *HashAlgorithm {
	p := new(HashAlgorithm)
	*p = x
	return p
}

func (x HashAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_protobom_v1_types_proto_enumTypes[0].Descriptor()
}

func (HashAlgorithm) Type() protoreflect.EnumType {
	return &file_protobom_v1_types_proto_enumTypes[0]
}

func (x HashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgorithm.Descriptor instead.
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_protobom_v1_types_proto_rawDescGZIP(), []int{0}
}

// Purpose represents different purposes or roles assigned to software entities within the Software Bill of Materials (SBOM).
// It categorizes the roles that software components can fulfill.
type Purpose int32

const (
	// Unknown purpose.
	Purpose_UNKNOWN_PURPOSE Purpose = 0
	// Application purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_APPLICATION Purpose = 1
	// Archive purpose. (SPDX2.3, SPDX3.0)
	Purpose_ARCHIVE Purpose = 2
	// BOM purpose. (SPDX3.0)
	Purpose_BOM Purpose = 3
	// Configuration purpose. (SPDX3.0)
	Purpose_CONFIGURATION Purpose = 4
	// Container purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_CONTAINER Purpose = 5
	// Data purpose. (CDX1.5, SPDX3.0)
	Purpose_DATA Purpose = 6
	// Device purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_DEVICE Purpose = 7
	// Device Driver purpose. (CDX1.5, SPDX3.0)
	Purpose_DEVICE_DRIVER Purpose = 8
	// Documentation purpose. (SPDX3.0)
	Purpose_DOCUMENTATION Purpose = 9
	// Evidence purpose. (SPDX3.0)
	Purpose_EVIDENCE Purpose = 10
	// Executable purpose. (SPDX3.0)
	Purpose_EXECUTABLE Purpose = 11
	// File purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_FILE Purpose = 12
	// Firmware purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_FIRMWARE Purpose = 13
	// Framework purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_FRAMEWORK Purpose = 14
	// Install purpose. (SPDX2.3, SPDX3.0)
	Purpose_INSTALL Purpose = 15
	// Library purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_LIBRARY Purpose = 16
	// Machine Learning Model purpose. (CDX1.5)
	Purpose_MACHINE_LEARNING_MODEL Purpose = 17
	// Manifest purpose. (SPDX3.0)
	Purpose_MANIFEST Purpose = 18
	// Model purpose. (SPDX3.0)
	Purpose_MODEL Purpose = 19
	// Module purpose. (SPDX3.0)
	Purpose_MODULE Purpose = 20
	// Operating System purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_OPERATING_SYSTEM Purpose = 21
	// Other purpose. (SPDX2.3, SPDX3.0)
	Purpose_OTHER Purpose = 22
	// Patch purpose. (SPDX3.0)
	Purpose_PATCH Purpose = 23
	// Platform purpose. (SPDX2.3, CDX1.5, SPDX3.0)
	Purpose_PLATFORM Purpose = 24
	// Requirement purpose. (SPDX3.0)
	Purpose_REQUIREMENT Purpose = 25
	// Source purpose. (SPDX2.3, SPDX3.0)
	Purpose_SOURCE Purpose = 26
	// Specification purpose. (SPDX3.0)
	Purpose_SPECIFICATION Purpose = 27
	// Test purpose. (SPDX3.0)
	Purpose_TEST Purpose = 28
)

// Enum value maps for Purpose.
var (
	Purpose_name = map[int32]string{
		0:  "UNKNOWN_PURPOSE",
		1:  "APPLICATION",
		2:  "ARCHIVE",
		3:  "BOM",
		4:  "CONFIGURATION",
		5:  "CONTAINER",
		6:  "DATA",
		7:  "DEVICE",
		8:  "DEVICE_DRIVER",
		9:  "DOCUMENTATION",
		10: "EVIDENCE",
		11: "EXECUTABLE",
		12: "FILE",
		13: "FIRMWARE",
		14: "FRAMEWORK",
		15: "INSTALL",
		16: "LIBRARY",
		17: "MACHINE_LEARNING_MODEL",
		18: "MANIFEST",
		19: "MODEL",
		20: "MODULE",
		21: "OPERATING_SYSTEM",
		22: "OTHER",
		23: "PATCH",
		24: "PLATFORM",
		25: "REQUIREMENT",
		26: "SOURCE",
		27: "SPECIFICATION",
		28: "TEST",
	}
	Purpose_value = map[string]int32{
		"UNKNOWN_PURPOSE":        0,
		"APPLICATION":            1,
		"ARCHIVE":                2,
		"BOM":                    3,
		"CONFIGURATION":          4,
		"CONTAINER":              5,
		"DATA":                   6,
		"DEVICE":                 7,
		"DEVICE_DRIVER":          8,
		"DOCUMENTATION":          9,
		"EVIDENCE":               10,
		"EXECUTABLE":             11,
		"FILE":                   12,
		"FIRMWARE":               13,
		"FRAMEWORK":              14,
		"INSTALL":                15,
		"LIBRARY":                16,
		"MACHINE_LEARNING_MODEL": 17,
		"MANIFEST":               18,
		"MODEL":                  19,
		"MODULE":                 20,
		"OPERATING_SYSTEM":       21,
		"OTHER":                  22,
		"PATCH":                  23,
		"PLATFORM":               24,
		"REQUIREMENT":            25,
		"SOURCE":                 26,
		"SPECIFICATION":          27,
		"TEST":                   28,
	}
)

func (x Purpose) Enum() *Purpose {
	p := new(Purpose)
	*p = x
	return p
}

func (x Purpose) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Purpose) Descriptor() protoreflect.EnumDescriptor {
	return file_protobom_v1_types_proto_enumTypes[1].Descriptor()
}

func (Purpose) Type() protoreflect.EnumType {
	return &file_protobom_v1_types_proto_enumTypes[1]
}

func (x Purpose) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Purpose.Descriptor instead.
func (Purpose) EnumDescriptor() ([]byte, []int) {
	return file_protobom_v1_types_proto_rawDescGZIP(), []int{1}
}

// SoftwareIdentifierType represents different types of identifiers used
// for software entities within the Software Bill of Materials (SBOM).
type SoftwareIdentifierType int32

const (
	// Unknown software identifier type.
	SoftwareIdentifierType_UNKNOWN_IDENTIFIER_TYPE SoftwareIdentifierType = 0
	// Package URL (PURL) identifier type.
	SoftwareIdentifierType_PURL SoftwareIdentifierType = 1
	// Common Platform Enumeration (CPE) version 2.2 identifier type.
	SoftwareIdentifierType_CPE22 SoftwareIdentifierType = 2
	// Common Platform Enumeration (CPE) version 2.3 identifier type.
	SoftwareIdentifierType_CPE23 SoftwareIdentifierType = 3
	// Git Object Identifier (OID) identifier type.
	SoftwareIdentifierType_GITOID SoftwareIdentifierType = 4
)

// Enum value maps for SoftwareIdentifierType.
var (
	SoftwareIdentifierType_name = map[int32]string{
		0: "UNKNOWN_IDENTIFIER_TYPE",
		1: "PURL",
		2: "CPE22",
		3: "CPE23",
		4: "GITOID",
	}
	SoftwareIdentifierType_value = map[string]int32{
		"UNKNOWN_IDENTIFIER_TYPE": 0,
		"PURL":                    1,
		"CPE22":                   2,
		"CPE23":                   3,
		"GITOID":                  4,
	}
)

func (x SoftwareIdentifierType) Enum() *SoftwareIdentifierType {
	p := new(SoftwareIdentifierType)
	*p = x
	return p
}

func (x SoftwareIdentifierType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SoftwareIdentifierType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobom_v1_types_proto_enumTypes[2].Descriptor()
}

func (SoftwareIdentifierType) Type() protoreflect.EnumType {
	return &file_protobom_v1_types_proto_enumTypes[2]
}

func (x SoftwareIdentifierType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SoftwareIdentifierType.Descriptor instead.
func (SoftwareIdentifierType) EnumDescriptor() ([]byte, []int) {
	return file_protobom_v1_types_proto_rawDescGZIP(), []int{2}
}

var File_protobom_v1_types_proto protoreflect.FileDescriptor

var file_protobom_v1_types_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x2a, 0xf0, 0x01, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44, 0x35, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x48, 0x41, 0x31, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32,
	0x35, 0x36, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x33, 0x38, 0x34, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x48, 0x41, 0x33, 0x5f, 0x32, 0x35, 0x36, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48,
	0x41, 0x33, 0x5f, 0x33, 0x38, 0x34, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x41, 0x33,
	0x5f, 0x35, 0x31, 0x32, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32,
	0x42, 0x5f, 0x32, 0x35, 0x36, 0x10, 0x09, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45,
	0x32, 0x42, 0x5f, 0x33, 0x38, 0x34, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b,
	0x45, 0x32, 0x42, 0x5f, 0x35, 0x31, 0x32, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x4c, 0x41,
	0x4b, 0x45, 0x33, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44, 0x32, 0x10, 0x0d, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x44, 0x4c, 0x45, 0x52, 0x33, 0x32, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x4d,
	0x44, 0x34, 0x10, 0x0f, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x44, 0x36, 0x10, 0x10, 0x12, 0x0a, 0x0a,
	0x06, 0x53, 0x48, 0x41, 0x32, 0x32, 0x34, 0x10, 0x11, 0x2a, 0xb7, 0x03, 0x0a, 0x07, 0x50, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x50, 0x55, 0x52, 0x50, 0x4f, 0x53, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x41,
	0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4f, 0x4d, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45,
	0x52, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x06, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x52, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12,
	0x0c, 0x0a, 0x08, 0x45, 0x56, 0x49, 0x44, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x0e, 0x0a,
	0x0a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0b, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x52, 0x4d, 0x57,
	0x41, 0x52, 0x45, 0x10, 0x0d, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x52, 0x41, 0x4d, 0x45, 0x57, 0x4f,
	0x52, 0x4b, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x10,
	0x0f, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x42, 0x52, 0x41, 0x52, 0x59, 0x10, 0x10, 0x12, 0x1a,
	0x0a, 0x16, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x52, 0x4e, 0x49,
	0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x10, 0x11, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41,
	0x4e, 0x49, 0x46, 0x45, 0x53, 0x54, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x10, 0x13, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x14, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x59, 0x53,
	0x54, 0x45, 0x4d, 0x10, 0x15, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x16,
	0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x17, 0x12, 0x0c, 0x0a, 0x08, 0x50,
	0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x51,
	0x55, 0x49, 0x52, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x19, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x10, 0x1a, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x1b, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x53,
	0x54, 0x10, 0x1c, 0x2a, 0x61, 0x0a, 0x16, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46,
	0x49, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55,
	0x52, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x50, 0x45, 0x32, 0x32, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x50, 0x45, 0x32, 0x33, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49,
	0x54, 0x4f, 0x49, 0x44, 0x10, 0x04, 0x42, 0x91, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x62, 0x6f, 0x6d, 0xa2,
	0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x17, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x6f, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protobom_v1_types_proto_rawDescOnce sync.Once
	file_protobom_v1_types_proto_rawDescData = file_protobom_v1_types_proto_rawDesc
)

func file_protobom_v1_types_proto_rawDescGZIP() []byte {
	file_protobom_v1_types_proto_rawDescOnce.Do(func() {
		file_protobom_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobom_v1_types_proto_rawDescData)
	})
	return file_protobom_v1_types_proto_rawDescData
}

var file_protobom_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protobom_v1_types_proto_goTypes = []any{
	(HashAlgorithm)(0),          // 0: protobom.v1.HashAlgorithm
	(Purpose)(0),                // 1: protobom.v1.Purpose
	(SoftwareIdentifierType)(0), // 2: protobom.v1.SoftwareIdentifierType
}
var file_protobom_v1_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobom_v1_types_proto_init() }
func file_protobom_v1_types_proto_init() {
	if File_protobom_v1_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobom_v1_types_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobom_v1_types_proto_goTypes,
		DependencyIndexes: file_protobom_v1_types_proto_depIdxs,
		EnumInfos:         file_protobom_v1_types_proto_enumTypes,
	}.Build()
	File_protobom_v1_types_proto = out.File
	file_protobom_v1_types_proto_rawDesc = nil
	file_protobom_v1_types_proto_goTypes = nil
	file_protobom_v1_types_proto_depIdxs = nil
}
