// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: content_manifest.proto

package protobuf

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

type EContentDeltaChunkDataLocation int32

const (
	EContentDeltaChunkDataLocation_k_EContentDeltaChunkDataLocationInProtobuf    EContentDeltaChunkDataLocation = 0
	EContentDeltaChunkDataLocation_k_EContentDeltaChunkDataLocationAfterProtobuf EContentDeltaChunkDataLocation = 1
)

// Enum value maps for EContentDeltaChunkDataLocation.
var (
	EContentDeltaChunkDataLocation_name = map[int32]string{
		0: "k_EContentDeltaChunkDataLocationInProtobuf",
		1: "k_EContentDeltaChunkDataLocationAfterProtobuf",
	}
	EContentDeltaChunkDataLocation_value = map[string]int32{
		"k_EContentDeltaChunkDataLocationInProtobuf":    0,
		"k_EContentDeltaChunkDataLocationAfterProtobuf": 1,
	}
)

func (x EContentDeltaChunkDataLocation) Enum() *EContentDeltaChunkDataLocation {
	p := new(EContentDeltaChunkDataLocation)
	*p = x
	return p
}

func (x EContentDeltaChunkDataLocation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EContentDeltaChunkDataLocation) Descriptor() protoreflect.EnumDescriptor {
	return file_content_manifest_proto_enumTypes[0].Descriptor()
}

func (EContentDeltaChunkDataLocation) Type() protoreflect.EnumType {
	return &file_content_manifest_proto_enumTypes[0]
}

func (x EContentDeltaChunkDataLocation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EContentDeltaChunkDataLocation) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EContentDeltaChunkDataLocation(num)
	return nil
}

// Deprecated: Use EContentDeltaChunkDataLocation.Descriptor instead.
func (EContentDeltaChunkDataLocation) EnumDescriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{0}
}

type ContentManifestPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mappings []*ContentManifestPayload_FileMapping `protobuf:"bytes,1,rep,name=mappings" json:"mappings,omitempty"`
}

func (x *ContentManifestPayload) Reset() {
	*x = ContentManifestPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentManifestPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentManifestPayload) ProtoMessage() {}

func (x *ContentManifestPayload) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentManifestPayload.ProtoReflect.Descriptor instead.
func (*ContentManifestPayload) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *ContentManifestPayload) GetMappings() []*ContentManifestPayload_FileMapping {
	if x != nil {
		return x.Mappings
	}
	return nil
}

type ContentManifestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepotId            *uint32 `protobuf:"varint,1,opt,name=depot_id,json=depotId" json:"depot_id,omitempty"`
	GidManifest        *uint64 `protobuf:"varint,2,opt,name=gid_manifest,json=gidManifest" json:"gid_manifest,omitempty"`
	CreationTime       *uint32 `protobuf:"varint,3,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	FilenamesEncrypted *bool   `protobuf:"varint,4,opt,name=filenames_encrypted,json=filenamesEncrypted" json:"filenames_encrypted,omitempty"`
	CbDiskOriginal     *uint64 `protobuf:"varint,5,opt,name=cb_disk_original,json=cbDiskOriginal" json:"cb_disk_original,omitempty"`
	CbDiskCompressed   *uint64 `protobuf:"varint,6,opt,name=cb_disk_compressed,json=cbDiskCompressed" json:"cb_disk_compressed,omitempty"`
	UniqueChunks       *uint32 `protobuf:"varint,7,opt,name=unique_chunks,json=uniqueChunks" json:"unique_chunks,omitempty"`
	CrcEncrypted       *uint32 `protobuf:"varint,8,opt,name=crc_encrypted,json=crcEncrypted" json:"crc_encrypted,omitempty"`
	CrcClear           *uint32 `protobuf:"varint,9,opt,name=crc_clear,json=crcClear" json:"crc_clear,omitempty"`
}

func (x *ContentManifestMetadata) Reset() {
	*x = ContentManifestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentManifestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentManifestMetadata) ProtoMessage() {}

func (x *ContentManifestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentManifestMetadata.ProtoReflect.Descriptor instead.
func (*ContentManifestMetadata) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *ContentManifestMetadata) GetDepotId() uint32 {
	if x != nil && x.DepotId != nil {
		return *x.DepotId
	}
	return 0
}

func (x *ContentManifestMetadata) GetGidManifest() uint64 {
	if x != nil && x.GidManifest != nil {
		return *x.GidManifest
	}
	return 0
}

func (x *ContentManifestMetadata) GetCreationTime() uint32 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *ContentManifestMetadata) GetFilenamesEncrypted() bool {
	if x != nil && x.FilenamesEncrypted != nil {
		return *x.FilenamesEncrypted
	}
	return false
}

func (x *ContentManifestMetadata) GetCbDiskOriginal() uint64 {
	if x != nil && x.CbDiskOriginal != nil {
		return *x.CbDiskOriginal
	}
	return 0
}

func (x *ContentManifestMetadata) GetCbDiskCompressed() uint64 {
	if x != nil && x.CbDiskCompressed != nil {
		return *x.CbDiskCompressed
	}
	return 0
}

func (x *ContentManifestMetadata) GetUniqueChunks() uint32 {
	if x != nil && x.UniqueChunks != nil {
		return *x.UniqueChunks
	}
	return 0
}

func (x *ContentManifestMetadata) GetCrcEncrypted() uint32 {
	if x != nil && x.CrcEncrypted != nil {
		return *x.CrcEncrypted
	}
	return 0
}

func (x *ContentManifestMetadata) GetCrcClear() uint32 {
	if x != nil && x.CrcClear != nil {
		return *x.CrcClear
	}
	return 0
}

type ContentManifestSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
}

func (x *ContentManifestSignature) Reset() {
	*x = ContentManifestSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentManifestSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentManifestSignature) ProtoMessage() {}

func (x *ContentManifestSignature) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentManifestSignature.ProtoReflect.Descriptor instead.
func (*ContentManifestSignature) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *ContentManifestSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ContentDeltaChunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DepotId           *uint32                          `protobuf:"varint,1,opt,name=depot_id,json=depotId" json:"depot_id,omitempty"`
	ManifestIdSource  *uint64                          `protobuf:"varint,2,opt,name=manifest_id_source,json=manifestIdSource" json:"manifest_id_source,omitempty"`
	ManifestIdTarget  *uint64                          `protobuf:"varint,3,opt,name=manifest_id_target,json=manifestIdTarget" json:"manifest_id_target,omitempty"`
	DeltaChunks       []*ContentDeltaChunks_DeltaChunk `protobuf:"bytes,4,rep,name=deltaChunks" json:"deltaChunks,omitempty"`
	ChunkDataLocation *EContentDeltaChunkDataLocation  `protobuf:"varint,5,opt,name=chunk_data_location,json=chunkDataLocation,enum=EContentDeltaChunkDataLocation,def=0" json:"chunk_data_location,omitempty"`
}

// Default values for ContentDeltaChunks fields.
const (
	Default_ContentDeltaChunks_ChunkDataLocation = EContentDeltaChunkDataLocation_k_EContentDeltaChunkDataLocationInProtobuf
)

func (x *ContentDeltaChunks) Reset() {
	*x = ContentDeltaChunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentDeltaChunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentDeltaChunks) ProtoMessage() {}

func (x *ContentDeltaChunks) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentDeltaChunks.ProtoReflect.Descriptor instead.
func (*ContentDeltaChunks) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{3}
}

func (x *ContentDeltaChunks) GetDepotId() uint32 {
	if x != nil && x.DepotId != nil {
		return *x.DepotId
	}
	return 0
}

func (x *ContentDeltaChunks) GetManifestIdSource() uint64 {
	if x != nil && x.ManifestIdSource != nil {
		return *x.ManifestIdSource
	}
	return 0
}

func (x *ContentDeltaChunks) GetManifestIdTarget() uint64 {
	if x != nil && x.ManifestIdTarget != nil {
		return *x.ManifestIdTarget
	}
	return 0
}

func (x *ContentDeltaChunks) GetDeltaChunks() []*ContentDeltaChunks_DeltaChunk {
	if x != nil {
		return x.DeltaChunks
	}
	return nil
}

func (x *ContentDeltaChunks) GetChunkDataLocation() EContentDeltaChunkDataLocation {
	if x != nil && x.ChunkDataLocation != nil {
		return *x.ChunkDataLocation
	}
	return Default_ContentDeltaChunks_ChunkDataLocation
}

type ContentManifestPayload_FileMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename    *string                                         `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Size        *uint64                                         `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Flags       *uint32                                         `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	ShaFilename []byte                                          `protobuf:"bytes,4,opt,name=sha_filename,json=shaFilename" json:"sha_filename,omitempty"`
	ShaContent  []byte                                          `protobuf:"bytes,5,opt,name=sha_content,json=shaContent" json:"sha_content,omitempty"`
	Chunks      []*ContentManifestPayload_FileMapping_ChunkData `protobuf:"bytes,6,rep,name=chunks" json:"chunks,omitempty"`
	Linktarget  *string                                         `protobuf:"bytes,7,opt,name=linktarget" json:"linktarget,omitempty"`
}

func (x *ContentManifestPayload_FileMapping) Reset() {
	*x = ContentManifestPayload_FileMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentManifestPayload_FileMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentManifestPayload_FileMapping) ProtoMessage() {}

