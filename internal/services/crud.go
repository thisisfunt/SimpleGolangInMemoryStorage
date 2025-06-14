package services

var data = make(map[string]string)

func Set(key string, value string) {
	data[key] = value
}

func Get(key string) string {
	return data[key]
}

func Delete(key string) {
	delete(data, key)
}

func GetAll() map[string]string {
	return data
}
