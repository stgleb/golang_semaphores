package receiver_type_benckmark

import "testing"

/*
Simple test suite intended to compare cost of using method with value receiver
or pointer receiver. When pointer receiver is used compiler issues & operator to
get an address of variable.

	func (p *P) foo() {
	}

	var v P
	v.foo() // (&v).foo()

Results show that average time of calling pointer receiver methods is the same,
in statistical meaning, as calling methods with value receiver.
*/


func BenchmarkPtrReceiver(b *testing.B) {
	receiver := Receiver{
		count: 0,
	}

	b.ResetTimer()
	for i:= 0;i < b.N; i++ {
		receiver.IncPtr()
	}
}


func BenchmarkValReceiver(b *testing.B) {
	receiver := Receiver{
		count: 0,
	}

	b.ResetTimer()
	for i:= 0;i < b.N; i++ {
		receiver.IncPtr()
	}
}
