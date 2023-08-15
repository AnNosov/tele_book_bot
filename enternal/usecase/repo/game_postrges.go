package repo

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/AnNosov/tele_bot/enternal/entity"
	"github.com/AnNosov/tele_bot/pkg/postgres"
)

type PgRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *PgRepo {
	return &PgRepo{pg}
}

func (r *PgRepo) GetBooksList() ([]entity.Book, error) { // solution for test, need 2 methods for names and tables
	rows, err := r.Postgres.DB.Query("select id, name from test.books")
	if err != nil {
		return nil, fmt.Errorf("GetBooksList: %w", err)
	}
	defer rows.Close()

	books := make([]entity.Book, 0)

	for rows.Next() {
		b := entity.Book{}

		err := rows.Scan(&b.Id, &b.Name)
		if err != nil {
			log.Println("GetBooksList: ", err)
			continue
		}
		books = append(books, b)
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("books from postgres is empty")
	}

	return books, nil
}

func (r *PgRepo) GetStepInfo(element *entity.Element) error {

	err := r.Postgres.DB.QueryRow("SELECT full_desc, type from test.element_info where id = $1", element.Id).Scan(&element.Desc, &element.Type)

	element.Desc = strings.ReplaceAll(element.Desc, `\n`, "\n") // замена на переход строки

	if err != nil {
		return fmt.Errorf("GetStepInfo: %w", err)
	}

	return nil
}

func (r *PgRepo) GetFirstStep(id int) (int, error) {
	var idFirstElement sql.NullInt64
	err := r.Postgres.DB.QueryRow("SELECT distinct first_element from test.books where id = $1", id).Scan(&idFirstElement)

	if err != nil {
		return 0, fmt.Errorf("GetFirstStep: %w", err)
	}

	if !idFirstElement.Valid {
		return 0, fmt.Errorf("GetFirstStep: first element is NULL")
	}

	return int(idFirstElement.Int64), nil
}

func (r *PgRepo) GetNextSteps(element *entity.Element) error {

	rows, err := r.Postgres.DB.Query("SELECT es.next_id, ei.short_desc from test.element_schema es, test.element_info ei where es.next_id = ei.id AND es.id = $1", element.Id)
	if err != nil { // сделать отдельную проверку что последующие элементы отсутствуют
		return fmt.Errorf("GetNextSteps: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var nElement sql.NullInt64
		var info sql.NullString

		err := rows.Scan(&nElement, &info)

		if err != nil {
			log.Println("GetNextSteps: ", err)
			continue
		}
		element.Next[int(nElement.Int64)] = info.String
	}

	return nil
}
