package main

import (
	"github.com/dhruvkb/awf/models"
	"github.com/dhruvkb/awf/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Hello, World!",
		})
	})

	app.Post("/waveform", func(c *fiber.Ctx) error {
		payload := models.Request{Counts: []int{1e3}}
		err := c.BodyParser(&payload)
		if err != nil {
			return err
		}

		path, err := utils.DownloadFile(payload.Identifier, payload.Url)
		if err != nil {
			return err
		}
		defer func() { err = os.Remove(path) }()

		peakSets := make(map[int]models.PeakSet)
		for _, count := range payload.Counts {
			peaks, err := utils.GetPeaks(path, count, payload.Duration)
			if err != nil {
				return err
			}
			normalisedPeaks := utils.Normalise(peaks)
			peakSets[count] = models.PeakSet{
				Length: len(normalisedPeaks),
				Peaks:  normalisedPeaks,
			}
		}

		return c.JSON(fiber.Map{
			"success":    true,
			"identifier": payload.Identifier,
			"peak_sets":  peakSets,
		})
	})

	log.Fatal(app.Listen(":8888"))
}
