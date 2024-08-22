function changeStatus (wartosc) {
  const status = {
    "-5": "nie zainteresowany",
    "-2": "raczej nie zainteresowany",
    0: "Skontaktuj",
    5: "Może zainteresowany",
    10: "Zainteresowany",
  }

  const StatusDOMEobj = document.querySelectorAll('.ustatus');

  StatusDOMEobj.forEach(obj => {
    const orginalvalue = parseInt(obj.textContent, 10);

    obj.textContent = status[orginalvalue] || obj.textContent;
    if (orginalvalue < 0) {
      obj.classList.add("bg-danger")
    } else if (orginalvalue > 0 && orginalvalue < 10) {
      obj.classList.add("bg-info")
    } else if (orginalvalue > 5 && orginalvalue < 11) {
obj.classList.add("bg-success")
    } else {
      obj.classList.add("bg-secondary")
    }
  })
};
changeStatus()
console.log("tekst")
