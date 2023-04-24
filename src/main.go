// main.go

package main

import (
    "strconv"
    "encoding/json"
    "github.com/astaxie/beego"
    "math/rand"
	//"time"
    "fmt"
    "log"
    "os"
    "net/http"
    _ "net/http/pprof"
)

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    /* This would match routes like the following:
       /sum/3/5
       /product/6/23
       ...
    */
    
    beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
    go func() {
            log.Println(http.ListenAndServe("localhost:6060", nil))
        }()
    beego.Run()
}

type mainController struct {
    beego.Controller
}


func (c *mainController) Get() {

    //Obtain the values of the route parameters defined in the route above    
    operation := c.Ctx.Input.Param(":operation")
    num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
    num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

    fmt.Print(rand.Float64() * rand.Float64())
    /*v := RandomSequence(2,502)
    fmt.Printf("size:%d\n", len(v))
    //sort.Ints(v)
    v = bubbleSort(v)
    fmt.Printf("size:%d\n", len(v))*/

    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)







    //Set the values for use in the template
    c.Data["operation"] = operation
    c.Data["num1"] = num1
    c.Data["num2"] = num2
    c.TplName = "result.html"

    // Perform the calculation depending on the 'operation' route parameter
    switch operation {
    case "sum":
        c.Data["result"] = add(num1, num2)
    case "product":
        c.Data["result"] = multiply(num1, num2)
    default:
        c.TplName = "invalid-route.html"
    }
}

func add(n1, n2 int) int {
    return n1 + n2
}

func multiply(n1, n2 int) int {
    return n1 * n2
}

/*
func RandomSequence(min,max int) []int{
	//计算序列的长度
	lenghth := max - min + 1
	
	//初始化一个长度为lenghth的原始切片，初始值从min到max
	initArr := make([]int,lenghth)
	for i :=0; i < lenghth; i++{
		initArr[i] = i + min
	}
 
	//初始化一个长度为lenghth的目标切片
	rtnArr := make([]int,lenghth)
 
	//初始化随机种子
	rand.Seed(time.Now().Unix())
	
	//生成目标序列
	for i :=0; i < lenghth; i++{
		//生成一个随机序号
		index := rand.Intn(lenghth - i)
		
		//将原始切片中序号index对应的值赋给目标切片
		rtnArr[i] = initArr[index]
		
		//替换掉原始切片中使用过的下标index对应的值
		initArr[index] = initArr[lenghth - i - 1]
	}
 
	return rtnArr
}

func bubbleSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }
    // 冒泡排序核心实现代码
    for i := 0; i < len(nums); i++ {
        flag := false
        for j := 0; j < len(nums) - i - 1; j++ {
            if nums[j] > nums[j+1] {
                nums[j], nums[j+1] = nums[j+1], nums[j]
                flag = true
            }
        }
        if !flag {
            break
        }
    }
    return nums
}
*/