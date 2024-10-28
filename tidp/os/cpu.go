package os

import (
	cpu "github.com/klauspost/cpuid/v2"
)

func CPUCores() int {
	return cpu.CPU.LogicalCores
}

func CPUBrandName() string {
	return cpu.CPU.BrandName
}

func CPUFrequency() int64 {
	return cpu.CPU.Hz
}
