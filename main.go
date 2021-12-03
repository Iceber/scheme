package main

import (
	"fmt"

	globalscheme "github.com/iceber/scheme/pkg/global"
	"github.com/iceber/scheme/pkg/plugin"
	"github.com/liushuochen/gotable"
)

func main() {
	if err := plugin.LoadPlugins("./plugins"); err != nil {
		panic(err)
	}

	table, err := gotable.Create("APIVersion", "Kind", "Struct")
	if err != nil {
		panic(err)
	}

	types := globalscheme.Scheme.AllKnownTypes()
	for gvk, typ := range types {
		table.AddRow([]string{gvk.GroupVersion().String(), gvk.Kind, typ.String()})
	}
	fmt.Println(table)
}
