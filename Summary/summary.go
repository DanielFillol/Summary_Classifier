package Summary

//InferredDecision is a simple struct that contains a legal summary the classification of it.
type InferredDecision struct {
	Summary string `json:"summary,omitempty"`
	Class   string `json:"classificcao,omitempty"`
}

//Classify attempts up to four times to classify a text of an legal summary.
//Every time it tries a longer length on the function ReturnSummaryClass
func Classify(summary string) (InferredDecision, error) {
	class, err := firstTry(summary)
	if err != nil {
		return InferredDecision{}, err
	}
	if class != "" {
		return InferredDecision{summary, class}, nil
	}

	class, err = secondTry(summary)
	if err != nil {
		return InferredDecision{}, err
	}
	if class != "" {
		return InferredDecision{summary, class}, nil
	}

	class, err = thirdTry(summary)
	if err != nil {
		return InferredDecision{}, err
	}
	if class != "" {
		return InferredDecision{summary, class}, nil
	}

	class, err = lastTry(summary)
	if err != nil {
		return InferredDecision{}, err
	}
	if class != "" {
		return InferredDecision{summary, class}, nil
	}

	return InferredDecision{summary, NotMapped}, nil
}

//The first try attempts to classify the summary utilizing only the las 16 char
func firstTry(summary string) (string, error) {
	class, err := ReturnSummaryClass(summary, 16)
	if err != nil {
		return "", err
	}

	if class != NotMapped {
		return class, nil
	}
	return "", nil
}

//The second try attempts to classify the summary utilizing 1/3 of total length
func secondTry(summary string) (string, error) {
	class, err := ReturnSummaryClass(summary, len(summary)/3)
	if err != nil {
		return "", err
	}

	if class != NotMapped {
		return class, nil
	}
	return "", nil
}

//The third try attempts to classify the summary utilizing 1/2 of total length
func thirdTry(summary string) (string, error) {
	class, err := ReturnSummaryClass(summary, len(summary)/2)
	if err != nil {
		return "", err
	}

	if class != NotMapped {
		return class, nil
	}
	return "", nil
}

//The third try attempts to classify the summary utilizing the full length
func lastTry(summary string) (string, error) {
	class, err := ReturnSummaryClass(summary, len(summary))
	if err != nil {
		return "", err
	}
	if class != NotMapped {
		return class, nil
	}
	return "", nil
}
