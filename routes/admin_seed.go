package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/smallbatch-apps/trivialize-go/database"
	"github.com/smallbatch-apps/trivialize-go/models"
)

var PredefinedTags = map[string]models.Tag{
	"Literature": {ID: uuid.MustParse("0cc875a9-7338-495a-94b1-5aa28bf2bbd9"), Text: "Literature", Colour: "#92400E", IconColour: "#FBBF24", IconSecondaryColour: "#B45309", Icon: "books"},
	"Film":       {ID: uuid.MustParse("eae15870-b4b2-440f-b36e-78ef517e8bbe"), Text: "Film", Colour: "", IconColour: "#1F2937", IconSecondaryColour: "#7DD3FC", Icon: "film"},
	"Sports":     {ID: uuid.MustParse("604abd8b-11e0-419a-ab91-ab8abf430244"), Text: "Sports", Colour: "", IconColour: "#FAFAF9", IconSecondaryColour: "#44403C", Icon: "futbol"},
	"Music":      {ID: uuid.MustParse("607b3a19-04e8-4e35-b2c5-07cedd919695"), Text: "Music", Colour: "", IconColour: "#7E22CE", IconSecondaryColour: "#6B21A8", Icon: "music"},
	"Science":    {ID: uuid.MustParse("c948e09c-2606-4eb3-9461-8334b3f71135"), Text: "Science", Colour: "", IconColour: "", IconSecondaryColour: "", Icon: "atom-alt"},
	"Geography":  {ID: uuid.MustParse("213467cd-10a0-4d9b-9f50-26374dd62963"), Text: "Geography", Colour: "", IconColour: "#22C55E", IconSecondaryColour: "#38BDF8", Icon: "globe-africa"},
	"Nature":     {ID: uuid.MustParse("bb1d66ce-c7cf-43d4-a672-0d1006264ba0"), Text: "Nature", Colour: "", IconColour: "#854D0E", IconSecondaryColour: "#FDE047", Icon: "kiwi-bird"},
	"Gaming":     {ID: uuid.MustParse("f14334ea-7e8f-4dd6-b3ba-e4970f4c16bc"), Text: "Gaming", Colour: "", IconColour: "#D6D3D1", IconSecondaryColour: "#EF4444", Icon: "gamepad-alt"},
	"Easy":       {ID: uuid.MustParse("454df36f-34a0-4e75-ad28-e614873fae93"), Text: "Easy", Colour: "", IconColour: "#BEF264", IconSecondaryColour: "#EF4444", Icon: "shapes"},
	"Difficult":  {ID: uuid.MustParse("ce3bab2f-fec9-4b43-9fc1-53fcd0041f21"), Text: "Difficult", Colour: "", IconColour: "#F9A8D4", IconSecondaryColour: "#EC4899", Icon: "brain"},
	"Television": {ID: uuid.MustParse("9a46261d-db91-4c2e-8f65-3aefff8aeef1"), Text: "Television", Colour: "", IconColour: "#713F12", IconSecondaryColour: "#EC4899", Icon: "tv-retro"},
	"History":    {ID: uuid.MustParse("786f4936-8378-4bcd-a18f-8a78b9637dfb"), Text: "History", Colour: "", IconColour: "#A8A29E", IconSecondaryColour: "#D6D3D1", Icon: "landmark"},
	"Local":      {ID: uuid.MustParse("ccb5a554-8f9d-4aa8-ad11-8bc69f19ce5f"), Text: "Local", Colour: "", IconColour: "#A8A29E", IconSecondaryColour: "#D6D3D1", Icon: "location-dot"},
}

type SeedQuestion struct {
	Text      string
	Answer    []models.Answer
	TagKeys   []string
	Documents []models.Document
}

var questionsToSeed = []SeedQuestion{
	{
		Text:    "In Doctor Who what does T.A.R.D.I.S. stand for?",
		Answer:  "Time And Relative Dimensions In Space",
		TagKeys: []string{"Television"},
		Documents: []models.Document{
			Title:    "The TARDIS - Doctor Who",
			Location: "tardis.jpeg",
		},
	},
	// ... add more questions ...
}

func InsertSeedData(c *fiber.Ctx) error {
	database.Database.Db.AutoMigrate(&models.Answer{}, &models.Company{}, &models.Document{}, &models.Event{}, &models.Question{}, &models.Series{}, &models.Tag{}, &models.User{})

	companyIdString := "c7c3c627-322f-41fc-a483-957428f79f5b"

	companyId, err := uuid.Parse(companyIdString)

	if err != nil {
		return c.SendStatus(500)
	}

	// for _, sq := range questionsToSeed {
	// 	question := models.Question{
	// 		Text:      sq.Text,
	// 		CompanyID: companyId,
	// 		Documents: []models.Document{sq.Document},
	// 		Tags:      []models.Tag{predefinedTags[sq.TagText]},
	// 	}

	// 	// Create the question and its associated document and tag
	// 	err := database.Database.Db.Create(&question).Error
	// 	if err != nil {
	// 		// handle error
	// 		continue
	// 	}

	// 	// Create the answer associated with the question
	// 	answer := models.Answer{
	// 		Text:       sq.Answer,
	// 		QuestionID: question.ID,
	// 		// Set other fields like Points, Sort, etc.
	// 	}

	// 	err = database.Database.Db.Create(&answer).Error
	// 	if err != nil {
	// 		// handle error
	// 	}
	// }
	return c.SendString("Insert Seed Data")
}
