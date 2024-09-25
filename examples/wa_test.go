package testcases

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

func GenFloatVector(dim int64) []float32 {
	vector := make([]float32, 0, dim)
	for j := 0; j < int(dim); j++ {
		vector = append(vector, rand.Float32())
	}
	return vector
}

func newContext(t *testing.T, timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(func() {
		cancel()
	})
	return ctx
}

func TestDynamic(t *testing.T) {
	ctx := newContext(t, time.Second*120)
	// connect
	mc, err := client.NewClient(ctx, client.Config{Address: "localhost:19530"})
	require.Nil(t, err)

	// create collection and enable dynamic
	name := "test_dynamic"
	fields := []*entity.Field{
		entity.NewField().WithName("pk").WithDataType(entity.FieldTypeInt64).WithIsPrimaryKey(true),
		entity.NewField().WithName("vec").WithDataType(entity.FieldTypeFloatVector).WithDim(128),
	}
	schema := &entity.Schema{
		CollectionName:     name,
		Fields:             fields,
		EnableDynamicField: true,
	}
	mc.CreateCollection(ctx, schema, 1)

	// insert data
	nb := 1000
	pkData := make([]int64, 0, nb)
	vecFloatValues := make([][]float32, 0, nb)
	dynamicDataX := make([]string, 0, nb)
	dynamicDataY := make([]int64, 0, nb)
	for i := 0; i < nb; i++ {
		pkData = append(pkData, int64(i))
		vec := GenFloatVector(128)
		vecFloatValues = append(vecFloatValues, vec)
		dynamicDataX = append(dynamicDataX, strconv.Itoa(i))
		dynamicDataY = append(dynamicDataY, int64(i))
	}

	mc.Insert(ctx, name, "",
		entity.NewColumnInt64("pk", pkData),
		entity.NewColumnFloatVector("vec", 128, vecFloatValues),
		entity.NewColumnVarChar("x", dynamicDataX),
		entity.NewColumnInt64("y", dynamicDataY),
	)
	mc.Flush(ctx, name, false)

	idx, _ := entity.NewIndexAUTOINDEX(entity.COSINE)
	mc.CreateIndex(ctx, name, "vec", idx, false)
	mc.LoadCollection(ctx, name, false)
	res, _ := mc.Query(ctx, name, []string{}, "pk == 0", []string{"x"})
	log.Println(res.GetColumn("x").FieldData())
}
