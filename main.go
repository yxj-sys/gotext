package main

import "fmt"
	
func main(){
	//实现判断一个整数，属于哪个范围
	var num int
	fmt.Printf("请输入一个整数：")
	fmt.Scanln(&num)
	if num > 0 {
		fmt.Printf("%d大于0\n",num)
	}else if num < 0 {
		fmt.Printf("%d小于0\n",num)
	} else {
		fmt.Printf("%d等于0\n",num)
	}


	//判断一个年份是否为闰年
	var year int
	fmt.Printf("请输入您要查询的年份：")
	fmt.Scanln(&year)
	if year % 400 == 0 && year % 4 == 0 || year % 100 != 0{
		fmt.Printf("%d 是闰年",year)
	} else {
		fmt.Printf("%d 不是闰年",year)
	}
	
	//判断一个整数是否为水仙花数
	var nums int
	//读取数值
	fmt.Printf("请输入一个整数：")
	fmt.Scanln(&nums)
	//提取该数据第一位数字
	var i = nums /100
	//提取该数据第二位数字
	var j = (nums-i*100)/10
	//提取该数据第三位数字
	var k =	nums-(i * 100 + j * 10)
	//套入水仙花数公式
	if nums ==(i*i*i)+(j*j*j)+(k*k*k) && nums != 0 {
		fmt.Printf("%d是水仙花数",nums)
	}else {
		fmt.Printf("%d不是水仙花数",nums)
	}



	//保存用户名和密码，判断用户名是否为“张三”，密码为“1234”。如果是，提示登录成功，否则登录失败
	var admin string
	var password int
	fmt.Println("请输入账号名")
	fmt.Scanln(&admin)
	fmt.Println("请输入密码")
	fmt.Scanln(&password)
	if admin == "张三" && password == 1234 {
		fmt.Printf("%v,恭喜您登录成功！",admin)
	} else {
		fmt.Print("登录失败！账号或密码错误！")
	}

	//编写一个程序，根据输入的月份和年份，求出该月的天数（1-12）
	var years int
	var month int
	fmt.Printf("请输入您要查询的年份：")
	fmt.Scanln(&years)
	fmt.Printf("请输入您要查询的月份：")
	fmt.Scanln(&month)
		switch month {
			case 1,3,5,7,8,10,12 :
				fmt.Printf("%d年%d月这个月一共有31天！",years,month)
			case 2 :
				fmt.Printf("%d年%d月这个月一共有28/29天！",years,month)
			default :
				fmt.Printf("%d年%d月这个月一共有30天！",years,month)
		}
	


	//开发一款软件，根据公式（身高-108）*2=体重，可以有上下10斤浮动，开观察测试者是否合适。
	var weight float64
	var height float64
	fmt.Printf("请输入您的身高：")
	fmt.Scanln(&weight)
	fmt.Printf("请输入您的体重：")
	fmt.Scanln(&height)
	var predict	= (weight - 108) * 2
	if height - predict >= 10 || predict - height >= 10 {
		fmt.Printf("您的体重不在标准范围之内")
	} else {
		fmt.Printf("您的体重符合科学体重哦~")
	}


	//判断考试成绩为什么等级，90-100之间为优秀；80-89之间优良；70-79之间为良好；60-69之间为及格；60以下为不及格。
	var mark byte
	fmt.Printf("请输入学生成绩：")
	fmt.Scanln(&mark)
	if mark >= 90 {
		fmt.Printf("成绩为%d,优秀！",mark)
	}else if mark >= 80 {
		fmt.Printf("成绩为%d,优良！",mark)
	}else if mark >= 70 {
		fmt.Printf("成绩为%d,良好！",mark)
	}else if mark >= 60 {
		fmt.Printf("成绩为%d,及格！",mark)	
	}else {
		fmt.Printf("成绩为%d,不及格！",mark)
	}	 

}