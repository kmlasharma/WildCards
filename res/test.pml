process test {
  sequence seq1 {
    iteration i {
      action act_1 {
        script { "{\"drugs\": [\"oj\"]}" }
      }
      action act_2 {
        script { "{\"drugs\": [\"milk\"]}" }
      }
      delay { "40" }
      loops { "3" }
    }
    action act_4 {
      script { "{\"drugs\": [\"coke\"]}" }
    }
    branch b1 {
      action act_5 {
        script { "{\"drugs\": [\"a\"]}" }
      }
      branch b2 {
        action act_7 {
          script { "{\"drugs\": [\"b\"]}" }
        }
        action act_8 {
          script { "{\"drugs\": [\"c\"]}" }
        }
      }
      action act_6 {
        script { "{\"drugs\": [\"d\"]}" }
      }
    }
  }
}
