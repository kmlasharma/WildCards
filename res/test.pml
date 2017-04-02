process test {
  sequence seq1 {
    action act_3 {
      script { "{\"drugs\": [\"coke\"]}" }
    }
    iteration i {
      selection s1 {
        action act_1 {
          script { "{\"drugs\": [\"oj\"]}" }
        }
        action act_2 {
          script { "{\"drugs\": [\"milk\"]}" }
        }
      }
      delay { "3 days" }
      loops { "3" }
    }
    action act_4 {
      script { "{\"drugs\": [\"coke\"]}" }
    }
  }
}
