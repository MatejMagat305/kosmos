package shapes

type Names map[string]bool

func (n Names) IsIn(name string) bool {
	_, ok := n[name]
	return ok
}

func (n Names) Add(name string) {
	n[name] = true
}

func (n Names) Delete(name string) {
	if n.IsIn(name) {
		delete(n, name)
	}
}

func (n Names) DeleteAll() {
	for s := range n {
		n.Delete(s)
	}
}