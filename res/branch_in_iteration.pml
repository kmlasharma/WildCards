process test {
	iteration iter_1{
	  branch branch1 {
	    action act_1 {
	      script { "{\"drugs\": [\"coke\"]}" }
	    }
	    action act_2 {
	      script { "{\"drugs\": [\"pepsi\"]}" }
	    }
	  }
	}
}
