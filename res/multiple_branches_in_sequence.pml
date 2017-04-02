process test {
	sequence seq1{
	  branch branch1 {
	    action act_1 {
	      script { "{\"drugs\": [\"coke\"]}" }
	    }
	    action act_2 {
	      script { "{\"drugs\": [\"pepsi\"]}" }
	    }
	  }

	  branch branch2 {
	    action act_3 {
	      script { "{\"drugs\": [\"milk\"]}" }
	    }

	    action act_4 {
	      script { "{\"drugs\": [\"oj\"]}" }
	    }
	  }
	}
}
