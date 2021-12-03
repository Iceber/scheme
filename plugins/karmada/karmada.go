package main

import (
	"fmt"

	globalscheme "github.com/iceber/scheme/pkg/global"
	clusterv1alpha1 "github.com/karmada-io/karmada/pkg/apis/cluster/v1alpha1"
	workv1alpha1 "github.com/karmada-io/karmada/pkg/apis/work/v1alpha1"
)

func init() {
	fmt.Println("register karmada")

	workv1alpha1.Install(globalscheme.Scheme)
	clusterv1alpha1.Install(globalscheme.Scheme)
}
