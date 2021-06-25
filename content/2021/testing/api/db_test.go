package api_test

//STARTDB,OMIT
import (
	"database/sql" //OMIT
	"testing"      //OMIT

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/imrenagi/gotalks/content/2021/testing/api"   //OMIT
	. "github.com/imrenagi/gotalks/content/2021/testing/api" //OMIT
	"github.com/stretchr/testify/assert"                     //OMIT
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New() // mocking db connection // HL
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{}) // HL
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return sqlDB, gormDB, mock // HL
}

//STOPDB,OMIT

//STARTTEST,OMIT
func TestUserFetcher(t *testing.T) {
	sqlmock, db, mock := DB(t) // HL

	f := api.UserFetcher{
		DB: db, //mock // HL
	}

	user := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("1", "imre")

	mock.ExpectQuery("SELECT (.+) FROM \"users\" Where id = ? ").
		WithArguments("1").
		WillReturnRows(user)

	u, err := f.FindByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "imre", u.Name)
	assert.Nil(t, mock.ExpectationsWereMet()) // verification // HL
}

//STOPTEST,OMIT
