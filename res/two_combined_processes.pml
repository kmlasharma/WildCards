process base {
  branch processes {
    sequence process_1 {
      sequence seq1_1 {
        iteration s_1 {
          selection s1_1 {
            action act_1_1 {
              script { "{\"drugs\": [\"oj\"]}" }
            }
            action test_act_2_1 {
              script { "{\"drugs\": [\"7up\"]}" }
            }
          }
          delay { "3 days" }
          loops { "3" }
        }
        branch branch1_1 {
          action act_4_1 {
            script { "{\"drugs\": [\"pepsi\"]}" }
          }
          action act_5_1 {
            script { "{\"drugs\": [\"milk\"]}" }
          }
        }
      }
    }
    sequence process_2 {
      sequence seq1_2 {
        action act_4_2 {
          script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
        }
        action act_5_2 {
          script { "{\"drugs\": [\"fanta\"]}" }
        }
        sequence2_2 {
          action act_6_2 {
            script { "{\"drugs\": [\"dr pepper\"]}" }
          }
      }
    }
  }
}
