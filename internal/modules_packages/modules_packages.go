package modules_packages

import (
	"fmt"

	"github.com/shivik/go-lang-cookbook/internal/modules_packages/pkgdemo/mathutil"
)

// ModulesPackagesDemo shows how to import and use a package inside the module.
func ModulesPackagesDemo() {
	n := 21
	fmt.Println("mathutil.Double:", mathutil.Double(n))
}
