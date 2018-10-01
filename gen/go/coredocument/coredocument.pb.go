// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coredocument/coredocument.proto

package coredocumentpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/centrifuge/precise-proofs/proofs/proto"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// `CoreDocument` is a document that can be sent to different nodes and anchored
// on chain. It handles all the generic features native Centrifuge Documents support:
//
// * Merkle Roots for the document data
// * Signatures on document data
// * Access Control
type CoreDocument struct {
	// # Identifiers
	// CoreDocument has two kinds of identifiers, the `document_identifier` is assigned
	// once per document and stays the same for the lifetime of the document.
	// document_identifier is the first ID ever used to anchor the document on chain and
	// is used internally to store all future versions. The `previous_identifier`, `current_identifier`, and the
	// `next_identifier` refer only to a particular version.
	//
	// 32 byte value
	DocumentIdentifier []byte `protobuf:"bytes,9,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	// previous_identifier refers to the previous state of the document.
	// 32 byte value
	PreviousIdentifier []byte `protobuf:"bytes,16,opt,name=previous_identifier,json=previousIdentifier,proto3" json:"previous_identifier,omitempty"`
	// current_identifier is the identifier used to refer to the current state of the document.
	// 32 byte value
	CurrentIdentifier []byte `protobuf:"bytes,3,opt,name=current_identifier,json=currentIdentifier,proto3" json:"current_identifier,omitempty"`
	// next_identifier is the identifier that is going to be used for the next version if any
	// party wants to update the state.
	NextIdentifier []byte `protobuf:"bytes,4,opt,name=next_identifier,json=nextIdentifier,proto3" json:"next_identifier,omitempty"`
	// # Merkle roots
	// document_root the root of all roots. It's signature_root along with a list of all signatures
	DocumentRoot []byte `protobuf:"bytes,7,opt,name=document_root,json=documentRoot,proto3" json:"document_root,omitempty"`
	// signing_root is the Merkle root calculated from all fields on the document that need
	// to be signed. This is all fields except for the signatures themselves and the `document_root`.
	SigningRoot []byte `protobuf:"bytes,10,opt,name=signing_root,json=signingRoot,proto3" json:"signing_root,omitempty"`
	// previous_root is the `document_root` of the previous version of the document
	PreviousRoot []byte `protobuf:"bytes,2,opt,name=previous_root,json=previousRoot,proto3" json:"previous_root,omitempty"`
	// data_root is the Merkle root calculated for the `embedded_data` & `embedded_salts` document (such as an invoice).
	DataRoot []byte `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	// Signatures
	// Signatures of the signature_root by centrifuge identities. This array should be sorted by the Centrifuge ID.
	Signatures []*Signature `protobuf:"bytes,6,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// Document a serialized document
	EmbeddedData      *any.Any `protobuf:"bytes,13,opt,name=embedded_data,json=embeddedData,proto3" json:"embedded_data,omitempty"`
	EmbeddedDataSalts *any.Any `protobuf:"bytes,14,opt,name=embedded_data_salts,json=embeddedDataSalts,proto3" json:"embedded_data_salts,omitempty"`
	// CoreDocumentSalts is inlined
	CoredocumentSalts    *CoreDocumentSalts `protobuf:"bytes,15,opt,name=coredocument_salts,json=coredocumentSalts,proto3" json:"coredocument_salts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CoreDocument) Reset()         { *m = CoreDocument{} }
func (m *CoreDocument) String() string { return proto.CompactTextString(m) }
func (*CoreDocument) ProtoMessage()    {}
func (*CoreDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_3ca8eeb62ad40595, []int{0}
}
func (m *CoreDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreDocument.Unmarshal(m, b)
}
func (m *CoreDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreDocument.Marshal(b, m, deterministic)
}
func (dst *CoreDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreDocument.Merge(dst, src)
}
func (m *CoreDocument) XXX_Size() int {
	return xxx_messageInfo_CoreDocument.Size(m)
}
func (m *CoreDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreDocument.DiscardUnknown(m)
}

var xxx_messageInfo_CoreDocument proto.InternalMessageInfo

func (m *CoreDocument) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CoreDocument) GetPreviousIdentifier() []byte {
	if m != nil {
		return m.PreviousIdentifier
	}
	return nil
}

