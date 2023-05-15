package permissions

func AllData() (data Router) {
	data.Name = "AllData"
	data.Path = "/alldata"
	data.Meta.Title = "数据管理"
	data.Component = "admin/getbook/GetbookView.vue"
	return data
}

func User() (data Router) {
	data.Name = "User"
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
	data.Children = append(data.Children, AllData())
	data.Children = append(data.Children, User())
	data.Children = append(data.Children, permissions())
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
