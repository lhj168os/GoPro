/*********************************
 * File： peopleInfo.go
 * Author： jay
 * Date： 2018-10-13
**********************************/
package peopleInfo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//个人信息结构体
type People_info struct {
	Id_     int    //个人id
	Name_   string //个人姓名
	Sex_    int    //性别
	Age_    int    //年龄
	Height_ int    //身高（CM）
	Weight_ int    //体重（KG）
}

/***************************************
 * People_info类型的方法
 * 函数功能：从文件中数据中获取指定id的个人信息到People_info中
 * 函数名：GetPepleCfgById
 * 参数：个人id， 文件路径
 * 返回值：错误码
***************************************/
func (p *People_info) GetPepleCfgById(id int, filePath string) (err error) {
	var peoples People_infos
	peoples.GetPepleInfoFromeTxt(filePath)
	*p, err = peoples.GetPepleInfoById(id)
	return
}

//所有人的数据存储结构
type People_infos struct {
	Peoples   []People_info //人物列表
	PeopleNum int           //人数
}

/***************************************
 * People_infos类型的方法
 * 函数功能：读取文件中所有的信息存进结构体
 * 函数名：GetPepleInfoFromeTxt
 * 参数：文件路径
 * 返回值：错误码
***************************************/
func (p *People_infos) GetPepleInfoFromeTxt(filePath string) (err error) {
	var people1 People_info
	file, err1 := os.Open(filePath) //文件操作，打开文件
	if err1 != nil {
		err = err1 //文件打开出错
		return
	}
	defer file.Close() //函数执行的最后关闭文件

	br := bufio.NewReader(file)
	p.PeopleNum = 0
	for {
		line, isPrefix, err1 := br.ReadLine() //按行读取

		if err1 != nil {
			if err1 != io.EOF {
				err = err1 //读取出错
			}
			//读取到文件最后
			break
		}

		if isPrefix { //一行数据字节太长
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		if p.PeopleNum == 0 {
			p.PeopleNum++ //统计人数
			continue
		}
		str := strings.Fields(string(line)) //去掉空白字符切割成字符串数组
		//对结构体赋值
		people1.Id_, err = strconv.Atoi(str[0])
		people1.Name_ = str[1]
		people1.Sex_, err = strconv.Atoi(str[2])
		people1.Age_, err = strconv.Atoi(str[3])
		people1.Height_, err = strconv.Atoi(str[4])
		people1.Weight_, err = strconv.Atoi(str[5])
		p.Peoples = append(p.Peoples, people1)
		p.PeopleNum++ //统计人数
	}
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：将结构体数组数据全部写入文件中
 * 函数名：SavePeopleInfoIntoTxt
 * 参数：文件路径
 * 返回值：错误码
***************************************/
func (p *People_infos) SavePeopleInfoIntoTxt(filePath string) (err error) {
	file, err := os.Create(filePath) //创建文件
	if err != nil {
		fmt.Println("Failed to open the output file ", filePath)
		return
	}
	defer file.Close()
	tableHead := "ID\t姓名\t性别(0男，1女)\t年龄\t身高(cm)\t体重（kg）"
	_, err = file.WriteString(tableHead + "\n") //按行把数据写进文件
	if err != nil {
		return
	}
	for _, peo := range p.Peoples {
		str := strconv.Itoa(peo.Id_) + "\t" + peo.Name_ + "\t" + strconv.Itoa(peo.Sex_) + "\t" + strconv.Itoa(peo.Age_) + "\t" + strconv.Itoa(peo.Height_) + "\t" + strconv.Itoa(peo.Weight_)
		_, err = file.WriteString(str + "\n") //按行把数据写进文件
		if err != nil {
			return
		}
	}
	fmt.Println("Data save successful.")
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：从结构体数组数据中查找指定id的个人信息
 * 函数名：GetPepleInfoById
 * 参数：个人Id
 * 返回值：单个人的信息People_info，错误码
***************************************/
func (p *People_infos) GetPepleInfoById(id int) (thisPeople People_info, err error) {
	for _, p := range p.Peoples {
		if p.Id_ == id {
			thisPeople = p
			return
		}
	}
	err = errors.New("Can't not find this people info where id is : " + strconv.Itoa(id))
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：从结构体数组数据中查找指定名字的个人信息
 * 函数名：GetPeopleInfoByName
 * 参数：个人名字
 * 返回值：对应名字的个人数据列表，错误码
***************************************/
func (p *People_infos) GetPeopleInfoByName(name string) (findPeoples People_infos, err error) {
	for _, p := range p.Peoples {
		if p.Name_ == name {
			findPeoples.Peoples = append(findPeoples.Peoples, p)
			findPeoples.PeopleNum++
		}
	}
	if findPeoples.PeopleNum == 0 {
		err = errors.New("Can't not find this people info where name is : " + name)
	}
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：增加一个人的信息到People_infos结构体中
 * 函数名：AddOnePeople
 * 参数：个人信息结构体
 * 返回值：错误码
***************************************/
func (p *People_infos) AddOnePeople(people People_info) (err error) {
	_, err1 := p.GetPepleInfoById(people.Id_)
	if err1 == nil {
		err = errors.New("Id: " + strconv.Itoa(people.Id_) + " already exists.")
		return
	}
	p.Peoples = append(p.Peoples, people)
	p.PeopleNum++
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：从People_infos结构体中删除一个人的信息
 * 函数名：RemoveOnePeople
 * 参数：个人id
 * 返回值：错误码
***************************************/
func (p *People_infos) RemoveOnePeople(id int) (err error) {
	for i, pe := range p.Peoples {
		if pe.Id_ == id {
			p.Peoples = append(p.Peoples[:i], p.Peoples[i+1:]...)
			return
		}
	}
	err = errors.New("Id: " + strconv.Itoa(id) + " not exists.")
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：在People_infos结构体中修改一个人的信息，其中id不能修改
 * 函数名：ReviseOnePeopleInfo
 * 参数：一个人的信息结构体
 * 返回值：错误码
***************************************/
func (p *People_infos) ReviseOnePeopleInfo(pe People_info) (err error) {
	for i, pp := range p.Peoples {
		if pp.Id_ == pe.Id_ {
			p.Peoples[i].Name_ = pe.Name_
			p.Peoples[i].Sex_ = pe.Sex_
			p.Peoples[i].Age_ = pe.Age_
			p.Peoples[i].Height_ = pe.Height_
			p.Peoples[i].Weight_ = pe.Weight_
			return
		}
	}
	err = errors.New("People not exist | id= " + strconv.Itoa(pe.Id_))
	return
}

/***************************************
 * People_infos类型的方法
 * 函数功能：该类的数据初始化
 * 函数名：Init
 * 参数：文件路径
 * 返回值：错误码
***************************************/
func (p *People_infos) Init(filePath string) (err error) {
	err1 := p.GetPepleInfoFromeTxt(filePath)
	if err1 != nil {
		err = errors.New(err1.Error() + "\nData Init failed.")
	}
	return
}
