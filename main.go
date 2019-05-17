package main

import (
	"fmt"
	"time"

	"github.com/boni/golang-mongo/config"
	"github.com/boni/golang-mongo/src/module/profile/model"
	"github.com/boni/golang-mongo/src/module/profile/repository"
)

func main() {
	fmt.Println("running")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewRepositoryMongo(db, "profile")
	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfile("U2", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U2"
	p.FirstName = "Tren"
	p.LastName = "Pra"
	p.Email = "bonnyiqbal12@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved..")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "dylan"
	p.LastName = "iqbal"
	p.Email = "bonnyiqbal12@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile update..")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile delete...")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profile {
		fmt.Println("---------------------")
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}
