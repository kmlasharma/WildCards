process test {
  sequence seq1 {
    action act_4 {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
  }
  iteration iter1 {
    action act_5 {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
    action act_6 {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
  }
}