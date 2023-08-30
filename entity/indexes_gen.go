// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate

package entity

import (
	"encoding/json"
	"errors"
	"fmt"
)

var _ Index = &IndexFlat{}

// IndexFlat idx type for FLAT
type IndexFlat struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexFlat) Name() string {
	return "Flat"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexFlat) IndexType() IndexType {
	return IndexType("FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexFlat) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexFlat) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexFlat) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexFlat create index with construction parameters
func NewIndexFlat(metricType MetricType) (*IndexFlat, error) {
	// auto generate parameters validation code, if any
	return &IndexFlat{
		metricType: metricType,
	}, nil
}

var _ Index = &IndexBinFlat{}

// IndexBinFlat idx type for BIN_FLAT
type IndexBinFlat struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexBinFlat) Name() string {
	return "BinFlat"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexBinFlat) IndexType() IndexType {
	return IndexType("BIN_FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexBinFlat) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexBinFlat) SupportBinary() bool {
	return 2&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexBinFlat) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexBinFlat create index with construction parameters
func NewIndexBinFlat(metricType MetricType,
	nlist int,
) (*IndexBinFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexBinFlat{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexIvfFlat{}

// IndexIvfFlat idx type for IVF_FLAT
type IndexIvfFlat struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexIvfFlat) Name() string {
	return "IvfFlat"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexIvfFlat) IndexType() IndexType {
	return IndexType("IVF_FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexIvfFlat) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexIvfFlat) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexIvfFlat) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfFlat create index with construction parameters
func NewIndexIvfFlat(metricType MetricType,
	nlist int,
) (*IndexIvfFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexIvfFlat{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexBinIvfFlat{}

// IndexBinIvfFlat idx type for BIN_IVF_FLAT
type IndexBinIvfFlat struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexBinIvfFlat) Name() string {
	return "BinIvfFlat"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexBinIvfFlat) IndexType() IndexType {
	return IndexType("BIN_IVF_FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexBinIvfFlat) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexBinIvfFlat) SupportBinary() bool {
	return 2&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexBinIvfFlat) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexBinIvfFlat create index with construction parameters
func NewIndexBinIvfFlat(metricType MetricType,
	nlist int,
) (*IndexBinIvfFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexBinIvfFlat{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexIvfSQ8{}

// IndexIvfSQ8 idx type for IVF_SQ8
type IndexIvfSQ8 struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexIvfSQ8) Name() string {
	return "IvfSQ8"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexIvfSQ8) IndexType() IndexType {
	return IndexType("IVF_SQ8")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexIvfSQ8) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexIvfSQ8) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexIvfSQ8) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfSQ8 create index with construction parameters
func NewIndexIvfSQ8(metricType MetricType,
	nlist int,
) (*IndexIvfSQ8, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexIvfSQ8{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexIvfPQ{}

// IndexIvfPQ idx type for IVF_PQ
type IndexIvfPQ struct { //auto generated fields
	nlist      int
	m          int
	nbits      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexIvfPQ) Name() string {
	return "IvfPQ"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexIvfPQ) IndexType() IndexType {
	return IndexType("IVF_PQ")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexIvfPQ) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexIvfPQ) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexIvfPQ) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
		"m":     fmt.Sprintf("%v", i.m),
		"nbits": fmt.Sprintf("%v", i.nbits),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfPQ create index with construction parameters
func NewIndexIvfPQ(metricType MetricType,
	nlist int,

	m int,

	nbits int,
) (*IndexIvfPQ, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	if nbits < 1 {
		return nil, errors.New("nbits has to be in range [1, 16]")
	}
	if nbits > 16 {
		return nil, errors.New("nbits has to be in range [1, 16]")
	}

	return &IndexIvfPQ{
		//auto generated setting
		nlist: nlist,
		//auto generated setting
		m: m,
		//auto generated setting
		nbits:      nbits,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexHNSW{}

// IndexHNSW idx type for HNSW
type IndexHNSW struct { //auto generated fields
	M              int
	efConstruction int
	metricType     MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexHNSW) Name() string {
	return "HNSW"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexHNSW) IndexType() IndexType {
	return IndexType("HNSW")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexHNSW) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexHNSW) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexHNSW) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"M":              fmt.Sprintf("%v", i.M),
		"efConstruction": fmt.Sprintf("%v", i.efConstruction),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexHNSW create index with construction parameters
func NewIndexHNSW(metricType MetricType,
	M int,

	efConstruction int,
) (*IndexHNSW, error) {
	// auto generate parameters validation code, if any
	if M < 4 {
		return nil, errors.New("M has to be in range [4, 64]")
	}
	if M > 64 {
		return nil, errors.New("M has to be in range [4, 64]")
	}

	if efConstruction < 8 {
		return nil, errors.New("efConstruction has to be in range [8, 512]")
	}
	if efConstruction > 512 {
		return nil, errors.New("efConstruction has to be in range [8, 512]")
	}

	return &IndexHNSW{
		//auto generated setting
		M: M,
		//auto generated setting
		efConstruction: efConstruction,
		metricType:     metricType,
	}, nil
}

var _ Index = &IndexIvfHNSW{}

// IndexIvfHNSW idx type for IVF_HNSW
type IndexIvfHNSW struct { //auto generated fields
	nlist          int
	M              int
	efConstruction int
	metricType     MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexIvfHNSW) Name() string {
	return "IvfHNSW"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexIvfHNSW) IndexType() IndexType {
	return IndexType("IVF_HNSW")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexIvfHNSW) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexIvfHNSW) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexIvfHNSW) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist":          fmt.Sprintf("%v", i.nlist),
		"M":              fmt.Sprintf("%v", i.M),
		"efConstruction": fmt.Sprintf("%v", i.efConstruction),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfHNSW create index with construction parameters
func NewIndexIvfHNSW(metricType MetricType,
	nlist int,

	M int,

	efConstruction int,
) (*IndexIvfHNSW, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	if M < 4 {
		return nil, errors.New("M has to be in range [4, 64]")
	}
	if M > 64 {
		return nil, errors.New("M has to be in range [4, 64]")
	}

	if efConstruction < 8 {
		return nil, errors.New("efConstruction has to be in range [8, 512]")
	}
	if efConstruction > 512 {
		return nil, errors.New("efConstruction has to be in range [8, 512]")
	}

	return &IndexIvfHNSW{
		//auto generated setting
		nlist: nlist,
		//auto generated setting
		M: M,
		//auto generated setting
		efConstruction: efConstruction,
		metricType:     metricType,
	}, nil
}

var _ Index = &IndexDISKANN{}

// IndexDISKANN idx type for DISKANN
type IndexDISKANN struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexDISKANN) Name() string {
	return "DISKANN"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexDISKANN) IndexType() IndexType {
	return IndexType("DISKANN")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexDISKANN) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexDISKANN) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexDISKANN) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexDISKANN create index with construction parameters
func NewIndexDISKANN(metricType MetricType) (*IndexDISKANN, error) {
	// auto generate parameters validation code, if any
	return &IndexDISKANN{
		metricType: metricType,
	}, nil
}

var _ Index = &IndexAUTOINDEX{}

// IndexAUTOINDEX idx type for AUTOINDEX
type IndexAUTOINDEX struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexAUTOINDEX) Name() string {
	return "AUTOINDEX"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexAUTOINDEX) IndexType() IndexType {
	return IndexType("AUTOINDEX")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexAUTOINDEX) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexAUTOINDEX) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexAUTOINDEX) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexAUTOINDEX create index with construction parameters
