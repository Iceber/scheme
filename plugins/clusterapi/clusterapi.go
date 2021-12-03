package main

import (
	"fmt"

	globalscheme "github.com/iceber/scheme/pkg/global"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

func init() {
	fmt.Println("register clusterapi")

	v1beta1.AddToScheme(globalscheme.Scheme)
}
