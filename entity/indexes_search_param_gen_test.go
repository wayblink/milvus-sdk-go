// Code generated by go generate; DO NOT EDIT
// This file is generated by go generate at 2023-02-13 22:25:33.586351 +0800 CST m=+0.014629626

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestIndexFlatSearchParam(t *testing.T) {
	

	t.Run("valid usage case", func(t *testing.T){
		
		
		idx0, err := NewIndexFlatSearchParam(
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
}

func TestIndexBinFlatSearchParam(t *testing.T) {
	
	var nprobe int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe = 10
		idx0, err := NewIndexBinFlatSearchParam(
			nprobe,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe = 0
		idx0, err := NewIndexBinFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe = 65537
		idx1, err := NewIndexBinFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexIvfFlatSearchParam(t *testing.T) {
	
	var nprobe int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe = 10
		idx0, err := NewIndexIvfFlatSearchParam(
			nprobe,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe = 0
		idx0, err := NewIndexIvfFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe = 65537
		idx1, err := NewIndexIvfFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexBinIvfFlatSearchParam(t *testing.T) {
	
	var nprobe int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe = 10
		idx0, err := NewIndexBinIvfFlatSearchParam(
			nprobe,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe = 0
		idx0, err := NewIndexBinIvfFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe = 65537
		idx1, err := NewIndexBinIvfFlatSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexIvfSQ8SearchParam(t *testing.T) {
	
	var nprobe int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe = 10
		idx0, err := NewIndexIvfSQ8SearchParam(
			nprobe,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe = 0
		idx0, err := NewIndexIvfSQ8SearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe = 65537
		idx1, err := NewIndexIvfSQ8SearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexIvfPQSearchParam(t *testing.T) {
	
	var nprobe int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe = 10
		idx0, err := NewIndexIvfPQSearchParam(
			nprobe,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe = 0
		idx0, err := NewIndexIvfPQSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe = 65537
		idx1, err := NewIndexIvfPQSearchParam(
			nprobe,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexHNSWSearchParam(t *testing.T) {
	
	var ef int

	t.Run("valid usage case", func(t *testing.T){
		
		ef = 16
		idx0, err := NewIndexHNSWSearchParam(
			ef,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		ef = 0
		idx0, err := NewIndexHNSWSearchParam(
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		ef = 32769
		idx1, err := NewIndexHNSWSearchParam(
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexIvfHNSWSearchParam(t *testing.T) {
	
	var nprobe int
	var ef int

	t.Run("valid usage case", func(t *testing.T){
		
		nprobe, ef = 10, 16
		idx0, err := NewIndexIvfHNSWSearchParam(
			nprobe,
			ef,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		nprobe, ef = 0, 16
		idx0, err := NewIndexIvfHNSWSearchParam(
			nprobe,
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		nprobe, ef = 65537, 16
		idx1, err := NewIndexIvfHNSWSearchParam(
			nprobe,
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
		nprobe, ef = 10, 0
		idx2, err := NewIndexIvfHNSWSearchParam(
			nprobe,
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx2)
		
		nprobe, ef = 10, 32769
		idx3, err := NewIndexIvfHNSWSearchParam(
			nprobe,
			ef,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx3)
		
	})
	
}

func TestIndexANNOYSearchParam(t *testing.T) {
	
	var search_k int

	t.Run("valid usage case", func(t *testing.T){
		
		search_k = -1
		idx0, err := NewIndexANNOYSearchParam(
			search_k,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
		search_k = 20
		idx1, err := NewIndexANNOYSearchParam(
			search_k,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx1)
		assert.NotNil(t, idx1.Params())
		
	})
	
}

func TestIndexDISKANNSearchParam(t *testing.T) {
	
	var search_list int

	t.Run("valid usage case", func(t *testing.T){
		
		search_list = 30
		idx0, err := NewIndexDISKANNSearchParam(
			search_list,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		search_list = 0
		idx0, err := NewIndexDISKANNSearchParam(
			search_list,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		search_list = 65537
		idx1, err := NewIndexDISKANNSearchParam(
			search_list,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
	})
	
}

func TestIndexAUTOINDEXSearchParam(t *testing.T) {
	
	var level int

	t.Run("valid usage case", func(t *testing.T){
		
		level = 1
		idx0, err := NewIndexAUTOINDEXSearchParam(
			level,
		)
		assert.Nil(t, err)
		assert.NotNil(t, idx0)
		assert.NotNil(t, idx0.Params())
		
	})
	
	t.Run("invalid usage case", func(t *testing.T){
		
		level = 0
		idx0, err := NewIndexAUTOINDEXSearchParam(
			level,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx0)
		
		level = 10
		idx1, err := NewIndexAUTOINDEXSearchParam(
			level,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx1)
		
		level = -1
		idx2, err := NewIndexAUTOINDEXSearchParam(
			level,
		)
		assert.NotNil(t, err)
		assert.Nil(t, idx2)
		
	})
	
}

func TestIndexTRIESearchParam(t *testing.T) {
	

	t.Run("valid usage case", func(t *testing.T){
		
	})
	
}

