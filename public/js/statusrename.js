function changeStatus (wartosc) {
  const status = {
    -5: "nie zainteresowany",
    -2: "raczej nie zainteresowany",
    0: "Skontaktuj",
    3: "Może zainteresowany",
    10: "Zainteresowany",
  }

  const StatusDOMEobj = document.querySelectorAll('.ustatus');

  StatusDOMEobj.forEach(obj => {
    obj.textContent = status[obj.textContent] || obj.textContent;
  })
};
changeStatus()
console.log("tekst")
