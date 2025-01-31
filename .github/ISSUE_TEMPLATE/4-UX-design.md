name: UX design issue
description: Create an issue for delivering wireframes, mockups or other design solutions.
title: ""
labels: ["type/ux"]
body:
  - type: input
    id: background
    attributes:
      label: Background / Why we're doing this
      description: Describe the problem and background of the issue. This could include research insights that inform the design changes, unmet user needs, or other usability issues.
      placeholder: Add UI improvements to make Grafana Alerting alert creation easier based on usability test results.
    validations:
      required: true
  - type: dropdown
    attributes:
      label: Is there existing research for this?
      description: Please link research results or insights in the Background section if you have any. If no research was conducted, you might want to consider usability testing your design later.
      options:
      - label: Yes, I have linked it
      - label: No research yet
        required: true
  - type: textarea
    id: problems-or-tasks
    attributes:
      label: Problems or tasks
      description: Describe problems the new design should solve or tasks the user needs to complete.
      placeholder: 
      value: |
        - A problem we're trying to solve
        - A task the user needs to accomplish
        - …
    validations:
      required: false
- type: textarea
  attributes:
    label: Deliverables
    description: Add a checklist of deliverables here. You can later add links to each deliverable.
    value: |
        - Figma mockup
        - Miro board
        - …
