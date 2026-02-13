# Community Tourist Assistant - Kanban Board

## Backlog

### As a Tourist/Local I want to create an account of type Tourist/Local so that I can contribute to the platform.

  - tags: [Sprint 3, Epic: ‘Sign Up or Log In’ page]
  - workload: Normal

### As a Tourist/Local I want to log in so that I can contribute using my existing account.

  - tags: [Sprint 3, Epic: ‘Sign Up or Log In’ page]
  - workload: Normal

### As a Tourist/Local I want to log out so that I can use a different account.

  - tags: [Sprint 3, Epic: ‘Sign Up or Log In’ page]
  - workload: Easy

### As a Tourist I want to see a place's star rating out of 5 so that I can compare it to other options.

  - tags: [Sprint 3, Epic: Scrollable 'Places' page]
  - workload: Easy

### As a Tourist I want to leave a star rating out of 5 so that I can raise the profile of a place.

  - tags: [Sprint 3, Epic: ‘Place’ details page]
  - workload: Normal

### As a Local I want to upload new places so that I can contribute to the platform.

  - tags: [Sprint 4, Epic: ‘Upload’ page]
  - workload: Hard

### As a Tourist/Local I want to see a list of my badges so that I know my contributions are worthwhile.

  - tags: [Sprint 4, Epic: ‘My Badges’ page]
  - workload: Normal

  ### As an Administrator I want to assign badges so that users are incentivised to contribute.

  - tags: [Sprint 4, Epic: ‘Manage Accounts’ admin page]
  - workload: Normal

### As a Tourist/Local I want to search the place list so that I can find details for a specific place.

  - tags: [Out of Scope, Epic: Scrollable 'Places' page]

### As a Tourist I want to open a place list item so that I can see more details.

  - tags: [Out of Scope, Epic: 'Place' details page]

### As a Tourist I want to see more photos when selecting a place so that I know what to expect.

  - tags: [Out of Scope, Epic: 'Place' details page]

### As a Tourist/Local I want to upload photos so that I can contribute to a place's details.

  - tags: [Out of Scope, Epic: Place' details page]

### As a Local I want to report incorrect place details so that they are flagged to an administrator.

  - tags: [Out of Scope, Epic: ‘Place’ details page]

### As a Local I want to report unrelated place photos so that they are flagged to an administrator.

  - tags: [Out of Scope, Epic: ‘Place’ details page]

### As an Administrator I want to see a list of users, including account status, date joined, contributions and badges, so that I can manage user accounts.

  - tags: [Out of Scope, Epic: ‘Manage Accounts’ admin page]

### As an Administrator I want to disable user accounts so that non-compliant users cannot contribute.

  - tags: [Out of Scope, Epic: ‘Manage Accounts’ admin page]

### As an Administrator I want to see a list of reported photos and place details so that I can take action or dismiss them.

  - tags: [Out of Scope, Epic: ‘Manage Places’ admin page]

### As an Administrator I want to edit reported place details that I can correct them.

  - tags: [Out of Scope, Epic: ‘Manage Places’ admin page]

### As an Administrator I want to remove photos so that I can address user reporting.

  - tags: [Out of Scope, Epic: ‘Manage Places’ admin page]

## In Progress

### Tech Debt: 'Places' Web App Calls Get Places API

  - due: 2026-02-25
  - tags: [Sprint 2, Epic: Scrollable 'Places' page]
  - workload: Easy
  - defaultExpanded: true
  - steps:
      - [ ] 1. Client sends getPlaces request to server with dummy postcode and receives JSON list of places
      - [ ] 2. Client displays place names as a scrollable list of cards

### As a Tourist I want to see cover photos of places so that I can visualise list items.

  - due: 2026-02-25
  - tags: [Sprint 2, Epic: Scrollable 'Places' page]
  - workload: Normal
  - defaultExpanded: true
  - steps:
      - [ ] 1. Add cover filename to Places table
      - [ ] 2. Display photo alongside place summary in the 'Places' list

## Done

### Tech Debt: Get Places API Unit Tests

  - tags: [Sprint 2, Epic: Scrollable 'Places' page]
  - workload: Easy

### As a Tourist I want to see a list of local places so that I can identify somewhere to visit.

  - tags: [Sprint 1, Epic: Scrollable 'Places' page]
  - workload: Hard
  - steps:
      - [x] 1. Create server with layered architecture
      - [x] 2. Server GET "/getPlaces?postcode=EX4 4QJ" returns JSON list with dummy data
      - [x] 3, Design Scrollable 'Places' page
      - [x] 4. Create web app
      - [ ] 5. Client sends getPlaces request to server with dummy postcode and receives JSON list of places
      - [ ] 6. Client displays place names as a scrollable list of cards

