package gutil

func GetDefaultString(target, defaultStr string) string {
	if target == "" {
		return defaultStr
	}
	return target
}

func GetDefaultInt(target, defaultStr int) int {
	if target == 0 {
		return defaultStr
	}
	return target
}

func GetDefaultFloat32(target, defaultStr float32) float32 {
	if target == 0 {
		return defaultStr
	}
	return target
}

func String(v string) *string { return &v }
