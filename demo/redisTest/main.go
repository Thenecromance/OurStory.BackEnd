package main

import (
	"context"
	"encoding/json"
	"fmt"
	SQL "github.com/Thenecromance/OurStories/SQL/redis"
	"github.com/Thenecromance/OurStories/application/dashboard/models"
	"log"
)

func main() {
	ctx := context.Background()
	cli := SQL.NewRedis()
	/*routes := []models.Route{
		{ID: 1, Name: "Home", Path: "/home2", Component: "/components/Home", AllowRole: Role.Admin},
		{ID: 2, Name: "Dashboard", Path: "/dashboard2", Component: "/components/Dashboard", AllowRole: Role.User},
		{ID: 3, Name: "Profile", Path: "/profile", Component: "/components/Profile", AllowRole: Role.Admin},
		{ID: 4, Name: "Settings", Path: "/settings", Component: "/components/Settings", AllowRole: Role.Master},
	}

	for _, route := range routes {
		routeJson, err := json.Marshal(route)
		if err != nil {
			log.Fatalf("序列化Route对象失败: %v", err)
		}

		// 将JSON字符串存储在Redis中，以ID为键
		err = cli.Set(ctx, "route:"+strconv.Itoa(route.ID), routeJson, 0).Err()
		if err != nil {
			log.Fatalf("存储Route对象失败: %v", err)
		}

		// 将Route对象ID添加到AllowRole对应的集合中
		sid := strconv.Itoa(route.AllowRole)
		fmt.Printf("AllowRole: %s\n", sid)
		err = cli.SAdd(ctx, "role:"+sid, route.ID).Err()
		if err != nil {
			log.Fatalf("添加Route对象ID到角色集合失败: %v", err)
		}
	}*/

	// 根据AllowRole获取Route对象ID列表
	role := "2"
	routeIDs, err := cli.SMembers(ctx, "role:"+role).Result()
	if err != nil {
		log.Fatalf("获取角色集合失败: %v", err)
	}

	fmt.Printf("角色为 %s 的Route对象ID列表: %v\n", role, routeIDs)

	// 根据ID获取Route对象
	for _, id := range routeIDs {
		val, err := cli.Get(ctx, "route:"+id).Result()
		if err != nil {
			log.Fatalf("获取Route对象失败: %v", err)
		}

		var route models.Route
		err = json.Unmarshal([]byte(val), &route)
		if err != nil {
			log.Fatalf("反序列化Route对象失败: %v", err)
		}

		fmt.Printf("从Redis中获取的Route对象: %+v\n", route)
	}
}
