package interfaces

import 

type IProfile interface{
	CreateProfile(profile *models.Profile)
	EditProfile(profile *models.Profile)
	DeleteProfile(profileId int)
	GetProfileById(profile int)
	GetProfileBySearch()
}