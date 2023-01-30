package main

import "16.Mediator/mediator"

func main() {
	//创建一个具体中介者
	tMediator := mediator.UnitedNationsSecurityCouncil{}
	//创建具体同事,并且让他认识中介者
	tColleageA := mediator.USA{
		UnitedNations: tMediator,
	}
	tColleageB := mediator.Iraq{
		UnitedNations: tMediator,
	}
	//让中介者认识每一个具体同事
	tMediator.USA = tColleageA
	tMediator.Iraq = tColleageB
	//A同事发送消息
	tColleageA.SendMessage("停止核武器研发，否则发动战争")
	tColleageB.SendMessage("我们没有研发核武器，也不怕战争")
}
