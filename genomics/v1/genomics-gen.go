// Package genomics provides access to the Genomics API.
//
// Usage example:
//
//   import "google.golang.org/api/genomics/v1"
//   ...
//   genomicsService, err := genomics.New(oauthHttpClient)
package genomics

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "genomics:v1"
const apiName = "genomics"
const apiVersion = "v1"
const basePath = "https://genomics.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data in Google BigQuery
	BigqueryScope = "https://www.googleapis.com/auth/bigquery"

	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your data in Google Cloud Storage
	DevstorageReadWriteScope = "https://www.googleapis.com/auth/devstorage.read_write"

	// View and manage Genomics data
	GenomicsScope = "https://www.googleapis.com/auth/genomics"

	// View Genomics data
	GenomicsReadonlyScope = "https://www.googleapis.com/auth/genomics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Callsets = NewCallsetsService(s)
	s.Datasets = NewDatasetsService(s)
	s.Operations = NewOperationsService(s)
	s.Readgroupsets = NewReadgroupsetsService(s)
	s.Reads = NewReadsService(s)
	s.References = NewReferencesService(s)
	s.Referencesets = NewReferencesetsService(s)
	s.Variants = NewVariantsService(s)
	s.Variantsets = NewVariantsetsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Callsets *CallsetsService

	Datasets *DatasetsService

	Operations *OperationsService

	Readgroupsets *ReadgroupsetsService

	Reads *ReadsService

	References *ReferencesService

	Referencesets *ReferencesetsService

	Variants *VariantsService

	Variantsets *VariantsetsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCallsetsService(s *Service) *CallsetsService {
	rs := &CallsetsService{s: s}
	return rs
}

type CallsetsService struct {
	s *Service
}

func NewDatasetsService(s *Service) *DatasetsService {
	rs := &DatasetsService{s: s}
	return rs
}

type DatasetsService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewReadgroupsetsService(s *Service) *ReadgroupsetsService {
	rs := &ReadgroupsetsService{s: s}
	rs.Coveragebuckets = NewReadgroupsetsCoveragebucketsService(s)
	return rs
}

type ReadgroupsetsService struct {
	s *Service

	Coveragebuckets *ReadgroupsetsCoveragebucketsService
}

func NewReadgroupsetsCoveragebucketsService(s *Service) *ReadgroupsetsCoveragebucketsService {
	rs := &ReadgroupsetsCoveragebucketsService{s: s}
	return rs
}

type ReadgroupsetsCoveragebucketsService struct {
	s *Service
}

func NewReadsService(s *Service) *ReadsService {
	rs := &ReadsService{s: s}
	return rs
}

type ReadsService struct {
	s *Service
}

func NewReferencesService(s *Service) *ReferencesService {
	rs := &ReferencesService{s: s}
	rs.Bases = NewReferencesBasesService(s)
	return rs
}

type ReferencesService struct {
	s *Service

	Bases *ReferencesBasesService
}

func NewReferencesBasesService(s *Service) *ReferencesBasesService {
	rs := &ReferencesBasesService{s: s}
	return rs
}

type ReferencesBasesService struct {
	s *Service
}

func NewReferencesetsService(s *Service) *ReferencesetsService {
	rs := &ReferencesetsService{s: s}
	return rs
}

type ReferencesetsService struct {
	s *Service
}

func NewVariantsService(s *Service) *VariantsService {
	rs := &VariantsService{s: s}
	return rs
}

type VariantsService struct {
	s *Service
}

func NewVariantsetsService(s *Service) *VariantsetsService {
	rs := &VariantsetsService{s: s}
	return rs
}

type VariantsetsService struct {
	s *Service
}

type CallSet struct {
	// Created: The date this call set was created in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// Id: The Google generated ID of the call set, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional call set information. This must be of the
	// form map (string key mapping to a list of string values).
	Info *CallSetInfo `json:"info,omitempty"`

	// Name: The call set name.
	Name string `json:"name,omitempty"`

	// SampleId: The sample ID this call set corresponds to.
	SampleId string `json:"sampleId,omitempty"`

	// VariantSetIds: The IDs of the variant sets this call set belongs to.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type CallSetInfo struct {
}

type CancelOperationRequest struct {
}

type CigarUnit struct {
	// Possible values:
	//   "OPERATION_UNSPECIFIED"
	//   "ALIGNMENT_MATCH"
	//   "INSERT"
	//   "DELETE"
	//   "SKIP"
	//   "CLIP_SOFT"
	//   "CLIP_HARD"
	//   "PAD"
	//   "SEQUENCE_MATCH"
	//   "SEQUENCE_MISMATCH"
	Operation string `json:"operation,omitempty"`

	// OperationLength: The number of genomic bases that the operation runs
	// for. Required.
	OperationLength int64 `json:"operationLength,omitempty,string"`

	// ReferenceSequence: `referenceSequence` is only used at mismatches
	// (`SEQUENCE_MISMATCH`) and deletions (`DELETE`). Filling this field
	// replaces SAM's MD tag. If the relevant information is not available,
	// this field is unset.
	ReferenceSequence string `json:"referenceSequence,omitempty"`
}

type CoverageBucket struct {
	// MeanCoverage: The average number of reads which are aligned to each
	// individual reference base in this bucket.
	MeanCoverage float64 `json:"meanCoverage,omitempty"`

	// Range: The genomic coordinate range spanned by this bucket.
	Range *Range `json:"range,omitempty"`
}

type Dataset struct {
	// CreateTime: The time this dataset was created, in seconds from the
	// epoch.
	CreateTime string `json:"createTime,omitempty"`

	// Id: The server-assigned dataset ID, immutable.
	Id string `json:"id,omitempty"`

	// Name: The dataset name.
	Name string `json:"name,omitempty"`

	// ProjectId: The Google Developers Console project ID that this dataset
	// belongs to.
	ProjectId string `json:"projectId,omitempty"`
}

type Empty struct {
}

type Experiment struct {
	// InstrumentModel: The instrument model used as part of this
	// experiment. This maps to sequencing technology in BAM.
	InstrumentModel string `json:"instrumentModel,omitempty"`

	// LibraryId: The library used as part of this experiment. Note: This is
	// not an actual ID within this repository, but rather an identifier for
	// a library which may be meaningful to some external system.
	LibraryId string `json:"libraryId,omitempty"`

	// PlatformUnit: The platform unit used as part of this experiment e.g.
	// flowcell-barcode.lane for Illumina or slide for SOLiD. Corresponds to
	// the @RG PU field in the SAM spec.
	PlatformUnit string `json:"platformUnit,omitempty"`

	// SequencingCenter: The sequencing center used as part of this
	// experiment.
	SequencingCenter string `json:"sequencingCenter,omitempty"`
}

type ExportReadGroupSetRequest struct {
	// ExportUri: Required. A Google Cloud Storage URI for the exported BAM
	// file. The currently authenticated user must have write access to the
	// new file. An error will be returned if the URI already contains data.
	ExportUri string `json:"exportUri,omitempty"`

	// ProjectId: Required. The Google Developers Console project ID that
	// owns this export.
	ProjectId string `json:"projectId,omitempty"`

	// ReferenceNames: The reference names to export. If this is not
	// specified, all reference sequences, including unmapped reads, are
	// exported. Use `*` to export only unmapped reads.
	ReferenceNames []string `json:"referenceNames,omitempty"`
}

type ExportVariantSetRequest struct {
	// BigqueryDataset: Required. The BigQuery dataset to export data to.
	// This dataset must already exist. Note that this is distinct from the
	// Genomics concept of "dataset".
	BigqueryDataset string `json:"bigqueryDataset,omitempty"`

	// BigqueryTable: Required. The BigQuery table to export data to. If the
	// table doesn't exist, it will be created. If it already exists, it
	// will be overwritten.
	BigqueryTable string `json:"bigqueryTable,omitempty"`

	// CallSetIds: If provided, only variant call information from the
	// specified call sets will be exported. By default all variant calls
	// are exported.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// Format: The format for the exported data.
	//
	// Possible values:
	//   "FORMAT_UNSPECIFIED"
	//   "FORMAT_BIGQUERY"
	Format string `json:"format,omitempty"`

	// ProjectId: Required. The Google Cloud project ID that owns the
	// destination BigQuery dataset. The caller must have WRITE access to
	// this project. This project will also own the resulting export job.
	ProjectId string `json:"projectId,omitempty"`
}

type ImportReadGroupSetsRequest struct {
	// DatasetId: Required. The ID of the dataset these read group sets will
	// belong to. The caller must have WRITE permissions to this dataset.
	DatasetId string `json:"datasetId,omitempty"`

	// PartitionStrategy: The partition strategy describes how read groups
	// are partitioned into read group sets.
	//
	// Possible values:
	//   "PARTITION_STRATEGY_UNSPECIFIED"
	//   "PER_FILE_PER_SAMPLE"
	//   "MERGE_ALL"
	PartitionStrategy string `json:"partitionStrategy,omitempty"`

	// ReferenceSetId: The reference set to which the imported read group
	// sets are aligned to, if any. The reference names of this reference
	// set must be a superset of those found in the imported file headers.
	// If no reference set id is provided, a best effort is made to
	// associate with a matching reference set.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// SourceUris: A list of URIs pointing at BAM files in Google Cloud
	// Storage.
	SourceUris []string `json:"sourceUris,omitempty"`
}

type ImportVariantsRequest struct {
	// Format: The format of the variant data being imported. If
	// unspecified, defaults to to `VCF`.
	//
	// Possible values:
	//   "FORMAT_UNSPECIFIED"
	//   "FORMAT_VCF"
	//   "FORMAT_COMPLETE_GENOMICS"
	Format string `json:"format,omitempty"`

	// NormalizeReferenceNames: Convert reference names to the canonical
	// representation. hg19 haploytypes (those reference names containing
	// "_hap") are not modified in any way. All other reference names are
	// modified according to the following rules: The reference name is
	// capitalized. The "chr" prefix is dropped for all autosomes and sex
	// chromsomes. For example "chr17" becomes "17" and "chrX" becomes "X".
	// All mitochondrial chromosomes ("chrM", "chrMT", etc) become "MT".
	NormalizeReferenceNames bool `json:"normalizeReferenceNames,omitempty"`

	// SourceUris: A list of URIs referencing variant files in Google Cloud
	// Storage. URIs can include wildcards [as described
	// here](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNa
	// mes). Note that recursive wildcards ('**') are not supported.
	SourceUris []string `json:"sourceUris,omitempty"`

	// VariantSetId: Required. The variant set to which variant data should
	// be imported.
	VariantSetId string `json:"variantSetId,omitempty"`
}

