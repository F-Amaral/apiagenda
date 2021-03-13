package agendarepositories

import (
	"context"
	"github.com/F-Amaral/apiagenda/internal/api/apierror"
	"github.com/F-Amaral/apiagenda/pkg/contracts"
	"github.com/F-Amaral/apiagenda/pkg/domain/entities"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{
		db: db,
	}
}

func (self contactRepository) Add(ctx context.Context, contact *entities.Contact) apierror.ApiError {
	contact.Deleted = false
	err := self.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(contact).Error
		if err != nil {
			return apierror.New(http.StatusInternalServerError, err.Error())
		}
		return nil
	})
	if err != nil {
		apiErr, ok := err.(apierror.ApiError)
		if ok {
			return apiErr
		}
		return apierror.New(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (self contactRepository) Remove(ctx context.Context, id string) apierror.ApiError {
	tx := self.db.Model(&entities.Contact{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted": true,
	})
	if tx.Error != nil {
		return apierror.New(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}

func (self contactRepository) GetById(ctx context.Context, id string) (*entities.Contact, apierror.ApiError) {
	contact := &entities.Contact{}
	tx := self.db.Where("id = ?", id).First(contact)
	if tx.Error != nil {
		return nil, apierror.New(http.StatusInternalServerError, tx.Error.Error())
	}
	return contact, nil
}

func (self contactRepository) Search(ctx context.Context, searchRequest *contracts.SearchRequest) ([]*entities.Contact, apierror.ApiError) {
	query := []string{}
	args := []interface{}{}
	response := []*entities.Contact{}

	if searchRequest.Deleted != nil {
		query = append(query, "contacts.deleted = ?")
		args = append(args, *searchRequest.Deleted)
	}

	if searchRequest.Id != nil {
		query = append(query, "contacts.id = ?")
		args = append(args, *searchRequest.Id)
	}

	if searchRequest.Name != nil {
		query = append(query, "contacts.name = ?")
		args = append(args, *searchRequest.Name)
	}

	if searchRequest.Email != nil {
		query = append(query, "contacts.email = ?")
		args = append(args, *searchRequest.Email)
	}

	tx := self.db.Table("contacts").Select("contacts.id, contacts.name, contacts.email, contacts.deleted")
	if len(query) > 0 {
		tx = tx.Where(strings.Join(query, " AND "), args...)
	}
	tx = tx.Limit(300)

	tx = tx.Scan(&response)
	if tx.Error != nil {
		return nil, apierror.New(http.StatusInternalServerError, tx.Error.Error())
	}
	return response, nil
}

func (self contactRepository) Update(ctx context.Context, contact *entities.Contact) apierror.ApiError {
	tx := self.db.Model(&entities.Contact{}).Where("id = ?", contact.Id).Updates(map[string]interface{}{
		"name":  contact.Name,
		"email": contact.Email,
	})
	if tx.Error != nil {
		return apierror.New(http.StatusInternalServerError, tx.Error.Error())
	}
	return nil
}
