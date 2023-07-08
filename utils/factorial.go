package utils

func Factorial(num int64) (result int64){
	if num == 1 {
		return 1
	}
	result = Factorial(num - 1) * num
	return
}
