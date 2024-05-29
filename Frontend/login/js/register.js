const emailInput = document.getElementById("inputEmail");
const passwordInput = document.getElementById("inputPassword");
const nameInput = document.getElementById("inputName");
const dupPasswordInput = document.getElementById("inputConfirmPassword");

const submit = document.getElementById("submit-register");

submit.addEventListener("click", (e) => {
  e.preventDefault();
  const name = nameInput.value;
  const dupPassword = dupPasswordInput.value;
  const email = emailInput.value;
  const password = passwordInput.value;

  if (password !== dupPassword) {
    console.log("Passwords do not match");
  } else {
    console.log("Passwords match");
  }

  console.log(email);
  console.log(password);
  console.log(name);
  console.log(dupPassword);
});
