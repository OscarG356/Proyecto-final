const emailInput = document.getElementById("inputEmail");
const passwordInput = document.getElementById("inputPassword");

const submit = document.getElementById("submit-login");

submit.addEventListener("click", (e) => {
  const email = emailInput.value;
  const password = passwordInput.value;

  console.log(email);
  console.log(password);
});
