process test {
  sequence Andy {
    action Mary {
      script { "{\"drugs\": [\"Plavix\", \"Lipitor\", \"Nexium\"]}" }
    }
  }
  sequence Mary {
    action Andy {}
  }
  sequence John {
    action John {}
  }
  sequence Beth {}
}
