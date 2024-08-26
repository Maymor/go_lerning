	package main
	import (
		"fmt"
		"math"
	)
	
	//  
	
	func main() {
		var a float64
		var b float64
		var operator string
		fmt.Println("Вы используете калькулятор. Для вычислений введите 2 нужных числа и выберите операцию.")
		fmt.Println("Введите первое число:")
		for {
			_, err1 := fmt.Scan(&a)
			if err1 == nil {
				break
			} else {
				fmt.Println("Неверно. Введите число:")
			}
		}
			fmt.Println("Введите второе число:")
		for {
			_, err2 := fmt.Scan(&b)
			if err2 == nil {
				break
			} else {
				fmt.Println("Неверно. Введите число:")
				fmt.Println(err2)
			}
		}
		return
		fmt.Println("Для выбора предложены следующие операции: '+' - сложение, '-' - вычитание, '*' - умножение, '/' - деление, '^' - возведение в степень. Выберите операцию:")
		fmt.Scanln(&operator)
		switch operator {
		case "+":
			fmt.Println("Результат: ", a+b)
		case "-":
			fmt.Println("Результат: ", a-b)
		case "*":
			fmt.Println("Результат: ", a*b)
		case "/":
			fmt.Println("Результат: ", a/b)
		case "^":
			fmt.Println("Результат: ", math.Pow(a, b))
		default:
			fmt.Println("К сожалениею, Вам не хватило умственных способностей, чтобы последовать оставленной инструкции, и Вы неверно выбрали операцию. Выполните команду запуска калькулятора снова для повторения попытки.")
	}
}
		