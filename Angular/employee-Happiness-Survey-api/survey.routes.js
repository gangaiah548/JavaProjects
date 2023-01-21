const surveyController = require("./survey.controller");
const router = require("express").Router();

//CRUD
router
  .post("/", surveyController.create)
  .get("/", surveyController.findAll);

module.exports = router;
