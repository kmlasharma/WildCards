process test {
  sequence seq_1 {
    delay { "2 days" }
    action act_1 {
      script { "{\"drugs\": [\"paracetamol\"]}" }
    }
    wait{ "Monday" }
    action act_2 {
      script { "{\"drugs\": [\"alcohol\"]}" }
    }
  }
}
