package site

type NavItem struct {
	Name string
	Href string
}

var NavItems = []NavItem{
	{
		Name: "Market",
		Href: "/",
	},
	{
		Name: "Positions",
		Href: "/positions",
	},
}
