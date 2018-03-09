package main

/*

간단한 switch문
바로 코딩하려는데 잠시 생각하게 만들어서 작성 해봄.
go는 여러개의 조건문일 때 if 보다는 switch문을 선호한다. 

 */
func main() {

	var name string
	var category = 2

	switch category {
	case 1:
		name = "taehapark"
	case 2:
		name ="kimmiae"
	case 3:
		name = "ace-t"
	default:
		name = "defaultName"
	}

	println(name )
}
