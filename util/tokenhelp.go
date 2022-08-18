package util

// 载荷 定义自己的东西
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 江定义的载荷加密生成token
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()                    //当前时间，获得token时的时间
	expireTime := nowTime.Add(3 * time.Hour) //有效时间，获得token后token的有效时间

	claims := Claims{ // 定义token里的信息
		username,
		password,
		jwt.StandardClaims{ // 加密的声明
			ExpiresAt: expireTime.Unix(),
			Issuer:    "its me",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 新建声明
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析token，拿到结构体，就可以验证结构体里面的用户及密码是否正确
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
