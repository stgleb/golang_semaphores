package receiver_type_benckmark


type Receiver struct {
	count int
}

func (receiver *Receiver) IncPtr() {
	receiver.count++
}


func (receiver Receiver) IncVal() {
}