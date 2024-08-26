	package main
	import (
		"time"
		"fmt"
		"math"
	)
	
	//  
	
	func main() {
		var a float64
		var b float64
		var operator string
		fmt.Println("Вы используете калькулятор. Для вычислений введите 2 нужных числа и выберите операцию.")
		FAILCASE1:
		fmt.Println("Введите первое число:")
		_, err1 := fmt.Scanln(&a)
		if err1 != nil {
			fmt.Println("К сожалениею, Вам не хватило умственных способностей, чтобы последовать оставленной инструкции, и Вы неверно ввели число. Попробуйте снова.")
			goto FAILCASE1
		}
		FAILCASE2:
		fmt.Println("Введите второе число:")
		_, err2 := fmt.Scanln(&b)
		if err2 != nil {
			fmt.Println("К сожалениею, Вам не хватило умственных способностей, чтобы последовать оставленной инструкции, и Вы неверно ввели число. Попробуйте снова.")
			goto FAILCASE2
		}
		FAILCASE3:
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
				goto FAILCASE3

		}
		time.Sleep(8 * time.Second)
	}
	
