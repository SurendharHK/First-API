package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)


type ProfileService struct{
	ProfileCollection *mongo.Collection
	ctx context.Context
}

func (p *ProfileService ) CreateProfile(profile *models.Profile){}
func (p *ProfileService ) EditProfile(profile *models.Profile){}
func (p *ProfileService ) DeleteProfile(Profileid int){}
func (p *ProfileService ) GetProfileById(profile int){}
func (p *ProfileService ) GetProfileBySearch(){}