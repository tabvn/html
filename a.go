package html

type A struct {
	Child []Node
	Value string
	URL   string
}

func (a A) Html() string {
	output := "<a href=\"" + a.URL + "\">"
	output += a.Value
	if a.Child != nil {
		for _, node := range a.Child {
			output += node.Html()
		}
	}
	output += "</a>"
	return output
}
