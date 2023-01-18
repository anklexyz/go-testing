//单元测试-覆盖率

package unittest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 单元测试
// 要以Test开头
func TestJudgePassLine(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}