type LinearAlignment struct {
	// Cigar: Represents the local alignment of this sequence (alignment
	// matches, indels, etc) against the reference.
	Cigar []*CigarUnit `json:"cigar,omitempty"`

	// MappingQuality: The mapping quality of this alignment. Represents how
	// likely the read maps to this position as opposed to other locations.
	MappingQuality int64 `json:"mappingQuality,omitempty"`

	// Position: The position of this alignment.
	Position *Position `json:"position,omitempty"`
}

type ListBasesResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Offset: The offset position (0-based) of the given `sequence` from
	// the start of this `Reference`. This value will differ for each page
	// in a paginated request.
	Offset int64 `json:"offset,omitempty,string"`

	// Sequence: A substring of the bases that make up this reference.
	Sequence string `json:"sequence,omitempty"`
}

type ListCoverageBucketsResponse struct {
	// BucketWidth: The length of each coverage bucket in base pairs. Note
	// that buckets at the end of a reference sequence may be shorter. This
	// value is omitted if the bucket width is infinity (the default
	// behaviour, with no range or `targetBucketWidth`).
	BucketWidth int64 `json:"bucketWidth,omitempty,string"`

	// CoverageBuckets: The coverage buckets. The list of buckets is sparse;
	// a bucket with 0 overlapping reads is not returned. A bucket never
	// crosses more than one reference sequence. Each bucket has width
	// `bucketWidth`, unless its end is the end of the reference sequence.
	CoverageBuckets []*CoverageBucket `json:"coverageBuckets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListDatasetsResponse struct {
	// Datasets: The list of matching Datasets.
	Datasets []*Dataset `json:"datasets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that match the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`
}

type MergeVariantsRequest struct {
	// VariantSetId: The destination variant set.
	VariantSetId string `json:"variantSetId,omitempty"`

	// Variants: The variants to be merged with existing variants.
	Variants []*Variant `json:"variants,omitempty"`
}

type Operation struct {
	// Done: If the value is false, it means the operation is still in
	// progress. If true, the operation is completed and the `result` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure.
	Error *Status `json:"error,omitempty"`

	// Metadata: An
	// [OperationMetadata][google.genomics.v1.OperationMetadata] object.
	// This will always be returned with the
	// [Operation][google.longrunning.Operation].
	Metadata OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that originally returns it. For example:
	// `operations/CJHU7Oi_ChDrveSpBRjfuL-qzoWAgEw`
	Name string `json:"name,omitempty"`

	// Response: If importing
	// [ReadGroupSets][google.genomics.v1.ReadGroupSet], an
	// [ImportReadGroupSetsResponse][google.genomics.v1.ImportReadGroupSetsRe
	// sponse] is returned. If importing
	// [Variants][google.genomics.v1.Variant], an
	// [ImportVariantsResponse][google.genomics.v1.ImportVariantsResponse]
	// is returned. For exports, an empty response is returned.
	Response OperationResponse `json:"response,omitempty"`
}

type OperationMetadata interface{}

type OperationResponse interface{}

type Position struct {
	// Position: The 0-based offset from the start of the forward strand for
	// that reference.
	Position int64 `json:"position,omitempty,string"`

	// ReferenceName: The name of the reference in whatever reference set is
	// being used.
	ReferenceName string `json:"referenceName,omitempty"`

	// ReverseStrand: Whether this position is on the reverse strand, as
	// opposed to the forward strand.
	ReverseStrand bool `json:"reverseStrand,omitempty"`
}

type Program struct {
	// CommandLine: The command line used to run this program.
	CommandLine string `json:"commandLine,omitempty"`

	// Id: The user specified locally unique ID of the program. Used along
	// with `prevProgramId` to define an ordering between programs.
	Id string `json:"id,omitempty"`

	// Name: The name of the program.
	Name string `json:"name,omitempty"`

	// PrevProgramId: The ID of the program run before this one.
	PrevProgramId string `json:"prevProgramId,omitempty"`

	// Version: The version of the program run.
	Version string `json:"version,omitempty"`
}

type Range struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive.
	End int64 `json:"end,omitempty,string"`

	// ReferenceName: The reference sequence name, for example `chr1`, `1`,
	// or `chrX`.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive.
	Start int64 `json:"start,omitempty,string"`
}

type Read struct {
	// AlignedQuality: The quality of the read sequence contained in this
	// alignment record. `alignedSequence` and `alignedQuality` may be
	// shorter than the full read sequence and quality. This will occur if
	// the alignment is part of a chimeric alignment, or if the read was
	// trimmed. When this occurs, the CIGAR for this read will begin/end
	// with a hard clip operator that will indicate the length of the
	// excised sequence.
	AlignedQuality []int64 `json:"alignedQuality,omitempty"`

	// AlignedSequence: The bases of the read sequence contained in this
	// alignment record, *without CIGAR operations applied*.
	// `alignedSequence` and `alignedQuality` may be shorter than the full
	// read sequence and quality. This will occur if the alignment is part
	// of a chimeric alignment, or if the read was trimmed. When this
	// occurs, the CIGAR for this read will begin/end with a hard clip
	// operator that will indicate the length of the excised sequence.
	AlignedSequence string `json:"alignedSequence,omitempty"`

	// Alignment: The linear alignment for this alignment record. This field
	// will be null if the read is unmapped.
	Alignment *LinearAlignment `json:"alignment,omitempty"`

	// DuplicateFragment: The fragment is a PCR or optical duplicate (SAM
	// flag 0x400)
	DuplicateFragment bool `json:"duplicateFragment,omitempty"`

	// FailedVendorQualityChecks: SAM flag 0x200
	FailedVendorQualityChecks bool `json:"failedVendorQualityChecks,omitempty"`

	// FragmentLength: The observed length of the fragment, equivalent to
	// TLEN in SAM.
	FragmentLength int64 `json:"fragmentLength,omitempty"`

	// FragmentName: The fragment name. Equivalent to QNAME (query template
	// name) in SAM.
	FragmentName string `json:"fragmentName,omitempty"`

	// Id: The unique ID for this read. This is a generated unique ID, not
	// to be confused with fragmentName.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read alignment information. This must be of
	// the form map (string key mapping to a list of string values).
	Info *ReadInfo `json:"info,omitempty"`

	// NextMatePosition: The mapping of the primary alignment of the
	// `(readNumber+1)%numberReads` read in the fragment. It replaces mate
	// position and mate strand in SAM.
	NextMatePosition *Position `json:"nextMatePosition,omitempty"`

	// NumberReads: The number of reads in the fragment (extension to SAM
	// flag 0x1).
	NumberReads int64 `json:"numberReads,omitempty"`

	// ProperPlacement: The orientation and the distance between reads from
	// the fragment are consistent with the sequencing protocol (SAM flag
	// 0x2)
	ProperPlacement bool `json:"properPlacement,omitempty"`

	// ReadGroupId: The ID of the read group this read belongs to. (Every
	// read must belong to exactly one read group.)
	ReadGroupId string `json:"readGroupId,omitempty"`

	// ReadGroupSetId: The ID of the read group set this read belongs to.
	// (Every read must belong to exactly one read group set.)
	ReadGroupSetId string `json:"readGroupSetId,omitempty"`

	// ReadNumber: The read number in sequencing. 0-based and less than
	// numberReads. This field replaces SAM flag 0x40 and 0x80.
	ReadNumber int64 `json:"readNumber,omitempty"`

	// SecondaryAlignment: Whether this alignment is secondary. Equivalent
	// to SAM flag 0x100. A secondary alignment represents an alternative to
	// the primary alignment for this read. Aligners may return secondary
	// alignments if a read can map ambiguously to multiple coordinates in
	// the genome. By convention, each read has one and only one alignment
	// where both `secondaryAlignment` and `supplementaryAlignment` are
	// false.
	SecondaryAlignment bool `json:"secondaryAlignment,omitempty"`

	// SupplementaryAlignment: Whether this alignment is supplementary.
	// Equivalent to SAM flag 0x800. Supplementary alignments are used in
	// the representation of a chimeric alignment. In a chimeric alignment,
	// a read is split into multiple linear alignments that map to different
	// reference contigs. The first linear alignment in the read will be
	// designated as the representative alignment; the remaining linear
	// alignments will be designated as supplementary alignments. These
	// alignments may have different mapping quality scores. In each linear
	// alignment in a chimeric alignment, the read will be hard clipped. The
	// `alignedSequence` and `alignedQuality` fields in the alignment record
	// will only represent the bases for its respective linear alignment.
	SupplementaryAlignment bool `json:"supplementaryAlignment,omitempty"`
}

type ReadInfo struct {
}

type ReadGroup struct {
	// DatasetId: The ID of the dataset this read group belongs to.
	DatasetId string `json:"datasetId,omitempty"`

	// Description: A free-form text description of this read group.
	Description string `json:"description,omitempty"`

	// Experiment: The experiment used to generate this read group.
	Experiment *Experiment `json:"experiment,omitempty"`

	// Id: The generated unique read group ID. Note: This is different than
	// the @RG ID field in the SAM spec. For that value, see the `name`
	// field.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read group information. This must be of the
	// form map (string key mapping to a list of string values).
	Info *ReadGroupInfo `json:"info,omitempty"`

	// Name: The read group name. This corresponds to the @RG ID field in
	// the SAM spec.
	Name string `json:"name,omitempty"`

	// PredictedInsertSize: The predicted insert size of this read group.
	// The insert size is the length the sequenced DNA fragment from
	// end-to-end, not including the adapters.
	PredictedInsertSize int64 `json:"predictedInsertSize,omitempty"`

	// Programs: The programs used to generate this read group. Programs are
	// always identical for all read groups within a read group set. For
	// this reason, only the first read group in a returned set will have
	// this field populated.
	Programs []*Program `json:"programs,omitempty"`

	// ReferenceSetId: The reference set the reads in this read group are
	// aligned to. Required if there are any read alignments.
	ReferenceSetId string `json:"referenceSetId,omitempty"`

	// SampleId: The sample this read group's data was generated from. Note:
	// This is not an actual ID within this repository, but rather an
	// identifier for a sample which may be meaningful to some external
	// system.
	SampleId string `json:"sampleId,omitempty"`
}