func (x *ContentManifestPayload_FileMapping) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentManifestPayload_FileMapping.ProtoReflect.Descriptor instead.
func (*ContentManifestPayload_FileMapping) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ContentManifestPayload_FileMapping) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *ContentManifestPayload_FileMapping) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *ContentManifestPayload_FileMapping) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *ContentManifestPayload_FileMapping) GetShaFilename() []byte {
	if x != nil {
		return x.ShaFilename
	}
	return nil
}

func (x *ContentManifestPayload_FileMapping) GetShaContent() []byte {
	if x != nil {
		return x.ShaContent
	}
	return nil
}

func (x *ContentManifestPayload_FileMapping) GetChunks() []*ContentManifestPayload_FileMapping_ChunkData {
	if x != nil {
		return x.Chunks
	}
	return nil
}

func (x *ContentManifestPayload_FileMapping) GetLinktarget() string {
	if x != nil && x.Linktarget != nil {
		return *x.Linktarget
	}
	return ""
}

type ContentManifestPayload_FileMapping_ChunkData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha          []byte  `protobuf:"bytes,1,opt,name=sha" json:"sha,omitempty"`
	Crc          *uint32 `protobuf:"fixed32,2,opt,name=crc" json:"crc,omitempty"`
	Offset       *uint64 `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	CbOriginal   *uint32 `protobuf:"varint,4,opt,name=cb_original,json=cbOriginal" json:"cb_original,omitempty"`
	CbCompressed *uint32 `protobuf:"varint,5,opt,name=cb_compressed,json=cbCompressed" json:"cb_compressed,omitempty"`
}

func (x *ContentManifestPayload_FileMapping_ChunkData) Reset() {
	*x = ContentManifestPayload_FileMapping_ChunkData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentManifestPayload_FileMapping_ChunkData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentManifestPayload_FileMapping_ChunkData) ProtoMessage() {}

func (x *ContentManifestPayload_FileMapping_ChunkData) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentManifestPayload_FileMapping_ChunkData.ProtoReflect.Descriptor instead.
func (*ContentManifestPayload_FileMapping_ChunkData) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ContentManifestPayload_FileMapping_ChunkData) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

func (x *ContentManifestPayload_FileMapping_ChunkData) GetCrc() uint32 {
	if x != nil && x.Crc != nil {
		return *x.Crc
	}
	return 0
}

func (x *ContentManifestPayload_FileMapping_ChunkData) GetOffset() uint64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *ContentManifestPayload_FileMapping_ChunkData) GetCbOriginal() uint32 {
	if x != nil && x.CbOriginal != nil {
		return *x.CbOriginal
	}
	return 0
}

func (x *ContentManifestPayload_FileMapping_ChunkData) GetCbCompressed() uint32 {
	if x != nil && x.CbCompressed != nil {
		return *x.CbCompressed
	}
	return 0
}

type ContentDeltaChunks_DeltaChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShaSource    []byte  `protobuf:"bytes,1,opt,name=sha_source,json=shaSource" json:"sha_source,omitempty"`
	ShaTarget    []byte  `protobuf:"bytes,2,opt,name=sha_target,json=shaTarget" json:"sha_target,omitempty"`
	SizeOriginal *uint32 `protobuf:"varint,3,opt,name=size_original,json=sizeOriginal" json:"size_original,omitempty"`
	PatchMethod  *uint32 `protobuf:"varint,4,opt,name=patch_method,json=patchMethod" json:"patch_method,omitempty"`
	Chunk        []byte  `protobuf:"bytes,5,opt,name=chunk" json:"chunk,omitempty"`
	SizeDelta    *uint32 `protobuf:"varint,6,opt,name=size_delta,json=sizeDelta" json:"size_delta,omitempty"`
}

func (x *ContentDeltaChunks_DeltaChunk) Reset() {
	*x = ContentDeltaChunks_DeltaChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_manifest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentDeltaChunks_DeltaChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentDeltaChunks_DeltaChunk) ProtoMessage() {}

func (x *ContentDeltaChunks_DeltaChunk) ProtoReflect() protoreflect.Message {
	mi := &file_content_manifest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentDeltaChunks_DeltaChunk.ProtoReflect.Descriptor instead.
func (*ContentDeltaChunks_DeltaChunk) Descriptor() ([]byte, []int) {
	return file_content_manifest_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ContentDeltaChunks_DeltaChunk) GetShaSource() []byte {
	if x != nil {
		return x.ShaSource
	}
	return nil
}

func (x *ContentDeltaChunks_DeltaChunk) GetShaTarget() []byte {
	if x != nil {
		return x.ShaTarget
	}
	return nil
}

func (x *ContentDeltaChunks_DeltaChunk) GetSizeOriginal() uint32 {
	if x != nil && x.SizeOriginal != nil {
		return *x.SizeOriginal
	}
	return 0
}

func (x *ContentDeltaChunks_DeltaChunk) GetPatchMethod() uint32 {
	if x != nil && x.PatchMethod != nil {
		return *x.PatchMethod
	}
	return 0
}

func (x *ContentDeltaChunks_DeltaChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *ContentDeltaChunks_DeltaChunk) GetSizeDelta() uint32 {
	if x != nil && x.SizeDelta != nil {
		return *x.SizeDelta
	}
	return 0
}

var File_content_manifest_proto protoreflect.FileDescriptor

var file_content_manifest_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x03, 0x0a, 0x16, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x73, 0x1a, 0x8e, 0x03, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x68,
	0x61, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x73, 0x68, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x68, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x45,
	0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x8d, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x03, 0x63, 0x72, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x62, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x62, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0xec, 0x02, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x67, 0x69, 0x64, 0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x67, 0x69, 0x64, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x6b,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x63, 0x62, 0x44, 0x69, 0x73, 0x6b, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12,
	0x2c, 0x0a, 0x12, 0x63, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x62, 0x44,
	0x69, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x63, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x72, 0x63, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x63, 0x5f, 0x63,
	0x6c, 0x65, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x72, 0x63, 0x43,
	0x6c, 0x65, 0x61, 0x72, 0x22, 0x38, 0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x94,
	0x04, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6d, 0x61,
	0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x49, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x49, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x40, 0x0a, 0x0b,
	0x64, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x12, 0x7b,
	0x0a, 0x13, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x45, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x2a, 0x6b, 0x5f,
	0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75,
	0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x52, 0x11, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xc7, 0x01, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x68, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x61,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x68, 0x61, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x73, 0x69, 0x7a, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x64,
	0x65, 0x6c, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x2a, 0x83, 0x01, 0x0a, 0x1e, 0x45, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x2a, 0x6b, 0x5f, 0x45, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x10, 0x00, 0x12, 0x31, 0x0a, 0x2d, 0x6b, 0x5f, 0x45, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x10, 0x01, 0x42, 0x05, 0x48, 0x01, 0x80,
	0x01, 0x00,
}

var (
	file_content_manifest_proto_rawDescOnce sync.Once
	file_content_manifest_proto_rawDescData = file_content_manifest_proto_rawDesc
)

func file_content_manifest_proto_rawDescGZIP() []byte {
	file_content_manifest_proto_rawDescOnce.Do(func() {
		file_content_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_manifest_proto_rawDescData)
	})
	return file_content_manifest_proto_rawDescData
}

var file_content_manifest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_content_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_content_manifest_proto_goTypes = []interface{}{
	(EContentDeltaChunkDataLocation)(0),                  // 0: EContentDeltaChunkDataLocation
	(*ContentManifestPayload)(nil),                       // 1: ContentManifestPayload
	(*ContentManifestMetadata)(nil),                      // 2: ContentManifestMetadata
	(*ContentManifestSignature)(nil),                     // 3: ContentManifestSignature
	(*ContentDeltaChunks)(nil),                           // 4: ContentDeltaChunks
	(*ContentManifestPayload_FileMapping)(nil),           // 5: ContentManifestPayload.FileMapping
	(*ContentManifestPayload_FileMapping_ChunkData)(nil), // 6: ContentManifestPayload.FileMapping.ChunkData
	(*ContentDeltaChunks_DeltaChunk)(nil),                // 7: ContentDeltaChunks.DeltaChunk
}
var file_content_manifest_proto_depIdxs = []int32{
	5, // 0: ContentManifestPayload.mappings:type_name -> ContentManifestPayload.FileMapping
	7, // 1: ContentDeltaChunks.deltaChunks:type_name -> ContentDeltaChunks.DeltaChunk
	0, // 2: ContentDeltaChunks.chunk_data_location:type_name -> EContentDeltaChunkDataLocation
	6, // 3: ContentManifestPayload.FileMapping.chunks:type_name -> ContentManifestPayload.FileMapping.ChunkData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_content_manifest_proto_init() }
func file_content_manifest_proto_init() {
	if File_content_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentManifestPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentManifestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentManifestSignature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentDeltaChunks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentManifestPayload_FileMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentManifestPayload_FileMapping_ChunkData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_content_manifest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentDeltaChunks_DeltaChunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_manifest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_manifest_proto_goTypes,
		DependencyIndexes: file_content_manifest_proto_depIdxs,
		EnumInfos:         file_content_manifest_proto_enumTypes,
		MessageInfos:      file_content_manifest_proto_msgTypes,
	}.Build()
	File_content_manifest_proto = out.File
	file_content_manifest_proto_rawDesc = nil
	file_content_manifest_proto_goTypes = nil
	file_content_manifest_proto_depIdxs = nil
}
