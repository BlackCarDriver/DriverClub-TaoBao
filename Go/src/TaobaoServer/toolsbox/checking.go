package toolsbox

import (
	"regexp"
)

//=========== request data checking ===================== 🍚
//goods name format
func CheckGoodsName(name string) bool {
	reg := regexp.MustCompile(`^[\p{Han}_a-zA-Z0-9]{2,15}$`)
	if !reg.MatchString(name) {
		return false
	}
	return true
}

//email format
func CheckEmail(email string) bool {
	reg := regexp.MustCompile(`^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`)
	if !reg.MatchString(email) {
		return false
	}
	return true
}

//user name format
func CheckUserName(name string) bool {
	reg := regexp.MustCompile(`^[\p{Han}_a-zA-Z0-9]{2,15}$`)
	if !reg.MatchString(name) {
		return false
	}
	return true
}

//user ID format
func CheckUserID(id string) bool {
	reg := regexp.MustCompile(`^[1-2][0-9][01][0-9][0-3][0-9][0-9]{2,3}$`)
	if !reg.MatchString(id) {
		return false
	}
	return true
}

//user password format
func CheckPassword(pw string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z._0-9]{6,20}$`)
	if !reg.MatchString(pw) {
		return false
	}
	return true
}

//sign up comfirm code format
func CheckComfirmCode(code string) bool {
	reg := regexp.MustCompile(`^[0-9]{6}$`)
	if !reg.MatchString(code) {
		return false
	}
	return true
}

//goods title
func CheckGoodsTitle(title string) bool {
	reg := regexp.MustCompile(`^.{5,45}$`)
	if !reg.MatchString(title) {
		return false
	}
	return true
}

//goods comment and user comment
func CheckComment(comment string) bool {
	reg := regexp.MustCompile(`^[\w\W]{2,200}$`)
	if !reg.MatchString(comment) {
		return false
	}
	return true
}

//private message checking
func CheckMessage(msg string) bool {
	reg := regexp.MustCompile(`^[\w\W]{2,150}$`)
	if !reg.MatchString(msg) {
		return false
	}
	return true
}

//feedback message
func CheckFeedbackDetail(describe string) bool {
	reg := regexp.MustCompile(`^[\w\W]{2,450}$`)
	if !reg.MatchString(describe) {
		return false
	}
	return true
}

//check grade
func CheckGrade(grade string) bool {
	reg := regexp.MustCompile(`^20[1-3][0-9]$`)
	if !reg.MatchString(grade) {
		return false
	}
	return true
}
