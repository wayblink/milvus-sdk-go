// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate

package entity

import (
	"errors"
)

var _ SearchParam = &IndexFlatSearchParam{}

// IndexFlatSearchParam search param struct for index type FLAT
type IndexFlatSearchParam struct { //auto generated fields
	baseSearchParams
}

// NewIndexFlatSearchParam create index search param
func NewIndexFlatSearchParam() (*IndexFlatSearchParam, error) {
	// auto generate parameters validation code, if any
	sp := &IndexFlatSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting

	return sp, nil
}

var _ SearchParam = &IndexBinFlatSearchParam{}

// IndexBinFlatSearchParam search param struct for index type BIN_FLAT
type IndexBinFlatSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexBinFlatSearchParam create index search param
func NewIndexBinFlatSearchParam(
	nprobe int,
) (*IndexBinFlatSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexBinFlatSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexIvfFlatSearchParam{}

// IndexIvfFlatSearchParam search param struct for index type IVF_FLAT
type IndexIvfFlatSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexIvfFlatSearchParam create index search param
func NewIndexIvfFlatSearchParam(
	nprobe int,
) (*IndexIvfFlatSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexIvfFlatSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexBinIvfFlatSearchParam{}

// IndexBinIvfFlatSearchParam search param struct for index type BIN_IVF_FLAT
type IndexBinIvfFlatSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexBinIvfFlatSearchParam create index search param
func NewIndexBinIvfFlatSearchParam(
	nprobe int,
) (*IndexBinIvfFlatSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexBinIvfFlatSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexIvfSQ8SearchParam{}

// IndexIvfSQ8SearchParam search param struct for index type IVF_SQ8
type IndexIvfSQ8SearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexIvfSQ8SearchParam create index search param
func NewIndexIvfSQ8SearchParam(
	nprobe int,
) (*IndexIvfSQ8SearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexIvfSQ8SearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexIvfPQSearchParam{}

// IndexIvfPQSearchParam search param struct for index type IVF_PQ
type IndexIvfPQSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexIvfPQSearchParam create index search param
func NewIndexIvfPQSearchParam(
	nprobe int,
) (*IndexIvfPQSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexIvfPQSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexHNSWSearchParam{}

// IndexHNSWSearchParam search param struct for index type HNSW
type IndexHNSWSearchParam struct { //auto generated fields
	baseSearchParams

	ef int
}

// NewIndexHNSWSearchParam create index search param
func NewIndexHNSWSearchParam(
	ef int,
) (*IndexHNSWSearchParam, error) {
	// auto generate parameters validation code, if any
	if ef < 1 {
		return nil, errors.New("ef has to be in range [1, 32768]")
	}
	if ef > 32768 {
		return nil, errors.New("ef has to be in range [1, 32768]")
	}

	sp := &IndexHNSWSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["ef"] = ef

	return sp, nil
}

var _ SearchParam = &IndexIvfHNSWSearchParam{}

// IndexIvfHNSWSearchParam search param struct for index type IVF_HNSW
type IndexIvfHNSWSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
	ef     int
}

// NewIndexIvfHNSWSearchParam create index search param
func NewIndexIvfHNSWSearchParam(
	nprobe int,

	ef int,
) (*IndexIvfHNSWSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	if ef < 1 {
		return nil, errors.New("ef has to be in range [1, 32768]")
	}
	if ef > 32768 {
		return nil, errors.New("ef has to be in range [1, 32768]")
	}

	sp := &IndexIvfHNSWSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe
	sp.params["ef"] = ef

	return sp, nil
}

var _ SearchParam = &IndexDISKANNSearchParam{}

// IndexDISKANNSearchParam search param struct for index type DISKANN
type IndexDISKANNSearchParam struct { //auto generated fields
	baseSearchParams

	search_list int
}

// NewIndexDISKANNSearchParam create index search param
func NewIndexDISKANNSearchParam(
	search_list int,
) (*IndexDISKANNSearchParam, error) {
	// auto generate parameters validation code, if any
	if search_list < 1 {
		return nil, errors.New("search_list has to be in range [1, 65535]")
	}
	if search_list > 65535 {
		return nil, errors.New("search_list has to be in range [1, 65535]")
	}

	sp := &IndexDISKANNSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["search_list"] = search_list

	return sp, nil
}

var _ SearchParam = &IndexAUTOINDEXSearchParam{}

// IndexAUTOINDEXSearchParam search param struct for index type AUTOINDEX
type IndexAUTOINDEXSearchParam struct { //auto generated fields
	baseSearchParams

	level int
}

// NewIndexAUTOINDEXSearchParam create index search param
func NewIndexAUTOINDEXSearchParam(
	level int,
) (*IndexAUTOINDEXSearchParam, error) {
	// auto generate parameters validation code, if any
	if level < 1 {
		return nil, errors.New("level has to be in range [1, 9223372036854775807]")
	}
	if level > 9223372036854775807 {
		return nil, errors.New("level has to be in range [1, 9223372036854775807]")
	}

	sp := &IndexAUTOINDEXSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["level"] = level

	return sp, nil
}

var _ SearchParam = &IndexGPUIvfFlatSearchParam{}

// IndexGPUIvfFlatSearchParam search param struct for index type GPU_IVF_FLAT
type IndexGPUIvfFlatSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexGPUIvfFlatSearchParam create index search param
func NewIndexGPUIvfFlatSearchParam(
	nprobe int,
) (*IndexGPUIvfFlatSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexGPUIvfFlatSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexGPUIvfPQSearchParam{}

// IndexGPUIvfPQSearchParam search param struct for index type GPU_IVF_PQ
type IndexGPUIvfPQSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe int
}

// NewIndexGPUIvfPQSearchParam create index search param
func NewIndexGPUIvfPQSearchParam(
	nprobe int,
) (*IndexGPUIvfPQSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	sp := &IndexGPUIvfPQSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe

	return sp, nil
}

var _ SearchParam = &IndexSCANNSearchParam{}

// IndexSCANNSearchParam search param struct for index type SCANN
type IndexSCANNSearchParam struct { //auto generated fields
	baseSearchParams

	nprobe    int
	reorder_k int
}

// NewIndexSCANNSearchParam create index search param
func NewIndexSCANNSearchParam(
	nprobe int,

	reorder_k int,
) (*IndexSCANNSearchParam, error) {
	// auto generate parameters validation code, if any
	if nprobe < 1 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}
	if nprobe > 65536 {
		return nil, errors.New("nprobe has to be in range [1, 65536]")
	}

	if reorder_k < 1 {
		return nil, errors.New("reorder_k has to be in range [1, 9223372036854775807]")
	}
	if reorder_k > 9223372036854775807 {
		return nil, errors.New("reorder_k has to be in range [1, 9223372036854775807]")
	}

	sp := &IndexSCANNSearchParam{
		baseSearchParams: newBaseSearchParams(),
	}

	//auto generated setting
	sp.params["nprobe"] = nprobe
	sp.params["reorder_k"] = reorder_k

	return sp, nil
}
