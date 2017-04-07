process test_1 {
  sequence seq1_1 {
    action act_4_1 {
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
    action act_5_1 {
      script { "{\"drugs\": [\"fanta\"]}" }
    }
  }
  sequence seq2_1 {
    action act_6_1 {
      script { "{\"drugs\": [\"dr pepper\"]}" }
    }
  }
}
