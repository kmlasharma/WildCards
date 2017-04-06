process test {
  sequence seq1 {
    delay { "30 sec" }
    action act_1 {
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
  }
}
