package main

import (
	"customermanage/model"
	"customermanage/service"
	"fmt"
)

type customerView struct {
	key             string                   // 接受用户输入
	loop            bool                     // 表示是否循环显示主菜单
	customerService *service.CustomerService // 添加一个字段customerService
}

// 显示所有客户信息
func (this *customerView) list() {
	// 首先，获取当前所有客户的信息（在切片中）
	customers := this.customerService.List()

	// 显示每个用户信息
	fmt.Println("----------------------------客户列表----------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n---------------------------客户列表完成-------------------------\n\n")
}

// 得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("--------------------添加客户-------------------")
	fmt.Println("姓名：")
	name := " "
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := " "
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := " "
	fmt.Scanln(&phone)
	fmt.Println("邮件：")
	email := " "
	fmt.Scanln(&email)
	// 构建一个新的Customer实例
	// 注意：id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("------------------------添加完成---------------------")
	} else {
		fmt.Println("------------------------添加失败---------------------")
	}

}

// 得到用户输入的id 删除该id的用户
func (this *customerView) delete() {
	fmt.Println("------------------------删除客户---------------------")
	fmt.Println("请选择待删除客户的编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		// 这里return 结束函数
		return
	}
	fmt.Println("你是否要真的删除id:%v客户(y/n)?", id)
	for {
		choice := " "
		fmt.Scanln(&choice)
		if choice == "n" || choice == "N" {
			// 这里return 结束函数
			return
		} else {
			if this.customerService.Delete(id) {
				fmt.Println("------------------------删除成功---------------------")
				fmt.Println()
				fmt.Println()
				return
			} else {
				fmt.Println("------------------------删除失败---------------------")
				return
			}
		}
	}

}

// 退出系统函数
func (this *customerView) exit() {
	fmt.Println("你确定要退出系统么？y/n")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出（y/n）:")
	}
	if this.key == "y" || this.key == "Y" {
		this.loop = false
	} else {
		fmt.Println("返回系统中...")
		fmt.Println()
	}

}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("------------------客户信息管理软件-----------------")
		fmt.Println("                    1.添加客户信息")
		fmt.Println("                    2.修改客户信息")
		fmt.Println("                    3.删除客户信息")
		fmt.Println("                    4.客户列表信息")
		fmt.Println("                    5.退        出")
		fmt.Println("请选择（1-5）：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户信息")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入（1-5）：")
		}
		if !this.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统....")
}

func main() {
	// 实例化customerview结构体
	customerView := customerView{
		key:  " ",
		loop: true,
	}
	// 对customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	// 显示主菜单
	customerView.mainMenu()
}
