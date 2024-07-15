package id

import "github.com/bwmarrin/snowflake"

var (
	inst *snowflake.Node
)

func init() {
	node, err := snowflake.NewNode(1)

	if err != nil {
		panic(err)
	}

	inst = node
}

func Generate() int64 {
	return inst.Generate().Int64()
}

func GenerateString() string {
	return inst.Generate().String()
}

func GenerateBase64() string {
	return inst.Generate().Base64()
}
