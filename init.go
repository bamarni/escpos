package escpos

// [Name]		Initialize printer
// [Format]
//			ASCII	   	ESC	  	@
//			Hex		1B		40
//			Decimal		27		64
// [Description]	Clears the data in the print buffer and resets the printer modes to the modes that were in effect when the power was turned on.
// 			Any macro definitions are not cleared.
// 			Offline response selection is not cleared.
// 			Contents of user NV memory are not cleared.
// 			NV graphics (NV bit image) and NV user memory are not cleared.
// 			The maintenance counter value is not affected by this command.
// 			The specifying of offline response isn't cleared.

func (e *Escpos) Init() error {
	return e.write([]byte{27, 64})
}
