async function getJSONData() {
  const response = await fetch("/status");
  const dataJSON = await response.json();

  if (response.ok) {
    const adresURL = dataJSON[0].StateDisplayLink;
    const actualURL = window.location.href;
    const actualLink = actualURL
      .replace(window.location.href, "")
      .replace(":" + window.location.port, "");

    if (adresURL !== actualLink) {
      console.log(actualLink);
      // window.location.assign(adresURL);
    } else {
      console.log("Game status not changed");
    }
  } else {
    console.log("Game status not recived, ERROR:", response.status);
  }
}

setInterval(getJSONData, 6000);
