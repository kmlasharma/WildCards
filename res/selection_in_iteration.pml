process test {
  iteration i {
    selection sel_1 {
      action act_1 {
        script { "{\"drugs\": [\"coke\"]}" }
      }
      action act_2 {
        script { "{\"drugs\": [\"7up\"]}" }
      }
    }
  }
}
