package pml

var (
	parentBranchName    = ""
	parentSequenceName  = ""
	parentSelectionName = ""
)

type ActionWrapper struct {
	action       *Action
	currentDelay Delay
}

type Params struct {
	drugPairs                []DrugPair
	sequentialActionWrappers []ActionWrapper
	parallelActionWrappers   []ActionWrapper
	selectionActionWrappers  []ActionWrapper
	currentDelay             Delay
}

func (p *Params) addAction(action *Action, inIter bool) {
	for _, wrapper := range p.parallelActionWrappers {
		action1 := wrapper.action
		actionDelay := wrapper.currentDelay
		for _, drugA := range action1.Drugs {
			for _, drugB := range action.Drugs {
				pair := DrugPair{
					DrugA:      drugA,
					DrugB:      drugB,
					Delay:      p.currentDelay - actionDelay,
					ddiType:    ParallelType,
					parentName: parentBranchName,
				}
				p.drugPairs = append(p.drugPairs, pair)
			}
		}
	}

	for _, wrapper := range p.sequentialActionWrappers {
		action1 := wrapper.action
		actionDelay := wrapper.currentDelay
		for _, drugA := range action1.Drugs {
			for _, drugB := range action.Drugs {
				pair := DrugPair{
					DrugA:      drugA,
					DrugB:      drugB,
					Delay:      p.currentDelay - actionDelay,
					ddiType:    SequentialType,
					parentName: parentSequenceName,
				}
				p.drugPairs = append(p.drugPairs, pair)
			}
		}
	}

	if inIter {
		for _, wrapper := range p.selectionActionWrappers {
			action1 := wrapper.action
			actionDelay := wrapper.currentDelay
			for _, drugA := range action1.Drugs {
				for _, drugB := range action.Drugs {
					pair := DrugPair{
						DrugA:      drugA,
						DrugB:      drugB,
						Delay:      p.currentDelay - actionDelay,
						ddiType:    RepeatedAlternativeDDIType,
						parentName: parentSelectionName,
					}
					p.drugPairs = append(p.drugPairs, pair)
				}
			}
		}
	} else {
		for _, wrapper := range p.selectionActionWrappers {
			action1 := wrapper.action
			actionDelay := wrapper.currentDelay
			for _, drugA := range action1.Drugs {
				for _, drugB := range action.Drugs {
					pair := DrugPair{
						DrugA:      drugA,
						DrugB:      drugB,
						Delay:      p.currentDelay - actionDelay,
						ddiType:    AlternativeNonDDIType,
						parentName: parentSelectionName,
					}
					p.drugPairs = append(p.drugPairs, pair)
				}
			}
		}
	}
}

func newParams() Params {
	return Params{
		drugPairs:                []DrugPair{},
		sequentialActionWrappers: []ActionWrapper{},
		parallelActionWrappers:   []ActionWrapper{},
		selectionActionWrappers:  []ActionWrapper{},
		currentDelay:             0,
	}
}

func (ele *Element) FindDrugPairs() []DrugPair {
	params := ele.parseElement(newParams(), false)
	return params.drugPairs
}

func (ele *Element) FindSequentialDrugPairs() (results []DrugPair) {
	pairs := ele.FindDrugPairs()
	for _, pair := range pairs {
		if pair.ddiType == SequentialType {
			results = append(results, pair)
		}
	}
	return
}

func (ele *Element) FindParallelDrugPairs() (results []DrugPair) {
	pairs := ele.FindDrugPairs()
	for _, pair := range pairs {
		if pair.ddiType == ParallelType {
			results = append(results, pair)
		}
	}
	return
}

func (ele *Element) FindAlternativeNonDDIDrugPairs() (results []DrugPair) {
	pairs := ele.FindDrugPairs()
	for _, pair := range pairs {
		if pair.ddiType == AlternativeNonDDIType {
			results = append(results, pair)
		}
	}
	return
}

func (ele *Element) FindRepeatedAlternativeDrugPairs() (results []DrugPair) {
	pairs := ele.FindDrugPairs()
	for _, pair := range pairs {
		if pair.ddiType == RepeatedAlternativeDDIType {
			results = append(results, pair)
		}
	}
	return
}

