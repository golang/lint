// Test of examples.

// Package examples ...
package examples

// Buf is a ...
type Buf []byte

// Append ...
func (*Buf) Append([]byte) {}

func (Buf) Reset() {}

func (Buf) Len() int { return 0 }

// DefaultBuf is a ...
var DefaultBuf Buf

func Example() {} // ok because is package-level.

func Example_suffix() // ok because refers to suffix annotation.

func Example_BadSuffix() // MATCH /Example_BadSuffix has malformed example suffix: BadSuffix/

func ExampleBuf() // ok because refers to known top-level type.

func ExampleBuf_Append() {} // ok because refers to known method.

func ExampleBuf_Clear() {} // MATCH /ExampleBuf_Clear refers to unknown field or method: Buf.Clear/

func ExampleBuf_suffix() {} // ok because refers to suffix annotation.

func ExampleBuf_Append_Bad() {} // MATCH /ExampleBuf_Append_Bad has malformed example suffix: Bad/

func ExampleBuf_Append_suffix() {} // ok because refers to known method with valid suffix.

func ExampleDefaultBuf() {} // ok because refers to top-level identifier.

func ExampleBuf_Reset() bool { return true } // MATCH /ExampleBuf_Reset should return nothing/

func ExampleBuf_Len(i int) {} // MATCH /ExampleBuf_Len should be niladic/

// "Puffer" is German for "Buffer".

func ExamplePuffer() // MATCH /ExamplePuffer refers to unknown identifier: Puffer/

func ExamplePuffer_Append() // MATCH /ExamplePuffer_Append refers to unknown identifier: Puffer/

func ExamplePuffer_suffix() // MATCH /ExamplePuffer_suffix refers to unknown identifier: Puffer/
