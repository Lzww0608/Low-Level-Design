package main


import "fmt"

type Folder struct {
	components []Component
	name 		string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Seraching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}


func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