func NewIndexAUTOINDEX(metricType MetricType) (*IndexAUTOINDEX, error) {
	// auto generate parameters validation code, if any
	return &IndexAUTOINDEX{
		metricType: metricType,
	}, nil
}

var _ Index = &IndexGPUIvfFlat{}

// IndexGPUIvfFlat idx type for GPU_IVF_FLAT
type IndexGPUIvfFlat struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexGPUIvfFlat) Name() string {
	return "GPUIvfFlat"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexGPUIvfFlat) IndexType() IndexType {
	return IndexType("GPU_IVF_FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexGPUIvfFlat) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexGPUIvfFlat) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexGPUIvfFlat) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexGPUIvfFlat create index with construction parameters
func NewIndexGPUIvfFlat(metricType MetricType,
	nlist int,
) (*IndexGPUIvfFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexGPUIvfFlat{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexGPUIvfPQ{}

// IndexGPUIvfPQ idx type for GPU_IVF_PQ
type IndexGPUIvfPQ struct { //auto generated fields
	nlist      int
	m          int
	nbits      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexGPUIvfPQ) Name() string {
	return "GPUIvfPQ"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexGPUIvfPQ) IndexType() IndexType {
	return IndexType("GPU_IVF_PQ")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexGPUIvfPQ) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexGPUIvfPQ) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexGPUIvfPQ) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
		"m":     fmt.Sprintf("%v", i.m),
		"nbits": fmt.Sprintf("%v", i.nbits),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexGPUIvfPQ create index with construction parameters
func NewIndexGPUIvfPQ(metricType MetricType,
	nlist int,

	m int,

	nbits int,
) (*IndexGPUIvfPQ, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	if nbits < 1 {
		return nil, errors.New("nbits has to be in range [1, 64]")
	}
	if nbits > 64 {
		return nil, errors.New("nbits has to be in range [1, 64]")
	}

	return &IndexGPUIvfPQ{
		//auto generated setting
		nlist: nlist,
		//auto generated setting
		m: m,
		//auto generated setting
		nbits:      nbits,
		metricType: metricType,
	}, nil
}

var _ Index = &IndexSCANN{}

// IndexSCANN idx type for IVF_FLAT
type IndexSCANN struct { //auto generated fields
	nlist      int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func (i *IndexSCANN) Name() string {
	return "SCANN"
}

// IndexType returns IndexType, implementing Index interface
func (i *IndexSCANN) IndexType() IndexType {
	return IndexType("IVF_FLAT")
}

// FieldName returns FieldName, implementing Index interface
func (i *IndexSCANN) FieldName() string {
	return "vec_field"
}

// SupportBinary returns whether index type support binary vector
func (i *IndexSCANN) SupportBinary() bool {
	return 0&2 > 0
}

// Params returns index construction params, implementing Index interface
func (i *IndexSCANN) Params() map[string]string {
	params := map[string]string{ //auto generated mapping
		"nlist": fmt.Sprintf("%v", i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string{
		"params":      string(bs),
		"index_type":  string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexSCANN create index with construction parameters
func NewIndexSCANN(metricType MetricType,
	nlist int,
) (*IndexSCANN, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist has to be in range [1, 65536]")
	}

	return &IndexSCANN{
		//auto generated setting
		nlist:      nlist,
		metricType: metricType,
	}, nil
}
