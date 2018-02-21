// Test that blank imports in library packages are flagged.

// Package foo ...
package foo

// The instructions need to go before the imports below so they will not be
// mistaken for documentation.

/* MATCH /blank import/ */ 
import _ "encoding/json"
import (
	"fmt"
	_ "os"
	_ "path"
	_ "encoding/json"
	_ "encoding/base64"
	_ "compress/zlib"
	_ "syscall"
	_ "path/filepath"
	"go/ast"
	_ "go/scanner" // Don't gripe about this or the following line.
	_ "go/token"
)





var (
	_ fmt.Stringer // for "fmt"
	_ ast.Node     // for "go/ast"
)
