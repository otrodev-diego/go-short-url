package security

import (
    "testing"
)

func TestJWT(t *testing.T) {
    t.Run("Generar y validar token", func(t *testing.T) {
        testCode := "abc123"

        token, err := GenerateToken(testCode)
        if err != nil {
            t.Fatalf("Error generando token: %v", err)
        }

        code, err := ValidateToken(token)
        if err != nil {
            t.Fatalf("Error validando token: %v", err)
        }

        if code != testCode {
            t.Errorf("Código esperado %s, obtenido %s", testCode, code)
        }
    })

    t.Run("Token inválido", func(t *testing.T) {
        _, err := ValidateToken("token.invalido")
        if err == nil {
            t.Error("Se esperaba error con token inválido")
        }
    })
}