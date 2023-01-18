package file_process

import (
	//monkey是一个Go单元测试中十分常用的打桩工具，它在运行时通过汇编语言重写可执行文件，将目标函数或方法的实现跳转到桩实现，其原理类似于热补丁。
	//monkey库很强大，但是使用时需注意以下事项：
	//monkey不支持内联函数，在测试的时候需要通过命令行参数-gcflags=-l关闭Go语言的内联优化。
	//monkey不是线程安全的，所以不要把它用到并发的单元测试中。
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}

func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)
	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}