func (m *CoreDocument) GetCurrentIdentifier() []byte {
	if m != nil {
		return m.CurrentIdentifier
	}
	return nil
}

func (m *CoreDocument) GetNextIdentifier() []byte {
	if m != nil {
		return m.NextIdentifier
	}
	return nil
}

func (m *CoreDocument) GetDocumentRoot() []byte {
	if m != nil {
		return m.DocumentRoot
	}
	return nil
}

func (m *CoreDocument) GetSigningRoot() []byte {
	if m != nil {
		return m.SigningRoot
	}
	return nil
}

func (m *CoreDocument) GetPreviousRoot() []byte {
	if m != nil {
		return m.PreviousRoot
	}
	return nil
}

func (m *CoreDocument) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

func (m *CoreDocument) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *CoreDocument) GetEmbeddedData() *any.Any {
	if m != nil {
		return m.EmbeddedData
	}
	return nil
}

func (m *CoreDocument) GetEmbeddedDataSalts() *any.Any {
	if m != nil {
		return m.EmbeddedDataSalts
	}
	return nil
}

func (m *CoreDocument) GetCoredocumentSalts() *CoreDocumentSalts {
	if m != nil {
		return m.CoredocumentSalts
	}
	return nil
}

type CoreDocumentSalts struct {
	DocumentIdentifier   []byte   `protobuf:"bytes,9,opt,name=document_identifier,json=documentIdentifier,proto3" json:"document_identifier,omitempty"`
	PreviousIdentifier   []byte   `protobuf:"bytes,1,opt,name=previous_identifier,json=previousIdentifier,proto3" json:"previous_identifier,omitempty"`
	CurrentIdentifier    []byte   `protobuf:"bytes,3,opt,name=current_identifier,json=currentIdentifier,proto3" json:"current_identifier,omitempty"`
	NextIdentifier       []byte   `protobuf:"bytes,4,opt,name=next_identifier,json=nextIdentifier,proto3" json:"next_identifier,omitempty"`
	PreviousRoot         []byte   `protobuf:"bytes,2,opt,name=previous_root,json=previousRoot,proto3" json:"previous_root,omitempty"`
	DataRoot             []byte   `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoreDocumentSalts) Reset()         { *m = CoreDocumentSalts{} }
func (m *CoreDocumentSalts) String() string { return proto.CompactTextString(m) }
func (*CoreDocumentSalts) ProtoMessage()    {}
func (*CoreDocumentSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_3ca8eeb62ad40595, []int{1}
}
func (m *CoreDocumentSalts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreDocumentSalts.Unmarshal(m, b)
}
func (m *CoreDocumentSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreDocumentSalts.Marshal(b, m, deterministic)
}
func (dst *CoreDocumentSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreDocumentSalts.Merge(dst, src)
}
func (m *CoreDocumentSalts) XXX_Size() int {
	return xxx_messageInfo_CoreDocumentSalts.Size(m)
}
func (m *CoreDocumentSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreDocumentSalts.DiscardUnknown(m)
}

var xxx_messageInfo_CoreDocumentSalts proto.InternalMessageInfo

func (m *CoreDocumentSalts) GetDocumentIdentifier() []byte {
	if m != nil {
		return m.DocumentIdentifier
	}
	return nil
}

func (m *CoreDocumentSalts) GetPreviousIdentifier() []byte {
	if m != nil {
		return m.PreviousIdentifier
	}
	return nil
}

func (m *CoreDocumentSalts) GetCurrentIdentifier() []byte {
	if m != nil {
		return m.CurrentIdentifier
	}
	return nil
}

func (m *CoreDocumentSalts) GetNextIdentifier() []byte {
	if m != nil {
		return m.NextIdentifier
	}
	return nil
}

func (m *CoreDocumentSalts) GetPreviousRoot() []byte {
	if m != nil {
		return m.PreviousRoot
	}
	return nil
}

func (m *CoreDocumentSalts) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

// Signature contains the entity ID, public key used and signature)
type Signature struct {
	// `entity_id` is the CentrifugeID of the entity signing the document.
	EntityId []byte `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// `public_key` is the public key of the `entity` used for signing the `CoreDocument`
	PublicKey            []byte               `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Signature            []byte               `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_coredocument_3ca8eeb62ad40595, []int{2}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetEntityId() []byte {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *Signature) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*CoreDocument)(nil), "coredocument.CoreDocument")
	proto.RegisterType((*CoreDocumentSalts)(nil), "coredocument.CoreDocumentSalts")
	proto.RegisterType((*Signature)(nil), "coredocument.Signature")
}

