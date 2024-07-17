package hook

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net"
)

type Int64SliceSupport struct {
}

func (ih Int64SliceSupport) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

func (ih Int64SliceSupport) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}
func (ih Int64SliceSupport) argsIsInt64s(args interface{}) bool {
	switch args.(type) {
	case []int64:
		return true
	default:
	}
	return false
}

func uint64ToLenBytes(v uint64, l int) (b []byte) {
	b = make([]byte, l)

	for i := 0; i < l; i++ {
		f := 8 * i
		b[i] = byte(v >> f)
	}

	return
}

func int64ToLenBytes(v int64, l int) (b []byte) {
	return uint64ToLenBytes(uint64(v), l)
}
func (ih Int64SliceSupport) transformInt64s(args []int64) []byte {
	var newValues []byte
	for _, value := range args {

		newValues = append(newValues, int64ToLenBytes(value, 8)...)
	}
	fmt.Println("newValues is ", newValues)
	return newValues
}
func (ih Int64SliceSupport) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {

		for idx, v := range cmd.Args() {
			if ih.argsIsInt64s(v) {
				cmd.Args()[idx] = ih.transformInt64s(v.([]int64))
			}
		}

		next(ctx, cmd)
		fmt.Println(cmd.String())
		/*fmt.Println("\n\n")
		for _, v := range cmd.Args() {
			fmt.Println(reflect.ValueOf(v), "\t", reflect.TypeOf(v))
		}

		print("hook-1 end\n\n")*/
		return nil
	}
}
