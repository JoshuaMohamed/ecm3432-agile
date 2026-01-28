# Community Tourist Assistant - Kanban Board

## Backlog

### As a Tourist I want to see cover photos of places so that I can visualise list items.

  - tags: [Epic: Scrollable 'Places' page]

### As a Tourist I want to see a place's star rating out of 5 so that I can compare it to other options.

  - tags: [Epic: Scrollable 'Places' page]

### As a Tourist/Local I want to search the place list so that I can find details for a specific place.

  - tags: [Epic: Scrollable 'Places' page]

### As a Tourist I want to open a place list item so that I can see more details.

  - tags: [Epic: 'Place' details page]

### As a Tourist I want to see more photos when selecting a place so that I know what to expect.

  - tags: [Epic: 'Place' details page]

### As a Tourist/Local I want to upload photos so that I can contribute to a place's details.

  - tags: [Epic: Place' details page]

### As a Tourist I want to leave a star rating out of 5 so that I can raise the profile of a place.

  - tags: [Epic: ‘Place’ details page]

### As a Local I want to report incorrect place details so that they are flagged to an administrator.

  - tags: [Epic: ‘Place’ details page]

### As a Local I want to report unrelated place photos so that they are flagged to an administrator.

  - tags: [Epic: ‘Place’ details page]

### As a Tourist/Local I want to create an account of type Tourist/Local so that I can contribute to the platform.

  - tags: [Epic: ‘Sign Up or Log In’ page]

### As a Tourist/Local I want to log in so that I can contribute using my existing account.

  - tags: [Epic: ‘Sign Up or Log In’ page]

### As a Tourist/Local I want to log out so that I can use a different account.

  - tags: [Epic: ‘Sign Up or Log In’ page]

### As a Tourist/Local I want to see a list of my badges so that I know my contributions are worthwhile.

  - tags: [Epic: ‘My Badges’ page]

### As a Local I want to upload new places so that I can contribute to the platform.

  - tags: [Epic: ‘Upload’ page]

### As an Administrator I want to see a list of users, including account status, date joined, contributions and badges, so that I can manage user accounts.

  - tags: [Epic: ‘Manage Accounts’ admin page]

### As an Administrator I want to disable user accounts so that non-compliant users cannot contribute.

  - tags: [Epic: ‘Manage Accounts’ admin page]

### As an Administrator I want to assign badges so that users are incentivised to contribute.

  - tags: [Epic: ‘Manage Accounts’ admin page]

### As an Administrator I want to see a list of reported photos and place details so that I can take action or dismiss them.

  - tags: [Epic: ‘Manage Places’ admin page]

### As an Administrator I want to edit reported place details that I can correct them.

  - tags: [Epic: ‘Manage Places’ admin page]

### As an Administrator I want to remove photos so that I can address user reporting.

  - tags: [Epic: ‘Manage Places’ admin page]

## In Progress

### As a Tourist I want to see a list of local places so that I can identify somewhere to visit.

  - due: 2026-02-04
  - tags: [Sprint 1, Epic: Scrollable 'Places' page]
  - workload: Hard
  - defaultExpanded: true
  - steps:
      - [ ] 1. Create server with layered architecture
      - [ ] 2. Server GET /api/v1/getPlaces?post_code=EX44QJ returns JSON list with dummy data
      - [ ] 3, Design Scrollable 'Places' page
      - [ ] 4. Create web app
      - [ ] 5. Client sends getPlaces request to server with dummy postcode and received JSON list of places
      - [ ] 6. Client displays place names as a scrollable list of cards

## Done

