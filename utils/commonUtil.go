package utils

func GetDBType(dbType int) string {
	switch dbType {
		case 1:
			return "MySQL"
		case 2:
			return "Oracle"
		case 3:
			return "SQLServer"
		default:
			return ""
	}
}
