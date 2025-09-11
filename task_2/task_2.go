package main	
import (
	"time"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)
///////指针/////////
// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// // 考察点 ：指针运算、切片操作。

func increase(num *int) {	
	*num += 10	
}

func doubleSlice(num *[]int) {
	for i := range *num {	
        (*num)[i] *= 2	
    }	
}

///////Goroutine/////////
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。

func oddOne() {
	for i := 1; i <= 10; i ++ {
		if i%2==1{
			fmt.Println("One:",i)
		}
	}
}

func oddTwo() {
	for i := 1; i <= 10; i ++ {
		if i%2==0{
			fmt.Println("Two:",i)
		}
	}
}

func	Renwu1(){
	start := time.Now()
	t:= 0
	for i := 0; i < 10000; i++ {
		t+=i * i
		if i%1000==0 {

			fmt.Printf("Execution 1 time:%v\n",i)
		}
	}
	end := time.Now() 
	fmt.Printf("sum = %d Execution 1 time:%v\n", t, end.Sub(start))
}

func	Renwu2(){
	start := time.Now()
	t:= 0
	for i := 0; i < 10000; i++ {
		t+=i +i+43*2
		if i%1000==0 {
			fmt.Printf("Execution 2 time:%v\n",i)
		}
	}
	end := time.Now() 
	fmt.Printf("sum = %d Execution 2 time:%v\n", t, end.Sub(start))
}

func	Renwu3(){
	start := time.Now()
	t:= 0
	for i := 1; i < 10000; i++ {
		t += i/i*23
		if i%1000==0 {
			fmt.Printf("Execution 3 time:%v\n",i)
		}
	}
	end := time.Now() 
	fmt.Printf("sum = %d Execution 3 time:%v\n", t, end.Sub(start))
}

//////面向对象/////////
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// // 考察点 ：组合的使用、方法接收者。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64{
	fmt.Println("Rectangle area:", r.width * r.height)
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64{
	fmt.Println("Rectangle perimeter:", 2 * (r.width + r.height))
	return 2 * (r.width + r.height)
}

func (c *Circle) Area() float64{
	fmt.Println("Circle area:", math.Pi * c.radius * c.radius)
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64{
	fmt.Println("Circle perimeter:", 2 * math.Pi * c.radius)
	return 2 * math.Pi * c.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}

///Channel/////////
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，
// 另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func producer(ch chan<- int) {
		for i:=0; i<=10; i++ {	
		fmt.Println("yifasong:", i)
		ch <- i		
	}
	close(ch)
}
func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("jieshou:", v)
	}
}

func pro(ch chan<- int){
	for i:=1; i<=100; i++ {
		fmt.Println("已发送:", i)
		ch <- i
	}
	close(ch)
}

///////锁机制/////////
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

type Counter struct {
	mutex sync.Mutex
	tag int 
}

func (c *Counter)Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.tag++
}

func main() {	
    fmt.Println("Hello, world!")

	// task_2:指针1
	var n int = 13
	num := &n
	increase(num)
	fmt.Println("Before increase:", *num)
    // task_2:指针2
    var slice []int = []int{1, 2, 3, 4, 5}
    doubleSlice(&slice)
	fmt.Println("After double:", slice)
	//
    


	// task_2:Goroutine1
	go oddOne()	
	go oddTwo()
	// 等待一段时间以确保协程完成
	time.Sleep(1 * time.Second)
   // task_2:Goroutine2
	func(){
		go Renwu1()
		go Renwu2()
		go Renwu3()
	}()
    // 等待一段时间以确保协程完成
	time.Sleep(3 * time.Second)



	// task_2:面向对象1
	Rectangle1 := &Rectangle{10, 20}
    Circle1 := &Circle{5}
	var s Shape = Rectangle1
	s.Area()
	s.Perimeter()
	s = Circle1
	s.Area()
	s.Perimeter()
	// task_2:面向对象2
	emp := &Employee{Person{"Tom", 25}, 1001}
	emp.PrintInfo()

    // task_2:Channel1
	var ch11 = make(chan int, 4)
	go producer(ch11)
	go consumer(ch11)
	// 等待一段时间以确保协程完成
	time.Sleep(3 * time.Second)
	// task_2:Channel2
	var ch21 = make(chan int, 10)
	go pro(ch21)
    func(){
		for {
            select {
				case v,ok := <-ch21:
					if ok == false {
						fmt.Println("通道关闭")
						return
					}
					fmt.Println("已接收:", v)
				case <-time.After(2 * time.Second):
					fmt.Println("超时")
					return
				default:
					fmt.Println("等待...")
		            time.Sleep(200*time.Millisecond)			
			}
		}
	}()



	// task_2:锁机制1
	var coun Counter
	for i:=0; i<10; i++ {
	     go func(){
             for j:=0; j<1000; j++ {
			     coun.Inc()
			 }
		 }()
		 
	}
	time.Sleep(100*time.Millisecond)
	fmt.Println("tag:", coun.tag)

	// task_2:锁机制2
	var wg sync.WaitGroup
	numWorker := 10	// 并发工作者数量
	wg.Add(numWorker)
	var getn int64 = 0
	for i:=0; i<10; i++ {
	     go func(){
			 defer wg.Done()
             for j:=0; j<1000; j++ {
			     atomic.AddInt64(&getn, 2)
			 }
		}()
    } 
	wg.Wait()
	//time.Sleep(100*time.Millisecond)
	fmt.Println("getn:", getn)
}