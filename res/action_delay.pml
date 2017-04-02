process test {
  sequence seq1 {
    action act_1 {
      delay { "30 sec" }
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
  }
}
