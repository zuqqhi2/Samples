package RApi

import (
  "testing"
)

func TestGet(t *testing.T) {
  actual := Get("https://app.rakuten.co.jp/services/api/IchibaItem/Search/20140222?format=json&keyword=%E6%A5%BD%E5%A4%A9&applicationId=a34f840ac7ab2fd31591e5a05ebcbe94&affiliateId=0d898f3c.ddb798c6.0d898f3d.7fdd16b0")
  if actual == "" {
    t.Errorf("got %s\n", actual)
  }  
} 


func TestGetRItems(t *testing.T) {
  api := RApi{}
  api.New()
 
  actual := api.GetRItems()
  if actual == "" {
    t.Errorf("got %s\n", actual)
  }
}
