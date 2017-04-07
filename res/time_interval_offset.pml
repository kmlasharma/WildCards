process test {
<<<<<<< HEAD
  sequence seq_1 {
    delay { "2 days" }
    action act_1 {
      script { "{\"drugs\": [\"paracetamol\"]}" }
    }
    wait{ "Monday" }
    action act_2 {
      script { "{\"drugs\": [\"alcohol\"]}" }
    }
  }
=======
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
>>>>>>> 890a5e818cf6adc3f46f151a260b0778357c8e3f
}
