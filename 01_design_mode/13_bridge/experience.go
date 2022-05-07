package bridge

import "fmt"

// Experience 经历
type Experience interface {
  Describe() string // 描述经历
}

// travelExperience 旅游经历
type travelExperience struct {
  subject  string
  traffic  Traffic
  location Location
}

// NewTravelExperience 创建旅游经历，包括主题、交通方式、地点
func NewTravelExperience(subject string, traffic Traffic, location Location) *travelExperience {
  return &travelExperience{
    subject:  subject,
    traffic:  traffic,
    location: location,
  }
}

// Describe 描述旅游经历
func (t *travelExperience) Describe() string {
  return fmt.Sprintf("%s is to %s %s and %s", t.subject, t.location.Name(), t.traffic.Transport(), t.location.PlaySports())
}

// adventureExperience 探险经历
type adventureExperience struct {
  survivalTraining string
  travelExperience
}

// NewAdventureExperience 创建探险经历，包括探险需要的培训，其他的与路由参数类似
func NewAdventureExperience(training string, subject string, traffic Traffic, location Location) *adventureExperience {
  return &adventureExperience{
    survivalTraining: training,
    travelExperience: *NewTravelExperience(subject, traffic, location),
  }
}

// Describe 描述探险经历
func (a *adventureExperience) Describe() string {
  return fmt.Sprintf("after %s, %s", a.survivalTraining, a.travelExperience.Describe())
}
