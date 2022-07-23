package internal

import (
	"log"
	"strconv"
)

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	membershipBuilder := NewMembershipGenerator()
	memberRepository := app.repository
	memberCount := len(memberRepository.data)

	id := strconv.Itoa(memberCount + 1)
	log.Println(id)

	membershipBuilder.
		SetID(id).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType)

	membership, err := membershipBuilder.GetMembership()
	if err != nil {
		return CreateResponse{}, err
	}

	return CreateResponse{membership.ID, membership.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	newMembership, err := NewMembershipGenerator().
		SetID(request.ID).
		SetUserName(request.UserName).
		SetMembershipType(request.MembershipType).
		GetMembership()
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{
		ID:             newMembership.ID,
		UserName:       newMembership.UserName,
		MembershipType: newMembership.MembershipType,
	}, nil
}

func (app *Application) Delete(id string) error {
	m := app.repository.data[id]
	mem, err := NewMembershipGenerator().
		SetID(m.ID).
		SetUserName(m.UserName).
		SetMembershipType(m.MembershipType).
		GetMembership()
	if err != nil {
		return err
	}
	err = app.repository.DeleteRepositoryData(*mem)
	if err != nil {
		return err
	}
	return nil
}
