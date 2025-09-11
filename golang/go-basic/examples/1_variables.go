package examples

import "fmt"

func Variables() {
	fmt.Println("--- Variables examples started ---")

	// IMPORTANT: Go 把變數型別放在變數名後面

	// --- 1. 基本宣告方式 ---
	var text string
	var number int

	// --- 2. 多變數宣告 ---
	var str1, str2, str3 string
	var num1, num2, num3 int

	// --- 3. 宣告變數並初始化 ---
	var greeting string = "Hello"

	// --- 4. 同時初始化多個變數 ---
	var letter1, letter2, letter3 string = "a", "b", "c"

	// --- 5. 省略型別宣告 ---
	var hi = "Hi"
	var letter4, letter5, letter6 = "d", "e", "f"

	// --- 6. 簡化宣告方式 ---

	/*
		簡短宣告用 := 取代 var 和 型別宣告
		但是只能在函式內使用，外部使用則會無法編譯
		所以全域變數使用 var 宣告
	*/
	golang := "Golang"
	letter7, letter8, letter9 := "g", "h", "i"

	// --- 7. 丟棄變數 ---

	/*
		_ (下劃線) 表示丟棄變數，任何賦予它的值都會被丟棄
		在這個例子中，我們將值 35 賦予b，並同時丟棄34：
	*/
	_, b := 34, 35

	// --- 8. 宣告但未使用的變數 ---
	// Golang 對於已宣告但未使用的變數會在編譯階段報錯
	// var i int

	// --- Output ---
	fmt.Println("1. text:", text, "number:", number)
	fmt.Println("2. str1:", str1, "str2:", str2, "str3:", str3, "num1:", num1, "num2:", num2, "num3:", num3)
	fmt.Println("3. greeting:", greeting)
	fmt.Println("4. letter1:", letter1, "letter2:", letter2, "letter3:", letter3)
	fmt.Println("5. hi:", hi, "letter4:", letter4, "letter5:", letter5, "letter6:", letter6)
	fmt.Println("6. golang:", golang, "letter7:", letter7, "letter8:", letter8, "letter9:", letter9)
	fmt.Println("7. b:", b)

	fmt.Println("--- Variables examples finished ---")
}
