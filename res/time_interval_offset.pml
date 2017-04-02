process test {
  wait{ "Monday" }
  iteration iter_1 {
    loops { "5" }
    action act_1 {
      script { "{\"drugs\": [\"coke\", \"7up\", \"pepsi\"]}" }
    }
    delay { "3 days" }
  }
}
