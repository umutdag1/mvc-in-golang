package databaser

var (
	inMemDB map[string]interface{}
)

func InitInMemDB() {
	inMemDB = make(map[string]interface{})
}

func GetInMemDB() map[string]interface{} {
	return inMemDB
}

func SetValToKey(key string, val interface{}) {
	inMemDB[key] = val
}

func GetValWithKeyByKey(key string) map[string]interface{} {
	res := make(map[string]interface{})
	if SearchValByKey(key) {
		res[key] = inMemDB[key]
	}
	return res
}

func SearchValByKey(key string) bool {
	_, isExist := inMemDB[key]
	return isExist
}
