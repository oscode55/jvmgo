/*
* @Author: myname
* @Date:   2018-10-19 16:50:51
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:27:36
 */
package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperandsInstruction }

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
