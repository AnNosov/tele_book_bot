package usecase

import (
	"fmt"

	"github.com/AnNosov/tele_bot/enternal/entity"
	"github.com/AnNosov/tele_bot/enternal/usecase/repo"
	"github.com/AnNosov/tele_bot/pkg/tlgrm"
)

type GameAction struct {
	repo  *repo.PgRepo
	Tlgrm *tlgrm.TlgBot
}

func New(pgRepo *repo.PgRepo, tBot *tlgrm.TlgBot) *GameAction {
	return &GameAction{repo: pgRepo, Tlgrm: tBot}
}

func (gAction *GameAction) BookList() ([]entity.Book, error) {
	books, err := gAction.repo.GetBooksList()
	if err != nil {
		return nil, fmt.Errorf("BookList: %w", err)
	}
	return books, nil
}

func (gAction *GameAction) FirstStepElement(idSchema int) (entity.Element, error) {
	id, err := gAction.repo.GetFirstStep(idSchema)
	if err != nil {
		return entity.Element{}, fmt.Errorf("StepInfo: %w", err)
	}

	var step entity.Element
	step.Id = id
	step.Next = make(map[int]string)

	if err := gAction.repo.GetStepInfo(&step); err != nil {
		return entity.Element{}, fmt.Errorf("StepInfo: %w", err)
	}

	if err := gAction.repo.GetNextSteps(&step); err != nil {
		return entity.Element{}, fmt.Errorf("StepInfo: %w", err)
	}

	return step, nil
}

func (gAction *GameAction) StepInfo(id int) (entity.Element, error) {
	var step entity.Element
	step.Id = id
	step.Next = make(map[int]string)

	if err := gAction.repo.GetStepInfo(&step); err != nil {
		return entity.Element{}, fmt.Errorf("StepInfo: %w", err)
	}

	if err := gAction.repo.GetNextSteps(&step); err != nil {
		return entity.Element{}, fmt.Errorf("StepInfo: %w", err)
	}

	return step, nil
}
