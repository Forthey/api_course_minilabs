import {
    clearSignForm,
    customValidator, getValidationStatus,
    setFormValue,
    submitSignForm,
    validateEmail,
    validatePassword,
    validateRepeatPassword
} from "./utils.js"

// Выписываем все айдишники HTMl-элементов в константы для переиспользования
const first_name_id = 'first_name'
const last_name_id = 'last_name'
const password_class = 'password'
const password_repeat_id = 'password-repeat'
const email_class = 'email'

const sign_button_class = 'sign_button'
const sign_up_form_id = 'sign_up_form'
const sign_in_link_id = 'sign_in_link'

const sign_in_form_id = 'sign_in_form'
const sign_up_link_id = 'sign_up_link'

const sign_in_form = document.getElementById(sign_in_form_id).getElementsByTagName('form')[0]
const first_name = document.getElementById(first_name_id)
const last_name = document.getElementById(last_name_id)
const email_fields = document.getElementsByClassName(email_class)
const password_fields = document.getElementsByClassName(password_class)
const password_repeat = document.getElementById(password_repeat_id)
const sign_buttons = document.getElementsByClassName(sign_button_class)
const switch_to_sign_in = document.getElementById(sign_in_link_id)

const sign_up_form = document.getElementById(sign_up_form_id).getElementsByTagName('form')[0]
const switch_to_sign_up = document.getElementById(sign_up_link_id)

first_name.oninput = (e) => setFormValue(first_name_id, e.target.value)

last_name.oninput = (e) => setFormValue(last_name_id, e.target.value)

for (let i = 0; i < email_fields.length; i++) {
    const email = email_fields.item(i)
    email.oninput = (e) => setFormValue(email_class, e.target.value, validateEmail)
    email.onblur = (_) => {
        sign_buttons.item(i).disabled = !getValidationStatus()
    }
}

for (let i = 0; i < password_fields.length; i++) {
    const password = password_fields.item(i)
    password.oninput = (e) => setFormValue(password_class, e.target.value, validatePassword)
    password.onblur = (e) => {
        customValidator(password, validatePassword, e.target.value)
        sign_buttons.item(i).disabled = !getValidationStatus()
    }

    sign_buttons.item(i).onclick = (e) => {
        e.preventDefault()
        submitSignForm()
    }
}

password_repeat.oninput = (e) => setFormValue(password_repeat_id, e.target.value, validateRepeatPassword)
password_repeat.onblur = (e) => {
    customValidator(password_repeat, validateRepeatPassword, e.target.value)
    sign_buttons.item(0).disabled = !getValidationStatus()
}

switch_to_sign_in.onclick = (e) => {
    sign_up_form.reset()
    clearSignForm()
    document.getElementById(sign_up_form_id).style.display = "none"
    document.getElementById(sign_in_form_id).style.display = ""
}

switch_to_sign_up.onclick = (e) => {
    sign_in_form.reset()
    clearSignForm()
    document.getElementById(sign_in_form_id).style.display = "none"
    document.getElementById(sign_up_form_id).style.display = ""
}

