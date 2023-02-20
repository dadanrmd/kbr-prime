package requestValidations

func IsNIKValid(nik string) bool {
	if len(nik) != 16 {
		return false
	}

	return true
}
