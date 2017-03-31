process test {
  sequence seq1 {
    action act_1 {
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
    delay { "30 sec" }
  }
  task t1 {
    action act_2 {
      script { "{}" }
    }
    delay { "20 min" }
  }
  delay { "5 hr" }
  delay { "4 day" }
  iteration iter1 {
    action act_3 {
      script { "{}" }
    }
    delay { "3 week" }
  }
}
