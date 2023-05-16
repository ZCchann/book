package permissions

import (
	"book/initalize/database/mysql/authority"
	"log"
	"reflect"
)

func AllData() (data Router) {
	data.Name = "AllData"
	data.Path = "/alldata"
	data.Meta.Title = "数据管理"
	data.Component = "admin/getbook/GetbookView.vue"
	return data
}

func User() (data Router) {
	data.Name = "UserManagement"
	data.Path = "/user"
	data.Meta.Title = "用户管理"
	data.Component = "admin/user/UserView.vue"
	return data
}
func permissions() (data Router) {
	data.Name = "Permissions"
	data.Path = "/permissions"
	data.Meta.Title = "权限管理"
	data.Component = "admin/permissions/permissions.vue"
	return data
}

func AdminMenu() (data Routers) {
	data.Name = "adminMenu"
	data.Path = "/admin"
	data.Meta.Title = "管理员菜单"
	data.Meta.IsTrue = 1
	return data
}

func newOrder() (data Router) {
	data.Name = "NewOrderView"
	data.Path = "/neworder"
	data.Meta.Title = "新增订单"
	data.Component = "order/newOrder/newOrderView.vue"
	return data
}

func orderList() (data Router) {
	data.Name = "OrderListView"
	data.Path = "/orderlist"
	data.Meta.Title = "已下单"
	data.Component = "order/OrderView.vue"
	return data
}

func OrderMenu() (data Routers) {
	data.Name = "orderMenu"
	data.Path = "/order"
	data.Meta.Title = "订单管理"
	data.Meta.IsTrue = 1
	data.Children = append(data.Children, newOrder())
	data.Children = append(data.Children, orderList())
	return data
}

// PermissionFiltering 权限过滤 根据数据库中的权限详情分配动态路由
func PermissionFiltering(uuid string) (routers []Routers, err error) {
	res, err := authority.GetRuleForUUID(uuid)
	if err != nil {
		log.Println(err)
		return
	}

	var adminMenu = AdminMenu()
	v := reflect.ValueOf(res)
	for i := 0; i < v.NumField(); i++ {
		ruleName := v.Type().Field(i).Name
		status := v.Field(i).Interface()
		s := false //临时变量 用于接收遍历结构体的布尔值
		if boolValue, ok := status.(bool); ok {
			s = boolValue
		}

		log.Println(ruleName, s)
		if ruleName == "DataManagement" && s == true {
			adminMenu.Children = append(adminMenu.Children, AllData())
		} else if ruleName == "PermissionManagement" && s == true {
			adminMenu.Children = append(adminMenu.Children, permissions())
		} else if ruleName == "UserManagement" && s == true {
			adminMenu.Children = append(adminMenu.Children, User())
		} else if ruleName == "OrderManagement" && s == true {
			routers = append(routers, OrderMenu())
		}

	}
	if len(adminMenu.Children) > 0 {
		routers = append(routers, adminMenu)
	}

	return routers, nil

}
