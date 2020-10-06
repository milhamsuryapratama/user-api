package helper

// ImageValidation ...
func ImageValidation(mime string) bool {
	if mime == "images/jpg" {
		return true
	} else if mime == "image/jpeg" {
		return true
	} else if mime == "image/png" {
		return true
	} else {
		return false
	}
}
