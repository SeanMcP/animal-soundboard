(function () {
  const app = document.getElementById("app");
  const audio = document.getElementById("audio");

  function clearActive() {
    document
      .querySelectorAll("button[data-active]")
      .forEach((n) => delete n.dataset.active);
    audio.src = "";
  }

  audio.onended = clearActive;

  app.addEventListener("click", (e) => {
    e.preventDefault();

    if (app.contains(e.target) && e.target.nodeName === "BUTTON") {
      clearActive();
      const button = e.target;
      button.dataset.active = true;
      audio.src = button.dataset.audioSrc;
      audio.play();
    }
  });
})();
