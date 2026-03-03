package main

import (
	"errors"
	"fmt"
)

// 基础练习：学生管理系统
type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

type StudentManager struct {
	students []Student
	nextID   int
}

func NewStudentManager() *StudentManager {
	return &StudentManager{
		students: make([]Student, 0),
		nextID:   1,
	}
}

// 添加学生
func (sm *StudentManager) AddStudent(name string, age int, grade string) Student {
	student := Student{
		ID:    sm.nextID,
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	sm.students = append(sm.students, student)
	sm.nextID++
	return student
}

// 删除学生
func (sm *StudentManager) DeleteStudent(id int) error {
	for i, student := range sm.students {
		if student.ID == id {
			sm.students = append(sm.students[:i], sm.students[i+1:]...)
			return nil
		}
	}
	return errors.New("学生不存在")
}

// 更新学生信息
func (sm *StudentManager) UpdateStudent(id int, name string, age int, grade string) error {
	for i, student := range sm.students {
		if student.ID == id {
			sm.students[i] = Student{
				ID:    id,
				Name:  name,
				Age:   age,
				Grade: grade,
			}
			return nil
		}
	}
	return errors.New("学生不存在")
}

// 根据ID查询学生
func (sm *StudentManager) GetStudentByID(id int) (Student, error) {
	for _, student := range sm.students {
		if student.ID == id {
			return student, nil
		}
	}
	return Student{}, errors.New("学生不存在")
}

// 获取所有学生
func (sm *StudentManager) GetAllStudents() []Student {
	return sm.students
}

// 进阶练习：银行账户系统
type Account struct {
	AccountNumber string
	Balance       float64
}

func NewAccount(accountNumber string) *Account {
	return &Account{
		AccountNumber: accountNumber,
		Balance:       0,
	}
}

// 存款
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("存款金额必须大于0")
	}
	a.Balance += amount
	return nil
}

// 取款
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("取款金额必须大于0")
	}
	if amount > a.Balance {
		return errors.New("账户余额不足")
	}
	a.Balance -= amount
	return nil
}

// 查询余额
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// 测试函数
func testStudentManager() {
	fmt.Println("=== 测试学生管理系统 ===")
	sm := NewStudentManager()

	// 添加学生
	student1 := sm.AddStudent("张三", 18, "高三")
	student2 := sm.AddStudent("李四", 17, "高二")
	student3 := sm.AddStudent("王五", 16, "高一")

	fmt.Printf("添加学生1: %v\n", student1)
	fmt.Printf("添加学生2: %v\n", student2)
	fmt.Printf("添加学生3: %v\n", student3)

	// 查询所有学生
	fmt.Println("\n所有学生:")
	for _, student := range sm.GetAllStudents() {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n", student.ID, student.Name, student.Age, student.Grade)
	}

	// 根据ID查询学生
	fmt.Println("\n根据ID查询学生:")
	student, err := sm.GetStudentByID(1)
	if err == nil {
		fmt.Printf("查询ID为1的学生: %v\n", student)
	} else {
		fmt.Printf("查询失败: %v\n", err)
	}

	// 更新学生信息
	fmt.Println("\n更新学生信息:")
	err = sm.UpdateStudent(2, "李四改", 18, "高三")
	if err == nil {
		updatedStudent, _ := sm.GetStudentByID(2)
		fmt.Printf("更新后学生2: %v\n", updatedStudent)
	} else {
		fmt.Printf("更新失败: %v\n", err)
	}

	// 删除学生
	fmt.Println("\n删除学生:")
	err = sm.DeleteStudent(3)
	if err == nil {
		fmt.Println("删除ID为3的学生成功")
	} else {
		fmt.Printf("删除失败: %v\n", err)
	}

	// 再次查询所有学生
	fmt.Println("\n删除后所有学生:")
	for _, student := range sm.GetAllStudents() {
		fmt.Printf("ID: %d, 姓名: %s, 年龄: %d, 年级: %s\n", student.ID, student.Name, student.Age, student.Grade)
	}
}

func testBankAccount() {
	fmt.Println("\n=== 测试银行账户系统 ===")
	account := NewAccount("6222021234567890123")

	// 查询初始余额
	fmt.Printf("初始余额: %.2f\n", account.GetBalance())

	// 存款
	fmt.Println("\n存款操作:")
	err := account.Deposit(1000)
	if err == nil {
		fmt.Printf("存款1000成功，当前余额: %.2f\n", account.GetBalance())
	} else {
		fmt.Printf("存款失败: %v\n", err)
	}

	// 再次存款
	err = account.Deposit(500)
	if err == nil {
		fmt.Printf("存款500成功，当前余额: %.2f\n", account.GetBalance())
	} else {
		fmt.Printf("存款失败: %v\n", err)
	}

	// 取款
	fmt.Println("\n取款操作:")
	err = account.Withdraw(800)
	if err == nil {
		fmt.Printf("取款800成功，当前余额: %.2f\n", account.GetBalance())
	} else {
		fmt.Printf("取款失败: %v\n", err)
	}

	// 测试余额不足
	err = account.Withdraw(1000)
	if err == nil {
		fmt.Printf("取款1000成功，当前余额: %.2f\n", account.GetBalance())
	} else {
		fmt.Printf("取款失败: %v\n", err)
	}

	// 查询最终余额
	fmt.Printf("\n最终余额: %.2f\n", account.GetBalance())
}

func main() {
	testStudentManager()
	testBankAccount()
}
