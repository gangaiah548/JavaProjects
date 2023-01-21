const mongoose = require("mongoose");

const SurveySchema = mongoose.Schema(
  {
    Name: String,
    OverallSatisfaction: String,
    Quality: Array,
    ValuedAtWork: String,
    Explanation: String,
    Feedback: String,
  },
  {
    timestamps: true,
    strict: false,
  }
);

module.exports = mongoose.model("Survey", SurveySchema);
