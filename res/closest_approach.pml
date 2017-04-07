process test {
  selection sel_1 {
    action act_1 {
      script { "{\"drugs\": [\"oj\"]}" }
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
  sequence seq_2 {
    action act_5 {
      script { "{\"drugs\": [\"pepsi\"]}" }
    }
    delay { "1 day" }
    action act_5 {
      script { "{\"drugs\": [\"flat7up\"]}" }
    }
  }
}
