# Authentication Routes:

- `/api/auth/signup`: Create a new user account.
- `/api/auth/login`: Authenticate and log in a user.
- `/api/auth/logout`: Log out the current user.

# Job Poster Routes:

- `/api/job-posters/:id`: Get details of a specific job poster's profile.
- `/api/job-posters/:id/jobs`: Get a list of jobs posted by a specific job poster.
- `/api/job-posters/:id/jobs/create`: Create a new job posting.
- `/api/job-posters/:id/jobs/:jobId/edit`: Update a specific job posting.
- `/api/job-posters/:id/jobs/:jobId/delete`: Delete a specific job posting.

# Job Seeker Routes:

- `/api/job-seekers/:id`: Get details of a specific job seeker's profile.
- `/api/job-seekers/:id/applications`: Get a list of job applications made by a specific job seeker.
- `/api/job-seekers/:id/search/jobs`: Search for jobs based on various parameters like job type, location, timings, pay, etc.
- `/api/job-seekers/:id/jobs/:jobId/apply`: Apply for a specific job posting.

# Job Routes:

- `/api/jobs/:jobId`: Get details of a specific job posting.
- `/api/jobs`: Get a list of available jobs based on user preferences.
