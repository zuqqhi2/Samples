package main

import (
  "fmt"
  //jubatus "github.com/jubatus/jubatus-go-client/lib/classifier"
  classifier "github.com/jubatus/jubatus-go-client/lib/classifier"
  common "github.com/jubatus/jubatus-go-client/lib/common"
)

func main() {
  cli, err := classifier.NewClassifierClient("localhost:9199", "hoge")
  if err != nil {
    fmt.Println(err)
    return
  }
  
  var male_train_data [3]map[string]string = [3]map[string]string{
    {"hair":"short", "top":"sweater", "bottom":"jeans"},
    {"hair":"short","top":"jacket","bottom":"chino"},
    {"hair":"long","top":"T shirt","bottom":"jeans"}, 
  }
  
  var female_train_data [3]map[string]string = [3]map[string]string{
    {"hair":"long", "top":"shirt", "bottom":"skirt"},
    {"hair":"short", "top":"T shirt", "bottom":"jeans"},
    {"hair":"long", "top":"jacket", "bottom":"skirt"},
  }

  for idx := 0; idx < len(male_train_data); idx++ {
    datum := common.NewDatum()
    datum.AddString("hair", male_train_data[idx]["hair"])
    datum.AddString("top", male_train_data[idx]["top"])
    datum.AddString("bottom", male_train_data[idx]["bottom"])
    cli.Train([]classifier.LabeledDatum{classifier.LabeledDatum{"male", datum}})
  }
  for idx := 0; idx < len(female_train_data); idx++ {
    datum := common.NewDatum()
    datum.AddString("hair", female_train_data[idx]["hair"])
    datum.AddString("top", female_train_data[idx]["top"])
    datum.AddString("bottom", female_train_data[idx]["bottom"])
    cli.Train([]classifier.LabeledDatum{classifier.LabeledDatum{"female", datum}})
  }

  datum1 := common.NewDatum()
  datum1.AddString("hair", "short")
  datum1.AddString("top", "T shirt")
  datum1.AddString("bottom", "jeans")

  ret := cli.Classify([]common.Datum{datum1})
  fmt.Println(ret)

  datum2 := common.NewDatum()
  datum2.AddString("hair", "long")
  datum2.AddString("top", "shirt")
  datum2.AddString("bottom", "skirt")

  ret = cli.Classify([]common.Datum{datum2})
  fmt.Println(ret)
  
}
