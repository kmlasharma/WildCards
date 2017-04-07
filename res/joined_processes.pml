process merged {
  branch mergedpathways {
    sequence test_1 {
      sequence seq1_1 {
        delay { "30 sec" }
        action act_1_1 {
          script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
        }
      }
    }
    sequence test_2 {
      sequence seq1_2 {
        action act_1_2 {
          script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
        }
        delay { "30 sec" }
      }
      task t1_2 {
        action act_2_2 {
          script { "{}" }
        }
        delay { "20 min" }
      }
      delay { "5 hr" }
      delay { "4 day" }
      iteration iter1_2 {
        action act_3_2 {
          script { "{}" }
        }
        delay { "3 week" }
      }
    }
    sequence test_3 {
      sequence seq1_3 {
        action act_4_3 {
          script { "{}" }
        }
      }
    }
    sequence test_4 {
      task t1_4 {
        action act_1_4 {
          script { "{}" }
        }
      }
      task t2_4 {
        action act_2_4 {
          script { "{}" }
        }
      }
    }
  }
}

