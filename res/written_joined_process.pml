  process merged {
    branch mergedpathways {
      sequence test_1 {
        sequence seq1_1 {
          delay { "30 secs" }
          action act_1_1 {
            script { "{\"drugs\":[\"coke\",\"7up\",\"pepsi\"]}" }
          }
        }
      }
      sequence test_2 {
        sequence seq1_2 {
          action act_1_2 {
            script { "{\"drugs\":[\"coke\",\"7up\",\"pepsi\"]}" }
          }
          delay { "30 secs" }
        }
        task t1_2 {
          action act_2_2 {
            script { "{\"drugs\":null}" }
          }
          delay { "30 secs" }
        }
        delay { "30 secs" }
        delay { "30 secs" }
        iteration iter1_2 {
          action act_3_2 {
            script { "{\"drugs\":null}" }
          }
          delay { "30 secs" }
        }
      }
      sequence test_3 {
        sequence seq1_3 {
          action act_4_3 {
            script { "{\"drugs\":null}" }
          }
        }
      }
    }
  }