func (ele *Element) parseElement(params Params, inIter bool) Params {
	parentSequenceName = ele.Name
	for _, child := range ele.Children {
		switch child.Type() {
		case ActionType:
			action := child.(*Action)
			params.addAction(action, inIter)
			params.sequentialActionWrappers = append(params.sequentialActionWrappers, ActionWrapper{action: action, currentDelay: params.currentDelay})
		case DelayType:
			delay := child.(Delay)
			params.currentDelay += delay
		case WaitType:
			offset := child.(Wait)
			params.currentDelay = calculateOffsetDelay(params.currentDelay, offset)
		default:
			var updatedParams Params
			el := child.(*Element)
			switch el.elementType {
			case BranchType:
				updatedParams = el.parseBranch(params, inIter)
			case IterationType:
				updatedParams = el.parseIteration(params)
			case SelectionType:
				updatedParams = el.parseSelection(params, inIter)
			default:
				updatedParams = el.parseElement(params, inIter)
			}
			params.drugPairs = updatedParams.drugPairs
			params.sequentialActionWrappers = append(params.sequentialActionWrappers, updatedParams.sequentialActionWrappers...)
			params.currentDelay += updatedParams.currentDelay
		}
	}
	return params
}

func (ele *Element) parseBranch(params Params, inIter bool) Params {
	// One element - current delay which is used for an direct delays
	parentBranchName = ele.Name
	delays := []Delay{params.currentDelay}
	for _, child := range ele.Children {
		switch child.Type() {
		case ActionType:
			action := child.(*Action)
			params.addAction(action, inIter)
			params.parallelActionWrappers = append(params.parallelActionWrappers, ActionWrapper{action: action, currentDelay: params.currentDelay})
		case DelayType:
			delay := child.(Delay)
			delays[0] += delay
		default:
			var updatedParams Params
			el := child.(*Element)
			switch el.elementType {
			case BranchType:
				updatedParams = el.parseBranch(params, inIter)
			case IterationType:
				updatedParams = el.parseIteration(params)
			case SelectionType:
				updatedParams = el.parseSelection(params, inIter)
			default:
				updatedParams = el.parseElement(params, inIter)
			}
			delays = append(delays, updatedParams.currentDelay)
			params.drugPairs = updatedParams.drugPairs
			params.parallelActionWrappers = append(params.parallelActionWrappers, updatedParams.sequentialActionWrappers...)
		}
	}
	// Set delay as max delay
	maxDelay := Delay(0)
	for _, delay := range delays {
		if delay > maxDelay {
			maxDelay = delay
		}
	}
	params.currentDelay = maxDelay
	params.sequentialActionWrappers = append(params.sequentialActionWrappers, params.parallelActionWrappers...)
	params.parallelActionWrappers = []ActionWrapper{}
	return params
}

func (ele *Element) parseSelection(params Params, inIter bool) Params {
	// One element - current delay which is used for an direct delays
	parentSelectionName = ele.Name
	delays := []Delay{params.currentDelay}
	for _, child := range ele.Children {
		switch child.Type() {
		case ActionType:
			action := child.(*Action)
			params.addAction(action, inIter)
			params.selectionActionWrappers = append(params.selectionActionWrappers, ActionWrapper{action: action, currentDelay: params.currentDelay})
		case DelayType:
			delay := child.(Delay)
			delays[0] += delay
		default:
			var updatedParams Params
			el := child.(*Element)
			switch el.elementType {
			case BranchType:
				updatedParams = el.parseBranch(params, inIter)
			case IterationType:
				updatedParams = el.parseIteration(params)
			case SelectionType:
				updatedParams = el.parseSelection(params, inIter)
			default:
				updatedParams = el.parseElement(params, inIter)
			}
			delays = append(delays, updatedParams.currentDelay)
			params.drugPairs = append(params.drugPairs, updatedParams.drugPairs...)
			params.selectionActionWrappers = append(params.selectionActionWrappers, updatedParams.sequentialActionWrappers...)
		}
	}

	// Set delay as max delay
	maxDelay := Delay(0)
	for _, delay := range delays {
		if delay > maxDelay {
			maxDelay = delay
		}
	}
	params.currentDelay = maxDelay
	params.sequentialActionWrappers = append(params.sequentialActionWrappers, params.selectionActionWrappers...)
	params.selectionActionWrappers = []ActionWrapper{}
	return params
}

func (ele *Element) parseIteration(params Params) Params {
	updatedParams := ele.parseElement(params, true)
	iterationDelay := updatedParams.currentDelay - params.currentDelay
	for _, pair := range updatedParams.drugPairs {
		delay := iterationDelay - pair.Delay
		var minDelay int
		if delay < pairDelay {
			minDelay := delay
		} else {
			minDelay := pair.Delay
		}
		newPair := DrugPair{
			DrugA:      pair.DrugB,
			DrugB:      pair.DrugA,
			Delay:      minDelay,
			ddiType:    pair.ddiType,
			parentName: pair.parentName,
		}
		updatedParams.drugPairs = append(updatedParams.drugPairs, newPair)
	}
	totalDelay := Delay(int(iterationDelay) * (ele.Loops - 1))
	updatedParams.currentDelay += totalDelay
	return updatedParams
}
