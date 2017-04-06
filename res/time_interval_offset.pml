process test {
        wait { "Wednesday morning" }
	iteration iter_1{
	  branch branch1 {
            sequence seq1 {
	      action act_1 {
	        script { "{\"drugs\": [\"pepsi\"]}" }
	      }
 	    }
	    sequence seq2 {
	      wait { "evening" }
	      action act_2 {
	        script { "{\"drugs\": [\"flat7up\"]}" }
	      }
	    }
	  }
	}
}
