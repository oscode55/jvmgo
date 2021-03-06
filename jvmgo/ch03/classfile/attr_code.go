/*
* @Author: myname
* @Date:   2018-09-13 00:00:28
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 17:44:36
 */
/*
    Code变长属性 只存在于method_info结构中
    Code_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 max_stackl
	    u2 max_locals;
	    u4 code_length;
	    u1 code[code_length];
	    u2 exception_table_length;
	    //字节码start_pc行到end_pc间出现异常catch_type,则跳到handler_pc行继续执行
	    {   u2 start_pc;//指令范围1
	    	u2 end_pc;//指令范围2
	    	u2 handler_pc;//指令范围3
	    	u2 catch_type;//异常类型1
	    } exception_table[exception_table_length];
	    u2 attributes_count;
	    attribute_info attributes[attributes_count];
    }
*/
package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
