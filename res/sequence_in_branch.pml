process test {
	branch branch1 {
		sequence seq_1 {
			action act_1 {
  				script { "{\"drugs\": [\"coke\"]}" }
	 		}
			action act_2 {
  				script { "{\"drugs\": [\"pepsi\"]}" }
			}
		}

		sequence seq_2 {
			action act_3 {
  				script { "{\"drugs\": [\"7up\"]}" }
	 		}
			action act_4 {
  				script { "{\"drugs\": [\"club\"]}" }
			}
		}
	}
}
