package main

import (
	"errors"
	"fmt"
	"math"
)

func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	m := fmt.Sprintf("int8  = %v > in64  = %v\n", i8, int64(i8))
	m += fmt.Sprintf("int   = %v > in8   = %v\n", i, int8(i))
	m += fmt.Sprintf("int8  = %v > float32 = %v\n", i8, float64(i8))
	m += fmt.Sprintf("float64 = %v > int   = %v\n", f64, int(f64))
	return m
}

// dubler means that return the value * 2
func doubler(v interface{}) (string, error) {
	//out1, out2 := v.(int) // out1 es el valor que se psa y out2 es el Assert
	// Por lo cual si se cumple out2 se pueden hacer cosas como lo siguiente
	/// el problema es que si no se hace bien se obtendra un runTimeExceptio , riesgo
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}
	return "", errors.New("unsupported type passed")
}
func doublerSwitch(v interface{}) (string, error) {

	switch t := v.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}

		return fmt.Sprint(t.(float32) * 2), nil
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:
		return fmt.Sprint(t * 2), nil
	case int32:
		return fmt.Sprint(t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint8:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	fmt.Println("::::::::Numeric Type Conversion:::::::")
	/// si hace overflow el resultado será números negativos
	/// note el int   = 128 > in8   = -128
	fmt.Println(convert())

	fmt.Println("::::::::Type Assertions and interface{}:::::::")
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("yum :", res)
	_, err := doubler(true)
	fmt.Println("true:", err)

	fmt.Println("::::::::Switch --->Type Assertions and interface{}:::::::")
	res, _ = doublerSwitch(-5)
	fmt.Println("-5 :", res)
	res, _ = doublerSwitch(-5)
	fmt.Println("-5 :", res)
	res, _ = doublerSwitch("yum")
	fmt.Println("yum :", res)
	res, _ = doublerSwitch("yum")
	fmt.Println("yum :", res)
	res, _ = doublerSwitch(true)
	fmt.Println("true:", res)
	res, _ = doublerSwitch(float32(3.14))
	fmt.Println("3.14:", res)

}