type ReadGroupInfo struct {
}

type ReadGroupSet struct {
	// DatasetId: The dataset ID.
	DatasetId string `json:"datasetId,omitempty"`

	// Filename: The filename of the original source file for this read
	// group set, if any.
	Filename string `json:"filename,omitempty"`

	// Id: The read group set ID.
	Id string `json:"id,omitempty"`

	// Info: A map of additional read group set information.
	Info *ReadGroupSetInfo `json:"info,omitempty"`

	// Name: The read group set name. By default this will be initialized to
	// the sample name of the sequenced data contained in this set.
	Name string `json:"name,omitempty"`

	// ReadGroups: The read groups in this set. There are typically 1-10
	// read groups in a read group set.
	ReadGroups []*ReadGroup `json:"readGroups,omitempty"`

	// ReferenceSetId: The reference set the reads in this read group set
	// are aligned to.
	ReferenceSetId string `json:"referenceSetId,omitempty"`
}

type ReadGroupSetInfo struct {
}

type Reference struct {
	// Id: The Google generated immutable ID of the reference.
	Id string `json:"id,omitempty"`

	// Length: The length of this reference's sequence.
	Length int64 `json:"length,omitempty,string"`

	// Md5checksum: MD5 of the upper-case sequence excluding all whitespace
	// characters (this is equivalent to SQ:M5 in SAM). This value is
	// represented in lower case hexadecimal format.
	Md5checksum string `json:"md5checksum,omitempty"`

	// Name: The name of this reference, for example `22`.
	Name string `json:"name,omitempty"`

	// NcbiTaxonId: ID from http://www.ncbi.nlm.nih.gov/taxonomy (e.g.
	// 9606->human) if not specified by the containing reference set.
	NcbiTaxonId int64 `json:"ncbiTaxonId,omitempty"`

	// SourceAccessions: All known corresponding accession IDs in INSDC
	// (GenBank/ENA/DDBJ) ideally with a version number, for example
	// `GCF_000001405.26`.
	SourceAccessions []string `json:"sourceAccessions,omitempty"`

	// SourceUri: The URI from which the sequence was obtained. Specifies a
	// FASTA format file/string with one name, sequence pair.
	SourceUri string `json:"sourceUri,omitempty"`
}

type ReferenceBound struct {
	// ReferenceName: The reference the bound is associate with.
	ReferenceName string `json:"referenceName,omitempty"`

	// UpperBound: An upper bound (inclusive) on the starting coordinate of
	// any variant in the reference sequence.
	UpperBound int64 `json:"upperBound,omitempty,string"`
}

type ReferenceSet struct {
	// AssemblyId: Public id of this reference set, such as `GRCh37`.
	AssemblyId string `json:"assemblyId,omitempty"`

	// Description: Free text description of this reference set.
	Description string `json:"description,omitempty"`

	// Id: The Google generated immutable ID of the reference set.
	Id string `json:"id,omitempty"`

	// Md5checksum: Order-independent MD5 checksum which identifies this
	// reference set. The checksum is computed by sorting all lower case
	// hexidecimal string `reference.md5checksum` (for all reference in this
	// set) in ascending lexicographic order, concatenating, and taking the
	// MD5 of that value. The resulting value is represented in lower case
	// hexadecimal format.
	Md5checksum string `json:"md5checksum,omitempty"`

	// NcbiTaxonId: ID from http://www.ncbi.nlm.nih.gov/taxonomy (e.g.
	// 9606->human) indicating the species which this assembly is intended
	// to model. Note that contained references may specify a different
	// `ncbiTaxonId`, as assemblies may contain reference sequences which do
	// not belong to the modeled species, e.g. EBV in a human reference
	// genome.
	NcbiTaxonId int64 `json:"ncbiTaxonId,omitempty"`

	// ReferenceIds: The IDs of the reference objects that are part of this
	// set. `Reference.md5checksum` must be unique within this set.
	ReferenceIds []string `json:"referenceIds,omitempty"`

	// SourceAccessions: All known corresponding accession IDs in INSDC
	// (GenBank/ENA/DDBJ) ideally with a version number, for example
	// `NC_000001.11`.
	SourceAccessions []string `json:"sourceAccessions,omitempty"`

	// SourceUri: The URI from which the references were obtained.
	SourceUri string `json:"sourceUri,omitempty"`
}

type SearchCallSetsRequest struct {
	// Name: Only return call sets for which a substring of the name matches
	// this string.
	Name string `json:"name,omitempty"`

	// PageSize: The maximum number of call sets to return. If unspecified,
	// defaults to 1000.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// VariantSetIds: Restrict the query to call sets within the given
	// variant sets. At least one ID must be provided.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type SearchCallSetsResponse struct {
	// CallSets: The list of matching call sets.
	CallSets []*CallSet `json:"callSets,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReadGroupSetsRequest struct {
	// DatasetIds: Restricts this query to read group sets within the given
	// datasets. At least one ID must be provided.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// Name: Only return read group sets for which a substring of the name
	// matches this string.
	Name string `json:"name,omitempty"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 256. The maximum value is 1024.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchReadGroupSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ReadGroupSets: The list of matching read group sets.
	ReadGroupSets []*ReadGroupSet `json:"readGroupSets,omitempty"`
}

type SearchReadsRequest struct {
	// End: The end position of the range on the reference, 0-based
	// exclusive. If specified, `referenceName` must also be specified.
	End int64 `json:"end,omitempty,string"`

	// PageSize: Specifies number of results to return in a single page. If
	// unspecified, it will default to 256. The maximum value is 2048.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReadGroupIds: The IDs of the read groups within which to search for
	// reads. All specified read groups must belong to the same read group
	// sets. Must specify one of `readGroupSetIds` or `readGroupIds`.
	ReadGroupIds []string `json:"readGroupIds,omitempty"`

	// ReadGroupSetIds: The IDs of the read groups sets within which to
	// search for reads. All specified read group sets must be aligned
	// against a common set of reference sequences; this defines the genomic
	// coordinates for the query. Must specify one of `readGroupSetIds` or
	// `readGroupIds`.
	ReadGroupSetIds []string `json:"readGroupSetIds,omitempty"`

	// ReferenceName: The reference sequence name, for example `chr1`, `1`,
	// or `chrX`. If set to *, only unmapped reads are returned.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The start position of the range on the reference, 0-based
	// inclusive. If specified, `referenceName` must also be specified.
	Start int64 `json:"start,omitempty,string"`
}

type SearchReadsResponse struct {
	// Alignments: The list of matching alignments sorted by mapped genomic
	// coordinate, if any, ascending in position within the same reference.
	// Unmapped reads, which have no position, are returned last and are
	// further sorted in ascending lexicographic order by fragment name.
	Alignments []*Read `json:"alignments,omitempty"`

	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type SearchReferenceSetsRequest struct {
	// Accessions: If present, return references for which the accession
	// matches any of these strings. Best to give a version number, for
	// example `GCF_000001405.26`. If only the main accession number is
	// given then all records with that main accession will be returned,
	// whichever version. Note that different versions will have different
	// sequences.
	Accessions []string `json:"accessions,omitempty"`

	// AssemblyId: If present, return reference sets for which a substring
	// of their `assemblyId` matches this string (case insensitive).
	AssemblyId string `json:"assemblyId,omitempty"`

	// Md5checksums: If present, return references for which the
	// `md5checksum` matches. See `ReferenceSet.md5checksum` for details.
	Md5checksums []string `json:"md5checksums,omitempty"`

	// PageSize: Specifies the maximum number of results to return in a
	// single page.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchReferenceSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ReferenceSets: The matching references sets.
	ReferenceSets []*ReferenceSet `json:"referenceSets,omitempty"`
}

type SearchReferencesRequest struct {
	// Accessions: If present, return references for which the accession
	// matches this string. Best to give a version number, for example
	// `GCF_000001405.26`. If only the main accession number is given then
	// all records with that main accession will be returned, whichever
	// version. Note that different versions will have different sequences.
	Accessions []string `json:"accessions,omitempty"`

	// Md5checksums: If present, return references for which the
	// `md5checksum` matches. See `Reference.md5checksum` for construction
	// details.
	Md5checksums []string `json:"md5checksums,omitempty"`

	// PageSize: Specifies the maximum number of results to return in a
	// single page.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReferenceSetId: If present, return only references which belong to
	// this reference set.
	ReferenceSetId string `json:"referenceSetId,omitempty"`
}

type SearchReferencesResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// References: The matching references.
	References []*Reference `json:"references,omitempty"`
}

type SearchVariantSetsRequest struct {
	// DatasetIds: Exactly one dataset ID must be provided here. Only
	// variant sets which belong to this dataset will be returned.
	DatasetIds []string `json:"datasetIds,omitempty"`

	// PageSize: The maximum number of variant sets to return in a request.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`
}

type SearchVariantSetsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// VariantSets: The variant sets belonging to the requested dataset.
	VariantSets []*VariantSet `json:"variantSets,omitempty"`
}

type SearchVariantsRequest struct {
	// CallSetIds: Only return variant calls which belong to call sets with
	// these ids. Leaving this blank returns all variant calls. If a variant
	// has no calls belonging to any of these call sets, it won't be
	// returned at all. Currently, variants with no calls from any call set
	// will never be returned.
	CallSetIds []string `json:"callSetIds,omitempty"`

	// End: The end of the window, 0-based exclusive. If unspecified or 0,
	// defaults to the length of the reference.
	End int64 `json:"end,omitempty,string"`

	// MaxCalls: The maximum number of calls to return. However, at least
	// one variant will always be returned, even if it has more calls than
	// this limit. If unspecified, defaults to 5000.
	MaxCalls int64 `json:"maxCalls,omitempty"`

	// PageSize: The maximum number of variants to return. If unspecified,
	// defaults to 5000.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The continuation token, which is used to page through
	// large result sets. To get the next page of results, set this
	// parameter to the value of `nextPageToken` from the previous response.
	PageToken string `json:"pageToken,omitempty"`

	// ReferenceName: Required. Only return variants in this reference
	// sequence.
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The beginning of the window (0-based, inclusive) for which
	// overlapping variants should be returned. If unspecified, defaults to
	// 0.
	Start int64 `json:"start,omitempty,string"`

	// VariantName: Only return variants which have exactly this name.
	VariantName string `json:"variantName,omitempty"`

	// VariantSetIds: At most one variant set ID must be provided. Only
	// variants from this variant set will be returned. If omitted, a call
	// set id must be included in the request.
	VariantSetIds []string `json:"variantSetIds,omitempty"`
}

