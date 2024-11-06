const formValues = {}
const formValidation = {}



/**
 *
 * @param {HTMLElement} htmlElement
 * @param {Function} validator
 * @param {string} value
 */
export const customValidator = (htmlElement, validator, value) => {
  if (validator(value)) {
    htmlElement.classList.add("valid");
    htmlElement.classList.remove("invalid");
  } else {
    htmlElement.classList.add("invalid");
    htmlElement.classList.remove("valid");
  }
}

/**
 * @param {string} password
 * @return {RegExpMatchArray}
 */
export const validatePassword = (password) => {
  const regExp = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/

  return password.match(regExp)
}

/**
 *
 * @param {string} password
 * @return {boolean}
 */
export const validateRepeatPassword = (password) => {
  return formValues["password"] === password
}


/**
 * @param {string} email
 * @returns {RegExpMatchArray}
 */
export const validateEmail = (email) => {
  const regExp = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

  return String(email)
    .toLowerCase()
    .match(regExp);
}


export const getValidationStatus = () => {
  return Object.values(formValidation).every((validationStatus) => !!validationStatus)
}


export const setFormValue = (valueKey, newValue, validator) => {
  formValues[valueKey] = newValue
  if (validator !== undefined) {
    formValidation[valueKey] = validator(newValue)
  }
}


export const submitSignForm = () => {
  console.log(formValues)
  if (!getValidationStatus()) {
    console.log("FORM IS INCORRECT")
    return false
  }
  console.log("FORM IS FINE")
  return true
}

export const clearSignForm = () => {
  for (let field in formValues) {
    if (formValues.hasOwnProperty(field)) {
      delete formValues[field]
    }
  }
  for (let field in formValidation) {
    if (formValidation.hasOwnProperty(field)) {
      delete formValidation[field]
    }
  }
}