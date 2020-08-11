package html

type DIV struct {
	Child []Node
	Value string
}

func (d DIV) Html() string {
	output := "<div>"
	output += d.Value
	if d.Child != nil {
		for _, node := range d.Child {
			output += node.Html()
		}
	}
	output += "</div>"

	return output
}

func (d *DIV) Add(node Node) *DIV {
	/*
		<div>
			<div>child 1</div>
		</div>
	*/
	d.Child = append(d.Child, node)
	return d
}
