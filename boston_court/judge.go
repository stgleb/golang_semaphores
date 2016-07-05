package boston_court

import "fmt"

type Judge struct {
	Name string
	Court Court
}

func (judge Judge) String() string {
	return fmt.Sprintf("<Name %s>", judge.Name)
}

func NewJudge(name string, court Court) Judge {
	return Judge{
		Name: name,
		Court: court,
	}
}

func (judge Judge) Run() {
	judge.Court.Entrance <- judge
	randSleep(20, 30)
	judge.Court.JudgeOut <- judge
}
