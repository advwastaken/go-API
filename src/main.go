package main

import (
  "errors"
  "strings"
  "net/http"
  "github.com/gin-gonic/gin"
)

type School struct {
  Name        string      `json:"name"`
  Location     string      `json:"Location"`
}

var topSchools = []School {
  { Name: "SVIS – Sainte Victoire International School", Location: "France" }, { Name: "The American School of Milan", Location: "Italy" }, { Name: "Ecole d’Humanité", Location: "Switzerland" }, { Name: "Akademeia High School", Location: "Poland" }, { Name: "Woodside Priory School", Location: "United states of America"},
}

func getAllSchools(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, topSchools)
}

func filterSchool(data string) (*School, error) {
  for i, school := range topSchools {
    if strings.EqualFold(data, school.Name) || strings.HasPrefix(data, school.Name) || strings.Contains(data, school.Name) {
      return &topSchools[i], nil
    }
  }
  return nil, errors.New("School Not found")
}

func getSchool(c * gin.Context) {
  data := c.Param("data")
  school, err := filterSchool(data)

  if err != nil {
    return
  }

  c.IndentedJSON(http.StatusOK, school)
}

func main() {
  router := gin.Default()
  router.GET("/schools", getAllSchools)
  router.GET("/school/:data", getSchool)
  router.Run("localhost:8080")
}