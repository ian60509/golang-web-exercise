package repository

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "user_server/model"
)

type UserRepositoryTestSuite struct { // 將測試用到的包成一包 suite
    suite.Suite
    db       *gorm.DB
    repo     UserRepository 
}

// 跑每一個 test function 都會設定
func (suite *UserRepositoryTestSuite) SetupTest() {
    // 使用 SQLite 的內存資料庫進行測試
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        suite.T().Fatal(err)
    } else {
		suite.T().Log("Connect to sqlite success")
	}

    db.AutoMigrate(&model.User{})
    suite.db = db
    suite.repo = NewUserRepository(db) // 測試的目標
}

// -------------------------- Test Function -------------------------------

func (suite *UserRepositoryTestSuite) TestCreate() {
    user := &model.User{Name: "John Doe", Email: "john@example.com"}
    err := suite.repo.Create(user)
    assert.NoError(suite.T(), err)
    assert.NotZero(suite.T(), user.ID) //ID 由資料庫自動生成後，會被賦值給 user.ID
	suite.T().Log("Create user success， ID: ", user.ID)
}

func (suite *UserRepositoryTestSuite) TestGetAll() {
    users := []model.User{
        {Name: "John Doe", Email: "john@example.com"},
        {Name: "Jane Doe", Email: "jane@example.com"},
    }
	//創建兩個 user
    for _, user := range users {
        err := suite.repo.Create(&user)
        assert.NoError(suite.T(), err)
    }

    retrievedUsers, err := suite.repo.GetAll()
    assert.NoError(suite.T(), err)
    assert.Len(suite.T(), retrievedUsers, 2)

	// 印出 GetAll 取得的 users
	for i, user := range retrievedUsers {
		assert.Equal(suite.T(), users[i].Name, user.Name)
		assert.Equal(suite.T(), users[i].Email, user.Email)
	}
}

func (suite *UserRepositoryTestSuite) TestGetByID() {
    user := &model.User{Name: "John Doe", Email: "john@example.com"}
    suite.repo.Create(user)

    retrievedUser, err := suite.repo.GetById(int(user.ID))
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), user.Name, retrievedUser.Name)
    assert.Equal(suite.T(), user.Email, retrievedUser.Email)
}

func (suite *UserRepositoryTestSuite) TestUpdate() {
    user := &model.User{Name: "John Doe", Email: "john@example.com"}
    suite.repo.Create(user)

	newName := "John Smith"
    user.Name = newName 
    err := suite.repo.Update(user) //將資料庫中的 user name 更新為 newName
    assert.NoError(suite.T(), err)

    updatedUser, err := suite.repo.GetById(int(user.ID))
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), newName, updatedUser.Name)
}

func (suite *UserRepositoryTestSuite) TestDelete() {
    user := &model.User{Name: "John Doe", Email: "john@example.com"}
    suite.repo.Create(user)

    err := suite.repo.Delete(user)
    assert.NoError(suite.T(), err)

    _, err = suite.repo.GetById(int(user.ID))
    assert.Error(suite.T(), err) //已刪除，應該要是Error
}

// -------------------------- 開始跑 Test Suit-------------------------------
func TestUserRepositoryTestSuite(t *testing.T) {
    suite.Run(t, new(UserRepositoryTestSuite)) //UserRepositoryTestSuite 這個自訂義的 struct 可以被傳入給所有的 Test Function
}
