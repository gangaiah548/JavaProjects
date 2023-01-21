import { Component, OnInit } from '@angular/core';
import { Model, StylesManager } from 'survey-core';

// const SURVEY_ID = 1;

StylesManager.applyTheme('modern');

const surveyJson = {
  title: 'Your opinion matter for us',
  description: 'Employee pulse on happiness',
  pages: [
    {
      elements: [
        {
          name: 'Name',
          title: 'Employee name:',
          type: 'text',
        },
        {
          name: 'OverallSatisfaction',
          title:
            'How would you describe your happiness on job during recent time?',
          type: 'radiogroup',
          isRequired: true,
          choices: [
            'Very satisfied',
            'Satisfied',
            'Neutral',
            'Dissatisfied',
            'Very dissatisfied',
          ],
        },
        {
          name: 'Quality',
          title: 'Rate the following?',
          type: 'matrix',
          isRequired: true,
          columns: [
            {
              value: 1,
              text: 'Very poor',
            },
            {
              value: 2,
              text: 'Poor',
            },
            {
              value: 3,
              text: 'Average',
            },
            {
              value: 4,
              text: 'Good',
            },
            {
              value: 5,
              text: 'Excellent',
            },
          ],
          rows: [
            {
              value: 'Salary',
              text: 'Salary',
            },
            {
              value: 'OverallBenefits',
              text: 'All benefits',
            },
            {
              value: 'HealthBenefits',
              text: 'Health care & Insurance benefits',
            },
            {
              value: 'PhysicalWorkEnvironment',
              text: 'Physical work environment',
            },
            {
              value: 'TrainingOpportunities',
              text: 'Learning opportunities',
            },
            {
              value: 'WorkingTimeFlexibility',
              text: 'Working Life flexibility',
            },
          ],
        },
        {
          name: 'ValuedAtWork',
          title: 'Do you getting satisfied work?',
          type: 'radiogroup',
          isRequired: true,
          choices: ['Yes', 'No'],
        },
        {
          name: 'Explanation',
          title: 'If no please explain',
          visibleIf: "{ValuedAtWork}='No'",
          type: 'comment',
        },
        {
          name: 'Feedback',
          title: 'Any Suggestions or Feedback',
          type: 'comment',
        },
      ],
    },
  ],
};

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'My First Survey';
  surveyModel: Model;
  logSurveyResults(sender) {
    console.log(sender.data);
  }
  saveSurveyResults(sender) {
    const request = new XMLHttpRequest();
    const url = 'http://localhost:3001/surveys';
    request.open('POST', url);
    request.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    request.send(JSON.stringify(sender.data));
  }
  ngOnInit() {
    const survey = new Model(surveyJson);
    survey.onComplete.add(this.saveSurveyResults);
    this.surveyModel = survey;
  }
}
