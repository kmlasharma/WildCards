process test {
  iteration iter_1 {
    selection sel_1 {
      action act_1 {
        script { "{\"drugs\": [\"coke\"]}" }
      }
      action act_2 {
        script { "{\"drugs\": [\"7up\"]}" }
      }
    }
    sequence seq_1 {
      action act_3 {
        script { "{\"drugs\": [\"caffeine\"]}" }
      }
      delay { "5 days" }
      action act_4 {
        script { "{\"drugs\": [\"alcohol\"]}" }
      }
    }
  }
}
