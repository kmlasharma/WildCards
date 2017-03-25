process test {
  sequence seq1 {
    action act_4 {
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
    action act_5 {
      script { "{\"drugs\": [\"fanta\"]}" }
    }
  }
  sequence seq2 {
    action act_6 {
      script { "{\"drugs\": [\"dr pepper\"]}" }
    }
  }
}
