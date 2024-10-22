package rest_err

import (
    "encoding/json"
    "errors"
    "net/http"
)

// Estrutura para o erro REST
type RestErr struct {
    Message string   `json:"message"`
    Err     string   `json:"error"`
    Code    int      `json:"code"`
    Causes  []Causes `json:"causes,omitempty"`
}

// Estrutura para as causas de erro
type Causes struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

// Função para criar um erro de Bad Request (400)
func NewBadRequestError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "bad_request",
        Code:    http.StatusBadRequest,
    }
}

// Função para criar um erro de Unauthorized (401)
func NewUnauthorizedError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "unauthorized",
        Code:    http.StatusUnauthorized,
    }
}

// Função para criar um erro de Forbidden (403)
func NewForbiddenError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "forbidden",
        Code:    http.StatusForbidden,
    }
}

// Função para criar um erro de Not Found (404)
func NewNotFoundError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "not_found",
        Code:    http.StatusNotFound,
    }
}

// Função para criar um erro de Conflict (409)
func NewConflictError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "conflict",
        Code:    http.StatusConflict,
    }
}

// Função para criar um erro de Unprocessable Entity (422)
func NewUnprocessableEntityError(message string) *RestErr {
    return &RestErr{
        Message: message,
        Err:     "unprocessable_entity",
        Code:    http.StatusUnprocessableEntity,
    }
}

// Função para criar um erro de Internal Server Error (500)
func NewInternalServerError(message string, err error) *RestErr {
    restErr := &RestErr{
        Message: message,
        Err:     "internal_server_error",
        Code:    http.StatusInternalServerError,
    }
    if err != nil {
        restErr.Causes = []Causes{
            {
                Field:   "system",
                Message: err.Error(),
            },
        }
    }
    return restErr
}

// Método para transformar o erro em JSON
func (e *RestErr) ToJSON() string {
    jsonData, _ := json.Marshal(e)
    return string(jsonData)
}

// Método para retornar um erro como um objeto Go padrão
func (e *RestErr) Error() string {
    return e.Message
}

// Função para converter um erro Go em um RestErr
func NewRestErrorFromBytes(bytes []byte) (*RestErr, error) {
    var apiErr RestErr
    if err := json.Unmarshal(bytes, &apiErr); err != nil {
        return nil, errors.New("invalid json")
    }
    return &apiErr, nil
}
