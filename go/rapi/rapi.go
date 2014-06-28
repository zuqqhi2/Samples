/*
  Rakuten Webservice API Interface 
*/
package RApi
 
import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

// APISetting Information
type RApi struct {
  BaseURL       string
  ApplicationId string
  AffiliateId   string
  RequestURL    string
}

// Read configuration
func (api *RApi) New() {
  file, _ := os.Open("apiconf.json")
  decoder := json.NewDecoder(file)
  //configuration := api{}
  err := decoder.Decode(&api)
  if err != nil {
    fmt.Println("error:", err)
  }
  
  url := api.BaseURL + "&keyword=" + "%E6%A5%BD%E5%A4%A9"
  url += "&applicationId=" + api.ApplicationId
  url += "&affiliateId=" + api.AffiliateId
  api.RequestURL = url
}

// Get items from Rakuten API  
func (api *RApi) GetRItems() string { 
  return Get(api.RequestURL)
}

// Get result from url with http access
func Get(url string) string {
    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
        return string(contents)
    }
    return ""
}
