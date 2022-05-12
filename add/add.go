package add

// Addをaddとしてしまうと外部（ここではmain.go）から参照できない
// 外部から参照する変数や関数は大文字から始める
func Add(a, b int) int {
    return a + b
}
