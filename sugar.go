package go_sugar

// IfExpr 类三元表达式
func IfExpr[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}
