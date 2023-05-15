package permissions

type Router struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Meta      meta   `json:"meta"`
}

type meta struct {
	Title  string `json:"title"`
	IsTrue int    `json:"isTrue"`
}

type Routers struct {
	Path     string   `json:"path"`
	Name     string   `json:"name"`
	Meta     meta     `json:"meta"`
	Children []Router `json:"children"`
}
