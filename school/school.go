package school

type School struct {
  Name        string      `json:"name"`
  Location     string      `json:"Location"`
}

var topSchools = []School {
  { Name: "SVIS – Sainte Victoire International School", Location: "France" }, { Name: "The American School of Milan", Location: "Italy" }, { Name: "Ecole d’Humanité", Location: "Switzerland" }, { Name: "Akademeia High School", Location: "Poland" }, { Name: "Woodside Priory School", Location: "United states of America"},
}