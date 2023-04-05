package authority

type Routers struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Component string `json:"component"`
}

func AllData() (data Routers) {
	data.Name = "AllData"
	data.Path = "/alldata"
	data.Title = "数据管理"
	data.Component = "admin/getbook/GetbookView.vue"
	return data
}

func User() (data Routers) {
	data.Name = "User"
	data.Path = "/user"
	data.Title = "用户管理"
	data.Component = "admin/user/UserView.vue"

	return data
}

func Admin() (ret []Routers) {
	ret = append(ret, AllData())
	ret = append(ret, User())
	return ret
}

func Supplier() (ret []Routers) {
	ret = append(ret, AllData())
	return ret
}
