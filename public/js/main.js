var addCtnr = document.getElementById("add-container");
var addBtn = document.getElementById("add-button");
var addSbmt = document.getElementById("add-submit");
addBtn.addEventListener("click", function(evt) {
  addCtnr.style.display = "block";
});
addSbmt.addEventListener("click", function(evt) {
  addCtnr.style.display = "none";
});

var searchCtnr = document.getElementById("search-container");
var searchBtn = document.getElementById("search-button")
var titleSbmt = document.getElementById("title-submit");
var descriptionSbmt = document.getElementById("title-submit");
var usernameSbmt = document.getElementById("title-submit");
searchBtn.addEventListener("click", function(evt) {
  searchCtnr.style.display = "block";
});
titleSbmt.addEventListener("click", function(evt) {
  searchCtnr.style.display = "none";
});
descriptionSbmt.addEventListener("click", function(evt) {
  searchCtnr.style.display = "none";
});
usernameSbmt.addEventListener("click", function(evt) {
  searchCtnr.style.display = "none";
});
//alert("Got to here");
