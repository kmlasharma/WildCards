process test {
  sequence Andy {
    action Mary {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
  }
  sequence Andy {
    action Mary {}
  }
  sequence John {}
}
