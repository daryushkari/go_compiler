package strUtl


func CheckHasString(s string, ar []string) bool{
	for _, x := range ar{
		if x == s{
			return true
		}
	}
	return false
}