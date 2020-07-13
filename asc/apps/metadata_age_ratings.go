package apps

import (
	"fmt"

	"github.com/aaronsky/asc-go/v1/asc/common"
)

// AppStoreAgeRating defines model for AppStoreAgeRating.
type AppStoreAgeRating string

// List of AppStoreAgeRating
const (
	FourPlus      AppStoreAgeRating = "FOUR_PLUS"
	NinePlus      AppStoreAgeRating = "NINE_PLUS"
	SeventeenPlus AppStoreAgeRating = "SEVENTEEN_PLUS"
	TwelvePlus    AppStoreAgeRating = "TWELVE_PLUS"
)

// BrazilAgeRating defines model for BrazilAgeRating.
type BrazilAgeRating string

// List of BrazilAgeRating
const (
	Eighteen BrazilAgeRating = "EIGHTEEN"
	Fourteen BrazilAgeRating = "FOURTEEN"
	L        BrazilAgeRating = "L"
	Sixteen  BrazilAgeRating = "SIXTEEN"
	Ten      BrazilAgeRating = "TEN"
	Twelve   BrazilAgeRating = "TWELVE"
)

// KidsAgeBand defines model for KidsAgeBand.
type KidsAgeBand string

// List of KidsAgeBand
const (
	FiveAndUnder KidsAgeBand = "FIVE_AND_UNDER"
	NineToEleven KidsAgeBand = "NINE_TO_ELEVEN"
	SixToEight   KidsAgeBand = "SIX_TO_EIGHT"
)

// AgeRatingDeclarationUpdateRequest defines model for AgeRatingDeclarationUpdateRequest.
type AgeRatingDeclarationUpdateRequest struct {
	Data struct {
		Attributes *struct {
			AlcoholTobaccoOrDrugUseOrReferences         *string      `json:"alcoholTobaccoOrDrugUseOrReferences,omitempty"`
			GamblingAndContests                         *bool        `json:"gamblingAndContests,omitempty"`
			GamblingSimulated                           *string      `json:"gamblingSimulated,omitempty"`
			HorrorOrFearThemes                          *string      `json:"horrorOrFearThemes,omitempty"`
			KidsAgeBand                                 *KidsAgeBand `json:"kidsAgeBand,omitempty"`
			MatureOrSuggestiveThemes                    *string      `json:"matureOrSuggestiveThemes,omitempty"`
			MedicalOrTreatmentInformation               *string      `json:"medicalOrTreatmentInformation,omitempty"`
			ProfanityOrCrudeHumor                       *string      `json:"profanityOrCrudeHumor,omitempty"`
			SexualContentGraphicAndNudity               *string      `json:"sexualContentGraphicAndNudity,omitempty"`
			SexualContentOrNudity                       *string      `json:"sexualContentOrNudity,omitempty"`
			UnrestrictedWebAccess                       *bool        `json:"unrestrictedWebAccess,omitempty"`
			ViolenceCartoonOrFantasy                    *string      `json:"violenceCartoonOrFantasy,omitempty"`
			ViolenceRealistic                           *string      `json:"violenceRealistic,omitempty"`
			ViolenceRealisticProlongedGraphicOrSadistic *string      `json:"violenceRealisticProlongedGraphicOrSadistic,omitempty"`
		} `json:"attributes,omitempty"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"data"`
}

// AgeRatingDeclarationResponse defines model for AgeRatingDeclarationResponse.
type AgeRatingDeclarationResponse struct {
	Data  AgeRatingDeclaration `json:"data"`
	Links common.DocumentLinks `json:"links"`
}

// UpdateAgeRatingDeclaration provides age-related information so the App Store can determine the age rating for your app.
func (s *Service) UpdateAgeRatingDeclaration(id string, body *AgeRatingDeclarationUpdateRequest) (*AgeRatingDeclarationResponse, *common.Response, error) {
	url := fmt.Sprintf("ageRatingDeclarations/%s", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.Patch(url, body, res)
	return res, resp, err
}
