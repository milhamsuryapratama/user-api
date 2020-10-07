package helper

// ImageValidation ...
func ImageValidation(mime string) bool {
	if mime == "images/jpg" || mime == "image/jpeg" || mime == "image/png" {
		return true
	}

	return false
}
