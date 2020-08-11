package html

type LI struct {
	Value string
	Child []Node
}
type UL struct {
	Child []Node
}

func (l LI) Html() string {
	output := "<li>"
	output += l.Value
	if l.Child != nil {
		for _, node := range l.Child {
			output += node.Html()
		}
	}
	output += "</li>"

	return output
}

func (l *LI) Add(node Node) *LI {
	l.Child = append(l.Child, node)
	return l
}

func (u UL) Html() string {
	output := "<ul>"
	if u.Child != nil {
		for _, node := range u.Child {
			output += node.Html()
		}
	}
	output += "</ul>"

	return output
}

func (u *UL) Add(node Node) *UL {
	u.Child = append(u.Child, node)
	return u
}
