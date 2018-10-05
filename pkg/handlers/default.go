package handlers 

// ResponseHandler - Default Response Handler
func ResponseHandler(w http.ResponseWriter, r *http.Request, m string, success bool, err *string, res interface{}, transaction *types.Transaction) {
	errorString := ""

	if err != nil {
		errorString = *err
	}

	responseStruct := response.DefaultResponse{
		Message:     m,
		Success:     success,
		Error:       errorString,
		Response:    &res,
		Transaction: nil,
		Endpoint:    r.URL.String(),
	}

	if transaction != nil {
		responseStruct.FormatTransactionResponse(transaction.Hash().String())
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false) // So we can have an & come through in our URL's
	parseErr := enc.Encode(responseStruct)

	if parseErr != nil {
		ErrorHandler(w, r, "Could not parse response JSON", parseErr, http.StatusInternalServerError)
		return
	}

	return
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New(r.URL.String() + " not found in available routes")
	ErrorHandler(w, r, "Invalid request, check parameters and try again", err, http.StatusNotFound)
	return
}

// ErrorHandler - Default Error Handler
func ErrorHandler(w http.ResponseWriter, r *http.Request, m string, e error, statusCode int) {
	w.WriteHeader(statusCode)

	var err string
	if e != nil {
		err = e.Error()
	} else {
		err = "Error message could not be parsed"
	}

	ResponseHandler(w, r, m, false, &err, nil, nil)

	return
}