type SearchVariantsResponse struct {
	// NextPageToken: The continuation token, which is used to page through
	// large result sets. Provide this value in a subsequent request to
	// return the next page of results. This field will be empty if there
	// aren't any additional results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Variants: The list of matching Variants.
	Variants []*Variant `json:"variants,omitempty"`
}

type Status struct {
	// Code: The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the [google.rpc.Status.details][google.rpc.Status.details] field,
	// or localized by the client.
	Message string `json:"message,omitempty"`
}

type StatusDetails interface{}

type UndeleteDatasetRequest struct {
}

type Variant struct {
	// AlternateBases: The bases that appear instead of the reference bases.
	AlternateBases []string `json:"alternateBases,omitempty"`

	// Calls: The variant calls for this particular variant. Each one
	// represents the determination of genotype with respect to this
	// variant.
	Calls []*VariantCall `json:"calls,omitempty"`

	// Created: The date this variant was created, in milliseconds from the
	// epoch.
	Created int64 `json:"created,omitempty,string"`

	// End: The end position (0-based) of this variant. This corresponds to
	// the first base after the last base in the reference allele. So, the
	// length of the reference allele is (end - start). This is useful for
	// variants that don't explicitly give alternate bases, for example
	// large deletions.
	End int64 `json:"end,omitempty,string"`

	// Filter: A list of filters (normally quality filters) this variant has
	// failed. `PASS` indicates this variant has passed all filters.
	Filter []string `json:"filter,omitempty"`

	// Id: The Google generated ID of the variant, immutable.
	Id string `json:"id,omitempty"`

	// Info: A map of additional variant information. This must be of the
	// form map (string key mapping to a list of string values).
	Info *VariantInfo `json:"info,omitempty"`

	// Names: Names for the variant, for example a RefSNP ID.
	Names []string `json:"names,omitempty"`

	// Quality: A measure of how likely this variant is to be real. A higher
	// value is better.
	Quality float64 `json:"quality,omitempty"`

	// ReferenceBases: The reference bases for this variant. They start at
	// the given position.
	ReferenceBases string `json:"referenceBases,omitempty"`

	// ReferenceName: The reference on which this variant occurs. (such as
	// `chr20` or `X`)
	ReferenceName string `json:"referenceName,omitempty"`

	// Start: The position at which this variant occurs (0-based). This
	// corresponds to the first base of the string of reference bases.
	Start int64 `json:"start,omitempty,string"`

	// VariantSetId: The ID of the variant set this variant belongs to.
	VariantSetId string `json:"variantSetId,omitempty"`
}

type VariantInfo struct {
}

type VariantCall struct {
	// CallSetId: The ID of the call set this variant call belongs to.
	CallSetId string `json:"callSetId,omitempty"`

	// CallSetName: The name of the call set this variant call belongs to.
	CallSetName string `json:"callSetName,omitempty"`

	// Genotype: The genotype of this variant call. Each value represents
	// either the value of the `referenceBases` field or a 1-based index
	// into `alternateBases`. If a variant had a `referenceBases` value of
	// `T` and an `alternateBases` value of `["A", "C"]`, and the `genotype`
	// was `[2, 1]`, that would mean the call represented the heterozygous
	// value `CA` for this variant. If the `genotype` was instead `[0, 1]`,
	// the represented value would be `TA`. Ordering of the genotype values
	// is important if the `phaseset` is present. If a genotype is not
	// called (that is, a `.` is present in the GT string) -1 is returned.
	Genotype []int64 `json:"genotype,omitempty"`

	// GenotypeLikelihood: The genotype likelihoods for this variant call.
	// Each array entry represents how likely a specific genotype is for
	// this call. The value ordering is defined by the GL tag in the VCF
	// spec. If Phred-scaled genotype likelihood scores (PL) are available
	// and log10(P) genotype likelihood scores (GL) are not, PL scores are
	// converted to GL scores. If both are available, PL scores are stored
	// in `info`.
	GenotypeLikelihood []float64 `json:"genotypeLikelihood,omitempty"`

	// Info: A map of additional variant call information. This must be of
	// the form map (string key mapping to a list of string values).
	Info *VariantCallInfo `json:"info,omitempty"`

	// Phaseset: If this field is present, this variant call's genotype
	// ordering implies the phase of the bases and is consistent with any
	// other variant calls in the same reference sequence which have the
	// same phaseset value. When importing data from VCF, if the genotype
	// data was phased but no phase set was specified this field will be set
	// to `*`.
	Phaseset string `json:"phaseset,omitempty"`
}

type VariantCallInfo struct {
}

type VariantSet struct {
	// DatasetId: The dataset to which this variant set belongs.
	DatasetId string `json:"datasetId,omitempty"`

	// Id: The Google-generated ID of the variant set.
	Id string `json:"id,omitempty"`

	// Metadata: The metadata associated with this variant set.
	Metadata []*VariantSetMetadata `json:"metadata,omitempty"`

	// ReferenceBounds: A list of all references used by the variants in a
	// variant set with associated coordinate upper bounds for each one.
	ReferenceBounds []*ReferenceBound `json:"referenceBounds,omitempty"`
}

type VariantSetMetadata struct {
	// Description: A textual description of this metadata.
	Description string `json:"description,omitempty"`

	// Id: User-provided ID field, not enforced by this API. Two or more
	// pieces of structured metadata with identical id and key fields are
	// considered equivalent.
	Id string `json:"id,omitempty"`

	// Info: Remaining structured metadata key-value pairs. This must be of
	// the form map (string key mapping to a list of string values).
	Info *VariantSetMetadataInfo `json:"info,omitempty"`

	// Key: The top-level key.
	Key string `json:"key,omitempty"`

	// Number: The number of values that can be included in a field
	// described by this metadata.
	Number int64 `json:"number,omitempty"`

	// Type: The type of data. Possible types include: Integer, Float, Flag,
	// Character, and String.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED"
	//   "INTEGER"
	//   "FLOAT"
	//   "FLAG"
	//   "CHARACTER"
	//   "STRING"
	Type string `json:"type,omitempty"`

	// Value: The value field for simple metadata
	Value string `json:"value,omitempty"`
}

type VariantSetMetadataInfo struct {
}

// method id "genomics.callsets.create":

type CallsetsCreateCall struct {
	s       *Service
	callset *CallSet
	opt_    map[string]interface{}
}

