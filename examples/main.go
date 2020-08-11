package main

import "github.com/tabvn/html"

func main() {
	container := html.DIV{
		Child: []html.Node{
			html.DIV{Value: "This is content of child 1"},
		},
	}

	container.Add(html.DIV{Value: "Child 2"})

	ul := html.UL{
		Child: []html.Node{
			html.LI{Value: "List item 1"},
		},
	}
	li := html.LI{Value: "List item 2"}
	li.Add(html.A{Value: "Click here", URL: "http://drupalexp.com"})
	ul.Add(li)
	container.Add(ul)
	println(container.Html())
}
