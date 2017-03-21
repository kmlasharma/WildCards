process test {
  sequence seq1 {
    sequence seq2 {
       action act_1 {
         script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
       }
     }
     action act_1 {
       script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
     }
  }
}