// Create: Creates a new call set.
func (r *CallsetsService) Create(callset *CallSet) *CallsetsCreateCall {
	c := &CallsetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.callset = callset
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsCreateCall) Fields(s ...googleapi.Field) *CallsetsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CallsetsCreateCall) Do() (*CallSet, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/callsets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CallSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new call set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.create",
	//   "path": "v1/callsets",
	//   "request": {
	//     "$ref": "CallSet"
	//   },
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.delete":

type CallsetsDeleteCall struct {
	s         *Service
	callSetId string
	opt_      map[string]interface{}
}

// Delete: Deletes a call set.
func (r *CallsetsService) Delete(callSetId string) *CallsetsDeleteCall {
	c := &CallsetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.callSetId = callSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsDeleteCall) Fields(s ...googleapi.Field) *CallsetsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CallsetsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/callsets/{callSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"callSetId": c.callSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a call set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.callsets.delete",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/callsets/{callSetId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.get":

type CallsetsGetCall struct {
	s         *Service
	callSetId string
	opt_      map[string]interface{}
}

// Get: Gets a call set by ID.
func (r *CallsetsService) Get(callSetId string) *CallsetsGetCall {
	c := &CallsetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.callSetId = callSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsGetCall) Fields(s ...googleapi.Field) *CallsetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CallsetsGetCall) Do() (*CallSet, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/callsets/{callSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"callSetId": c.callSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CallSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a call set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.callsets.get",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/callsets/{callSetId}",
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.callsets.patch":

type CallsetsPatchCall struct {
	s         *Service
	callSetId string
	callset   *CallSet
	opt_      map[string]interface{}
}

// Patch: Updates a call set. This method supports patch semantics.
func (r *CallsetsService) Patch(callSetId string, callset *CallSet) *CallsetsPatchCall {
	c := &CallsetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.callSetId = callSetId
	c.callset = callset
	return c
}

// UpdateMask sets the optional parameter "updateMask": An optional mask
// specifying which fields to update. At this time, the only mutable
// field is [name][google.genomics.v1.CallSet.name]. The only acceptable
// value is "name". If unspecified, all mutable fields will be updated.
func (c *CallsetsPatchCall) UpdateMask(updateMask string) *CallsetsPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsPatchCall) Fields(s ...googleapi.Field) *CallsetsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CallsetsPatchCall) Do() (*CallSet, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/callsets/{callSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"callSetId": c.callSetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *CallSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a call set. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.callsets.patch",
	//   "parameterOrder": [
	//     "callSetId"
	//   ],
	//   "parameters": {
	//     "callSetId": {
	//       "description": "The ID of the call set to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "An optional mask specifying which fields to update. At this time, the only mutable field is [name][google.genomics.v1.CallSet.name]. The only acceptable value is \"name\". If unspecified, all mutable fields will be updated.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/callsets/{callSetId}",
	//   "request": {
	//     "$ref": "CallSet"
	//   },
	//   "response": {
	//     "$ref": "CallSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.callsets.search":

type CallsetsSearchCall struct {
	s                     *Service
	searchcallsetsrequest *SearchCallSetsRequest
	opt_                  map[string]interface{}
}

// Search: Gets a list of call sets matching the criteria. Implements
// [GlobalAllianceApi.searchCallSets](http://ga4gh.org/documentation/api/
// v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchCallSets).
func (r *CallsetsService) Search(searchcallsetsrequest *SearchCallSetsRequest) *CallsetsSearchCall {
	c := &CallsetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchcallsetsrequest = searchcallsetsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CallsetsSearchCall) Fields(s ...googleapi.Field) *CallsetsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *CallsetsSearchCall) Do() (*SearchCallSetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchcallsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/callsets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchCallSetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of call sets matching the criteria. Implements [GlobalAllianceApi.searchCallSets](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchCallSets).",
	//   "httpMethod": "POST",
	//   "id": "genomics.callsets.search",
	//   "path": "v1/callsets/search",
	//   "request": {
	//     "$ref": "SearchCallSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchCallSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.datasets.create":

type DatasetsCreateCall struct {
	s       *Service
	dataset *Dataset
	opt_    map[string]interface{}
}

// Create: Creates a new dataset.
func (r *DatasetsService) Create(dataset *Dataset) *DatasetsCreateCall {
	c := &DatasetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.dataset = dataset
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsCreateCall) Fields(s ...googleapi.Field) *DatasetsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsCreateCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.create",
	//   "path": "v1/datasets",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.delete":

type DatasetsDeleteCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Delete: Deletes a dataset.
func (r *DatasetsService) Delete(datasetId string) *DatasetsDeleteCall {
	c := &DatasetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsDeleteCall) Fields(s ...googleapi.Field) *DatasetsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a dataset.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.datasets.delete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.get":

type DatasetsGetCall struct {
	s         *Service
	datasetId string
	opt_      map[string]interface{}
}

// Get: Gets a dataset by ID.
func (r *DatasetsService) Get(datasetId string) *DatasetsGetCall {
	c := &DatasetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsGetCall) Fields(s ...googleapi.Field) *DatasetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsGetCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a dataset by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.get",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/datasets/{datasetId}",
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.datasets.list":

type DatasetsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists datasets within a project.
func (r *DatasetsService) List() *DatasetsListCall {
	c := &DatasetsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results returned by this request. If unspecified, defaults to 50.
// The maximum value is 1024.
func (c *DatasetsListCall) PageSize(pageSize int64) *DatasetsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// `nextPageToken` from the previous response.
func (c *DatasetsListCall) PageToken(pageToken string) *DatasetsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ProjectId sets the optional parameter "projectId": Required. The
// project to list datasets for.
func (c *DatasetsListCall) ProjectId(projectId string) *DatasetsListCall {
	c.opt_["projectId"] = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsListCall) Fields(s ...googleapi.Field) *DatasetsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsListCall) Do() (*ListDatasetsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["projectId"]; ok {
		params.Set("projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListDatasetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists datasets within a project.",
	//   "httpMethod": "GET",
	//   "id": "genomics.datasets.list",
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results returned by this request. If unspecified, defaults to 50. The maximum value is 1024.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `nextPageToken` from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Required. The project to list datasets for.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/datasets",
	//   "response": {
	//     "$ref": "ListDatasetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.datasets.patch":

type DatasetsPatchCall struct {
	s         *Service
	datasetId string
	dataset   *Dataset
	opt_      map[string]interface{}
}

// Patch: Updates a dataset. This method supports patch semantics.
func (r *DatasetsService) Patch(datasetId string, dataset *Dataset) *DatasetsPatchCall {
	c := &DatasetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.dataset = dataset
	return c
}

// UpdateMask sets the optional parameter "updateMask": An optional mask
// specifying which fields to update. At this time, the only mutable
// field is [name][google.genomics.v1.Dataset.name]. The only acceptable
// value is "name". If unspecified, all mutable fields will be updated.
func (c *DatasetsPatchCall) UpdateMask(updateMask string) *DatasetsPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsPatchCall) Fields(s ...googleapi.Field) *DatasetsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsPatchCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.dataset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets/{datasetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a dataset. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.datasets.patch",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "An optional mask specifying which fields to update. At this time, the only mutable field is [name][google.genomics.v1.Dataset.name]. The only acceptable value is \"name\". If unspecified, all mutable fields will be updated.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/datasets/{datasetId}",
	//   "request": {
	//     "$ref": "Dataset"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.datasets.undelete":

type DatasetsUndeleteCall struct {
	s                      *Service
	datasetId              string
	undeletedatasetrequest *UndeleteDatasetRequest
	opt_                   map[string]interface{}
}

// Undelete: Undeletes a dataset by restoring a dataset which was
// deleted via this API. This operation is only possible for a week
// after the deletion occurred.
func (r *DatasetsService) Undelete(datasetId string, undeletedatasetrequest *UndeleteDatasetRequest) *DatasetsUndeleteCall {
	c := &DatasetsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.datasetId = datasetId
	c.undeletedatasetrequest = undeletedatasetrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatasetsUndeleteCall) Fields(s ...googleapi.Field) *DatasetsUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatasetsUndeleteCall) Do() (*Dataset, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.undeletedatasetrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/datasets/{datasetId}:undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"datasetId": c.datasetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Dataset
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Undeletes a dataset by restoring a dataset which was deleted via this API. This operation is only possible for a week after the deletion occurred.",
	//   "httpMethod": "POST",
	//   "id": "genomics.datasets.undelete",
	//   "parameterOrder": [
	//     "datasetId"
	//   ],
	//   "parameters": {
	//     "datasetId": {
	//       "description": "The ID of the dataset to be undeleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/datasets/{datasetId}:undelete",
	//   "request": {
	//     "$ref": "UndeleteDatasetRequest"
	//   },
	//   "response": {
	//     "$ref": "Dataset"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.operations.cancel":

type OperationsCancelCall struct {
	s                      *Service
	name                   string
	canceloperationrequest *CancelOperationRequest
	opt_                   map[string]interface{}
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
// The server makes a best effort to cancel the operation, but success
// is not guaranteed. Clients may use
// [Operations.GetOperation][google.longrunning.Operations.GetOperation]
// or
// [Operations.ListOperations][google.longrunning.Operations.ListOperatio
// ns] to check whether the cancellation succeeded or the operation
// completed despite cancellation.
func (r *OperationsService) Cancel(name string, canceloperationrequest *CancelOperationRequest) *OperationsCancelCall {
	c := &OperationsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.canceloperationrequest = canceloperationrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsCancelCall) Fields(s ...googleapi.Field) *OperationsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsCancelCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceloperationrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation. The server makes a best effort to cancel the operation, but success is not guaranteed. Clients may use [Operations.GetOperation][google.longrunning.Operations.GetOperation] or [Operations.ListOperations][google.longrunning.Operations.ListOperations] to check whether the cancellation succeeded or the operation completed despite cancellation.",
	//   "httpMethod": "POST",
	//   "id": "genomics.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "request": {
	//     "$ref": "CancelOperationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.operations.delete":

type OperationsDeleteCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Delete: This method is not implemented. To cancel an operation,
// please use
// [Operations.CancelOperation][google.longrunning.Operations.CancelOpera
// tion].
func (r *OperationsService) Delete(name string) *OperationsDeleteCall {
	c := &OperationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsDeleteCall) Fields(s ...googleapi.Field) *OperationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "This method is not implemented. To cancel an operation, please use [Operations.CancelOperation][google.longrunning.Operations.CancelOperation].",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.operations.get":

type OperationsGetCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Gets the latest state of a long-running operation. Clients may
// use this method to poll the operation result at intervals as
// recommended by the API service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the latest state of a long-running operation. Clients may use this method to poll the operation result at intervals as recommended by the API service.",
	//   "httpMethod": "GET",
	//   "id": "genomics.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.operations.list":

type OperationsListCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// List: Lists operations that match the specified filter in the
// request.
func (r *OperationsService) List(name string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": A string for filtering
// [Operations][google.longrunning.Operation]. The following filter
// fields are supported: * projectId: Required. Corresponds to
// [OperationMetadata.projectId][google.genomics.v1.OperationMetadata.pro
// ject_id]. * createTime: The time this job was created, in seconds
// from the [epoch](http://en.wikipedia.org/wiki/Unix_time). Can use
// `>=` and/or `= 1432140000` * `projectId = my-project AND createTime
// >= 1432140000 AND createTime <= 1432150000 AND status = RUNNING`
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return. If unspecified, defaults to 256. The maximum
// value is 2048.
func (c *OperationsListCall) PageSize(pageSize int64) *OperationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsListCall) Do() (*ListOperationsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListOperationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request.",
	//   "httpMethod": "GET",
	//   "id": "genomics.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "A string for filtering [Operations][google.longrunning.Operation]. The following filter fields are supported: * projectId: Required. Corresponds to [OperationMetadata.projectId][google.genomics.v1.OperationMetadata.project_id]. * createTime: The time this job was created, in seconds from the [epoch](http://en.wikipedia.org/wiki/Unix_time). Can use `\u003e=` and/or `= 1432140000` * `projectId = my-project AND createTime \u003e= 1432140000 AND createTime \u003c= 1432150000 AND status = RUNNING`",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation collection.",
	//       "location": "path",
	//       "pattern": "^operations$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return. If unspecified, defaults to 256. The maximum value is 2048.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard List page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.delete":

type ReadgroupsetsDeleteCall struct {
	s              *Service
	readGroupSetId string
	opt_           map[string]interface{}
}

// Delete: Deletes a read group set.
func (r *ReadgroupsetsService) Delete(readGroupSetId string) *ReadgroupsetsDeleteCall {
	c := &ReadgroupsetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.readGroupSetId = readGroupSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsDeleteCall) Fields(s ...googleapi.Field) *ReadgroupsetsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/{readGroupSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a read group set.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.readgroupsets.delete",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set to be deleted. The caller must have WRITE permissions to the dataset associated with this read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/readgroupsets/{readGroupSetId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.export":

type ReadgroupsetsExportCall struct {
	s                         *Service
	readGroupSetId            string
	exportreadgroupsetrequest *ExportReadGroupSetRequest
	opt_                      map[string]interface{}
}

// Export: Exports a read group set to a BAM file in Google Cloud
// Storage. Note that currently there may be some differences between
// exported BAM files and the original BAM file at the time of import.
// In particular, comments in the input file header will not be
// preserved, some custom tags will be converted to strings, and
// original reference sequence order is not necessarily preserved.
func (r *ReadgroupsetsService) Export(readGroupSetId string, exportreadgroupsetrequest *ExportReadGroupSetRequest) *ReadgroupsetsExportCall {
	c := &ReadgroupsetsExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.readGroupSetId = readGroupSetId
	c.exportreadgroupsetrequest = exportreadgroupsetrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsExportCall) Fields(s ...googleapi.Field) *ReadgroupsetsExportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsExportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportreadgroupsetrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/{readGroupSetId}:export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports a read group set to a BAM file in Google Cloud Storage. Note that currently there may be some differences between exported BAM files and the original BAM file at the time of import. In particular, comments in the input file header will not be preserved, some custom tags will be converted to strings, and original reference sequence order is not necessarily preserved.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.export",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "Required. The ID of the read group set to export.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/readgroupsets/{readGroupSetId}:export",
	//   "request": {
	//     "$ref": "ExportReadGroupSetRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.get":

type ReadgroupsetsGetCall struct {
	s              *Service
	readGroupSetId string
	opt_           map[string]interface{}
}

// Get: Gets a read group set by ID.
func (r *ReadgroupsetsService) Get(readGroupSetId string) *ReadgroupsetsGetCall {
	c := &ReadgroupsetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.readGroupSetId = readGroupSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsGetCall) Fields(s ...googleapi.Field) *ReadgroupsetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsGetCall) Do() (*ReadGroupSet, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/{readGroupSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ReadGroupSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a read group set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readgroupsets.get",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/readgroupsets/{readGroupSetId}",
	//   "response": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.readgroupsets.import":

type ReadgroupsetsImportCall struct {
	s                          *Service
	importreadgroupsetsrequest *ImportReadGroupSetsRequest
	opt_                       map[string]interface{}
}

// Import: Creates read group sets by asynchronously importing the
// provided information. Note that currently comments in the input file
// header are **not** imported and some custom tags will be converted to
// strings, rather than preserving tag types. The caller must have WRITE
// permissions to the dataset.
func (r *ReadgroupsetsService) Import(importreadgroupsetsrequest *ImportReadGroupSetsRequest) *ReadgroupsetsImportCall {
	c := &ReadgroupsetsImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.importreadgroupsetsrequest = importreadgroupsetsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsImportCall) Fields(s ...googleapi.Field) *ReadgroupsetsImportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsImportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importreadgroupsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets:import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates read group sets by asynchronously importing the provided information. Note that currently comments in the input file header are **not** imported and some custom tags will be converted to strings, rather than preserving tag types. The caller must have WRITE permissions to the dataset.",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.import",
	//   "path": "v1/readgroupsets:import",
	//   "request": {
	//     "$ref": "ImportReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.patch":

type ReadgroupsetsPatchCall struct {
	s              *Service
	readGroupSetId string
	readgroupset   *ReadGroupSet
	opt_           map[string]interface{}
}

// Patch: Updates a read group set. This method supports patch
// semantics.
func (r *ReadgroupsetsService) Patch(readGroupSetId string, readgroupset *ReadGroupSet) *ReadgroupsetsPatchCall {
	c := &ReadgroupsetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.readGroupSetId = readGroupSetId
	c.readgroupset = readgroupset
	return c
}

// UpdateMask sets the optional parameter "updateMask": An optional mask
// specifying which fields to update. At this time, mutable fields are
// [referenceSetId][google.genomics.v1.ReadGroupSet.reference_set_id]
// and [name][google.genomics.v1.ReadGroupSet.name]. Acceptable values
// are "referenceSetId" and "name". If unspecified, all mutable fields
// will be updated.
func (c *ReadgroupsetsPatchCall) UpdateMask(updateMask string) *ReadgroupsetsPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsPatchCall) Fields(s ...googleapi.Field) *ReadgroupsetsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsPatchCall) Do() (*ReadGroupSet, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.readgroupset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/{readGroupSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ReadGroupSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a read group set. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.readgroupsets.patch",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "readGroupSetId": {
	//       "description": "The ID of the read group set to be updated. The caller must have WRITE permissions to the dataset associated with this read group set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "An optional mask specifying which fields to update. At this time, mutable fields are [referenceSetId][google.genomics.v1.ReadGroupSet.reference_set_id] and [name][google.genomics.v1.ReadGroupSet.name]. Acceptable values are \"referenceSetId\" and \"name\". If unspecified, all mutable fields will be updated.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/readgroupsets/{readGroupSetId}",
	//   "request": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "response": {
	//     "$ref": "ReadGroupSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.readgroupsets.search":

type ReadgroupsetsSearchCall struct {
	s                          *Service
	searchreadgroupsetsrequest *SearchReadGroupSetsRequest
	opt_                       map[string]interface{}
}

// Search: Searches for read group sets matching the criteria.
// Implements
// [GlobalAllianceApi.searchReadGroupSets](http://ga4gh.org/documentation
// /api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReadGroupSets).
func (r *ReadgroupsetsService) Search(searchreadgroupsetsrequest *SearchReadGroupSetsRequest) *ReadgroupsetsSearchCall {
	c := &ReadgroupsetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreadgroupsetsrequest = searchreadgroupsetsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsSearchCall) Fields(s ...googleapi.Field) *ReadgroupsetsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsSearchCall) Do() (*SearchReadGroupSetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreadgroupsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReadGroupSetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for read group sets matching the criteria. Implements [GlobalAllianceApi.searchReadGroupSets](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReadGroupSets).",
	//   "httpMethod": "POST",
	//   "id": "genomics.readgroupsets.search",
	//   "path": "v1/readgroupsets/search",
	//   "request": {
	//     "$ref": "SearchReadGroupSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadGroupSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.readgroupsets.coveragebuckets.list":

type ReadgroupsetsCoveragebucketsListCall struct {
	s              *Service
	readGroupSetId string
	opt_           map[string]interface{}
}

// List: Lists fixed width coverage buckets for a read group set, each
// of which correspond to a range of a reference sequence. Each bucket
// summarizes coverage information across its corresponding genomic
// range. Coverage is defined as the number of reads which are aligned
// to a given base in the reference sequence. Coverage buckets are
// available at several precomputed bucket widths, enabling retrieval of
// various coverage 'zoom levels'. The caller must have READ permissions
// for the target read group set.
func (r *ReadgroupsetsCoveragebucketsService) List(readGroupSetId string) *ReadgroupsetsCoveragebucketsListCall {
	c := &ReadgroupsetsCoveragebucketsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.readGroupSetId = readGroupSetId
	return c
}

// End sets the optional parameter "end": The end position of the range
// on the reference, 0-based exclusive. If specified, `referenceName`
// must also be specified. If unset or 0, defaults to the length of the
// reference.
func (c *ReadgroupsetsCoveragebucketsListCall) End(end int64) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["end"] = end
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return in a single page. If unspecified, defaults to
// 1024. The maximum value is 2048.
func (c *ReadgroupsetsCoveragebucketsListCall) PageSize(pageSize int64) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// `nextPageToken` from the previous response.
func (c *ReadgroupsetsCoveragebucketsListCall) PageToken(pageToken string) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ReferenceName sets the optional parameter "referenceName": The name
// of the reference to query, within the reference set associated with
// this query.
func (c *ReadgroupsetsCoveragebucketsListCall) ReferenceName(referenceName string) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["referenceName"] = referenceName
	return c
}

// Start sets the optional parameter "start": The start position of the
// range on the reference, 0-based inclusive. If specified,
// `referenceName` must also be specified. Defaults to 0.
func (c *ReadgroupsetsCoveragebucketsListCall) Start(start int64) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["start"] = start
	return c
}

// TargetBucketWidth sets the optional parameter "targetBucketWidth":
// The desired width of each reported coverage bucket in base pairs.
// This will be rounded down to the nearest precomputed bucket width;
// the value of which is returned as `bucketWidth` in the response.
// Defaults to infinity (each bucket spans an entire reference sequence)
// or the length of the target range, if specified. The smallest
// precomputed `bucketWidth` is currently 2048 base pairs; this is
// subject to change.
func (c *ReadgroupsetsCoveragebucketsListCall) TargetBucketWidth(targetBucketWidth int64) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["targetBucketWidth"] = targetBucketWidth
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadgroupsetsCoveragebucketsListCall) Fields(s ...googleapi.Field) *ReadgroupsetsCoveragebucketsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadgroupsetsCoveragebucketsListCall) Do() (*ListCoverageBucketsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["end"]; ok {
		params.Set("end", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["referenceName"]; ok {
		params.Set("referenceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start"]; ok {
		params.Set("start", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["targetBucketWidth"]; ok {
		params.Set("targetBucketWidth", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/readgroupsets/{readGroupSetId}/coveragebuckets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"readGroupSetId": c.readGroupSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListCoverageBucketsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists fixed width coverage buckets for a read group set, each of which correspond to a range of a reference sequence. Each bucket summarizes coverage information across its corresponding genomic range. Coverage is defined as the number of reads which are aligned to a given base in the reference sequence. Coverage buckets are available at several precomputed bucket widths, enabling retrieval of various coverage 'zoom levels'. The caller must have READ permissions for the target read group set.",
	//   "httpMethod": "GET",
	//   "id": "genomics.readgroupsets.coveragebuckets.list",
	//   "parameterOrder": [
	//     "readGroupSetId"
	//   ],
	//   "parameters": {
	//     "end": {
	//       "description": "The end position of the range on the reference, 0-based exclusive. If specified, `referenceName` must also be specified. If unset or 0, defaults to the length of the reference.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of results to return in a single page. If unspecified, defaults to 1024. The maximum value is 2048.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `nextPageToken` from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "readGroupSetId": {
	//       "description": "Required. The ID of the read group set over which coverage is requested.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "referenceName": {
	//       "description": "The name of the reference to query, within the reference set associated with this query. Optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "start": {
	//       "description": "The start position of the range on the reference, 0-based inclusive. If specified, `referenceName` must also be specified. Defaults to 0.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "targetBucketWidth": {
	//       "description": "The desired width of each reported coverage bucket in base pairs. This will be rounded down to the nearest precomputed bucket width; the value of which is returned as `bucketWidth` in the response. Defaults to infinity (each bucket spans an entire reference sequence) or the length of the target range, if specified. The smallest precomputed `bucketWidth` is currently 2048 base pairs; this is subject to change.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/readgroupsets/{readGroupSetId}/coveragebuckets",
	//   "response": {
	//     "$ref": "ListCoverageBucketsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.reads.search":

type ReadsSearchCall struct {
	s                  *Service
	searchreadsrequest *SearchReadsRequest
	opt_               map[string]interface{}
}

// Search: Gets a list of reads for one or more read group sets. Reads
// search operates over a genomic coordinate space of reference sequence
// & position defined over the reference sequences to which the
// requested read group sets are aligned. If a target positional range
// is specified, search returns all reads whose alignment to the
// reference genome overlap the range. A query which specifies only read
// group set IDs yields all reads in those read group sets, including
// unmapped reads. All reads returned (including reads on subsequent
// pages) are ordered by genomic coordinate (reference sequence &
// position). Reads with equivalent genomic coordinates are returned in
// a deterministic order. Implements
// [GlobalAllianceApi.searchReads](http://ga4gh.org/documentation/api/v0.
// 5.1/ga4gh_api.html#/schema/org.ga4gh.searchReads).
func (r *ReadsService) Search(searchreadsrequest *SearchReadsRequest) *ReadsSearchCall {
	c := &ReadsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreadsrequest = searchreadsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReadsSearchCall) Fields(s ...googleapi.Field) *ReadsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReadsSearchCall) Do() (*SearchReadsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreadsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/reads/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReadsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of reads for one or more read group sets. Reads search operates over a genomic coordinate space of reference sequence \u0026 position defined over the reference sequences to which the requested read group sets are aligned. If a target positional range is specified, search returns all reads whose alignment to the reference genome overlap the range. A query which specifies only read group set IDs yields all reads in those read group sets, including unmapped reads. All reads returned (including reads on subsequent pages) are ordered by genomic coordinate (reference sequence \u0026 position). Reads with equivalent genomic coordinates are returned in a deterministic order. Implements [GlobalAllianceApi.searchReads](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReads).",
	//   "httpMethod": "POST",
	//   "id": "genomics.reads.search",
	//   "path": "v1/reads/search",
	//   "request": {
	//     "$ref": "SearchReadsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReadsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.get":

type ReferencesGetCall struct {
	s           *Service
	referenceId string
	opt_        map[string]interface{}
}

// Get: Gets a reference. Implements
// [GlobalAllianceApi.getReference](http://ga4gh.org/documentation/api/v0
// .5.1/ga4gh_api.html#/schema/org.ga4gh.getReference).
func (r *ReferencesService) Get(referenceId string) *ReferencesGetCall {
	c := &ReferencesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.referenceId = referenceId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesGetCall) Fields(s ...googleapi.Field) *ReferencesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReferencesGetCall) Do() (*Reference, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/references/{referenceId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"referenceId": c.referenceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Reference
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a reference. Implements [GlobalAllianceApi.getReference](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.getReference).",
	//   "httpMethod": "GET",
	//   "id": "genomics.references.get",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "referenceId": {
	//       "description": "The ID of the reference.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/references/{referenceId}",
	//   "response": {
	//     "$ref": "Reference"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.search":

type ReferencesSearchCall struct {
	s                       *Service
	searchreferencesrequest *SearchReferencesRequest
	opt_                    map[string]interface{}
}

// Search: Searches for references which match the given criteria.
// Implements
// [GlobalAllianceApi.searchReferences](http://ga4gh.org/documentation/ap
// i/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReferences).
func (r *ReferencesService) Search(searchreferencesrequest *SearchReferencesRequest) *ReferencesSearchCall {
	c := &ReferencesSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreferencesrequest = searchreferencesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesSearchCall) Fields(s ...googleapi.Field) *ReferencesSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReferencesSearchCall) Do() (*SearchReferencesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreferencesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/references/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReferencesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for references which match the given criteria. Implements [GlobalAllianceApi.searchReferences](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReferences).",
	//   "httpMethod": "POST",
	//   "id": "genomics.references.search",
	//   "path": "v1/references/search",
	//   "request": {
	//     "$ref": "SearchReferencesRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReferencesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.references.bases.list":

type ReferencesBasesListCall struct {
	s           *Service
	referenceId string
	opt_        map[string]interface{}
}

// List: Lists the bases in a reference, optionally restricted to a
// range. Implements
// [GlobalAllianceApi.getReferenceBases](http://ga4gh.org/documentation/a
// pi/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.getReferenceBases).
func (r *ReferencesBasesService) List(referenceId string) *ReferencesBasesListCall {
	c := &ReferencesBasesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.referenceId = referenceId
	return c
}

// End sets the optional parameter "end": The end position (0-based,
// exclusive) of this query. Defaults to the length of this reference.
func (c *ReferencesBasesListCall) End(end int64) *ReferencesBasesListCall {
	c.opt_["end"] = end
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the
// maximum number of bases to return in a single page.
func (c *ReferencesBasesListCall) PageSize(pageSize int64) *ReferencesBasesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The continuation
// token, which is used to page through large result sets. To get the
// next page of results, set this parameter to the value of
// `nextPageToken` from the previous response.
func (c *ReferencesBasesListCall) PageToken(pageToken string) *ReferencesBasesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Start sets the optional parameter "start": The start position
// (0-based) of this query. Defaults to 0.
func (c *ReferencesBasesListCall) Start(start int64) *ReferencesBasesListCall {
	c.opt_["start"] = start
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesBasesListCall) Fields(s ...googleapi.Field) *ReferencesBasesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReferencesBasesListCall) Do() (*ListBasesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["end"]; ok {
		params.Set("end", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["start"]; ok {
		params.Set("start", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/references/{referenceId}/bases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"referenceId": c.referenceId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListBasesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the bases in a reference, optionally restricted to a range. Implements [GlobalAllianceApi.getReferenceBases](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.getReferenceBases).",
	//   "httpMethod": "GET",
	//   "id": "genomics.references.bases.list",
	//   "parameterOrder": [
	//     "referenceId"
	//   ],
	//   "parameters": {
	//     "end": {
	//       "description": "The end position (0-based, exclusive) of this query. Defaults to the length of this reference.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Specifies the maximum number of bases to return in a single page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The continuation token, which is used to page through large result sets. To get the next page of results, set this parameter to the value of `nextPageToken` from the previous response.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "referenceId": {
	//       "description": "The ID of the reference.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "start": {
	//       "description": "The start position (0-based) of this query. Defaults to 0.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/references/{referenceId}/bases",
	//   "response": {
	//     "$ref": "ListBasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.referencesets.get":

type ReferencesetsGetCall struct {
	s              *Service
	referenceSetId string
	opt_           map[string]interface{}
}

// Get: Gets a reference set. Implements
// [GlobalAllianceApi.getReferenceSet](http://ga4gh.org/documentation/api
// /v0.5.1/ga4gh_api.html#/schema/org.ga4gh.getReferenceSet").
func (r *ReferencesetsService) Get(referenceSetId string) *ReferencesetsGetCall {
	c := &ReferencesetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.referenceSetId = referenceSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesetsGetCall) Fields(s ...googleapi.Field) *ReferencesetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReferencesetsGetCall) Do() (*ReferenceSet, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/referencesets/{referenceSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"referenceSetId": c.referenceSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ReferenceSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a reference set. Implements [GlobalAllianceApi.getReferenceSet](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.getReferenceSet\").",
	//   "httpMethod": "GET",
	//   "id": "genomics.referencesets.get",
	//   "parameterOrder": [
	//     "referenceSetId"
	//   ],
	//   "parameters": {
	//     "referenceSetId": {
	//       "description": "The ID of the reference set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/referencesets/{referenceSetId}",
	//   "response": {
	//     "$ref": "ReferenceSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.referencesets.search":

type ReferencesetsSearchCall struct {
	s                          *Service
	searchreferencesetsrequest *SearchReferenceSetsRequest
	opt_                       map[string]interface{}
}

// Search: Searches for reference sets which match the given criteria.
// Implements
// [GlobalAllianceApi.searchReferenceSets](http://ga4gh.org/documentation
// /api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReferenceSets)
func (r *ReferencesetsService) Search(searchreferencesetsrequest *SearchReferenceSetsRequest) *ReferencesetsSearchCall {
	c := &ReferencesetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchreferencesetsrequest = searchreferencesetsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReferencesetsSearchCall) Fields(s ...googleapi.Field) *ReferencesetsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReferencesetsSearchCall) Do() (*SearchReferenceSetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchreferencesetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/referencesets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchReferenceSetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Searches for reference sets which match the given criteria. Implements [GlobalAllianceApi.searchReferenceSets](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchReferenceSets)",
	//   "httpMethod": "POST",
	//   "id": "genomics.referencesets.search",
	//   "path": "v1/referencesets/search",
	//   "request": {
	//     "$ref": "SearchReferenceSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchReferenceSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variants.create":

type VariantsCreateCall struct {
	s       *Service
	variant *Variant
	opt_    map[string]interface{}
}

// Create: Creates a new variant.
func (r *VariantsService) Create(variant *Variant) *VariantsCreateCall {
	c := &VariantsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.variant = variant
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsCreateCall) Fields(s ...googleapi.Field) *VariantsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsCreateCall) Do() (*Variant, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variant)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new variant.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.create",
	//   "path": "v1/variants",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.delete":

type VariantsDeleteCall struct {
	s         *Service
	variantId string
	opt_      map[string]interface{}
}

// Delete: Deletes a variant.
func (r *VariantsService) Delete(variantId string) *VariantsDeleteCall {
	c := &VariantsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsDeleteCall) Fields(s ...googleapi.Field) *VariantsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a variant.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.variants.delete",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variants/{variantId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.get":

type VariantsGetCall struct {
	s         *Service
	variantId string
	opt_      map[string]interface{}
}

// Get: Gets a variant by ID.
func (r *VariantsService) Get(variantId string) *VariantsGetCall {
	c := &VariantsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsGetCall) Fields(s ...googleapi.Field) *VariantsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsGetCall) Do() (*Variant, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a variant by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variants.get",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "variantId": {
	//       "description": "The ID of the variant.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variants/{variantId}",
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variants.import":

type VariantsImportCall struct {
	s                     *Service
	importvariantsrequest *ImportVariantsRequest
	opt_                  map[string]interface{}
}

// Import: Creates variant data by asynchronously importing the provided
// information. The variants for import will be merged with any existing
// data and each other according to the behavior of mergeVariants. In
// particular, this means for merged VCF variants that have conflicting
// INFO fields, some data will be arbitrarily discarded. As a special
// case, for single-sample VCF files, QUAL and FILTER fields will be
// moved to the call level; these are sometimes interpreted in a
// call-specific context. Imported VCF headers are appended to the
// metadata already in a variant set.
func (r *VariantsService) Import(importvariantsrequest *ImportVariantsRequest) *VariantsImportCall {
	c := &VariantsImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.importvariantsrequest = importvariantsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsImportCall) Fields(s ...googleapi.Field) *VariantsImportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsImportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importvariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants:import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates variant data by asynchronously importing the provided information. The variants for import will be merged with any existing data and each other according to the behavior of mergeVariants. In particular, this means for merged VCF variants that have conflicting INFO fields, some data will be arbitrarily discarded. As a special case, for single-sample VCF files, QUAL and FILTER fields will be moved to the call level; these are sometimes interpreted in a call-specific context. Imported VCF headers are appended to the metadata already in a variant set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.import",
	//   "path": "v1/variants:import",
	//   "request": {
	//     "$ref": "ImportVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/devstorage.read_write",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.merge":

type VariantsMergeCall struct {
	s                    *Service
	mergevariantsrequest *MergeVariantsRequest
	opt_                 map[string]interface{}
}

// Merge: Merges the given variants with existing variants. Each variant
// will be merged with an existing variant that matches its reference
// sequence, start, end, reference bases, and alternative bases. If no
// such variant exists, a new one will be created. When variants are
// merged, the call information from the new variant is added to the
// existing variant, and other fields (such as key/value pairs) are
// discarded.
func (r *VariantsService) Merge(mergevariantsrequest *MergeVariantsRequest) *VariantsMergeCall {
	c := &VariantsMergeCall{s: r.s, opt_: make(map[string]interface{})}
	c.mergevariantsrequest = mergevariantsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsMergeCall) Fields(s ...googleapi.Field) *VariantsMergeCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsMergeCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.mergevariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants:merge")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Merges the given variants with existing variants. Each variant will be merged with an existing variant that matches its reference sequence, start, end, reference bases, and alternative bases. If no such variant exists, a new one will be created. When variants are merged, the call information from the new variant is added to the existing variant, and other fields (such as key/value pairs) are discarded.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.merge",
	//   "path": "v1/variants:merge",
	//   "request": {
	//     "$ref": "MergeVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.patch":

type VariantsPatchCall struct {
	s         *Service
	variantId string
	variant   *Variant
	opt_      map[string]interface{}
}

// Patch: Updates a variant. This method supports patch semantics.
// Returns the modified variant without its calls.
func (r *VariantsService) Patch(variantId string, variant *Variant) *VariantsPatchCall {
	c := &VariantsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantId = variantId
	c.variant = variant
	return c
}

// UpdateMask sets the optional parameter "updateMask": An optional mask
// specifying which fields to update. At this time, mutable fields are
// [names][google.genomics.v1.Variant.names] and
// [info][google.genomics.v1.Variant.info]. Acceptable values are
// "names" and "info". If unspecified, all mutable fields will be
// updated.
func (c *VariantsPatchCall) UpdateMask(updateMask string) *VariantsPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsPatchCall) Fields(s ...googleapi.Field) *VariantsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsPatchCall) Do() (*Variant, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variant)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants/{variantId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantId": c.variantId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Variant
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a variant. This method supports patch semantics. Returns the modified variant without its calls.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.variants.patch",
	//   "parameterOrder": [
	//     "variantId"
	//   ],
	//   "parameters": {
	//     "updateMask": {
	//       "description": "An optional mask specifying which fields to update. At this time, mutable fields are [names][google.genomics.v1.Variant.names] and [info][google.genomics.v1.Variant.info]. Acceptable values are \"names\" and \"info\". If unspecified, all mutable fields will be updated.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "variantId": {
	//       "description": "The ID of the variant to be updated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variants/{variantId}",
	//   "request": {
	//     "$ref": "Variant"
	//   },
	//   "response": {
	//     "$ref": "Variant"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variants.search":

type VariantsSearchCall struct {
	s                     *Service
	searchvariantsrequest *SearchVariantsRequest
	opt_                  map[string]interface{}
}

// Search: Gets a list of variants matching the criteria. Implements
// [GlobalAllianceApi.searchVariants](http://ga4gh.org/documentation/api/
// v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchVariants).
func (r *VariantsService) Search(searchvariantsrequest *SearchVariantsRequest) *VariantsSearchCall {
	c := &VariantsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchvariantsrequest = searchvariantsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsSearchCall) Fields(s ...googleapi.Field) *VariantsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsSearchCall) Do() (*SearchVariantsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchvariantsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variants/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchVariantsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a list of variants matching the criteria. Implements [GlobalAllianceApi.searchVariants](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchVariants).",
	//   "httpMethod": "POST",
	//   "id": "genomics.variants.search",
	//   "path": "v1/variants/search",
	//   "request": {
	//     "$ref": "SearchVariantsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchVariantsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variantsets.create":

type VariantsetsCreateCall struct {
	s          *Service
	variantset *VariantSet
	opt_       map[string]interface{}
}

// Create: Creates a new variant set.
func (r *VariantsetsService) Create(variantset *VariantSet) *VariantsetsCreateCall {
	c := &VariantsetsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantset = variantset
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsCreateCall) Fields(s ...googleapi.Field) *VariantsetsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsCreateCall) Do() (*VariantSet, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variantset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *VariantSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new variant set.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.create",
	//   "path": "v1/variantsets",
	//   "request": {
	//     "$ref": "VariantSet"
	//   },
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.delete":

type VariantsetsDeleteCall struct {
	s            *Service
	variantSetId string
	opt_         map[string]interface{}
}

// Delete: Deletes the contents of a variant set. The variant set object
// is not deleted.
func (r *VariantsetsService) Delete(variantSetId string) *VariantsetsDeleteCall {
	c := &VariantsetsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantSetId = variantSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsDeleteCall) Fields(s ...googleapi.Field) *VariantsetsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets/{variantSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantSetId": c.variantSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the contents of a variant set. The variant set object is not deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "genomics.variantsets.delete",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "The ID of the variant set to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variantsets/{variantSetId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.export":

type VariantsetsExportCall struct {
	s                       *Service
	variantSetId            string
	exportvariantsetrequest *ExportVariantSetRequest
	opt_                    map[string]interface{}
}

// Export: Exports variant set data to an external destination.
func (r *VariantsetsService) Export(variantSetId string, exportvariantsetrequest *ExportVariantSetRequest) *VariantsetsExportCall {
	c := &VariantsetsExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantSetId = variantSetId
	c.exportvariantsetrequest = exportvariantsetrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsExportCall) Fields(s ...googleapi.Field) *VariantsetsExportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsExportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportvariantsetrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets/{variantSetId}:export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantSetId": c.variantSetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports variant set data to an external destination.",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.export",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "Required. The ID of the variant set that contains variant data which should be exported. The caller must have READ access to this variant set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variantsets/{variantSetId}:export",
	//   "request": {
	//     "$ref": "ExportVariantSetRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/bigquery",
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.get":

type VariantsetsGetCall struct {
	s            *Service
	variantSetId string
	opt_         map[string]interface{}
}

// Get: Gets a variant set by ID.
func (r *VariantsetsService) Get(variantSetId string) *VariantsetsGetCall {
	c := &VariantsetsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantSetId = variantSetId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsGetCall) Fields(s ...googleapi.Field) *VariantsetsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsGetCall) Do() (*VariantSet, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets/{variantSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantSetId": c.variantSetId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *VariantSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a variant set by ID.",
	//   "httpMethod": "GET",
	//   "id": "genomics.variantsets.get",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "variantSetId": {
	//       "description": "Required. The ID of the variant set.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variantsets/{variantSetId}",
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}

// method id "genomics.variantsets.patch":

type VariantsetsPatchCall struct {
	s            *Service
	variantSetId string
	variantset   *VariantSet
	opt_         map[string]interface{}
}

// Patch: Updates a variant set. This method supports patch semantics.
func (r *VariantsetsService) Patch(variantSetId string, variantset *VariantSet) *VariantsetsPatchCall {
	c := &VariantsetsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.variantSetId = variantSetId
	c.variantset = variantset
	return c
}

// UpdateMask sets the optional parameter "updateMask": An optional mask
// specifying which fields to update. At this time, the only mutable
// field is [metadata][google.genomics.v1.VariantSet.metadata]. The only
// acceptable value is "metadata". If unspecified, all mutable fields
// will be updated.
func (c *VariantsetsPatchCall) UpdateMask(updateMask string) *VariantsetsPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsPatchCall) Fields(s ...googleapi.Field) *VariantsetsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsPatchCall) Do() (*VariantSet, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.variantset)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets/{variantSetId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"variantSetId": c.variantSetId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *VariantSet
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a variant set. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "genomics.variantsets.patch",
	//   "parameterOrder": [
	//     "variantSetId"
	//   ],
	//   "parameters": {
	//     "updateMask": {
	//       "description": "An optional mask specifying which fields to update. At this time, the only mutable field is [metadata][google.genomics.v1.VariantSet.metadata]. The only acceptable value is \"metadata\". If unspecified, all mutable fields will be updated.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "variantSetId": {
	//       "description": "The ID of the variant to be updated (must already exist).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/variantsets/{variantSetId}",
	//   "request": {
	//     "$ref": "VariantSet"
	//   },
	//   "response": {
	//     "$ref": "VariantSet"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics"
	//   ]
	// }

}

// method id "genomics.variantsets.search":

type VariantsetsSearchCall struct {
	s                        *Service
	searchvariantsetsrequest *SearchVariantSetsRequest
	opt_                     map[string]interface{}
}

// Search: Returns a list of all variant sets matching search criteria.
// Implements
// [GlobalAllianceApi.searchVariantSets](http://ga4gh.org/documentation/a
// pi/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchVariantSets).
func (r *VariantsetsService) Search(searchvariantsetsrequest *SearchVariantSetsRequest) *VariantsetsSearchCall {
	c := &VariantsetsSearchCall{s: r.s, opt_: make(map[string]interface{})}
	c.searchvariantsetsrequest = searchvariantsetsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *VariantsetsSearchCall) Fields(s ...googleapi.Field) *VariantsetsSearchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *VariantsetsSearchCall) Do() (*SearchVariantSetsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.searchvariantsetsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/variantsets/search")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SearchVariantSetsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of all variant sets matching search criteria. Implements [GlobalAllianceApi.searchVariantSets](http://ga4gh.org/documentation/api/v0.5.1/ga4gh_api.html#/schema/org.ga4gh.searchVariantSets).",
	//   "httpMethod": "POST",
	//   "id": "genomics.variantsets.search",
	//   "path": "v1/variantsets/search",
	//   "request": {
	//     "$ref": "SearchVariantSetsRequest"
	//   },
	//   "response": {
	//     "$ref": "SearchVariantSetsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/genomics",
	//     "https://www.googleapis.com/auth/genomics.readonly"
	//   ]
	// }

}
