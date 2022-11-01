package service

import (
	"customermanage/model"
)

// CustomerService完成对customer的操作（增删改查）
type CustomerService struct {
	customers []model.Customer

	// 声明一个字段，表示当前切片含有多少个客户
	// 该字段后面还可以用作新客户的id+1
	customerNum int
}

// 编写一个方法，返回*CustomerService
func NewCustomerService() *CustomerService {
	// 为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}

	// 注意细节  要使用CustomerService结构体的实例对象去调用结构体内部的属性、方法
	customerService.customerNum = 1

	// 创建结构体实例customer
	customer := model.NewCustomer(1, "张三", "男", 20, "15349840707", "zhangsan@.com")

	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到customers切片
// 一定使用*CustomerService  否则customer覆盖，永远只有一个
func (this *CustomerService) Add(customer model.Customer) bool {

	// 我们确定一个分配id的规则，就是添加的顺序
	this.customerNum++
	customer.Id = this.customerNum

	this.customers = append(this.customers, customer)
	return true
}

// 根据id在切片中删除客户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	// 下标-1  表示没有这个客户
	if index == -1 {
		return false
	}
	// 如何在切片中删除一个元素
	this.customers = append(this.customers[0:index], this.customers[index+1:]...)
	return true
}

// 根据id查找客户在切片中对应的下标，如果没有该客户，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for _, v := range this.customers {
		if v.Id == id {
			// 找到下标
			index = id
		}
	}
	return index
}

// 根据id修改客户信息  id不可改
