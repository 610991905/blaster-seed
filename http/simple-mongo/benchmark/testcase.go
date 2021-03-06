package benchmark

import (
	"github.com/lordking/blaster/log"
)

//TestCase 测试用例定义
type TestCase struct {
	BaseURL   string
	RandLimit int
	TesterNO  int
	Count     int
}

//GetTestRun 实现接口
func (t *TestCase) GetTestRun(no, cnt int) {
	log.Debugf("Tester No: %d\n", no)
	log.Debugf("Count: %d\n", cnt)

	t.TesterNO = no
	t.Count = cnt
}

//GetCompletedSynchrony 实现接口
func (t *TestCase) GetCompletedSynchrony(x int) {

	log.Debugf("Compleate test: %d\n", x)
}
