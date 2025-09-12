package examples

import (
	"errors"
	"fmt"
)

func Basic() {
	Variables()

	Constants()

	Types()
}

// 變數範例
func Variables() {
	fmt.Println("--- Variables examples started ---")

	// * Go 把變數型別放在變數名後面

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
	* 簡短宣告用 := 取代 var 和 型別宣告
	* 但是只能在函式內使用，外部使用則會無法編譯
	* 所以全域變數使用 var 宣告
	 */
	golang := "Golang"
	letter7, letter8, letter9 := "g", "h", "i"

	// --- 7. 丟棄變數 ---

	/*
	* _ (下劃線) 表示丟棄變數，任何賦予它的值都會被丟棄
	* 在這個例子中，我們將值 35 賦予b，並同時丟棄34：
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

// 常數範例
func Constants() {
	fmt.Println("--- Constants example started ---")

	const i = 100
	const MaxThread = 10
	const Prefix = "Golang_"

	// 如果需要，也可以明確指定型別
	const Truth bool = true
	const Greeting string = "Hello"

	//* 若指定給 float32 自動縮短為 32bit，指定給 float64 自動縮短為 64bit
	const Pi float32 = 3.1415926

	// --- Output ---
	fmt.Println("i:", i, "MaxThread:", MaxThread, "Prefix:", Prefix)
	fmt.Println("Truth:", Truth, "Greeting:", Greeting, "Pi:", Pi)

	fmt.Println("--- Constants example finished ---")
}

// 基礎型別範例
func Types() {
	fmt.Println("--- Types example started ---")

	// Boolean
	// * Golang中 bool 預設值為 false
	var isActive bool                   // 全域變數宣告
	var enabled, disabled = true, false // 忽略型別宣告
	valid := false                      // 簡短宣告
	isActive = true                     // 賦值

	// Int
	/*
	* 整數型別有無符號和帶符號兩種
	* Go 同時支援 int 和uint，這兩種型別的長度相同，但具體長度取決於不同編譯器的實現
	*
	* Go 裡面也有直接定義好位數的型別：rune, int8, int16, int32, int64 和 byte, uint8, uint16, uint32, uint64
	* 其中 rune 是int32的別稱，byte是 uint8 的別稱
	 */
	// * 需要注意的一點是，這些型別的變數之間不允許互相賦值或操作，不然會在編譯時引起編譯器報錯
	// var num1 int8
	// var num2 int16
	// c := num1 + num2
	// * 另外，儘管 int 的長度是 32 bit, 但 int 與 int32 並不可以互用

	// Float
	// * 浮點數的型別有 float32 和 float64 兩種（沒有 float 型別），預設是 float64

	// Complex
	/*
	* Go 還支援複數。它的預設型別是complex128（64 位實數+64 位虛數）
	* 如果需要小一些的，也有complex64(32 位實數+32 位虛數)
	*
	* 複數的形式為RE + IMi，其中 RE 是實數部分，IM是虛數部分，而最後的 i 是虛數單位
	 */
	var complex1 complex64 = 5 + 5i //output: (5+5i)

	// String
	/*
	 * Go 中的字串都是採用UTF-8字符集編碼
	 * 字串是用一對雙引號（""）或反引號（`  ` ）括起來定義，它的型別是string
	 */
	var frenchHello string                 // 宣告變數為字串的一般方法
	var emptyString string = ""            // 宣告了一個字串變數，初始化為空字串
	no, yes, maybe := "no", "yes", "maybe" // 簡短宣告，同時宣告多個變數
	japaneseHello := "Konichiwa"           // 同上
	frenchHello = "Bonjour"                // 常規賦值
	// 在 Go 中字串是不可變的，例如下面的程式碼編譯時會報錯：cannot assign to s[0]
	// var s string = "hello"
	// s[0] = 'c'
	// 但如果真的想要修改怎麼辦呢？下面的程式碼可以實現：
	s := "hello"
	c := []byte(s) // 將字串 s 轉換為 []byte 型別
	c[0] = 'c'
	s2 := string(c) // 再轉換回 string 型別
	fmt.Printf("%s\n", s2)
	// Go 中可以使用+運算子來連線兩個字串：
	string1 := "hello,"
	m := " world"
	a := s + m
	fmt.Printf("%s\n", a)
	// 修改字串也可寫為：
	string2 := "hello"
	s = "c" + s[1:] // 字串雖不能更改，但可進行切片(slice)操作
	fmt.Printf("%s\n", s)
	//如果要宣告一個多行的字串怎麼辦？可以透過` 來宣告：
	str := `hello 
		world`
	/*
	* `  括起的字串為 Raw 字串，即字串在程式碼中的形式就是列印時的形式，它沒有字元轉義，換行也將原樣輸出。例如本例中會輸出：
	* hello
	* 	world
	 */

	// Error
	// * Go 內建有一個 error 型別，專門用來處理錯誤資訊，Go 的 package 裡面還專門有一個套件 errors 來處理錯誤：
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

	// --- Output ---
	fmt.Println("Boolean - isActive:", isActive, "enabled:", enabled, "disabled:", disabled, "valid:", valid)
	fmt.Println("Complex - complex1:", complex1)
	fmt.Println("String - frenchHello:", frenchHello, "emptyString:", emptyString, "no:", no, "yes:", yes, "maybe:", maybe, "japaneseHello:", japaneseHello, "string1:", string1, "string2:", string2, "str:", str)
	fmt.Println("Error - err:", err)

	fmt.Println("--- Types example finished ---")
}
