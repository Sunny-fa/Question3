package main

import "fmt"

type character struct{
	name string
	HP float32
	MP float32
	CE float32
}

func round(x,y int)(int,int) {
	var A character
	A.name = "萌新"
	A.HP = 100
	A.MP = 0
	A.CE = 10
	var B character
	B.name = "菜鸡"
	B.HP = 300
	B.MP = 50
	B.CE = 20
	for i := 1; i <= 5; i++ {
		B.HP = B.HP - A.CE
		for random() == 0 {
			B.HP = B.HP - A.CE
		}
		A.HP = A.HP - B.CE
		B.MP = B.MP + 10
		if B.MP >= 50 {
			A.CE = A.CE * 0.9
		}
		if A.HP <= 0 {
			y = y + 1
			break
		} else if B.HP <= 0 {
			x = x + 1
			break
		}
	}
	if A.HP > 0 && B.HP > 0 {
		for i := 1; i <= 5; i++ {
			A.HP = A.HP - 20
			B.MP = B.MP + 10
			if B.MP >= 50 {
				A.CE = A.CE * 0.9
			}
			B.HP = B.HP - A.CE
			for random() == 0 {
				B.HP = B.HP - A.CE
			}
			if A.HP <= 0 {
				y = y + 1
				break
			} else if B.HP <= 0 {
				x = x + 1
				break
			}

		}
	}
	return x,y
}
func random() int {
	ch:=make(chan int,1)
	for{
		select{
		case ch <- 0:
		case ch <- 1:
		}
		i:=<-ch
		return i

	}

}
func main(){
	var winA=0
	var winB=0

	for i:=1;i<=10;i++{

		winA,winB=round(winA,winB)

	}
	winA=winA/10
	winB=winB/10
	fmt.Println("萌新的胜率：",winA)
	fmt.Println("菜鸡的胜率：",winB)
}