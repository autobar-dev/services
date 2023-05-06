use fancy_regex::Regex;

pub fn verify_clear_password(password: String) -> bool {
    // lowercase required
    // uppercase required
    // digit required
    // special character required (!, @, #, $, %,^,&,*,(,), -, _, +, =, [, {, ], }, \, |, ;, :, ',
    // ", \, , <, ., >, /, ?)
    // minimum 8 characters
    let exp = r#"^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z])(?=.*[!@#$%^&*()\-_+=\[{\]}\\|;:'",<.>/?]).{8,}$"#;

    let re = Regex::new(exp).unwrap();

    re.is_match(password.as_str()).unwrap()
}
