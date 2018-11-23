package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"peopleInfo"
	"strconv"
	"strings"
)

/***************************************
 * 判错误函数
 * 函数功能：判断操作返回的错误码，若有错误则退出程序
 * 函数名：ifError
 * 参数：错误码
 * 返回值：无
***************************************/
func ifError(err error) {
	if err != nil {
		panic("\n" + err.Error())
	}
}

/***************************************
 * 判判断个人信息是否正确的函数
 * 函数功能：判断所增加或者修改时输入的个人信息是否合理
 * 函数名：ifInfoRight
 * 参数：个人信息结构体
 * 返回值：bool值， true表示正确，false表示错误
***************************************/
func ifInfoRight(p peopleInfo.People_info) (ron bool) {
	if p.Id_ <= 0 {
		return false
	}
	if p.Sex_ != 1 && p.Sex_ != 2 {
		return false
	}
	if p.Age_ <= 0 || p.Age_ > 200 {
		return false
	}
	if p.Height_ <= 0 || p.Height_ > 400 {
		return false
	}
	if p.Weight_ <= 0 || p.Weight_ > 300 {
		return false
	}
	return true
}

//程序主函数
func main() {
	var pes peopleInfo.People_infos
	filePath := "../people_info.txt" //文件的存放路径，相对路径
	err1 := pes.Init(filePath)       //初始化函数，即从文件中读取信息解析到结构体中
	ifError(err1)
	defer pes.SavePeopleInfoIntoTxt(filePath) //程序的最后需要把结构体写回到文件中进行保存

	//提示操作命令
	fmt.Println("Enter following commands to operate the people info database:\n" +
		"list   -- View the existing people info\n" +
		"add <id><name><sex><age><height><weight> -- Add a people to database\n" +
		"remove <id> -- Remove the specified people from database\n" +
		"revise <id> -- revise the specified people info in database\n" +
		"find <i or n><id or name> -- Find the specified people from database\n" +
		"q or e --Exit the system")

	r := bufio.NewReader(os.Stdin) //从标准输入中获取字符串
	for {
		var pe peopleInfo.People_info
		var err error
		fmt.Print("Enter command -> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" { //退出程序
			break
		}
		tokens := strings.Split(line, " ") //切割字符串获取各个参数
		switch tokens[0] {
		case "list": //列出所有数据
			{
				for _, pe := range pes.Peoples {
					fmt.Println(pe)
				}
			}
		case "add": //增
			{
				if len(tokens) == 7 {
					pe.Id_, err = strconv.Atoi(tokens[1])
					pe.Name_ = tokens[2]
					pe.Sex_, err = strconv.Atoi(tokens[3])
					pe.Age_, err = strconv.Atoi(tokens[4])
					pe.Height_, err = strconv.Atoi(tokens[5])
					pe.Weight_, err = strconv.Atoi(tokens[6])
					if !ifInfoRight(pe) {
						err = errors.New("People info input error.")
						break
					}
					err = pes.AddOnePeople(pe)
				} else {
					err = errors.New("Command error.")
				}
			}
		case "remove": //删
			{
				if len(tokens) == 2 {
					var id int
					id, err = strconv.Atoi(tokens[1])
					err = pes.RemoveOnePeople(id)
				} else {
					err = errors.New("Command error.")
				}
			}
		case "revise": //改
			{
				if len(tokens) == 7 {
					pe.Id_, err = strconv.Atoi(tokens[1])
					pe.Name_ = tokens[2]
					pe.Sex_, err = strconv.Atoi(tokens[3])
					pe.Age_, err = strconv.Atoi(tokens[4])
					pe.Height_, err = strconv.Atoi(tokens[5])
					pe.Weight_, err = strconv.Atoi(tokens[6])
					if !ifInfoRight(pe) {
						err = errors.New("People info input error.")
						break
					}
					err = pes.ReviseOnePeopleInfo(pe)
				} else {
					err = errors.New("Command error.")
				}
			}
		case "find": //查
			{
				if len(tokens) == 3 {
					if tokens[1] == "-i" {
						var id int
						id, err = strconv.Atoi(tokens[2])
						pe, err = pes.GetPepleInfoById(id)
						if err == nil {
							fmt.Println(pe)
						}
					} else if tokens[1] == "-n" {
						var ps peopleInfo.People_infos
						ps, err = pes.GetPeopleInfoByName(tokens[2])
						if err == nil {
							for _, op := range ps.Peoples {
								fmt.Println(op)
							}
						}
					} else {
						err = errors.New("Command error.")
					}
				} else {
					err = errors.New("Command error.")
				}
			}
		default:
			{
				err = errors.New("Command error.")
			}
		}
		if err != nil {
			fmt.Println(err)
		}
	}
	return
}
