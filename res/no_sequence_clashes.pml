process test {
  sequence mySeq {
    action act_4 {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
  }
  sequence myOtherSeq {
    action act_5 {

    }
  }
}
