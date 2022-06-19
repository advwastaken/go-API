package main

import (
  "errors"
  "strings"
  "net/http"
  "go-api/school"
  "github.com/gin-gonic/gin"
)

var topSchools = new(school.topSchools)

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