package utils

import "math"

// type Regular[T any] struct {
// 	Data T
// }

// 定义一个泛型函数，用于过滤数组中的元素。
func Filter[T any](arr []T, fn func(s T) bool) []T {
	var result []T
	for _, item := range arr {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func GroupBy[T any](arr []T, fn func(s T) string) map[string][]T {
	result := make(map[string][]T)
	for _, item := range arr {
		key := fn(item)
		result[key] = append(result[key], item)
	}
	return result
}

func Sum[T any](arr []T, fn func(s T) float64, digit float64) float64 {
	var result float64
	for _, item := range arr {
		result += fn(item)
	}
	if digit >= 0 {
		return math.Round(result*math.Pow(10, digit)) / math.Pow(10, digit)
	} else {
		return result
	}
}

// IsExistArray 判断字符串是否存在于数组，字符串转换为 map， 时间复杂度O（1）
func IsExistArray(str string, arr []string) bool {
	arrMap := make(map[string]bool)
	for _, item := range arr {
		arrMap[item] = true
	}
	if _, found := arrMap[str]; found {
		return true
	} else {
		return false
	}
}
