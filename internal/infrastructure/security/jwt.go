package security

import (
    "github.com/golang-jwt/jwt"
    "time"
)

var secretKey = []byte("tu_clave_secreta")

func GenerateToken(code string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "code": code,
        "exp":  time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["code"].(string), nil
    }

    return "", err
}