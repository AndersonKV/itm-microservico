package auth

import (
	"errors"
	"testing"

	"github.com/AndersonKV/auth_service/internal/models"
)

// ==========================
// MockRepository
// ==========================

type MockRepository struct {
    users map[string]models.User
}


 
func NewMockRepository() *MockRepository {
    return &MockRepository{
        users: make(map[string]models.User),
    }
}

 
// CreateUser simula criação de usuário
func (m *MockRepository) CreateUser(user models.User) error {
    if _, exists := m.users[user.Email]; exists {
        return errors.New("usuario já existe")
    }
    m.users[user.Email] = user
    return nil
}

// GetUserByEmail simula busca de usuário por email
func (m *MockRepository) GetUserByEmail(email string) (models.User, error) {
    user, exists := m.users[email]
    if !exists {
        return models.User{}, errors.New("usuario nao encontrado")
    }
    return user, nil
}

// ==========================
// Testes
// ==========================

func TestRegisterUser_Success(t *testing.T) {
repo := NewMockRepository()
service := NewAuthService(repo)  

    err := service.Register("user1", "user1@email.com", "senha123", "foto.png" )
    if err != nil {
        t.Errorf("Esperava nil, recebeu: %v", err)
    }
}

func TestRegisterUser_DuplicateEmail(t *testing.T) {
    repo := NewMockRepository()
    service := NewAuthService(repo)

    // Primeiro registro
    _ = service.Register("user1", "user1@email.com", "senha123", "foto.png" )
    
    // Segundo registro com mesmo email
    err := service.Register("user2", "user1@email.com", "senha456", "foto2.png" )
    if err == nil {
        t.Errorf("Esperava erro de usuário duplicado, recebeu nil")
    }
}

func TestLoginUser_Success(t *testing.T) {
    repo := NewMockRepository()
    service := NewAuthService(repo)

    // Registrar usuário
    _ = service.Register("user1", "user1@email.com", "senha123", "foto.png")

    token, err := service.Login("user1@email.com", "senha123")
    if err != nil {
        t.Errorf("Esperava nil, recebeu: %v", err)
    }
    if token == "" {
        t.Errorf("Esperava token válido, recebeu vazio")
    }
}

func TestLoginUser_WrongPassword(t *testing.T) {
    repo := NewMockRepository()
    service := NewAuthService(repo)

    _ = service.Register("user1", "user1@email.com", "senha123", "foto.png" )

    _, err := service.Login("user1@email.com", "senhaErrada")
    if err == nil {
        t.Errorf("Esperava erro de senha incorreta, recebeu nil")
    }
}

func TestLoginUser_NotFound(t *testing.T) {
    repo := NewMockRepository()
    service := NewAuthService(repo)

    _, err := service.Login("naoexiste@email.com", "senha123")
    if err == nil {
        t.Errorf("Esperava erro de usuário não encontrado, recebeu nil")
    }
}
