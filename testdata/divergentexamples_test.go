// Test of examples with divergent packages.

// Package bytes_test ...
package bytes_test

import "bytes"

func Example() {} // ok because is package-level.

func Example_suffix() // ok because refers to suffix annotation.

func Example_BadSuffix() // MATCH /Example_BadSuffix has malformed example suffix: BadSuffix/

func ExampleBuffer() // ok because refers to bytes.Buffer.

func ExampleBuffer_Len() {} // ok because refers to known method.

func ExampleBuffer_Clear() {} // MATCH /ExampleBuffer_Clear refers to unknown field or method: Buffer.Clear/

func ExampleBuffer_suffix() {} // ok because refers to suffix annotation.

func ExampleBuffer_Write_Bad() {} // MATCH /ExampleBuffer_Write_Bad has malformed example suffix: Bad/

func ExampleBuffer_Write_suffix() {} // ok because refers to known method with valid suffix.

func ExampleErrTooLarge() {} // ok because refers to bytes.ErrTooLarge.

func ExampleBuffer_Reset() bool { return true } // MATCH /ExampleBuffer_Reset should return nothing/

func ExampleBuffer_Grow(i int) {} // MATCH /ExampleBuffer_Grow should be niladic/

// "Puffer" is German for "Buffer".

func ExamplePuffer() // MATCH /ExamplePuffer refers to unknown identifier: Puffer/

func ExamplePuffer_Len() // MATCH /ExamplePuffer_Len refers to unknown identifier: Puffer/

func ExamplePuffer_suffix() // MATCH /ExamplePuffer_suffix refers to unknown identifier: Puffer/
