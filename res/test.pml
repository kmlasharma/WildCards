process test {
  sequence seq1 {
    iteration s {
      selection s1 {
        action act_1 {
          script { "{\"drugs\": [\"oj\"]}" }
        }
        action act_2 {
          script { "{\"drugs\": [\"7up\"]}" }
        }
      }
      loops { "3" }
    }
    branch branch1 {
      action act_4 {
        script { "{\"drugs\": [\"pepsi\"]}" }
      }
      action act_5 {
        script { "{\"drugs\": [\"milk\"]}" }
      }
    }
    action act_6 {
      script { "{\"drugs\": [\"caffeine\"]}" }
    }
    action act_7 {
      script { "{\"drugs\": [\"alcohol\"]}" }
    }
  }
}
