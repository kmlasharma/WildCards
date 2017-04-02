package pml

type ActionWrapper struct {
	action       Action
	currentDelay Delay
}

func (ele *Element) FindDrugPairs() []DrugPair {
	pairs, _, _ := ele.parsePossibleDDIs([]ActionWrapper{}, Delay(0), ele.Name)
	return pairs
}

func (ele *Element) parseIterationPossibleDDIs(actions []ActionWrapper, delay Delay, parentName string) (pairs []DrugPair, newActions []ActionWrapper, newDelay Delay) {
	newDelay = delay
	newPairs, newActions, updatedNewDelay := ele.parsePossibleDDIs(actions, newDelay, parentName)
	iterationDelay := updatedNewDelay - delay
	for _, pair := range newPairs {
		delay := iterationDelay - pair.delay
		newPair := DrugPair{DrugA: pair.DrugB, DrugB: pair.DrugA, delay: delay, ddiType: pair.ddiType, parentName: pair.parentName}
		pairs = append(pairs, pair)
		pairs = append(pairs, newPair)
	}
	newDelay += Delay(int(iterationDelay) * ele.Loops) // Account for loops
	return
}

func (ele *Element) parseBranchPossibleDDIs(actions []ActionWrapper, oldParallelActions []ActionWrapper, delay Delay, parentName string) (pairs []DrugPair, newParallelActions []ActionWrapper, newDelay Delay) {
	newParallelActions = oldParallelActions
	newDelay = delay
	for _, child := range ele.Children {
		if child.IsSubElementType() {
			el := child.(*Element)
			var newPairs []DrugPair
			if el.elementType == BranchType {
				newPairs, newParallelActions, newDelay = el.parseBranchPossibleDDIs(actions, newParallelActions, newDelay, ele.Name)
				pairs = append(pairs, newPairs...)
			} else {
				newPairs, newParallelActions, newDelay = el.parseBranchPossibleDDIs(actions, newParallelActions, newDelay, parentName)
				pairs = append(pairs, newPairs...)
			}
		} else if child.Type() == ActionType {
			action1 := child.(Action)
			for _, wrapper := range newParallelActions {
				action2 := wrapper.action
				actionDelay := wrapper.currentDelay
				for _, drugA := range action2.Drugs {
					for _, drugB := range action1.Drugs {
						pair := DrugPair{DrugA: drugA, DrugB: drugB, delay: newDelay - actionDelay, ddiType: ParallelType, parentName: parentName}
						pairs = append(pairs, pair)
					}
				}
			}
			for _, wrapper := range actions {
				action2 := wrapper.action
				actionDelay := wrapper.currentDelay
				for _, drugA := range action2.Drugs {
					for _, drugB := range action1.Drugs {
						pair := DrugPair{DrugA: drugA, DrugB: drugB, delay: newDelay - actionDelay, ddiType: SequentialType, parentName: parentName}
						pairs = append(pairs, pair)
					}
				}
			}
			wrapper := ActionWrapper{action: action1, currentDelay: newDelay}
			newParallelActions = append(newParallelActions, wrapper)
		}
	}
	return
}

func (ele *Element) parsePossibleDDIs(oldActions []ActionWrapper, currentDelay Delay, parentName string) (pairs []DrugPair, newActions []ActionWrapper, newDelay Delay) {
	newActions = oldActions
	newDelay = currentDelay
	for _, child := range ele.Children {
		if child.IsSubElementType() {
			el := child.(*Element)
			var newPairs []DrugPair
			if el.elementType == BranchType {
				parallelActions := []ActionWrapper{}
				newPairs, parallelActions, newDelay = el.parseBranchPossibleDDIs(newActions, parallelActions, newDelay, el.Name)
				newActions = append(newActions, parallelActions...)
				pairs = append(pairs, newPairs...)
			} else if el.elementType == IterationType {
				newPairs, newActions, newDelay = el.parseIterationPossibleDDIs(newActions, newDelay, el.Name)
				pairs = append(pairs, newPairs...)
			} else {
				newPairs, newActions, newDelay = el.parsePossibleDDIs(newActions, newDelay, el.Name)
				pairs = append(pairs, newPairs...)
			}
		} else if child.Type() == ActionType {
			action1 := child.(Action)
			// For all previous actions, check for DDI
			// with delay = newDelay
			for _, wrapper := range newActions {
				action2 := wrapper.action
				actionDelay := wrapper.currentDelay
				for _, drugA := range action2.Drugs {
					for _, drugB := range action1.Drugs {
						pair := DrugPair{DrugA: drugA, DrugB: drugB, delay: newDelay - actionDelay, parentName: parentName}
						pairs = append(pairs, pair)
					}
				}
			}
			wrapper := ActionWrapper{action: action1, currentDelay: newDelay}
			newActions = append(newActions, wrapper)
		} else if child.Type() == DelayType {
			delay := child.(Delay)
			newDelay += delay
		}
	}
	return
}
