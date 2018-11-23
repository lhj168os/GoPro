package peopleInfo

import (
	"testing"
)

func TestOps(t *testing.T) {
	filePath := "people_info.txt"
	var p1 People_info
	err1 := p1.GetPepleCfgById(18, filePath)
	if err1 != nil {
		t.Error("GetPepleCfgById failed.")
	}

	var p2 People_infos
	err2 := p2.GetPepleInfoFromeTxt(filePath)
	if err2 != nil {
		t.Error("GetPepleInfoFromeTxt failed.")
	}

	var p3 People_infos
	err3 := p3.Init(filePath)
	if err3 != nil {
		t.Error("Init failed.")
	} else {
		_, err4 := p3.GetPepleInfoById(18)
		if err4 != nil {
			t.Error("GetPepleInfoById failed.")
		}

		_, err5 := p3.GetPeopleInfoByName("jay")
		if err5 != nil {
			t.Error("GetPeopleInfoByName failed.")
		}

		var p4 People_info
		p4.Id_ = 1001
		p4.Age_ = 25
		p4.Height_ = 180
		p4.Name_ = "ABC"
		p4.Sex_ = 1
		p4.Weight_ = 55
		p3.AddOnePeople(p4)
		_, err6 := p3.GetPepleInfoById(p4.Id_)
		if err6 != nil {
			t.Error("AddOnePeople or GetPepleInfoById failed.")
		}

		p4.Name_ = "acc"
		err7 := p3.ReviseOnePeopleInfo(p4)
		if err7 != nil {
			t.Error("ReviseOnePeopleInfo or AddOnePeople failed.")
		}
		p5, err8 := p3.GetPepleInfoById(p4.Id_)
		if err8 != nil || p5.Name_ != p4.Name_ {
			t.Error("ReviseOnePeopleInfo or GetPepleInfoById failed.")
		}

		p3.RemoveOnePeople(p4.Id_)
		_, err9 := p3.GetPepleInfoById(p4.Id_)
		if err9 == nil {
			t.Error("RemoveOnePeople or GetPepleInfoById failed.")
		}

		err10 := p3.SavePeopleInfoIntoTxt(filePath)
		if err10 != nil {
			t.Error("SavePeopleInfoIntoTxt")
		}
	}

}
