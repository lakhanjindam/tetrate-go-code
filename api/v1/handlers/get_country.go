package v1handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	v1models "github.com/tetrate-go-code/api/v1/models"
	"github.com/tetrate-go-code/pkg/utils"
)

const CountryAPIEndpoint = "https://restcountries.com/v3.1/name"

// GetCountryHandler godoc
//
// @Summary		Get country information
// @Description	Get information about a country by name
// @Tags			v1
// @Accept		json
// @Produce		json
// @Param		name	query	string	true	"Country name"
// @Success		200	{object}	v1models.CountryResponse	"OK"
// @Failure		400	{object}	string	"Bad request"
// @Failure		404	{object}	string	"Country not found"
// @Failure		500	{object}	string	"Internal server error"
// @Router		/api/v1/country [get]
func GetCountryHandler(c *gin.Context) {
	request := v1models.CountryRequest{}

	// Bind the query parameters to the request struct
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build the API URL with the country name
	url := fmt.Sprintf("%s/%s?fullText=true", CountryAPIEndpoint, request.Name)

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to call country API: %v", err)})
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
			return
		}
		c.JSON(resp.StatusCode, gin.H{"error": "Country API returned an error"})
		return
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read country data"})
		return
	}

	// Unmarshal the JSON body into a struct
	var countryData []v1models.Country
	err = json.Unmarshal(body, &countryData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse country data"})
		return
	}

	// No data present for the country
	if len(countryData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No countries found"})
		return
	}

	populationStr := utils.FormatNumberWithSuffix(countryData[0].Population)

	data := v1models.CountryData{
		Name:        countryData[0].Name,
		Independent: countryData[0].Independent,
		UNMember:    countryData[0].UNMember,
		Currencies:  countryData[0].Currencies,
		Capital:     countryData[0].Capital,
		Region:      countryData[0].Region,
		Languages:   countryData[0].Languages,
		Population:  populationStr,
		LatLng:      countryData[0].LatLng,
		Continents:  countryData[0].Continents,
	}

	response := v1models.CountryResponse{
		Data: data,
	}

	c.JSON(http.StatusOK, response)
}
