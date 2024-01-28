$(document).ready(function () {
  $("#qsLogoutBtn").click(function (e) {
    Cookies.remove("auth-session");
  });
});
