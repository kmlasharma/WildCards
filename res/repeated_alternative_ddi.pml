process test {
  iteration iter_1 {
    delay { "2 days" }
    selection sel_1 {
      action act_1 {
        script { "{\"drugs\": [\"caffeine\"]}" }
      }
      action act_2 {
        script { "{\"drugs\": [\"alcohol\"]}" }
      }
    }
    delay { "1 day" }
  }
}
