package utils

// 判断sqlite是否是唯一约束错误
func IsUniqueConstraintError(err error) string {
	if err != nil {
		// 判断是否是sqlite的唯一约束错误
		if err.Error() == "UNIQUE constraint failed: admins.phone" {
			return "手机号已存在"
		}
		if err.Error() == "UNIQUE constraint failed: admins.email" {
			return "邮箱已存在"
		}
	}
	return ""
}
