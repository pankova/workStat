package config

// Config - структура, настроек доступа к апи, от пользователя требуется ввести токен доступа, который можно получить так:
//
// 1. Вставить в строку браузера 
//    https://oauth.vk.com/authorize?client_id=5812131&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=friends&response_type=token&v=5.62
// 2. Откроется окно 
//    "Пожалуйста, не копируйте данные из адресной строки для сторонних сайтов. Таким образом Вы можете потерять доступ к Вашему аккаунту."
// 3. Из адресной строки вида 
//    https://oauth.vk.com/blank.html#access_token=38f3d5532ceabcec3fc0r60d46f0d9dd784b5c6f2b9e8eb2c6daa46b270r36f96b10au2bm33c478849180&expires_in=86400&user_id=16935367
//    скопировать часть 38f3d5532ceabcec3fc0r60d46f0d9dd784b5c6f2b9e8eb2c6daa46b270r36f96b10au2bm33c478849180
// 4. Вставить ее ниже после "&access_token="

type Config struct {
	Tocken string
	Common string
	Version string
}

func NewConfig() Config {
	config := Config{
		Tocken: "&access_token=38f3d5532ceabcec3fc0140d46f0d9dd784b5c6f2b9e8eb2c6daa46b270236f96b10a22b433c478849180",
		Common: "https://api.vk.com/method/",
		Version: "v=5.62"}
	return config
}