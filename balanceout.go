package sprint


func BalanceOut(arr []bool) []bool {
	TrueValue := 0
	FalseValue := 0

	for _, v := range arr {
		if v == true {
			TrueValue++
		} else {
			FalseValue++
		}
	}

	if TrueValue > FalseValue {
		for i := 0; i < (TrueValue - FalseValue); i++ {
			arr = append(arr, false)
		}
	} else if FalseValue > TrueValue {
		for i := 0; i < (FalseValue - TrueValue); i++ {
			arr = append(arr, true)
		}

	}
	return arr

}



