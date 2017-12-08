package main

import "fmt"

const MAX  int=3
func main()  {
	test2()
	//a :=[]int{10,100,200}
	//var i int
	//var ptr [MAX]*int;
	//for i=0;i<MAX;i++{
	//	ptr[i] = &a[i] /*整数地址赋值给指针数组*/
	//}
	//for i=0;i<MAX;i++{
	//	fmt.Printf("a[%d] = %d\n",i,*ptr[i])
	//}
	//var numbers  = make([]int,3,5)
	//fmt.Println(numbers)
	//hello(numbers)
	//var a int=20
	//var ip *int
	//ip= &a
	//
	//fmt.Printf("a 变量的地址是：%x\n",&a)
	//fmt.Printf("ip 变量存储的指针的地址：%d\n", *ip)
	//
}
func hello(x [] int)  {
	fmt.Printf("len = %d, cap = %d, slice=%v\n",len(x),cap(x),x)
}
func test()  {
	numbers :=[]int{0,1,2,3,4,5,6,7,8}
	hello(numbers)
	/*打印原始切片*/
	fmt.Println("numbers ==",numbers)
	/**/
}
func test2()  {
	var numbers [] int
	hello(numbers)
	/*允许追加空切片*/
	numbers = append(numbers,0)
	hello(numbers)
	/*向切片添加一个元素*/
	numbers = append(numbers,1)
	hello(numbers)
	/*同时添加多个元素*/
	numbers = append(numbers,2,3,4)
	hello(numbers)

	/*创建切片numbers1是之前切片容量的两倍*/
	numbers1 := make([]int,len(numbers),(cap(numbers))*2)

	/*拷贝number到number1*/
	copy(numbers1,numbers)
	hello(numbers1)
}
