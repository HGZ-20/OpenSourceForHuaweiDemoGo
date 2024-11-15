package main

import (
	"fmt"
)

var (
	version     = "v5.0.0"
	productline = "htrunk2.csi.gaussdb_kernel"
	productname = "opengaussgo"
	versionid   = "r1"
)

func main() {
	fmt.Printf("%s-%s.%s.%s\n", version, productline, productname, versionid)
}