func init() {
	proto.RegisterFile("coredocument/coredocument.proto", fileDescriptor_coredocument_3ca8eeb62ad40595)
}

var fileDescriptor_coredocument_3ca8eeb62ad40595 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0x8d, 0x8e, 0xe5, 0x6d, 0xda, 0xad, 0x1e, 0x12, 0xa1, 0x80, 0x5a, 0x95, 0xc3,
	0xaa, 0x49, 0x4b, 0xa4, 0x72, 0x80, 0x13, 0x82, 0xb2, 0xcb, 0xb4, 0xcb, 0xd4, 0xc1, 0x85, 0x4b,
	0x94, 0x3f, 0x6e, 0x65, 0xd1, 0xc4, 0x91, 0xe3, 0x20, 0xf2, 0x29, 0xb8, 0x70, 0x87, 0x23, 0x5f,
	0x83, 0x0f, 0xc5, 0x1d, 0xd9, 0x8e, 0x13, 0xa7, 0x45, 0xe2, 0x82, 0x38, 0x55, 0x7e, 0x9f, 0xdf,
	0xf3, 0xd8, 0xaf, 0xf2, 0xa8, 0x30, 0x89, 0x29, 0xc3, 0x09, 0x8d, 0xcb, 0x14, 0x67, 0xdc, 0x37,
	0x0f, 0x5e, 0xce, 0x28, 0xa7, 0xc8, 0x31, 0x67, 0xe3, 0x47, 0x1b, 0x4a, 0x37, 0x5b, 0xec, 0x4b,
	0x2d, 0x2a, 0xd7, 0x7e, 0x98, 0x55, 0x0a, 0x1c, 0x4f, 0x76, 0x25, 0x4e, 0x52, 0x5c, 0xf0, 0x30,
	0xcd, 0x6b, 0xe0, 0x3c, 0x67, 0x38, 0x26, 0x05, 0xbe, 0xcc, 0x19, 0xa5, 0xeb, 0xc2, 0x6f, 0x7f,
	0x38, 0x55, 0x07, 0x05, 0xce, 0xbe, 0xf6, 0xc0, 0x79, 0x4b, 0x19, 0xbe, 0xaa, 0x6f, 0x45, 0x3e,
	0x9c, 0xe9, 0x17, 0x04, 0x24, 0xc1, 0x19, 0x27, 0x6b, 0x82, 0x99, 0x6b, 0x4f, 0xad, 0xb9, 0xb3,
	0x42, 0x5a, 0xba, 0x6e, 0x14, 0x61, 0xc8, 0x19, 0xfe, 0x44, 0x68, 0x59, 0x98, 0x86, 0x53, 0x65,
	0xd0, 0x92, 0x61, 0xb8, 0x04, 0x14, 0x97, 0x8c, 0xed, 0x5c, 0x70, 0x28, 0xf9, 0x51, 0xad, 0x18,
	0xf8, 0x39, 0x9c, 0x64, 0xf8, 0x73, 0x87, 0xbd, 0x27, 0xd9, 0xa1, 0x18, 0x1b, 0xe0, 0x05, 0x0c,
	0x9a, 0x97, 0x33, 0x4a, 0xb9, 0x7b, 0x5f, 0x60, 0xcb, 0xde, 0xf7, 0x9f, 0xbf, 0xc0, 0x5a, 0x39,
	0x5a, 0x5b, 0x51, 0xca, 0xd1, 0x1c, 0x9c, 0x82, 0x6c, 0x32, 0x92, 0x6d, 0x14, 0x0a, 0x26, 0xda,
	0xaf, 0x25, 0x49, 0x5e, 0xc0, 0xa0, 0x59, 0x4f, 0xa2, 0x07, 0x0a, 0xfd, 0xa1, 0x52, 0xb5, 0x26,
	0xd9, 0x19, 0xd8, 0x49, 0xc8, 0x43, 0xc5, 0xf5, 0x4c, 0xee, 0x58, 0xcc, 0x25, 0xf3, 0x0a, 0x40,
	0xc4, 0x87, 0xbc, 0x64, 0xb8, 0x70, 0x8f, 0xa6, 0x87, 0xf3, 0xfe, 0xe2, 0xa1, 0xd7, 0x29, 0xc3,
	0x9d, 0xd6, 0xf5, 0x83, 0x0c, 0x07, 0x7a, 0x0d, 0x03, 0x9c, 0x46, 0x38, 0x49, 0x70, 0x12, 0x88,
	0x50, 0x77, 0x30, 0xb5, 0xe6, 0xfd, 0xc5, 0x03, 0x4f, 0x55, 0xc2, 0xd3, 0x95, 0xf0, 0xde, 0x64,
	0x55, 0xb3, 0xbb, 0x76, 0x5c, 0x85, 0x3c, 0x44, 0x37, 0x70, 0xd6, 0x49, 0x08, 0x8a, 0x70, 0xcb,
	0x0b, 0x77, 0xf8, 0xf7, 0x9c, 0x91, 0x99, 0x73, 0x27, 0x5c, 0xe8, 0x3d, 0x20, 0xf3, 0xed, 0x75,
	0xd6, 0x89, 0xcc, 0x9a, 0x74, 0xd7, 0x32, 0x6b, 0x26, 0xcd, 0x4d, 0xac, 0x89, 0x49, 0x65, 0xf6,
	0xe5, 0x00, 0x46, 0x7b, 0xfc, 0x3f, 0xeb, 0xa6, 0xf5, 0xdf, 0xbb, 0xf9, 0xec, 0x8f, 0x2d, 0xda,
	0xa9, 0xcf, 0xe3, 0xbd, 0xfa, 0xb4, 0xbd, 0x99, 0x7d, 0xb3, 0xc0, 0x6e, 0x8a, 0x21, 0x50, 0x91,
	0xcd, 0xab, 0x80, 0x24, 0xf5, 0x3a, 0xc7, 0x6a, 0x70, 0x9d, 0xa0, 0xa7, 0x00, 0x79, 0x19, 0x6d,
	0x49, 0x1c, 0x7c, 0xc4, 0x55, 0x7d, 0x93, 0xad, 0x26, 0x37, 0xb8, 0x42, 0x4f, 0xc0, 0x6e, 0xfa,
	0x54, 0xaf, 0xd6, 0x0e, 0xd0, 0x4b, 0xb0, 0x9b, 0x3f, 0x13, 0xb9, 0x4c, 0x7f, 0x31, 0xde, 0xeb,
	0xc4, 0x3b, 0x4d, 0xac, 0x5a, 0x78, 0xf9, 0x02, 0x4e, 0x63, 0x9a, 0x76, 0xbe, 0xf9, 0x52, 0x7e,
	0x44, 0x7d, 0xba, 0x15, 0xf6, 0x5b, 0xeb, 0xc3, 0xd0, 0x44, 0xf2, 0x28, 0x3a, 0x92, 0xb9, 0xcf,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x8b, 0xff, 0xc1, 0x20, 0x05, 0x00, 0x00,
}
