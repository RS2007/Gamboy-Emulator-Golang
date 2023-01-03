package cpu

import (
	"fmt"
	"gameboy/utils"
)

var REGISTERS = [...]string{"A", "F", "B", "C", "D", "E", "H", "L"}

type CPU struct {
	Ram           [1024 * 8]int8
	Registers8Bit map[string]int8
	Sp            int16 // stack pointer
	Pc            int16 // program counter
}

func (cpu *CPU) reset() error {
	utils.MemsetInt8(cpu.Ram[:], 0)
	_initaliseRegisters(cpu.Registers8Bit)
	cpu.Pc = 0
	cpu.Sp = 0
	return fmt.Errorf("Error encountered during reset")
}

func (cpu *CPU) step() error {
	return fmt.Errorf("Error while stepping")
}

func _initaliseRegisters(registerMap map[string]int8) {
	for _, registerName := range REGISTERS {
		registerMap[registerName] = 0
	}
}

func _fetchOpcode(cpu *CPU) int8 {
	return cpu.Ram[cpu.Pc+1]
}
