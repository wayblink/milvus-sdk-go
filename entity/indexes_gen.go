// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate at 2023-02-13 16:16:51.475873 +0800 CST m=+0.012486210

package entity

import (
	"errors"
	"encoding/json"
	"fmt"
)



var _ Index = &IndexFlat{}

// IndexFlat idx type for FLAT
type IndexFlat struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexFlat) Name() string {
	return "Flat"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexFlat) IndexType() IndexType {
	return IndexType("FLAT")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexFlat) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexFlat) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexFlat create index with construction parameters
func NewIndexFlat(metricType MetricType, ) (*IndexFlat, error) {
	// auto generate parameters validation code, if any
	return &IndexFlat{ 
	metricType: metricType,
	}, nil
}


var _ Index = &IndexBinFlat{}

// IndexBinFlat idx type for BIN_FLAT
type IndexBinFlat struct { //auto generated fields
	nlist int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexBinFlat) Name() string {
	return "BinFlat"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexBinFlat) IndexType() IndexType {
	return IndexType("BIN_FLAT")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexBinFlat) SupportBinary() bool {
	return 2 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexBinFlat) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexBinFlat create index with construction parameters
func NewIndexBinFlat(metricType MetricType, 
	nlist int,
) (*IndexBinFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	return &IndexBinFlat{ 
	//auto generated setting
	nlist: nlist,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexIvfFlat{}

// IndexIvfFlat idx type for IVF_FLAT
type IndexIvfFlat struct { //auto generated fields
	nlist int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexIvfFlat) Name() string {
	return "IvfFlat"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexIvfFlat) IndexType() IndexType {
	return IndexType("IVF_FLAT")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexIvfFlat) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexIvfFlat) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfFlat create index with construction parameters
func NewIndexIvfFlat(metricType MetricType, 
	nlist int,
) (*IndexIvfFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	return &IndexIvfFlat{ 
	//auto generated setting
	nlist: nlist,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexBinIvfFlat{}

// IndexBinIvfFlat idx type for BIN_IVF_FLAT
type IndexBinIvfFlat struct { //auto generated fields
	nlist int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexBinIvfFlat) Name() string {
	return "BinIvfFlat"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexBinIvfFlat) IndexType() IndexType {
	return IndexType("BIN_IVF_FLAT")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexBinIvfFlat) SupportBinary() bool {
	return 2 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexBinIvfFlat) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexBinIvfFlat create index with construction parameters
func NewIndexBinIvfFlat(metricType MetricType, 
	nlist int,
) (*IndexBinIvfFlat, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	return &IndexBinIvfFlat{ 
	//auto generated setting
	nlist: nlist,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexIvfSQ8{}

// IndexIvfSQ8 idx type for IVF_SQ8
type IndexIvfSQ8 struct { //auto generated fields
	nlist int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexIvfSQ8) Name() string {
	return "IvfSQ8"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexIvfSQ8) IndexType() IndexType {
	return IndexType("IVF_SQ8")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexIvfSQ8) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexIvfSQ8) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexIvfSQ8 create index with construction parameters
func NewIndexIvfSQ8(metricType MetricType, 
	nlist int,
) (*IndexIvfSQ8, error) {
	// auto generate parameters validation code, if any
	if nlist < 1 {
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	return &IndexIvfSQ8{ 
	//auto generated setting
	nlist: nlist,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexIvfPQ{}

// IndexIvfPQ idx type for IVF_PQ
type IndexIvfPQ struct { //auto generated fields
	nlist int
	m int
	nbits int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexIvfPQ) Name() string {
	return "IvfPQ"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexIvfPQ) IndexType() IndexType {
	return IndexType("IVF_PQ")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexIvfPQ) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexIvfPQ) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
		"m": fmt.Sprintf("%v",i.m),
		"nbits": fmt.Sprintf("%v",i.nbits),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
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
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	
	
	if nbits < 1 {
		return nil, errors.New("nbits not valid")
	}
	if nbits > 16 {
		return nil, errors.New("nbits not valid")
	}
	
	return &IndexIvfPQ{ 
	//auto generated setting
	nlist: nlist,
	//auto generated setting
	m: m,
	//auto generated setting
	nbits: nbits,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexHNSW{}

// IndexHNSW idx type for HNSW
type IndexHNSW struct { //auto generated fields
	M int
	efConstruction int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexHNSW) Name() string {
	return "HNSW"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexHNSW) IndexType() IndexType {
	return IndexType("HNSW")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexHNSW) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexHNSW) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"M": fmt.Sprintf("%v",i.M),
		"efConstruction": fmt.Sprintf("%v",i.efConstruction),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
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
		return nil, errors.New("M not valid")
	}
	if M > 64 {
		return nil, errors.New("M not valid")
	}
	
	if efConstruction < 8 {
		return nil, errors.New("efConstruction not valid")
	}
	if efConstruction > 512 {
		return nil, errors.New("efConstruction not valid")
	}
	
	return &IndexHNSW{ 
	//auto generated setting
	M: M,
	//auto generated setting
	efConstruction: efConstruction,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexIvfHNSW{}

// IndexIvfHNSW idx type for IVF_HNSW
type IndexIvfHNSW struct { //auto generated fields
	nlist int
	M int
	efConstruction int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexIvfHNSW) Name() string {
	return "IvfHNSW"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexIvfHNSW) IndexType() IndexType {
	return IndexType("IVF_HNSW")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexIvfHNSW) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexIvfHNSW) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"nlist": fmt.Sprintf("%v",i.nlist),
		"M": fmt.Sprintf("%v",i.M),
		"efConstruction": fmt.Sprintf("%v",i.efConstruction),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
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
		return nil, errors.New("nlist not valid")
	}
	if nlist > 65536 {
		return nil, errors.New("nlist not valid")
	}
	
	if M < 4 {
		return nil, errors.New("M not valid")
	}
	if M > 64 {
		return nil, errors.New("M not valid")
	}
	
	if efConstruction < 8 {
		return nil, errors.New("efConstruction not valid")
	}
	if efConstruction > 512 {
		return nil, errors.New("efConstruction not valid")
	}
	
	return &IndexIvfHNSW{ 
	//auto generated setting
	nlist: nlist,
	//auto generated setting
	M: M,
	//auto generated setting
	efConstruction: efConstruction,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexANNOY{}

// IndexANNOY idx type for ANNOY
type IndexANNOY struct { //auto generated fields
	n_trees int
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexANNOY) Name() string {
	return "ANNOY"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexANNOY) IndexType() IndexType {
	return IndexType("ANNOY")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexANNOY) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexANNOY) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
		"n_trees": fmt.Sprintf("%v",i.n_trees),
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexANNOY create index with construction parameters
func NewIndexANNOY(metricType MetricType, 
	n_trees int,
) (*IndexANNOY, error) {
	// auto generate parameters validation code, if any
	if n_trees < 1 {
		return nil, errors.New("n_trees not valid")
	}
	if n_trees > 1024 {
		return nil, errors.New("n_trees not valid")
	}
	
	return &IndexANNOY{ 
	//auto generated setting
	n_trees: n_trees,
	metricType: metricType,
	}, nil
}


var _ Index = &IndexDISKANN{}

// IndexDISKANN idx type for DISKANN
type IndexDISKANN struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexDISKANN) Name() string {
	return "DISKANN"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexDISKANN) IndexType() IndexType {
	return IndexType("DISKANN")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexDISKANN) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexDISKANN) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexDISKANN create index with construction parameters
func NewIndexDISKANN(metricType MetricType, ) (*IndexDISKANN, error) {
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
func(i *IndexAUTOINDEX) Name() string {
	return "AUTOINDEX"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexAUTOINDEX) IndexType() IndexType {
	return IndexType("AUTOINDEX")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexAUTOINDEX) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexAUTOINDEX) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexAUTOINDEX create index with construction parameters
func NewIndexAUTOINDEX(metricType MetricType, ) (*IndexAUTOINDEX, error) {
	// auto generate parameters validation code, if any
	return &IndexAUTOINDEX{ 
	metricType: metricType,
	}, nil
}


var _ Index = &IndexTRIE{}

// IndexTRIE idx type for TRIE
type IndexTRIE struct { //auto generated fields
	metricType MetricType
}

// Name returns index type name, implementing Index interface
func(i *IndexTRIE) Name() string {
	return "TRIE"
}

// IndexType returns IndexType, implementing Index interface
func(i *IndexTRIE) IndexType() IndexType {
	return IndexType("TRIE")
}

// SupportBinary returns whether index type support binary vector
func(i *IndexTRIE) SupportBinary() bool {
	return 0 & 2 > 0
}

// Params returns index construction params, implementing Index interface
func(i *IndexTRIE) Params() map[string]string {
	params := map[string]string {//auto generated mapping 
	}
	bs, _ := json.Marshal(params)
	return map[string]string {
		"params": string(bs),
		"index_type": string(i.IndexType()),
		"metric_type": string(i.metricType),
	}
}

// NewIndexTRIE create index with construction parameters
func NewIndexTRIE(metricType MetricType, ) (*IndexTRIE, error) {
	// auto generate parameters validation code, if any
	return &IndexTRIE{ 
	metricType: metricType,
	}, nil
}

