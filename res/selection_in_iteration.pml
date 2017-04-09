process test {
  iteration iter_1 {
    selection sel_1 {
      action act_1 {
        script { "{\"drugs\": [\"caffeine\"]}" }
      }
      action act_2 {
        script { "{\"drugs\": [\"alcohol\"]}" }
      }
    }
  }
}
